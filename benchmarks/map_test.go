package benchmarks

import (
	"fmt"
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

func BenchmarkMap(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = mapLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = mapGo(values)
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := mapLibrary(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := mapGo(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
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

func BenchmarkMap2(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = map2Library(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = map2Go(values)
			}
		})
	}
}

func TestMap2(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := map2Library(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := map2Go(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
