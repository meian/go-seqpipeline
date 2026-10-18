package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func toSeqLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for v := range seq.ToSeq(source) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToSeqLibrary(b *testing.B) {
	b.Run("class=Unit/api=ToSeq/phase=Evaluate/ref=IteratorRange/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeqLibrary(values)
		}
	})
	b.Run("class=Unit/api=ToSeq/phase=Evaluate/ref=IteratorRange/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeqLibrary(values)
		}
	})
}

func TestToSeqLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toSeqLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func toSeqGo(values []int) int64 {
	standard := slices.Values(values)
	var sum int64
	for v := range standard {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToSeqGo(b *testing.B) {
	b.Run("class=Unit/api=ToSeq/phase=Evaluate/ref=IteratorRange/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeqGo(values)
		}
	})
	b.Run("class=Unit/api=ToSeq/phase=Evaluate/ref=IteratorRange/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeqGo(values)
		}
	})
}

func TestToSeqGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toSeqGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func toSeq2Library(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	var sum int64
	for k, v := range seq.ToSeq2(pairs) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkToSeq2Library(b *testing.B) {
	b.Run("class=Unit/api=ToSeq2/phase=Evaluate/ref=Iterator2Range/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeq2Library(values)
		}
	})
	b.Run("class=Unit/api=ToSeq2/phase=Evaluate/ref=Iterator2Range/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeq2Library(values)
		}
	})
}

func TestToSeq2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toSeq2Library(values)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func toSeq2Go(values []int) int64 {
	standardPairs := slices.All(values)
	var sum int64
	for k, v := range standardPairs {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkToSeq2Go(b *testing.B) {
	b.Run("class=Unit/api=ToSeq2/phase=Evaluate/ref=Iterator2Range/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeq2Go(values)
		}
	})
	b.Run("class=Unit/api=ToSeq2/phase=Evaluate/ref=Iterator2Range/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toSeq2Go(values)
		}
	})
}

func TestToSeq2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := toSeq2Go(values)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
