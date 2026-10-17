#!/usr/bin/env python3
"""Go benchmark の生データから中央値・ばらつき・対照との差を再集計する。"""
import argparse
import re
import statistics
from pathlib import Path


def summarize(raw, count):
    groups = {}
    pattern = re.compile(r"^Benchmark(\S+?)(?:-\d+)?\s+\d+\s+([\d.]+) ns/op\s+([\d.]+) B/op\s+([\d.]+) allocs/op$")
    for line in raw.splitlines():
        match = pattern.match(line.strip())
        if match:
            name, *numbers = match.groups()
            groups.setdefault(name, []).append(tuple(map(float, numbers)))
    expected = set()
    for kind in ('Seq', 'Seq2'):
        for n in (64, 4096):
            for depth in (2, 6):
                for mode in ('All', 'Early8'):
                    for impl in ('Loop', 'Functions', 'Pipeline'):
                        expected.add(f'Pipeline/{kind}/N{n}/D{depth}/{mode}/{impl}')
            for work in (0, 32):
                prefix = f'Cache/{kind}/N{n}/W{work}'
                expected.update((prefix+'/Fill', prefix+'/Warm'))
                for reuse in (1, 4, 16):
                    expected.update(f'{prefix}/R{reuse}/{impl}' for impl in ('Raw', 'Cached'))
    for n in (64, 4096):
        for impl in ('Seq', 'ToChannel', 'Direct', 'FromChannel', 'RoundTrip'):
            for capacity in ((0,) if impl in ('Seq', 'ToChannel') else (0, 64)):
                expected.add(f'Channel/N{n}/B{capacity}/{impl}')
    if set(groups) != expected:
        raise ValueError(f'missing={expected-set(groups)}, unexpected={set(groups)-expected}')
    rows = {}
    for name, samples in groups.items():
        if len(samples) != count:
            raise ValueError(f'{name}: expected {count} samples, got {len(samples)}')
        times = [s[0] for s in samples]
        rows[name] = dict(ns=statistics.median(times), low=min(times), high=max(times),
                          cv=100*statistics.stdev(times)/statistics.mean(times),
                          bytes=statistics.median(s[1] for s in samples),
                          allocs=statistics.median(s[2] for s in samples))
    return rows


def comparison(row, base):
    delta = row['ns']-base['ns']
    return f"{delta:+.2f} / {row['ns']/base['ns']:.3f} / {100*delta/base['ns']:+.1f}%"


def render(rows):
    lines = ['# 測定結果', '',
             '時間は ns/op の中央値、範囲は最小–最大、CV は標本標準偏差 / 平均。  ',
             '差・倍率・増減率は、それぞれの対照の中央値から計算する。  ',
             'B/op と allocs/op も測定回の中央値を示す。  ', '']
    for group in ('Pipeline', 'Channel', 'Cache'):
        lines += [f'## {group}', '', '| ケース | ns/op | 最小–最大 | CV % | B/op | allocs/op | 対照 | 差 ns / 倍率 / 増減率 |',
                  '| --- | ---: | ---: | ---: | ---: | ---: | --- | --- |']
        for name, row in rows.items():
            if not name.startswith(group+'/'):
                continue
            prefix, impl = name.rsplit('/', 1)
            bases = []
            if group == 'Pipeline' and impl == 'Pipeline':
                bases = [prefix+'/Loop', prefix+'/Functions']
            elif group == 'Channel':
                if impl in ('FromChannel', 'RoundTrip'):
                    bases = [prefix+'/Direct']
                elif impl == 'ToChannel':
                    bases = [prefix+'/Seq']
            elif group == 'Cache':
                if impl == 'Cached':
                    bases = [prefix+'/Raw']
                elif impl == 'Warm':
                    bases = [prefix+'/R1/Raw']
            labels = '<br>'.join(b.rsplit('/', 1)[-1] if impl != 'Warm' else 'R1/Raw' for b in bases) or '—'
            comparisons = '<br>'.join(comparison(row, rows[b]) for b in bases) or '—'
            lines.append(f"| {name[len(group)+1:]} | {row['ns']:.2f} | {row['low']:.2f}–{row['high']:.2f} | {row['cv']:.1f} | {row['bytes']:g} | {row['allocs']:g} | {labels} | {comparisons} |")
        lines.append('')
    return '\n'.join(lines)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('raw', type=Path)
    parser.add_argument('--count', type=int, default=5)
    parser.add_argument('--output', type=Path, required=True)
    args = parser.parse_args()
    rows = summarize(args.raw.read_text(), args.count)
    args.output.write_text(render(rows))
    print(f'{len(rows)} cases, {args.count} samples each -> {args.output}')


if __name__ == '__main__':
    main()
