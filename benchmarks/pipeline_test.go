package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func scenarioPipelineSeqLibrary(values []int, depth, limit int) []pair {
	s := seq.FromSlice(values)
	for j := 1; j < depth; j++ {
		s = s.Pipe(seq.Map(mapValue))
	}
	s = s.Pipe(seq.Filter(keepValue))
	if limit > 0 {
		s = s.Pipe(seq.Take(limit).Operator)
	}
	var out []pair
	for v := range s {
		out = append(out, pair{value: v})
	}
	return out
}

func BenchmarkScenarioPipelineSeqLibrary(b *testing.B) {
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=2/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=2/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=6/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=6/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 6, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLibrary(values, 6, 8)
		}
	})
}

func TestScenarioPipelineSeqLibrary(t *testing.T) {
	t.Run("N64Depth2All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLibrary(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth2Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLibrary(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLibrary(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLibrary(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLibrary(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLibrary(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLibrary(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLibrary(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLibrary(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLibrary(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLibrary(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
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

func BenchmarkScenarioPipelineSeqLoop(b *testing.B) {
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=2/mode=All/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=2/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=6/mode=All/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=6/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 6, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=All/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=All/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqLoop(values, 6, 8)
		}
	})
}

func TestScenarioPipelineSeqLoop(t *testing.T) {
	t.Run("N64Depth2All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLoop(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth2Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLoop(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLoop(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLoop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLoop(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLoop(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLoop(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLoop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqLoop(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLoop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqLoop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
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

func BenchmarkScenarioPipelineSeqFunctions(b *testing.B) {
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=2/mode=All/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=2/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=6/mode=All/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=64/depth=6/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 6, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=All/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=2/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=All/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq/N=4096/depth=6/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeqFunctions(values, 6, 8)
		}
	})
}

func TestScenarioPipelineSeqFunctions(t *testing.T) {
	t.Run("N64Depth2All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqFunctions(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth2Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqFunctions(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqFunctions(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqFunctions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqFunctions(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqFunctions(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqFunctions(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqFunctions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})

		}
		got := scenarioPipelineSeqFunctions(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqFunctions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{value: value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeqFunctions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func scenarioPipelineSeq2Library(values []int, depth, limit int) []pair {
	s := seq.FromSeq2(slices.All(values))
	for j := 1; j < depth; j++ {
		s = s.Pipe2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) }))
	}
	s = s.Pipe2(seq.Filter2(func(k, v int) bool { return keepValue(v) }))
	if limit > 0 {
		s = s.Pipe2(seq.Take(limit).Operator2)
	}
	var out []pair
	for k, v := range s {
		out = append(out, pair{k, v})
	}
	return out
}

func BenchmarkScenarioPipelineSeq2Library(b *testing.B) {
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=2/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=2/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=6/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=6/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 6, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=2/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=2/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=All/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=Early8/impl=Pipeline", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Library(values, 6, 8)
		}
	})
}

func TestScenarioPipelineSeq2Library(t *testing.T) {
	t.Run("N64Depth2All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Library(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth2Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Library(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Library(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Library(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Library(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Library(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Library(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Library(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Library(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Library(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Library(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
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

func BenchmarkScenarioPipelineSeq2Loop(b *testing.B) {
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=2/mode=All/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=2/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=6/mode=All/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=6/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 6, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=2/mode=All/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=2/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=All/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=Early8/impl=Loop", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Loop(values, 6, 8)
		}
	})
}

func TestScenarioPipelineSeq2Loop(t *testing.T) {
	t.Run("N64Depth2All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Loop(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth2Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Loop(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Loop(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Loop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Loop(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Loop(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Loop(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Loop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Loop(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Loop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Loop(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
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

func BenchmarkScenarioPipelineSeq2Functions(b *testing.B) {
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=2/mode=All/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=2/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=6/mode=All/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=64/depth=6/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 6, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=2/mode=All/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 2, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=2/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 2, 8)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=All/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 6, 0)
		}
	})
	b.Run("class=ScenarioPipeline/kind=Seq2/N=4096/depth=6/mode=Early8/impl=Functions", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioPipelineSeq2Functions(values, 6, 8)
		}
	})
}

func TestScenarioPipelineSeq2Functions(t *testing.T) {
	t.Run("N64Depth2All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Functions(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth2Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Functions(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6All", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Functions(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Depth6Early8", func(t *testing.T) {
		values := input(64)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Functions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Functions(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth2Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Functions(values, 2, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6All", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Functions(values, 6, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Depth6Early8", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Functions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for k, v := range values {
			value := (v + 1) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})

		}
		got := scenarioPipelineSeq2Functions(values, 2, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Functions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for k, v := range values {
			value := (v + 5) & 65535
			if value%3 == 0 {
				continue
			}
			want = append(want, pair{k, value})
			if len(want) == 8 {
				break
			}
		}
		got := scenarioPipelineSeq2Functions(values, 6, 8)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}
