package benchmarks

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func zipOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for k, v := range seq.Zip(source).Operator(source) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipOperatorLibrary(b *testing.B) {
	b.Run("class=Unit/api=Zip.Operator/phase=Evaluate/ref=ZipIndex/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipOperatorLibrary(values)
		}
	})
	b.Run("class=Unit/api=Zip.Operator/phase=Evaluate/ref=ZipIndex/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipOperatorLibrary(values)
		}
	})
}

func TestZipOperatorLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipOperatorLibrary(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipOperatorLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipOperatorLibrary(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipOperatorLibrary(values)
		const want int64 = 120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipOperatorLibrary(values)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipOperatorLibrary(values)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipOperatorGo(values []int) int64 {
	var sum int64
	for i, k := range values {
		sum += int64(k) + int64(values[i])
	}
	return sum
}

func BenchmarkZipOperatorGo(b *testing.B) {
	b.Run("class=Unit/api=Zip.Operator/phase=Evaluate/ref=ZipIndex/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipOperatorGo(values)
		}
	})
	b.Run("class=Unit/api=Zip.Operator/phase=Evaluate/ref=ZipIndex/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipOperatorGo(values)
		}
	})
}

func TestZipOperatorGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipOperatorGo(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipOperatorGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipOperatorGo(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipOperatorGo(values)
		const want int64 = 120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipOperatorGo(values)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipOperatorGo(values)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipAlignByLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	var sum int64
	for k, v := range seq.Zip(short).AlignBy(seq.AlignLeft).Operator(source) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipAlignByLibrary(b *testing.B) {
	b.Run("class=Unit/api=Zip.AlignBy/phase=Evaluate/ref=ZipLeft0/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipAlignByLibrary(values)
		}
	})
	b.Run("class=Unit/api=Zip.AlignBy/phase=Evaluate/ref=ZipLeft0/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipAlignByLibrary(values)
		}
	})
}

func TestZipAlignByLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipAlignByLibrary(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipAlignByLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipAlignByLibrary(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipAlignByLibrary(values)
		const want int64 = 88
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipAlignByLibrary(values)
		const want int64 = 2512
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipAlignByLibrary(values)
		const want int64 = 10482688
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipAlignByGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, k := range values {
		v := 0
		if i < len(shortValues) {
			v = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipAlignByGo(b *testing.B) {
	b.Run("class=Unit/api=Zip.AlignBy/phase=Evaluate/ref=ZipLeft0/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipAlignByGo(values)
		}
	})
	b.Run("class=Unit/api=Zip.AlignBy/phase=Evaluate/ref=ZipLeft0/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipAlignByGo(values)
		}
	})
}

func TestZipAlignByGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipAlignByGo(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipAlignByGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipAlignByGo(values)
		const want int64 = 7
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipAlignByGo(values)
		const want int64 = 88
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipAlignByGo(values)
		const want int64 = 2512
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipAlignByGo(values)
		const want int64 = 10482688
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipWithDefaultValueLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	var sum int64
	for k, v := range seq.Zip(short).AlignBy(seq.AlignLeft).WithDefaultValue(7).Operator(source) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipWithDefaultValueLibrary(b *testing.B) {
	b.Run("class=Unit/api=Zip.WithDefaultValue/phase=Evaluate/ref=ZipLeft7/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultValueLibrary(values)
		}
	})
	b.Run("class=Unit/api=Zip.WithDefaultValue/phase=Evaluate/ref=ZipLeft7/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultValueLibrary(values)
		}
	})
}

func TestZipWithDefaultValueLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 130
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 2736
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 10497024
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipWithDefaultValueGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, k := range values {
		v := 7
		if i < len(shortValues) {
			v = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipWithDefaultValueGo(b *testing.B) {
	b.Run("class=Unit/api=Zip.WithDefaultValue/phase=Evaluate/ref=ZipLeft7/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultValueGo(values)
		}
	})
	b.Run("class=Unit/api=Zip.WithDefaultValue/phase=Evaluate/ref=ZipLeft7/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultValueGo(values)
		}
	})
}

func TestZipWithDefaultValueGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultValueGo(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipWithDefaultValueGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipWithDefaultValueGo(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipWithDefaultValueGo(values)
		const want int64 = 130
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipWithDefaultValueGo(values)
		const want int64 = 2736
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipWithDefaultValueGo(values)
		const want int64 = 10497024
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipWithDefaultKeyLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	var sum int64
	for k, v := range seq.Zip(source).AlignBy(seq.AlignRight).WithDefaultKey(7).Operator(short) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipWithDefaultKeyLibrary(b *testing.B) {
	b.Run("class=Unit/api=Zip.WithDefaultKey/phase=Evaluate/ref=ZipRight7/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultKeyLibrary(values)
		}
	})
	b.Run("class=Unit/api=Zip.WithDefaultKey/phase=Evaluate/ref=ZipRight7/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultKeyLibrary(values)
		}
	})
}

func TestZipWithDefaultKeyLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 130
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 2736
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 10497024
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipWithDefaultKeyGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, v := range values {
		k := 7
		if i < len(shortValues) {
			k = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipWithDefaultKeyGo(b *testing.B) {
	b.Run("class=Unit/api=Zip.WithDefaultKey/phase=Evaluate/ref=ZipRight7/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultKeyGo(values)
		}
	})
	b.Run("class=Unit/api=Zip.WithDefaultKey/phase=Evaluate/ref=ZipRight7/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipWithDefaultKeyGo(values)
		}
	})
}

func TestZipWithDefaultKeyGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultKeyGo(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipWithDefaultKeyGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipWithDefaultKeyGo(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipWithDefaultKeyGo(values)
		const want int64 = 130
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipWithDefaultKeyGo(values)
		const want int64 = 2736
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipWithDefaultKeyGo(values)
		const want int64 = 10497024
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyOperatorLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	var sum int64
	for k, v := range seq.Zip(source).WithDefaultKey(7).Operator(source) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyOperatorLibrary(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.Operator/phase=Evaluate/ref=ZipIndex/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyOperatorLibrary(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.Operator/phase=Evaluate/ref=ZipIndex/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyOperatorLibrary(values)
		}
	})
}

func TestZipKeyOperatorLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyOperatorLibrary(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyOperatorLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyOperatorLibrary(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyOperatorLibrary(values)
		const want int64 = 120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyOperatorLibrary(values)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyOperatorLibrary(values)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyOperatorGo(values []int) int64 {
	var sum int64
	for i, k := range values {
		sum += int64(k) + int64(values[i])
	}
	return sum
}

func BenchmarkZipKeyOperatorGo(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.Operator/phase=Evaluate/ref=ZipIndex/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyOperatorGo(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.Operator/phase=Evaluate/ref=ZipIndex/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyOperatorGo(values)
		}
	})
}

func TestZipKeyOperatorGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyOperatorGo(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyOperatorGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyOperatorGo(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyOperatorGo(values)
		const want int64 = 120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyOperatorGo(values)
		const want int64 = 4032
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyOperatorGo(values)
		const want int64 = 16773120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyAlignByLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	var sum int64
	for k, v := range seq.Zip(source).WithDefaultKey(7).AlignBy(seq.AlignRight).Operator(short) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyAlignByLibrary(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.AlignBy/phase=Evaluate/ref=ZipRight7/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyAlignByLibrary(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.AlignBy/phase=Evaluate/ref=ZipRight7/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyAlignByLibrary(values)
		}
	})
}

func TestZipKeyAlignByLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyAlignByLibrary(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyAlignByLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyAlignByLibrary(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyAlignByLibrary(values)
		const want int64 = 130
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyAlignByLibrary(values)
		const want int64 = 2736
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyAlignByLibrary(values)
		const want int64 = 10497024
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyAlignByGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, v := range values {
		k := 7
		if i < len(shortValues) {
			k = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyAlignByGo(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.AlignBy/phase=Evaluate/ref=ZipRight7/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyAlignByGo(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.AlignBy/phase=Evaluate/ref=ZipRight7/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyAlignByGo(values)
		}
	})
}

func TestZipKeyAlignByGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyAlignByGo(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyAlignByGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyAlignByGo(values)
		const want int64 = 14
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyAlignByGo(values)
		const want int64 = 130
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyAlignByGo(values)
		const want int64 = 2736
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyAlignByGo(values)
		const want int64 = 10497024
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyWithDefaultValueLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	var sum int64
	for k, v := range seq.Zip(short).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultValue(9).Operator(source) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyWithDefaultValueLibrary(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.WithDefaultValue/phase=Evaluate/ref=ZipLeft9/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultValueLibrary(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.WithDefaultValue/phase=Evaluate/ref=ZipLeft9/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultValueLibrary(values)
		}
	})
}

func TestZipKeyWithDefaultValueLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 16
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 142
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 2800
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 10501120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyWithDefaultValueGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, k := range values {
		v := 9
		if i < len(shortValues) {
			v = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyWithDefaultValueGo(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.WithDefaultValue/phase=Evaluate/ref=ZipLeft9/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultValueGo(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.WithDefaultValue/phase=Evaluate/ref=ZipLeft9/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultValueGo(values)
		}
	})
}

func TestZipKeyWithDefaultValueGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 16
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 142
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 2800
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 10501120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyWithDefaultKeyLibrary(values []int) int64 {
	source := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	var sum int64
	for k, v := range seq.Zip(source).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultKey(9).Operator(short) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyWithDefaultKeyLibrary(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.WithDefaultKey/phase=Evaluate/ref=ZipRight9/impl=Library/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultKeyLibrary(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.WithDefaultKey/phase=Evaluate/ref=ZipRight9/impl=Library/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultKeyLibrary(values)
		}
	})
}

func TestZipKeyWithDefaultKeyLibrary(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 16
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 142
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 2800
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 10501120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyWithDefaultKeyGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, v := range values {
		k := 9
		if i < len(shortValues) {
			k = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyWithDefaultKeyGo(b *testing.B) {
	b.Run("class=Unit/api=ZipKey.WithDefaultKey/phase=Evaluate/ref=ZipRight9/impl=Go/N=64", func(b *testing.B) {
		values := input(64)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultKeyGo(values)
		}
	})
	b.Run("class=Unit/api=ZipKey.WithDefaultKey/phase=Evaluate/ref=ZipRight9/impl=Go/N=4096", func(b *testing.B) {
		values := input(4096)

		b.ReportAllocs()
		for b.Loop() {
			checksum = zipKeyWithDefaultKeyGo(values)
		}
	})
}

func TestZipKeyWithDefaultKeyGo(t *testing.T) {
	t.Run("Values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		values := []int(nil)

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 0
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Single", func(t *testing.T) {
		values := []int{7}

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 16
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		values := []int{17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 142
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N64", func(t *testing.T) {
		values := input(64)

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 2800
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("N4096", func(t *testing.T) {
		values := input(4096)

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 10501120
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
