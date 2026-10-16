package seq_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestSkipOperator(t *testing.T) {
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
			starts, reads := 0, 0
			src := seq.Seq[int](func(yield func(int) bool) {
				starts++
				for _, v := range values {
					reads++
					if !yield(v) {
						return
					}
				}
			})
			iterator := seq.Skip(tc.n).Operator(src)
			assertEqual(t, []int{starts, reads}, []int{0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, seq.ToSlice(iterator), tc.want)
				assertEqual(t, []int{starts, reads}, []int{round, len(values) * round})
			}
		})
	}
}

func TestSkipOperator2(t *testing.T) {
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
			starts, reads := 0, 0
			src := seq.Seq2[int, string](func(yield func(int, string) bool) {
				starts++
				for _, p := range pairs {
					reads++
					if !yield(p.Key, p.Value) {
						return
					}
				}
			})
			iterator := seq.Skip(tc.n).Operator2(src)
			assertEqual(t, []int{starts, reads}, []int{0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, collectPairs(seq.ToSeq2(iterator)), tc.want)
				assertEqual(t, []int{starts, reads}, []int{round, len(pairs) * round})
			}
		})
	}
}

func TestSkipOperatorEmptyAndNil(t *testing.T) {
	for _, n := range []int{-1, 0, 1} {
		for _, src := range []seq.Seq[int]{nil, func(func(int) bool) {}} {
			iterator := seq.Skip(n).Operator(src)
			for round := 0; round < 2; round++ {
				assertEqual(t, seq.ToSlice(iterator), []int(nil))
			}
		}
	}
}

func TestSkipOperator2EmptyAndNil(t *testing.T) {
	for _, n := range []int{-1, 0, 1} {
		for _, src := range []seq.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
			iterator := seq.Skip(n).Operator2(src)
			for round := 0; round < 2; round++ {
				assertEqual(t, collectPairs(seq.ToSeq2(iterator)), []pair[int, string](nil))
			}
		}
	}
}

func TestSkipOperatorDownstreamStop(t *testing.T) {
	for _, n := range []int{-1, 0, 2} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			reads, stops := 0, 0
			src := seq.Seq[int](func(yield func(int) bool) {
				for i := 1; i <= 5; i++ {
					reads++
					if !yield(i) {
						stops++
						return
					}
				}
			})
			iterator := seq.Skip(n).Operator(src)
			for round := 1; round <= 2; round++ {
				for v := range iterator {
					assertEqual(t, v, max(n, 0)+1)
					break
				}
				assertEqual(t, []int{reads, stops}, []int{(max(n, 0) + 1) * round, round})
			}
		})
	}
}

func TestSkipOperator2DownstreamStop(t *testing.T) {
	for _, n := range []int{-1, 0, 2} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			reads, stops := 0, 0
			src := seq.Seq2[int, int](func(yield func(int, int) bool) {
				for i := 1; i <= 5; i++ {
					reads++
					if !yield(i, i*10) {
						stops++
						return
					}
				}
			})
			iterator := seq.Skip(n).Operator2(src)
			for round := 1; round <= 2; round++ {
				for k, v := range iterator {
					assertEqual(t, k, max(n, 0)+1)
					assertEqual(t, v, (max(n, 0)+1)*10)
					break
				}
				assertEqual(t, []int{reads, stops}, []int{(max(n, 0) + 1) * round, round})
			}
		})
	}
}

func TestSkipConfigurationReuse(t *testing.T) {
	skip := seq.Skip(1)
	var op seq.Operator[int, int] = skip.Operator
	var op2 seq.Operator2[[]int, string, []int, string] = skip.Operator2
	values := seq.Seq[int](func(yield func(int) bool) {
		for _, v := range []int{1, 2, 3} {
			if !yield(v) {
				return
			}
		}
	})
	pairs := seq.Seq2[[]int, string](func(yield func([]int, string) bool) {
		if yield([]int{1}, "a") {
			yield([]int{2}, "b")
		}
	})
	first, second := op(values), skip.Operator(values)
	paired := op2(pairs)
	texts := skip.Operator(seq.Seq[string](func(yield func(string) bool) {
		if yield("a") {
			yield("b")
		}
	}))
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(first), []int{2, 3})
		assertEqual(t, seq.ToSlice(second), []int{2, 3})
		assertEqual(t, seq.ToSlice(texts), []string{"b"})
		want := []pair[[]int, string]{{Key: []int{2}, Value: "b"}}
		assertEqual(t, collectPairs(seq.ToSeq2(paired)), want)
	}
}
