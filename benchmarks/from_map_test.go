package benchmarks

import (
	"maps"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromMapLibrary(values map[int]int) int64 {
	source := seq.FromMap(values)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkFromMapLibrary(b *testing.B) {
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Library/N=64", func(b *testing.B) {
		mapping := mapInput(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapLibrary(mapping)
		}
	})
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Library/N=4096", func(b *testing.B) {
		mapping := mapInput(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapLibrary(mapping)
		}
	})
}

func TestFromMapLibrary(t *testing.T) {
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
}

func fromMapGo(values map[int]int) int64 {
	var sum int64
	for k, v := range maps.All(values) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkFromMapGo(b *testing.B) {
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Go/N=64", func(b *testing.B) {
		mapping := mapInput(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapGo(mapping)
		}
	})
	b.Run("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Go/N=4096", func(b *testing.B) {
		mapping := mapInput(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromMapGo(mapping)
		}
	})
}

func TestFromMapGo(t *testing.T) {
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
}
