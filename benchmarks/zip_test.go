package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func zipOperatorLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	source := inputSeq.
		Pipe2(op.Zip(seq.ToSeq(inputSeq)).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipOperatorGo(values []int) int64 {
	var sum int64
	for i, k := range slices.All(values) {
		sum += int64(k) + int64(values[i])
	}
	return sum
}

func BenchmarkZipOperator(b *testing.B) {
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

func TestZipOperator(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipOperatorLibrary(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipOperatorGo(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipAlignByLibrary(values []int) int64 {
	short := seq.FromSlice(values[:len(values)/2])
	source := seq.FromSlice(values).
		Pipe2(op.Zip(seq.ToSeq(short)).AlignBy(op.AlignLeft).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipAlignByGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, k := range slices.All(values) {
		v := 0
		if i < len(shortValues) {
			v = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipAlignBy(b *testing.B) {
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

func TestZipAlignBy(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipAlignByLibrary(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipAlignByGo(values)
		const want int64 = 27
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipWithDefaultValueLibrary(values []int) int64 {
	short := seq.FromSlice(values[:len(values)/2])
	source := seq.FromSlice(values).
		Pipe2(op.Zip(seq.ToSeq(short)).AlignBy(op.AlignLeft).WithDefaultValue(7).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipWithDefaultValueGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, k := range slices.All(values) {
		v := 7
		if i < len(shortValues) {
			v = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipWithDefaultValue(b *testing.B) {
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

func TestZipWithDefaultValue(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultValueLibrary(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultValueGo(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipWithDefaultKeyLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	source := short.
		Pipe2(op.Zip(seq.ToSeq(inputSeq)).AlignBy(op.AlignRight).WithDefaultKey(7).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipWithDefaultKeyGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, v := range slices.All(values) {
		k := 7
		if i < len(shortValues) {
			k = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipWithDefaultKey(b *testing.B) {
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

func TestZipWithDefaultKey(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultKeyLibrary(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipWithDefaultKeyGo(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyOperatorLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	source := inputSeq.
		Pipe2(op.Zip(seq.ToSeq(inputSeq)).WithDefaultKey(7).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipKeyOperatorGo(values []int) int64 {
	var sum int64
	for i, k := range slices.All(values) {
		sum += int64(k) + int64(values[i])
	}
	return sum
}

func BenchmarkZipKeyOperator(b *testing.B) {
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

func TestZipKeyOperator(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyOperatorLibrary(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyOperatorGo(values)
		const want int64 = 42
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyAlignByLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	source := short.
		Pipe2(op.Zip(seq.ToSeq(inputSeq)).WithDefaultKey(7).AlignBy(op.AlignRight).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipKeyAlignByGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, v := range slices.All(values) {
		k := 7
		if i < len(shortValues) {
			k = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyAlignBy(b *testing.B) {
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

func TestZipKeyAlignBy(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyAlignByLibrary(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyAlignByGo(values)
		const want int64 = 48
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyWithDefaultValueLibrary(values []int) int64 {
	short := seq.FromSlice(values[:len(values)/2])
	source := seq.FromSlice(values).
		Pipe2(op.Zip(seq.ToSeq(short)).WithDefaultKey(7).AlignBy(op.AlignLongest).WithDefaultValue(9).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipKeyWithDefaultValueGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, k := range slices.All(values) {
		v := 9
		if i < len(shortValues) {
			v = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyWithDefaultValue(b *testing.B) {
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

func TestZipKeyWithDefaultValue(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultValueLibrary(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultValueGo(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}

func zipKeyWithDefaultKeyLibrary(values []int) int64 {
	inputSeq := seq.FromSlice(values)
	short := seq.FromSlice(values[:len(values)/2])
	source := short.
		Pipe2(op.Zip(seq.ToSeq(inputSeq)).WithDefaultKey(7).AlignBy(op.AlignLongest).WithDefaultKey(9).Operator)
	var sum int64
	for k, v := range source {
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipKeyWithDefaultKeyGo(values []int) int64 {
	shortValues := values[:len(values)/2]
	var sum int64
	for i, v := range slices.All(values) {
		k := 9
		if i < len(shortValues) {
			k = shortValues[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func BenchmarkZipKeyWithDefaultKey(b *testing.B) {
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

func TestZipKeyWithDefaultKey(t *testing.T) {
	t.Parallel()

	t.Run("Library", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultKeyLibrary(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
	t.Run("Go", func(t *testing.T) {
		t.Parallel()

		values := []int{1, 2, 3, 4, 5, 6}

		got := zipKeyWithDefaultKeyGo(values)
		const want int64 = 54
		if got != want {
			t.Fatalf("got %d; want %d", got, want)
		}
	})
}
