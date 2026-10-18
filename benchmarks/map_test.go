package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func mapLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for v := range seq.Map(func(v int) int { return (v + 1) & 65535 })(source) {
		sum += int64(v)
	}
	return sum
}

func BenchmarkMapLibrary(b *testing.B) {
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapLibrary(values)
		}
	})
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapLibrary(values)
		}
	})
}

func TestMapLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := mapLibrary(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := mapLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := mapLibrary(values)
		const want int64 = 8
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := mapLibrary(values)
		const want int64 = 131143
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := mapLibrary(values)
		const want int64 = 2080
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := mapLibrary(values)
		const want int64 = 8390656
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func mapGo(values []int) int64 {
	var sum int64
	for _, v := range values {
		sum += int64((v + 1) & 65535)
	}
	return sum
}

func BenchmarkMapGo(b *testing.B) {
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapGo(values)
		}
	})
	b.Run("class=Unit/api=Map/phase=Evaluate/ref=MapLoop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = mapGo(values)
		}
	})
}

func TestMapGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := mapGo(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := mapGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := mapGo(values)
		const want int64 = 8
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := mapGo(values)
		const want int64 = 131143
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := mapGo(values)
		const want int64 = 2080
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := mapGo(values)
		const want int64 = 8390656
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func map2Library(values []int) int64 {
	pairs := seq.FromSeq2(slices.All(values))
	var sum int64
	for k, v := range seq.Map2(func(k, v int) (int, int) { return k, (v + 1) & 65535 })(pairs) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkMap2Library(b *testing.B) {
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Library(values)
		}
	})
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Library(values)
		}
	})
}

func TestMap2Library(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := map2Library(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := map2Library(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := map2Library(values)
		const want int64 = 8
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := map2Library(values)
		const want int64 = 131198
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := map2Library(values)
		const want int64 = 4096
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := map2Library(values)
		const want int64 = 16777216
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func map2Go(values []int) int64 {
	var sum int64
	for k, v := range values {
		sum += int64(k) + int64((v+1)&65535)
	}
	return sum
}

func BenchmarkMap2Go(b *testing.B) {
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Go(values)
		}
	})
	b.Run("class=Unit/api=Map2/phase=Evaluate/ref=Map2Loop/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = map2Go(values)
		}
	})
}

func TestMap2Go(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := map2Go(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := map2Go(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := map2Go(values)
		const want int64 = 8
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := map2Go(values)
		const want int64 = 131198
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := map2Go(values)
		const want int64 = 4096
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := map2Go(values)
		const want int64 = 16777216
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
