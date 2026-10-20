package op_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

func TestOperator(t *testing.T) {
	var factory func(func(int) string) op.Operator[int, string] = op.Map[int, string]
	transform := factory(strconv.Itoa)
	assertEqual(t, slices.Collect(transform(slices.Values([]int{1, 22}))), []string{"1", "22"})
	assertEqual(t, slices.Collect(transform(nil)), []string(nil))
}

func TestOperator2(t *testing.T) {
	var factory func(func(int, string) (string, int)) op.Operator2[int, string, string, int] = op.Map2[int, string, string, int]
	transform := factory(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	want := []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(transform(slices.All([]string{"a", "bbb"}))), want)
	assertEqual(t, collectPairs(transform(nil)), []pair[string, int](nil))
}

func TestToSeqOperator(t *testing.T) {
	var transform op.ToSeqOperator[int, string, string] = op.Values[int, string]
	assertEqual(t, slices.Collect(transform(slices.All([]string{"a", "bbb"}))), []string{"a", "bbb"})
	assertEqual(t, slices.Collect(transform(nil)), []string(nil))
}

func TestToSeq2Operator(t *testing.T) {
	var transform op.ToSeq2Operator[int, int, string] = op.Zip(slices.Values([]string{"a", "b"})).Operator
	want := []pair[int, string]{{Key: 1, Value: "a"}, {Key: 2, Value: "b"}}
	assertEqual(t, collectPairs(transform(slices.Values([]int{1, 2, 3}))), want)
	assertEqual(t, collectPairs(transform(nil)), []pair[int, string](nil))
}
