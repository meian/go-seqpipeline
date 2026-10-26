package benchmarks

import (
	"context"
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
	b.Run("class=Unit/api=FromChannel/phase=Evaluate/ref=ChannelReceive/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromChannelLibrary(values)
		}
	})
	b.Run("class=Unit/api=FromChannel/phase=Evaluate/ref=ChannelReceive/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromChannelLibrary(values)
		}
	})
	b.Run("class=Unit/api=FromChannel/phase=Evaluate/ref=ChannelReceive/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromChannelGo(values)
		}
	})
	b.Run("class=Unit/api=FromChannel/phase=Evaluate/ref=ChannelReceive/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromChannelGo(values)
		}
	})
}

func TestFromChannel(t *testing.T) {
	t.Run("Library", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := fromChannelLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := fromChannelGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
