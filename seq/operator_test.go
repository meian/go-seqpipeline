package seq_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func takeWith[T any](n int, predicate func(T) bool) seq.Operator[T, T] {
	return func(src seq.Seq[T]) seq.Seq[T] {
		return src.Pipe(seq.Filter(predicate)).Pipe(seq.Take(n).Operator)
	}
}

func TestOperatorDirectCall(t *testing.T) {
	// 生成関数の戻り値が Operator 型であることも確認する。
	var mapFactory func(func(int) string) seq.Operator[int, string] = seq.Map[int, string]
	var filterFactory func(func(string) bool) seq.Operator[string, string] = seq.Filter[string]
	mapped := mapFactory(strconv.Itoa)
	filtered := filterFactory(func(v string) bool { return len(v) > 1 })
	var limit seq.Operator[string, string] = seq.Take(1).Operator
	var custom seq.Operator[int, string] = repeat
	src := seq.FromSlice([]int{1, 22, 333})
	assertEqual(t, seq.ToSlice(limit(filtered(mapped(src)))), []string{"22"})
	assertEqual(t, seq.ToSlice(src.Pipe(custom).Pipe(filtered).Pipe(limit)), []string{"22"})
}

func TestComposedOperator(t *testing.T) {
	reads := 0
	op := takeWith(2, func(v int) bool { return v%2 == 0 })
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 5, 6}).
		Pipe(seq.Map(func(v int) int {
			reads++
			return v
		})).Pipe(op)
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
	var mapFactory func(func(int, int) (string, string)) seq.Operator2[int, int, string, string] = seq.Map2[int, int, string, string]
	var filterFactory func(func(string, string) bool) seq.Operator2[string, string, string, string] = seq.Filter2[string, string]
	mapped := mapFactory(func(k, v int) (string, string) { return strconv.Itoa(k), strconv.Itoa(v) })
	filtered := filterFactory(func(k, v string) bool { return len(v) > 1 })
	var limit seq.Operator2[string, string, string, string] = seq.Take2(1).Operator
	var custom seq.Operator2[int, int, string, int] = keyText
	src := seq.FromSeq2(slices.All([]int{1, 22, 333}))
	want := []pair[string, string]{{Key: "1", Value: "22"}}
	direct := limit(filtered(mapped(src)))
	assertEqual(t, collectPairs(seq.ToSeq2(direct)), want)
	pipeline := src.Pipe2(custom).
		Pipe2(seq.Map2(func(k string, v int) (string, string) { return k, strconv.Itoa(v) })).
		Pipe2(filtered).Pipe2(limit)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
}

func TestComposedOperator2(t *testing.T) {
	reads := 0
	var op seq.Operator2[int, int, int, int] = func(src seq.Seq2[int, int]) seq.Seq2[int, int] {
		return src.Pipe2(seq.Filter2(func(k, v int) bool { return v%2 == 0 })).
			Pipe2(seq.Take2(2).Operator)
	}
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3, 4, 5, 6})).
		Pipe2(seq.Map2(func(k, v int) (int, int) {
			reads++
			return k, v
		})).Pipe2(op)
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
	var op seq.ToSeq2Operator[string, []int, string] = withLength
	src := seq.FromSlice([]string{"a", "bbb"})
	want := []pair[[]int, string]{{Key: []int{1}, Value: "a"}, {Key: []int{3}, Value: "bbb"}}
	assertEqual(t, collectPairs(seq.ToSeq2(op(src))), want)
	assertEqual(t, collectPairs(seq.ToSeq2(src.Pipe2(op))), want)
	assertEqual(t, collectPairs(seq.ToSeq2(op(nil))), []pair[[]int, string](nil))
}

// TestToSeqOperatorDirectCall は、メソッド式の代入、直接呼び出し、Pipe 接続を確認する。
func TestToSeqOperatorDirectCall(t *testing.T) {
	var op seq.ToSeqOperator[int, string, string] = seq.Seq2[int, string].Values
	src := seq.FromSeq2(slices.All([]string{"a", "bbb"}))
	assertEqual(t, seq.ToSlice(op(src)), []string{"a", "bbb"})
	assertEqual(t, seq.ToSlice(src.Pipe(op)), []string{"a", "bbb"})
	assertEqual(t, seq.ToSlice(op(nil)), []string(nil))
}
