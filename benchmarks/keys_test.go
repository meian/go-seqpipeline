package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func keysLibrary(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	var sum int64
	for v := range seq.Keys(pairs) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkKeysLibrary(b *testing.B) {
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=KeyLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysLibrary(values)
		}
	})
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=KeyLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysLibrary(values)
		}
	})
}

func TestKeysLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := keysLibrary(values)
		const want int64 = 15
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := keysLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := keysLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := keysLibrary(values)
		const want int64 = 55
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := keysLibrary(values)
		const want int64 = 2016
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := keysLibrary(values)
		const want int64 = 8386560
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func keysGo(values []int) int64 {
	var sum int64
	for k := range values {
		sum += int64(k)
	}
	return sum
}

func BenchmarkKeysGo(b *testing.B) {
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=KeyLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysGo(values)
		}
	})
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=KeyLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysGo(values)
		}
	})
}

func TestKeysGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := keysGo(values)
		const want int64 = 15
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := keysGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := keysGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := keysGo(values)
		const want int64 = 55
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := keysGo(values)
		const want int64 = 2016
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := keysGo(values)
		const want int64 = 8386560
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
