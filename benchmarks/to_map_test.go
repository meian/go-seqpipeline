package benchmarks

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedMap map[int]int

func toMapLibrary(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	collectedMap = seq.ToMap(pairs)
	var sum int64
	for k, v := range maps.All(collectedMap) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func toMapGo(values []int) int64 {
	standardPairs := slices.All(values)
	collectedMap = maps.Collect(standardPairs)
	var sum int64
	for k, v := range maps.All(collectedMap) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkToMap(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toMapLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toMapGo(values)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	t.Parallel()

	// 検証対象がベンチマーク用の collectedMap を共有するため、子テストは直列に実行する。
	cases := []struct {
		values []int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6}, want: 36},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			got := toMapLibrary(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			got := toMapGo(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
