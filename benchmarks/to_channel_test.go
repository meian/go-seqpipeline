package benchmarks

import (
	"context"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func toChannelLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for v := range seq.ToChannel(context.Background(), source) {
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

func TestToChannelLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toChannelLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := toChannelLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := toChannelLibrary(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := toChannelLibrary(values)
		const want int64 = 60
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := toChannelLibrary(values)
		const want int64 = 2016
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := toChannelLibrary(values)
		const want int64 = 8386560
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func toChannelGo(values []int) int64 {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	var sum int64
	for v := range ch {
		sum += int64(v)
	}
	return sum
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

func TestToChannelGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toChannelGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := toChannelGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := toChannelGo(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := toChannelGo(values)
		const want int64 = 60
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := toChannelGo(values)
		const want int64 = 2016
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := toChannelGo(values)
		const want int64 = 8386560
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
