package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func skipOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for v := range seq.Skip(8).Operator(source) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkSkipOperatorLibrary(b *testing.B) {
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
}

func TestSkipOperatorLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := skipOperatorLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := skipOperatorLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := skipOperatorLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := skipOperatorLibrary(values)
		const want int64 = 30
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := skipOperatorLibrary(values)
		const want int64 = 1988
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := skipOperatorLibrary(values)
		const want int64 = 8386532
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func skipOperatorGo(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for _, v := range values[end:] {
		sum += int64(v)
	}
	return sum
}

func BenchmarkSkipOperatorGo(b *testing.B) {
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

func TestSkipOperatorGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := skipOperatorGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := skipOperatorGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := skipOperatorGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := skipOperatorGo(values)
		const want int64 = 30
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := skipOperatorGo(values)
		const want int64 = 1988
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := skipOperatorGo(values)
		const want int64 = 8386532
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func skipOperator2Library(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	var sum int64
	for k, v := range seq.Skip(8).Operator2(pairs) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkSkipOperator2Library(b *testing.B) {
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
}

func TestSkipOperator2Library(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := skipOperator2Library(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := skipOperator2Library(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := skipOperator2Library(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := skipOperator2Library(values)
		const want int64 = 57
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := skipOperator2Library(values)
		const want int64 = 3976
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := skipOperator2Library(values)
		const want int64 = 16773064
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func skipOperator2Go(values []int) int64 {
	end := min(8, len(values))
	var sum int64
	for k, v := range values[end:] {
		sum += int64(k+end) + int64(v)
	}
	return sum
}

func BenchmarkSkipOperator2Go(b *testing.B) {
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

func TestSkipOperator2Go(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := skipOperator2Go(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := skipOperator2Go(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := skipOperator2Go(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := skipOperator2Go(values)
		const want int64 = 57
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := skipOperator2Go(values)
		const want int64 = 3976
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := skipOperator2Go(values)
		const want int64 = 16773064
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
