package benchmarks

import (
	"maps"
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

func TestConversionUnitResults(t *testing.T) {
	checkUnitCases(t, conversionUnitCases, conversionUnitReferences)
}

func conversionUnitReferences(values []int) map[string]unitReference {
	standard := slices.Values(values)
	standardPairs := slices.All(values)
	mapping := make(map[int]int, len(values))
	for k, v := range values {
		mapping[k] = v
	}
	iterate := func() int64 {
		var sum int64
		for v := range standard {
			sum += int64(v)
		}
		return sum
	}
	iteratePairs := func() int64 {
		var sum int64
		for k, v := range standardPairs {
			sum += int64(k) + int64(v)
		}
		return sum
	}
	return map[string]unitReference{
		"FromSlice": {"SliceRange", func() int64 { return rangeValues(values) }},
		"FromSeq":   {"IteratorRange", iterate},
		"ToSeq":     {"IteratorRange", iterate},
		"ToSlice": {"SlicesCollect", func() int64 {
			collectedValues = slices.Collect(standard)
			return rangeValues(collectedValues)
		}},
		"FromMap": {"MapRange", func() int64 {
			var sum int64
			for k, v := range mapping {
				sum += int64(k) + int64(v)
			}
			return sum
		}},
		"FromSeq2": {"Iterator2Range", iteratePairs},
		"ToSeq2":   {"Iterator2Range", iteratePairs},
		"ToMap": {"MapsCollect", func() int64 {
			collectedMap = maps.Collect(standardPairs)
			var sum int64
			for k, v := range collectedMap {
				sum += int64(k) + int64(v)
			}
			return sum
		}},
	}
}
