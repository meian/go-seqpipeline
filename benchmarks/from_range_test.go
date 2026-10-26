package benchmarks

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func fromRangeSignedLibrary(start, stop int, step int64) int64 {
	source := seq.FromRange(start, stop, step)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func fromRangeSignedGo(start, stop int, step int64) int64 {
	var sum int64
	if step > 0 {
		for v := start; v <= stop; v += int(step) {
			sum += int64(v)
		}
	} else {
		for v := start; v >= stop; v -= int(-step) {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkFromRangeSigned(b *testing.B) {
	b.Run("class=Unit/api=FromRange.Signed/phase=Ascending/ref=IntegerLoop/impl=Library/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedLibrary(1, 64, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Ascending/ref=IntegerLoop/impl=Library/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedLibrary(1, 4096, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Descending/ref=IntegerLoop/impl=Library/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedLibrary(64, 1, -1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Descending/ref=IntegerLoop/impl=Library/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedLibrary(4096, 1, -1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Ascending/ref=IntegerLoop/impl=Go/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedGo(1, 64, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Ascending/ref=IntegerLoop/impl=Go/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedGo(1, 4096, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Descending/ref=IntegerLoop/impl=Go/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedGo(64, 1, -1)
		}
	})
	b.Run("class=Unit/api=FromRange.Signed/phase=Descending/ref=IntegerLoop/impl=Go/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeSignedGo(4096, 1, -1)
		}
	})
}

func TestFromRangeSigned(t *testing.T) {
	t.Parallel()

	t.Run("Library/Ascending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeSignedLibrary(1, 6, 1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Library/Descending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeSignedLibrary(6, 1, -1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go/Ascending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeSignedGo(1, 6, 1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go/Descending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeSignedGo(6, 1, -1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func fromRangeUnsignedLibrary(start, stop uint, step int64) int64 {
	source := seq.FromRange(start, stop, step)
	var sum int64
	for v := range source {
		sum += int64(v)
	}
	return sum
}

func fromRangeUnsignedGo(start, stop uint, step int64) int64 {
	var sum int64
	if step > 0 {
		for v := start; v <= stop; v += uint(step) {
			sum += int64(v)
		}
	} else {
		for v := start; v >= stop; v -= uint(-step) {
			sum += int64(v)
		}
	}
	return sum
}

func BenchmarkFromRangeUnsigned(b *testing.B) {
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Ascending/ref=IntegerLoop/impl=Library/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedLibrary(1, 64, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Ascending/ref=IntegerLoop/impl=Library/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedLibrary(1, 4096, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Descending/ref=IntegerLoop/impl=Library/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedLibrary(64, 1, -1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Descending/ref=IntegerLoop/impl=Library/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedLibrary(4096, 1, -1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Ascending/ref=IntegerLoop/impl=Go/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedGo(1, 64, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Ascending/ref=IntegerLoop/impl=Go/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedGo(1, 4096, 1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Descending/ref=IntegerLoop/impl=Go/N=64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedGo(64, 1, -1)
		}
	})
	b.Run("class=Unit/api=FromRange.Unsigned/phase=Descending/ref=IntegerLoop/impl=Go/N=4096", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			checksum = fromRangeUnsignedGo(4096, 1, -1)
		}
	})
}

func TestFromRangeUnsigned(t *testing.T) {
	t.Parallel()

	t.Run("Library/Ascending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeUnsignedLibrary(1, 6, 1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Library/Descending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeUnsignedLibrary(6, 1, -1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go/Ascending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeUnsignedGo(1, 6, 1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go/Descending", func(t *testing.T) {
		t.Parallel()

		got := fromRangeUnsignedGo(6, 1, -1)
		const want int64 = 21
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
