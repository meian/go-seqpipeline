package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func pipeUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var mapped, total, keys int64
	for k, v := range values {
		total += int64(v)
		keys += int64(k)
		mapped += int64(mapValue(v))
	}
	return []unitCase{
		{"Seq.Pipe", "Evaluate", func() int64 { return sumSeq(source.Pipe(seq.Map(mapValue))) }, mapped},
		{"Seq.Pipe2", "Evaluate", func() int64 { return sumSeq2(source.Pipe2(seq.Zip(source).Operator)) }, total * 2},
		{"Seq2.Pipe", "Evaluate", func() int64 { return sumSeq(pairs.Pipe(seq.Values)) }, total},
		{"Seq2.Pipe2", "Evaluate", func() int64 {
			return sumSeq2(pairs.Pipe2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })))
		}, keys + mapped},
	}
}

func TestPipeUnitResults(t *testing.T) { checkUnitCases(t, pipeUnitCases, pipeUnitReferences) }

func pipeUnitReferences(values []int) map[string]unitReference {
	transforms := transformUnitReferences(values)
	return map[string]unitReference{
		"Seq.Pipe":   transforms["Map"],
		"Seq.Pipe2":  {"ZipIndex", func() int64 { return zipIndex(values, values) }},
		"Seq2.Pipe":  transforms["Values"],
		"Seq2.Pipe2": transforms["Map2"],
	}
}
