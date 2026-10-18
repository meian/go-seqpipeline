package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func limitUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var taken, skipped, takenKeys, skippedKeys int64
	for k, v := range values {
		if k < 8 {
			taken += int64(v)
			takenKeys += int64(k)
		}
		if k >= 8 {
			skipped += int64(v)
			skippedKeys += int64(k)
		}
	}
	return []unitCase{
		{"Take.Operator", "Early8", func() int64 { return sumSeq(seq.Take(8).Operator(source)) }, taken},
		{"Skip.Operator", "Evaluate", func() int64 { return sumSeq(seq.Skip(8).Operator(source)) }, skipped},
		{"Take.Operator2", "Early8", func() int64 { return sumSeq2(seq.Take(8).Operator2(pairs)) }, taken + takenKeys},
		{"Skip.Operator2", "Evaluate", func() int64 { return sumSeq2(seq.Skip(8).Operator2(pairs)) }, skipped + skippedKeys},
	}
}

func TestLimitUnitResults(t *testing.T) { checkUnitCases(t, limitUnitCases, limitUnitReferences) }

func limitUnitReferences(values []int) map[string]unitReference {
	end := min(8, len(values))
	return map[string]unitReference{
		"Take.Operator": {"TakeSlice", func() int64 { return rangeValues(values[:end]) }},
		"Skip.Operator": {"SkipSlice", func() int64 { return rangeValues(values[end:]) }},
		"Take.Operator2": {"Take2Slice", func() int64 {
			var sum int64
			for k, v := range values[:end] {
				sum += int64(k) + int64(v)
			}
			return sum
		}},
		"Skip.Operator2": {"Skip2Slice", func() int64 {
			var sum int64
			for k, v := range values[end:] {
				sum += int64(k+end) + int64(v)
			}
			return sum
		}},
	}
}
