"""測定値を加工せず、欠落・不整合を検出して Markdown 化することを確認する。"""
import csv
import os
import signal
import subprocess
import time
import tempfile
import unittest
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
        comparison.base, comparison.target = comparison.target, comparison.base
        with self.assertRaisesRegex(ValueError, 'unit comparison must be Go / Library'):
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
                     'func BenchmarkMapGo(b *testing.B) {}\n'
                     'func TestMapGo(t *testing.T) {}\n'}

    def test_link_identifies_processing_at_measured_commit(self):
        sources = publish.SourceLinks(self.code, 'a'*40)
        self.assertEqual(sources.link(self.case, 'Go'),
                         '[Go](https://github.com/meian/go-seqpipeline/blob/'+'a'*40+
                         '/benchmarks/map_test.go#L1)')
        sources.verify([self.case])

    def test_missing_benchmark_or_test_entry_is_rejected(self):
        for prefix in ('Benchmark', 'Test'):
            code = {path: '\n'.join(line for line in source.splitlines()
                                   if not line.startswith('func '+prefix))
                    for path, source in self.code.items()}
            with self.subTest(prefix=prefix):
                with self.assertRaisesRegex(ValueError, 'missing processing function or entry'):
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
