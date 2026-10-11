package seq_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestDefinedFunctionTypes(t *testing.T) {
	src := slices.Values([]string{"1", "2"})
	pipeline := seq.Seq[string](src)
	var got []string
	for v := range pipeline {
		got = append(got, v)
	}
	assertEqual(t, got, []string{"1", "2"})
	got = nil
	pipeline(func(v string) bool {
		got = append(got, v)
		return false
	})
	assertEqual(t, got, []string{"1"})
	assertEqual(t, slices.Collect(iter.Seq[string](pipeline)), []string{"1", "2"})

	src2 := func(yield func(string, int) bool) {
		if !yield("0", 1) {
			return
		}
		yield("1", 3)
	}
	pipeline2 := seq.Seq2[string, int](src2)
	var pairs []pair[string, int]
	for k, v := range pipeline2 {
		pairs = append(pairs, pair[string, int]{Key: k, Value: v})
	}
	assertEqual(t, pairs, []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}})
	pairs = nil
	pipeline2(func(k string, v int) bool {
		pairs = append(pairs, pair[string, int]{Key: k, Value: v})
		return false
	})
	assertEqual(t, pairs, []pair[string, int]{{Key: "0", Value: 1}})
	assertEqual(t, collectPairs(seq.ToSeq2(seq.FromSeq2(iter.Seq2[string, int](pipeline2)))), collectPairs(seq.ToSeq2(pipeline2)))
}
