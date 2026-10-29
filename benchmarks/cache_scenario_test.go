package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func cacheValue(v, work int) int {
	for j := 0; j < work; j++ {
		v = (v*1664525 + 1013904223) & 65535
	}
	return v
}

func cacheSeq(values []int, work int) seq.Seq[int] {
	return seq.FromSlice(values).
		Pipe(op.Map(func(v int) int { return cacheValue(v, work) }))
}

func cacheSeq2(values []int, work int) seq.Seq2[int, int] {
	return seq.FromSeq2(slices.All(values)).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, cacheValue(v, work) }))
}

var scenarioCacheSeqResult seq.Seq[int]

func scenarioCacheSeqFill(values []int, work int) seq.Seq[int] {
	source := cacheSeq(values, work).Pipe(op.Cache)
	// 最初の値を返す前に全件保存されるため、保存済み結果の走査は先頭で止める。
	for range source {
		break
	}
	return source
}

func BenchmarkScenarioCacheSeqFill(b *testing.B) {
	cases := []struct {
		n    int
		work int
	}{
		{64, 0},
		{64, 32},
		{4096, 0},
		{4096, 32},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq/N=%d/work=%d/phase=Fill", tc.n, tc.work)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				scenarioCacheSeqResult = scenarioCacheSeqFill(values, tc.work)
			}
		})
	}
}

func TestScenarioCacheSeqFill(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	cached := scenarioCacheSeqFill(values, 0)
	clear(values)
	got := seq.ToSlice(cached)
	want := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func scenarioCacheSeqWarm(source seq.Seq[int]) int64 {
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkScenarioCacheSeqWarm(b *testing.B) {
	cases := []struct {
		n    int
		work int
	}{
		{64, 0},
		{64, 32},
		{4096, 0},
		{4096, 32},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq/N=%d/work=%d/phase=Warm", tc.n, tc.work)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			source := scenarioCacheSeqFill(values, tc.work)
			b.ReportAllocs()
			for b.Loop() {
				checksum = scenarioCacheSeqWarm(source)
			}
		})
	}
}

func TestScenarioCacheSeqWarm(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	cached := scenarioCacheSeqFill(values, 0)
	got := scenarioCacheSeqWarm(cached)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func scenarioCacheSeqCached(values []int, work int, reuse int) int64 {
	source := cacheSeq(values, work).Pipe(op.Cache)
	var sum int64
	for r := 0; r < reuse; r++ {
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func scenarioCacheSeqUncached(values []int, work int, reuse int) int64 {
	source := cacheSeq(values, work)
	var sum int64
	for r := 0; r < reuse; r++ {
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkScenarioCacheSeqReuse(b *testing.B) {
	cases := []struct {
		n     int
		work  int
		reuse int
	}{
		{64, 0, 1},
		{64, 0, 4},
		{64, 0, 16},
		{64, 32, 1},
		{64, 32, 4},
		{64, 32, 16},
		{4096, 0, 1},
		{4096, 0, 4},
		{4096, 0, 16},
		{4096, 32, 1},
		{4096, 32, 4},
		{4096, 32, 16},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq/N=%d/work=%d/phase=Reuse/reuse=%d/impl=Cached", tc.n, tc.work, tc.reuse)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = scenarioCacheSeqCached(values, tc.work, tc.reuse)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq/N=%d/work=%d/phase=Reuse/reuse=%d/impl=Uncached", tc.n, tc.work, tc.reuse)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = scenarioCacheSeqUncached(values, tc.work, tc.reuse)
			}
		})
	}
}

func TestScenarioCacheSeqReuse(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		work   int
		passes int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6}, want: 42, work: 0, passes: 2},
	}
	for _, tc := range cases {
		t.Run("Cached", func(t *testing.T) {
			t.Parallel()

			got := scenarioCacheSeqCached(slices.Clone(tc.values), tc.work, tc.passes)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Uncached", func(t *testing.T) {
			t.Parallel()

			got := scenarioCacheSeqUncached(slices.Clone(tc.values), tc.work, tc.passes)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}

var scenarioCacheSeq2Result seq.Seq2[int, int]

func scenarioCacheSeq2Fill(values []int, work int) seq.Seq2[int, int] {
	source := cacheSeq2(values, work).Pipe2(op.Cache2)
	for range source {
		break
	}
	return source
}

func BenchmarkScenarioCacheSeq2Fill(b *testing.B) {
	cases := []struct {
		n    int
		work int
	}{
		{64, 0},
		{64, 32},
		{4096, 0},
		{4096, 32},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq2/N=%d/work=%d/phase=Fill", tc.n, tc.work)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				scenarioCacheSeq2Result = scenarioCacheSeq2Fill(values, tc.work)
			}
		})
	}
}

func TestScenarioCacheSeq2Fill(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	cached := scenarioCacheSeq2Fill(values, 0)
	clear(values)
	var got []pair
	for k, v := range cached {
		got = append(got, pair{k, v})
	}
	want := []pair{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func scenarioCacheSeq2Warm(source seq.Seq2[int, int]) int64 {
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkScenarioCacheSeq2Warm(b *testing.B) {
	cases := []struct {
		n    int
		work int
	}{
		{64, 0},
		{64, 32},
		{4096, 0},
		{4096, 32},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq2/N=%d/work=%d/phase=Warm", tc.n, tc.work)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			source := scenarioCacheSeq2Fill(values, tc.work)
			b.ReportAllocs()
			for b.Loop() {
				checksum = scenarioCacheSeq2Warm(source)
			}
		})
	}
}

func TestScenarioCacheSeq2Warm(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	cached := scenarioCacheSeq2Fill(values, 0)
	got := scenarioCacheSeq2Warm(cached)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func scenarioCacheSeq2Cached(values []int, work int, reuse int) int64 {
	source := cacheSeq2(values, work).Pipe2(op.Cache2)
	var sum int64
	for r := 0; r < reuse; r++ {
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func scenarioCacheSeq2Uncached(values []int, work int, reuse int) int64 {
	source := cacheSeq2(values, work)
	var sum int64
	for r := 0; r < reuse; r++ {
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func BenchmarkScenarioCacheSeq2Reuse(b *testing.B) {
	cases := []struct {
		n     int
		work  int
		reuse int
	}{
		{64, 0, 1},
		{64, 0, 4},
		{64, 0, 16},
		{64, 32, 1},
		{64, 32, 4},
		{64, 32, 16},
		{4096, 0, 1},
		{4096, 0, 4},
		{4096, 0, 16},
		{4096, 32, 1},
		{4096, 32, 4},
		{4096, 32, 16},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq2/N=%d/work=%d/phase=Reuse/reuse=%d/impl=Cached", tc.n, tc.work, tc.reuse)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = scenarioCacheSeq2Cached(values, tc.work, tc.reuse)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioCache/kind=Seq2/N=%d/work=%d/phase=Reuse/reuse=%d/impl=Uncached", tc.n, tc.work, tc.reuse)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = scenarioCacheSeq2Uncached(values, tc.work, tc.reuse)
			}
		})
	}
}

func TestScenarioCacheSeq2Reuse(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		work   int
		passes int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6}, want: 72, work: 0, passes: 2},
	}
	for _, tc := range cases {
		t.Run("Cached", func(t *testing.T) {
			t.Parallel()

			got := scenarioCacheSeq2Cached(slices.Clone(tc.values), tc.work, tc.passes)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Uncached", func(t *testing.T) {
			t.Parallel()

			got := scenarioCacheSeq2Uncached(slices.Clone(tc.values), tc.work, tc.passes)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
