package benchmarks

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func zipUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	// 不足する左右の値を実際に通し、設定メソッドを評価費用と一緒に測る。
	short := seq.FromSlice(values[:len(values)/2])
	var total, shortTotal int64
	for _, v := range values {
		total += int64(v)
	}
	for _, v := range values[:len(values)/2] {
		shortTotal += int64(v)
	}
	missing := int64(len(values) - len(values)/2)
	data := &unitInput{
		values: values,
		source: source,
		short:  short,
	}
	return []unitCase{
		{"Zip.Operator", "Evaluate", unitRun{data, runUnitZipOperatorEvaluateLibrary}, total * 2},
		{"Zip.AlignBy", "Evaluate", unitRun{data, runUnitZipAlignByEvaluateLibrary}, total + shortTotal},
		{"Zip.WithDefaultValue", "Evaluate", unitRun{data, runUnitZipWithDefaultValueEvaluateLibrary}, total + shortTotal + missing*7},
		{"Zip.WithDefaultKey", "Evaluate", unitRun{data, runUnitZipWithDefaultKeyEvaluateLibrary}, total + shortTotal + missing*7},
		{"ZipKey.Operator", "Evaluate", unitRun{data, runUnitZipKeyOperatorEvaluateLibrary}, total * 2},
		{"ZipKey.AlignBy", "Evaluate", unitRun{data, runUnitZipKeyAlignByEvaluateLibrary}, total + shortTotal + missing*7},
		{"ZipKey.WithDefaultValue", "Evaluate", unitRun{data, runUnitZipKeyWithDefaultValueEvaluateLibrary}, total + shortTotal + missing*9},
		{"ZipKey.WithDefaultKey", "Evaluate", unitRun{data, runUnitZipKeyWithDefaultKeyEvaluateLibrary}, total + shortTotal + missing*9},
	}
}

func zipIndex(left, right []int) int64 {
	var sum int64
	for i, k := range left {
		sum += int64(k) + int64(right[i])
	}
	return sum
}

func zipLeft(left, right []int, fallback int) int64 {
	var sum int64
	for i, k := range left {
		v := fallback
		if i < len(right) {
			v = right[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipRight(left, right []int, fallback int) int64 {
	var sum int64
	for i, v := range right {
		k := fallback
		if i < len(left) {
			k = left[i]
		}
		sum += int64(k) + int64(v)
	}
	return sum
}

func zipUnitReferences(values []int) map[string]unitReference {
	short := values[:len(values)/2]
	data := &unitInput{
		values:      values,
		shortValues: short,
	}
	return map[string]unitReference{
		"Zip.Operator":            {"ZipIndex", unitRun{data, runUnitZipOperatorEvaluateGo}},
		"Zip.AlignBy":             {"ZipLeft0", unitRun{data, runUnitZipAlignByEvaluateGo}},
		"Zip.WithDefaultValue":    {"ZipLeft7", unitRun{data, runUnitZipWithDefaultValueEvaluateGo}},
		"Zip.WithDefaultKey":      {"ZipRight7", unitRun{data, runUnitZipWithDefaultKeyEvaluateGo}},
		"ZipKey.Operator":         {"ZipIndex", unitRun{data, runUnitZipKeyOperatorEvaluateGo}},
		"ZipKey.AlignBy":          {"ZipRight7", unitRun{data, runUnitZipKeyAlignByEvaluateGo}},
		"ZipKey.WithDefaultValue": {"ZipLeft9", unitRun{data, runUnitZipKeyWithDefaultValueEvaluateGo}},
		"ZipKey.WithDefaultKey":   {"ZipRight9", unitRun{data, runUnitZipKeyWithDefaultKeyEvaluateGo}},
	}
}

func runUnitZipOperatorEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.source).Operator(data.source))
}

func benchmarkUnitZipOperatorEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipOperatorEvaluateLibrary)
}

func TestUnitZipOperatorEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Zip.Operator", "Evaluate", "Library", runUnitZipOperatorEvaluateLibrary)
}

func runUnitZipOperatorEvaluateGo(data *unitInput) int64 { return zipIndex(data.values, data.values) }

func benchmarkUnitZipOperatorEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipOperatorEvaluateGo)
}

func TestUnitZipOperatorEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Zip.Operator", "Evaluate", "Go", runUnitZipOperatorEvaluateGo)
}

func runUnitZipAlignByEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.short).AlignBy(seq.AlignLeft).Operator(data.source))
}

func benchmarkUnitZipAlignByEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipAlignByEvaluateLibrary)
}

func TestUnitZipAlignByEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Zip.AlignBy", "Evaluate", "Library", runUnitZipAlignByEvaluateLibrary)
}

func runUnitZipAlignByEvaluateGo(data *unitInput) int64 {
	return zipLeft(data.values, data.shortValues, 0)
}

func benchmarkUnitZipAlignByEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipAlignByEvaluateGo)
}

func TestUnitZipAlignByEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Zip.AlignBy", "Evaluate", "Go", runUnitZipAlignByEvaluateGo)
}

func runUnitZipWithDefaultValueEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.short).AlignBy(seq.AlignLeft).WithDefaultValue(7).Operator(data.source))
}

func benchmarkUnitZipWithDefaultValueEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipWithDefaultValueEvaluateLibrary)
}

func TestUnitZipWithDefaultValueEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Zip.WithDefaultValue", "Evaluate", "Library", runUnitZipWithDefaultValueEvaluateLibrary)
}

func runUnitZipWithDefaultValueEvaluateGo(data *unitInput) int64 {
	return zipLeft(data.values, data.shortValues, 7)
}

func benchmarkUnitZipWithDefaultValueEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipWithDefaultValueEvaluateGo)
}

func TestUnitZipWithDefaultValueEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Zip.WithDefaultValue", "Evaluate", "Go", runUnitZipWithDefaultValueEvaluateGo)
}

func runUnitZipWithDefaultKeyEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.source).AlignBy(seq.AlignRight).WithDefaultKey(7).Operator(data.short))
}

func benchmarkUnitZipWithDefaultKeyEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipWithDefaultKeyEvaluateLibrary)
}

func TestUnitZipWithDefaultKeyEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Zip.WithDefaultKey", "Evaluate", "Library", runUnitZipWithDefaultKeyEvaluateLibrary)
}

func runUnitZipWithDefaultKeyEvaluateGo(data *unitInput) int64 {
	return zipRight(data.shortValues, data.values, 7)
}

func benchmarkUnitZipWithDefaultKeyEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipWithDefaultKeyEvaluateGo)
}

func TestUnitZipWithDefaultKeyEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Zip.WithDefaultKey", "Evaluate", "Go", runUnitZipWithDefaultKeyEvaluateGo)
}

func runUnitZipKeyOperatorEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.source).WithDefaultKey(7).Operator(data.source))
}

func benchmarkUnitZipKeyOperatorEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyOperatorEvaluateLibrary)
}

func TestUnitZipKeyOperatorEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ZipKey.Operator", "Evaluate", "Library", runUnitZipKeyOperatorEvaluateLibrary)
}

func runUnitZipKeyOperatorEvaluateGo(data *unitInput) int64 {
	return zipIndex(data.values, data.values)
}

func benchmarkUnitZipKeyOperatorEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyOperatorEvaluateGo)
}

func TestUnitZipKeyOperatorEvaluateGo(t *testing.T) {
	testUnitOperation(t, "ZipKey.Operator", "Evaluate", "Go", runUnitZipKeyOperatorEvaluateGo)
}

func runUnitZipKeyAlignByEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.source).WithDefaultKey(7).AlignBy(seq.AlignRight).Operator(data.short))
}

func benchmarkUnitZipKeyAlignByEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyAlignByEvaluateLibrary)
}

func TestUnitZipKeyAlignByEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ZipKey.AlignBy", "Evaluate", "Library", runUnitZipKeyAlignByEvaluateLibrary)
}

func runUnitZipKeyAlignByEvaluateGo(data *unitInput) int64 {
	return zipRight(data.shortValues, data.values, 7)
}

func benchmarkUnitZipKeyAlignByEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyAlignByEvaluateGo)
}

func TestUnitZipKeyAlignByEvaluateGo(t *testing.T) {
	testUnitOperation(t, "ZipKey.AlignBy", "Evaluate", "Go", runUnitZipKeyAlignByEvaluateGo)
}

func runUnitZipKeyWithDefaultValueEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.short).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultValue(9).Operator(data.source))
}

func benchmarkUnitZipKeyWithDefaultValueEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyWithDefaultValueEvaluateLibrary)
}

func TestUnitZipKeyWithDefaultValueEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ZipKey.WithDefaultValue", "Evaluate", "Library", runUnitZipKeyWithDefaultValueEvaluateLibrary)
}

func runUnitZipKeyWithDefaultValueEvaluateGo(data *unitInput) int64 {
	return zipLeft(data.values, data.shortValues, 9)
}

func benchmarkUnitZipKeyWithDefaultValueEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyWithDefaultValueEvaluateGo)
}

func TestUnitZipKeyWithDefaultValueEvaluateGo(t *testing.T) {
	testUnitOperation(t, "ZipKey.WithDefaultValue", "Evaluate", "Go", runUnitZipKeyWithDefaultValueEvaluateGo)
}

func runUnitZipKeyWithDefaultKeyEvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Zip(data.source).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultKey(9).Operator(data.short))
}

func benchmarkUnitZipKeyWithDefaultKeyEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyWithDefaultKeyEvaluateLibrary)
}

func TestUnitZipKeyWithDefaultKeyEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ZipKey.WithDefaultKey", "Evaluate", "Library", runUnitZipKeyWithDefaultKeyEvaluateLibrary)
}

func runUnitZipKeyWithDefaultKeyEvaluateGo(data *unitInput) int64 {
	return zipRight(data.shortValues, data.values, 9)
}

func benchmarkUnitZipKeyWithDefaultKeyEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitZipKeyWithDefaultKeyEvaluateGo)
}

func TestUnitZipKeyWithDefaultKeyEvaluateGo(t *testing.T) {
	testUnitOperation(t, "ZipKey.WithDefaultKey", "Evaluate", "Go", runUnitZipKeyWithDefaultKeyEvaluateGo)
}
