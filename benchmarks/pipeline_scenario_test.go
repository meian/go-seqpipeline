package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func scenarioPipelineSeqLibrary(values []int, depth, limit int) []pair {
	s := seq.FromSlice(values)
	for j := 1; j < depth; j++ {
		s = s.Pipe(op.Map(mapValue))
	}
	s = s.Pipe(op.Filter(keepValue))
	if limit > 0 {
		s = s.Pipe(op.Take(limit).Operator)
	}
	var out []pair
	for v := range s {
		out = append(out, pair{value: v})
	}
	return out
}

func scenarioPipelineSeqLoop(values []int, depth, limit int) []pair {
	var out []pair
	for _, v := range values {
		for j := 1; j < depth; j++ {
			v = (v + 1) & 65535
		}
		if v%3 == 0 {
			continue
		}
		out = append(out, pair{value: v})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

func scenarioPipelineSeqFunctions(values []int, depth, limit int) []pair {
	var out []pair
	for _, v := range values {
		for j := 1; j < depth; j++ {
			v = mapValue(v)
		}
		if !keepValue(v) {
			continue
		}
		out = append(out, pair{value: v})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

func BenchmarkScenarioPipelineSeq(b *testing.B) {
	cases := []struct {
		n     int
		depth int
		mode  string
		limit int
	}{
		{64, 2, "All", 0},
		{64, 2, "Early8", 8},
		{64, 6, "All", 0},
		{64, 6, "Early8", 8},
		{4096, 2, "All", 0},
		{4096, 2, "Early8", 8},
		{4096, 6, "All", 0},
		{4096, 6, "Early8", 8},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioPipeline/kind=Seq/N=%d/depth=%d/mode=%s/impl=Pipeline", tc.n, tc.depth, tc.mode)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioPipelineSeqLibrary(values, tc.depth, tc.limit)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioPipeline/kind=Seq/N=%d/depth=%d/mode=%s/impl=Loop", tc.n, tc.depth, tc.mode)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioPipelineSeqLoop(values, tc.depth, tc.limit)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioPipeline/kind=Seq/N=%d/depth=%d/mode=%s/impl=Functions", tc.n, tc.depth, tc.mode)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioPipelineSeqFunctions(values, tc.depth, tc.limit)
			}
		})
	}
}

func TestScenarioPipelineSeq(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	want := []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}}
	cases := []struct {
		name string
		run  func([]int, int, int) []pair
	}{
		{"Library", scenarioPipelineSeqLibrary},
		{"Loop", scenarioPipelineSeqLoop},
		{"Functions", scenarioPipelineSeqFunctions},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.run(slices.Clone(values), 2, 0)
			if !slices.Equal(got, want) {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func scenarioPipelineSeq2Library(values []int, depth, limit int) []pair {
	s := seq.FromSeq2(slices.All(values))
	for j := 1; j < depth; j++ {
		s = s.Pipe2(op.Map2(func(k, v int) (int, int) { return k, mapValue(v) }))
	}
	s = s.Pipe2(op.Filter2(func(k, v int) bool { return keepValue(v) }))
	if limit > 0 {
		s = s.Pipe2(op.Take(limit).Operator2)
	}
	var out []pair
	for k, v := range s {
		out = append(out, pair{k, v})
	}
	return out
}

func scenarioPipelineSeq2Loop(values []int, depth, limit int) []pair {
	var out []pair
	for k, v := range values {
		for j := 1; j < depth; j++ {
			v = (v + 1) & 65535
		}
		if v%3 == 0 {
			continue
		}
		out = append(out, pair{k, v})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

func scenarioPipelineSeq2Functions(values []int, depth, limit int) []pair {
	var out []pair
	for k, v := range values {
		for j := 1; j < depth; j++ {
			v = mapValue(v)
		}
		if !keepValue(v) {
			continue
		}
		out = append(out, pair{k, v})
		if limit > 0 && len(out) == limit {
			break
		}
	}
	return out
}

func BenchmarkScenarioPipelineSeq2(b *testing.B) {
	cases := []struct {
		n     int
		depth int
		mode  string
		limit int
	}{
		{64, 2, "All", 0},
		{64, 2, "Early8", 8},
		{64, 6, "All", 0},
		{64, 6, "Early8", 8},
		{4096, 2, "All", 0},
		{4096, 2, "Early8", 8},
		{4096, 6, "All", 0},
		{4096, 6, "Early8", 8},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioPipeline/kind=Seq2/N=%d/depth=%d/mode=%s/impl=Pipeline", tc.n, tc.depth, tc.mode)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioPipelineSeq2Library(values, tc.depth, tc.limit)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioPipeline/kind=Seq2/N=%d/depth=%d/mode=%s/impl=Loop", tc.n, tc.depth, tc.mode)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioPipelineSeq2Loop(values, tc.depth, tc.limit)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioPipeline/kind=Seq2/N=%d/depth=%d/mode=%s/impl=Functions", tc.n, tc.depth, tc.mode)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioPipelineSeq2Functions(values, tc.depth, tc.limit)
			}
		})
	}
}

func TestScenarioPipelineSeq2(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	want := []pair{{0, 2}, {2, 4}, {3, 5}, {5, 7}}
	cases := []struct {
		name string
		run  func([]int, int, int) []pair
	}{
		{"Library", scenarioPipelineSeq2Library},
		{"Loop", scenarioPipelineSeq2Loop},
		{"Functions", scenarioPipelineSeq2Functions},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.run(slices.Clone(values), 2, 0)
			if !slices.Equal(got, want) {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}
