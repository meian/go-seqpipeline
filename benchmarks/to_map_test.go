package benchmarks

import (
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
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapLibrary(values)
		}
	})
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapLibrary(values)
		}
	})
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapGo(values)
		}
	})
	b.Run("class=Unit/api=ToMap/phase=Collect/ref=MapsCollect/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = toMapGo(values)
		}
	})
}

func TestToMap(t *testing.T) {
	t.Parallel()

	// 検証対象がベンチマーク用の collectedMap を共有するため、子テストは直列に実行する。
	t.Run("Library", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toMapLibrary(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toMapGo(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
