"""Rebase and merge narrowly classified Dependabot updates after fresh checks."""

import argparse
import base64
import json
import os
from pathlib import Path
import re
import subprocess
import sys
from urllib.parse import quote

import yaml

BOT = "dependabot[bot]"
ACTION_APP = 15368
WORKFLOW = ".github/workflows/dependabot-automerge.yml"
REQUIRED = {"gofmt", "build", "test", "Static checks", "README policy"}
USES = re.compile(
    r"(?P<prefix>\s*(?:-\s+)?uses:\s+)(?P<name>[\w.-]+/[\w./-]+)@"
    r"(?P<ref>[\w.-]+)(?:\s+#\s*(?P<version>v?\d+(?:\.\d+){1,2}))?\s*"
)
PIN = re.compile(r"(?P<name>[\w.-]+)==(?P<ref>\d+(?:\.\d+){1,2})")


def api(path, method="GET", data=None):
    args = ["gh", "api", "-H", "X-GitHub-Api-Version: 2026-03-10", "--method", method, path]
    if data is not None:
        args += ["--input", "-"]
    result = subprocess.run(
        args, input=json.dumps(data) if data is not None else None,
        check=True, capture_output=True, text=True, encoding="utf-8",
    )
    return json.loads(result.stdout)


def items(path, key=None):
    """Read every page, including check/status collections wrapped in an object."""
    result = subprocess.run(
        ["gh", "api", "--paginate", "--slurp", path],
        check=True, capture_output=True, text=True, encoding="utf-8",
    )
    pages = json.loads(result.stdout)
    return [item for page in pages for item in (page[key] if key else page)]


def metadata(commit):
    """Use the signed commit trailer, as dependabot/fetch-metadata does."""
    if (commit.get("author", {}).get("login") != BOT
            or not commit["commit"]["verification"]["verified"]
            or len(commit["parents"]) != 1):
        raise ValueError("Dependabot の署名付き単一親コミットではありません")
    fragments = re.findall(r"^---\n(.*?)\n\.\.\.(?:\n|$)",
                           commit["commit"]["message"], re.M | re.S)
    if len(fragments) != 1:
        raise ValueError("Dependabot metadata を一意に取得できません")
    dependencies = yaml.safe_load(fragments[0])["updated-dependencies"]
    result = {}
    for dependency in dependencies:
        name = dependency["dependency-name"]
        version = str(dependency["dependency-version"])
        kind = dependency.get("update-type")
        if (name in result or kind not in (None, "version-update:semver-minor",
                                           "version-update:semver-patch")):
            raise ValueError("major・分類不能・重複した依存 metadata があります")
        result[name] = version
    if not result:
        raise ValueError("依存 metadata が空です")
    return result


def minor_or_patch(before, after):
    def version(value):
        match = re.fullmatch(r"v?(\d+)\.(\d+)(?:\.(\d+))?", value)
        if not match:
            raise ValueError(f"安定版の minor / patch と分類できません: {value}")
        return tuple(int(part or 0) for part in match.groups())

    old, new = version(before), version(after)
    if old[0] != new[0] or old >= new:
        raise ValueError(f"minor / patch 更新ではありません: {before} -> {after}")


def read_file(repo, path, ref):
    data = api(f"repos/{repo}/contents/{quote(path, safe='/')}?ref={ref}")
    if data["type"] != "file" or data["encoding"] != "base64":
        raise ValueError(f"通常ファイルを取得できません: {path}")
    return base64.b64decode(data["content"], validate=False).decode("utf-8")


def classify(repo, pr):
    """Accept only existing pinned requirements and same-action uses replacements."""
    root = f"repos/{repo}"
    commits = items(f"{root}/pulls/{pr['number']}/commits?per_page=100")
    if len(commits) != 1 or commits[0]["sha"] != pr["head"]["sha"]:
        raise ValueError("PR は単一の Dependabot コミットではありません")
    dependencies = metadata(api(f"{root}/commits/{pr['head']['sha']}"))
    files = items(f"{root}/pulls/{pr['number']}/files?per_page=100")
    if not files or len(files) != pr["changed_files"]:
        raise ValueError("全変更ファイルを取得できません")
    changed = set()
    parent = commits[0]["parents"][0]["sha"]
    for file in files:
        path = file["filename"]
        workflow = re.fullmatch(r"\.github/workflows/[^/]+\.ya?ml", path)
        if file["status"] != "modified" or not (workflow or path == "ci/scripts/requirements.txt"):
            raise ValueError(f"自動分類の対象外ファイルです: {path}")
        old = read_file(repo, path, parent).splitlines(keepends=True)
        new = read_file(repo, path, pr["head"]["sha"]).splitlines(keepends=True)
        if len(old) != len(new):
            raise ValueError(f"行の追加・削除があります: {path}")
        for before, after in zip(old, new):
            if before == after:
                continue
            pattern = USES if workflow else PIN
            left, right = pattern.fullmatch(before.rstrip("\n")), pattern.fullmatch(after.rstrip("\n"))
            if not left or not right or left["name"] != right["name"]:
                raise ValueError(f"依存の版番号以外の変更があります: {path}")
            if workflow and left["prefix"] != right["prefix"]:
                raise ValueError(f"uses の配置が変更されています: {path}")
            name = right["name"]
            if not workflow:
                normalize = lambda value: re.sub(r"[-_.]+", "-", value).lower()
                matches = [dependency for dependency in dependencies
                           if normalize(dependency) == normalize(name)]
                if len(matches) != 1:
                    raise ValueError(f"依存名を一意に照合できません: {name}")
                name = matches[0]
            versions = [left["ref"], right["ref"]]
            if workflow and any(re.fullmatch(r"[0-9a-f]{40}", ref) for ref in versions):
                if not all(re.fullmatch(r"[0-9a-f]{40}", ref) for ref in versions):
                    raise ValueError("SHA 固定方法が変更されています")
                versions = [left["version"], right["version"]]
                if not all(versions):
                    raise ValueError("SHA 更新の版番号コメントがありません")
                for match, version in zip((left, right), versions):
                    upstream = "/".join(name.split("/")[:2])
                    resolved = api(f"repos/{upstream}/commits/{version}")["sha"]
                    if resolved != match["ref"]:
                        raise ValueError("版番号タグと固定 SHA が一致しません")
            expected = dependencies.get(name, "")
            candidates = {right["ref"], versions[1]}
            if workflow:
                expected = expected.removeprefix("v")
                candidates = {value.removeprefix("v") for value in candidates}
            if not expected or expected not in candidates:
                raise ValueError(f"差分と metadata が一致しません: {name}")
            minor_or_patch(*versions)
            changed.add(name)
    if changed != set(dependencies):
        raise ValueError("グループの全依存と差分が一致しません")
    return sorted(changed)


def required_checks(repo):
    rules = items(f"repos/{repo}/rules/branches/main?per_page=100")
    required = set()
    strict = False
    for rule in rules:
        if rule["type"] == "required_status_checks":
            params = rule["parameters"]
            strict |= params["strict_required_status_checks_policy"]
            required.update((check["context"], check.get("integration_id"))
                            for check in params["required_status_checks"])
    if not strict or not REQUIRED <= {name for name, _ in required}:
        raise ValueError("main の strict な必須チェック設定を確認できません")
    return required


def checks_pass(repo, head, required):
    """Keep the newest workflow per event; a skipped push lint cannot satisfy PR lint."""
    root = f"repos/{repo}"
    runs = items(f"{root}/actions/runs?head_sha={head}&per_page=100", "workflow_runs")
    latest = {}
    for run in runs:
        if run["path"] == WORKFLOW:
            continue
        key = (run["workflow_id"], run["event"])
        if key not in latest or run["id"] > latest[key]["id"]:
            latest[key] = run
    suites = {run["check_suite_id"]: run for run in latest.values()}
    ignored_suites = {run["check_suite_id"] for run in runs} - suites.keys()
    if not any(run["event"] == "pull_request" and run["path"] == ".github/workflows/ci.yml"
               for run in latest.values()):
        raise ValueError("最新 head の PR CI が未作成です")
    if any(run["status"] != "completed" or run["conclusion"] != "success"
           for run in latest.values()):
        raise ValueError("最新 head の workflow が実行中または未成功です")
    checks = items(f"{root}/commits/{head}/check-runs?filter=latest&per_page=100", "check_runs")
    successes = set()
    for check in checks:
        app = check["app"]["id"]
        suite = check["check_suite"]["id"]
        if app == ACTION_APP and suite in ignored_suites:
            continue
        if app == ACTION_APP and suite not in suites:
            raise ValueError("チェックに対応する workflow を取得できません")
        skipped_push_lint = (app == ACTION_APP and check["name"] == "Static checks"
                             and suites[suite]["event"] == "push"
                             and check["conclusion"] == "skipped")
        if skipped_push_lint:
            continue
        if check["status"] != "completed" or check["conclusion"] != "success":
            raise ValueError(f"チェックが未成功です: {check['name']}")
        successes.add((check["name"], app))
    statuses = api(f"{root}/commits/{head}/status?per_page=100")
    # The combined endpoint is capped at 100; use the paginated list to avoid omission.
    status_items = items(f"{root}/commits/{head}/statuses?per_page=100")
    newest = {}
    for status in status_items:
        newest.setdefault(status["context"], status)
    if statuses["total_count"] != len(newest):
        raise ValueError("全 commit status を取得できません")
    for status in newest.values():
        if status["state"] != "success":
            raise ValueError(f"commit status が未成功です: {status['context']}")
        successes.add((status["context"], None))
    for name, app in required:
        if not any(context == name and (app is None or app == source)
                   for context, source in successes):
            raise ValueError(f"必須チェックが未成功または未作成です: {name}")


def process(repo, number, dry_run):
    root = f"repos/{repo}"
    pr = api(f"{root}/pulls/{number}")
    if (pr["state"] != "open" or pr["draft"] or pr["user"]["login"] != BOT
            or pr["base"]["ref"] != "main" or pr["head"]["repo"]["full_name"] != repo):
        raise ValueError("同一リポジトリの main 向け Dependabot PR ではありません")
    dependencies = classify(repo, pr)
    required = required_checks(repo)
    head = pr["head"]["sha"]
    main = api(f"{root}/git/ref/heads/main")["object"]["sha"]
    comparison = api(f"{root}/compare/{main}...{head}")
    if comparison["merge_base_commit"]["sha"] != main:
        marker = f"<!-- dependabot-rebase:{head}:{main} -->"
        comments = items(f"{root}/issues/{number}/comments?per_page=100")
        if any(comment["user"]["login"] == "github-actions[bot]" and marker in comment["body"]
               for comment in comments):
            return "保留: rebase 依頼済み。新しい head と CI を待ちます"
        if not dry_run:
            fresh = api(f"{root}/pulls/{number}")
            current_main = api(f"{root}/git/ref/heads/main")["object"]["sha"]
            if fresh["head"]["sha"] != head or current_main != main:
                raise ValueError("head / main が変化しました。次回に再判定します")
            api(f"{root}/issues/{number}/comments", "POST",
                {"body": f"@dependabot rebase\n\n{marker}"})
        return "rebase 依頼予定" if dry_run else "rebase を依頼しました。新しい head と CI を待ちます"
    if pr["mergeable"] is not True or pr["mergeable_state"] != "clean":
        raise ValueError("競合・保護条件・マージ可能状態を確認できません")
    checks_pass(repo, head, required)
    fresh = api(f"{root}/pulls/{number}")
    current_main = api(f"{root}/git/ref/heads/main")["object"]["sha"]
    if (fresh["head"]["sha"] != head or current_main != main
            or fresh["state"] != "open" or fresh["draft"]
            or fresh["base"]["ref"] != "main"
            or fresh["mergeable_state"] != "clean"):
        raise ValueError("head / main または PR の状態が変化しました。次回に再判定します")
    if dry_run:
        return "マージ予定: " + ", ".join(dependencies)
    # sha guards head races; bypass_rules=false and strict rules guard base races.
    checks_pass(repo, head, required)
    endpoint = f"{root}/pulls/{number}/merge-async"
    result = api(endpoint, "PUT", {"sha": head, "merge_method": "squash",
                                   "merge_action": "direct_merge", "bypass_rules": False})
    if result["status"] == "pending":
        identifier = result["details"]["uuid"]
        result = api(f"{endpoint}/{identifier}")
        if result["status"] == "pending":
            return f"マージ処理中: {identifier}（完了未確認）"
    if result["status"] != "merged":
        raise ValueError(f"GitHub がマージを保留しました: {result['details']['message']}")
    return f"マージしました: {result['details']['sha']}"


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo", required=True)
    parser.add_argument("--pr", type=int)
    parser.add_argument("--dry-run", action="store_true")
    args = parser.parse_args(argv)
    if not re.fullmatch(r"[\w.-]+/[\w.-]+", args.repo):
        parser.error("--repo must be OWNER/REPO")
    if not args.dry_run and os.environ.get("GITHUB_ACTIONS") != "true":
        parser.error("変更操作は GitHub Actions で実行してください。ローカル確認は --dry-run を指定します")
    summary = ["# Dependabot 自動マージ\n"]
    failed = False
    try:
        numbers = [args.pr] if args.pr else [
            pr["number"] for pr in items(f"repos/{args.repo}/pulls?state=open&base=main&per_page=100")
            if pr["user"]["login"] == BOT
        ]
        for number in numbers:
            try:
                result = process(args.repo, number, args.dry_run)
            except (ValueError, KeyError, TypeError, AttributeError, yaml.YAMLError) as error:
                result = f"保留: {error}"
            except (OSError, subprocess.CalledProcessError) as error:
                detail = error.stderr if isinstance(error, subprocess.CalledProcessError) else str(error)
                result = f"API / 実行エラーで保留: {detail}"
                failed = True
            summary.append(f"- #{number}: {result}")
    except (OSError, subprocess.CalledProcessError) as error:
        summary.append(f"API / 実行エラーで保留: {error}")
        failed = True
    text = "\n".join(summary) + "\n"
    print(text)
    if os.environ.get("GITHUB_STEP_SUMMARY"):
        with Path(os.environ["GITHUB_STEP_SUMMARY"]).open("a", encoding="utf-8") as stream:
            stream.write(text)
    return int(failed)


if __name__ == "__main__":
    sys.exit(main())
