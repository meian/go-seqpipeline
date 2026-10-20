package op_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

func TestKeys(t *testing.T) {
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
	var ope op.ToSeqOperator[int, int, int] = op.Keys[int, int]
	projected := ope(src)
	assertEqual(t, started, 0)
	for i := 0; i < 2; i++ {
		assertEqual(t, slices.Collect(projected), want)
	}
	assertEqual(t, started, 2)
	assertEqual(t, visited, 8)
	assertEqual(t, stopped, 2)

	for _, mode := range []string{"yield", "break", "Take"} {
		t.Run(mode, func(t *testing.T) {
			visited, stopped = 0, 0
			var got []int
			switch mode {
			case "yield":
				projected(func(v int) bool {
					got = append(got, v)
					return false
				})
			case "break":
				for v := range projected {
					got = append(got, v)
					break
				}
			case "Take":
				got = slices.Collect(op.Take(1).Operator(projected))
			}
			assertEqual(t, got, want[:1])
			assertEqual(t, visited, 1)
			assertEqual(t, stopped, 1)
		})
	}

	visited, stopped = 0, 0
	assertEqual(t, slices.Collect(op.Take(0).Operator(projected)), []int(nil))
	assertEqual(t, visited, 0)
	assertEqual(t, stopped, 0)

	for _, input := range []iter.Seq2[int, int]{
		nil,
		iter.Seq2[int, int](nil),
		func(func(int, int) bool) {},
	} {
		output := op.Keys(input)
		if output == nil {
			t.Fatal("projection must return a non-nil iterator")
		}
		assertEqual(t, slices.Collect(output), []int(nil))
	}
}

func TestValues(t *testing.T) {
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
	var ope op.ToSeqOperator[int, int, int] = op.Values[int, int]
	projected := ope(src)
	assertEqual(t, started, 0)
	for i := 0; i < 2; i++ {
		assertEqual(t, slices.Collect(projected), want)
	}
	assertEqual(t, started, 2)
	assertEqual(t, visited, 8)
	assertEqual(t, stopped, 2)

	for _, mode := range []string{"yield", "break", "Take"} {
		t.Run(mode, func(t *testing.T) {
			visited, stopped = 0, 0
			var got []int
			switch mode {
			case "yield":
				projected(func(v int) bool {
					got = append(got, v)
					return false
				})
			case "break":
				for v := range projected {
					got = append(got, v)
					break
				}
			case "Take":
				got = slices.Collect(op.Take(1).Operator(projected))
			}
			assertEqual(t, got, want[:1])
			assertEqual(t, visited, 1)
			assertEqual(t, stopped, 1)
		})
	}

	visited, stopped = 0, 0
	assertEqual(t, slices.Collect(op.Take(0).Operator(projected)), []int(nil))
	assertEqual(t, visited, 0)
	assertEqual(t, stopped, 0)

	for _, input := range []iter.Seq2[int, int]{
		nil,
		iter.Seq2[int, int](nil),
		func(func(int, int) bool) {},
	} {
		output := op.Values(input)
		if output == nil {
			t.Fatal("projection must return a non-nil iterator")
		}
		assertEqual(t, slices.Collect(output), []int(nil))
	}
}

func TestKeysNonComparable(t *testing.T) {
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

func TestValuesNonComparable(t *testing.T) {
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
