package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func valuesLibrary(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	var sum int64
	for v := range seq.Values(pairs) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkValuesLibrary(b *testing.B) {
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=SliceRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesLibrary(values)
		}
	})
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=SliceRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesLibrary(values)
		}
	})
}

func TestValuesLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := valuesLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func valuesGo(values []int) int64 {
	var sum int64
	for _, v := range values {
		sum += int64(v)
	}
	return sum
}

func BenchmarkValuesGo(b *testing.B) {
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=SliceRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesGo(values)
		}
	})
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=SliceRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesGo(values)
		}
	})
}

func TestValuesGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := valuesGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
