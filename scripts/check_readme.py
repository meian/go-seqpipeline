"""Keep README introductions within their documented size and example budgets."""

import re
import sys
from pathlib import Path

MAX_LINES = 100
MAX_EXAMPLES = 2
README_PATHS = ("README.md", "README.ja.md")
INSTALL_COMMAND = "go get github.com/meian/go-seqpipeline"
MOVE_HINT = (
    "Move shared semantics to doc/guide.md and doc/guide.ja.md, "
    "API contracts to GoDoc, and usage examples to Go Examples."
)


def check_readme(text):
    """Return policy violations, counting all lines and fenced examples."""
    errors = []
    lines = text.splitlines()
    if len(lines) > MAX_LINES:
        errors.append(f"{len(lines)} lines exceeds the {MAX_LINES}-line limit.")

    fence = None
    examples = 0
    for line in lines:
        if fence is None:
            opening = re.fullmatch(r" {0,3}(`{3,}|~{3,})(.*)", line)
            if opening:
                marker, info = opening.groups()
                fence = (marker, info.strip())
                body = []
        else:
            marker, info = fence
            closing_pattern = rf" {{0,3}}{marker[0]}{{{len(marker)},}}[ \t]*"
            if re.fullmatch(closing_pattern, line):
                # Only the standalone installation command is exempt, regardless
                # of how other examples are labelled (Go, shell, or unlabelled).
                installation = (
                    info in ("sh", "shell", "bash")
                    and "\n".join(body).strip() == INSTALL_COMMAND
                )
                if not installation:
                    examples += 1
                fence = None
            else:
                body.append(line)

    if fence is not None:
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
