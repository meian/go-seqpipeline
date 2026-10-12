package seq_test

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFromSlice(t *testing.T) {
	values := []int{3, 1, 2}
	source := seq.FromSlice(values)
	values[0] = 4
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(source), []int{4, 1, 2})
	}
	assertEqual(t, seq.ToSlice(seq.FromSlice([]int(nil))), []int(nil))
}

func TestFromSeq(t *testing.T) {
	source := seq.FromSeq(slices.Values([]int{3, 1, 2}))
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(source), []int{3, 1, 2})
	}
}

func TestFromMap(t *testing.T) {
	values := map[string]int{"a": 1, "b": 2}
	pipeline := seq.FromMap(values)
	values["c"] = 3
	for round := 0; round < 2; round++ {
		got := make(map[string]int)
		pairs := collectPairs(seq.ToSeq2(pipeline))
		for _, p := range pairs {
			got[p.Key] = p.Value
		}
		assertEqual(t, len(pairs), len(values))
		assertEqual(t, got, values)
	}
}

func TestFromSeq2(t *testing.T) {
	source := seq.FromSeq2(slices.All([]string{"b", "a"}))
	want := []pair[int, string]{{Key: 0, Value: "b"}, {Key: 1, Value: "a"}}
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(source)), want)
	}
}
