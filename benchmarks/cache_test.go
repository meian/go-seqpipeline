package benchmarks

import (
	"fmt"
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
	return []unitCase{
		{"Cache", "Fill", func() int64 { s := seq.Cache(source); cachedConsumer = func() int64 { return sumSeq(s) }; return 0 }, 0},
		{"Cache", "Warm", func() int64 { return sumSeq(warm) }, total},
		{"Cache2", "Fill", func() int64 { s := seq.Cache2(pairs); cachedConsumer = func() int64 { return sumSeq2(s) }; return 0 }, 0},
		{"Cache2", "Warm", func() int64 { return sumSeq2(warm2) }, total + keys},
	}
}

func TestCacheUnitResults(t *testing.T) {
	for _, n := range []int{0, 1, 64, 4096} {
		for _, tc := range cacheUnitCases(input(n)) {
			t.Run(fmt.Sprintf("%s/N=%d/phase=%s", tc.api, n, tc.phase), func(t *testing.T) {
				if got := tc.run(); got != tc.want {
					t.Fatalf("got %d; want %d", got, tc.want)
				}
				if tc.phase == "Fill" {
					want := int64(n * (n - 1) / 2)
					if tc.api == "Cache2" {
						want *= 2
					}
					if got := cachedConsumer(); got != want {
						t.Fatalf("filled cache: got %d; want %d", got, want)
					}
				}
			})
		}
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

// cached=true の呼び出しが Cache/Cache2 の作成・充填に相当する。
func cacheConsumer(values []int, kind string, work int, cached bool) func() int64 {
	if kind == "Seq" {
		s := cacheSeq(values, work)
		if cached {
			s = seq.Cache(s)
		}
		return func() int64 { return sumSeq(s) }
	}
	s := cacheSeq2(values, work)
	if cached {
		s = seq.Cache2(s)
	}
	return func() int64 { return sumSeq2(s) }
}

func TestCacheEquivalent(t *testing.T) {
	for _, kind := range []string{"Seq", "Seq2"} {
		for _, n := range []int{0, 1, 64, 4096} {
			for _, work := range []int{0, 32} {
				t.Run(fmt.Sprintf("%s/N%d/W%d", kind, n, work), func(t *testing.T) {
					values := input(n)
					var want int64
					for k, v := range values {
						want += int64(cacheValue(v, work))
						if kind == "Seq2" {
							want += int64(k)
						}
					}
					for _, cached := range []bool{false, true} {
						consume := cacheConsumer(values, kind, work, cached)
						for r := 0; r < 16; r++ {
							if got := consume(); got != want {
								t.Fatalf("cached=%v got %d; want %d", cached, got, want)
							}
						}
					}
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

func BenchmarkScenarioCache(b *testing.B) {
	for _, kind := range []string{"Seq", "Seq2"} {
		for _, n := range []int{64, 4096} {
			values := input(n)
			for _, work := range []int{0, 32} {
				prefix := fmt.Sprintf("kind=%s/N=%d/work=%d", kind, n, work)
				b.Run(prefix+"/phase=Warm", func(b *testing.B) {
					consume := cacheConsumer(values, kind, work, true)
					b.ReportAllocs()
					for b.Loop() {
						checksum = consume()
					}
				})
				b.Run(prefix+"/phase=Fill", func(b *testing.B) {
					b.ReportAllocs()
					for b.Loop() {
						// 保存のみ。下流による反復は含めず、結果をグローバルへ保持する。
						cachedConsumer = cacheConsumer(values, kind, work, true)
					}
				})
				for _, reuse := range []int{1, 4, 16} {
					for _, cached := range []bool{false, true} {
						impl := "Uncached"
						if cached {
							impl = "Cached"
						}
						b.Run(fmt.Sprintf("%s/phase=Reuse/reuse=%d/impl=%s", prefix, reuse, impl), func(b *testing.B) {
							b.ReportAllocs()
							for b.Loop() {
								consume := cacheConsumer(values, kind, work, cached)
								var sum int64
								for r := 0; r < reuse; r++ {
									sum += consume()
								}
								checksum = sum
							}
						})
					}
				}
			}
		}
	}
}

var cachedConsumer func() int64
