package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func mapLibrary(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(op.Map(func(v int) int { return (v + 1) & 65535 }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func mapGo(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		sum += int64((v + 1) & 65535)
	}
	return sum
}

func BenchmarkMapLibrary(b *testing.B) {
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapLibrary(values)
		}
	})
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapLibrary(values)
		}
	})
}

func BenchmarkMapGo(b *testing.B) {
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapGo(values)
		}
	})
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapGo(values)
		}
	})
}

func TestMapLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := mapLibrary(values)
	const want int64 = 27
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestMapGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := mapGo(values)
	const want int64 = 27
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func map2Library(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, (v + 1) & 65535 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func map2Go(values []int) int64 {
	var sum int64
	for k, v := range slices.All(values) {
		sum += int64(k) + int64((v+1)&65535)
	}
	return sum
}

func BenchmarkMap2Library(b *testing.B) {
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Library(values)
		}
	})
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Library(values)
		}
	})
}

func BenchmarkMap2Go(b *testing.B) {
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Go(values)
		}
	})
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Go(values)
		}
	})
}

func TestMap2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := map2Library(values)
	const want int64 = 42
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestMap2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := map2Go(values)
	const want int64 = 42
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
