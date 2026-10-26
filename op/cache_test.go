package op_test

import (
	"iter"
	"slices"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

func TestCache(t *testing.T) {
	t.Parallel()
	values := []int{3, 1, 1, 2}
	starts, reads, finishes := 0, 0, 0
	src := iter.Seq[int](func(yield func(int) bool) {
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
	cached := op.Cache(src)
	assertEqual(t, []int{starts, reads, finishes}, []int{0, 0, 0})
	values[0] = 99
	yields := 0
	cached(func(v int) bool {
		yields++
		assertEqual(t, v, 99)
		assertEqual(t, []int{reads, finishes}, []int{4, 1})
		return false
	})
	assertEqual(t, yields, 1)
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
	values[0] = 77
	for round := 0; round < 2; round++ {
		assertEqual(t, slices.Collect(cached), []int{99, 1, 1, 2})
	}
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
}

func TestCacheEmptyAndNil(t *testing.T) {
	t.Parallel()
	for _, src := range []iter.Seq[int]{nil, func(func(int) bool) {}} {
		cached := op.Cache(src)
		if cached == nil {
			t.Fatal("Cache returned nil")
		}
		for round := 0; round < 2; round++ {
			assertEqual(t, slices.Collect(cached), []int(nil))
		}
	}
}

func TestCacheShallowCopy(t *testing.T) {
	t.Parallel()
	value := 1
	values := []*int{&value}
	cached := op.Cache(slices.Values(values))
	assertEqual(t, slices.Collect(cached), []*int{&value})
	other := 2
	values[0] = &other
	value = 3
	assertEqual(t, slices.Collect(cached), []*int{&value})
	assertEqual(t, *slices.Collect(cached)[0], 3)
}

func TestCacheInterleavedAndConcurrent(t *testing.T) {
	t.Parallel()
	cached := op.Cache(slices.Values([]int{3, 1, 1, 2}))
	next, stop := iter.Pull(cached)
	defer stop()
	v, ok := next()
	assertEqual(t, v, 3)
	assertEqual(t, ok, true)
	assertEqual(t, slices.Collect(cached), []int{3, 1, 1, 2})
	v, ok = next()
	assertEqual(t, v, 1)
	assertEqual(t, ok, true)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Go(func() { assertEqual(t, slices.Collect(cached), []int{3, 1, 1, 2}) })
	}
	wg.Wait()
}

// 初回評価中の反復も完成した同一キャッシュを待ち、上流を重複実行しない。
func TestCacheConcurrentFirstIteration(t *testing.T) {
	t.Parallel()
	var starts atomic.Int32
	started, release := make(chan struct{}), make(chan struct{})
	cached := op.Cache(iter.Seq[int](func(yield func(int) bool) {
		starts.Add(1)
		close(started)
		<-release
		yield(1)
		yield(2)
	}))
	results := make(chan []int, 4)
	for i := 0; i < cap(results); i++ {
		go func() { results <- slices.Collect(cached) }()
	}
	<-started
	assertEqual(t, starts.Load(), int32(1))
	close(release)
	for i := 0; i < cap(results); i++ {
		assertEqual(t, <-results, []int{1, 2})
	}
	assertEqual(t, starts.Load(), int32(1))
}

func TestCachePanic(t *testing.T) {
	t.Parallel()
	starts := 0
	cached := op.Cache(iter.Seq[int](func(yield func(int) bool) {
		starts++
		yield(1)
		panic("cache source failed")
	}))
	for i := 0; i < 2; i++ {
		yields := 0
		func() {
			defer func() { assertEqual(t, recover(), any("cache source failed")) }()
			cached(func(int) bool { yields++; return true })
		}()
		assertEqual(t, yields, 0)
	}
	assertEqual(t, starts, 1)
}

func TestCache2(t *testing.T) {
	t.Parallel()
	values := []pair[int, string]{
		{Key: 3, Value: "c"}, {Key: 1, Value: "a"},
		{Key: 1, Value: "a"}, {Key: 1, Value: "b"},
	}
	starts, reads, finishes := 0, 0, 0
	src := iter.Seq2[int, string](func(yield func(int, string) bool) {
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
	cached := op.Cache2(src)
	assertEqual(t, []int{starts, reads, finishes}, []int{0, 0, 0})
	values[0] = pair[int, string]{Key: 99, Value: "changed"}
	want := slices.Clone(values)
	yields := 0
	cached(func(k int, v string) bool {
		yields++
		assertEqual(t, pair[int, string]{Key: k, Value: v}, want[0])
		assertEqual(t, []int{reads, finishes}, []int{4, 1})
		return false
	})
	assertEqual(t, yields, 1)
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
	values[0] = pair[int, string]{Key: 77, Value: "later"}
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(cached), want)
	}
	assertEqual(t, []int{starts, reads, finishes}, []int{1, 4, 1})
}

func TestCache2EmptyAndNil(t *testing.T) {
	t.Parallel()
	for _, src := range []iter.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
		cached := op.Cache2(src)
		if cached == nil {
			t.Fatal("Cache2 returned nil")
		}
		for round := 0; round < 2; round++ {
			assertEqual(t, collectPairs(cached), []pair[int, string](nil))
		}
	}
}

func TestCache2ShallowCopyAndNonComparableKeys(t *testing.T) {
	t.Parallel()
	keys, values := [][]int{{1}}, [][]int{{10}}
	cached := op.Cache2(iter.Seq2[[]int, []int](func(yield func([]int, []int) bool) {
		yield(keys[0], values[0])
	}))
	key, value := keys[0], values[0]
	assertEqual(t, collectPairs(cached), []pair[[]int, []int]{{Key: []int{1}, Value: []int{10}}})
	keys[0], values[0] = []int{2}, []int{20}
	key[0], value[0] = 3, 30
	want := []pair[[]int, []int]{{Key: []int{3}, Value: []int{30}}}
	assertEqual(t, collectPairs(cached), want)
}

func TestCache2InterleavedAndConcurrent(t *testing.T) {
	t.Parallel()
	cached := op.Cache2(slices.All([]string{"a", "b"}))
	want := []pair[int, string]{{Key: 0, Value: "a"}, {Key: 1, Value: "b"}}
	next, stop := iter.Pull2(cached)
	defer stop()
	k, v, ok := next()
	assertEqual(t, pair[int, string]{Key: k, Value: v}, want[0])
	assertEqual(t, ok, true)
	assertEqual(t, collectPairs(cached), want)
	k, v, ok = next()
	assertEqual(t, pair[int, string]{Key: k, Value: v}, want[1])
	assertEqual(t, ok, true)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Go(func() { assertEqual(t, collectPairs(cached), want) })
	}
	wg.Wait()
}

// 初回評価中の反復も完成した同一キャッシュを待ち、上流を重複実行しない。
func TestCache2ConcurrentFirstIteration(t *testing.T) {
	t.Parallel()
	var starts atomic.Int32
	started, release := make(chan struct{}), make(chan struct{})
	cached := op.Cache2(iter.Seq2[int, string](func(yield func(int, string) bool) {
		starts.Add(1)
		close(started)
		<-release
		yield(1, "a")
		yield(2, "b")
	}))
	results := make(chan []pair[int, string], 4)
	for i := 0; i < cap(results); i++ {
		go func() { results <- collectPairs(cached) }()
	}
	<-started
	assertEqual(t, starts.Load(), int32(1))
	close(release)
	want := []pair[int, string]{{Key: 1, Value: "a"}, {Key: 2, Value: "b"}}
	for i := 0; i < cap(results); i++ {
		assertEqual(t, <-results, want)
	}
	assertEqual(t, starts.Load(), int32(1))
}

func TestCache2Panic(t *testing.T) {
	t.Parallel()
	starts := 0
	cached := op.Cache2(iter.Seq2[int, string](func(yield func(int, string) bool) {
		starts++
		yield(1, "a")
		panic("cache2 source failed")
	}))
	for i := 0; i < 2; i++ {
		yields := 0
		func() {
			defer func() { assertEqual(t, recover(), any("cache2 source failed")) }()
			cached(func(int, string) bool { yields++; return true })
		}()
		assertEqual(t, yields, 0)
	}
	assertEqual(t, starts, 1)
}
