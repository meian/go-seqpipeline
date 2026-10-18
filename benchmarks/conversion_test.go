package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedValues []int

var collectedMap map[int]int

func conversionUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	standard := slices.Values(values)
	standardPairs := slices.All(values)
	mapping := make(map[int]int, len(values))
	var total, keys int64
	for k, v := range values {
		mapping[k] = v
		total += int64(v)
		keys += int64(k)
	}
	return []unitCase{
		{"FromSlice", "Evaluate", func() int64 { return sumSeq(seq.FromSlice(values)) }, total},
		{"FromSeq", "Evaluate", func() int64 { return sumSeq(seq.FromSeq(standard)) }, total},
		{"ToSeq", "Evaluate", func() int64 {
			var sum int64
			for v := range seq.ToSeq(source) {
				sum += int64(v)
			}
			return sum
		}, total},
		{"ToSlice", "Collect", func() int64 {
			collectedValues = seq.ToSlice(source)
			var sum int64
			for _, v := range collectedValues {
				sum += int64(v)
			}
			return sum
		}, total},
		{"FromMap", "Evaluate", func() int64 { return sumSeq2(seq.FromMap(mapping)) }, keys + total},
		{"FromSeq2", "Evaluate", func() int64 { return sumSeq2(seq.FromSeq2(standardPairs)) }, keys + total},
		{"ToSeq2", "Evaluate", func() int64 {
			var sum int64
			for k, v := range seq.ToSeq2(pairs) {
				sum += int64(k) + int64(v)
			}
			return sum
		}, keys + total},
		{"ToMap", "Collect", func() int64 {
			collectedMap = seq.ToMap(pairs)
			var sum int64
			for k, v := range collectedMap {
				sum += int64(k) + int64(v)
			}
			return sum
		}, keys + total},
	}
}

func TestConversionUnitResults(t *testing.T) { checkUnitCases(t, conversionUnitCases) }
