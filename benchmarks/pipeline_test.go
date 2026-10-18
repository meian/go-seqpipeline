package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

// depth は Map の数 + Filter の 1 段。途中終了時はさらに Take を接続する。
func runScenarioPipelinePipeline(values []int, kind string, depth, limit int) []pair {
	var out []pair
	if kind == "Seq" {
		s := seq.FromSlice(values)
		for j := 1; j < depth; j++ {
			s = s.Pipe(seq.Map(mapValue))
		}
		s = s.Pipe(seq.Filter(keepValue))
		if limit > 0 {
			s = s.Pipe(seq.Take(limit).Operator)
		}
		for v := range s {
			out = append(out, pair{value: v})
		}
	} else {
		s := seq.FromSeq2(slices.All(values))
		for j := 1; j < depth; j++ {
			s = s.Pipe2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) }))
		}
		s = s.Pipe2(seq.Filter2(func(k, v int) bool { return keepValue(v) }))
		if limit > 0 {
			s = s.Pipe2(seq.Take(limit).Operator2)
		}
		for k, v := range s {
			out = append(out, pair{k, v})
		}
	}
	return out
}

func runScenarioPipelineLoop(values []int, kind string, depth, limit int) []pair {
	var out []pair
	for k, v := range values {
		for j := 1; j < depth; j++ {
			v = (v + 1) & 65535
		}
		if v%3 == 0 {
			continue
		}
		if kind == "Seq" {
			k = 0
		}
		out = append(out, pair{k, v})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

// 要素ごとの変換・判定を通常の関数に分け、インライン化を許す。
func runScenarioPipelineFunctions(values []int, kind string, depth, limit int) []pair {
	var out []pair
	for k, v := range values {
		for j := 1; j < depth; j++ {
			v = mapValue(v)
		}
		if !keepValue(v) {
			continue
		}
		if kind == "Seq" {
			k = 0
		}
		out = append(out, pair{k, v})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

// 変換回数から結果を直接計算し、測定対象の実装とは独立に照合する。
func pipelineExpected(values []int, kind string, depth, limit int) []pair {
	var out []pair
	for k, v := range values {
		value := (v + depth - 1) & 65535
		if value%3 == 0 {
			continue
		}
		if kind == "Seq" {
			k = 0
		}
		out = append(out, pair{k, value})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

func testScenarioPipeline(t *testing.T, run func([]int, string, int, int) []pair) {
	t.Helper()
	for _, n := range []int{0, 1, 64, 4096} {
		for _, kind := range []string{"Seq", "Seq2"} {
			for _, depth := range []int{2, 6} {
				for _, limit := range []int{0, 8} {
					t.Run(fmt.Sprintf("%s/N%d/D%d/L%d", kind, n, depth, limit), func(t *testing.T) {
						values := input(n)
						if got, want := run(values, kind, depth, limit), pipelineExpected(values, kind, depth, limit); !slices.Equal(got, want) {
							t.Fatalf("got %v; want %v", got, want)
						}
					})
				}
			}
		}
	}
}

func TestScenarioPipelineLoop(t *testing.T) { testScenarioPipeline(t, runScenarioPipelineLoop) }

func TestScenarioPipelineFunctions(t *testing.T) {
	testScenarioPipeline(t, runScenarioPipelineFunctions)
}

func TestScenarioPipelinePipeline(t *testing.T) { testScenarioPipeline(t, runScenarioPipelinePipeline) }

func benchmarkScenarioPipelineLoop(b *testing.B, values []int, kind string, depth, limit int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioPipelineLoop(values, kind, depth, limit)
	}
}

func benchmarkScenarioPipelineFunctions(b *testing.B, values []int, kind string, depth, limit int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioPipelineFunctions(values, kind, depth, limit)
	}
}

func benchmarkScenarioPipelinePipeline(b *testing.B, values []int, kind string, depth, limit int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioPipelinePipeline(values, kind, depth, limit)
	}
}

func BenchmarkScenarioPipeline(b *testing.B) {
	for _, kind := range []string{"Seq", "Seq2"} {
		for _, n := range []int{64, 4096} {
			values := input(n)
			for _, depth := range []int{2, 6} {
				for _, limit := range []int{0, 8} {
					mode := "All"
					if limit > 0 {
						mode = "Early8"
					}
					for _, implementation := range []struct {
						name  string
						entry func(*testing.B, []int, string, int, int)
					}{
						{"Loop", benchmarkScenarioPipelineLoop}, {"Functions", benchmarkScenarioPipelineFunctions}, {"Pipeline", benchmarkScenarioPipelinePipeline},
					} {
						b.Run(fmt.Sprintf("kind=%s/N=%d/depth=%d/mode=%s/impl=%s", kind, n, depth, mode, implementation.name), func(b *testing.B) { implementation.entry(b, values, kind, depth, limit) })
					}
				}
			}
		}
	}
}
