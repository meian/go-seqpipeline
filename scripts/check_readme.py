"""Keep README introductions within their documented size and example budgets."""

import sys
from pathlib import Path

from markdown_it import MarkdownIt

MAX_LINES = 100
MAX_EXAMPLES = 2
README_PATHS = ("README.md", "README.ja.md")
INSTALL_COMMAND = "go get github.com/meian/go-seqpipeline"
MOVE_HINT = (
    "Move shared semantics to doc/guide.md and doc/guide.ja.md, "
    "API contracts to GoDoc, and usage examples to Go Examples."
)


def check_readme(text):
    """Return policy violations, counting all lines and Markdown code blocks."""
    errors = []
    lines = text.splitlines()
    if len(lines) > MAX_LINES:
        errors.append(f"{len(lines)} lines exceeds the {MAX_LINES}-line limit.")

    examples = 0
    for token in MarkdownIt("commonmark").parse(text):
        if token.type not in ("fence", "code_block"):
            continue
        # Parsed tokens include code inside blockquotes and lists, but do not
        # interpret literal fences inside a code block as additional examples.
        installation = (
            token.type == "fence"
            and token.info.strip() in ("sh", "shell", "bash")
            and token.content.strip() == INSTALL_COMMAND
        )
        if not installation:
            examples += 1
        if token.type == "fence":
            # The source span includes the opener and, when present, the closer.
            # Container prefixes are already removed from the parsed content.
            source_lines = token.map[1] - token.map[0]
            if source_lines == token.content.count("\n") + 1:
                errors.append("Unclosed code fence.")

    if examples > MAX_EXAMPLES:
        errors.append(
            f"{examples} code examples exceeds the {MAX_EXAMPLES}-example limit "
            "(installation excluded)."
        )
    return errors


def main(paths):
    failed = False
    for name in paths:
        try:
            errors = check_readme(Path(name).read_text(encoding="utf-8"))
        except (OSError, UnicodeError) as error:
            errors = [str(error)]
        if errors:
            failed = True
            for error in errors:
                print(f"{name}: {error} {MOVE_HINT}", file=sys.stderr)
    return int(failed)


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:] or README_PATHS))
