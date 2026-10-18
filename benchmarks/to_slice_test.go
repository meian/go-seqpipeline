package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedValues []int

func toSliceLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	collectedValues = seq.ToSlice(source)
	var sum int64
	for _, v := range collectedValues {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToSliceLibrary(b *testing.B) {
	b.Run("class=Unit/api=ToSlice/phase=Collect/ref=SlicesCollect/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSliceLibrary(values)
		}
	})
	b.Run("class=Unit/api=ToSlice/phase=Collect/ref=SlicesCollect/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSliceLibrary(values)
		}
	})
}

func TestToSliceLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toSliceLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func toSliceGo(values []int) int64 {
	standard := slices.Values(values)
	collectedValues = slices.Collect(standard)
	var sum int64
	for _, v := range collectedValues {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToSliceGo(b *testing.B) {
	b.Run("class=Unit/api=ToSlice/phase=Collect/ref=SlicesCollect/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSliceGo(values)
		}
	})
	b.Run("class=Unit/api=ToSlice/phase=Collect/ref=SlicesCollect/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSliceGo(values)
		}
	})
}

func TestToSliceGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toSliceGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
