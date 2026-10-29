package benchmarks

import (
	"fmt"
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
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Filter/phase=Evaluate/ref=FilterLoop/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = filterLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Filter/phase=Evaluate/ref=FilterLoop/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = filterGo(values)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	const want int64 = 12
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", filterLibrary},
		{"Go", filterGo},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.run(slices.Clone(values))
			if got != want {
				t.Fatalf("got %d; want %d", got, want)
			}
		})
	}
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
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Filter2/phase=Evaluate/ref=Filter2Loop/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = filter2Library(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Filter2/phase=Evaluate/ref=Filter2Loop/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = filter2Go(values)
			}
		})
	}
}

func TestFilter2(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6}
	const want int64 = 20
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", filter2Library},
		{"Go", filter2Go},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.run(slices.Clone(values))
			if got != want {
				t.Fatalf("got %d; want %d", got, want)
			}
		})
	}
}
