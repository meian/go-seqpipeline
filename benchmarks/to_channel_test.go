package benchmarks

import (
	"context"
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

func BenchmarkToChannelLibrary(b *testing.B) {
	b.Run("class=Unit/api=ToChannel/phase=Evaluate/ref=ChannelSend/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toChannelLibrary(values)
		}
	})
	b.Run("class=Unit/api=ToChannel/phase=Evaluate/ref=ChannelSend/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toChannelLibrary(values)
		}
	})
}

func BenchmarkToChannelGo(b *testing.B) {
	b.Run("class=Unit/api=ToChannel/phase=Evaluate/ref=ChannelSend/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toChannelGo(values)
		}
	})
	b.Run("class=Unit/api=ToChannel/phase=Evaluate/ref=ChannelSend/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toChannelGo(values)
		}
	})
}

func TestToChannelLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toChannelLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestToChannelGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toChannelGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
