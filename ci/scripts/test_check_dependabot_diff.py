from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

import check_dependabot_diff as diff


REQUIREMENTS = "ci/scripts/requirements.txt"
WORKFLOW = ".github/workflows/ci.yml"


def dependency(name="packaging", version="26.1", **changes):
    item = {
        "dependencyName": name, "newVersion": version,
        "updateType": "version-update:semver-minor", "packageEcosystem": "pip",
        "directory": "/ci/scripts", "targetBranch": "main",
    }
    return dict(item, **changes)


class DiffTests(unittest.TestCase):
    def validate(self, contents, dependencies=None, tags=None, mode=None):
        if dependencies is None:
            dependencies = [dependency()]
        original_run = diff.run

        with tempfile.TemporaryDirectory() as root:
            def run(*args):
                if args[:2] == ("gh", "api"):
                    self.assertIn(args[2], tags or {})
                    return tags[args[2]] + "\n"
                return original_run("git", "-C", root, *args[1:])

            run("git", "init", "--quiet")
            for side in (0, 1):
                for name, versions in contents.items():
                    path = Path(root, name)
                    if versions[side] is None:
                        path.unlink(missing_ok=True)
                    else:
                        path.parent.mkdir(parents=True, exist_ok=True)
                        path.write_text(versions[side], encoding="utf-8")
                        if side == 1 and mode is not None:
                            path.chmod(mode)
                run("git", "add", "-A")
                run("git", "-c", "user.name=Test", "-c", "user.email=test@example.com",
                    "-c", "commit.gpgsign=false", "commit", "--quiet", "--allow-empty", "-m", str(side))
            head = run("git", "rev-parse", "HEAD").strip()
            with patch.object(diff, "run", side_effect=run):
                return diff.validate(dependencies, head)

    def test_grouped_requirements_and_repeated_actions(self):
        contents = {REQUIREMENTS: ("packaging==26.0\nPyYAML==6.0.2\n",
                                   "packaging==26.1\nPyYAML==6.0.3\n")}
        dependencies = [dependency(), dependency("PyYAML", "6.0.3", updateType="version-update:semver-patch")]
        self.assertEqual(self.validate(contents, dependencies), ["PyYAML", "packaging"])
        contents = {path: ("    - uses: actions/checkout@v6.0.0\n",
                           "    - uses: actions/checkout@v6.0.1\n")
                    for path in (WORKFLOW, ".github/workflows/lint.yml")}
        action = dependency("actions/checkout", "6.0.1", packageEcosystem="github_actions", directory="/")
        self.assertEqual(self.validate(contents, [action]), ["actions/checkout"])

    def test_every_group_member_must_be_present_and_eligible(self):
        contents = {REQUIREMENTS: ("packaging==26.0\nPyYAML==6.0.2\n", "packaging==26.1\nPyYAML==6.0.3\n")}
        cases = [[dependency()],
                 [dependency(), dependency("PyYAML", "6.0.3", updateType="version-update:semver-major")]]
        for dependencies in cases:
            with self.subTest(dependencies=dependencies), self.assertRaises(ValueError):
                self.validate(contents, dependencies)

    def test_normalized_requirement_names(self):
        contents = {REQUIREMENTS: ("pack_name==1.0.0\n", "pack_name==1.0.1\n")}
        self.assertEqual(self.validate(contents, [dependency("pack-name", "1.0.1")]), ["pack-name"])

    def test_metadata_must_describe_every_changed_dependency(self):
        contents = {REQUIREMENTS: ("packaging==26.0\n", "packaging==26.1\n")}
        cases = [[], [dependency(version="26.2")], [dependency(), dependency("PyYAML", "6.0.3")],
                 [dependency(), dependency("Packaging", "26.1")]]
        for key, value in (("updateType", "version-update:semver-major"), ("updateType", "unknown"),
                           ("packageEcosystem", "gomod"), ("directory", "/"), ("targetBranch", "dev")):
            cases.append([dependency(**{key: value})])
        incomplete = dependency()
        del incomplete["updateType"]
        cases.append([incomplete])
        for dependencies in cases:
            with self.subTest(dependencies=dependencies), self.assertRaises(ValueError):
                self.validate(contents, dependencies)

    def test_only_stable_minor_or_patch_increases_are_allowed(self):
        cases = [("1.2.3", "1.2.4", True), ("1.2", "1.3.0", True), ("0.1.0", "0.2.0", True),
                 ("1.9.9", "2.0.0", False), ("1.2", "1.2.0", False), ("1.2.3", "1.2.3", False),
                 ("1.3.0", "1.2.9", False), ("1.2.3", "1.3.0-rc.1", False),
                 ("main", "1.3.0", False), ("1.2.3", "1.2.4+build", False)]
        for before, after, accepted in cases:
            contents = {REQUIREMENTS: (f"packaging=={before}\n", f"packaging=={after}\n")}
            with self.subTest(before=before, after=after):
                if accepted:
                    self.assertEqual(self.validate(contents, [dependency(version=after)]), ["packaging"])
                else:
                    with self.assertRaises(ValueError):
                        self.validate(contents, [dependency(version=after)])

    def test_other_changes_are_rejected(self):
        before = "    - uses: actions/setup-go@v6.0.0\n"
        after = "    - uses: actions/setup-go@v6.0.1\n"
        action = dependency("actions/setup-go", "6.0.1", packageEcosystem="github_actions", directory="/")
        cases = [("go.mod", "go 1.27.2\n", "go 1.27.3\n"),
                 (WORKFLOW, before + "      go-version: 1.27.2\n", after + "      go-version: 1.27.3\n"),
                 (WORKFLOW, before + "      with:\n        token: old\n", after + "      with:\n        token: new\n"),
                 (WORKFLOW, before, after + "    - run: echo added\n"),
                 (WORKFLOW, before + "    - run: echo removed\n", after),
                 (WORKFLOW, before, after.replace("    - uses:", "      uses:")),
                 (WORKFLOW, before, after.replace("actions/setup-go", "other/setup-go")),
                 (REQUIREMENTS, "packaging==26.0\n", "packaging>=26.1\n"),
                 (REQUIREMENTS, "packaging==26.0\n", "packaging==26.1"),
                 (REQUIREMENTS, "packaging==26.0\n", "packaging==26.1\nnew==1.0.0\n")]
        for path, old, new in cases:
            with self.subTest(path=path, new=new), self.assertRaises(ValueError):
                self.validate({path: (old, new)}, [action] if path == WORKFLOW else [dependency()])

    def test_added_removed_renamed_and_empty_diffs_are_rejected(self):
        before, after = "packaging==26.0\n", "packaging==26.1\n"
        cases = [{REQUIREMENTS: (None, after)}, {REQUIREMENTS: (before, None)},
                 {REQUIREMENTS: (before, None), "ci/scripts/renamed.txt": (None, after)},
                 {REQUIREMENTS: (before, before)}]
        for contents in cases:
            with self.subTest(contents=contents), self.assertRaises(ValueError):
                self.validate(contents)
        with self.assertRaises(ValueError):
            self.validate({REQUIREMENTS: (before, after)}, mode=0o755)

    def test_sha_pins_require_matching_version_tags_and_metadata(self):
        before, after = "d" * 40, "e" * 40
        old = f"    - uses: actions/checkout@{before} # v6.0.0\n"
        new = f"    - uses: actions/checkout@{after} # v6.0.1\n"
        tags = {"repos/actions/checkout/commits/v6.0.0": before,
                "repos/actions/checkout/commits/v6.0.1": after}
        for version in (after, "6.0.1"):
            action = dependency("actions/checkout", version, packageEcosystem="github_actions", directory="/")
            with self.subTest(version=version):
                self.assertEqual(self.validate({WORKFLOW: (old, new)}, [action], tags), ["actions/checkout"])
        for version in ("6.0.2", "f" * 40):
            action = dependency("actions/checkout", version, packageEcosystem="github_actions", directory="/")
            with self.subTest(version=version), self.assertRaises(ValueError):
                self.validate({WORKFLOW: (old, new)}, [action], tags)
        action = dependency("actions/checkout", after, packageEcosystem="github_actions", directory="/")
        cases = [(old.split(" #")[0] + "\n", new, tags),
                 (old, "    - uses: actions/checkout@v6.0.1\n", tags),
                 (old, new, {key: "f" * 40 for key in tags})]
        for old, new, values in cases:
            with self.subTest(old=old, new=new, tags=values), self.assertRaises(ValueError):
                self.validate({WORKFLOW: (old, new)}, [action], values)


if __name__ == "__main__":
    unittest.main()
