package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func composeSeqLibrary(values []int) int64 {
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

func BenchmarkComposeSeqLibrary(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=ComposedLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqLibrary(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=ComposedLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqLibrary(values)
		}
	})
}

func TestComposeSeqLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeqLibrary(values)
	const want int64 = 72
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func composeSeqGo(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		v++
		v++
		k, v := v, v+1
		sum += int64(k + v)
	}
	return sum
}

func BenchmarkComposeSeqGo(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=ComposedLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqGo(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq/phase=Evaluate/ref=ComposedLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeqGo(values)
		}
	})
}

func TestComposeSeqGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeqGo(values)
	const want int64 = 72
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func composeSeq2Library(values []int) int64 {
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

func BenchmarkComposeSeq2Library(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Library(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Library(values)
		}
	})
}

func TestComposeSeq2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeq2Library(values)
	const want int64 = 54
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func composeSeq2Go(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		k, v := v, v+1
		v++
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkComposeSeq2Go(b *testing.B) {
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Go(values)
		}
	})
	b.Run("class=Unit/api=Compose.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = composeSeq2Go(values)
		}
	})
}

func TestComposeSeq2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := composeSeq2Go(values)
	const want int64 = 54
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func compose2SeqLibrary(values []int) int64 {
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

func BenchmarkCompose2SeqLibrary(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=ComposedLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqLibrary(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=ComposedLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqLibrary(values)
		}
	})
}

func TestCompose2SeqLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2SeqLibrary(values)
	const want int64 = 90
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func compose2SeqGo(values []int) int64 {
	var sum int64
	for k, v := range slices.All(values) {
		v = k + v
		v++
		k, v = v, v+1
		sum += int64(k + v)
	}
	return sum
}

func BenchmarkCompose2SeqGo(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=ComposedLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqGo(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq/phase=Evaluate/ref=ComposedLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2SeqGo(values)
		}
	})
}

func TestCompose2SeqGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2SeqGo(values)
	const want int64 = 90
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func compose2Seq2Library(values []int) int64 {
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

func BenchmarkCompose2Seq2Library(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Library(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Library(values)
		}
	})
}

func TestCompose2Seq2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2Seq2Library(values)
	const want int64 = 48
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func compose2Seq2Go(values []int) int64 {
	var sum int64
	for k, v := range slices.All(values) {
		v++
		v++
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkCompose2Seq2Go(b *testing.B) {
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Go(values)
		}
	})
	b.Run("class=Unit/api=Compose2.Seq2/phase=Evaluate/ref=ComposedLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)
		b.ReportAllocs()
		for b.Loop() {
			checksum = compose2Seq2Go(values)
		}
	})
}

func TestCompose2Seq2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	got := compose2Seq2Go(values)
	const want int64 = 48
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
