package op_test

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// 採用・不採用・重複を含む入力で、順序・再評価・途中停止と全件不採用を確認する。
func TestFilter(t *testing.T) {
	var state inputTrace
	calls := 0
	minimum := 1
	src := traceValues([]int{1, 3, 2, 3, 1}, &state)
	filtered := op.Filter(func(v int) bool {
		calls++
		return v > minimum
	})(src)
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	want := []int{3, 2, 3}
	for i := 1; i <= 2; i++ {
		assertEqual(t, slices.Collect(filtered), want)
		assertEqual(t, []int{state.reads, calls}, []int{5 * i, 5 * i})
	}
	state, calls = inputTrace{}, 0
	filtered(func(v int) bool {
		assertEqual(t, v, want[0])
		return false
	})
	assertEqual(t, state, inputTrace{starts: 1, reads: 2, stops: 1})
	assertEqual(t, calls, 2)
	minimum = 3
	state, calls = inputTrace{}, 0
	assertEqual(t, slices.Collect(filtered), []int(nil))
	assertEqual(t, state, inputTrace{starts: 1, reads: 5})
	assertEqual(t, calls, 5)
}

func TestFilterNilInput(t *testing.T) {
	iterator := op.Filter(func(v int) bool {
		t.Fatal("Filter callback called")
		return true
	})(nil)
	for range iterator {
		t.Fatal("nil input yielded a value")
	}
}

// 採用・不採用・重複を含む入力で、順序・再評価・途中停止と全件不採用を確認する。
func TestFilter2(t *testing.T) {
	var state inputTrace
	calls := 0
	minimum := 1
	src := tracePairs([]pair[int, string]{
		{1, "aa"}, {3, "b"}, {3, "cc"}, {2, "ddd"}, {3, "cc"}, {1, "d"},
	}, &state)
	filtered := op.Filter2(func(k int, v string) bool {
		calls++
		return k > minimum && len(v) > 1
	})(src)
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	want := []pair[int, string]{{3, "cc"}, {2, "ddd"}, {3, "cc"}}
	for i := 1; i <= 2; i++ {
		assertEqual(t, collectPairs(filtered), want)
		assertEqual(t, []int{state.reads, calls}, []int{6 * i, 6 * i})
	}
	state, calls = inputTrace{}, 0
	filtered(func(k int, v string) bool {
		assertEqual(t, pair[int, string]{k, v}, want[0])
		return false
	})
	assertEqual(t, state, inputTrace{starts: 1, reads: 3, stops: 1})
	assertEqual(t, calls, 3)
	minimum = 3
	state, calls = inputTrace{}, 0
	assertEqual(t, collectPairs(filtered), []pair[int, string](nil))
	assertEqual(t, state, inputTrace{starts: 1, reads: 6})
	assertEqual(t, calls, 6)
}

func TestFilter2NilInput(t *testing.T) {
	iterator := op.Filter2(func(k int, v string) bool {
		t.Fatal("Filter2 callback called")
		return true
	})(nil)
	for range iterator {
		t.Fatal("nil input yielded a value")
	}
}
