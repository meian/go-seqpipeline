package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

// 構築を計測から除く。nil の b は通常テスト用。
func composeIterations(b *testing.B) int {
	if b == nil {
		return 1
	}
	b.ResetTimer()
	return b.N
}

func composeSeqComposed(values []int, b *testing.B) (sum int64, output []int) {
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	for range composeIterations(b) {
		source := seq.FromSlice(values).Pipe(transform.Operator)
		if b == nil {
			for v := range source {
				output = append(output, v)
			}
		}
		for v := range source {
			sum += int64(v)
		}
	}
	return sum, output
}

func composeSeqPipe(values []int, b *testing.B) (sum int64, output []int) {
	step1 := op.Map(func(v int) int { return v + 1 })
	step2 := op.Map(func(v int) int { return v * 2 })
	step3 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step4 := op.MapToSeq(func(k, v int) int { return k + v })
	for range composeIterations(b) {
		source := seq.FromSlice(values).
			Pipe(step1).
			Pipe(step2).
			Pipe2(step3).
			Pipe(step4)
		if b == nil {
			for v := range source {
				output = append(output, v)
			}
		}
		for v := range source {
			sum += int64(v)
		}
	}
	return sum, output
}

func BenchmarkComposeSeq(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = composeSeqPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = composeSeqComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = composeSeqPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = composeSeqComposed(values, b)
	})
}

// 順序・重複を含む出力列と加算結果を、実際の測定関数で両実装とも確認する。
func TestComposeSeq(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	want := []int{19, 11, 19, -1}
	composedSum, composed := composeSeqComposed(values, nil)
	pipeSum, pipe := composeSeqPipe(values, nil)
	if !slices.Equal(composed, want) || !slices.Equal(pipe, want) {
		t.Fatalf("Composed %v; Pipe %v; want %v", composed, pipe, want)
	}
	const wantSum int64 = 48
	if composedSum != wantSum || pipeSum != wantSum {
		t.Fatalf("Composed sum %d; Pipe sum %d; want %d", composedSum, pipeSum, wantSum)
	}
}

func composeSeq2Composed(values []int, b *testing.B) (sum int64, output []pair) {
	transform := op.Compose().
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
	for range composeIterations(b) {
		source := seq.FromSlice(values).Pipe2(transform.Operator)
		if b == nil {
			for k, v := range source {
				output = append(output, pair{k, v})
			}
		}
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum, output
}

func composeSeq2Pipe(values []int, b *testing.B) (sum int64, output []pair) {
	step1 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step2 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
	step3 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
	step4 := op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 })
	for range composeIterations(b) {
		source := seq.FromSlice(values).
			Pipe2(step1).
			Pipe2(step2).
			Pipe2(step3).
			Pipe2(step4)
		if b == nil {
			for k, v := range source {
				output = append(output, pair{k, v})
			}
		}
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum, output
}

func BenchmarkComposeSeq2(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = composeSeq2Pipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = composeSeq2Composed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = composeSeq2Pipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = composeSeq2Composed(values, b)
	})
}

// 順序・重複を含む出力列と加算結果を、実際の測定関数で両実装とも確認する。
func TestComposeSeq2(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	want := []pair{{11, 30}, {7, 22}, {11, 30}, {1, 10}}
	composedSum, composed := composeSeq2Composed(values, nil)
	pipeSum, pipe := composeSeq2Pipe(values, nil)
	if !slices.Equal(composed, want) || !slices.Equal(pipe, want) {
		t.Fatalf("Composed %v; Pipe %v; want %v", composed, pipe, want)
	}
	const wantSum int64 = 122
	if composedSum != wantSum || pipeSum != wantSum {
		t.Fatalf("Composed sum %d; Pipe sum %d; want %d", composedSum, pipeSum, wantSum)
	}
}

func compose2SeqComposed(values []int, b *testing.B) (sum int64, output []int) {
	transform := op.Compose2().
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	for range composeIterations(b) {
		source := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		if b == nil {
			for v := range source {
				output = append(output, v)
			}
		}
		for v := range source {
			sum += int64(v)
		}
	}
	return sum, output
}

func compose2SeqPipe(values []int, b *testing.B) (sum int64, output []int) {
	step1 := op.MapToSeq(func(k, v int) int { return k + v })
	step2 := op.Map(func(v int) int { return v * 2 })
	step3 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step4 := op.MapToSeq(func(k, v int) int { return k + v })
	for range composeIterations(b) {
		source := seq.FromSeq2(slices.All(values)).
			Pipe(step1).
			Pipe(step2).
			Pipe2(step3).
			Pipe(step4)
		if b == nil {
			for v := range source {
				output = append(output, v)
			}
		}
		for v := range source {
			sum += int64(v)
		}
	}
	return sum, output
}

func BenchmarkCompose2Seq(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = compose2SeqPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = compose2SeqComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = compose2SeqPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = compose2SeqComposed(values, b)
	})
}

// 順序・重複を含む出力列と加算結果を、実際の測定関数で両実装とも確認する。
func TestCompose2Seq(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	want := []int{15, 11, 23, 7}
	composedSum, composed := compose2SeqComposed(values, nil)
	pipeSum, pipe := compose2SeqPipe(values, nil)
	if !slices.Equal(composed, want) || !slices.Equal(pipe, want) {
		t.Fatalf("Composed %v; Pipe %v; want %v", composed, pipe, want)
	}
	const wantSum int64 = 56
	if composedSum != wantSum || pipeSum != wantSum {
		t.Fatalf("Composed sum %d; Pipe sum %d; want %d", composedSum, pipeSum, wantSum)
	}
}

func compose2Seq2Composed(values []int, b *testing.B) (sum int64, output []pair) {
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
	for range composeIterations(b) {
		source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		if b == nil {
			for k, v := range source {
				output = append(output, pair{k, v})
			}
		}
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum, output
}

func compose2Seq2Pipe(values []int, b *testing.B) (sum int64, output []pair) {
	step1 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
	step2 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
	step3 := op.MapToSeq(func(k, v int) int { return k + v })
	step4 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	for range composeIterations(b) {
		source := seq.FromSeq2(slices.All(values)).
			Pipe2(step1).
			Pipe2(step2).
			Pipe(step3).
			Pipe2(step4)
		if b == nil {
			for k, v := range source {
				output = append(output, pair{k, v})
			}
		}
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum, output
}

func BenchmarkCompose2Seq2(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = compose2Seq2Pipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum, _ = compose2Seq2Composed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = compose2Seq2Pipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum, _ = compose2Seq2Composed(values, b)
	})
}

// 順序・重複を含む出力列と加算結果を、実際の測定関数で両実装とも確認する。
func TestCompose2Seq2(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 3, -2}
	want := []pair{{11, 14}, {9, 12}, {15, 18}, {7, 10}}
	composedSum, composed := compose2Seq2Composed(values, nil)
	pipeSum, pipe := compose2Seq2Pipe(values, nil)
	if !slices.Equal(composed, want) || !slices.Equal(pipe, want) {
		t.Fatalf("Composed %v; Pipe %v; want %v", composed, pipe, want)
	}
	const wantSum int64 = 96
	if composedSum != wantSum || pipeSum != wantSum {
		t.Fatalf("Composed sum %d; Pipe sum %d; want %d", composedSum, pipeSum, wantSum)
	}
}
