package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func takeWith[T any](n int, predicate func(T) bool) op.Operator[T, T] {
	return func(src iter.Seq[T]) iter.Seq[T] {
		return op.Take(n).Operator(op.Filter(predicate)(src))
	}
}

func TestOperatorDirectCall(t *testing.T) {
	// 生成関数の戻り値が Operator 型であることも確認する。
	var mapFactory func(func(int) string) op.Operator[int, string] = op.Map[int, string]
	var filterFactory func(func(string) bool) op.Operator[string, string] = op.Filter[string]
	mapped := mapFactory(strconv.Itoa)
	filtered := filterFactory(func(v string) bool { return len(v) > 1 })
	var limit op.Operator[string, string] = op.Take(1).Operator
	var custom op.Operator[int, string] = repeat
	src := seq.FromSlice([]int{1, 22, 333})
	assertEqual(t, slices.Collect(limit(filtered(mapped(seq.ToSeq(src))))), []string{"22"})
	assertEqual(t, seq.ToSlice(src.Pipe(custom).Pipe(filtered).Pipe(limit)), []string{"22"})
}

func TestComposedOperator(t *testing.T) {
	reads := 0
	ope := takeWith(2, func(v int) bool { return v%2 == 0 })
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 5, 6}).
		Pipe(op.Map(func(v int) int {
			reads++
			return v
		})).Pipe(ope)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []int{2, 4})
		assertEqual(t, reads, round*4)
	}
	reads = 0
	for range pipeline {
		break
	}
	assertEqual(t, reads, 2)
}

func TestOperator2DirectCall(t *testing.T) {
	// 生成関数の戻り値が Operator2 型であることも確認する。
	var mapFactory func(func(int, int) (string, string)) op.Operator2[int, int, string, string] = op.Map2[int, int, string, string]
	var filterFactory func(func(string, string) bool) op.Operator2[string, string, string, string] = op.Filter2[string, string]
	mapped := mapFactory(func(k, v int) (string, string) { return strconv.Itoa(k), strconv.Itoa(v) })
	filtered := filterFactory(func(k, v string) bool { return len(v) > 1 })
	var limit op.Operator2[string, string, string, string] = op.Take(1).Operator2
	var custom op.Operator2[int, int, string, int] = keyText
	src := seq.FromSeq2(slices.All([]int{1, 22, 333}))
	want := []pair[string, string]{{Key: "1", Value: "22"}}
	direct := limit(filtered(mapped(seq.ToSeq2(src))))
	assertEqual(t, collectPairs(direct), want)
	pipeline := src.Pipe2(custom).
		Pipe2(op.Map2(func(k string, v int) (string, string) { return k, strconv.Itoa(v) })).
		Pipe2(filtered).Pipe2(limit)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
}

func TestComposedOperator2(t *testing.T) {
	reads := 0
	var ope op.Operator2[int, int, int, int] = func(src iter.Seq2[int, int]) iter.Seq2[int, int] {
		return op.Take(2).Operator2(op.Filter2(func(k, v int) bool { return v%2 == 0 })(src))
	}
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3, 4, 5, 6})).
		Pipe2(op.Map2(func(k, v int) (int, int) {
			reads++
			return k, v
		})).Pipe2(ope)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		want := []pair[int, int]{{Key: 1, Value: 2}, {Key: 3, Value: 4}}
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
		assertEqual(t, reads, round*4)
	}
	reads = 0
	for range pipeline {
		break
	}
	assertEqual(t, reads, 2)
}

// TestToSeq2OperatorDirectCall は、名前付き関数型への代入、直接呼び出し、Pipe2 接続を確認する。
func TestToSeq2OperatorDirectCall(t *testing.T) {
	var ope op.ToSeq2Operator[string, []int, string] = withLength
	src := seq.FromSlice([]string{"a", "bbb"})
	want := []pair[[]int, string]{{Key: []int{1}, Value: "a"}, {Key: []int{3}, Value: "bbb"}}
	assertEqual(t, collectPairs(ope(seq.ToSeq(src))), want)
	assertEqual(t, collectPairs(seq.ToSeq2(src.Pipe2(ope))), want)
	assertEqual(t, collectPairs(ope(nil)), []pair[[]int, string](nil))
}

// TestToSeqOperatorDirectCall は、関数の代入、直接呼び出し、Pipe 接続を確認する。
func TestToSeqOperatorDirectCall(t *testing.T) {
	var ope op.ToSeqOperator[int, string, string] = op.Values[int, string]
	src := seq.FromSeq2(slices.All([]string{"a", "bbb"}))
	assertEqual(t, slices.Collect(ope(seq.ToSeq2(src))), []string{"a", "bbb"})
	assertEqual(t, seq.ToSlice(src.Pipe(ope)), []string{"a", "bbb"})
	assertEqual(t, slices.Collect(ope(nil)), []string(nil))
}

func repeat(src iter.Seq[int]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for v := range src {
			for i := 0; i < 2; i++ {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	}
}

func keyText(src iter.Seq2[int, int]) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range src {
			if !yield(strconv.Itoa(k), v) {
				return
			}
		}
	}
}
