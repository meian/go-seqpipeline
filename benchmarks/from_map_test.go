package benchmarks

import (
	"fmt"
	"maps"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromMapLibrary(values map[int]int) int64 {
	source := seq.FromMap(values)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func fromMapGo(values map[int]int) int64 {
	var sum int64
	for k, v := range maps.All(values) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkFromMap(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			mapping := mapInput(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = fromMapLibrary(mapping)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=FromMap/phase=Evaluate/ref=MapRange/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			mapping := mapInput(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = fromMapGo(mapping)
			}
		})
	}
}

func TestFromMap(t *testing.T) {
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

			mapping := make(map[int]int, len(tc.values))
			for k, v := range tc.values {
				mapping[k] = v
			}
			got := fromMapLibrary(mapping)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			mapping := make(map[int]int, len(tc.values))
			for k, v := range tc.values {
				mapping[k] = v
			}
			got := fromMapGo(mapping)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
