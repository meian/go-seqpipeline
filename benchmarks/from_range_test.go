// #nosec G115 -- 固定入力は 1〜4096、step は ±1 で、整数変換先の範囲に収まる。
package benchmarks

import (
	"fmt"
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
	cases := []struct {
		phase string
		n     int
		start int
		stop  int
		step  int64
	}{
		{"Ascending", 64, 1, 64, 1},
		{"Ascending", 4096, 1, 4096, 1},
		{"Descending", 64, 64, 1, -1},
		{"Descending", 4096, 4096, 1, -1},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=Unit/api=FromRange.Signed/phase=%s/ref=IntegerLoop/impl=Library/N=%d", tc.phase, tc.n)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				checksum = fromRangeSignedLibrary(tc.start, tc.stop, tc.step)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=Unit/api=FromRange.Signed/phase=%s/ref=IntegerLoop/impl=Go/N=%d", tc.phase, tc.n)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				checksum = fromRangeSignedGo(tc.start, tc.stop, tc.step)
			}
		})
	}
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
	cases := []struct {
		phase string
		n     int
		start uint
		stop  uint
		step  int64
	}{
		{"Ascending", 64, 1, 64, 1},
		{"Ascending", 4096, 1, 4096, 1},
		{"Descending", 64, 64, 1, -1},
		{"Descending", 4096, 4096, 1, -1},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=Unit/api=FromRange.Unsigned/phase=%s/ref=IntegerLoop/impl=Library/N=%d", tc.phase, tc.n)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				checksum = fromRangeUnsignedLibrary(tc.start, tc.stop, tc.step)
			}
		})
	}
	for _, tc := range cases {
		name := fmt.Sprintf("class=Unit/api=FromRange.Unsigned/phase=%s/ref=IntegerLoop/impl=Go/N=%d", tc.phase, tc.n)
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				checksum = fromRangeUnsignedGo(tc.start, tc.stop, tc.step)
			}
		})
	}
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
