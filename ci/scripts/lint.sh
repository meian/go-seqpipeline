#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "$0")/../.." && pwd)
golangci_version=v2.15.0
reviewdog_version=v0.21.2
golangci_dir="$root/bin/lint/golangci-lint-$golangci_version"
reviewdog_dir="$root/bin/lint/reviewdog-$reviewdog_version"

install_tool() {
  local repo=$1 version=$2 destination=$3
  curl --fail --silent --show-error --location \
    "https://raw.githubusercontent.com/$repo/$version/install.sh" \
    -o "$workdir/install.sh"
  sh "$workdir/install.sh" -b "$destination" "$version"
}

case "${1:-run}" in
  install)
    workdir=$(mktemp -d)
    trap 'rm -rf "$workdir"' EXIT
    install_tool golangci/golangci-lint "$golangci_version" "$golangci_dir"
    install_tool reviewdog/reviewdog "$reviewdog_version" "$reviewdog_dir"
    "$golangci_dir/golangci-lint" version
    "$reviewdog_dir/reviewdog" -version
    exit 0
    ;;
  run) ;;
  *)
    printf 'Usage: %s [install|run]\n' "$0" >&2
    exit 2
    ;;
esac

cd "$root"
"$golangci_dir/golangci-lint" config verify
lint_status=0
"$golangci_dir/golangci-lint" run ./... || lint_status=$?
if [ "$lint_status" -ne 0 ] && [ "$lint_status" -ne 1 ]; then
  exit "$lint_status"
fi

"$reviewdog_dir/reviewdog" -f=checkstyle -reporter=local \
  -filter-mode=nofilter -fail-level=error < bin/lint/report.xml | tee bin/lint/report.txt
