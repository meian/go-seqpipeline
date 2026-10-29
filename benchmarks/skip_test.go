package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func skipOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(op.Skip(8).Operator)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func skipOperatorGo(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for v := range slices.Values(values[end:]) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkSkipOperator(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Skip.Operator/phase=Evaluate/ref=SkipSlice/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = skipOperatorLibrary(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Skip.Operator/phase=Evaluate/ref=SkipSlice/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = skipOperatorGo(values)
			}
		})
	}
}

func TestSkipOperator(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 19},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := skipOperatorLibrary(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := skipOperatorGo(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func skipOperator2Library(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Skip(8).Operator2)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func skipOperator2Go(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for k, v := range slices.All(values[end:]) {
		sum += int64(k+end) + int64(v)
	}
	return sum
}

func BenchmarkSkipOperator2(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Skip.Operator2/phase=Evaluate/ref=Skip2Slice/impl=Library/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = skipOperator2Library(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Skip.Operator2/phase=Evaluate/ref=Skip2Slice/impl=Go/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)

			b.ReportAllocs()
			for b.Loop() {
				checksum = skipOperator2Go(values)
			}
		})
	}
}

func TestSkipOperator2(t *testing.T) {
	t.Parallel()

	cases := []struct {
		values []int
		want   int64
	}{
		{values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 36},
	}
	for _, tc := range cases {
		t.Run("Library", func(t *testing.T) {
			t.Parallel()

			got := skipOperator2Library(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
	for _, tc := range cases {
		t.Run("Go", func(t *testing.T) {
			t.Parallel()

			got := skipOperator2Go(slices.Clone(tc.values))
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}
