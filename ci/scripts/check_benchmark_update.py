"""Create a benchmark update issue when main differs from the published snapshot."""

import argparse
import datetime
import json
import os
from pathlib import Path, PurePosixPath
import re
import subprocess
import sys

ENVIRONMENT = "benchmarks/results/environment.md"
MARKER = "<!-- go-seqpipeline:benchmark-update -->"
RUNTIME_CONFIGS = {
    "go.mod", "go.sum", "go.work", "go.work.sum", ".go-version", ".tool-versions",
}


def run(*args, input=None):
    return subprocess.run(
        args, input=input, check=True, capture_output=True, text=True,
        encoding="utf-8",
    ).stdout


def read_baseline(repository, target):
    """Read the measured commit, distinguishing an initial run from a broken report."""
    paths = run("git", "ls-tree", "-r", "--name-only", "-z", target, "--", "benchmarks")
    paths = paths.split("\0")
    if ENVIRONMENT not in paths:
        if any(path == "benchmarks/REPORT.md" or path.startswith("benchmarks/results/")
               for path in paths if path):
            raise ValueError("測定結果はありますが、実行環境レポートがありません")
        return None

    report = run("git", "show", f"{target}:{ENVIRONMENT}")
    rows = [line for line in report.splitlines()
            if line.startswith("| 対象コード（ライブラリ・測定） |")]
    pattern = (
        r"\| 対象コード（ライブラリ・測定） \| \[([0-9a-f]{12})\]"
        r"\(https://github\.com/" + re.escape(repository)
        + r"/commit/([0-9a-f]{40})\) \|"
    )
    match = re.fullmatch(pattern, rows[0]) if len(rows) == 1 else None
    if not match or match[1] != match[2][:12]:
        raise ValueError("実行環境レポートの測定対象コミットが不正または一意ではありません")
    return match[2]


def ensure_commit(commit):
    """Fetch a measured branch commit if a squash merge left it outside main's history."""
    try:
        run("git", "cat-file", "-e", f"{commit}^{{commit}}")
    except subprocess.CalledProcessError:
        run("git", "fetch", "--no-tags", "origin", commit)
        run("git", "cat-file", "-e", f"{commit}^{{commit}}")


def is_runtime_path(name):
    """Select library sources and runtime settings, excluding tests and documentation."""
    path = PurePosixPath(name)
    if name in RUNTIME_CONFIGS:
        return True
    if "testdata" in path.parts or path.parts[0] in {"benchmarks", "ci", "doc", ".github"}:
        return False
    return path.suffix == ".go" and path.name != "doc.go" and not path.name.endswith("_test.go")


def changed_files(baseline, target):
    """Compare snapshots directly, retaining deletions when a source is renamed to a test."""
    paths = run(
        "git", "diff", "--no-ext-diff", "--no-renames", "--name-only", "-z",
        baseline, target, "--",
    )
    return sorted(path for path in paths.split("\0") if path and is_runtime_path(path))


def open_update_issue(repository):
    """Find an open marked issue across every API page, excluding pull requests."""
    pages = json.loads(run(
        "gh", "api", "--method", "GET", "--paginate", "--slurp",
        f"repos/{repository}/issues?state=open&per_page=100",
    ))
    if not isinstance(pages, list) or not pages or any(not isinstance(page, list) for page in pages):
        raise ValueError("Issue API のページ一覧が不正です")
    for page in pages:
        for issue in page:
            if "pull_request" not in issue and MARKER in (issue.get("body") or "").splitlines():
                return issue["html_url"]
    return None


def issue_body(repository, baseline, target, files):
    root = f"https://github.com/{repository}"
    report = (f"[直近のレポート]({root}/blob/{target}/benchmarks/REPORT.md)"
              if baseline else "直近のレポート: 未作成")
    if baseline:
        comparison = (
            f"- 前回の測定対象: `{baseline}`\n"
            f"- 今回の確認対象: `{target}`\n"
            f"- [コミット比較]({root}/compare/{baseline}..{target})\n\n"
            "### 検出した実装・実行設定の変更\n\n"
            + "\n".join(f"- `{path}`" for path in files)
        )
    else:
        comparison = (
            "測定結果がまだないため、初回の測定・レポート作成を行う。\n\n"
            f"- 前回の測定対象: なし\n- 今回の確認対象: `{target}`"
        )
    return f"""{MARKER}
## 定期ベンチマーク更新

{comparison}

## 作業内容

- [ ] 現行 API と実装差分を確認し、ベンチマークコード・比較条件を見直す
- [ ] 測定コードをコミットし、通常テスト・等価性テストを確認する
- [ ] 全ケースを複数回測定し、集計・実測結果・概要レポートを更新する
- [ ] 実際の測定対象コミットと測定日時を `benchmarks/results/environment.md` に記録する
- [ ] レポート更新 PR を main に反映してから、この Issue を完了する

## 比較基準の更新

測定には更新作業時点のコードを使用し、今回の確認対象コミットから変更した場合は実際に測定したコミットを記録する。
`benchmarks/measure.sh` は測定対象コミットと測定日時を含むレポートを生成する。
更新結果が main に反映されると、そのレポートの測定対象コミットが次回の比較基準になる。
この Issue の作成・クローズだけでは比較基準は進まない。

## 参照

- [初回ベンチマークの方針 #32]({root}/issues/32)
- {report}
- [更新手順]({root}/blob/{target}/doc/ai/benchmark.md)
"""


def create_issue(repository, body):
    today = datetime.datetime.now(datetime.timezone(datetime.timedelta(hours=9))).date()
    payload = {"title": f"[定期更新] ベンチマークを更新する ({today})", "body": body}
    issue = json.loads(run(
        "gh", "api", "--method", "POST", f"repos/{repository}/issues", "--input", "-",
        input=json.dumps(payload, ensure_ascii=False),
    ))
    return issue["html_url"]


def check(repository, dry_run):
    """Inspect a fixed main snapshot and create at most one pending update issue."""
    target = run("git", "rev-parse", "--verify", "refs/remotes/origin/main^{commit}").strip()
    baseline = read_baseline(repository, target)
    measured = f"`{baseline}`" if baseline else "初回測定なし"
    summary = f"# ベンチマーク更新チェック\n\n確認対象: `{target}`\n\n比較基準: {measured}\n\n"
    files = []
    if baseline:
        ensure_commit(baseline)
        files = changed_files(baseline, target)
        if not files:
            return summary + "結果: 対象差分なし。Issue は作成しません。\n"

    existing = open_update_issue(repository)
    if existing:
        return summary + f"結果: 未完了の更新 Issue があるため作成を見送ります。\n\n{existing}\n"
    body = issue_body(repository, baseline, target, files)
    if dry_run:
        return summary + "結果: dry-run。以下の Issue を作成予定です。\n\n" + body
    url = create_issue(repository, body)
    return summary + f"結果: 更新 Issue を作成しました。\n\n{url}\n"


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo", required=True, help="GitHub repository: OWNER/REPO")
    parser.add_argument("--dry-run", action="store_true", help="判定と Issue 本文の確認のみ")
    args = parser.parse_args(argv)
    if not re.fullmatch(r"[A-Za-z0-9_.-]+/[A-Za-z0-9_.-]+", args.repo):
        parser.error("--repo must be OWNER/REPO")
    try:
        summary = check(args.repo, args.dry_run)
        status = 0
    except (OSError, ValueError, KeyError, TypeError, subprocess.CalledProcessError) as error:
        detail = error.stderr if isinstance(error, subprocess.CalledProcessError) else str(error)
        summary = f"# ベンチマーク更新チェック失敗\n\n{detail}\n"
        status = 1
    print(summary, file=sys.stderr if status else sys.stdout)
    if os.environ.get("GITHUB_STEP_SUMMARY"):
        with Path(os.environ["GITHUB_STEP_SUMMARY"]).open("a", encoding="utf-8") as stream:
            stream.write(summary)
    return status


if __name__ == "__main__":
    sys.exit(main())
