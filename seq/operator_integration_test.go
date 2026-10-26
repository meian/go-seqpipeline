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

// Map・Filter の生成関数の戻り値の型を確認し、名前付き関数型を直接適用・Pipe 接続する。
func TestOperatorDirectCall(t *testing.T) {
	t.Parallel()

	var mapFactory func(func(int) string) op.Operator[int, string] = op.Map[int, string]
	var filterFactory func(func(string) bool) op.Operator[string, string] = op.Filter[string]
	mapped := mapFactory(strconv.Itoa)
	filtered := filterFactory(func(v string) bool { return len(v) > 1 })
	var limit op.Operator[string, string] = op.Take(1).Operator
	var custom op.Operator[int, string] = repeat
	src := seq.FromSlice([]int{1, 22, 333})
	assertEqual(t, slices.Collect(limit(filtered(mapped(seq.ToSeq(src))))), []string{"22"})
	pipeline := src.Pipe(mapped).Pipe(filtered).Pipe(limit)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"22"})
	}
	assertEqual(t, seq.ToSlice(src.Pipe(custom).Pipe(filtered).Pipe(limit)), []string{"22"})
}

func TestComposedOperator(t *testing.T) {
	t.Parallel()

	reads := 0
	ope := takeWith(2, func(v int) bool { return v%2 == 0 })
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 5, 6}).
		Pipe(op.Map(func(v int) int {
			reads++
			return v
		})).Pipe(ope)
	assertEqual(t, reads, 0)
	for i := 1; i <= 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []int{2, 4})
		assertEqual(t, reads, i*4)
	}
	reads = 0
	for range pipeline {
		break
	}
	assertEqual(t, reads, 2)
}

// 独自の変換関数を Map・Take に接続し、遅延評価と再評価ごとの上流停止を確認する。
func TestTransformPipeline(t *testing.T) {
	t.Parallel()

	ope := repeat
	calls := 0
	pipeline := seq.FromSlice([]int{1, 2, 3}).
		Pipe(op.Map(func(v int) int {
			calls++
			return v
		})).Pipe(ope).Pipe(op.Take(1).Operator)
	assertEqual(t, calls, 0)
	for i := 1; i <= 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1"})
		assertEqual(t, calls, i)
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

// Map2・Filter2 の生成関数の戻り値の型を確認し、名前付き関数型を直接適用・Pipe2 接続する。
func TestOperator2DirectCall(t *testing.T) {
	t.Parallel()

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
	for i := 0; i < 2; i++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	}
}

func TestComposedOperator2(t *testing.T) {
	t.Parallel()

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
	for i := 1; i <= 2; i++ {
		want := []pair[int, int]{{Key: 1, Value: 2}, {Key: 3, Value: 4}}
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
		assertEqual(t, reads, i*4)
	}
	reads = 0
	for range pipeline {
		break
	}
	assertEqual(t, reads, 2)
}

// 独自の組変換を直接適用・Pipe2 接続でき、接続時の遅延評価と下流の停止を保つ。
func TestExternalTransforms2(t *testing.T) {
	t.Parallel()

	ope := func(src iter.Seq2[string, int]) iter.Seq2[string, string] {
		return func(yield func(string, string) bool) {
			for k, v := range src {
				if !yield(k, strconv.Itoa(v)) {
					return
				}
			}
		}
	}
	direct := ope(seq.ToSeq2(seq.FromMap(map[string]int{"1": 10})))
	assertEqual(t, collectPairs(direct), []pair[string, string]{{Key: "1", Value: "10"}})
	reads := 0
	src := func(yield func(int, int) bool) {
		for i := 1; i <= 3; i++ {
			reads++
			if !yield(i, i*10) {
				return
			}
		}
	}
	pipeline := seq.FromSeq2(src).Pipe2(keyText).Pipe2(ope)
	assertEqual(t, reads, 0)
	for k, v := range seq.ToSeq2(pipeline) {
		assertEqual(t, k, "1")
		assertEqual(t, v, "10")
		break
	}
	assertEqual(t, reads, 1)
	want := []pair[string, string]{
		{Key: "1", Value: "10"},
		{Key: "2", Value: "20"},
		{Key: "3", Value: "30"},
	}
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
}

// 単一値を組に変換する名前付き関数型を直接適用・Pipe2 接続し、nil 入力も確認する。
func TestToSeq2OperatorDirectCall(t *testing.T) {
	t.Parallel()

	var ope op.ToSeq2Operator[string, []int, string] = withLength
	src := seq.FromSlice([]string{"a", "bbb"})
	want := []pair[[]int, string]{{Key: []int{1}, Value: "a"}, {Key: []int{3}, Value: "bbb"}}
	assertEqual(t, collectPairs(ope(seq.ToSeq(src))), want)
	assertEqual(t, collectPairs(seq.ToSeq2(src.Pipe2(ope))), want)
	assertEqual(t, collectPairs(ope(nil)), []pair[[]int, string](nil))
}

// 組を単一値に変換する名前付き関数型を Pipe に接続できる。
func TestToSeqOperatorPipeline(t *testing.T) {
	t.Parallel()

	var ope op.ToSeqOperator[int, string, string] = op.Values[int, string]
	src := seq.FromSeq2(slices.All([]string{"a", "bbb"}))
	assertEqual(t, seq.ToSlice(src.Pipe(ope)), []string{"a", "bbb"})
}
