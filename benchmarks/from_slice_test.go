package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromSliceLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func fromSliceGo(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkFromSlice(b *testing.B) {
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

func TestFromSlice(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := fromSliceLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := fromSliceGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
