import base64
import contextlib
import io
import subprocess
from copy import deepcopy
import unittest
from unittest.mock import patch

import yaml

import dependabot_automerge as automerge


REPO = "owner/repo"
ROOT = f"repos/{REPO}"
HEAD = "a" * 40
MAIN = "b" * 40
PARENT = "c" * 40


def commit(dependencies=None):
    if dependencies is None:
        dependencies = [{"dependency-name": "packaging", "dependency-version": "26.1",
                         "update-type": "version-update:semver-minor"}]
    trailer = yaml.safe_dump({"updated-dependencies": dependencies}, sort_keys=False)
    return {
        "sha": HEAD, "author": {"login": automerge.BOT},
        "parents": [{"sha": PARENT}],
        "commit": {"verification": {"verified": True},
                   "message": f"Bump dependencies\n\n---\n{trailer}...\n"},
    }


def pull_request():
    return {
        "number": 7, "state": "open", "draft": False,
        "user": {"login": automerge.BOT}, "base": {"ref": "main"},
        "head": {"sha": HEAD, "repo": {"full_name": REPO}},
        "changed_files": 1, "mergeable": True, "mergeable_state": "clean",
    }


def workflow(run_id=1, event="pull_request", path=".github/workflows/ci.yml"):
    return {"id": run_id, "workflow_id": 10, "event": event, "path": path,
            "check_suite_id": run_id * 100, "status": "completed", "conclusion": "success"}


def check(name, suite=100):
    return {"name": name, "app": {"id": automerge.ACTION_APP},
            "check_suite": {"id": suite}, "status": "completed", "conclusion": "success"}


class MetadataTests(unittest.TestCase):
    def test_signed_metadata_and_missing_kind(self):
        self.assertEqual(automerge.metadata(commit()), {"packaging": "26.1"})
        dependency = {"dependency-name": "packaging", "dependency-version": "26.1"}
        self.assertEqual(automerge.metadata(commit([dependency])), {"packaging": "26.1"})

    def test_untrusted_or_unclassifiable_metadata_is_rejected(self):
        unsigned = commit()
        unsigned["commit"]["verification"]["verified"] = False
        human = commit()
        human["author"]["login"] = "human"
        merge = commit()
        merge["parents"].append({"sha": MAIN})
        absent = commit()
        absent["commit"]["message"] = "No metadata"
        repeated = commit()
        repeated["commit"]["message"] *= 2
        cases = {"unsigned": unsigned, "human": human, "merge": merge,
                 "absent": absent, "multiple trailers": repeated, "empty": commit([])}
        for kind in ("version-update:semver-major", "unknown", ""):
            cases[kind] = commit([{"dependency-name": "packaging", "dependency-version": "26.1",
                                   "update-type": kind}])
        duplicate = [{"dependency-name": "packaging", "dependency-version": "26.1"}] * 2
        cases["duplicate dependency"] = commit(duplicate)
        for name, value in cases.items():
            with self.subTest(name=name), self.assertRaises(ValueError):
                automerge.metadata(value)

    def test_minor_and_patch_boundaries(self):
        cases = [
            ("1.2.3", "1.2.4", True), ("v1.2.3", "v1.3.0", True),
            ("1.2", "1.2.1", True), ("0.1.0", "0.2.0", True),
            ("1.9.9", "2.0.0", False), ("1.2", "1.2.0", False),
            ("1.2.3", "1.2.3", False), ("1.3.0", "1.2.9", False),
            ("v1", "v1.1", False), ("1.2.3", "1.3.0-rc.1", False),
            ("main", "1.3.0", False), ("1.2.3", "1.2.4+build", False),
        ]
        for before, after, accepted in cases:
            with self.subTest(before=before, after=after):
                if accepted:
                    self.assertIsNone(automerge.minor_or_patch(before, after))
                else:
                    with self.assertRaises(ValueError):
                        automerge.minor_or_patch(before, after)


class FileTests(unittest.TestCase):
    def test_file_decoding_and_unexpected_content_types(self):
        data = {"type": "file", "encoding": "base64",
                "content": base64.b64encode(b"packaging==26.1\n").decode()}
        with patch.object(automerge, "api", return_value=data) as api:
            self.assertEqual(automerge.read_file(REPO, "ci/a file.txt", HEAD), "packaging==26.1\n")
            api.assert_called_with(f"{ROOT}/contents/ci/a%20file.txt?ref={HEAD}")
            for field, value in (("type", "dir"), ("encoding", "none")):
                invalid = dict(data, **{field: value})
                api.return_value = invalid
                with self.subTest(field=field), self.assertRaises(ValueError):
                    automerge.read_file(REPO, "ci/a file.txt", HEAD)


class ClassificationTests(unittest.TestCase):
    def classify(self, contents, dependencies, *, status="modified", changed_files=None, tags=None):
        pr = pull_request()
        pr["changed_files"] = len(contents) if changed_files is None else changed_files
        files = [{"filename": path, "status": status} for path in contents]
        signed = commit([{"dependency-name": name, "dependency-version": version}
                         for name, version in dependencies.items()])
        responses = {f"{ROOT}/pulls/7/commits?per_page=100": [signed],
                     f"{ROOT}/pulls/7/files?per_page=100": files}
        def api_response(path):
            if path == f"{ROOT}/commits/{HEAD}":
                return signed
            return {"sha": tags[path]}
        def read_file(repo, path, ref):
            self.assertEqual(repo, REPO)
            return contents[path][0 if ref == PARENT else 1]
        with patch.object(automerge, "items", side_effect=responses.__getitem__), \
                patch.object(automerge, "api", side_effect=api_response), \
                patch.object(automerge, "read_file", side_effect=read_file):
            return automerge.classify(REPO, pr)

    def test_grouped_requirements_and_repeated_action_uses(self):
        contents = {"ci/scripts/requirements.txt": (
            "packaging==26.0\nPyYAML==6.0.2\n", "packaging==26.1\nPyYAML==6.0.3\n")}
        self.assertEqual(self.classify(contents, {"packaging": "26.1", "PyYAML": "6.0.3"}),
                         ["PyYAML", "packaging"])
        line = "    - uses: actions/checkout@v6.0.0\n"
        updated = "    - uses: actions/checkout@v6.0.1\n"
        contents = {f".github/workflows/{name}.yml": (line, updated) for name in ("ci", "lint")}
        self.assertEqual(self.classify(contents, {"actions/checkout": "6.0.1"}),
                         ["actions/checkout"])

    def test_python_dependency_names_are_normalized_and_unambiguous(self):
        for filename, dependency in (("PyYAML", "pyyaml"), ("pack_name", "pack-name"),
                                     ("pack...name", "pack-name")):
            contents = {"ci/scripts/requirements.txt":
                        (f"{filename}==1.0.0\n", f"{filename}==1.0.1\n")}
            with self.subTest(filename=filename, dependency=dependency):
                self.assertEqual(self.classify(contents, {dependency: "1.0.1"}), [dependency])
        with self.assertRaisesRegex(ValueError, "一意"):
            self.classify({"ci/scripts/requirements.txt": ("pack_name==1.0.0\n", "pack_name==1.0.1\n")},
                          {"pack-name": "1.0.1", "pack_name": "1.0.1"})

    def test_non_dependency_changes_and_incomplete_groups_are_rejected(self):
        workflow_path = ".github/workflows/ci.yml"
        action = "    - uses: actions/setup-go@v6.0.0\n"
        action_update = "    - uses: actions/setup-go@v6.0.1\n"
        cases = [
            ("go.mod", "go 1.27.2\n", "go 1.27.3\n", {"go": "1.27.3"}),
            (workflow_path, action + "      go-version: 1.27.2\n",
             action_update + "      go-version: 1.27.3\n", {"actions/setup-go": "6.0.1"}),
            (workflow_path, action + "      token: old\n", action_update + "      token: new\n",
             {"actions/setup-go": "6.0.1"}),
            (workflow_path, action, action_update + "    - run: echo added\n",
             {"actions/setup-go": "6.0.1"}),
            (workflow_path, action + "    - run: echo removed\n", action_update,
             {"actions/setup-go": "6.0.1"}),
            (workflow_path, action, "      uses: actions/setup-go@v6.0.1\n",
             {"actions/setup-go": "6.0.1"}),
            ("README.md", "old\n", "new\n", {"packaging": "26.1"}),
            ("ci/scripts/requirements.txt", "packaging==26.0\n", "packaging==26.1\n",
             {"packaging": "26.2"}),
            ("ci/scripts/requirements.txt", "packaging==26.0\n", "packaging==26.1\n",
             {"packaging": "26.1", "PyYAML": "6.0.3"}),
            ("ci/scripts/requirements.txt", "packaging==26.0\n", "packaging==27.0\n",
             {"packaging": "27.0"}),
        ]
        for path, before, after, dependencies in cases:
            with self.subTest(path=path, after=after, dependencies=dependencies), self.assertRaises(ValueError):
                self.classify({path: (before, after)}, dependencies)
        contents = {workflow_path: (action, action_update)}
        for status in ("added", "removed", "renamed"):
            with self.subTest(status=status), self.assertRaises(ValueError):
                self.classify(contents, {"actions/setup-go": "6.0.1"}, status=status)
        with self.assertRaises(ValueError):
            self.classify(contents, {"actions/setup-go": "6.0.1"}, changed_files=2)

    def test_sha_pins_require_matching_stable_version_tags(self):
        path = ".github/workflows/ci.yml"
        before, after = "d" * 40, "e" * 40
        contents = {path: (f"- uses: actions/checkout@{before} # v6.0.0\n",
                           f"- uses: actions/checkout@{after} # v6.0.1\n")}
        tags = {"repos/actions/checkout/commits/v6.0.0": before,
                "repos/actions/checkout/commits/v6.0.1": after}
        for version in (after, "6.0.1"):
            with self.subTest(metadata_version=version):
                self.assertEqual(self.classify(contents, {"actions/checkout": version}, tags=tags),
                                 ["actions/checkout"])
        with self.assertRaisesRegex(ValueError, "metadata"):
            self.classify(contents, {"actions/checkout": "6.0.2"}, tags=tags)
        for old_line, new_line, tag_values in (
            (f"- uses: actions/checkout@{before}\n", contents[path][1], tags),
            (contents[path][0], "- uses: actions/checkout@v6.0.1\n", tags),
            (contents[path][0], contents[path][1], {key: "f" * 40 for key in tags}),
        ):
            with self.subTest(old=old_line, new=new_line, tags=tag_values), self.assertRaises(ValueError):
                self.classify({path: (old_line, new_line)}, {"actions/checkout": after}, tags=tag_values)

    def test_extra_commits_are_rejected_before_reading_files(self):
        with patch.object(automerge, "items", return_value=[commit(), commit()]), \
                patch.object(automerge, "api") as api:
            with self.assertRaises(ValueError):
                automerge.classify(REPO, pull_request())
            api.assert_not_called()


class CheckTests(unittest.TestCase):
    def verify(self, *, runs=None, checks=None, statuses=None, total=None):
        runs = [workflow()] if runs is None else runs
        checks = [check(name) for name in automerge.REQUIRED] if checks is None else checks
        statuses = [] if statuses is None else statuses
        if total is None:
            total = len({status["context"] for status in statuses})
        responses = {f"{ROOT}/actions/runs?head_sha={HEAD}&per_page=100": runs,
                     f"{ROOT}/commits/{HEAD}/check-runs?filter=latest&per_page=100": checks,
                     f"{ROOT}/commits/{HEAD}/statuses?per_page=100": statuses}
        with patch.object(automerge, "items", side_effect=lambda path, key=None: responses[path]), \
                patch.object(automerge, "api", return_value={"total_count": total}):
            automerge.checks_pass(REPO, HEAD, {(name, automerge.ACTION_APP) for name in automerge.REQUIRED})

    def test_all_required_checks_and_latest_status_must_succeed(self):
        self.verify(statuses=[{"context": "external", "state": "success"},
                              {"context": "external", "state": "failure"}])
        failures = [([], None, None), ([workflow()], [], None)]
        for conclusion in ("failure", "cancelled", "skipped", "neutral", None):
            checks = [check(name) for name in automerge.REQUIRED]
            checks[0]["conclusion"] = conclusion
            failures.append(([workflow()], checks, None))
        pending = check("test")
        pending["status"] = "in_progress"
        failures.append(([workflow()], [pending], None))
        extra = check("optional")
        extra["conclusion"] = "failure"
        failures.append(([workflow()], [check(name) for name in automerge.REQUIRED] + [extra], None))
        failures.append(([workflow()], None, [{"context": "external", "state": "pending"}]))
        failures.append(([workflow()], None, [{"context": "external", "state": "failure"},
                                             {"context": "external", "state": "success"}]))
        for runs, checks, statuses in failures:
            with self.subTest(runs=runs, checks=checks, statuses=statuses), self.assertRaises(ValueError):
                self.verify(runs=runs, checks=checks, statuses=statuses)
        with self.assertRaises(ValueError):
            self.verify(statuses=[{"context": "external", "state": "success"}], total=2)

    def test_workflow_pending_failure_and_cancellation_are_not_success(self):
        for status, conclusion in (("in_progress", None), ("completed", "failure"),
                                   ("completed", "cancelled")):
            run = workflow()
            run.update(status=status, conclusion=conclusion)
            with self.subTest(status=status, conclusion=conclusion), self.assertRaises(ValueError):
                self.verify(runs=[run])

    def test_push_lint_skip_requires_successful_pr_lint(self):
        checks = [check(name) for name in automerge.REQUIRED]
        skipped = check("Static checks", suite=200)
        skipped["conclusion"] = "skipped"
        runs = [workflow(), workflow(2, event="push")]
        self.verify(runs=runs, checks=checks + [skipped])
        with self.assertRaises(ValueError):
            self.verify(runs=runs, checks=[entry for entry in checks if entry["name"] != "Static checks"] + [skipped])
        skipped["name"] = "test"
        with self.assertRaises(ValueError):
            self.verify(runs=runs, checks=checks + [skipped])

    def test_old_run_success_cannot_satisfy_new_run_and_retry_can_recover(self):
        old, new = workflow(), workflow(2)
        old["conclusion"] = "failure"
        old_check = check("test")
        old_check["conclusion"] = "failure"
        self.verify(runs=[old, new], checks=[old_check] + [check(name, suite=200) for name in automerge.REQUIRED])
        with self.assertRaises(ValueError):
            self.verify(runs=[old, new], checks=[check(name) for name in automerge.REQUIRED])
        new["status"], new["conclusion"] = "queued", None
        with self.assertRaises(ValueError):
            self.verify(runs=[workflow(), new])
        own = workflow(3, path=automerge.WORKFLOW)
        own["status"], own["conclusion"] = "in_progress", None
        self.verify(runs=[workflow(), own], checks=[check(name) for name in automerge.REQUIRED] + [check("automerge", suite=300)])

    def test_unmapped_actions_suites_are_not_silently_ignored(self):
        for conclusion, status in (("success", "completed"), ("failure", "completed"),
                                   (None, "in_progress")):
            unknown = check("other workflow", suite=999)
            unknown.update(conclusion=conclusion, status=status)
            checks = [check(name) for name in automerge.REQUIRED] + [unknown]
            with self.subTest(conclusion=conclusion), self.assertRaisesRegex(ValueError, "workflow"):
                self.verify(checks=checks)

    def test_wrong_app_cannot_satisfy_required_context(self):
        checks = [check(name) for name in automerge.REQUIRED]
        checks[0]["app"]["id"] = 123
        with self.assertRaises(ValueError):
            self.verify(checks=checks)


class ProtectionTests(unittest.TestCase):
    def test_strict_required_sources_without_private_ruleset_fields(self):
        params = {"strict_required_status_checks_policy": True,
                  "required_status_checks": [{"context": name, "integration_id": automerge.ACTION_APP}
                                              for name in automerge.REQUIRED]}
        rule = {"type": "required_status_checks", "parameters": params, "ruleset_id": 9}
        with patch.object(automerge, "items", return_value=[rule]) as items, \
                patch.object(automerge, "api", side_effect=AssertionError("private ruleset lookup")):
            self.assertEqual(automerge.required_checks(REPO),
                             {(name, automerge.ACTION_APP) for name in automerge.REQUIRED})
            items.assert_called_with(f"{ROOT}/rules/branches/main?per_page=100")
            params["strict_required_status_checks_policy"] = False
            with self.assertRaises(ValueError):
                automerge.required_checks(REPO)
            params["strict_required_status_checks_policy"] = True
            params["required_status_checks"].pop()
            with self.assertRaises(ValueError):
                automerge.required_checks(REPO)


class ProcessTests(unittest.TestCase):
    def run_process(self, *, pr=None, fresh=None, main=MAIN, fresh_main=MAIN,
                    behind=False, comments=None, dry_run=False, merge_results=None, check_error=None):
        pr = pull_request() if pr is None else pr
        fresh = deepcopy(pr) if fresh is None else fresh
        prs, mains = iter([pr, fresh]), iter([main, fresh_main])
        merge_results = iter(merge_results or [{"status": "merged", "details": {"sha": "merged-sha"}}])
        def response(path, method="GET", data=None):
            if path == f"{ROOT}/pulls/7" and method == "GET":
                return next(prs)
            if path == f"{ROOT}/git/ref/heads/main":
                return {"object": {"sha": next(mains)}}
            if path == f"{ROOT}/compare/{main}...{HEAD}":
                return {"merge_base_commit": {"sha": PARENT if behind else main}}
            if method == "POST":
                return {"id": 1}
            if path == f"{ROOT}/pulls/7/merge-async" and method == "PUT":
                return next(merge_results)
            if path == f"{ROOT}/pulls/7/merge-async/request-id" and method == "GET":
                return next(merge_results)
            raise AssertionError(f"Unexpected API: {method} {path}")
        with patch.object(automerge, "api", side_effect=response) as api, \
                patch.object(automerge, "items", return_value=comments or []), \
                patch.object(automerge, "classify", return_value=["packaging"]), \
                patch.object(automerge, "required_checks", return_value={("test", automerge.ACTION_APP)}), \
                patch.object(automerge, "checks_pass", side_effect=check_error) as checks:
            try:
                result = automerge.process(REPO, 7, dry_run)
            finally:
                self.last_api, self.last_checks = api, checks
        return result

    def test_merge_uses_validated_sha_and_checks_again_before_write(self):
        self.assertIn("マージしました", self.run_process())
        self.last_api.assert_called_with(f"{ROOT}/pulls/7/merge-async", "PUT",
                                         {"sha": HEAD, "merge_method": "squash",
                                          "merge_action": "direct_merge", "bypass_rules": False})
        self.assertEqual(self.last_checks.call_count, 2)
        self.assertIn("マージ予定", self.run_process(dry_run=True))
        self.assertFalse(any(call.args[1:] for call in self.last_api.call_args_list))

    def test_async_merge_reports_only_confirmed_completion_and_polls_once(self):
        pending = {"status": "pending", "details": {"uuid": "request-id"}}
        merged = {"status": "merged", "details": {"sha": "merged-sha"}}
        self.assertIn("マージしました: merged-sha", self.run_process(merge_results=[pending, merged]))
        self.last_api.assert_called_with(f"{ROOT}/pulls/7/merge-async/request-id")
        result = self.run_process(merge_results=[pending, pending])
        self.assertIn("完了未確認", result)
        polls = [call for call in self.last_api.call_args_list
                 if call.args[0] == f"{ROOT}/pulls/7/merge-async/request-id"]
        self.assertEqual(len(polls), 1)
        for status in ("failed", "enqueued"):
            result = {"status": status, "details": {"message": "Not merged"}}
            for responses in ([result], [pending, result]):
                with self.subTest(status=status, responses=responses), self.assertRaises(ValueError):
                    self.run_process(merge_results=responses)

    def test_failed_checks_prevent_merge(self):
        with self.assertRaises(ValueError):
            self.run_process(check_error=ValueError("CI failed"))
        self.assertFalse(any(call.args[1:] for call in self.last_api.call_args_list))

    def test_non_dependabot_draft_closed_fork_and_wrong_base_are_rejected(self):
        cases = []
        for field, value in (("state", "closed"), ("draft", True)):
            pr = pull_request()
            pr[field] = value
            cases.append(pr)
        for field, key, value in (("user", "login", "human"), ("base", "ref", "release")):
            pr = pull_request()
            pr[field][key] = value
            cases.append(pr)
        pr = pull_request()
        pr["head"]["repo"]["full_name"] = "other/repo"
        cases.append(pr)
        for pr in cases:
            with self.subTest(pr=pr), self.assertRaises(ValueError):
                self.run_process(pr=pr)
            self.last_checks.assert_not_called()

    def test_head_main_or_pr_state_change_prevents_merge(self):
        fresh = pull_request()
        fresh["head"]["sha"] = "changed-head"
        states = [fresh]
        for field, value in (("state", "closed"), ("draft", True), ("mergeable_state", "blocked")):
            fresh = pull_request()
            fresh[field] = value
            states.append(fresh)
        for fresh in states:
            with self.subTest(fresh=fresh), self.assertRaises(ValueError):
                self.run_process(fresh=fresh)
            self.assertFalse(any(call.args[1:] for call in self.last_api.call_args_list))
        with self.assertRaises(ValueError):
            self.run_process(fresh_main="changed-main")
        for mergeable, state in ((False, "dirty"), (None, "unknown"), (True, "blocked")):
            pr = pull_request()
            pr.update(mergeable=mergeable, mergeable_state=state)
            with self.subTest(state=state), self.assertRaises(ValueError):
                self.run_process(pr=pr)
            self.last_checks.assert_not_called()

    def test_rebase_requests_are_deduplicated_by_head_main_and_bot_author(self):
        marker = f"<!-- dependabot-rebase:{HEAD}:{MAIN} -->"
        self.assertIn("rebase を依頼", self.run_process(behind=True))
        self.last_api.assert_called_with(f"{ROOT}/issues/7/comments", "POST",
                                         {"body": f"@dependabot rebase\n\n{marker}"})
        self.last_checks.assert_not_called()
        for author, body, requested in (("github-actions[bot]", marker, False),
                                        ("human", marker, True),
                                        ("github-actions[bot]", marker.replace(MAIN, PARENT), True)):
            comments = [{"user": {"login": author}, "body": body}]
            with self.subTest(author=author, body=body):
                self.run_process(behind=True, comments=comments)
                mutations = [call for call in self.last_api.call_args_list if call.args[1:]]
                self.assertEqual(bool(mutations), requested)
        self.assertIn("予定", self.run_process(behind=True, dry_run=True))
        self.assertFalse(any(call.args[1:] for call in self.last_api.call_args_list))
        fresh = pull_request()
        fresh["head"]["sha"] = "changed-head"
        with self.assertRaises(ValueError):
            self.run_process(behind=True, fresh=fresh)
        with self.assertRaises(ValueError):
            self.run_process(behind=True, fresh_main="changed-main")


class MainTests(unittest.TestCase):
    def test_one_api_failure_reports_error_and_continues_remaining_prs(self):
        prs = [{"number": number, "user": {"login": automerge.BOT}} for number in (1, 2)]
        failure = subprocess.CalledProcessError(1, ["gh", "api"], stderr="HTTP403")
        output = io.StringIO()
        with patch.object(automerge, "items", return_value=prs), \
                patch.object(automerge, "process", side_effect=[failure, "保留: CI 待ち"]) as process, \
                patch.dict("os.environ", {"GITHUB_STEP_SUMMARY": ""}), \
                contextlib.redirect_stdout(output):
            self.assertEqual(automerge.main(["--repo", REPO, "--dry-run"]), 1)
        self.assertEqual([call.args for call in process.call_args_list], [(REPO, 1, True), (REPO, 2, True)])
        self.assertIn("#1: API / 実行エラーで保留: HTTP403", output.getvalue())
        self.assertIn("#2: 保留: CI 待ち", output.getvalue())


if __name__ == "__main__":
    unittest.main()
