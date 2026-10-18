package benchmarks

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromSliceLibrary(values []int) int64 {
	var sum int64
	for v := range seq.FromSlice(values) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkFromSliceLibrary(b *testing.B) {
	b.Run("class=Unit/api=FromSlice/phase=Evaluate/ref=SliceRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSliceLibrary(values)
		}
	})
	b.Run("class=Unit/api=FromSlice/phase=Evaluate/ref=SliceRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSliceLibrary(values)
		}
	})
}

func TestFromSliceLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := fromSliceLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func fromSliceGo(values []int) int64 {
	var sum int64
	for _, v := range values {
		sum += int64(v)
	}
	return sum
}

func BenchmarkFromSliceGo(b *testing.B) {
	b.Run("class=Unit/api=FromSlice/phase=Evaluate/ref=SliceRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSliceGo(values)
		}
	})
	b.Run("class=Unit/api=FromSlice/phase=Evaluate/ref=SliceRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSliceGo(values)
		}
	})
}

func TestFromSliceGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := fromSliceGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
