package benchmarks

import (
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
	b.Run("class=Unit/api=Skip.Operator/phase=Evaluate/ref=SkipSlice/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperatorLibrary(values)
		}
	})
	b.Run("class=Unit/api=Skip.Operator/phase=Evaluate/ref=SkipSlice/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperatorLibrary(values)
		}
	})
	b.Run("class=Unit/api=Skip.Operator/phase=Evaluate/ref=SkipSlice/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperatorGo(values)
		}
	})
	b.Run("class=Unit/api=Skip.Operator/phase=Evaluate/ref=SkipSlice/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperatorGo(values)
		}
	})
}

func TestSkipOperator(t *testing.T) {
	t.Run("Library", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := skipOperatorLibrary(values)
		const want int64 = 19
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := skipOperatorGo(values)
		const want int64 = 19
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
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
	b.Run("class=Unit/api=Skip.Operator2/phase=Evaluate/ref=Skip2Slice/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperator2Library(values)
		}
	})
	b.Run("class=Unit/api=Skip.Operator2/phase=Evaluate/ref=Skip2Slice/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperator2Library(values)
		}
	})
	b.Run("class=Unit/api=Skip.Operator2/phase=Evaluate/ref=Skip2Slice/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperator2Go(values)
		}
	})
	b.Run("class=Unit/api=Skip.Operator2/phase=Evaluate/ref=Skip2Slice/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = skipOperator2Go(values)
		}
	})
}

func TestSkipOperator2(t *testing.T) {
	t.Run("Library", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := skipOperator2Library(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := skipOperator2Go(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
