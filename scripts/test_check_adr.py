from pathlib import Path
import tempfile
import unittest

from check_adr import REQUIRED_SECTIONS, check


class CheckAdrTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.adr_dir = self.root / "doc" / "adr"
        self.adr_dir.mkdir(parents=True)
        self.path = self.adr_dir / "0001-example.md"
        self.summary = self.adr_dir / "current-decisions.md"
        sections = "".join(f"\n## {name}\n\n内容。\n" for name in REQUIRED_SECTIONS)
        self.path.write_text(
            "# ADR-0001: 判断\n\n- 決定日: 2026-10-25\n- 状態: 採用\n" + sections,
            encoding="utf-8",
        )
        self.summary.write_text(
            "| 判断 | 適用範囲 | 根拠 |\n"
            "| --- | --- | --- |\n"
            "| 判断 | 対象 | [ADR-0001](0001-example.md) |\n",
            encoding="utf-8",
        )

    def test_valid_active_decision(self):
        self.assertEqual(check(self.root), [])

    def test_active_decision_must_be_listed(self):
        self.summary.write_text("| 判断 | 適用範囲 | 根拠 |\n", encoding="utf-8")
        self.assertTrue(any("サマリにありません" in error for error in check(self.root)))

    def test_retired_decision_must_not_be_listed(self):
        content = self.path.read_text(encoding="utf-8")
        self.path.write_text(content.replace("状態: 採用", "状態: 撤回"), encoding="utf-8")
        self.assertTrue(any("採用中ではありません" in error for error in check(self.root)))

    def test_link_and_required_fields(self):
        content = self.path.read_text(encoding="utf-8")
        self.path.write_text(content.replace("## 理由\n", "## 理由なし\n"), encoding="utf-8")
        summary = self.summary.read_text(encoding="utf-8")
        self.summary.write_text(summary.replace("0001-example.md", "0001-other.md"), encoding="utf-8")
        errors = check(self.root)
        self.assertTrue(any("理由 がありません" in error for error in errors))
        self.assertTrue(any("リンク先が違います" in error for error in errors))


if __name__ == "__main__":
    unittest.main()
