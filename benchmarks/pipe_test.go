package benchmarks

import (
	"iter"
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

func BenchmarkSeqPipeLibrary(b *testing.B) {
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
}

func TestSeqPipeLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seqPipeLibrary(values)
	const want int64 = 27
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func seqPipeGo(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		sum += int64((v + 1) & 65535)
	}
	return sum
}

func BenchmarkSeqPipeGo(b *testing.B) {
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

func TestSeqPipeGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seqPipeGo(values)
	const want int64 = 27
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func seqPipe2Library(values []int) int64 {
	source := seq.FromSlice(values).
		Pipe2(func(input iter.Seq[int]) iter.Seq2[int, int] {
			return func(yield func(int, int) bool) {
				for v := range input {
					if !yield(v, (v+1)&65535) {
						return
					}
				}
			}
		})
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkSeqPipe2Library(b *testing.B) {
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
}

func TestSeqPipe2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seqPipe2Library(values)
	const want int64 = 48
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func seqPipe2Go(values []int) int64 {
	var sum int64
	for v := range slices.Values(values) {
		sum += int64(v) + int64((v+1)&65535)
	}
	return sum
}

func BenchmarkSeqPipe2Go(b *testing.B) {
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

func TestSeqPipe2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seqPipe2Go(values)
	const want int64 = 48
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func seq2PipeLibrary(values []int) int64 {
	source := seq.FromSeq2(slices.All(values)).
		Pipe(op.Values)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func BenchmarkSeq2PipeLibrary(b *testing.B) {
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
}

func TestSeq2PipeLibrary(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seq2PipeLibrary(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func seq2PipeGo(values []int) int64 {
	var sum int64
	for _, v := range slices.All(values) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkSeq2PipeGo(b *testing.B) {
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

func TestSeq2PipeGo(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seq2PipeGo(values)
	const want int64 = 21
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
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

func BenchmarkSeq2Pipe2Library(b *testing.B) {
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
}

func TestSeq2Pipe2Library(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seq2Pipe2Library(values)
	const want int64 = 42
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}

func seq2Pipe2Go(values []int) int64 {
	var sum int64
	for k, v := range slices.All(values) {
		sum += int64(k) + int64((v+1)&65535)
	}
	return sum
}

func BenchmarkSeq2Pipe2Go(b *testing.B) {
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

func TestSeq2Pipe2Go(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}

	got := seq2Pipe2Go(values)
	const want int64 = 42
	if got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
}
