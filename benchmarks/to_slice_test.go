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
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toSliceLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := toSliceLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := toSliceLibrary(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := toSliceLibrary(values)
		const want int64 = 60
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := toSliceLibrary(values)
		const want int64 = 2016
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := toSliceLibrary(values)
		const want int64 = 8386560
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
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
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toSliceGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := toSliceGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := toSliceGo(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := toSliceGo(values)
		const want int64 = 60
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := toSliceGo(values)
		const want int64 = 2016
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := toSliceGo(values)
		const want int64 = 8386560
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
