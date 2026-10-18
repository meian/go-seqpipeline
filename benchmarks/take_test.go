package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func takeOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for v := range seq.Take(8).Operator(source) {
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
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := takeOperatorLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := takeOperatorLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := takeOperatorLibrary(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := takeOperatorLibrary(values)
		const want int64 = 30
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := takeOperatorLibrary(values)
		const want int64 = 28
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := takeOperatorLibrary(values)
		const want int64 = 28
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func takeOperatorGo(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for _, v := range values[:end] {
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
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := takeOperatorGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := takeOperatorGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := takeOperatorGo(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := takeOperatorGo(values)
		const want int64 = 30
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := takeOperatorGo(values)
		const want int64 = 28
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := takeOperatorGo(values)
		const want int64 = 28
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func takeOperator2Library(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	var sum int64
	for k, v := range seq.Take(8).Operator2(pairs) {
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
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := takeOperator2Library(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := takeOperator2Library(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := takeOperator2Library(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := takeOperator2Library(values)
		const want int64 = 58
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := takeOperator2Library(values)
		const want int64 = 56
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := takeOperator2Library(values)
		const want int64 = 56
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func takeOperator2Go(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for k, v := range values[:end] {
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
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := takeOperator2Go(values)
		const want int64 = 36
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := takeOperator2Go(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := takeOperator2Go(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := takeOperator2Go(values)
		const want int64 = 58
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := takeOperator2Go(values)
		const want int64 = 56
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := takeOperator2Go(values)
		const want int64 = 56
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
