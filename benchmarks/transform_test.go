package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func transformUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var total, mapped, filtered, keys int64
	for k, v := range values {
		total += int64(v)
		keys += int64(k)
		mapped += int64(mapValue(v))
		if keepValue(v) {
			filtered += int64(v)
		}
	}
	return []unitCase{
		{"Map", "Evaluate", func() int64 { return sumSeq(seq.Map(mapValue)(source)) }, mapped},
		{"Filter", "Evaluate", func() int64 { return sumSeq(seq.Filter(keepValue)(source)) }, filtered},
		{"Map2", "Evaluate", func() int64 { return sumSeq2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })(pairs)) }, keys + mapped},
		{"Filter2", "Evaluate", func() int64 { return sumSeq2(seq.Filter2(func(k, v int) bool { return keepValue(v) })(pairs)) }, filtered * 2},
		{"Keys", "Evaluate", func() int64 { return sumSeq(seq.Keys(pairs)) }, keys},
		{"Values", "Evaluate", func() int64 { return sumSeq(seq.Values(pairs)) }, total},
	}
}

func TestTransformUnitResults(t *testing.T) { checkUnitCases(t, transformUnitCases) }
