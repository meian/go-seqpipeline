import contextlib
import io
from pathlib import Path
import tempfile
from textwrap import indent
import unittest

from check_readme import INSTALL_COMMAND, MAX_LINES, check_readme, main


def block(body="fmt.Println(1)", language="go", marker="```"):
    return f"{marker}{language}\n{body}\n{marker}\n"


class CheckReadmeTests(unittest.TestCase):
    def test_normal_readme(self):
        text = "# Library\n\n" + block(INSTALL_COMMAND, "sh") + block() * 2
        self.assertEqual(check_readme(text), [])

    def test_line_boundary(self):
        for count in (MAX_LINES - 1, MAX_LINES, MAX_LINES + 1):
            with self.subTest(lines=count):
                errors = check_readme("text\n" * count)
                self.assertEqual(bool(errors), count > MAX_LINES)
                if errors:
                    self.assertIn(f"{count} lines", errors[0])

    def test_blank_lines_tables_and_code_count(self):
        text = "\n" * (MAX_LINES - 5) + "| a | b |\n| --- | --- |\n" + block()
        self.assertEqual(len(text.splitlines()), MAX_LINES)
        self.assertEqual(check_readme(text), [])
        self.assertIn("lines exceeds", check_readme(text + "\n")[0])

    def test_final_line_without_newline_counts(self):
        self.assertIn("lines exceeds", check_readme("\n" * MAX_LINES + "last")[0])

    def test_third_example_fails(self):
        errors = check_readme(block() * 3)
        self.assertIn("3 code examples", errors[0])

    def test_non_go_examples_count(self):
        for language in ("", "sh", "text", "python"):
            with self.subTest(language=language):
                errors = check_readme(block() * 2 + block("other command", language))
                self.assertIn("3 code examples", errors[0])

    def test_only_standalone_installation_is_exempt(self):
        text = block() * 2
        for language in ("sh", "shell", "bash"):
            with self.subTest(language=language):
                self.assertEqual(check_readme(text + block(INSTALL_COMMAND, language)), [])
        self.assertTrue(check_readme(text + block(INSTALL_COMMAND + "\ngo test ./...", "sh")))
        self.assertTrue(check_readme(text + block(INSTALL_COMMAND, "go")))

    def test_fence_variants(self):
        for marker in ("~~~", "````"):
            with self.subTest(marker=marker):
                self.assertTrue(check_readme(block(marker=marker) * 3))
        # Shorter or different markers inside a fence do not close it.
        self.assertEqual(check_readme("````go\n```\n~~~\n`````\n"), [])
        self.assertEqual(check_readme("  ~~~go\nx\n   ~~~~\n"), [])

    def test_nested_fences_count(self):
        containers = {
            "blockquote": indent(block(), "> "),
            "nested blockquote": indent(block(), "> > "),
            "list": "- Example\n\n" + indent(block(), "    "),
            "ordered list": "1. Example\n\n" + indent(block(), "    "),
            "list with quote": "- Example\n\n" + indent(block(), "    > "),
            "quote with list": "> - Example\n>\n" + indent(block(), ">     "),
            "fence beside list marker": "- ```go\n  fmt.Println(1)\n  ```\n",
        }
        for name, example in containers.items():
            with self.subTest(container=name):
                self.assertEqual(check_readme((example + "\n") * 2), [])
                errors = check_readme((example + "\n") * 3)
                self.assertTrue(any("3 code examples" in error for error in errors))
                # A nested example must also count alongside top-level ones.
                self.assertTrue(check_readme(block() * 2 + "\n" + example))

    def test_indented_code_counts(self):
        for prefix in ("    ", "\t"):
            with self.subTest(prefix=repr(prefix)):
                example = "Example\n\n" + indent("fmt.Println(1)\n", prefix) + "\n"
                self.assertEqual(check_readme(example * 2), [])
                self.assertIn("3 code examples", check_readme(example * 3)[0])

    def test_nested_installation_is_exempt(self):
        installation = block(INSTALL_COMMAND, "sh")
        for example in (indent(installation, "> "), "- Install\n\n" + indent(installation, "    ")):
            with self.subTest(example=example):
                self.assertEqual(check_readme(block() * 2 + "\n" + example), [])

    def test_fence_text_inside_code_is_not_an_extra_example(self):
        for example in (block(block() * 3, marker="~~~~"), indent(block() * 3, "    ")):
            with self.subTest(example=example):
                self.assertEqual(check_readme(example), [])

    def test_nested_unclosed_fence_fails(self):
        errors = check_readme("> ```go\n> example\n\noutside\n")
        self.assertIn("Unclosed code fence", errors[0])

    def test_unclosed_fence_fails(self):
        self.assertIn("Unclosed code fence", check_readme("```go\nexample\n")[0])


class MainTests(unittest.TestCase):
    def test_both_languages_are_checked_and_failures_explain_destination(self):
        with tempfile.TemporaryDirectory() as directory:
            paths = [Path(directory) / name for name in ("README.md", "README.ja.md")]
            for path in paths:
                path.write_text(block() * 3, encoding="utf-8")
            output = io.StringIO()
            with contextlib.redirect_stderr(output):
                self.assertEqual(main(paths), 1)
            for path in paths:
                self.assertIn(str(path), output.getvalue())
            self.assertIn("doc/guide.md", output.getvalue())
            self.assertIn("GoDoc", output.getvalue())
            self.assertIn("Go Examples", output.getvalue())

    def test_success_and_missing_file_exit_status(self):
        with tempfile.TemporaryDirectory() as directory:
            path = Path(directory) / "README.md"
            path.write_text("# Library\n", encoding="utf-8")
            self.assertEqual(main([path]), 0)
            with contextlib.redirect_stderr(io.StringIO()):
                self.assertEqual(main([path.with_name("missing.md")]), 1)


if __name__ == "__main__":
    unittest.main()
