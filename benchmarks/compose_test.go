package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func composeSeqComposed(values []int) int64 {
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
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
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkComposeSeq(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqPipe(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqComposed(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqPipe(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqComposed(values)
		}
	})
}

func TestComposeSeq(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	composed := composeSeqComposed(values)
	pipe := composeSeqPipe(values)
	const want int64 = 48
	if composed != want || pipe != want {
		t.Fatalf("Composed %d; Pipe %d; want %d", composed, pipe, want)
	}
}

func composeSeq2Composed(values []int) int64 {
	transform := op.Compose().
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
	source := seq.FromSlice(values).Pipe2(transform.Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func composeSeq2Pipe(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkComposeSeq2(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Pipe(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Composed(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Pipe(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Composed(values)
		}
	})
}

func TestComposeSeq2(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	composed := composeSeq2Composed(values)
	pipe := composeSeq2Pipe(values)
	const want int64 = 122
	if composed != want || pipe != want {
		t.Fatalf("Composed %d; Pipe %d; want %d", composed, pipe, want)
	}
}

func compose2SeqComposed(values []int) int64 {
	transform := op.Compose2().
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
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
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkCompose2Seq(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqPipe(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqComposed(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqPipe(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqComposed(values)
		}
	})
}

func TestCompose2Seq(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	composed := compose2SeqComposed(values)
	pipe := compose2SeqPipe(values)
	const want int64 = 56
	if composed != want || pipe != want {
		t.Fatalf("Composed %d; Pipe %d; want %d", composed, pipe, want)
	}
}

func compose2Seq2Composed(values []int) int64 {
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
	source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func compose2Seq2Pipe(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkCompose2Seq2(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Pipe(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Composed(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Pipe(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Composed(values)
		}
	})
}

func TestCompose2Seq2(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	composed := compose2Seq2Composed(values)
	pipe := compose2Seq2Pipe(values)
	const want int64 = 96
	if composed != want || pipe != want {
		t.Fatalf("Composed %d; Pipe %d; want %d", composed, pipe, want)
	}
}
