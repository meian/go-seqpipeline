package benchmarks

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromMapLibrary(values map[int]int) int64 {
	var sum int64
	for k, v := range seq.FromMap(values) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkFromMapLibrary(b *testing.B) {
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapLibrary(mapping)
		}
	})
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapLibrary(mapping)
		}
	})
}

func TestFromMapLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapLibrary(mapping)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapLibrary(mapping)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapLibrary(mapping)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapLibrary(mapping)
		const want int64 = 115
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapLibrary(mapping)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapLibrary(mapping)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func fromMapGo(values map[int]int) int64 {
	var sum int64
	for k, v := range values {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkFromMapGo(b *testing.B) {
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapGo(mapping)
		}
	})
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapGo(mapping)
		}
	})
}

func TestFromMapGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapGo(mapping)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapGo(mapping)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapGo(mapping)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapGo(mapping)
		const want int64 = 115
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapGo(mapping)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)
		mapping := make(map[int]int, len(values))
		for k, v := range values {
			mapping[k] = v
		}
		got := fromMapGo(mapping)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
