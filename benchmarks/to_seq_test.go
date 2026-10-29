package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func toSeqLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	source := seq.ToSeq(inputSeq)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func toSeqGo(values []int) int64 {
	standard := slices.Values(values)
	var sum int64
	for v := range standard {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToSeq(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToSeq/phase=Evaluate/ref=IteratorRange/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toSeqLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToSeq/phase=Evaluate/ref=IteratorRange/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toSeqGo(values)
			}
		})
	}
}

func TestToSeq(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6}, want: 21},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := toSeqLibrary(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := toSeqGo(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func toSeq2Library(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	source := seq.ToSeq2(pairs)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func toSeq2Go(values []int) int64 {
	standardPairs := slices.All(values)
	var sum int64
	for k, v := range standardPairs {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkToSeq2(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToSeq2/phase=Evaluate/ref=Iterator2Range/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toSeq2Library(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToSeq2/phase=Evaluate/ref=Iterator2Range/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toSeq2Go(values)
			}
		})
	}
}

func TestToSeq2(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6}, want: 36},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := toSeq2Library(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := toSeq2Go(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
