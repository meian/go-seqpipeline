package seq_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

type keyTextOperator struct{}

func (keyTextOperator) Apply(src seq.Seq2[int, int]) seq.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range src {
			if !yield(strconv.Itoa(k), v) {
				return
			}
		}
	}
}

var _ seq.Operator2[int, int, string, int] = keyTextOperator{}

// 外部パッケージはライブラリ側の変更なしに Operator を追加できる。
type repeatOperator struct{}

func (repeatOperator) Apply(src seq.Seq[int]) seq.Seq[string] {
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

var _ seq.Operator[int, string] = repeatOperator{}

func TestOperatorFunc(t *testing.T) {
	op := seq.OperatorFunc[int, string](func(src seq.Seq[int]) seq.Seq[string] {
		return func(yield func(string) bool) {
			for v := range src {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	var _ seq.Operator[int, string] = op
	assertEqual(t, seq.ToSlice(op.Apply(seq.FromSlice([]int{1, 2}))), []string{"1", "2"})
}

func TestOperator2Func(t *testing.T) {
	op := seq.Operator2Func(keyTextOperator{}.Apply)
	var _ seq.Operator2[int, int, string, int] = op
	got := collectPairs(seq.ToSeq2(op.Apply(seq.FromSeq2(slices.All([]int{4, 5})))))
	assertEqual(t, got, []pair[string, int]{{Key: "0", Value: 4}, {Key: "1", Value: 5}})
}

type takeWithOperator[T any] struct {
	n         int
	predicate func(T) bool
}

func (op takeWithOperator[T]) Apply(src seq.Seq[T]) seq.Seq[T] {
	return src.Pipe(seq.Filter(op.predicate)).Pipe(seq.Take(op.n).Type)
}

var _ seq.Operator[int, int] = takeWithOperator[int]{}

func TestComposedOperator(t *testing.T) {
	reads := 0
	op := takeWithOperator[int]{n: 2, predicate: func(v int) bool { return v%2 == 0 }}
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 5, 6}).
		Pipe(seq.Map(func(v int) int { reads++; return v })).Pipe(op.Apply)
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

func TestComposedOperator2(t *testing.T) {
	reads := 0
	op := seq.Operator2Func(func(src seq.Seq2[int, int]) seq.Seq2[int, int] {
		return src.Pipe(seq.Filter2(func(k, v int) bool { return v%2 == 0 })).
			Pipe(seq.Take2(2).Type)
	})
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
