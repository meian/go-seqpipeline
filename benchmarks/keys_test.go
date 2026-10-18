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
	values := []int{1, 2, 3, 4, 5, 6}

	got := keysLibrary(values)
	const want int64 = 15
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
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
	values := []int{1, 2, 3, 4, 5, 6}

	got := keysGo(values)
	const want int64 = 15
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
