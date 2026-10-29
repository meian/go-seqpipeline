package benchmarks

import (
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func composeSeqComposed(values []int) int64 {
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	source := seq.FromSlice(values).Pipe(transform.Operator)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func composeSeqPipe(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkComposeSeq(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/impl=Composed/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = composeSeqComposed(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = composeSeqPipe(values)
			}
		})
	}
}

func TestComposeSeq(t *testing.T) {
	t.Parallel()

	t.Run("Composed", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := composeSeqComposed(values)
		const want int64 = 72
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Pipe", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := composeSeqPipe(values)
		const want int64 = 72
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func composeSeq2Composed(values []int) int64 {
	transform := op.Compose().
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 }))
	source := seq.FromSlice(values).Pipe2(transform.Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func composeSeq2Pipe(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkComposeSeq2(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/impl=Composed/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = composeSeq2Composed(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = composeSeq2Pipe(values)
			}
		})
	}
}

func TestComposeSeq2(t *testing.T) {
	t.Parallel()

	t.Run("Composed", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := composeSeq2Composed(values)
		const want int64 = 138
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Pipe", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := composeSeq2Pipe(values)
		const want int64 = 138
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func compose2SeqComposed(values []int) int64 {
	transform := op.Compose2().
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	source := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func compose2SeqPipe(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkCompose2Seq(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/impl=Composed/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = compose2SeqComposed(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = compose2SeqPipe(values)
			}
		})
	}
}

func TestCompose2Seq(t *testing.T) {
	t.Parallel()

	t.Run("Composed", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := compose2SeqComposed(values)
		const want int64 = 90
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Pipe", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := compose2SeqPipe(values)
		const want int64 = 90
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func compose2Seq2Composed(values []int) int64 {
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 }))
	source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func compose2Seq2Pipe(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkCompose2Seq2(b *testing.B) {
	sizes := []int{64, 4096}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/impl=Composed/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = compose2Seq2Composed(values)
			}
		})
	}
	for _, n := range sizes {
		name := fmt.Sprintf("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=%d", n)
		b.Run(name, func(b *testing.B) {
			values := input(n)
			b.ReportAllocs()
			for b.Loop() {
				checksum = compose2Seq2Pipe(values)
			}
		})
	}
}

func TestCompose2Seq2(t *testing.T) {
	t.Parallel()

	t.Run("Composed", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := compose2Seq2Composed(values)
		const want int64 = 102
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Pipe", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}
		got := compose2Seq2Pipe(values)
		const want int64 = 102
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
