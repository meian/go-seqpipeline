package benchmarks

import (
	"maps"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func keysLibrary(values map[int]int) int64 {
	source := seq.FromMap(values).
		Pipe(op.Keys)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkKeysLibrary(b *testing.B) {
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=MapKeysIterator/impl=Library/N=64", func(b *testing.B) {
		values := mapInput(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysLibrary(values)
		}
	})
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=MapKeysIterator/impl=Library/N=4096", func(b *testing.B) {
		values := mapInput(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysLibrary(values)
		}
	})
}

func TestKeysLibrary(t *testing.T) {
	values := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}

	got := keysLibrary(values)
	const want int64 = 15
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func keysGo(values map[int]int) int64 {
	var sum int64
	for k := range maps.All(values) {
		sum += int64(k)
	}
	return sum
}

func BenchmarkKeysGo(b *testing.B) {
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=MapKeysIterator/impl=Go/N=64", func(b *testing.B) {
		values := mapInput(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysGo(values)
		}
	})
	b.Run("class=Unit/api=Keys/phase=Evaluate/ref=MapKeysIterator/impl=Go/N=4096", func(b *testing.B) {
		values := mapInput(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = keysGo(values)
		}
	})
}

func TestKeysGo(t *testing.T) {
	values := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}

	got := keysGo(values)
	const want int64 = 15
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
