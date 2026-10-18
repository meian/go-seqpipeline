package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func limitUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var taken, skipped int64
	for k, v := range values {
		if k < 8 {
			taken += int64(v)
		}
		if k >= 8 {
			skipped += int64(v)
		}
	}
	return []unitCase{
		{"Take.Operator", "Early8", func() int64 { return sumSeq(seq.Take(8).Operator(source)) }, taken},
		{"Skip.Operator", "Evaluate", func() int64 { return sumSeq(seq.Skip(8).Operator(source)) }, skipped},
		{"Take.Operator2", "Early8", func() int64 { return sumSeq2(seq.Take(8).Operator2(pairs)) }, taken * 2},
		{"Skip.Operator2", "Evaluate", func() int64 { return sumSeq2(seq.Skip(8).Operator2(pairs)) }, skipped * 2},
	}
}

func TestLimitUnitResults(t *testing.T) { checkUnitCases(t, limitUnitCases) }
