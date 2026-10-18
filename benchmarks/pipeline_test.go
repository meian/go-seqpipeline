package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

// depth は Map の数 + Filter の 1 段。途中終了時はさらに Take を接続する。
func pipeline(values []int, kind string, depth, limit int) []pair {
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

func baseline(values []int, kind string, depth, limit int) []pair {
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
func functionCalls(values []int, kind string, depth, limit int) []pair {
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

func TestPipelineEquivalent(t *testing.T) {
	for _, n := range []int{0, 1, 64, 4096} {
		for _, kind := range []string{"Seq", "Seq2"} {
			for _, depth := range []int{2, 6} {
				for _, limit := range []int{0, 8} {
					name := fmt.Sprintf("%s/N%d/D%d/L%d", kind, n, depth, limit)
					t.Run(name, func(t *testing.T) {
						want := baseline(input(n), kind, depth, limit)
						for _, got := range [][]pair{functionCalls(input(n), kind, depth, limit), pipeline(input(n), kind, depth, limit)} {
							if !slices.Equal(got, want) {
								t.Fatalf("got %v; want %v", got, want)
							}
						}
					})
				}
			}
		}
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
					for _, impl := range []string{"Loop", "Functions", "Pipeline"} {
						b.Run(fmt.Sprintf("kind=%s/N=%d/depth=%d/mode=%s/impl=%s", kind, n, depth, mode, impl), func(b *testing.B) {
							b.ReportAllocs()
							for b.Loop() {
								switch impl {
								case "Pipeline":
									output = pipeline(values, kind, depth, limit)
								case "Functions":
									output = functionCalls(values, kind, depth, limit)
								default:
									output = baseline(values, kind, depth, limit)
								}
							}
						})
					}
				}
			}
		}
	}
}
