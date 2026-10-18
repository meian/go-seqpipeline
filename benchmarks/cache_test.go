package benchmarks

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func cacheUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var total, keys int64
	for k, v := range values {
		total += int64(v)
		keys += int64(k)
	}
	warm := seq.Cache(source)
	warm2 := seq.Cache2(pairs)
	data := &unitInput{
		values: values,
		source: source,
		pairs:  pairs,
		warm:   warm,
		warm2:  warm2,
	}
	return []unitCase{
		{"Cache", "Fill", unitRun{data, runUnitCacheFillLibrary}, 0},
		{"Cache", "Warm", unitRun{data, runUnitCacheWarmLibrary}, total},
		{"Cache2", "Fill", unitRun{data, runUnitCache2FillLibrary}, 0},
		{"Cache2", "Warm", unitRun{data, runUnitCache2WarmLibrary}, total + keys},
	}
}

// 値依存の算術処理を繰り返し、上流の負荷を変える。保持値は常に int。
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

func TestCacheEquivalent(t *testing.T) {
	for _, kind := range []string{"Seq", "Seq2"} {
		for _, n := range []int{0, 1, 64, 4096} {
			for _, work := range []int{0, 32} {
				t.Run(fmt.Sprintf("%s/N%d/W%d", kind, n, work), func(t *testing.T) {
					values := input(n)
					// チェックサムに加え、保存した値・キーの順序も比較する。
					if got := seq.ToSlice(seq.Cache(cacheSeq(values, work))); !slices.Equal(got, seq.ToSlice(cacheSeq(values, work))) {
						t.Fatal("Seq cache differs")
					}
					var got, wantPairs []pair
					for k, v := range seq.Cache2(cacheSeq2(values, work)) {
						got = append(got, pair{k, v})
					}
					for k, v := range cacheSeq2(values, work) {
						wantPairs = append(wantPairs, pair{k, v})
					}
					if !slices.Equal(got, wantPairs) {
						t.Fatal("Seq2 cache differs")
					}
				})
			}
		}
	}
}

type cachedScenario struct {
	single seq.Seq[int]
	pairs  seq.Seq2[int, int]
}

var scenarioCachedResult cachedScenario

func runScenarioCacheFill(values []int, kind string, work int) cachedScenario {
	if kind == "Seq" {
		return cachedScenario{single: seq.Cache(cacheSeq(values, work))}
	}
	return cachedScenario{pairs: seq.Cache2(cacheSeq2(values, work))}
}

func runScenarioCacheWarm(cached cachedScenario) int64 {
	var sum int64
	if cached.single != nil {
		for v := range cached.single {
			sum += int64(v)
		}
	} else {
		for k, v := range cached.pairs {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func runScenarioCacheCached(values []int, kind string, work, reuse int) int64 {
	cached := runScenarioCacheFill(values, kind, work)
	var sum int64
	for r := 0; r < reuse; r++ {
		sum += runScenarioCacheWarm(cached)
	}
	return sum
}

func runScenarioCacheUncached(values []int, kind string, work, reuse int) int64 {
	var sum int64
	if kind == "Seq" {
		source := cacheSeq(values, work)
		for r := 0; r < reuse; r++ {
			sum += sumSeq(source)
		}
	} else {
		source := cacheSeq2(values, work)
		for r := 0; r < reuse; r++ {
			sum += sumSeq2(source)
		}
	}
	return sum
}

func benchmarkScenarioCacheFill(b *testing.B, values []int, kind string, work int) {
	b.ReportAllocs()
	for b.Loop() {
		scenarioCachedResult = runScenarioCacheFill(values, kind, work)
	}
}

func benchmarkScenarioCacheWarm(b *testing.B, values []int, kind string, work int) {
	cached := runScenarioCacheFill(values, kind, work)
	b.ReportAllocs()
	for b.Loop() {
		checksum = runScenarioCacheWarm(cached)
	}
}

func benchmarkScenarioCacheCached(b *testing.B, values []int, kind string, work, reuse int) {
	b.ReportAllocs()
	for b.Loop() {
		checksum = runScenarioCacheCached(values, kind, work, reuse)
	}
}

func benchmarkScenarioCacheUncached(b *testing.B, values []int, kind string, work, reuse int) {
	b.ReportAllocs()
	for b.Loop() {
		checksum = runScenarioCacheUncached(values, kind, work, reuse)
	}
}

func testScenarioCache(t *testing.T, run func([]int, string, int, int) int64) {
	t.Helper()
	for _, kind := range []string{"Seq", "Seq2"} {
		for _, n := range []int{0, 1, 64, 4096} {
			values := input(n)
			for _, work := range []int{0, 32} {
				var want int64
				for k, v := range values {
					want += int64(cacheValue(v, work))
					if kind == "Seq2" {
						want += int64(k)
					}
				}
				for _, reuse := range []int{1, 4, 16} {
					t.Run(fmt.Sprintf("%s/N%d/W%d/R%d", kind, n, work, reuse), func(t *testing.T) {
						if got := run(values, kind, work, reuse); got != want*int64(reuse) {
							t.Fatalf("got %d; want %d", got, want*int64(reuse))
						}
					})
				}
			}
		}
	}
}

func TestScenarioCacheFill(t *testing.T) {
	testScenarioCache(t, func(values []int, kind string, work, reuse int) int64 {
		cached := runScenarioCacheFill(values, kind, work)
		var sum int64
		for r := 0; r < reuse; r++ {
			sum += runScenarioCacheWarm(cached)
		}
		return sum
	})
}

func TestScenarioCacheWarm(t *testing.T) {
	testScenarioCache(t, func(values []int, kind string, work, reuse int) int64 {
		cached := runScenarioCacheFill(values, kind, work)
		var sum int64
		for r := 0; r < reuse; r++ {
			sum += runScenarioCacheWarm(cached)
		}
		return sum
	})
}

func TestScenarioCacheCached(t *testing.T) { testScenarioCache(t, runScenarioCacheCached) }

func TestScenarioCacheUncached(t *testing.T) { testScenarioCache(t, runScenarioCacheUncached) }

func BenchmarkScenarioCache(b *testing.B) {
	for _, kind := range []string{"Seq", "Seq2"} {
		for _, n := range []int{64, 4096} {
			values := input(n)
			for _, work := range []int{0, 32} {
				prefix := fmt.Sprintf("kind=%s/N=%d/work=%d", kind, n, work)
				b.Run(prefix+"/phase=Warm", func(b *testing.B) { benchmarkScenarioCacheWarm(b, values, kind, work) })
				b.Run(prefix+"/phase=Fill", func(b *testing.B) { benchmarkScenarioCacheFill(b, values, kind, work) })
				for _, reuse := range []int{1, 4, 16} {
					for _, implementation := range []struct {
						name  string
						entry func(*testing.B, []int, string, int, int)
					}{
						{"Uncached", benchmarkScenarioCacheUncached}, {"Cached", benchmarkScenarioCacheCached},
					} {
						b.Run(fmt.Sprintf("%s/phase=Reuse/reuse=%d/impl=%s", prefix, reuse, implementation.name), func(b *testing.B) { implementation.entry(b, values, kind, work, reuse) })
					}
				}
			}
		}
	}
}

var cachedConsumer func() int64

func cacheUnitReferences(values []int) map[string]unitReference {
	standard := slices.Values(values)
	standardPairs := slices.All(values)
	warm := slices.Collect(standard)
	warm2 := collectUnitPairs(standardPairs)
	data := &unitInput{
		values:        values,
		standard:      standard,
		standardPairs: standardPairs,
		warmValues:    warm,
		warmPairs:     warm2,
	}
	return map[string]unitReference{
		"Cache/Fill":  {"SlicesCollect", unitRun{data, runUnitCacheFillGo}},
		"Cache/Warm":  {"SliceRange", unitRun{data, runUnitCacheWarmGo}},
		"Cache2/Fill": {"PairAppend", unitRun{data, runUnitCache2FillGo}},
		"Cache2/Warm": {"PairRange", unitRun{data, runUnitCache2WarmGo}},
	}
}

func collectUnitPairs(source iter.Seq2[int, int]) []pair {
	var cached []pair
	for k, v := range source {
		cached = append(cached, pair{k, v})
	}
	return cached
}

func sumUnitPairs(cached []pair) int64 {
	var sum int64
	for _, p := range cached {
		sum += int64(p.key) + int64(p.value)
	}
	return sum
}

func runUnitCacheFillLibrary(data *unitInput) int64 {
	s := seq.Cache(data.source)
	cachedConsumer = func() int64 { return sumSeq(s) }
	return 0
}

func benchmarkUnitCacheFillLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCacheFillLibrary)
}

func TestUnitCacheFillLibrary(t *testing.T) {
	testUnitOperation(t, "Cache", "Fill", "Library", runUnitCacheFillLibrary)
}

func runUnitCacheFillGo(data *unitInput) int64 {
	cached := slices.Collect(data.standard)
	cachedConsumer = func() int64 { return rangeValues(cached) }
	return 0
}

func benchmarkUnitCacheFillGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCacheFillGo)
}

func TestUnitCacheFillGo(t *testing.T) {
	testUnitOperation(t, "Cache", "Fill", "Go", runUnitCacheFillGo)
}

func runUnitCacheWarmLibrary(data *unitInput) int64 { return sumSeq(data.warm) }

func benchmarkUnitCacheWarmLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCacheWarmLibrary)
}

func TestUnitCacheWarmLibrary(t *testing.T) {
	testUnitOperation(t, "Cache", "Warm", "Library", runUnitCacheWarmLibrary)
}

func runUnitCacheWarmGo(data *unitInput) int64 { return rangeValues(data.warmValues) }

func benchmarkUnitCacheWarmGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCacheWarmGo)
}

func TestUnitCacheWarmGo(t *testing.T) {
	testUnitOperation(t, "Cache", "Warm", "Go", runUnitCacheWarmGo)
}

func runUnitCache2FillLibrary(data *unitInput) int64 {
	s := seq.Cache2(data.pairs)
	cachedConsumer = func() int64 { return sumSeq2(s) }
	return 0
}

func benchmarkUnitCache2FillLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCache2FillLibrary)
}

func TestUnitCache2FillLibrary(t *testing.T) {
	testUnitOperation(t, "Cache2", "Fill", "Library", runUnitCache2FillLibrary)
}

func runUnitCache2FillGo(data *unitInput) int64 {
	cached := collectUnitPairs(data.standardPairs)
	cachedConsumer = func() int64 { return sumUnitPairs(cached) }
	return 0
}

func benchmarkUnitCache2FillGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCache2FillGo)
}

func TestUnitCache2FillGo(t *testing.T) {
	testUnitOperation(t, "Cache2", "Fill", "Go", runUnitCache2FillGo)
}

func runUnitCache2WarmLibrary(data *unitInput) int64 { return sumSeq2(data.warm2) }

func benchmarkUnitCache2WarmLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCache2WarmLibrary)
}

func TestUnitCache2WarmLibrary(t *testing.T) {
	testUnitOperation(t, "Cache2", "Warm", "Library", runUnitCache2WarmLibrary)
}

func runUnitCache2WarmGo(data *unitInput) int64 { return sumUnitPairs(data.warmPairs) }

func benchmarkUnitCache2WarmGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitCache2WarmGo)
}

func TestUnitCache2WarmGo(t *testing.T) {
	testUnitOperation(t, "Cache2", "Warm", "Go", runUnitCache2WarmGo)
}
