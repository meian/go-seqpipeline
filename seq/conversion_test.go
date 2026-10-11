package seq_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestNilConversions(t *testing.T) {
	var zero seq.Seq[int]
	var zero2 seq.Seq2[int, string]
	if zero != nil || zero2 != nil {
		t.Fatal("zero values must be nil")
	}
	for _, s := range []seq.Seq[int]{zero, seq.Seq[int](nil), seq.FromSeq[int](nil)} {
		if seq.ToSeq(s) != nil {
			t.Fatal("ToSeq must preserve nil")
		}
	}
	for _, s := range []seq.Seq2[int, string]{zero2, seq.Seq2[int, string](nil), seq.FromSeq2[int, string](nil)} {
		if seq.ToSeq2(s) != nil {
			t.Fatal("ToSeq2 must preserve nil")
		}
	}
}

func TestStandardIteratorConversions(t *testing.T) {
	s := seq.FromSlice([]int{3, 1, 2})
	assertEqual(t, slices.Sorted(seq.ToSeq(s)), []int{1, 2, 3})
	pairs := seq.FromMap(map[string]int{"a": 1, "b": 2})
	assertEqual(t, maps.Collect(seq.ToSeq2(pairs)), map[string]int{"a": 1, "b": 2})
}

func TestToSeq2PreservesPairs(t *testing.T) {
	src := func(yield func(string, int) bool) {
		if !yield("a", 1) {
			return
		}
		yield("a", 2)
	}
	pairs := seq.FromSeq2(src)
	want := []pair[string, int]{{Key: "a", Value: 1}, {Key: "a", Value: 2}}
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pairs)), want)
	}
}

func TestToMap(t *testing.T) {
	reads := 0
	src := func(yield func(string, int) bool) {
		for _, p := range []pair[string, int]{{Key: "a", Value: 1}, {Key: "b", Value: 2}, {Key: "a", Value: 3}} {
			reads++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
	source := seq.FromSeq2(src)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		var got map[string]int = seq.ToMap(source)
		assertEqual(t, got, map[string]int{"a": 3, "b": 2})
		assertEqual(t, reads, round*3)
		got["a"] = 99
	}
}

func TestToMapEmpty(t *testing.T) {
	for _, src := range []seq.Seq2[string, int]{seq.FromMap(map[string]int{}), seq.FromMap(map[string]int(nil)), seq.FromSeq2(func(func(string, int) bool) {})} {
		got := seq.ToMap(src)
		if got == nil {
			t.Fatal("ToMap returned nil")
		}
		assertEqual(t, len(got), 0)
		got["a"] = 1
		assertEqual(t, seq.ToMap(src), map[string]int{})
	}
}

func TestToSlice(t *testing.T) {
	values := []int{3, 1, 2}
	source := seq.FromSlice(values)
	got := seq.ToSlice(source)
	assertEqual(t, got, values)
	got[0] = 99
	assertEqual(t, seq.ToSlice(source), []int{3, 1, 2})
}

func TestToSeqLazyAndStop(t *testing.T) {
	starts, reads, stops := 0, 0, 0
	source := seq.FromSeq(func(yield func(int) bool) {
		starts++
		for _, v := range []int{3, 1, 2} {
			reads++
			if !yield(v) {
				stops++
				return
			}
		}
	})
	iterator := seq.ToSeq(source)
	assertEqual(t, []int{starts, reads, stops}, []int{0, 0, 0})
	for round := 1; round <= 2; round++ {
		for v := range iterator {
			assertEqual(t, v, 3)
			break
		}
		assertEqual(t, []int{starts, reads, stops}, []int{round, round, round})
	}
}

func TestNilCollections(t *testing.T) {
	for _, tc := range []struct {
		name    string
		collect func()
	}{
		{name: "ToSlice", collect: func() { _ = seq.ToSlice[int](nil) }},
		{name: "slices.Collect", collect: func() { _ = slices.Collect[int](nil) }},
		{name: "ToMap", collect: func() { _ = seq.ToMap[string, int](nil) }},
		{name: "maps.Collect", collect: func() { _ = maps.Collect[string, int](nil) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("nil iterator must panic")
				}
			}()
			tc.collect()
		})
	}
}

func TestToSliceEmpty(t *testing.T) {
	for _, src := range []seq.Seq[int]{seq.FromSlice([]int{}), seq.FromSlice([]int(nil)), seq.FromSeq(func(func(int) bool) {})} {
		assertEqual(t, seq.ToSlice(src), []int(nil))
	}
}
