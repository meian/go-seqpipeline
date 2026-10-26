package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func filterLibrary(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(op.Filter(func(v int) bool { return v%3 != 0 }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func filterGo(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		if v%3 != 0 {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkFilter(b *testing.B) {
	b.Run("class=Unit/api=Filter/phase=Evaluate/ref=FilterLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filterLibrary(values)
		}
	})
	b.Run("class=Unit/api=Filter/phase=Evaluate/ref=FilterLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filterLibrary(values)
		}
	})
	b.Run("class=Unit/api=Filter/phase=Evaluate/ref=FilterLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filterGo(values)
		}
	})
	b.Run("class=Unit/api=Filter/phase=Evaluate/ref=FilterLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filterGo(values)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := filterLibrary(values)
		const want int64 = 12
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := filterGo(values)
		const want int64 = 12
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func filter2Library(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Filter2(func(k, v int) bool { return v%3 != 0 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func filter2Go(values []int) int64 {
	var sum int64
	for k, v := range slices.All(values) {
		if v%3 != 0 {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func BenchmarkFilter2(b *testing.B) {
	b.Run("class=Unit/api=Filter2/phase=Evaluate/ref=Filter2Loop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filter2Library(values)
		}
	})
	b.Run("class=Unit/api=Filter2/phase=Evaluate/ref=Filter2Loop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filter2Library(values)
		}
	})
	b.Run("class=Unit/api=Filter2/phase=Evaluate/ref=Filter2Loop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filter2Go(values)
		}
	})
	b.Run("class=Unit/api=Filter2/phase=Evaluate/ref=Filter2Loop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = filter2Go(values)
		}
	})
}

func TestFilter2(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := filter2Library(values)
		const want int64 = 20
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := filter2Go(values)
		const want int64 = 20
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
