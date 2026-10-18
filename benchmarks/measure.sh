#!/bin/sh
# raw はリポジトリ外の一時入力のみ。成功・失敗・中断のいずれでも削除する。
set -eu
cd "$(dirname "$0")/.."
GO_BINARY=${GO_BINARY:-go}
BENCHSTAT=${BENCHSTAT:-benchstat}
BENCHTIME=${BENCHTIME:-100ms}
COUNT=${COUNT:-10}
export GOMAXPROCS=${GOMAXPROCS:-2}
command -v "$GO_BINARY" >/dev/null
command -v "$BENCHSTAT" >/dev/null
workdir=$(mktemp -d /tmp/seq-benchmark.XXXXXX)
trap 'rm -rf "$workdir"' EXIT
trap 'exit 130' INT
trap 'exit 143' TERM
export GOCACHE=${GOCACHE:-"$workdir/go-cache"}
"$GO_BINARY" test ./benchmarks -count=1
# 各巡で Go / Library を一組ずつ測り、長い連続測定の偏りを減らす。
round=0
while [ "$round" -lt "$COUNT" ]; do
 "$GO_BINARY" test ./benchmarks -run '^$' -bench . -benchmem \
   -benchtime="$BENCHTIME" -count=1 >> "$workdir/raw.txt"
 round=$((round + 1))
done
python3 benchmarks/publish.py validate "$workdir/raw.txt" "$COUNT"
"$BENCHSTAT" -format=csv - < "$workdir/raw.txt" > "$workdir/all.csv"
compare() {
 name=$1
 filter=$2
 columns=$3
 "$BENCHSTAT" -filter "$filter" -col "$columns" -format=csv - \
   < "$workdir/raw.txt" > "$workdir/$name.csv"
}
compare unit '.name:Unit' '/impl@(Go Library)'
compare pipeline-loop '.name:ScenarioPipeline' '/impl@(Loop Pipeline)'
compare pipeline-functions '.name:ScenarioPipeline' '/impl@(Functions Pipeline)'
compare channel-from '.name:ScenarioChannel' '/impl@(Direct FromChannel)'
compare channel-roundtrip '.name:ScenarioChannel' '/impl@(Direct RoundTrip)'
compare channel-to '.name:ScenarioChannel' '/impl@(Seq ToChannel)'
compare cache-reuse '.name:ScenarioCache /phase:Reuse' '/impl@(Uncached Cached)'
"$GO_BINARY" version > "$workdir/go-version.txt"
"$GO_BINARY" env GOOS GOARCH GOARM64 GOAMD64 GOGC GOMEMLIMIT GOEXPERIMENT > "$workdir/go-env.txt"
git rev-parse HEAD > "$workdir/commit.txt"
python3 benchmarks/publish.py report "$workdir" "$BENCHTIME" "$COUNT" "$GOMAXPROCS"
# 検証・全ファイルの生成が成功してから公開結果を置き換える。
mkdir -p benchmarks/results
cp "$workdir/published/results/"*.md benchmarks/results/
cp "$workdir/published/REPORT.md" benchmarks/REPORT.md
# trap で raw を含む一時ディレクトリ全体を削除する。
