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

func TestZipUnitResults(t *testing.T) { checkUnitCases(t, zipUnitCases) }
