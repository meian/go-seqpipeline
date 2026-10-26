package benchmarks

import (
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
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Fill", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeqResult = scenarioCacheSeqFill(values, 0)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Fill", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeqResult = scenarioCacheSeqFill(values, 32)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Fill", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeqResult = scenarioCacheSeqFill(values, 0)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Fill", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeqResult = scenarioCacheSeqFill(values, 32)
		}
	})
}

func TestScenarioCacheSeqFill(t *testing.T) {
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
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Warm", func(b *testing.B) {
		values := input(64)
		source := scenarioCacheSeqFill(values, 0)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqWarm(source)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Warm", func(b *testing.B) {
		values := input(64)
		source := scenarioCacheSeqFill(values, 32)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqWarm(source)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Warm", func(b *testing.B) {
		values := input(4096)
		source := scenarioCacheSeqFill(values, 0)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqWarm(source)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Warm", func(b *testing.B) {
		values := input(4096)
		source := scenarioCacheSeqFill(values, 32)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqWarm(source)
		}
	})
}

func TestScenarioCacheSeqWarm(t *testing.T) {
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
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 32, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqCached(values, 32, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=0/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=64/work=32/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 32, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=0/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq/N=4096/work=32/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeqUncached(values, 32, 16)
		}
	})
}

func TestScenarioCacheSeqReuse(t *testing.T) {
	t.Run("Cached", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		got := scenarioCacheSeqCached(values, 0, 2)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Uncached", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		got := scenarioCacheSeqUncached(values, 0, 2)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
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
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Fill", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeq2Result = scenarioCacheSeq2Fill(values, 0)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Fill", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeq2Result = scenarioCacheSeq2Fill(values, 32)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Fill", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeq2Result = scenarioCacheSeq2Fill(values, 0)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Fill", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			scenarioCacheSeq2Result = scenarioCacheSeq2Fill(values, 32)
		}
	})
}

func TestScenarioCacheSeq2Fill(t *testing.T) {
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
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Warm", func(b *testing.B) {
		values := input(64)
		source := scenarioCacheSeq2Fill(values, 0)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Warm(source)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Warm", func(b *testing.B) {
		values := input(64)
		source := scenarioCacheSeq2Fill(values, 32)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Warm(source)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Warm", func(b *testing.B) {
		values := input(4096)
		source := scenarioCacheSeq2Fill(values, 0)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Warm(source)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Warm", func(b *testing.B) {
		values := input(4096)
		source := scenarioCacheSeq2Fill(values, 32)
		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Warm(source)
		}
	})
}

func TestScenarioCacheSeq2Warm(t *testing.T) {
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
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 32, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=1/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=4/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=16/impl=Cached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Cached(values, 32, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=0/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=64/work=32/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 32, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 0, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 0, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=0/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 0, 16)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=1/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 32, 1)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=4/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 32, 4)
		}
	})
	b.Run("class=ScenarioCache/kind=Seq2/N=4096/work=32/phase=Reuse/reuse=16/impl=Uncached", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = scenarioCacheSeq2Uncached(values, 32, 16)
		}
	})
}

func TestScenarioCacheSeq2Reuse(t *testing.T) {
	t.Run("Cached", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		got := scenarioCacheSeq2Cached(values, 0, 2)
		const want int64 = 72
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Uncached", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		got := scenarioCacheSeq2Uncached(values, 0, 2)
		const want int64 = 72
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
