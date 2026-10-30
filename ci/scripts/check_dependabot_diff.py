"""Match verified Dependabot metadata against dependency-only Git changes."""

import json
import os
import re
import subprocess
import sys

USES = re.compile(
    r"(?P<prefix>[ \t]*(?:-[ \t]+)?uses:[ \t]+)"
    r"(?P<name>[\w.-]+/[\w./-]+)@(?P<ref>[\w.-]+)"
    r"(?:[ \t]+#[ \t]*(?P<version>v?[0-9]+(?:\.[0-9]+){1,2}))?[ \t]*"
)
PIN = re.compile(r"(?P<name>[\w.-]+)==(?P<ref>[0-9]+(?:\.[0-9]+){1,2})")


def run(*args):
    return subprocess.run(args, check=True, capture_output=True).stdout.decode("utf-8")


def canonical(name, ecosystem):
    return re.sub(r"[-_.]+", "-", name).lower() if ecosystem == "pip" else name


def minor_or_patch(before, after):
    versions = []
    for value in (before, after):
        match = re.fullmatch(r"v?([0-9]+)\.([0-9]+)(?:\.([0-9]+))?", value or "")
        if not match:
            raise ValueError(f"安定版の minor / patch と分類できません: {value}")
        versions.append(tuple(int(part or 0) for part in match.groups()))
    if versions[0][0] != versions[1][0] or versions[0] >= versions[1]:
        raise ValueError(f"minor / patch 更新ではありません: {before} -> {after}")


def validate(dependencies, head):
    if not isinstance(dependencies, list) or not dependencies:
        raise ValueError("依存 metadata が空または不正です")
    expected = {}
    for dependency in dependencies:
        if not isinstance(dependency, dict):
            raise ValueError("依存 metadata が不正です")
        name = dependency.get("dependencyName")
        ecosystem = dependency.get("packageEcosystem")
        version = dependency.get("newVersion")
        if (dependency.get("updateType") not in ("version-update:semver-minor",
                                                 "version-update:semver-patch")
                or dependency.get("targetBranch") != "main"
                or (ecosystem, dependency.get("directory")) not in (
                    ("github_actions", "/"), ("pip", "/ci/scripts"))
                or not isinstance(name, str) or not name or name.lower() == "unknown"
                or not isinstance(version, str) or not version):
            raise ValueError("major・対象外・分類不能の依存 metadata があります")
        key = (ecosystem, canonical(name, ecosystem))
        if key in expected:
            raise ValueError(f"依存 metadata が重複しています: {name}")
        expected[key] = dependency
    if not re.fullmatch(r"[0-9a-f]{40}", head):
        raise ValueError("PR head SHA が不正です")
    parent = head + "^"
    files = run("git", "diff", "--no-renames", "--name-status", "-z", parent, head, "--").split("\0")[:-1]
    if not files or run("git", "diff", "--summary", parent, head, "--"):
        raise ValueError("変更が空、またはファイルの追加・削除・mode 変更があります")
    changed = set()
    for status, path in zip(files[::2], files[1::2]):
        workflow = re.fullmatch(r"\.github/workflows/[^/]+\.ya?ml", path)
        if status != "M" or not (workflow or path == "ci/scripts/requirements.txt"):
            raise ValueError(f"対象外のファイル変更です: {path}")
        ecosystem = "github_actions" if workflow else "pip"
        old = run("git", "show", f"{parent}:{path}").splitlines(keepends=True)
        new = run("git", "show", f"{head}:{path}").splitlines(keepends=True)
        if len(old) != len(new):
            raise ValueError(f"行の追加・削除があります: {path}")
        for before, after in zip(old, new):
            if before == after:
                continue
            pattern = USES if workflow else PIN
            left = pattern.fullmatch(before.removesuffix("\n"))
            right = pattern.fullmatch(after.removesuffix("\n"))
            if (not left or not right or left["name"] != right["name"]
                    or (workflow and left["prefix"] != right["prefix"])
                    or before.endswith("\n") != after.endswith("\n")):
                raise ValueError(f"依存の版番号以外の変更があります: {path}")
            key = (ecosystem, canonical(right["name"], ecosystem))
            if key not in expected:
                raise ValueError(f"差分に対応する metadata がありません: {right['name']}")
            versions = [left["ref"], right["ref"]]
            if workflow and any(re.fullmatch(r"[0-9a-f]{40}", ref) for ref in versions):
                if not all(re.fullmatch(r"[0-9a-f]{40}", ref) for ref in versions):
                    raise ValueError("SHA 固定方法が変更されています")
                versions = [left["version"], right["version"]]
                if not all(versions):
                    raise ValueError("SHA 更新の版番号コメントがありません")
                upstream = "/".join(right["name"].split("/")[:2])
                for match, version in zip((left, right), versions):
                    resolved = run("gh", "api", f"repos/{upstream}/commits/{version}", "--jq", ".sha").strip()
                    if resolved != match["ref"]:
                        raise ValueError("版番号タグと固定 SHA が一致しません")
            actual = {right["ref"], versions[1]}
            version = expected[key]["newVersion"]
            if workflow:
                actual = {value.removeprefix("v") for value in actual}
                version = version.removeprefix("v")
            if version not in actual:
                raise ValueError(f"差分と metadata の版番号が一致しません: {right['name']}")
            minor_or_patch(*versions)
            changed.add(key)
    if changed != expected.keys():
        raise ValueError("グループの全依存と差分が一致しません")
    return sorted(expected[key]["dependencyName"] for key in changed)


def main():
    try:
        dependencies = json.loads(os.environ["DEPENDENCIES_JSON"])
        names = validate(dependencies, os.environ["PR_HEAD_SHA"])
    except (ValueError, KeyError, OSError, subprocess.CalledProcessError) as error:
        print(f"保留: {error}", file=sys.stderr)
        return 1
    print("差分確認済み: " + ", ".join(names))
    return 0


if __name__ == "__main__":
    sys.exit(main())
