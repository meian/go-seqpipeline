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

func BenchmarkComposeSeqComposed(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqComposed(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqComposed(values)
		}
	})
}

func BenchmarkComposeSeqPipe(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqPipe(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqPipe(values)
		}
	})
}

func TestComposeSeqComposed(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeqComposed(values)
	const want int64 = 72
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestComposeSeqPipe(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeqPipe(values)
	const want int64 = 72
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func composeSeq2Composed(values []int) int64 {
	transform := op.Compose().
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 }))
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
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkComposeSeq2Composed(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Composed(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Composed(values)
		}
	})
}

func BenchmarkComposeSeq2Pipe(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Pipe(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Pipe(values)
		}
	})
}

func TestComposeSeq2Composed(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeq2Composed(values)
	const want int64 = 54
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestComposeSeq2Pipe(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeq2Pipe(values)
	const want int64 = 54
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
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

func BenchmarkCompose2SeqComposed(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqComposed(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqComposed(values)
		}
	})
}

func BenchmarkCompose2SeqPipe(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqPipe(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqPipe(values)
		}
	})
}

func TestCompose2SeqComposed(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2SeqComposed(values)
	const want int64 = 90
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestCompose2SeqPipe(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2SeqPipe(values)
	const want int64 = 90
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func compose2Seq2Composed(values []int) int64 {
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 }))
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
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, v + 1 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkCompose2Seq2Composed(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Composed(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Composed(values)
		}
	})
}

func BenchmarkCompose2Seq2Pipe(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Pipe(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Pipe(values)
		}
	})
}

func TestCompose2Seq2Composed(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2Seq2Composed(values)
	const want int64 = 48
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func TestCompose2Seq2Pipe(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2Seq2Pipe(values)
	const want int64 = 48
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
