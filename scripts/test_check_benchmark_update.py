import contextlib
import io
import json
import os
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch

import check_benchmark_update as update

REPOSITORY = "meian/go-seqpipeline"
COMMIT = "a" * 40
TARGET = "b" * 40


def report(commit=COMMIT, repository=REPOSITORY):
    return (
        "# 実行環境\n\n| 項目 | 値 |\n| --- | --- |\n"
        f"| 対象コード（ライブラリ・測定） | [{commit[:12]}]"
        f"(https://github.com/{repository}/commit/{commit}) |\n"
    )


class ReadBaselineTests(unittest.TestCase):
    def test_existing_report_uses_full_measured_commit(self):
        with patch.object(update, "run", side_effect=[update.ENVIRONMENT + "\0", report()]):
            self.assertEqual(update.read_baseline(REPOSITORY, TARGET), COMMIT)

    def test_missing_results_are_initial_measurement(self):
        with patch.object(update, "run", return_value="benchmarks/map_test.go\0"):
            self.assertIsNone(update.read_baseline(REPOSITORY, TARGET))

    def test_missing_environment_with_existing_results_fails(self):
        for path in ("benchmarks/REPORT.md", "benchmarks/results/unit.md"):
            with self.subTest(path=path):
                with patch.object(update, "run", return_value=path + "\0"):
                    with self.assertRaisesRegex(ValueError, "実行環境レポートがありません"):
                        update.read_baseline(REPOSITORY, TARGET)

    def test_invalid_ambiguous_or_foreign_commit_fails(self):
        reports = (
            "# no commit\n", report() * 2, report(commit="a" * 12),
            report(repository="other/repository"), report().replace("[aaaaaaaaaaaa]", "[bbbbbbbbbbbb]"),
        )
        for text in reports:
            with self.subTest(report=text):
                with patch.object(update, "run", side_effect=[update.ENVIRONMENT + "\0", text]):
                    with self.assertRaisesRegex(ValueError, "測定対象コミットが不正"):
                        update.read_baseline(REPOSITORY, TARGET)

    def test_git_failure_is_not_initial_measurement(self):
        error = subprocess.CalledProcessError(1, ["git"], stderr="cannot read tree")
        with patch.object(update, "run", side_effect=error):
            with self.assertRaises(subprocess.CalledProcessError):
                update.read_baseline(REPOSITORY, TARGET)


class EnsureCommitTests(unittest.TestCase):
    def test_existing_object_needs_no_fetch(self):
        with patch.object(update, "run", return_value="") as run:
            update.ensure_commit(COMMIT)
        run.assert_called_once_with("git", "cat-file", "-e", f"{COMMIT}^{{commit}}")

    def test_missing_object_is_fetched_and_verified(self):
        error = subprocess.CalledProcessError(1, ["git"], stderr="missing object")
        with patch.object(update, "run", side_effect=[error, "", ""]) as run:
            update.ensure_commit(COMMIT)
        self.assertEqual(run.call_args_list[1].args, ("git", "fetch", "--no-tags", "origin", COMMIT))
        self.assertEqual(run.call_args_list[2], run.call_args_list[0])

    def test_fetch_failure_is_not_no_difference(self):
        error = subprocess.CalledProcessError(1, ["git"], stderr="fetch failed")
        with patch.object(update, "run", side_effect=error):
            with self.assertRaises(subprocess.CalledProcessError):
                update.ensure_commit(COMMIT)


class IsRuntimePathTests(unittest.TestCase):
    def test_library_and_dependency_sources_and_configs_are_included(self):
        for path in (
            "seq/map.go", "newpackage/implementation.go", "root.go", "vendor/example/dep.go",
            "go.mod", "go.sum", "go.work", "go.work.sum", ".go-version", ".tool-versions",
        ):
            with self.subTest(path=path):
                self.assertTrue(update.is_runtime_path(path))

    def test_docs_tests_benchmarks_and_automation_are_excluded(self):
        for path in (
            "README.md", "doc.go", "seq/doc.go", "seq/map_test.go", "seq/testdata/input.go",
            "benchmarks/map_test.go", "benchmarks/helper.go", "benchmarks/results/environment.md",
            "benchmarks/measure.sh", "scripts/check_benchmark_update.py", "doc/sample.go",
            ".github/workflows/ci.yml", "testdata/example.go",
        ):
            with self.subTest(path=path):
                self.assertFalse(update.is_runtime_path(path))


class ChangedFilesTests(unittest.TestCase):
    def test_selects_runtime_paths_and_keeps_spaces(self):
        paths = "README.md\0seq/z.go\0seq/a file.go\0seq/a_test.go\0go.mod\0"
        with patch.object(update, "run", return_value=paths) as run:
            result = update.changed_files(COMMIT, TARGET)
        self.assertEqual(result, ["go.mod", "seq/a file.go", "seq/z.go"])
        self.assertIn("--no-renames", run.call_args.args)
        self.assertNotIn(f"{COMMIT}...{TARGET}", run.call_args.args)

    def test_git_failure_propagates(self):
        error = subprocess.CalledProcessError(1, ["git"], stderr="diff failed")
        with patch.object(update, "run", side_effect=error):
            with self.assertRaises(subprocess.CalledProcessError):
                update.changed_files(COMMIT, TARGET)


class OpenUpdateIssueTests(unittest.TestCase):
    def test_finds_marked_issue_on_later_page_without_matching_planning_issue_or_pr(self):
        pages = [
            [{"number": 33, "body": "benchmark update planning"}, {"body": None},
             {"body": update.MARKER, "pull_request": {}}],
            [{"body": update.MARKER + "\nbody", "html_url": "https://github.com/example/issues/50"}],
        ]
        with patch.object(update, "run", return_value=json.dumps(pages)) as run:
            self.assertEqual(update.open_update_issue(REPOSITORY), "https://github.com/example/issues/50")
        self.assertIn("--paginate", run.call_args.args)
        self.assertIn("--slurp", run.call_args.args)
        self.assertIn("state=open", run.call_args.args[-1])

    def test_unmarked_or_quoted_marker_is_not_an_update_issue(self):
        pages = [[{"body": "quoted: " + update.MARKER}]]
        with patch.object(update, "run", return_value=json.dumps(pages)):
            self.assertIsNone(update.open_update_issue(REPOSITORY))

    def test_api_failure_is_not_an_empty_issue_list(self):
        error = subprocess.CalledProcessError(1, ["gh"], stderr="API unavailable")
        with patch.object(update, "run", side_effect=error):
            with self.assertRaises(subprocess.CalledProcessError):
                update.open_update_issue(REPOSITORY)

    def test_invalid_page_structure_is_not_an_empty_issue_list(self):
        for response in ("{}", "[]", '[{"body": null}]', "not JSON"):
            with self.subTest(response=response):
                with patch.object(update, "run", return_value=response):
                    with self.assertRaises(ValueError):
                        update.open_update_issue(REPOSITORY)


class IssueBodyTests(unittest.TestCase):
    def test_existing_snapshot_includes_comparison_and_completion_instructions(self):
        body = update.issue_body(REPOSITORY, COMMIT, TARGET, ["seq/map.go", "go.mod"])
        self.assertTrue(body.startswith(update.MARKER + "\n"))
        self.assertIn(f"/compare/{COMMIT}..{TARGET}", body)
        self.assertIn("- `seq/map.go`", body)
        self.assertIn("- `go.mod`", body)
        self.assertIn("/issues/32", body)
        self.assertIn(f"/blob/{TARGET}/benchmarks/REPORT.md", body)
        self.assertIn("測定日時", body)
        self.assertIn("main に反映", body)
        self.assertIn("作成・クローズだけでは比較基準は進まない", body)

    def test_initial_measurement_has_no_missing_report_or_comparison_link(self):
        body = update.issue_body(REPOSITORY, None, TARGET, [])
        self.assertIn("初回の測定", body)
        self.assertIn(TARGET, body)
        self.assertNotIn("/compare/", body)
        self.assertNotIn("/benchmarks/REPORT.md)", body)


class CreateIssueTests(unittest.TestCase):
    def test_posts_literal_body_as_json(self):
        body = "Japanese 日本語\n`literal` $(not a command)\n"
        with patch.object(update, "run", return_value='{"html_url":"https://example.test/issue"}') as run:
            self.assertEqual(update.create_issue(REPOSITORY, body), "https://example.test/issue")
        payload = json.loads(run.call_args.kwargs["input"])
        self.assertEqual(payload["body"], body)
        self.assertTrue(payload["title"].startswith("[定期更新]"))
        self.assertEqual(run.call_args.args[-2:], ("--input", "-"))


class CheckTests(unittest.TestCase):
    def setUp(self):
        self.patches = {}
        for name, result in (
            ("run", TARGET + "\n"), ("read_baseline", COMMIT), ("ensure_commit", None),
            ("changed_files", ["seq/map.go"]), ("open_update_issue", None),
            ("create_issue", "https://example.test/issue"),
        ):
            patcher = patch.object(update, name, return_value=result)
            self.patches[name] = patcher.start()
            self.addCleanup(patcher.stop)

    def test_no_changes_skip_issue_lookup_and_creation(self):
        self.patches["changed_files"].return_value = []
        self.assertIn("対象差分なし", update.check(REPOSITORY, False))
        self.patches["open_update_issue"].assert_not_called()
        self.patches["create_issue"].assert_not_called()

    def test_existing_issue_prevents_creation(self):
        self.patches["open_update_issue"].return_value = "https://example.test/existing"
        self.assertIn("https://example.test/existing", update.check(REPOSITORY, False))
        self.patches["create_issue"].assert_not_called()

    def test_dry_run_previews_current_and_measured_commits_without_posting(self):
        summary = update.check(REPOSITORY, True)
        self.assertIn("dry-run", summary)
        self.assertIn(COMMIT, summary)
        self.assertIn(TARGET, summary)
        self.assertIn("seq/map.go", summary)
        self.patches["create_issue"].assert_not_called()

    def test_initial_measurement_skips_diff_but_checks_duplicates(self):
        self.patches["read_baseline"].return_value = None
        self.assertIn("初回の測定", update.check(REPOSITORY, True))
        self.patches["ensure_commit"].assert_not_called()
        self.patches["changed_files"].assert_not_called()
        self.patches["open_update_issue"].assert_called_once_with(REPOSITORY)
        self.patches["create_issue"].assert_not_called()

    def test_changes_create_issue_using_captured_snapshot(self):
        self.assertIn("https://example.test/issue", update.check(REPOSITORY, False))
        repository, body = self.patches["create_issue"].call_args.args
        self.assertEqual(repository, REPOSITORY)
        self.assertIn(COMMIT, body)
        self.assertIn(TARGET, body)
        self.patches["read_baseline"].assert_called_once_with(REPOSITORY, TARGET)
        self.patches["changed_files"].assert_called_once_with(COMMIT, TARGET)


class MainTests(unittest.TestCase):
    def test_dry_run_and_summary(self):
        with tempfile.TemporaryDirectory() as directory:
            summary = Path(directory) / "summary.md"
            with patch.dict(os.environ, {"GITHUB_STEP_SUMMARY": str(summary)}):
                with patch.object(update, "check", return_value="dry-run summary\n") as check:
                    with contextlib.redirect_stdout(io.StringIO()):
                        self.assertEqual(update.main(["--repo", REPOSITORY, "--dry-run"]), 0)
            check.assert_called_once_with(REPOSITORY, True)
            self.assertEqual(summary.read_text(), "dry-run summary\n")

    def test_failures_have_nonzero_exit_and_summary(self):
        for error in (
            ValueError("invalid baseline"),
            subprocess.CalledProcessError(1, ["gh"], stderr="API failed"),
        ):
            with self.subTest(error=error), tempfile.TemporaryDirectory() as directory:
                summary = Path(directory) / "summary.md"
                with patch.dict(os.environ, {"GITHUB_STEP_SUMMARY": str(summary)}):
                    with patch.object(update, "check", side_effect=error):
                        with contextlib.redirect_stderr(io.StringIO()):
                            self.assertEqual(update.main(["--repo", REPOSITORY]), 1)
                self.assertIn("チェック失敗", summary.read_text())

    def test_invalid_repository_is_rejected_before_git_or_api_access(self):
        with patch.object(update, "check") as check, contextlib.redirect_stderr(io.StringIO()):
            with self.assertRaises(SystemExit) as error:
                update.main(["--repo", "../invalid/repository"])
        self.assertEqual(error.exception.code, 2)
        check.assert_not_called()


if __name__ == "__main__":
    unittest.main()
