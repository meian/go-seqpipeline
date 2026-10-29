package seq_test

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFromSlice(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 2}
	source := seq.FromSlice(values)
	values[0] = 4
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(source), []int{4, 1, 2})
	}
	assertEqual(t, seq.ToSlice(seq.FromSlice([]int(nil))), []int(nil))
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	values := []int{3, 1, 2}
	source := seq.FromSlice(values)
	got := seq.ToSlice(source)
	assertEqual(t, got, values)
	got[0] = 99
	assertEqual(t, seq.ToSlice(source), []int{3, 1, 2})
}

func TestToSliceEmpty(t *testing.T) {
	t.Parallel()

	for _, src := range []seq.Seq[int]{
		seq.FromSlice([]int{}),
		seq.FromSlice([]int(nil)),
		seq.FromSeq(func(func(int) bool) {}),
	} {
		assertEqual(t, seq.ToSlice(src), []int(nil))
	}
}

func TestToSliceNil(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatal("ToSlice(nil) must panic")
		}
	}()
	seq.ToSlice[int](nil)
}

func TestFromSeq(t *testing.T) {
	t.Parallel()

	source := seq.FromSeq(slices.Values([]int{3, 1, 2}))
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(source), []int{3, 1, 2})
	}
}

func TestFromSeqNil(t *testing.T) {
	t.Parallel()

	if seq.FromSeq[int](nil) != nil {
		t.Fatal("FromSeq must preserve nil")
	}
}

func TestToSeqNil(t *testing.T) {
	t.Parallel()

	var zero seq.Seq[int]
	if seq.ToSeq(zero) != nil {
		t.Fatal("ToSeq must preserve nil")
	}
}

func TestToSeq(t *testing.T) {
	t.Parallel()

	starts, reads, stops := 0, 0, 0
	source := seq.FromSeq(func(yield func(int) bool) {
		starts++
		for _, v := range []int{3, 1, 3} {
			reads++
			if !yield(v) {
				stops++
				return
			}
		}
	})
	iterator := seq.ToSeq(source)
	assertEqual(t, []int{starts, reads, stops}, []int{0, 0, 0})
	for i := 1; i <= 2; i++ {
		for v := range iterator {
			assertEqual(t, v, 3)
			break
		}
		assertEqual(t, []int{starts, reads, stops}, []int{i, i, i})
	}
	assertEqual(t, slices.Collect(iterator), []int{3, 1, 3})
	assertEqual(t, []int{starts, reads, stops}, []int{3, 5, 2})
}

func TestFromMap(t *testing.T) {
	t.Parallel()

	values := map[string]int{"a": 1, "b": 2}
	pipeline := seq.FromMap(values)
	values["c"] = 3
	for i := 0; i < 2; i++ {
		got := make(map[string]int)
		pairs := collectPairs(seq.ToSeq2(pipeline))
		for _, p := range pairs {
			got[p.Key] = p.Value
		}
		assertEqual(t, len(pairs), len(values))
		assertEqual(t, got, values)
	}
}

func TestToMap(t *testing.T) {
	t.Parallel()

	reads := 0
	src := func(yield func(string, int) bool) {
		for _, p := range []pair[string, int]{
			{Key: "a", Value: 1},
			{Key: "b", Value: 2},
			{Key: "a", Value: 3},
		} {
			reads++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
	source := seq.FromSeq2(src)
	assertEqual(t, reads, 0)
	for i := 1; i <= 2; i++ {
		got := seq.ToMap(source)
		assertEqual(t, got, map[string]int{"a": 3, "b": 2})
		assertEqual(t, reads, i*3)
		got["a"] = 99
	}
}

func TestToMapEmpty(t *testing.T) {
	t.Parallel()

	for _, src := range []seq.Seq2[string, int]{
		seq.FromMap(map[string]int{}),
		seq.FromMap(map[string]int(nil)),
		seq.FromSeq2(func(func(string, int) bool) {}),
	} {
		got := seq.ToMap(src)
		if got == nil {
			t.Fatal("ToMap returned nil")
		}
		assertEqual(t, len(got), 0)
		got["a"] = 1
		assertEqual(t, seq.ToMap(src), map[string]int{})
	}
}

func TestToMapNil(t *testing.T) {
	t.Parallel()

	defer func() {
		if recover() == nil {
			t.Fatal("ToMap(nil) must panic")
		}
	}()
	seq.ToMap[string, int](nil)
}

func TestFromSeq2(t *testing.T) {
	t.Parallel()

	source := seq.FromSeq2(slices.All([]string{"b", "a"}))
	want := []pair[int, string]{{Key: 0, Value: "b"}, {Key: 1, Value: "a"}}
	for i := 0; i < 2; i++ {
		assertEqual(t, collectPairs(seq.ToSeq2(source)), want)
	}
}

func TestFromSeq2Nil(t *testing.T) {
	t.Parallel()

	if seq.FromSeq2[int, string](nil) != nil {
		t.Fatal("FromSeq2 must preserve nil")
	}
}

func TestToSeq2Nil(t *testing.T) {
	t.Parallel()

	var zero seq.Seq2[int, string]
	if seq.ToSeq2(zero) != nil {
		t.Fatal("ToSeq2 must preserve nil")
	}
}

func TestToSeq2(t *testing.T) {
	t.Parallel()

	src := func(yield func(string, int) bool) {
		if !yield("a", 1) {
			return
		}
		yield("a", 2)
	}
	pairs := seq.FromSeq2(src)
	want := []pair[string, int]{{Key: "a", Value: 1}, {Key: "a", Value: 2}}
	for i := 0; i < 2; i++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pairs)), want)
	}
}
