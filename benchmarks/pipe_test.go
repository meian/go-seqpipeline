package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func seqPipeLibrary(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe(op.Map(func(v int) int { return (v + 1) & 65535 }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func seqPipeGo(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		sum += int64((v + 1) & 65535)
	}
	return sum
}

func BenchmarkSeqPipe(b *testing.B) {
	b.Run("class=Unit/api=Seq.Pipe/phase=Evaluate/ref=MapLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipeLibrary(values)
		}
	})
	b.Run("class=Unit/api=Seq.Pipe/phase=Evaluate/ref=MapLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipeLibrary(values)
		}
	})
	b.Run("class=Unit/api=Seq.Pipe/phase=Evaluate/ref=MapLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipeGo(values)
		}
	})
	b.Run("class=Unit/api=Seq.Pipe/phase=Evaluate/ref=MapLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipeGo(values)
		}
	})
}

func TestSeqPipe(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seqPipeLibrary(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seqPipeGo(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func seqPipe2Library(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe2(op.MapToSeq2(func(v int) (int, int) { return v, (v + 1) & 65535 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func seqPipe2Go(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		sum += int64(v) + int64((v+1)&65535)
	}
	return sum
}

func BenchmarkSeqPipe2(b *testing.B) {
	b.Run("class=Unit/api=Seq.Pipe2/phase=Evaluate/ref=ValuePair/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipe2Library(values)
		}
	})
	b.Run("class=Unit/api=Seq.Pipe2/phase=Evaluate/ref=ValuePair/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipe2Library(values)
		}
	})
	b.Run("class=Unit/api=Seq.Pipe2/phase=Evaluate/ref=ValuePair/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipe2Go(values)
		}
	})
	b.Run("class=Unit/api=Seq.Pipe2/phase=Evaluate/ref=ValuePair/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seqPipe2Go(values)
		}
	})
}

func TestSeqPipe2(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seqPipe2Library(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seqPipe2Go(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func seq2PipeLibrary(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe(op.MapToSeq(func(_, v int) int { return v }))
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func seq2PipeGo(values []int) int64 {
	var sum int64
	for _, v := range slices.All(values) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkSeq2Pipe(b *testing.B) {
	b.Run("class=Unit/api=Seq2.Pipe/phase=Evaluate/ref=Slice2Values/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2PipeLibrary(values)
		}
	})
	b.Run("class=Unit/api=Seq2.Pipe/phase=Evaluate/ref=Slice2Values/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2PipeLibrary(values)
		}
	})
	b.Run("class=Unit/api=Seq2.Pipe/phase=Evaluate/ref=Slice2Values/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2PipeGo(values)
		}
	})
	b.Run("class=Unit/api=Seq2.Pipe/phase=Evaluate/ref=Slice2Values/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2PipeGo(values)
		}
	})
}

func TestSeq2Pipe(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seq2PipeLibrary(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seq2PipeGo(values)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func seq2Pipe2Library(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe2(op.Map2(func(k, v int) (int, int) { return k, (v + 1) & 65535 }))
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func seq2Pipe2Go(values []int) int64 {
	var sum int64
	for k, v := range slices.All(values) {
		sum += int64(k) + int64((v+1)&65535)
	}
	return sum
}

func BenchmarkSeq2Pipe2(b *testing.B) {
	b.Run("class=Unit/api=Seq2.Pipe2/phase=Evaluate/ref=Map2Loop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2Pipe2Library(values)
		}
	})
	b.Run("class=Unit/api=Seq2.Pipe2/phase=Evaluate/ref=Map2Loop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2Pipe2Library(values)
		}
	})
	b.Run("class=Unit/api=Seq2.Pipe2/phase=Evaluate/ref=Map2Loop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2Pipe2Go(values)
		}
	})
	b.Run("class=Unit/api=Seq2.Pipe2/phase=Evaluate/ref=Map2Loop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = seq2Pipe2Go(values)
		}
	})
}

func TestSeq2Pipe2(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seq2Pipe2Library(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := seq2Pipe2Go(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
