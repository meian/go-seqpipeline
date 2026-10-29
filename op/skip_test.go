package op_test

import (
	"iter"
	"math"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

func TestSkipOperator(t *testing.T) {
	t.Parallel()
	values := []int{3, 1, 1, 2}
	for _, tc := range []struct {
		name string
		n    int
		want []int
	}{
		{name: "minimum", n: math.MinInt, want: values},
		{name: "negative", n: -1, want: values},
		{name: "zero", n: 0, want: values},
		{name: "prefix", n: 1, want: values[1:]},
		{name: "all", n: len(values)},
		{name: "larger", n: len(values) + 1},
		{name: "maximum", n: math.MaxInt},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var state inputTrace
			src := traceValues(values, &state)
			iterator := op.Skip(tc.n).Operator(src)
			assertEqual(t, []int{state.starts, state.reads}, []int{0, 0})
			for i := 1; i <= 2; i++ {
				assertEqual(t, slices.Collect(iterator), tc.want)
				assertEqual(t, []int{state.starts, state.reads}, []int{i, len(values) * i})
			}
		})
	}
}

func TestSkipOperatorEmptyAndNil(t *testing.T) {
	t.Parallel()
	for _, src := range []iter.Seq[int]{nil, func(func(int) bool) {}} {
		assertEqual(t, slices.Collect(op.Skip(1).Operator(src)), []int(nil))
	}
}

func TestSkipOperatorDownstreamStop(t *testing.T) {
	t.Parallel()
	for _, n := range []int{-1, 0, 2} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			t.Parallel()
			var state inputTrace
			src := traceValues([]int{1, 2, 3, 4, 5}, &state)
			iterator := op.Skip(n).Operator(src)
			for i := 1; i <= 2; i++ {
				for v := range iterator {
					assertEqual(t, v, max(n, 0)+1)
					break
				}
				assertEqual(t, []int{state.reads, state.stops}, []int{(max(n, 0) + 1) * i, i})
			}
		})
	}
}

func TestSkipOperator2(t *testing.T) {
	t.Parallel()
	pairs := []pair[int, string]{
		{Key: 3, Value: "c"},
		{Key: 1, Value: "a"},
		{Key: 1, Value: "a"},
		{Key: 2, Value: "b"},
	}
	for _, tc := range []struct {
		name string
		n    int
		want []pair[int, string]
	}{
		{name: "minimum", n: math.MinInt, want: pairs},
		{name: "negative", n: -1, want: pairs},
		{name: "zero", n: 0, want: pairs},
		{name: "prefix", n: 1, want: pairs[1:]},
		{name: "all", n: len(pairs)},
		{name: "larger", n: len(pairs) + 1},
		{name: "maximum", n: math.MaxInt},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var state inputTrace
			src := tracePairs(pairs, &state)
			iterator := op.Skip(tc.n).Operator2(src)
			assertEqual(t, []int{state.starts, state.reads}, []int{0, 0})
			for i := 1; i <= 2; i++ {
				assertEqual(t, collectPairs(iterator), tc.want)
				assertEqual(t, []int{state.starts, state.reads}, []int{i, len(pairs) * i})
			}
		})
	}
}

func TestSkipOperator2EmptyAndNil(t *testing.T) {
	t.Parallel()
	for _, src := range []iter.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
		assertEqual(t, collectPairs(op.Skip(1).Operator2(src)), []pair[int, string](nil))
	}
}

func TestSkipOperator2DownstreamStop(t *testing.T) {
	t.Parallel()
	for _, n := range []int{-1, 0, 2} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			t.Parallel()
			var state inputTrace
			src := tracePairs([]pair[int, int]{{1, 10}, {2, 20}, {3, 30}, {4, 40}, {5, 50}}, &state)
			iterator := op.Skip(n).Operator2(src)
			for i := 1; i <= 2; i++ {
				for k, v := range iterator {
					assertEqual(t, k, max(n, 0)+1)
					assertEqual(t, v, (max(n, 0)+1)*10)
					break
				}
				assertEqual(t, []int{state.reads, state.stops}, []int{(max(n, 0) + 1) * i, i})
			}
		})
	}
}

func TestSkipConfigurationReuse(t *testing.T) {
	t.Parallel()
	skip := op.Skip(1)
	ope := op.Operator[int, int](skip.Operator)
	ope2 := op.Operator2[[]int, string, []int, string](skip.Operator2)
	values := slices.Values([]int{1, 2, 3})
	pairs := iter.Seq2[[]int, string](func(yield func([]int, string) bool) {
		if yield([]int{1}, "a") {
			yield([]int{2}, "b")
		}
	})
	first, second := ope(values), skip.Operator(values)
	paired := ope2(pairs)
	texts := skip.Operator(slices.Values([]string{"a", "b"}))
	for i := 0; i < 2; i++ {
		assertEqual(t, slices.Collect(first), []int{2, 3})
		assertEqual(t, slices.Collect(second), []int{2, 3})
		assertEqual(t, slices.Collect(texts), []string{"b"})
		want := []pair[[]int, string]{{Key: []int{2}, Value: "b"}}
		assertEqual(t, collectPairs(paired), want)
	}
}
