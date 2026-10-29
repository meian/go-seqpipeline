package benchmarks

import (
	"fmt"
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

func keysGo(values map[int]int) int64 {
	var sum int64
	for k := range maps.All(values) {
		sum += int64(k)
	}
	return sum
}

func BenchmarkKeys(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Keys/phase=Evaluate/ref=MapKeysIterator/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := mapInput(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = keysLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Keys/phase=Evaluate/ref=MapKeysIterator/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := mapInput(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = keysGo(values)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values map[int]int
		want   int64
	}{
		{values: map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5, 5: 6}, want: 15},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := keysLibrary(maps.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := keysGo(maps.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
