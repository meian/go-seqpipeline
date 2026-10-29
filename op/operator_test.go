package op_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

func TestOperator(t *testing.T) {
	t.Parallel()
	var factory func(func(int) string) op.Operator[int, string] = op.Map //nolint:staticcheck // ST1023/QF1011: ジェネリック関数の型推論に代入先の型が必要。
	ope := factory(strconv.Itoa)
	assertEqual(t, slices.Collect(ope(slices.Values([]int{1, 22}))), []string{"1", "22"})
	assertEqual(t, slices.Collect(ope(nil)), []string(nil))
}

func TestOperator2(t *testing.T) {
	t.Parallel()
	var factory func(func(int, string) (string, int)) op.Operator2[int, string, string, int] = op.Map2 //nolint:staticcheck // ST1023/QF1011: ジェネリック関数の型推論に代入先の型が必要。
	ope := factory(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	want := []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(ope(slices.All([]string{"a", "bbb"}))), want)
	assertEqual(t, collectPairs(ope(nil)), []pair[string, int](nil))
}

func TestToSeqOperator(t *testing.T) {
	t.Parallel()
	var ope op.ToSeqOperator[int, string, string] = op.Values
	assertEqual(t, slices.Collect(ope(slices.All([]string{"a", "bbb"}))), []string{"a", "bbb"})
	assertEqual(t, slices.Collect(ope(nil)), []string(nil))
}

func TestToSeq2Operator(t *testing.T) {
	t.Parallel()
	var ope op.ToSeq2Operator[int, int, string] = op.Zip(slices.Values([]string{"a", "b"})).Operator
	want := []pair[int, string]{{Key: 1, Value: "a"}, {Key: 2, Value: "b"}}
	assertEqual(t, collectPairs(ope(slices.Values([]int{1, 2, 3}))), want)
	assertEqual(t, collectPairs(ope(nil)), []pair[int, string](nil))
}
