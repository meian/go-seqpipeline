#!/usr/bin/env python3
"""Check the structural consistency of ADRs and the current decisions summary."""

from datetime import date
from pathlib import Path
import re
import sys


ADR_NAME = re.compile(r"([0-9]{4})-[a-z0-9]+(?:-[a-z0-9]+)*\.md")
ADR_LINK = re.compile(r"\[ADR-([0-9]{4})\]\(([^)]+)\)")
REQUIRED_SECTIONS = (
    "背景",
    "決定",
    "理由",
    "採用しなかった選択肢",
    "影響",
    "関連 ADR",
    "根拠",
)
STATUSES = {"採用", "置換済み", "撤回"}


def check(root):
    adr_dir = root / "doc" / "adr"
    summary_path = adr_dir / "current-decisions.md"
    errors = []
    records = {}

    for path in sorted(adr_dir.glob("*.md")):
        if path.name in {"README.md", summary_path.name}:
            continue
        match = ADR_NAME.fullmatch(path.name)
        if not match:
            errors.append(f"{path.name}: ADR のファイル名が不正です")
            continue
        number = match.group(1)
        if number in records:
            errors.append(f"ADR-{number}: 番号が重複しています")
            continue

        content = path.read_text(encoding="utf-8")
        if not re.search(rf"(?m)^# ADR-{number}: .+$", content):
            errors.append(f"{path.name}: タイトルと番号が一致しません")
        decision_date = re.search(r"(?m)^- 決定日: (\d{4}-\d{2}-\d{2})$", content)
        if not decision_date:
            errors.append(f"{path.name}: 決定日がありません")
        else:
            try:
                date.fromisoformat(decision_date.group(1))
            except ValueError:
                errors.append(f"{path.name}: 決定日が不正です")
        status = re.search(r"(?m)^- 状態: (.+)$", content)
        if not status or status.group(1) not in STATUSES:
            errors.append(f"{path.name}: 状態が不正です")
        for section in REQUIRED_SECTIONS:
            if not re.search(rf"(?m)^## {re.escape(section)}$", content):
                errors.append(f"{path.name}: {section} がありません")
        records[number] = (path.name, status.group(1) if status else None)

    if not summary_path.is_file():
        return errors + [f"{summary_path}: サマリがありません"]

    listed = set()
    for line in summary_path.read_text(encoding="utf-8").splitlines():
        if not line.startswith("|"):
            continue
        cells = [cell.strip() for cell in line.strip("|").split("|")]
        if cells == ["判断", "適用範囲", "根拠"] or all(
            re.fullmatch(r":?-{3,}:?", cell) for cell in cells
        ):
            continue
        if len(cells) != 3 or not cells[0] or not cells[1]:
            errors.append("current-decisions.md: 判断と適用範囲が必要です")
            continue
        links = ADR_LINK.findall(cells[2])
        if not links:
            errors.append("current-decisions.md: 根拠となる ADR のリンクが必要です")
        for number, target in links:
            if number not in records:
                errors.append(f"current-decisions.md: ADR-{number} がありません")
            elif target != records[number][0]:
                errors.append(f"current-decisions.md: ADR-{number} のリンク先が違います")
            elif records[number][1] != "採用":
                errors.append(f"current-decisions.md: ADR-{number} は採用中ではありません")
            else:
                listed.add(number)

    for number, (_, status) in records.items():
        if status == "採用" and number not in listed:
            errors.append(f"ADR-{number}: 採用中の判断がサマリにありません")
    return errors


if __name__ == "__main__":
    errors = check(Path(__file__).resolve().parents[1])
    for error in errors:
        print(error, file=sys.stderr)
    sys.exit(bool(errors))
