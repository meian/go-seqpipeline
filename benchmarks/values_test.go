package benchmarks

import (
	"maps"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func valuesLibrary(values map[int]int) int64 {
	source := seq.FromMap(values).
		Pipe(op.Values)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkValuesLibrary(b *testing.B) {
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=MapValuesIterator/impl=Library/N=64", func(b *testing.B) {
		values := mapInput(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesLibrary(values)
		}
	})
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=MapValuesIterator/impl=Library/N=4096", func(b *testing.B) {
		values := mapInput(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesLibrary(values)
		}
	})
}

func TestValuesLibrary(t *testing.T) {
	values := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}

	got := valuesLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func valuesGo(values map[int]int) int64 {
	var sum int64
	for _, v := range maps.All(values) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkValuesGo(b *testing.B) {
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=MapValuesIterator/impl=Go/N=64", func(b *testing.B) {
		values := mapInput(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesGo(values)
		}
	})
	b.Run("class=Unit/api=Values/phase=Evaluate/ref=MapValuesIterator/impl=Go/N=4096", func(b *testing.B) {
		values := mapInput(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = valuesGo(values)
		}
	})
}

func TestValuesGo(t *testing.T) {
	values := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}

	got := valuesGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
