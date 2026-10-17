package seq_test

import (
	"iter"
	"slices"
	"sync"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestCache(t *testing.T) {
	values := []int{3, 1, 1, 2}
	starts, reads, finishes := 0, 0, 0
	src := seq.Seq[int](func(yield func(int) bool) {
		starts++
		if starts > 1 {
			t.Fatal("single-use source restarted")
		}
		defer func() { finishes++ }()
		for _, v := range values {
			reads++
			if !yield(v) {
				return
			}
		}
	})
	cached := seq.Cache(src)
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
	values[0] = 99
	yields := 0
	cached(func(v int) bool {
		yields++
		assertEqual(t, v, 3)
		return false
	})
	assertEqual(t, yields, 1)
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(cached), []int{3, 1, 1, 2})
	}
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
}

func TestCache2(t *testing.T) {
	values := []pair[int, string]{
		{Key: 3, Value: "c"}, {Key: 1, Value: "a"},
		{Key: 1, Value: "a"}, {Key: 1, Value: "b"},
	}
	starts, reads, finishes := 0, 0, 0
	src := seq.Seq2[int, string](func(yield func(int, string) bool) {
		starts++
		if starts > 1 {
			t.Fatal("single-use source restarted")
		}
		defer func() { finishes++ }()
		for _, p := range values {
			reads++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	cached := seq.Cache2(src)
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
	want := slices.Clone(values)
	values[0] = pair[int, string]{Key: 99, Value: "changed"}
	yields := 0
	cached(func(k int, v string) bool {
		yields++
		assertEqual(t, pair[int, string]{Key: k, Value: v}, want[0])
		return false
	})
	assertEqual(t, yields, 1)
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(cached)), want)
	}
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
}

func TestCacheEmptyAndNil(t *testing.T) {
	for _, src := range []seq.Seq[int]{nil, func(func(int) bool) {}} {
		cached := seq.Cache(src)
		if cached == nil {
			t.Fatal("Cache returned nil")
		}
		for round := 0; round < 2; round++ {
			assertEqual(t, seq.ToSlice(cached), []int(nil))
		}
	}
}

func TestCache2EmptyAndNil(t *testing.T) {
	for _, src := range []seq.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
		cached := seq.Cache2(src)
		if cached == nil {
			t.Fatal("Cache2 returned nil")
		}
		for round := 0; round < 2; round++ {
			assertEqual(t, collectPairs(seq.ToSeq2(cached)), []pair[int, string](nil))
		}
	}
}

func TestCacheShallowCopy(t *testing.T) {
	value := 1
	values := []*int{&value}
	cached := seq.Cache(seq.FromSlice(values))
	other := 2
	values[0] = &other
	value = 3
	assertEqual(t, seq.ToSlice(cached), []*int{&value})
	assertEqual(t, *seq.ToSlice(cached)[0], 3)
}

func TestCache2ShallowCopyAndNonComparableKeys(t *testing.T) {
	keys, values := [][]int{{1}}, [][]int{{10}}
	cached := seq.Cache2(seq.Seq2[[]int, []int](func(yield func([]int, []int) bool) {
		yield(keys[0], values[0])
	}))
	key, value := keys[0], values[0]
	keys[0], values[0] = []int{2}, []int{20}
	key[0], value[0] = 3, 30
	want := []pair[[]int, []int]{{Key: []int{3}, Value: []int{30}}}
	assertEqual(t, collectPairs(seq.ToSeq2(cached)), want)
}

func TestCacheInterleavedAndConcurrent(t *testing.T) {
	cached := seq.Cache(seq.FromSlice([]int{3, 1, 1, 2}))
	next, stop := iter.Pull(seq.ToSeq(cached))
	defer stop()
	v, ok := next()
	assertEqual(t, v, 3)
	assertEqual(t, ok, true)
	assertEqual(t, seq.ToSlice(cached), []int{3, 1, 1, 2})
	v, ok = next()
	assertEqual(t, v, 1)
	assertEqual(t, ok, true)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Go(func() { assertEqual(t, seq.ToSlice(cached), []int{3, 1, 1, 2}) })
	}
	wg.Wait()
}

func TestCache2InterleavedAndConcurrent(t *testing.T) {
	cached := seq.Cache2(seq.FromSeq2(slices.All([]string{"a", "b"})))
	want := []pair[int, string]{{Key: 0, Value: "a"}, {Key: 1, Value: "b"}}
	next, stop := iter.Pull2(seq.ToSeq2(cached))
	defer stop()
	k, v, ok := next()
	assertEqual(t, pair[int, string]{Key: k, Value: v}, want[0])
	assertEqual(t, ok, true)
	assertEqual(t, collectPairs(seq.ToSeq2(cached)), want)
	k, v, ok = next()
	assertEqual(t, pair[int, string]{Key: k, Value: v}, want[1])
	assertEqual(t, ok, true)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Go(func() { assertEqual(t, collectPairs(seq.ToSeq2(cached)), want) })
	}
	wg.Wait()
}
