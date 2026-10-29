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

	values := []int{1, 2, 3, 4, 5, 6}
	const want int64 = 27
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", mapLibrary},
		{"Go", mapGo},
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

	values := []int{1, 2, 3, 4, 5, 6}
	const want int64 = 42
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", map2Library},
		{"Go", map2Go},
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
