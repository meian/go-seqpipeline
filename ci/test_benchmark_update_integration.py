"""Exercise update decisions against real Git snapshots and a simulated GitHub API."""

import json
from pathlib import Path
import subprocess
import tempfile
import unittest
from unittest.mock import patch

import check_benchmark_update as update

REPOSITORY = "meian/go-seqpipeline"
ISSUE_URL = "https://github.com/meian/go-seqpipeline/issues/100"


class BenchmarkUpdateIntegrationTests(unittest.TestCase):
    def setUp(self):
        self.directory = tempfile.TemporaryDirectory()
        self.addCleanup(self.directory.cleanup)
        self.root = Path(self.directory.name)
        self.git("init", "-b", "main")
        self.git("config", "user.name", "Benchmark test")
        self.git("config", "user.email", "benchmark@example.test")
        self.write("seq/map.go", "package seq\n\nfunc Map() {}\n")
        self.write("go.mod", "module example.test/seq\n\ngo 1.27.1\n")
        self.measured = self.commit()
        self.publish(self.measured)
        self.issues = []
        self.posts = []
        self.api_calls = []
        self.api_error = None
        self.runner = patch.object(update, "run", side_effect=self.command)
        self.runner.start()
        self.addCleanup(self.runner.stop)

    def git(self, *args):
        return subprocess.run(
            ["git", *args], cwd=self.root, check=True, capture_output=True, text=True,
        ).stdout.strip()

    def write(self, name, text):
        path = self.root / name
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text, encoding="utf-8")

    def commit(self):
        self.git("add", "-A")
        self.git("-c", "commit.gpgsign=false", "-c", "core.hooksPath=/dev/null", "commit", "-m", "Fixture")
        head = self.git("rev-parse", "HEAD")
        self.git("update-ref", "refs/remotes/origin/main", head)
        return head

    def publish(self, commit):
        self.write(update.ENVIRONMENT, (
            "# 実行環境\n\n"
            f"| 対象コード（ライブラリ・測定） | [{commit[:12]}]"
            f"(https://github.com/{REPOSITORY}/commit/{commit}) |\n"
        ))
        self.commit()

    def command(self, *args, input=None):
        if args[0] == "git":
            return subprocess.run(
                args, cwd=self.root, input=input, check=True, capture_output=True,
                text=True, encoding="utf-8",
            ).stdout
        self.assertEqual(args[0:2], ("gh", "api"))
        self.api_calls.append(args)
        if self.api_error:
            raise self.api_error
        if "GET" in args:
            return json.dumps([self.issues])
        self.assertIn("POST", args)
        self.posts.append(json.loads(input))
        self.issues.append({"html_url": ISSUE_URL, "body": self.posts[-1]["body"]})
        return json.dumps({"html_url": ISSUE_URL})

    def change_source(self):
        self.write("seq/map.go", "package seq\n\nfunc Map() { println(1) }\n")
        return self.commit()

    def test_unchanged_runtime_needs_no_issue_api_call(self):
        summary = update.check(REPOSITORY, False)
        self.assertIn("対象差分なし", summary)
        self.assertEqual(self.api_calls, [])

    def test_docs_tests_and_results_only_do_not_create_issue(self):
        self.write("README.md", "# Changed docs\n")
        self.write("seq/doc.go", "// Package seq documents the library.\npackage seq\n")
        self.write("seq/map_test.go", "package seq\n")
        self.write("benchmarks/map_test.go", "package benchmarks\n")
        self.write("benchmarks/results/unit.md", "# New report\n")
        self.commit()
        self.assertIn("対象差分なし", update.check(REPOSITORY, False))
        self.assertEqual(self.api_calls, [])

    def test_source_change_creates_issue_without_updating_report_or_git_refs(self):
        target = self.change_source()
        before = (self.root / update.ENVIRONMENT).read_text()
        summary = update.check(REPOSITORY, False)
        self.assertIn(ISSUE_URL, summary)
        self.assertEqual(len(self.posts), 1)
        body = self.posts[0]["body"]
        self.assertIn(self.measured, body)
        self.assertIn(target, body)
        self.assertIn("- `seq/map.go`", body)
        self.assertEqual((self.root / update.ENVIRONMENT).read_text(), before)
        self.assertEqual(self.git("rev-parse", "HEAD"), target)
        self.assertEqual(self.git("status", "--porcelain"), "")

    def test_new_package_is_detected(self):
        self.write("newpackage/implementation.go", "package newpackage\n")
        self.commit()
        update.check(REPOSITORY, False)
        self.assertIn("newpackage/implementation.go", self.posts[0]["body"])

    def test_go_version_change_creates_issue(self):
        self.write("go.mod", "module example.test/seq\n\ngo 1.28.0\n")
        self.commit()
        update.check(REPOSITORY, False)
        self.assertIn("- `go.mod`", self.posts[0]["body"])

    def test_deleted_source_creates_issue(self):
        (self.root / "seq/map.go").unlink()
        self.commit()
        update.check(REPOSITORY, False)
        self.assertIn("- `seq/map.go`", self.posts[0]["body"])

    def test_source_renamed_to_test_is_detected_as_deletion(self):
        (self.root / "seq/map.go").rename(self.root / "seq/map_test.go")
        self.commit()
        update.check(REPOSITORY, False)
        self.assertIn("- `seq/map.go`", self.posts[0]["body"])
        self.assertNotIn("- `seq/map_test.go`", self.posts[0]["body"])

    def test_reverted_implementation_needs_no_update(self):
        self.change_source()
        self.write("seq/map.go", "package seq\n\nfunc Map() {}\n")
        self.commit()
        self.assertIn("対象差分なし", update.check(REPOSITORY, False))
        self.assertEqual(self.api_calls, [])

    def test_repeat_check_skips_existing_issue(self):
        self.change_source()
        update.check(REPOSITORY, False)
        summary = update.check(REPOSITORY, False)
        self.assertIn("未完了の更新 Issue", summary)
        self.assertIn(ISSUE_URL, summary)
        self.assertEqual(len(self.posts), 1)

    def test_closing_issue_alone_does_not_advance_baseline(self):
        self.change_source()
        update.check(REPOSITORY, False)
        self.issues.clear()
        update.check(REPOSITORY, False)
        self.assertEqual(len(self.posts), 2)
        self.assertIn(self.measured, self.posts[1]["body"])

    def test_published_measurement_advances_baseline_and_only_new_changes_trigger(self):
        measured = self.change_source()
        update.check(REPOSITORY, False)
        self.issues.clear()
        self.publish(measured)
        summary = update.check(REPOSITORY, False)
        self.assertIn("対象差分なし", summary)
        self.assertIn(measured, summary)
        self.assertEqual(len(self.posts), 1)
        self.write("seq/filter.go", "package seq\n\nfunc Filter() {}\n")
        self.commit()
        update.check(REPOSITORY, False)
        self.assertEqual(len(self.posts), 2)
        self.assertIn(measured, self.posts[-1]["body"])
        self.assertIn("- `seq/filter.go`", self.posts[-1]["body"])
        self.assertNotIn("- `seq/map.go`", self.posts[-1]["body"])

    def test_dry_run_shows_body_without_posting_or_writing(self):
        self.change_source()
        summary = update.check(REPOSITORY, True)
        self.assertIn("dry-run", summary)
        self.assertIn(update.MARKER, summary)
        self.assertIn("seq/map.go", summary)
        self.assertEqual(len(self.api_calls), 1)
        self.assertEqual(self.posts, [])
        self.assertEqual(self.git("status", "--porcelain"), "")

    def test_initial_measurement_creates_only_one_issue(self):
        (self.root / update.ENVIRONMENT).unlink()
        self.commit()
        update.check(REPOSITORY, False)
        update.check(REPOSITORY, False)
        self.assertEqual(len(self.posts), 1)
        self.assertIn("初回の測定", self.posts[0]["body"])

    def test_invalid_report_fails_before_issue_api_access(self):
        self.write(update.ENVIRONMENT, "# broken report\n")
        self.commit()
        with self.assertRaises(ValueError):
            update.check(REPOSITORY, False)
        self.assertEqual(self.api_calls, [])

    def test_api_failure_propagates_without_creating_issue(self):
        self.change_source()
        self.api_error = subprocess.CalledProcessError(1, ["gh"], stderr="API unavailable")
        with self.assertRaises(subprocess.CalledProcessError):
            update.check(REPOSITORY, False)
        self.assertEqual(self.posts, [])

    def test_snapshot_comparison_does_not_require_shared_history(self):
        self.git("checkout", "--orphan", "squashed-main")
        self.commit()
        self.change_source()
        target = self.git("rev-parse", "refs/remotes/origin/main")
        ancestor = subprocess.run(
            ["git", "merge-base", "--is-ancestor", self.measured, target], cwd=self.root,
        )
        self.assertEqual(ancestor.returncode, 1)
        update.check(REPOSITORY, False)
        self.assertIn("- `seq/map.go`", self.posts[0]["body"])


if __name__ == "__main__":
    unittest.main()
