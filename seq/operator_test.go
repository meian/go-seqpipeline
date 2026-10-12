package seq_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

// 外部パッケージはライブラリ側の変更なしに変換関数を追加できる。
func repeat(src seq.Seq[int]) seq.Seq[string] {
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

func TestTransform(t *testing.T) {
	op := func(src seq.Seq[int]) seq.Seq[string] {
		return func(yield func(string) bool) {
			for v := range src {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	}
	assertEqual(t, seq.ToSlice(seq.FromSlice([]int{1, 2}).Pipe(op)), []string{"1", "2"})
}

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
		Pipe(seq.Map(func(v int) int { reads++; return v })).Pipe(op)
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

func keyText(src seq.Seq2[int, int]) seq.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range src {
			if !yield(strconv.Itoa(k), v) {
				return
			}
		}
	}
}

func TestTransform2(t *testing.T) {
	op := keyText
	got := collectPairs(seq.ToSeq2(seq.FromSeq2(slices.All([]int{4, 5})).Pipe(op)))
	assertEqual(t, got, []pair[string, int]{{Key: "0", Value: 4}, {Key: "1", Value: 5}})
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
	assertEqual(t, collectPairs(seq.ToSeq2(limit(filtered(mapped(src))))), want)
	pipeline := src.Pipe(custom).
		Pipe(seq.Map2(func(k string, v int) (string, string) { return k, strconv.Itoa(v) })).
		Pipe(filtered).Pipe(limit)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
}

func TestComposedOperator2(t *testing.T) {
	reads := 0
	var op seq.Operator2[int, int, int, int] = func(src seq.Seq2[int, int]) seq.Seq2[int, int] {
		return src.Pipe(seq.Filter2(func(k, v int) bool { return v%2 == 0 })).
			Pipe(seq.Take2(2).Operator)
	}
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3, 4, 5, 6})).
		Pipe(seq.Map2(func(k, v int) (int, int) { reads++; return k, v })).Pipe(op)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), []pair[int, int]{{Key: 1, Value: 2}, {Key: 3, Value: 4}})
		assertEqual(t, reads, round*4)
	}
	reads = 0
	for range pipeline {
		break
	}
	assertEqual(t, reads, 2)
}
