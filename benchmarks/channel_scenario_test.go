package benchmarks

import (
	"context"
	"fmt"
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

func scenarioChannelToChannel(values []int, capacity int) []pair {
	var out []pair
	source := channelTransform(seq.FromSlice(values))
	channel := seq.ToChannel(context.Background(), source)
	for v := range channel {
		out = append(out, pair{value: v})
	}
	return out
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

func scenarioChannelSeq(values []int, capacity int) []pair {
	var out []pair
	source := channelTransform(seq.FromSlice(values))
	for v := range source {
		out = append(out, pair{value: v})
	}
	return out
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

func BenchmarkScenarioChannel(b *testing.B) {
	sizes := []int{64, 4096}
	cases := []struct {
		n        int
		capacity int
	}{
		{64, 0},
		{64, 64},
		{4096, 0},
		{4096, 64},
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=ScenarioChannel/N=%d/buffer=0/impl=Seq", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioChannelSeq(values, 0)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=ScenarioChannel/N=%d/buffer=0/impl=ToChannel", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioChannelToChannel(values, 0)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioChannel/N=%d/buffer=%d/impl=Direct", tc.n, tc.capacity)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioChannelDirect(values, tc.capacity)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioChannel/N=%d/buffer=%d/impl=FromChannel", tc.n, tc.capacity)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioChannelFromChannel(values, tc.capacity)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=ScenarioChannel/N=%d/buffer=%d/impl=RoundTrip", tc.n, tc.capacity)
		b.Run(name, func(b *testing.B) {
			values := input(tc.n)
			b.ReportAllocs()
			for b.Loop() {
				output = scenarioChannelRoundTrip(values, tc.capacity)
			}
		})
	}
}

func TestScenarioChannel(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values   []int
		capacity int
		want     []pair
	}{
		{
			values:   []int{1, 2, 3, 4, 5, 6},
			capacity: 64,
			want:     []pair{{value: 2}, {value: 4}, {value: 5}, {value: 7}},
		},
	}
	for _, tc := range cases {
		t.Run("Seq", func(t *testing.T) {
			t.Parallel()

			got := scenarioChannelSeq(slices.Clone(tc.values), 0)
			if !slices.Equal(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("ToChannel", func(t *testing.T) {
			t.Parallel()

			got := scenarioChannelToChannel(slices.Clone(tc.values), 0)
			if !slices.Equal(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Direct", func(t *testing.T) {
			t.Parallel()

			got := scenarioChannelDirect(slices.Clone(tc.values), tc.capacity)
			if !slices.Equal(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("FromChannel", func(t *testing.T) {
			t.Parallel()

			got := scenarioChannelFromChannel(slices.Clone(tc.values), tc.capacity)
			if !slices.Equal(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("RoundTrip", func(t *testing.T) {
			t.Parallel()

			got := scenarioChannelRoundTrip(slices.Clone(tc.values), tc.capacity)
			if !slices.Equal(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
