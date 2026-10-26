package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromSeqLibrary(values []int) int64 {
	standard := slices.Values(values)
	source := seq.FromSeq(standard)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func fromSeqGo(values []int) int64 {
	standard := slices.Values(values)
	var sum int64
	for v := range standard {
		sum += int64(v)
	}
	return sum
}

func BenchmarkFromSeq(b *testing.B) {
	b.Run("class=Unit/api=FromSeq/phase=Evaluate/ref=IteratorRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeqLibrary(values)
		}
	})
	b.Run("class=Unit/api=FromSeq/phase=Evaluate/ref=IteratorRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeqLibrary(values)
		}
	})
	b.Run("class=Unit/api=FromSeq/phase=Evaluate/ref=IteratorRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeqGo(values)
		}
	})
	b.Run("class=Unit/api=FromSeq/phase=Evaluate/ref=IteratorRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeqGo(values)
		}
	})
}

func TestFromSeq(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := fromSeqLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := fromSeqGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func fromSeq2Library(values []int) int64 {
	standardPairs := slices.All(values)
	source := seq.FromSeq2(standardPairs)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func fromSeq2Go(values []int) int64 {
	standardPairs := slices.All(values)
	var sum int64
	for k, v := range standardPairs {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkFromSeq2(b *testing.B) {
	b.Run("class=Unit/api=FromSeq2/phase=Evaluate/ref=Iterator2Range/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeq2Library(values)
		}
	})
	b.Run("class=Unit/api=FromSeq2/phase=Evaluate/ref=Iterator2Range/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeq2Library(values)
		}
	})
	b.Run("class=Unit/api=FromSeq2/phase=Evaluate/ref=Iterator2Range/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeq2Go(values)
		}
	})
	b.Run("class=Unit/api=FromSeq2/phase=Evaluate/ref=Iterator2Range/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = fromSeq2Go(values)
		}
	})
}

func TestFromSeq2(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := fromSeq2Library(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := fromSeq2Go(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
