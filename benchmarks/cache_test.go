package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var cacheFillLibraryResult seq.Seq[int]

func cacheFillLibrary(values []int) seq.Seq[int] {
	return seq.Cache(seq.FromSlice(values))
}

func BenchmarkCacheFillLibrary(b *testing.B) {
	b.Run("class=Unit/api=Cache/phase=Fill/ref=SlicesCollect/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			cacheFillLibraryResult = cacheFillLibrary(values)
		}
	})
	b.Run("class=Unit/api=Cache/phase=Fill/ref=SlicesCollect/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			cacheFillLibraryResult = cacheFillLibrary(values)
		}
	})
}

func TestCacheFillLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cacheFillLibrary(values)
	got := slices.Collect(seq.ToSeq(cached))
	if !slices.Equal(got, values) {
		t.Fatalf("got %v; want %v", got, values)
	}
}

var cacheFillGoResult []int

func cacheFillGo(values []int) []int {
	return slices.Collect(slices.Values(values))
}

func BenchmarkCacheFillGo(b *testing.B) {
	b.Run("class=Unit/api=Cache/phase=Fill/ref=SlicesCollect/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			cacheFillGoResult = cacheFillGo(values)
		}
	})
	b.Run("class=Unit/api=Cache/phase=Fill/ref=SlicesCollect/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			cacheFillGoResult = cacheFillGo(values)
		}
	})
}

func TestCacheFillGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cacheFillGo(values)
	got := cached
	if !slices.Equal(got, values) {
		t.Fatalf("got %v; want %v", got, values)
	}
}

func cacheWarmLibrary(cached seq.Seq[int]) int64 {
	var sum int64
	for v := range cached {
		sum += int64(v)
	}
	return sum
}

func BenchmarkCacheWarmLibrary(b *testing.B) {
	b.Run("class=Unit/api=Cache/phase=Warm/ref=SliceRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		cached := cacheFillLibrary(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cacheWarmLibrary(cached)
		}
	})
	b.Run("class=Unit/api=Cache/phase=Warm/ref=SliceRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		cached := cacheFillLibrary(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cacheWarmLibrary(cached)
		}
	})
}

func TestCacheWarmLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cacheFillLibrary(values)
	got := cacheWarmLibrary(cached)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func cacheWarmGo(cached []int) int64 {
	var sum int64
	for _, v := range cached {
		sum += int64(v)
	}
	return sum
}

func BenchmarkCacheWarmGo(b *testing.B) {
	b.Run("class=Unit/api=Cache/phase=Warm/ref=SliceRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		cached := cacheFillGo(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cacheWarmGo(cached)
		}
	})
	b.Run("class=Unit/api=Cache/phase=Warm/ref=SliceRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		cached := cacheFillGo(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cacheWarmGo(cached)
		}
	})
}

func TestCacheWarmGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cacheFillGo(values)
	got := cacheWarmGo(cached)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

var cache2FillLibraryResult seq.Seq2[int, int]

func cache2FillLibrary(values []int) seq.Seq2[int, int] {
	return seq.Cache2(seq.FromSeq2(slices.All(values)))
}

func BenchmarkCache2FillLibrary(b *testing.B) {
	b.Run("class=Unit/api=Cache2/phase=Fill/ref=PairAppend/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			cache2FillLibraryResult = cache2FillLibrary(values)
		}
	})
	b.Run("class=Unit/api=Cache2/phase=Fill/ref=PairAppend/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			cache2FillLibraryResult = cache2FillLibrary(values)
		}
	})
}

func TestCache2FillLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cache2FillLibrary(values)
	var got []pair
	for k, v := range cached {
		got = append(got, pair{k, v})
	}
	want := []pair{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

var cache2FillGoResult []pair

func cache2FillGo(values []int) []pair {
	var cached []pair
	for k, v := range slices.All(values) {
		cached = append(cached, pair{k, v})
	}
	return cached
}

func BenchmarkCache2FillGo(b *testing.B) {
	b.Run("class=Unit/api=Cache2/phase=Fill/ref=PairAppend/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			cache2FillGoResult = cache2FillGo(values)
		}
	})
	b.Run("class=Unit/api=Cache2/phase=Fill/ref=PairAppend/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			cache2FillGoResult = cache2FillGo(values)
		}
	})
}

func TestCache2FillGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cache2FillGo(values)
	got := cached
	want := []pair{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func cache2WarmLibrary(cached seq.Seq2[int, int]) int64 {
	var sum int64
	for k, v := range cached {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkCache2WarmLibrary(b *testing.B) {
	b.Run("class=Unit/api=Cache2/phase=Warm/ref=PairRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		cached := cache2FillLibrary(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cache2WarmLibrary(cached)
		}
	})
	b.Run("class=Unit/api=Cache2/phase=Warm/ref=PairRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		cached := cache2FillLibrary(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cache2WarmLibrary(cached)
		}
	})
}

func TestCache2WarmLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cache2FillLibrary(values)
	got := cache2WarmLibrary(cached)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func cache2WarmGo(cached []pair) int64 {
	var sum int64
	for _, p := range cached {
		sum += int64(p.key) + int64(p.value)
	}
	return sum
}

func BenchmarkCache2WarmGo(b *testing.B) {
	b.Run("class=Unit/api=Cache2/phase=Warm/ref=PairRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		cached := cache2FillGo(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cache2WarmGo(cached)
		}
	})
	b.Run("class=Unit/api=Cache2/phase=Warm/ref=PairRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		cached := cache2FillGo(values)
		b.ReportAllocs()
		for b.Loop() {
			checksum = cache2WarmGo(cached)
		}
	})
}

func TestCache2WarmGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	cached := cache2FillGo(values)
	got := cache2WarmGo(cached)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func cacheValue(v, work int) int {
	for j := 0; j < work; j++ {
		v = (v*1664525 + 1013904223) & 65535
	}
	return v
}

func cacheSeq(values []int, work int) seq.Seq[int] {
	return seq.FromSlice(values).Pipe(seq.Map(func(v int) int { return cacheValue(v, work) }))
}

func cacheSeq2(values []int, work int) seq.Seq2[int, int] {
	return seq.FromSeq2(slices.All(values)).Pipe2(seq.Map2(func(k, v int) (int, int) { return k, cacheValue(v, work) }))
}

var scenarioCacheSeqResult seq.Seq[int]

func scenarioCacheSeqFill(values []int, work int) seq.Seq[int] {
	return seq.Cache(cacheSeq(values, work))
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
	got := seq.ToSlice(cached)
	want := values
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
	source := scenarioCacheSeqFill(values, work)
	var sum int64
	for r := 0; r < reuse; r++ {
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkScenarioCacheSeqCached(b *testing.B) {
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
}

func TestScenarioCacheSeqCached(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioCacheSeqCached(values, 0, 2)
	const want int64 = 42
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
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

func BenchmarkScenarioCacheSeqUncached(b *testing.B) {
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

func TestScenarioCacheSeqUncached(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioCacheSeqUncached(values, 0, 2)
	const want int64 = 42
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

var scenarioCacheSeq2Result seq.Seq2[int, int]

func scenarioCacheSeq2Fill(values []int, work int) seq.Seq2[int, int] {
	return seq.Cache2(cacheSeq2(values, work))
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
	source := scenarioCacheSeq2Fill(values, work)
	var sum int64
	for r := 0; r < reuse; r++ {
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func BenchmarkScenarioCacheSeq2Cached(b *testing.B) {
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
}

func TestScenarioCacheSeq2Cached(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioCacheSeq2Cached(values, 0, 2)
	const want int64 = 72
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
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

func BenchmarkScenarioCacheSeq2Uncached(b *testing.B) {
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

func TestScenarioCacheSeq2Uncached(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioCacheSeq2Uncached(values, 0, 2)
	const want int64 = 72
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
