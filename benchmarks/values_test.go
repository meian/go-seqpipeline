package benchmarks

import (
	"fmt"
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

func valuesGo(values map[int]int) int64 {
	var sum int64
	for _, v := range maps.All(values) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkValues(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Values/phase=Evaluate/ref=MapValuesIterator/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := mapInput(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = valuesLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Values/phase=Evaluate/ref=MapValuesIterator/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := mapInput(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = valuesGo(values)
			}
		})
	}
}

func TestValues(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values map[int]int
		want   int64
	}{
		{values: map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}, want: 21},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := valuesLibrary(maps.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := valuesGo(maps.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
