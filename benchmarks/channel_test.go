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
	data := &unitInput{
		values: values,
		source: source,
	}
	return []unitCase{
		{"FromChannel", "Evaluate", unitRun{data, runUnitFromChannelEvaluateLibrary}, total},
		{"ToChannel", "Evaluate", unitRun{data, runUnitToChannelEvaluateLibrary}, total},
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

func channelTransform(source seq.Seq[int]) seq.Seq[int] {
	return source.Pipe(seq.Map(mapValue)).Pipe(seq.Filter(keepValue))
}

func runScenarioChannelSeq(values []int, capacity int) []pair {
	var out []pair
	for v := range channelTransform(seq.FromSlice(values)) {
		out = append(out, pair{value: v})
	}
	return out
}

func runScenarioChannelToChannel(values []int, capacity int) []pair {
	var out []pair
	source := channelTransform(seq.FromSlice(values))
	for v := range seq.ToChannel(context.Background(), source) {
		out = append(out, pair{value: v})
	}
	return out
}

func runScenarioChannelDirect(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	for v := range ch {
		v = mapValue(v)
		if keepValue(v) {
			out = append(out, pair{value: v})
		}
	}
	<-done
	return out
}

func runScenarioChannelFromChannel(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	for v := range channelTransform(seq.FromChannel(context.Background(), ch)) {
		out = append(out, pair{value: v})
	}
	<-done
	return out
}

func runScenarioChannelRoundTrip(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	source := channelTransform(seq.FromChannel(context.Background(), ch))
	for v := range seq.ToChannel(context.Background(), source) {
		out = append(out, pair{value: v})
	}
	<-done
	return out
}

func testScenarioChannel(t *testing.T, run func([]int, int) []pair) {
	t.Helper()
	for _, n := range []int{0, 1, 64, 4096} {
		values := input(n)
		want := pipelineExpected(values, "Seq", 2, 0)
		for _, capacity := range []int{0, 64} {
			t.Run(fmt.Sprintf("N%d/B%d", n, capacity), func(t *testing.T) {
				if got := run(values, capacity); !slices.Equal(got, want) {
					t.Fatalf("got %v; want %v", got, want)
				}
			})
		}
	}
}

func channelUnitReferences(values []int) map[string]unitReference {
	data := &unitInput{
		values: values,
	}
	return map[string]unitReference{
		"FromChannel": {"ChannelReceive", unitRun{data, runUnitFromChannelEvaluateGo}},
		"ToChannel":   {"ChannelSend", unitRun{data, runUnitToChannelEvaluateGo}},
	}
}

func benchmarkScenarioChannelSeq(b *testing.B, values []int, capacity int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioChannelSeq(values, capacity)
	}
}

func TestScenarioChannelSeq(t *testing.T) { testScenarioChannel(t, runScenarioChannelSeq) }

func benchmarkScenarioChannelToChannel(b *testing.B, values []int, capacity int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioChannelToChannel(values, capacity)
	}
}

func TestScenarioChannelToChannel(t *testing.T) { testScenarioChannel(t, runScenarioChannelToChannel) }

func benchmarkScenarioChannelDirect(b *testing.B, values []int, capacity int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioChannelDirect(values, capacity)
	}
}

func TestScenarioChannelDirect(t *testing.T) { testScenarioChannel(t, runScenarioChannelDirect) }

func benchmarkScenarioChannelFromChannel(b *testing.B, values []int, capacity int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioChannelFromChannel(values, capacity)
	}
}

func TestScenarioChannelFromChannel(t *testing.T) {
	testScenarioChannel(t, runScenarioChannelFromChannel)
}

func benchmarkScenarioChannelRoundTrip(b *testing.B, values []int, capacity int) {
	b.ReportAllocs()
	for b.Loop() {
		output = runScenarioChannelRoundTrip(values, capacity)
	}
}

func TestScenarioChannelRoundTrip(t *testing.T) { testScenarioChannel(t, runScenarioChannelRoundTrip) }

func BenchmarkScenarioChannel(b *testing.B) {
	for _, n := range []int{64, 4096} {
		values := input(n)
		for _, implementation := range []struct {
			name  string
			entry func(*testing.B, []int, int)
		}{
			{"Seq", benchmarkScenarioChannelSeq},
			{"ToChannel", benchmarkScenarioChannelToChannel},
			{"Direct", benchmarkScenarioChannelDirect},
			{"FromChannel", benchmarkScenarioChannelFromChannel},
			{"RoundTrip", benchmarkScenarioChannelRoundTrip},
		} {
			capacities := []int{0}
			if implementation.name != "Seq" && implementation.name != "ToChannel" {
				capacities = []int{0, 64}
			}
			for _, capacity := range capacities {
				b.Run(fmt.Sprintf("N=%d/buffer=%d/impl=%s", n, capacity, implementation.name), func(b *testing.B) { implementation.entry(b, values, capacity) })
			}
		}
	}
}

func runUnitFromChannelEvaluateLibrary(data *unitInput) int64 {
	ch, done := producer(data.values, 0)
	sum := sumSeq(seq.FromChannel(context.Background(), ch))
	<-done
	return sum
}

func benchmarkUnitFromChannelEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromChannelEvaluateLibrary)
}

func TestUnitFromChannelEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "FromChannel", "Evaluate", "Library", runUnitFromChannelEvaluateLibrary)
}

func runUnitFromChannelEvaluateGo(data *unitInput) int64 {
	ch, done := producer(data.values, 0)
	var sum int64
	for v := range ch {
		sum += int64(v)
	}
	<-done
	return sum
}

func benchmarkUnitFromChannelEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromChannelEvaluateGo)
}

func TestUnitFromChannelEvaluateGo(t *testing.T) {
	testUnitOperation(t, "FromChannel", "Evaluate", "Go", runUnitFromChannelEvaluateGo)
}

func runUnitToChannelEvaluateLibrary(data *unitInput) int64 {
	var sum int64
	for v := range seq.ToChannel(context.Background(), data.source) {
		sum += int64(v)
	}
	return sum
}

func benchmarkUnitToChannelEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToChannelEvaluateLibrary)
}

func TestUnitToChannelEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ToChannel", "Evaluate", "Library", runUnitToChannelEvaluateLibrary)
}

func runUnitToChannelEvaluateGo(data *unitInput) int64 {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range data.values {
			ch <- v
		}
	}()
	var sum int64
	for v := range ch {
		sum += int64(v)
	}
	return sum
}

func benchmarkUnitToChannelEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToChannelEvaluateGo)
}

func TestUnitToChannelEvaluateGo(t *testing.T) {
	testUnitOperation(t, "ToChannel", "Evaluate", "Go", runUnitToChannelEvaluateGo)
}
