package benchmarks

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func channelUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	var total int64
	for _, v := range values {
		total += int64(v)
	}
	return []unitCase{
		{"FromChannel", "Evaluate", func() int64 {
			ch, done := producer(values, 0)
			sum := sumSeq(seq.FromChannel(context.Background(), ch))
			<-done
			return sum
		}, total},
		{"ToChannel", "Evaluate", func() int64 {
			var sum int64
			for v := range seq.ToChannel(context.Background(), source) {
				sum += int64(v)
			}
			return sum
		}, total},
	}
}

func TestChannelUnitResults(t *testing.T) { checkUnitCases(t, channelUnitCases) }

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
		for v := range seq.ToChannel(ctx, s) {
			collect(v)
		}
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
