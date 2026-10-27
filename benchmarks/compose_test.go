package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

// オペレーターを関数引数経由で間接呼び出しせず、同じ関数内で構築・反復する。
// b.N ループの前で計測をリセットし、nil の b は通常テストでの 1 回評価に使う。

func composeSeqDepth2BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose().
			Pipe(op.Map(func(v int) int { return v + 1 })).
			Pipe(op.Map(func(v int) int { return v * 2 }))
		source := seq.FromSlice(values).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth2BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe(op.Map(func(v int) int { return v + 1 })).
			Pipe(op.Map(func(v int) int { return v * 2 }))
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth2ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe(op.Map(func(v int) int { return v * 2 }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth2ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.Map(func(v int) int { return v + 1 })
	step2 := op.Map(func(v int) int { return v * 2 })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe(step1).
			Pipe(step2)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth4BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose().
			Pipe(op.Map(func(v int) int { return v + 1 })).
			Pipe(op.Map(func(v int) int { return v * 2 })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
		source := seq.FromSlice(values).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth4BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe(op.Map(func(v int) int { return v + 1 })).
			Pipe(op.Map(func(v int) int { return v * 2 })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth4ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { return v + 1 })).
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func composeSeqDepth4ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.Map(func(v int) int { return v + 1 })
	step2 := op.Map(func(v int) int { return v * 2 })
	step3 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step4 := op.MapToSeq(func(k, v int) int { return k + v })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe(step1).
			Pipe(step2).
			Pipe2(step3).
			Pipe(step4)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkComposeSeq(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeqDepth4ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeqDepth4ReuseComposed(values, b)
	})
}

// 順序・重複を含む出力と、構築込み・構築済みの測定関数が返す加算結果を確認する。
func TestComposeSeq(t *testing.T) {
	t.Parallel()
	t.Run("depth=2/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose().
			Pipe(op.Map(func(v int) int { return v + 1 })).
			Pipe(op.Map(func(v int) int { return v * 2 }))
		step1 := op.Map(func(v int) int { return v + 1 })
		step2 := op.Map(func(v int) int { return v * 2 })
		composed := seq.FromSlice(values).Pipe(transform.Operator)
		pipe := seq.FromSlice(values).
			Pipe(step1).
			Pipe(step2)
		want := []int{8, 4, 8, -2}
		var composedOutput, pipeOutput []int
		for v := range composed {
			composedOutput = append(composedOutput, v)
		}
		for v := range pipe {
			pipeOutput = append(pipeOutput, v)
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 18
		if got := composeSeqDepth2BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeqDepth2BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := composeSeqDepth2ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeqDepth2ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
	t.Run("depth=4/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose().
			Pipe(op.Map(func(v int) int { return v + 1 })).
			Pipe(op.Map(func(v int) int { return v * 2 })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
		step1 := op.Map(func(v int) int { return v + 1 })
		step2 := op.Map(func(v int) int { return v * 2 })
		step3 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
		step4 := op.MapToSeq(func(k, v int) int { return k + v })
		composed := seq.FromSlice(values).Pipe(transform.Operator)
		pipe := seq.FromSlice(values).
			Pipe(step1).
			Pipe(step2).
			Pipe2(step3).
			Pipe(step4)
		want := []int{19, 11, 19, -1}
		var composedOutput, pipeOutput []int
		for v := range composed {
			composedOutput = append(composedOutput, v)
		}
		for v := range pipe {
			pipeOutput = append(pipeOutput, v)
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 48
		if got := composeSeqDepth4BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeqDepth4BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := composeSeqDepth4ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeqDepth4ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
}

func composeSeq2Depth2BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose().
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 }))
		source := seq.FromSlice(values).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth2BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 }))
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth2ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose().
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth2ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step2 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe2(step1).
			Pipe2(step2)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth4BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose().
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
		source := seq.FromSlice(values).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth4BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth4ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose().
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func composeSeq2Depth4ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step2 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
	step3 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
	step4 := op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSlice(values).
			Pipe2(step1).
			Pipe2(step2).
			Pipe2(step3).
			Pipe2(step4)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func BenchmarkComposeSeq2(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = composeSeq2Depth4ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = composeSeq2Depth4ReuseComposed(values, b)
	})
}

// 順序・重複を含む出力と、構築込み・構築済みの測定関数が返す加算結果を確認する。
func TestComposeSeq2(t *testing.T) {
	t.Parallel()
	t.Run("depth=2/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose().
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 }))
		step1 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
		step2 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
		composed := seq.FromSlice(values).Pipe2(transform.Operator)
		pipe := seq.FromSlice(values).
			Pipe2(step1).
			Pipe2(step2)
		want := []pair{{4, 12}, {2, 8}, {4, 12}, {-1, 2}}
		var composedOutput, pipeOutput []pair
		for k, v := range composed {
			composedOutput = append(composedOutput, pair{k, v})
		}
		for k, v := range pipe {
			pipeOutput = append(pipeOutput, pair{k, v})
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 43
		if got := composeSeq2Depth2BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeq2Depth2BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := composeSeq2Depth2ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeq2Depth2ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
	t.Run("depth=4/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose().
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 }))
		step1 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
		step2 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
		step3 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
		step4 := op.Map2(func(k, v int) (int, int) { return k + 3, v * 2 })
		composed := seq.FromSlice(values).Pipe2(transform.Operator)
		pipe := seq.FromSlice(values).
			Pipe2(step1).
			Pipe2(step2).
			Pipe2(step3).
			Pipe2(step4)
		want := []pair{{11, 30}, {7, 22}, {11, 30}, {1, 10}}
		var composedOutput, pipeOutput []pair
		for k, v := range composed {
			composedOutput = append(composedOutput, pair{k, v})
		}
		for k, v := range pipe {
			pipeOutput = append(pipeOutput, pair{k, v})
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 122
		if got := composeSeq2Depth4BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeq2Depth4BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := composeSeq2Depth4ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := composeSeq2Depth4ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
}

func compose2SeqDepth2BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose2().
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe(op.Map(func(v int) int { return v * 2 }))
		source := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth2BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe(op.Map(func(v int) int { return v * 2 }))
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth2ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose2().
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe(op.Map(func(v int) int { return v * 2 }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth2ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.MapToSeq(func(k, v int) int { return k + v })
	step2 := op.Map(func(v int) int { return v * 2 })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe(step1).
			Pipe(step2)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth4BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose2().
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe(op.Map(func(v int) int { return v * 2 })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
		source := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth4BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe(op.Map(func(v int) int { return v * 2 })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth4ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose2().
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe(op.Map(func(v int) int { return v * 2 })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func compose2SeqDepth4ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.MapToSeq(func(k, v int) int { return k + v })
	step2 := op.Map(func(v int) int { return v * 2 })
	step3 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	step4 := op.MapToSeq(func(k, v int) int { return k + v })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe(step1).
			Pipe(step2).
			Pipe2(step3).
			Pipe(step4)
		for v := range source {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkCompose2Seq(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2SeqDepth4ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2SeqDepth4ReuseComposed(values, b)
	})
}

// 順序・重複を含む出力と、構築込み・構築済みの測定関数が返す加算結果を確認する。
func TestCompose2Seq(t *testing.T) {
	t.Parallel()
	t.Run("depth=2/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose2().
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe(op.Map(func(v int) int { return v * 2 }))
		step1 := op.MapToSeq(func(k, v int) int { return k + v })
		step2 := op.Map(func(v int) int { return v * 2 })
		composed := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		pipe := seq.FromSeq2(slices.All(values)).
			Pipe(step1).
			Pipe(step2)
		want := []int{6, 4, 10, 2}
		var composedOutput, pipeOutput []int
		for v := range composed {
			composedOutput = append(composedOutput, v)
		}
		for v := range pipe {
			pipeOutput = append(pipeOutput, v)
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 22
		if got := compose2SeqDepth2BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2SeqDepth2BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := compose2SeqDepth2ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2SeqDepth2ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
	t.Run("depth=4/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose2().
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe(op.Map(func(v int) int { return v * 2 })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v }))
		step1 := op.MapToSeq(func(k, v int) int { return k + v })
		step2 := op.Map(func(v int) int { return v * 2 })
		step3 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
		step4 := op.MapToSeq(func(k, v int) int { return k + v })
		composed := seq.FromSeq2(slices.All(values)).Pipe(transform.Operator)
		pipe := seq.FromSeq2(slices.All(values)).
			Pipe(step1).
			Pipe(step2).
			Pipe2(step3).
			Pipe(step4)
		want := []int{15, 11, 23, 7}
		var composedOutput, pipeOutput []int
		for v := range composed {
			composedOutput = append(composedOutput, v)
		}
		for v := range pipe {
			pipeOutput = append(pipeOutput, v)
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 56
		if got := compose2SeqDepth4BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2SeqDepth4BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := compose2SeqDepth4ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2SeqDepth4ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
}

func compose2Seq2Depth2BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose2().
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 }))
		source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth2BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 }))
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth2ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth2ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
	step2 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe2(step1).
			Pipe2(step2)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth4BuildComposed(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		transform := op.Compose2().
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
		source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth4BuildPipe(values []int, b *testing.B) int64 {
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth4ReuseComposed(values []int, b *testing.B) int64 {
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
		Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func compose2Seq2Depth4ReusePipe(values []int, b *testing.B) int64 {
	step1 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
	step2 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
	step3 := op.MapToSeq(func(k, v int) int { return k + v })
	step4 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
	iterations := 1
	if b != nil {
		iterations = b.N
		b.ResetTimer()
	}
	var sum int64
	for range iterations {
		source := seq.FromSeq2(slices.All(values)).
			Pipe2(step1).
			Pipe2(step2).
			Pipe(step3).
			Pipe2(step4)
		for k, v := range source {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func BenchmarkCompose2Seq2(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=2/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth2ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4BuildPipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Build/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4BuildComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4ReuseComposed(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Pipe/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4ReusePipe(values, b)
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=DirectPipe/depth=4/scope=Reuse/impl=Composed/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		checksum = compose2Seq2Depth4ReuseComposed(values, b)
	})
}

// 順序・重複を含む出力と、構築込み・構築済みの測定関数が返す加算結果を確認する。
func TestCompose2Seq2(t *testing.T) {
	t.Parallel()
	t.Run("depth=2/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose2().
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 }))
		step1 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
		step2 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
		composed := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		pipe := seq.FromSeq2(slices.All(values)).
			Pipe2(step1).
			Pipe2(step2)
		want := []pair{{2, 9}, {4, 5}, {6, 9}, {8, -1}}
		var composedOutput, pipeOutput []pair
		for k, v := range composed {
			composedOutput = append(composedOutput, pair{k, v})
		}
		for k, v := range pipe {
			pipeOutput = append(pipeOutput, pair{k, v})
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 42
		if got := compose2Seq2Depth2BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2Seq2Depth2BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := compose2Seq2Depth2ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2Seq2Depth2ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
	t.Run("depth=4/BuildAndReuse/PipeAndComposed", func(t *testing.T) {
		t.Parallel()
		values := []int{3, 1, 3, -2}
		transform := op.Compose2().
			Pipe2(op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })).
			Pipe2(op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })).
			Pipe(op.MapToSeq(func(k, v int) int { return k + v })).
			Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, v + 3 }))
		step1 := op.Map2(func(k, v int) (int, int) { return k + 1, v * 2 })
		step2 := op.Map2(func(k, v int) (int, int) { return k * 2, v + 3 })
		step3 := op.MapToSeq(func(k, v int) int { return k + v })
		step4 := op.MapToSeq2(func(v int) (int, int) { return v, v + 3 })
		composed := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
		pipe := seq.FromSeq2(slices.All(values)).
			Pipe2(step1).
			Pipe2(step2).
			Pipe(step3).
			Pipe2(step4)
		want := []pair{{11, 14}, {9, 12}, {15, 18}, {7, 10}}
		var composedOutput, pipeOutput []pair
		for k, v := range composed {
			composedOutput = append(composedOutput, pair{k, v})
		}
		for k, v := range pipe {
			pipeOutput = append(pipeOutput, pair{k, v})
		}
		if !slices.Equal(composedOutput, want) || !slices.Equal(pipeOutput, want) {
			t.Fatalf("Composed %v; Pipe %v; want %v", composedOutput, pipeOutput, want)
		}
		const wantSum int64 = 96
		if got := compose2Seq2Depth4BuildComposed(values, nil); got != wantSum {
			t.Fatalf("Build/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2Seq2Depth4BuildPipe(values, nil); got != wantSum {
			t.Fatalf("Build/Pipe sum %d; want %d", got, wantSum)
		}
		if got := compose2Seq2Depth4ReuseComposed(values, nil); got != wantSum {
			t.Fatalf("Reuse/Composed sum %d; want %d", got, wantSum)
		}
		if got := compose2Seq2Depth4ReusePipe(values, nil); got != wantSum {
			t.Fatalf("Reuse/Pipe sum %d; want %d", got, wantSum)
		}
	})
}
