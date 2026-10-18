package benchmarks

import (
	"context"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

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

func scenarioChannelSeq(values []int, capacity int) []pair {
	var out []pair
	for v := range channelTransform(seq.FromSlice(values)) {
		out = append(out, pair{value: v})
	}
	return out
}

func BenchmarkScenarioChannelSeq(b *testing.B) {
	b.Run("class=ScenarioChannel/N=64/buffer=0/impl=Seq", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelSeq(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=0/impl=Seq", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelSeq(values, 0)
		}
	})
}

func TestScenarioChannelSeq(t *testing.T) {
	t.Run("N64Buffer0", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelSeq(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer0", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelSeq(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelSeq(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelSeq(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelSeq(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func scenarioChannelToChannel(values []int, capacity int) []pair {
	var out []pair
	source := channelTransform(seq.FromSlice(values))
	for v := range seq.ToChannel(context.Background(), source) {
		out = append(out, pair{value: v})
	}
	return out
}

func BenchmarkScenarioChannelToChannel(b *testing.B) {
	b.Run("class=ScenarioChannel/N=64/buffer=0/impl=ToChannel", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelToChannel(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=0/impl=ToChannel", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelToChannel(values, 0)
		}
	})
}

func TestScenarioChannelToChannel(t *testing.T) {
	t.Run("N64Buffer0", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelToChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer0", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelToChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelToChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelToChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelToChannel(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func scenarioChannelDirect(values []int, capacity int) []pair {
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

func BenchmarkScenarioChannelDirect(b *testing.B) {
	b.Run("class=ScenarioChannel/N=64/buffer=0/impl=Direct", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelDirect(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=64/buffer=64/impl=Direct", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelDirect(values, 64)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=0/impl=Direct", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelDirect(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=64/impl=Direct", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelDirect(values, 64)
		}
	})
}

func TestScenarioChannelDirect(t *testing.T) {
	t.Run("N64Buffer0", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Buffer64", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer0", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer64", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelDirect(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func scenarioChannelFromChannel(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	for v := range channelTransform(seq.FromChannel(context.Background(), ch)) {
		out = append(out, pair{value: v})
	}
	<-done
	return out
}

func BenchmarkScenarioChannelFromChannel(b *testing.B) {
	b.Run("class=ScenarioChannel/N=64/buffer=0/impl=FromChannel", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelFromChannel(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=64/buffer=64/impl=FromChannel", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelFromChannel(values, 64)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=0/impl=FromChannel", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelFromChannel(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=64/impl=FromChannel", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelFromChannel(values, 64)
		}
	})
}

func TestScenarioChannelFromChannel(t *testing.T) {
	t.Run("N64Buffer0", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Buffer64", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer0", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer64", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelFromChannel(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func scenarioChannelRoundTrip(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	source := channelTransform(seq.FromChannel(context.Background(), ch))
	for v := range seq.ToChannel(context.Background(), source) {
		out = append(out, pair{value: v})
	}
	<-done
	return out
}

func BenchmarkScenarioChannelRoundTrip(b *testing.B) {
	b.Run("class=ScenarioChannel/N=64/buffer=0/impl=RoundTrip", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelRoundTrip(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=64/buffer=64/impl=RoundTrip", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelRoundTrip(values, 64)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=0/impl=RoundTrip", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelRoundTrip(values, 0)
		}
	})
	b.Run("class=ScenarioChannel/N=4096/buffer=64/impl=RoundTrip", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			output = scenarioChannelRoundTrip(values, 64)
		}
	})
}

func TestScenarioChannelRoundTrip(t *testing.T) {
	t.Run("N64Buffer0", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N64Buffer64", func(t *testing.T) {
		values := input(64)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer0", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("N4096Buffer64", func(t *testing.T) {
		values := input(4096)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		var want []pair
		for _, v := range values {
			value := (v + 1) & 65535
			if value%3 != 0 {
				want = append(want, pair{value: value})
			}
		}
		got := scenarioChannelRoundTrip(values, 64)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}
