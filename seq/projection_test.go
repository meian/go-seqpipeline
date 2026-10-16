package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestSeq2Keys(t *testing.T) {
	want := []int{3, 1, 3, 2}
	started, visited, stopped := 0, 0, 0
	src := seq.Seq2[int, int](func(yield func(int, int) bool) {
		started++
		defer func() { stopped++ }()
		for _, p := range []pair[int, int]{{3, 20}, {1, 10}, {3, 30}, {2, 20}} {
			visited++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	projected := src.Keys()
	assertEqual(t, started, 0)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(projected), want)
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
				got = seq.ToSlice(projected.Pipe(seq.Take(1).Operator))
			}
			assertEqual(t, got, want[:1])
			assertEqual(t, visited, 1)
			assertEqual(t, stopped, 1)
		})
	}

	visited, stopped = 0, 0
	assertEqual(t, seq.ToSlice(projected.Pipe(seq.Take(0).Operator)), []int(nil))
	assertEqual(t, visited, 0)
	assertEqual(t, stopped, 0)

	pipeline := projected.
		Pipe(seq.Filter(func(v int) bool { return v >= want[0] })).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Take(2).Operator)
	pipelineWant := []string{strconv.Itoa(want[0]), strconv.Itoa(want[2])}
	assertEqual(t, slices.Collect(seq.ToSeq(pipeline)), pipelineWant)

	for _, input := range []seq.Seq2[int, int]{
		nil,
		seq.FromSeq2(iter.Seq2[int, int](nil)),
		func(func(int, int) bool) {},
	} {
		output := input.Keys()
		if output == nil {
			t.Fatal("projection must return a non-nil iterator")
		}
		assertEqual(t, seq.ToSlice(output), []int(nil))
	}
}

func TestSeq2Values(t *testing.T) {
	want := []int{20, 10, 30, 20}
	started, visited, stopped := 0, 0, 0
	src := seq.Seq2[int, int](func(yield func(int, int) bool) {
		started++
		defer func() { stopped++ }()
		for _, p := range []pair[int, int]{{3, 20}, {1, 10}, {3, 30}, {2, 20}} {
			visited++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	projected := src.Values()
	assertEqual(t, started, 0)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(projected), want)
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
				got = seq.ToSlice(projected.Pipe(seq.Take(1).Operator))
			}
			assertEqual(t, got, want[:1])
			assertEqual(t, visited, 1)
			assertEqual(t, stopped, 1)
		})
	}

	visited, stopped = 0, 0
	assertEqual(t, seq.ToSlice(projected.Pipe(seq.Take(0).Operator)), []int(nil))
	assertEqual(t, visited, 0)
	assertEqual(t, stopped, 0)

	pipeline := projected.
		Pipe(seq.Filter(func(v int) bool { return v >= want[0] })).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Take(2).Operator)
	pipelineWant := []string{strconv.Itoa(want[0]), strconv.Itoa(want[2])}
	assertEqual(t, slices.Collect(seq.ToSeq(pipeline)), pipelineWant)

	for _, input := range []seq.Seq2[int, int]{
		nil,
		seq.FromSeq2(iter.Seq2[int, int](nil)),
		func(func(int, int) bool) {},
	} {
		output := input.Values()
		if output == nil {
			t.Fatal("projection must return a non-nil iterator")
		}
		assertEqual(t, seq.ToSlice(output), []int(nil))
	}
}

func TestSeq2KeysNonComparable(t *testing.T) {
	keys := [][]int{{2}, {1}, {2}}
	values := []map[string]int{{"a": 2}, {"b": 1}, {"a": 2}}
	src := seq.Seq2[[]int, map[string]int](func(yield func([]int, map[string]int) bool) {
		for i, k := range keys {
			if !yield(k, values[i]) {
				return
			}
		}
	})
	assertEqual(t, seq.ToSlice(src.Keys()), keys)
}

func TestSeq2ValuesNonComparable(t *testing.T) {
	keys := [][]int{{2}, {1}, {2}}
	values := []map[string]int{{"a": 2}, {"b": 1}, {"a": 2}}
	src := seq.Seq2[[]int, map[string]int](func(yield func([]int, map[string]int) bool) {
		for i, k := range keys {
			if !yield(k, values[i]) {
				return
			}
		}
	})
	assertEqual(t, seq.ToSlice(src.Values()), values)
}
