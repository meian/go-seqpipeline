package benchmarks

import (
	"maps"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedMap map[int]int

func toMapLibrary(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	collectedMap = seq.ToMap(pairs)
	var sum int64
	for k, v := range collectedMap {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkToMapLibrary(b *testing.B) {
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapLibrary(values)
		}
	})
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapLibrary(values)
		}
	})
}

func TestToMapLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toMapLibrary(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := toMapLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := toMapLibrary(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := toMapLibrary(values)
		const want int64 = 115
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := toMapLibrary(values)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := toMapLibrary(values)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func toMapGo(values []int) int64 {
	standardPairs := slices.All(values)
	collectedMap = maps.Collect(standardPairs)
	var sum int64
	for k, v := range collectedMap {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkToMapGo(b *testing.B) {
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapGo(values)
		}
	})
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapGo(values)
		}
	})
}

func TestToMapGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toMapGo(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := toMapGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := toMapGo(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := toMapGo(values)
		const want int64 = 115
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := toMapGo(values)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := toMapGo(values)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
