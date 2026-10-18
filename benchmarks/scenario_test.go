package benchmarks

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

type pair struct{ key, value int }

var output []pair
var checksum int64

func input(n int) []int {
	values := make([]int, n)
	for i := range values {
		values[i] = i
	}
	return values
}
func mapValue(v int) int   { return (v + 1) & 65535 }
func keepValue(v int) bool { return v%3 != 0 }

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

// producer は送信専用の 1 goroutine。close 後、呼び出し元が終了を join する。
func producer(values []int, capacity int) (<-chan int, <-chan struct{}) {
	ch := make(chan int, capacity)
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch, done
}

// 全ケースで同じ Map/Filter と append による出力収集を実行する。
func channelRun(values []int, capacity int, impl string) []pair {
	ctx := context.Background()
	var out []pair
	collect := func(v int) { out = append(out, pair{value: v}) }
	transform := func(s seq.Seq[int]) seq.Seq[int] { return s.Pipe(seq.Map(mapValue)).Pipe(seq.Filter(keepValue)) }
	switch impl {
	case "Seq":
		for v := range transform(seq.FromSlice(values)) {
			collect(v)
		}
	case "ToChannel":
		s := transform(seq.FromSlice(values))
		// 入力を最後まで評価した goroutine の終了も待つ。
		done := make(chan struct{})
		joined := seq.Seq[int](func(yield func(int) bool) {
			defer close(done)
			for v := range s {
				if !yield(v) {
					return
				}
			}
		})
		for v := range seq.ToChannel(ctx, joined) {
			collect(v)
		}
		<-done
	default:
		ch, done := producer(values, capacity)
		switch impl {
		case "Direct":
			for v := range ch {
				v = mapValue(v)
				if keepValue(v) {
					collect(v)
				}
			}
		case "FromChannel":
			for v := range transform(seq.FromChannel(ctx, ch)) {
				collect(v)
			}
		case "RoundTrip":
			for v := range seq.ToChannel(ctx, transform(seq.FromChannel(ctx, ch))) {
				collect(v)
			}
		}
		<-done
	}
	return out
}

func TestChannelEquivalent(t *testing.T) {
	for _, n := range []int{0, 1, 64, 4096} {
		values := input(n)
		want := baseline(values, "Seq", 2, 0)
		for _, capacity := range []int{0, 64} {
			for _, impl := range []string{"Seq", "ToChannel", "Direct", "FromChannel", "RoundTrip"} {
				t.Run(fmt.Sprintf("N%d/B%d/%s", n, capacity, impl), func(t *testing.T) {
					if got := channelRun(values, capacity, impl); !slices.Equal(got, want) {
						t.Fatalf("got %v; want %v", got, want)
					}
				})
			}
		}
	}
}

func BenchmarkScenarioChannel(b *testing.B) {
	for _, n := range []int{64, 4096} {
		values := input(n)
		for _, impl := range []string{"Seq", "ToChannel", "Direct", "FromChannel", "RoundTrip"} {
			capacities := []int{0}
			if impl != "Seq" && impl != "ToChannel" {
				capacities = []int{0, 64}
			}
			for _, capacity := range capacities {
				b.Run(fmt.Sprintf("N=%d/buffer=%d/impl=%s", n, capacity, impl), func(b *testing.B) {
					b.ReportAllocs()
					for b.Loop() {
						output = channelRun(values, capacity, impl)
					}
				})
			}
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
func sumSeq(s seq.Seq[int]) int64 {
	var sum int64
	for v := range s {
		sum += int64(v)
	}
	return sum
}
func sumSeq2(s seq.Seq2[int, int]) int64 {
	var sum int64
	for k, v := range s {
		sum += int64(k) + int64(v)
	}
	return sum
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
