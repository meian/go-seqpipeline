package op_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// 順序と重複を保ってキーを取り出し、遅延評価・再評価と下流停止時の解放を確認する。
func TestKeys(t *testing.T) {
	t.Parallel()
	want := []int{3, 1, 3, 2}
	started, visited, stopped := 0, 0, 0
	src := iter.Seq2[int, int](func(yield func(int, int) bool) {
		started++
		defer func() { stopped++ }()
		for _, p := range []pair[int, int]{{3, 20}, {1, 10}, {3, 30}, {2, 20}} {
			visited++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	var ope op.ToSeqOperator[int, int, int] = op.Keys
	projected := ope(src)
	assertEqual(t, started, 0)
	for i := 1; i <= 2; i++ {
		assertEqual(t, slices.Collect(projected), want)
		assertEqual(t, []int{started, visited, stopped}, []int{i, 4 * i, i})
	}

	started, visited, stopped = 0, 0, 0
	var got []int
	for v := range projected {
		got = append(got, v)
		break
	}
	assertEqual(t, got, want[:1])
	assertEqual(t, []int{started, visited, stopped}, []int{1, 1, 1})
}

func TestKeysEmptyInput(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		src  iter.Seq2[int, int]
	}{
		{name: "nil"},
		{name: "empty", src: func(func(int, int) bool) {}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			output := op.Keys(tc.src)
			if output == nil {
				t.Fatal("Keys must return a non-nil iterator")
			}
			assertEqual(t, slices.Collect(output), []int(nil))
		})
	}
}

func TestKeysNonComparable(t *testing.T) {
	t.Parallel()
	keys := [][]int{{2}, {1}, {2}}
	values := []map[string]int{{"a": 2}, {"b": 1}, {"a": 2}}
	src := iter.Seq2[[]int, map[string]int](func(yield func([]int, map[string]int) bool) {
		for i, k := range keys {
			if !yield(k, values[i]) {
				return
			}
		}
	})
	assertEqual(t, slices.Collect(op.Keys(src)), keys)
}

// 順序と重複を保って値を取り出し、遅延評価・再評価と下流停止時の解放を確認する。
func TestValues(t *testing.T) {
	t.Parallel()
	want := []int{20, 10, 30, 20}
	started, visited, stopped := 0, 0, 0
	src := iter.Seq2[int, int](func(yield func(int, int) bool) {
		started++
		defer func() { stopped++ }()
		for _, p := range []pair[int, int]{{3, 20}, {1, 10}, {3, 30}, {2, 20}} {
			visited++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	var ope op.ToSeqOperator[int, int, int] = op.Values
	projected := ope(src)
	assertEqual(t, started, 0)
	for i := 1; i <= 2; i++ {
		assertEqual(t, slices.Collect(projected), want)
		assertEqual(t, []int{started, visited, stopped}, []int{i, 4 * i, i})
	}

	started, visited, stopped = 0, 0, 0
	var got []int
	for v := range projected {
		got = append(got, v)
		break
	}
	assertEqual(t, got, want[:1])
	assertEqual(t, []int{started, visited, stopped}, []int{1, 1, 1})
}

func TestValuesEmptyInput(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		src  iter.Seq2[int, int]
	}{
		{name: "nil"},
		{name: "empty", src: func(func(int, int) bool) {}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			output := op.Values(tc.src)
			if output == nil {
				t.Fatal("Values must return a non-nil iterator")
			}
			assertEqual(t, slices.Collect(output), []int(nil))
		})
	}
}

func TestValuesNonComparable(t *testing.T) {
	t.Parallel()
	keys := [][]int{{2}, {1}, {2}}
	values := []map[string]int{{"a": 2}, {"b": 1}, {"a": 2}}
	src := iter.Seq2[[]int, map[string]int](func(yield func([]int, map[string]int) bool) {
		for i, k := range keys {
			if !yield(k, values[i]) {
				return
			}
		}
	})
	assertEqual(t, slices.Collect(op.Values(src)), values)
}
