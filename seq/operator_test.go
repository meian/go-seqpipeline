package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

type keyTextOperator struct{}

func (keyTextOperator) Apply(src iter.Seq2[int, int]) iter.Seq2[string, int] {
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

func (repeatOperator) Apply(src iter.Seq[int]) iter.Seq[string] {
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
	op := seq.OperatorFunc[int, string](func(src iter.Seq[int]) iter.Seq[string] {
		return func(yield func(string) bool) {
			for v := range src {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	var _ seq.Operator[int, string] = op
	assertEqual(t, slices.Collect(op.Apply(slices.Values([]int{1, 2}))), []string{"1", "2"})
}

func TestOperator2Func(t *testing.T) {
	op := seq.Operator2Func(keyTextOperator{}.Apply)
	var _ seq.Operator2[int, int, string, int] = op
	got := collectPairs(op.Apply(slices.All([]int{4, 5})))
	assertEqual(t, got, []pair[string, int]{{Key: "0", Value: 4}, {Key: "1", Value: 5}})
}
