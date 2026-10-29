package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func takeOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(op.Take(8).Operator)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func takeOperatorGo(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for v := range slices.Values(values[:end]) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkTakeOperator(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Take.Operator/phase=Early8/ref=TakeSlice/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = takeOperatorLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Take.Operator/phase=Early8/ref=TakeSlice/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = takeOperatorGo(values)
			}
		})
	}
}

func TestTakeOperator(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	const want int64 = 36
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", takeOperatorLibrary},
		{"Go", takeOperatorGo},
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

func takeOperator2Library(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Take(8).Operator2)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func takeOperator2Go(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for k, v := range slices.All(values[:end]) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkTakeOperator2(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Take.Operator2/phase=Early8/ref=Take2Slice/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = takeOperator2Library(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Take.Operator2/phase=Early8/ref=Take2Slice/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = takeOperator2Go(values)
			}
		})
	}
}

func TestTakeOperator2(t *testing.T) {
	t.Parallel()

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	const want int64 = 64
	cases := []struct {
		name string
		run  func([]int) int64
	}{
		{"Library", takeOperator2Library},
		{"Go", takeOperator2Go},
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
