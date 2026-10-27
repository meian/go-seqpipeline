"""測定値を加工せず、欠落・不整合を検出して Markdown 化することを確認する。"""
import csv
import os
import signal
import subprocess
import time
import tempfile
import unittest
from unittest import mock
from pathlib import Path

import publish


class PublishTest(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory(prefix='benchmark-publish-')
        self.addCleanup(self.temporary.cleanup)
        self.directory = Path(self.temporary.name)
        self.single = 'FutureGo/class=Unit/api=FutureAPI/N=8/phase=Evaluate/ref=FutureLoop/impl=Go-2'
        self.library = 'FutureLibrary/class=Unit/api=FutureAPI/N=8/phase=Evaluate/ref=FutureLoop/impl=Library-2'
        self.base = 'FutureLoop/class=ScenarioFuture/N=8/impl=Loop-2'
        self.target = 'FuturePipeline/class=ScenarioFuture/N=8/impl=Pipeline-2'
        self.alone = 'FutureWarm/class=ScenarioFuture/N=8/phase=Warm-2'
        self.all_rows = []
        self.comparison_rows = []
        self.unit_rows = []
        for metric in publish.METRICS:
            self.all_rows += [['', '-', ''], ['', metric, 'CI']]
            self.all_rows += [[name, value, '2%'] for name, value in
                             ((self.single, 1), (self.library, 2), (self.base, 2), (self.target, 4), (self.alone, 8))]
            self.comparison_rows += [['', 'Loop', '', 'Pipeline', '', '', ''],
                                     ['', metric, 'CI', metric, 'CI', 'vs base', 'P'],
                                     ['*/class=ScenarioFuture/N=8-2', 2, '2%', 4, '2%', '~', 'p=0.123 n=10']]
            self.unit_rows += [['', 'Go', '', 'Library', '', '', ''],
                               ['', metric, 'CI', metric, 'CI', 'vs base', 'P'],
                               ['*/class=Unit/api=FutureAPI/N=8/phase=Evaluate/ref=FutureLoop-2',
                                1, '2%', 2, '2%', '+100.00%', 'p=0.000 n=10']]
        self.write('unit.csv', self.unit_rows)
        self.write('all.csv', self.all_rows)
        self.write('future.csv', self.comparison_rows)

    def write(self, name, rows):
        with (self.directory/name).open('w', newline='') as stream:
            csv.writer(stream).writerows(rows)

    def test_render_covers_future_api_scenario_and_all_metrics(self):
        destination = self.directory/'output/results'
        snapshot = publish.render(self.directory, destination)
        self.assertEqual(len(snapshot.seen), 15)
        unit = (destination/'unit.md').read_text()
        scenario = (destination/'scenario.md').read_text()
        self.assertIn('## その他', unit)
        self.assertIn('### FutureAPI：構築・評価', unit)
        self.assertIn('Go：FutureLoop。', unit)
        self.assertIn('| 8 | 1,000,000,000 ±2% | 2,000,000,000 ±2% | 2.000× |', unit)
        self.assertIn('| 8 | Go | 1 ±2% | 1 ±2% |', unit)
        self.assertIn('| 8 | Library | 2 ±2% | 2 ±2% |', unit)
        self.assertIn('| 8 | sec/op | +100.00% | p=0.000 n=10 |', unit)
        self.assertIn('## Future', scenario)
        self.assertIn('2.000×', scenario)
        self.assertIn('| 8 | sec/op | ~ | p=0.123 n=10 |', scenario)
        self.assertIn('保存済み反復 Warm', scenario)
        self.assertIn('| 8 | Pipeline | 4 ±2% | 4 ±2% |', scenario)

    def compose_rows(self):
        all_rows, comparisons = [], []
        for metric in publish.METRICS:
            all_rows += [['', '-', ''], ['', metric, 'CI']]
            comparisons += [['', 'Pipe', '', 'Composed', '', '', ''],
                            ['', metric, 'CI', metric, 'CI', 'vs base', 'P']]
            for api_index, api in enumerate(publish.COMPOSE_APIS, 1):
                for depth in (2, 4):
                    for scope in ('Build', 'Reuse'):
                        for n in (64, 4096):
                            case = (f'class=Unit/api={api}/N={n}/phase=Evaluate/ref=DirectPipe'
                                    f'/depth={depth}/scope={scope}')
                            value = api_index*10000+depth*1000+n+(scope == 'Reuse')
                            all_rows += [[f'{api}/{case}/impl={impl}-2', v, '2%']
                                         for impl, v in (('Pipe', value), ('Composed', value*2))]
                            comparisons += [[f'*/{case}-2', value, '2%', value*2, '2%',
                                             '+100.00%', 'p=0.000 n=10']]
        return all_rows, comparisons

    def test_compose_compares_only_equal_depth_scope_and_input(self):
        all_rows, comparisons = self.compose_rows()
        self.write('all.csv', self.all_rows+all_rows)
        self.write('unit-compose.csv', comparisons)
        destination = self.directory/'output/results'
        snapshot = publish.render(self.directory, destination)
        unit = (destination/'unit.md').read_text()
        self.assertEqual(len(snapshot.seen), 207)
        self.assertEqual(unit.count('| N | Pipe ns/op | Composed ns/op | 倍率 |'), 16)
        for api in publish.COMPOSE_APIS:
            for depth in (2, 4):
                for scope in ('構築込み', '構築済みを利用'):
                    self.assertIn(f'### {api}：段数={depth} / {scope}', unit)
        self.assertEqual(unit.count('2.000×'), 33)
        # 段数・構築範囲の片方を別条件に変えると、中央値の不一致で拒否する。
        for dimension, changed in (('depth=2', 'depth=4'), ('scope=Build', 'scope=Reuse')):
            rows = [row[:] for row in comparisons]
            rows[2][0] = rows[2][0].replace(dimension, changed)
            self.write('unit-compose.csv', rows)
            with self.subTest(dimension=dimension):
                with self.assertRaisesRegex(ValueError, 'inconsistent median'):
                    publish.Snapshot(self.directory)

    def test_compose_partial_update_preserves_other_results_and_metadata(self):
        all_rows, comparisons = self.compose_rows()
        self.write('all.csv', all_rows)
        self.write('unit-compose.csv', comparisons)
        (self.directory/'unit.csv').unlink()
        (self.directory/'future.csv').unlink()
        (self.directory/'commit.txt').write_text('a'*40+'\n')
        (self.directory/'go-version.txt').write_text('go version test\n')
        (self.directory/'go-env.txt').write_text('darwin\narm64\nv8.0\n\n\n\n\n')
        cases = publish.Snapshot(self.directory).cases
        raw = 'cpu: Test CPU\n'+''.join(
            'Benchmark/class='+case[0]+'/'+ '/'.join(k+'='+v for k, v in case[1])+
            '-2 1 1 ns/op 0 B/op 0 allocs/op\n' for case in cases)+'PASS\n'
        (self.directory/'raw.txt').write_text(raw)
        sources = publish.SourceLinks({'benchmarks/compose_test.go':
                                       (publish.ROOT/'compose_test.go').read_text()}, 'a'*40)
        original = (publish.ROOT/'results/unit.md').read_text()
        untouched = publish.replace_compose_sections(original, 'COMPOSE\n')
        destination = self.directory/'partial/results'
        with mock.patch.object(publish.SourceLinks, 'from_commit', return_value=sources):
            publish.render_compose(self.directory, destination, '500ms', '1', '2')
        updated = (destination/'unit.md').read_text()
        self.assertEqual(publish.replace_compose_sections(updated, 'COMPOSE\n'), untouched)
        self.assertIn('500ms / 1', updated)
        self.assertIn('Test CPU', updated)
        self.assertIn('go version test', updated)
        self.assertIn('/blob/'+'a'*40+'/benchmarks/compose_test.go#L', updated)
        self.assertEqual(updated.count('### Compose / Compose2 の測定条件'), 1)
        self.assertEqual(list(destination.iterdir()), [destination/'unit.md'])
        self.assertFalse((destination.parent/'REPORT.md').exists())
        # 全件 CSV の一つの実装を欠落させても、部分公開を進めない。
        self.write('all.csv', [row for row in all_rows
                              if not (row and str(row[0]).endswith('scope=Build/impl=Pipe-2')
                                      and '/api=Compose.Seq/N=64/' in str(row[0])
                                      and '/depth=2/' in str(row[0]))])
        (self.directory/'unit-compose.csv').unlink()
        with self.assertRaisesRegex(ValueError, 'exactly the 64'):
            publish.render_compose(self.directory, destination, '500ms', '1', '2')
        self.assertEqual((destination/'unit.md').read_text(), updated)

    def test_compose_sections_must_exist_and_be_contiguous(self):
        for current in ('# Empty\n', '### Compose.Seq：old\n### Map：keep\n### Compose2.Seq：old\n'):
            with self.subTest(current=current):
                with self.assertRaisesRegex(ValueError, 'Compose sections'):
                    publish.replace_compose_sections(current, 'new\n')

    def test_result_tables_keep_conditions_and_comparisons_separate(self):
        snapshot = publish.Snapshot(self.directory)
        # 概要の代表ケースも、同じ条件を持つ比較から自動生成する。
        comparison = next(c for c in snapshot.comparisons if c.base[0] != 'Unit')
        dimensions = {'kind': 'Seq', 'N': '4096', 'depth': '2', 'mode': 'All'}
        for label, attribute in [('Loop', 'base'), ('Pipeline', 'target')]:
            original = getattr(comparison, attribute)
            case = ('ScenarioPipeline', tuple(sorted(dict(dimensions, impl=label).items())))
            snapshot.cases.append(case)
            for metric in publish.METRICS:
                snapshot.values[case, metric] = snapshot.values[original, metric]
            setattr(comparison, attribute, case)
        overview = publish.render_overview(snapshot)
        self.assertIn('| Loop | 2,000,000,000 | 1× | 2 |', overview)
        self.assertIn('| Pipeline | 4,000,000,000 | ~ | 4 |', overview)
        for text in (overview, publish.render_unit(snapshot), publish.render_scenario(snapshot)):
            for line in text.splitlines():
                if line.startswith('|'):
                    self.assertNotIn('→', line)
                    self.assertNotIn('<br>', line)
                    self.assertNotIn(' / ', line)
                    self.assertLessEqual(len(line.strip('|').split('|')), 5)
        self.assertIn('<summary>検定：Loop と Pipeline</summary>', publish.render_scenario(snapshot))

    def test_channel_capacity_is_shown_only_for_channel_input(self):
        snapshot = publish.Snapshot(self.directory)
        original = next(c for c in snapshot.comparisons if c.base[0] == 'ScenarioFuture')
        values = snapshot.values
        snapshot.cases = []
        snapshot.values = {}
        snapshot.comparisons = []
        for base, target in [('Seq', 'ToChannel'), ('Direct', 'FromChannel')]:
            cases = [publish.key(f'ScenarioChannel/N=4096/buffer=0/impl={label}')
                     for label in (base, target)]
            for case, source in zip(cases, (original.base, original.target)):
                snapshot.cases.append(case)
                for metric in publish.METRICS:
                    snapshot.values[case, metric] = values[source, metric]
            snapshot.comparisons.append(publish.Comparison(*cases, original.stats.copy()))
        overview = publish.render_overview(snapshot)
        self.assertIn('## Channel へ書き出す', overview)
        self.assertNotIn('容量', overview)
        detail = publish.render_scenario(snapshot)
        seq_input, channel_input = detail.split('### Seq 入力')[1].split('### Channel 入力')
        self.assertNotIn('入力容量', seq_input)
        self.assertIn('| 4096 | 2,000,000,000 ±2% | 4,000,000,000 ±2% | 2.000× |', seq_input)
        self.assertIn('| N | 入力容量 | Direct ns/op | FromChannel ns/op | 倍率 |', channel_input)
        self.assertIn('| 4096 | 0 | 2,000,000,000 ±2% | 4,000,000,000 ±2% | 2.000× |', channel_input)
        snapshot.verify()

    def test_public_reports_exclude_maintainer_instructions(self):
        destination = self.directory/'output/results'
        publish.render(self.directory, destination)
        for path in [destination/'unit.md', destination/'scenario.md', destination.parent/'REPORT.md']:
            with self.subTest(path=path.name):
                for term in ('raw', 'CSV', 'Git 管理', '定期更新', '世代管理', 'doc/ai', '生成手順'):
                    self.assertNotIn(term, path.read_text())

    def test_environment_contains_conditions_without_workflow(self):
        (self.directory/'commit.txt').write_text('a'*40+'\n')
        (self.directory/'go-version.txt').write_text('go version test\n')
        (self.directory/'go-env.txt').write_text('darwin\narm64\nv8.0\n\n\n\n\n')
        (self.directory/'raw.txt').write_text('cpu: Test CPU\nBenchmarkUnit/api=X-2 1 1 ns/op 0 B/op 0 allocs/op\nPASS\n')
        destination = self.directory/'environment.md'
        publish.environment(self.directory, '100ms', '1', '2', destination)
        text = destination.read_text()
        self.assertIn('Test CPU', text)
        self.assertIn('go version test', text)
        self.assertIn('100ms / 1', text)
        for term in ('raw', 'CSV', 'Git 管理', '後続コミット', '一時ディレクトリ', '世代管理'):
            self.assertNotIn(term, text)

    def test_missing_unit_comparison_is_rejected(self):
        (self.directory/'unit.csv').unlink()
        with self.assertRaisesRegex(ValueError, 'missing unit comparisons'):
            publish.render_unit(publish.Snapshot(self.directory))

    def test_unit_comparison_with_wrong_implementation_is_rejected(self):
        snapshot = publish.Snapshot(self.directory)
        comparison = next(c for c in snapshot.comparisons if c.base[0] == 'Unit')
        for api, base, target, expected in (
            ('FutureAPI', 'Library', 'Go', 'Go / Library'),
            ('FutureAPI', 'Pipe', 'Composed', 'Go / Library'),
            ('Compose.Seq', 'Go', 'Library', 'Pipe / Composed'),
        ):
            with self.subTest(api=api, base=base, target=target):
                dimensions = dict(comparison.base[1], api=api)
                comparison.base = ('Unit', tuple(sorted(dict(dimensions, impl=base).items())))
                comparison.target = ('Unit', tuple(sorted(dict(dimensions, impl=target).items())))
                with self.assertRaisesRegex(ValueError, 'unit comparison must be '+expected):
                    publish.render_unit(snapshot)

    def test_missing_metric_is_rejected(self):
        self.write('all.csv', self.all_rows[:-1])
        with self.assertRaisesRegex(ValueError, 'missing metric'):
            publish.Snapshot(self.directory)

    def test_duplicate_aggregate_is_rejected(self):
        self.write('all.csv', self.all_rows+[self.all_rows[-1]])
        with self.assertRaisesRegex(ValueError, 'duplicate aggregate'):
            publish.Snapshot(self.directory)

    def test_comparison_median_and_interval_must_match(self):
        for column, replacement, error in ((1, 3, 'median'), (2, '9%', 'interval')):
            with self.subTest(column=column):
                rows = [row[:] for row in self.comparison_rows]
                rows[2][column] = replacement
                self.write('future.csv', rows)
                with self.assertRaisesRegex(ValueError, error):
                    publish.Snapshot(self.directory)

    def test_missing_comparison_statistics_is_rejected(self):
        self.write('future.csv', self.comparison_rows[:-1])
        with self.assertRaisesRegex(ValueError, 'missing comparison statistics'):
            publish.Snapshot(self.directory)

    def test_unpublished_metric_is_rejected(self):
        snapshot = publish.Snapshot(self.directory)
        publish.render_unit(snapshot)
        with self.assertRaisesRegex(ValueError, 'unpublished metrics'):
            snapshot.verify()

    def test_unpublished_comparison_statistics_is_rejected(self):
        snapshot = publish.Snapshot(self.directory)
        for case in snapshot.cases:
            for metric in publish.METRICS:
                snapshot.cell(case, metric)
        with self.assertRaisesRegex(ValueError, 'unpublished comparison statistics'):
            snapshot.verify()

    def test_table_escapes_markdown_delimiters(self):
        self.assertIn('A&#124;B<br>C', publish.table(['heading'], [['A|B\nC']]))


class SourceLinksTest(unittest.TestCase):
    def setUp(self):
        self.case = publish.key('Unit/api=Map/N=64/phase=Evaluate/ref=MapLoop/impl=Go')
        self.code = {'benchmarks/map_test.go':
                     'func mapGo(values []int) int64 { return 0 }\n'
                     'func BenchmarkMap(b *testing.B) {}\n'
                     'func TestMap(t *testing.T) {}\n'}

    def test_link_identifies_processing_at_measured_commit(self):
        sources = publish.SourceLinks(self.code, 'a'*40)
        self.assertEqual(sources.link(self.case, 'Go'),
                         '[Go](https://github.com/meian/go-seqpipeline/blob/'+'a'*40+
                         '/benchmarks/map_test.go#L1)')
        sources.verify([self.case])

    def test_grouped_entries_link_each_processing_implementation(self):
        for condition, suffix, implementations in [
            ('Unit/api=Map/phase=Evaluate', 'Map', ('Library', 'Go')),
            ('Unit/api=Compose2.Seq2/phase=Evaluate/depth=4/scope=Build', 'Compose2Seq2', ('Composed', 'Pipe')),
            ('ScenarioPipeline/kind=Seq', 'ScenarioPipelineSeq', ('Pipeline', 'Loop', 'Functions')),
            ('ScenarioPipeline/kind=Seq2', 'ScenarioPipelineSeq2', ('Pipeline', 'Loop', 'Functions')),
            ('ScenarioChannel', 'ScenarioChannel', ('Seq', 'ToChannel', 'Direct', 'FromChannel', 'RoundTrip')),
            ('ScenarioCache/kind=Seq/phase=Reuse', 'ScenarioCacheSeqReuse', ('Cached', 'Uncached')),
            ('ScenarioCache/kind=Seq2/phase=Reuse', 'ScenarioCacheSeq2Reuse', ('Cached', 'Uncached')),
        ]:
            cases = [publish.key(condition+'/impl='+implementation) for implementation in implementations]
            source = ''.join('func '+publish.processing_function(case)+'() {}\n' for case in cases)
            source += f'func Benchmark{suffix}(b *testing.B) {{}}\nfunc Test{suffix}(t *testing.T) {{}}\n'
            sources = publish.SourceLinks({'benchmarks/example_test.go': source}, 'a'*40)
            with self.subTest(condition=condition):
                for line, (case, implementation) in enumerate(zip(cases, implementations), 1):
                    self.assertEqual(sources.link(case, implementation),
                                     f'[{implementation}](https://github.com/meian/go-seqpipeline/blob/'+
                                     'a'*40+f'/benchmarks/example_test.go#L{line})')
                sources.verify(cases)

    def test_missing_benchmark_or_test_entry_is_rejected(self):
        for prefix in ('Benchmark', 'Test'):
            code = {path: '\n'.join(line for line in source.splitlines()
                                   if not line.startswith('func '+prefix))
                    for path, source in self.code.items()}
            with self.subTest(prefix=prefix):
                with self.assertRaisesRegex(ValueError, 'missing processing function or entry'):
                    publish.SourceLinks(code, 'a'*40).link(self.case, 'Go')

    def test_missing_processing_function_is_rejected(self):
        code = {'benchmarks/map_test.go':
                'func BenchmarkMap(b *testing.B) {}\nfunc TestMap(t *testing.T) {}\n'}
        with self.assertRaisesRegex(ValueError, 'missing processing function or entry: mapGo'):
            publish.SourceLinks(code, 'a'*40).link(self.case, 'Go')

    def test_unlinked_processing_is_rejected(self):
        with self.assertRaisesRegex(ValueError, 'unlinked processing functions'):
            publish.SourceLinks(self.code, 'a'*40).verify([self.case])

    def test_scenario_names_identify_implementation_or_phase(self):
        for name, function in [
            ('ScenarioChannel/N=64/buffer=0/impl=Direct', 'scenarioChannelDirect'),
            ('ScenarioPipeline/kind=Seq/N=64/depth=2/mode=All/impl=Loop', 'scenarioPipelineSeqLoop'),
            ('ScenarioCache/kind=Seq/N=64/work=0/phase=Fill', 'scenarioCacheSeqFill'),
            ('ScenarioCache/kind=Seq/N=64/work=0/phase=Warm', 'scenarioCacheSeqWarm'),
            ('ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=4/impl=Cached', 'scenarioCacheSeqCached'),
        ]:
            with self.subTest(name=name):
                self.assertEqual(publish.processing_function(publish.key(name)), function)

    def test_cache_fill_and_warm_keep_separate_entries(self):
        for kind in ('Seq', 'Seq2'):
            for phase in ('Fill', 'Warm'):
                with self.subTest(kind=kind, phase=phase):
                    case = publish.key(f'ScenarioCache/kind={kind}/phase={phase}')
                    self.assertEqual(publish.entry_suffix(case), f'ScenarioCache{kind}{phase}')

    def test_compose_names_identify_depth_scope_and_connection(self):
        for api, stem in [('Compose.Seq', 'composeSeq'), ('Compose.Seq2', 'composeSeq2'),
                          ('Compose2.Seq', 'compose2Seq'), ('Compose2.Seq2', 'compose2Seq2')]:
            for depth in (2, 4):
                for scope in ('Build', 'Reuse'):
                    for impl in ('Pipe', 'Composed'):
                        with self.subTest(api=api, depth=depth, scope=scope, impl=impl):
                            name = (f'Unit/api={api}/N=64/phase=Evaluate/ref=DirectPipe'
                                    f'/depth={depth}/scope={scope}/impl={impl}')
                            self.assertEqual(publish.processing_function(publish.key(name)),
                                             f'{stem}Depth{depth}{scope}{impl}')


class MeasureCleanupTest(unittest.TestCase):
    def exercise(self, interrupted):
        with tempfile.TemporaryDirectory(prefix='benchmark-cleanup-') as directory:
            directory = Path(directory)
            executable = directory/'go'
            marker = directory/'cache-path'
            executable.write_text('#!/bin/sh\nprintf "%s" "$GOCACHE" > "$MARKER"\n'+
                                  ('exec sleep 30\n' if interrupted else 'exit 2\n'))
            executable.chmod(0o755)
            env = dict(os.environ, GO_BINARY=str(executable), BENCHSTAT='true', MARKER=str(marker))
            env.pop('GOCACHE', None)
            process = subprocess.Popen(['sh', str(publish.ROOT/'measure.sh')], env=env,
                                       start_new_session=True, stdout=subprocess.DEVNULL,
                                       stderr=subprocess.DEVNULL)
            try:
                deadline = time.monotonic()+5
                while not marker.exists() and time.monotonic() < deadline:
                    time.sleep(0.01)
                self.assertTrue(marker.exists())
                temporary = Path(marker.read_text()).parent
                if interrupted:
                    os.killpg(process.pid, signal.SIGTERM)
                self.assertNotEqual(process.wait(timeout=5), 0)
                self.assertFalse(temporary.exists())
            finally:
                if process.poll() is None:
                    os.killpg(process.pid, signal.SIGKILL)
                    process.wait()

    def test_error_removes_intermediate_directory(self):
        self.exercise(False)

    def test_interrupt_removes_intermediate_directory(self):
        self.exercise(True)


if __name__ == '__main__':
    unittest.main()
