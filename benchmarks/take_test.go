package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func takeOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(seq.Take(8).Operator)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkTakeOperatorLibrary(b *testing.B) {
	b.Run("class=Unit/api=Take.Operator/phase=Early8/ref=TakeSlice/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperatorLibrary(values)
		}
	})
	b.Run("class=Unit/api=Take.Operator/phase=Early8/ref=TakeSlice/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperatorLibrary(values)
		}
	})
}

func TestTakeOperatorLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := takeOperatorLibrary(values)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func takeOperatorGo(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for v := range slices.Values(values[:end]) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkTakeOperatorGo(b *testing.B) {
	b.Run("class=Unit/api=Take.Operator/phase=Early8/ref=TakeSlice/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperatorGo(values)
		}
	})
	b.Run("class=Unit/api=Take.Operator/phase=Early8/ref=TakeSlice/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperatorGo(values)
		}
	})
}

func TestTakeOperatorGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := takeOperatorGo(values)
	const want int64 = 36
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func takeOperator2Library(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(seq.Take(8).Operator2)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkTakeOperator2Library(b *testing.B) {
	b.Run("class=Unit/api=Take.Operator2/phase=Early8/ref=Take2Slice/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperator2Library(values)
		}
	})
	b.Run("class=Unit/api=Take.Operator2/phase=Early8/ref=Take2Slice/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperator2Library(values)
		}
	})
}

func TestTakeOperator2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := takeOperator2Library(values)
	const want int64 = 64
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func takeOperator2Go(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for k, v := range slices.All(values[:end]) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkTakeOperator2Go(b *testing.B) {
	b.Run("class=Unit/api=Take.Operator2/phase=Early8/ref=Take2Slice/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperator2Go(values)
		}
	})
	b.Run("class=Unit/api=Take.Operator2/phase=Early8/ref=Take2Slice/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = takeOperator2Go(values)
		}
	})
}

func TestTakeOperator2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := takeOperator2Go(values)
	const want int64 = 64
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
