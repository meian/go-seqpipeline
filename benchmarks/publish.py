#!/usr/bin/env python3
"""benchstat の集計済み CSV を整形する。統計計算は benchstat に任せる。"""
import csv
import datetime
import platform
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent


def validate(path, count):
    counts = {}
    for line in path.read_text().splitlines():
        if line.startswith('Benchmark') and ' ns/op' in line:
            name = line.split()[0]
            counts[name] = counts.get(name, 0) + 1
    if not counts or any(n != count for n in counts.values()) or '\nPASS\n' not in path.read_text():
        raise ValueError('missing or incomplete benchmark samples')
    print(f'{len(counts)} cases x {count} samples')
    return len(counts)


def comparisons(path):
    metric = None
    rows = {}
    for row in csv.reader(path.open()):
        if len(row) > 1 and row[0] == '' and row[1] in ('sec/op', 'B/op', 'allocs/op'):
            metric = row[1]
        elif metric and row and row[0].startswith('Scenario'):
            name = re.sub(r'-\d+$', '', row[0])
            rows[name, metric] = row
    return rows


def report(workdir, benchtime, count, procs):
    commit = (workdir/'commit.txt').read_text().strip()
    version = (workdir/'go-version.txt').read_text().strip()
    raw = (workdir/'raw.txt').read_text()
    cpu = next(line.removeprefix('cpu: ') for line in raw.splitlines() if line.startswith('cpu: '))
    cases = validate(workdir/'raw.txt', int(count))
    env_names = ['GOOS', 'GOARCH', 'GOARM64', 'GOAMD64', 'GOGC', 'GOMEMLIMIT', 'GOEXPERIMENT']
    env_values = (workdir/'go-env.txt').read_text().splitlines()
    benchstat = next(line.strip() for line in (workdir/'benchstat-version.txt').read_text().splitlines() if '\tmod\t' in line)
    timestamp = datetime.datetime.now(datetime.timezone.utc).isoformat(timespec='seconds')
    env = f'''# 測定環境

測定日時は {timestamp}、対象コミットは `{commit}`。  
対象コードと測定コードをこのコミットに固定し、文書・集計結果だけを後続コミットで公開する。  

| 項目 | 値 |
| --- | --- |
| Go | {version} |
| OS | {platform.platform()} |
| CPU | {cpu} |
| GOMAXPROCS | {procs} |
| benchtime / count | {benchtime} / {count} |
| ケース数 | {cases} |
| benchstat | {benchstat} |
'''
    for name, value in zip(env_names, env_values):
        env += f'| {name} | {value or "既定値"} |\n'
    env += '''
通常のコンパイラ最適化・インライン化を有効にし、CPU 固定・温度制御・測定順のランダム化は行っていない。  
測定中には別のテストやビルドを実行せず、OS や他プロセスの負荷は制御していない。  
Go のキャッシュと測定の一時ファイルには書き込み可能な一時ディレクトリを使用した。  
信頼区間は benchstat の既定の 95%、同一スナップショット内の比較検定の有意水準は 0.05。  
raw は集計時の一時入力のみで、集計後に削除する。  
'''
    (ROOT/'results/environment.md').write_text(env)
    selected = [
        ('pipeline-loop', 'ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=All', 'Loop → Pipeline'),
        ('pipeline-functions', 'ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=All', 'Functions → Pipeline'),
        ('pipeline-loop', 'ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=All', 'Loop → Pipeline'),
        ('pipeline-loop', 'ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=Early8', 'Loop → Pipeline'),
        ('channel-from', 'ScenarioChannel/N=4096/buffer=0', 'Direct → FromChannel'),
        ('channel-from', 'ScenarioChannel/N=4096/buffer=64', 'Direct → FromChannel'),
        ('channel-to', 'ScenarioChannel/N=4096/buffer=0', 'Seq → ToChannel'),
        ('channel-roundtrip', 'ScenarioChannel/N=4096/buffer=64', 'Direct → RoundTrip'),
        ('cache-reuse', 'ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=1', 'Uncached → Cached'),
        ('cache-reuse', 'ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=4', 'Uncached → Cached'),
        ('cache-reuse', 'ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=16', 'Uncached → Cached'),
    ]
    text = '''# 性能概要

このレポートは測定時点の性能スナップショットである。  
[測定環境・対象コミット](results/environment.md)、[測定方法と範囲](METHOD.md)、[今回の API 対応表](API.md) を参照。  
全 Unit / Scenario の時間・割り当て量と信頼区間は [詳細集計](results/all.txt)（[CSV](results/all.csv)）に分離している。  

## Scenario の代表値

N4096 の代表例を示し、全件・途中終了、キャッシュの初回費用を含む利用回数を区別する。  
時間は benchstat の中央値、倍率は同じ行の対象 / 対照であり、過去の測定とは比較しない。  
倍率が異なっていても検定が `~` の場合は差を断定しない。  

| 条件 | 対照 → 対象 | 対照 ns/op | 対象 ns/op | 倍率 | 同一測定内の検定 |
| --- | --- | ---: | ---: | ---: | --- |
'''
    for filename, key, label in selected:
        row = comparisons(ROOT/f'results/{filename}.csv')[key, 'sec/op']
        base, target = float(row[1]), float(row[3])
        result = '差は不明確' if row[5] == '~' else row[5]
        text += f'| {key.removeprefix("Scenario")} | {label} | {base*1e9:.2f} | {target*1e9:.2f} | {target/base:.3f} | {result} ({row[6]}) |\n'
    text += '''
## 読み方

軽い変換では、接続する段数が多いほど関数・反復・構築の費用が表れやすい。  
Early8 は必要な出力がそろうと上流も止まるため、入力件数が多くても全件は処理しない。  
Channel は送受信・goroutine 起動・終了処理を含み、入力容量で対照の絶対時間も変わる。  
Cache は初回保存の時間と要素数に比例する保持メモリを必要とし、上流負荷と利用回数によって再評価を省く効果が変わる。  
Warm の測定値は保存費用を含まないため、利用全体の判断には Reuse の同じ回数の Uncached / Cached を比較する。  

これらは整数と軽い算術を使った条件の結果で、外部 I/O や別の型・出力方式へ一般化しない。  
短い測定・OS・GC・温度等の影響があり、細かい順位や統計的に不明確な差は結論にしない。  
詳細な実装間比較は [Loop](results/pipeline-loop.txt)、[Functions](results/pipeline-functions.txt)、[FromChannel](results/channel-from.txt)、[往復変換](results/channel-roundtrip.txt)、[ToChannel](results/channel-to.txt)、[Cache 再利用](results/cache-reuse.txt) を参照。  

更新は専用の定期更新 Issue で現在の結果を置き換える。  
通常の機能変更 PR やリリースで追加測定を義務付けず、手順は [運用ガイド](../doc/ai/benchmark.md) にまとめる。  
'''
    (ROOT/'REPORT.md').write_text(text)


if __name__ == '__main__':
    if sys.argv[1] == 'validate':
        validate(Path(sys.argv[2]), int(sys.argv[3]))
    elif sys.argv[1] == 'report':
        report(Path(sys.argv[2]), *sys.argv[3:])
    else:
        raise SystemExit('use validate or report')
