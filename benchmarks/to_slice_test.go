package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedValues []int

func toSliceLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	collectedValues = seq.ToSlice(source)
	var sum int64
	for v := range slices.Values(collectedValues) {
		sum += int64(v)
	}
	return sum
}

func toSliceGo(values []int) int64 {
	standard := slices.Values(values)
	collectedValues = slices.Collect(standard)
	var sum int64
	for v := range slices.Values(collectedValues) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkToSlice(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToSlice/phase=Collect/ref=SlicesCollect/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toSliceLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=ToSlice/phase=Collect/ref=SlicesCollect/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = toSliceGo(values)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	// 検証対象がベンチマーク用の collectedValues を共有するため、子テストは直列に実行する。
	t.Run("Library", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toSliceLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := toSliceGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
