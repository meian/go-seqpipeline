package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func transformUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var total, mapped, filtered, filteredKeys, keys int64
	for k, v := range values {
		total += int64(v)
		keys += int64(k)
		mapped += int64(mapValue(v))
		if keepValue(v) {
			filtered += int64(v)
			filteredKeys += int64(k)
		}
	}
	return []unitCase{
		{"Map", "Evaluate", func() int64 { return sumSeq(seq.Map(mapValue)(source)) }, mapped},
		{"Filter", "Evaluate", func() int64 { return sumSeq(seq.Filter(keepValue)(source)) }, filtered},
		{"Map2", "Evaluate", func() int64 { return sumSeq2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })(pairs)) }, keys + mapped},
		{"Filter2", "Evaluate", func() int64 { return sumSeq2(seq.Filter2(func(k, v int) bool { return keepValue(v) })(pairs)) }, filtered + filteredKeys},
		{"Keys", "Evaluate", func() int64 { return sumSeq(seq.Keys(pairs)) }, keys},
		{"Values", "Evaluate", func() int64 { return sumSeq(seq.Values(pairs)) }, total},
	}
}

func TestTransformUnitResults(t *testing.T) {
	checkUnitCases(t, transformUnitCases, transformUnitReferences)
}

func transformUnitReferences(values []int) map[string]unitReference {
	return map[string]unitReference{
		"Map": {"MapLoop", func() int64 {
			var sum int64
			for _, v := range values {
				sum += int64(mapValue(v))
			}
			return sum
		}},
		"Filter": {"FilterLoop", func() int64 {
			var sum int64
			for _, v := range values {
				if keepValue(v) {
					sum += int64(v)
				}
			}
			return sum
		}},
		"Map2": {"Map2Loop", func() int64 {
			var sum int64
			for k, v := range values {
				sum += int64(k) + int64(mapValue(v))
			}
			return sum
		}},
		"Filter2": {"Filter2Loop", func() int64 {
			var sum int64
			for k, v := range values {
				if keepValue(v) {
					sum += int64(k) + int64(v)
				}
			}
			return sum
		}},
		"Keys": {"KeyLoop", func() int64 {
			var sum int64
			for k := range values {
				sum += int64(k)
			}
			return sum
		}},
		"Values": {"SliceRange", func() int64 { return rangeValues(values) }},
	}
}
