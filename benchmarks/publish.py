#!/usr/bin/env python3
"""benchstat CSV を公開用 Markdown に整形する。統計計算は行わない。"""
import csv
import datetime
import math
import platform
import re
import sys
from dataclasses import dataclass
from pathlib import Path

ROOT = Path(__file__).resolve().parent
METRICS = ('sec/op', 'B/op', 'allocs/op')


def validate(path, count):
    counts = {}
    text = path.read_text()
    for line in text.splitlines():
        if line.startswith('Benchmark') and ' ns/op' in line:
            name = line.split()[0]
            counts[name] = counts.get(name, 0) + 1
    if not counts or any(n != count for n in counts.values()) or '\nPASS\n' not in text:
        raise ValueError('missing or incomplete benchmark samples')
    print(f'{len(counts)} cases x {count} samples')
    return len(counts)


def key(name):
    parts = re.sub(r'-\d+$', '', name).split('/')
    dimensions = dict(part.split('=', 1) for part in parts[1:])
    return parts[0], tuple(sorted(dimensions.items()))


def csv_metrics(path):
    metric = None
    with path.open(newline='') as stream:
        for row in csv.reader(stream):
            if len(row) > 1 and row[0] == '' and row[1] in METRICS:
                metric = row[1]
            elif metric and row and row[0].startswith(('Unit/', 'Scenario')):
                yield key(row[0]), metric, row


@dataclass(frozen=True)
class Value:
    median: float
    interval: str


@dataclass
class Comparison:
    base: tuple
    target: tuple
    stats: dict


class Snapshot:
    def __init__(self, directory):
        self.values = {}
        self.seen = set()
        self.stats_seen = set()
        for case, metric, row in csv_metrics(directory/'all.csv'):
            if (case, metric) in self.values:
                raise ValueError(f'duplicate aggregate: {case} {metric}')
            self.values[case, metric] = Value(float(row[1]), row[2])
        if not self.values:
            raise ValueError('no aggregates')
        self.cases = list(dict.fromkeys(case for case, _ in self.values))
        for case in self.cases:
            if any((case, metric) not in self.values for metric in METRICS):
                raise ValueError(f'missing metric: {case}')
        self.comparisons = []
        for path in sorted(directory.glob('*.csv')):
            if path.name == 'all.csv':
                continue
            labels = None
            with path.open(newline='') as stream:
                for row in csv.reader(stream):
                    if len(row) == 7 and row[0] == '' and row[1] not in METRICS:
                        labels = row[1], row[3]
                        break
            if labels is None:
                raise ValueError(f'missing comparison columns: {path}')
            groups = {}
            for case, metric, row in csv_metrics(path):
                kind, dimensions = case
                dimensions = dict(dimensions)
                cases = [
                    (kind, tuple(sorted(dict(dimensions, impl=label).items())))
                    for label in labels
                ]
                for identity, offset in zip(cases, (1, 3)):
                    expected = self.values[identity, metric]
                    if not math.isclose(expected.median, float(row[offset]), rel_tol=1e-12, abs_tol=0):
                        raise ValueError(f'inconsistent median: {path} {identity}')
                    if expected.interval != row[offset+1]:
                        raise ValueError(f'inconsistent interval: {path} {identity}')
                comparison = groups.setdefault(case, Comparison(*cases, {}))
                if metric in comparison.stats:
                    raise ValueError(f'duplicate comparison metric: {path} {case}')
                comparison.stats[metric] = row[5], row[6]
            for comparison in groups.values():
                if set(comparison.stats) != set(METRICS):
                    raise ValueError('missing comparison statistics')
            self.comparisons.extend(groups.values())

    def cell(self, case, metric, interval=True):
        self.seen.add((case, metric))
        value = self.values[case, metric]
        number = value.median * (1e9 if metric == 'sec/op' else 1)
        # 極小の非ゼロ値も 0 に丸めない。
        text = f'{number:.6g}' if 0 < abs(number) < 0.000001 else f'{number:,.6f}'.rstrip('0').rstrip('.')
        return f'{text} ±{value.interval}' if interval else text

    def ratio(self, comparison):
        base = self.values[comparison.base, 'sec/op'].median
        target = self.values[comparison.target, 'sec/op'].median
        return f'{target/base:.3f}×' if base else '—（対照が 0）'

    def verify(self):
        if self.seen != set(self.values):
            raise ValueError(f'unpublished metrics: {set(self.values)-self.seen}')
        expected = {(c.base, c.target, metric) for c in self.comparisons for metric in METRICS}
        if self.stats_seen != expected:
            raise ValueError('unpublished comparison statistics')


def table(headers, rows):
    def escape(value):
        return str(value).replace('|', '&#124;').replace('\n', '<br>')
    lines = ['| '+' | '.join(headers)+' |', '| '+' | '.join('---' for _ in headers)+' |']
    lines.extend('| '+' | '.join(map(escape, row))+' |' for row in rows)
    return '\n'.join(lines)+'\n\n'


def condition(case, omit=()):
    dimensions = dict(case[1])
    labels = {'kind': '', 'N': 'N=', 'depth': '段数=', 'mode': '',
              'buffer': '入力容量=', 'work': '負荷=', 'reuse': '利用回数=', 'phase': ''}
    values = {'All': '全件', 'Early8': '先頭8件', 'Evaluate': '構築・評価',
              'Collect': '収集', 'Fill': '初回保存 Fill', 'Warm': '保存済み反復 Warm', 'Reuse': '保存費用込み Reuse'}
    order = list(labels)+sorted(set(dimensions)-set(labels)-{'api', 'impl'})
    return ' / '.join(labels.get(k, k+'=')+values.get(dimensions[k], dimensions[k])
                      for k in order if k in dimensions and k not in omit)


def preface(title):
    return f'''# {title}

[性能概要](../REPORT.md) / [測定条件](../METHOD.md) / [実行環境](environment.md)。  
値は中央値（`±` は 95% 信頼区間）で、割り当て量は保持メモリ量とは異なる。  

'''


def unit_category(api):
    if api.startswith(('Zip', 'Take', 'Skip')):
        return '件数制御・結合'
    if api in ('Cache', 'Cache2'):
        return 'キャッシュ'
    if 'Channel' in api:
        return 'Channel'
    if api.startswith(('From', 'To')):
        return '入力変換・出力収集'
    if '.Pipe' in api:
        return 'パイプ接続'
    if api in ('Map', 'Map2', 'Filter', 'Filter2', 'Keys', 'Values'):
        return '変換・選別・射影'
    return 'その他'


def condition_groups(cases):
    """共通条件を表の外へ出し、件数・段数などは独立した列にする。"""
    groups = {}
    for case in cases:
        groups.setdefault(condition(case, omit=('N', 'depth', 'buffer', 'reuse')), []).append(case)
    for group in groups.values():
        group.sort(key=lambda case: (
            *(int(dict(case[1]).get(k, 0)) for k in ('N', 'depth', 'buffer', 'reuse')),
            dict(case[1]).get('impl', ''),
        ))
    return groups.items()


def render_unit(snapshot):
    text = preface('API 別の性能')
    groups = {}
    for case in snapshot.cases:
        if case[0] == 'Unit':
            groups.setdefault(unit_category(dict(case[1])['api']), []).append(case)
    for group, cases in groups.items():
        cases.sort(key=lambda case: (dict(case[1])['api'], dict(case[1]).get('phase', ''),
                                     int(dict(case[1]).get('N', 0))))
        text += f'## {group}\n\n'
        # 同じ API の設定メソッドでも、キー既定値の有無は区別する。
        sections = {}
        for case in cases:
            detail = condition(case, omit=('N',))
            if dict(case[1])['api'].startswith('ZipKey.'):
                detail += ' / キー既定値あり'
            sections.setdefault(detail, []).append(case)
        for detail, section in sections.items():
            text += f'### {detail or "既定条件"}\n\n'
            rows = [[dict(case[1])['api'].replace('ZipKey.', 'Zip.', 1),
                     dict(case[1]).get('N', '—'),
                     *(snapshot.cell(case, metric) for metric in METRICS)] for case in section]
            text += table(['API', 'N', 'ns/op', 'B/op', 'allocs/op'], rows)
    return text.rstrip('\n')+'\n'


def render_scenario(snapshot):
    text = preface('利用場面別の性能')
    text += '''倍率は各表の右側の実装の時間 / 左側の実装の時間で、1 を超えると右側が遅い。  
検定の `~` は差が不明確で、増減率の正値は右側の時間・割り当ての増加を表す。  
Fill / Warm は保存費用込みの Reuse と計測範囲が異なる。  

'''
    groups = {}
    for case in snapshot.cases:
        if case[0] != 'Unit':
            groups.setdefault(case[0], []).append(case)
    order = {'ScenarioPipeline': 0, 'ScenarioChannel': 1, 'ScenarioCache': 2}
    for group in sorted(groups, key=lambda name: (order.get(name, 3), name)):
        text += f'## {group.removeprefix("Scenario")}\n\n'
        for detail, cases in condition_groups(groups[group]):
            text += f'### {detail or "入力条件別"}\n\n'
            axes = [k for k in ('N', 'depth', 'buffer', 'reuse') if any(k in dict(c[1]) for c in cases)]
            labels = {'N': 'N', 'depth': '段数', 'buffer': '入力容量', 'reuse': '利用回数'}
            headers = [labels[k] for k in axes]

            def fields(case):
                return [dict(case[1]).get(k, '—') for k in axes]

            pairs = [c for c in snapshot.comparisons if c.base in cases]
            by_implementations = {}
            for pair in pairs:
                labels = (dict(pair.base[1])['impl'], dict(pair.target[1])['impl'])
                by_implementations.setdefault(labels, []).append(pair)
            covered = set()
            for (base, target), comparisons in by_implementations.items():
                text += f'#### 時間：{base} と {target}\n\n'
                rows = []
                for c in comparisons:
                    covered.update((c.base, c.target))
                    rows.append([*fields(c.base), snapshot.cell(c.base, 'sec/op'),
                                 snapshot.cell(c.target, 'sec/op'), snapshot.ratio(c)])
                text += table([*headers, base+' ns/op', target+' ns/op', '倍率'], rows)
                text += f'<details>\n<summary>検定：{base} と {target}</summary>\n\n'
                rows = []
                for c in comparisons:
                    for metric in METRICS:
                        snapshot.stats_seen.add((c.base, c.target, metric))
                        delta, probability = c.stats[metric]
                        rows.append([*fields(c.base), metric, delta, probability])
                text += table([*headers, '指標', '増減率', '検定'], rows)
                text += '</details>\n\n'
            alone = [case for case in cases if case not in covered]
            if alone:
                text += '#### 時間\n\n'
                text += table([*headers, '実装', 'ns/op'],
                              [[*fields(case), dict(case[1]).get('impl', '—'),
                                snapshot.cell(case, 'sec/op')] for case in alone])
            text += '#### メモリ割り当て\n\n'
            text += table([*headers, '実装', 'B/op', 'allocs/op'],
                          [[*fields(case), dict(case[1]).get('impl', '—'),
                            snapshot.cell(case, 'B/op'), snapshot.cell(case, 'allocs/op')] for case in cases])
    return text.rstrip('\n')+'\n'


def render_overview(snapshot):
    text = '''# 性能概要

[API 別の全結果](results/unit.md) / [利用場面別の全結果](results/scenario.md) / [測定条件](METHOD.md) / [実行環境](results/environment.md)。  
時間・割り当て量は中央値で、時間比は各表の先頭の実装を 1 とする（`~` は差が不明確）。  

'''
    # 代表例の選択だけを固定し、詳細表は全ケースから自動生成する。
    choices = [
        ('ScenarioPipeline', {'kind': 'Seq', 'N': '4096', 'depth': '2', 'mode': 'All'}, 'Loop', 'Pipeline'),
        ('ScenarioChannel', {'N': '4096', 'buffer': '64'}, 'Direct', 'FromChannel'),
        ('ScenarioChannel', {'N': '4096', 'buffer': '0'}, 'Seq', 'ToChannel'),
        ('ScenarioCache', {'kind': 'Seq', 'N': '4096', 'work': '0', 'phase': 'Reuse', 'reuse': '1'}, 'Uncached', 'Cached'),
        ('ScenarioCache', {'kind': 'Seq2', 'N': '4096', 'work': '32', 'phase': 'Reuse', 'reuse': '16'}, 'Uncached', 'Cached'),
    ]
    for kind, dimensions, base, target in choices:
        for comparison in snapshot.comparisons:
            dims = dict(comparison.base[1])
            if (comparison.base[0] == kind and dims == dict(dimensions, impl=base)
                    and dict(comparison.target[1])['impl'] == target):
                title = {'ScenarioPipeline': 'パイプライン', 'FromChannel': 'Channel から読み込む',
                         'ToChannel': 'Channel へ書き出す'}.get(kind if kind == 'ScenarioPipeline' else target)
                if kind == 'ScenarioCache':
                    title = 'キャッシュ：軽い処理を 1 回使う' if dimensions['reuse'] == '1' else 'キャッシュ：重い処理を 16 回使う'
                text += f'## {title}\n\n'
                text += condition(comparison.base)+'。  \n\n'
                ratio = '~' if comparison.stats['sec/op'][0] == '~' else snapshot.ratio(comparison)
                rows = [[base, snapshot.cell(comparison.base, 'sec/op', False), '1×',
                         snapshot.cell(comparison.base, 'B/op', False)],
                        [target, snapshot.cell(comparison.target, 'sec/op', False), ratio,
                         snapshot.cell(comparison.target, 'B/op', False)]]
                text += table(['実装', 'ns/op', '時間比', 'B/op'], rows)
                break
    text += '''軽い変換では、パイプライン接続や channel 変換の追加費用が目立つ。  
キャッシュは初回保存の費用と保持メモリを要するが、重い上流を繰り返し使うと再評価を省ける。  
この整数処理での結果を、外部 I/O や他の型の速度保証には使わない。  
'''
    return text.rstrip('\n')+'\n'


def render(directory, destination):
    snapshot = Snapshot(directory)
    unit = render_unit(snapshot)
    scenario = render_scenario(snapshot)
    overview = render_overview(snapshot)
    snapshot.verify()
    destination.mkdir(parents=True, exist_ok=True)
    (destination/'unit.md').write_text(unit)
    (destination/'scenario.md').write_text(scenario)
    (destination.parent/'REPORT.md').write_text(overview)
    print(f'{len(snapshot.cases)} cases / {len(snapshot.values)} metrics published')
    return snapshot


def environment(workdir, benchtime, count, procs, destination):
    commit = (workdir/'commit.txt').read_text().strip()
    version = (workdir/'go-version.txt').read_text().strip()
    raw = (workdir/'raw.txt').read_text()
    cpu = next(line.removeprefix('cpu: ') for line in raw.splitlines() if line.startswith('cpu: '))
    validate(workdir/'raw.txt', int(count))
    env_names = ['GOOS', 'GOARCH', 'GOARM64', 'GOAMD64', 'GOGC', 'GOMEMLIMIT', 'GOEXPERIMENT']
    env_values = (workdir/'go-env.txt').read_text().splitlines()
    timestamp = datetime.datetime.now(datetime.timezone.utc).isoformat(timespec='seconds')
    settings = dict(zip(env_names, env_values))
    env = f'''# 実行環境

| 項目 | 値 |
| --- | --- |
| 測定日時（UTC） | {timestamp} |
| 対象コード（ライブラリ・測定） | [{commit[:12]}](https://github.com/meian/go-seqpipeline/commit/{commit}) |
| Go | {version} |
| OS / CPU | {platform.platform()} / {cpu} |
| GOMAXPROCS | {procs} |
| 測定時間 / 測定回数 | {benchtime} / {count} |
| GOGC / GOMEMLIMIT / GOEXPERIMENT | {settings.get('GOGC') or '既定値'} / {settings.get('GOMEMLIMIT') or '既定値'} / {settings.get('GOEXPERIMENT') or '既定値'} |

通常の最適化は有効で、CPU・温度・OS 負荷は固定していない。  
'''
    destination.write_text(env)


if __name__ == '__main__':
    if sys.argv[1] == 'validate':
        validate(Path(sys.argv[2]), int(sys.argv[3]))
    elif sys.argv[1] == 'render':
        render(Path(sys.argv[2]), Path(sys.argv[3]))
    elif sys.argv[1] == 'report':
        workdir = Path(sys.argv[2])
        destination = workdir/'published/results'
        snapshot = render(workdir, destination)
        if len(snapshot.cases) != validate(workdir/'raw.txt', int(sys.argv[4])):
            raise ValueError('raw and aggregate case counts differ')
        environment(workdir, *sys.argv[3:], destination/'environment.md')
    else:
        raise SystemExit('use validate, render or report')
