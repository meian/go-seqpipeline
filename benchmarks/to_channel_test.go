package benchmarks

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func toChannelLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	source := seq.ToChannel(context.Background(), inputSeq)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func toChannelGo(values []int) int64 {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range slices.Values(values) {
			ch <- v
		}
	}()
	var sum int64
	for v := range ch {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToChannel(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToChannel/phase=Evaluate/ref=ChannelSend/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toChannelLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToChannel/phase=Evaluate/ref=ChannelSend/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toChannelGo(values)
			}
		})
	}
}

func TestToChannel(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6}, want: 21},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := toChannelLibrary(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := toChannelGo(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
