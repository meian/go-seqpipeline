package benchmarks

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromChannelLibrary(values []int) int64 {
	ch, done := producer(values, 0)
	source := seq.FromChannel(context.Background(), ch)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	<-done
	return sum
}

func fromChannelGo(values []int) int64 {
	ch, done := producer(values, 0)
	var sum int64
	for v := range ch {
		sum += int64(v)
	}
	<-done
	return sum
}

func BenchmarkFromChannel(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=FromChannel/phase=Evaluate/ref=ChannelReceive/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = fromChannelLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=FromChannel/phase=Evaluate/ref=ChannelReceive/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = fromChannelGo(values)
			}
		})
	}
}

func TestFromChannel(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	const want int64 = 21
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", fromChannelLibrary},
		{"Go", fromChannelGo},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.run(slices.Clone(values))
			if got != want {
				t.Fatalf("got %d; want %d", got, want)
			}
		})
	}
}
