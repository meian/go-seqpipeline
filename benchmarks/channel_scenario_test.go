package benchmarks

import (
	"context"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func channelTransform(source seq.Seq[int]) seq.Seq[int] {
	return source.
		Pipe(op.Map(mapValue)).
		Pipe(op.Filter(keepValue))
}

func scenarioChannelSeq(values []int, capacity int) []pair {
	var out []pair
	source := channelTransform(seq.FromSlice(values))
	for v := range source {
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
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioChannelSeq(values, 0)
	want := []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func scenarioChannelToChannel(values []int, capacity int) []pair {
	var out []pair
	source := channelTransform(seq.FromSlice(values))
	channel := seq.ToChannel(context.Background(), source)
	for v := range channel {
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
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioChannelToChannel(values, 0)
	want := []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
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
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioChannelDirect(values, 64)
	want := []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func scenarioChannelFromChannel(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	source := channelTransform(seq.FromChannel(context.Background(), ch))
	for v := range source {
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
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioChannelFromChannel(values, 64)
	want := []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func scenarioChannelRoundTrip(values []int, capacity int) []pair {
	ch, done := producer(values, capacity)
	var out []pair
	source := channelTransform(seq.FromChannel(context.Background(), ch))
	channel := seq.ToChannel(context.Background(), source)
	for v := range channel {
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
	values := []int{1, 2, 3, 4, 5, 6}
	got := scenarioChannelRoundTrip(values, 64)
	want := []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
