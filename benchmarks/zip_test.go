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
	return []unitCase{
		{"Zip.Operator", "Evaluate", func() int64 { return sumSeq2(seq.Zip(source).Operator(source)) }, total * 2},
		{"Zip.AlignBy", "Evaluate", func() int64 { return sumSeq2(seq.Zip(short).AlignBy(seq.AlignLeft).Operator(source)) }, total + shortTotal},
		{"Zip.WithDefaultValue", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(short).AlignBy(seq.AlignLeft).WithDefaultValue(7).Operator(source))
		}, total + shortTotal + missing*7},
		{"Zip.WithDefaultKey", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(source).AlignBy(seq.AlignRight).WithDefaultKey(7).Operator(short))
		}, total + shortTotal + missing*7},
		{"ZipKey.Operator", "Evaluate", func() int64 { return sumSeq2(seq.Zip(source).WithDefaultKey(7).Operator(source)) }, total * 2},
		{"ZipKey.AlignBy", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(source).WithDefaultKey(7).AlignBy(seq.AlignRight).Operator(short))
		}, total + shortTotal + missing*7},
		{"ZipKey.WithDefaultValue", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(short).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultValue(9).Operator(source))
		}, total + shortTotal + missing*9},
		{"ZipKey.WithDefaultKey", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(source).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultKey(9).Operator(short))
		}, total + shortTotal + missing*9},
	}
}

func TestZipUnitResults(t *testing.T) { checkUnitCases(t, zipUnitCases, zipUnitReferences) }

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
	return map[string]unitReference{
		"Zip.Operator":            {"ZipIndex", func() int64 { return zipIndex(values, values) }},
		"Zip.AlignBy":             {"ZipLeft0", func() int64 { return zipLeft(values, short, 0) }},
		"Zip.WithDefaultValue":    {"ZipLeft7", func() int64 { return zipLeft(values, short, 7) }},
		"Zip.WithDefaultKey":      {"ZipRight7", func() int64 { return zipRight(short, values, 7) }},
		"ZipKey.Operator":         {"ZipIndex", func() int64 { return zipIndex(values, values) }},
		"ZipKey.AlignBy":          {"ZipRight7", func() int64 { return zipRight(short, values, 7) }},
		"ZipKey.WithDefaultValue": {"ZipLeft9", func() int64 { return zipLeft(values, short, 9) }},
		"ZipKey.WithDefaultKey":   {"ZipRight9", func() int64 { return zipRight(short, values, 9) }},
	}
}
