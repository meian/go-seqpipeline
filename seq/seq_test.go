package seq_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func TestSeqDefinedFunctionType(t *testing.T) {
	t.Parallel()

	src := slices.Values([]string{"1", "2"})
	pipeline := seq.Seq[string](src)
	var got []string
	for v := range pipeline {
		got = append(got, v)
	}
	assertEqual(t, got, []string{"1", "2"})
	got = nil
	pipeline(func(v string) bool {
		got = append(got, v)
		return false
	})
	assertEqual(t, got, []string{"1"})
	assertEqual(t, slices.Collect(iter.Seq[string](pipeline)), []string{"1", "2"})
}

// 入力を評価する変換関数も接続時に呼び出し、入力・戻り値と nil を保持する。
func TestSeqPipe(t *testing.T) {
	t.Parallel()

	reads, calls := 0, 0
	var src seq.Seq[int] = func(yield func(int) bool) {
		reads++
		yield(7)
	}
	var pipe func(op.Operator[int, string]) seq.Seq[string] = src.Pipe //nolint:staticcheck // ST1023/QF1011: ジェネリックメソッドの型推論に代入先の型が必要。
	result := pipe(func(got iter.Seq[int]) iter.Seq[string] {
		calls++
		assertEqual(t, reads, 0)
		for v := range got {
			assertEqual(t, v, 7)
		}
		return func(yield func(string) bool) { yield("ok") }
	})
	assertEqual(t, calls, 1)
	assertEqual(t, reads, 1)
	assertEqual(t, seq.ToSlice(result), []string{"ok"})
	var zero seq.Seq[int]
	var nilCalls int
	nilResult := zero.Pipe(func(got iter.Seq[int]) iter.Seq[string] {
		nilCalls++
		if got != nil {
			t.Fatal("nil input changed")
		}
		return nil
	})
	assertEqual(t, nilCalls, 1)
	if nilResult != nil {
		t.Fatal("nil result changed")
	}
}

// 単一値から組への変換関数を接続時に呼び出し、入力・戻り値と nil を保持する。
func TestSeqPipe2(t *testing.T) {
	t.Parallel()

	reads, calls := 0, 0
	var src seq.Seq[int] = func(yield func(int) bool) {
		reads++
		yield(7)
	}
	var pipe func(op.ToSeq2Operator[int, []int, string]) seq.Seq2[[]int, string] = src.Pipe2 //nolint:staticcheck // ST1023/QF1011: ジェネリックメソッドの型推論に代入先の型が必要。
	result := pipe(func(got iter.Seq[int]) iter.Seq2[[]int, string] {
		calls++
		assertEqual(t, reads, 0)
		for v := range got {
			assertEqual(t, v, 7)
		}
		return func(yield func([]int, string) bool) { yield([]int{1}, "ok") }
	})
	assertEqual(t, calls, 1)
	assertEqual(t, reads, 1)
	assertEqual(t, collectPairs(seq.ToSeq2(result)), []pair[[]int, string]{{Key: []int{1}, Value: "ok"}})
	var zero seq.Seq[int]
	var nilCalls int
	nilResult := zero.Pipe2(func(got iter.Seq[int]) iter.Seq2[[]int, string] {
		nilCalls++
		if got != nil {
			t.Fatal("nil input changed")
		}
		return nil
	})
	assertEqual(t, nilCalls, 1)
	if nilResult != nil {
		t.Fatal("nil result changed")
	}
}

func TestSeq2DefinedFunctionType(t *testing.T) {
	t.Parallel()

	src2 := func(yield func(string, int) bool) {
		if !yield("0", 1) {
			return
		}
		yield("1", 3)
	}
	pipeline2 := seq.Seq2[string, int](src2)
	var pairs []pair[string, int]
	for k, v := range pipeline2 {
		pairs = append(pairs, pair[string, int]{Key: k, Value: v})
	}
	assertEqual(t, pairs, []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}})
	pairs = nil
	pipeline2(func(k string, v int) bool {
		pairs = append(pairs, pair[string, int]{Key: k, Value: v})
		return false
	})
	assertEqual(t, pairs, []pair[string, int]{{Key: "0", Value: 1}})
	roundTrip := seq.FromSeq2(iter.Seq2[string, int](pipeline2))
	got := collectPairs(seq.ToSeq2(roundTrip))
	want := collectPairs(seq.ToSeq2(pipeline2))
	assertEqual(t, got, want)
}

// 組から単一値への変換関数を接続時に呼び出し、入力・戻り値と nil を保持する。
func TestSeq2Pipe(t *testing.T) {
	t.Parallel()

	reads, calls := 0, 0
	var src seq.Seq2[int, bool] = func(yield func(int, bool) bool) {
		reads++
		yield(7, true)
	}
	var pipe func(op.ToSeqOperator[int, bool, string]) seq.Seq[string] = src.Pipe //nolint:staticcheck // ST1023/QF1011: ジェネリックメソッドの型推論に代入先の型が必要。
	result := pipe(func(got iter.Seq2[int, bool]) iter.Seq[string] {
		calls++
		assertEqual(t, reads, 0)
		for k, v := range got {
			assertEqual(t, k, 7)
			assertEqual(t, v, true)
		}
		return func(yield func(string) bool) { yield("ok") }
	})
	assertEqual(t, calls, 1)
	assertEqual(t, reads, 1)
	assertEqual(t, seq.ToSlice(result), []string{"ok"})
	var zero seq.Seq2[int, bool]
	var nilCalls int
	nilResult := zero.Pipe(func(got iter.Seq2[int, bool]) iter.Seq[string] {
		nilCalls++
		if got != nil {
			t.Fatal("nil input changed")
		}
		return nil
	})
	assertEqual(t, nilCalls, 1)
	if nilResult != nil {
		t.Fatal("nil result changed")
	}
}

// 組を変換する関数を接続時に呼び出し、入力・戻り値と nil を保持する。
func TestSeq2Pipe2(t *testing.T) {
	t.Parallel()

	reads, calls := 0, 0
	var src seq.Seq2[int, bool] = func(yield func(int, bool) bool) {
		reads++
		yield(7, true)
	}
	var pipe func(op.Operator2[int, bool, []int, string]) seq.Seq2[[]int, string] = src.Pipe2 //nolint:staticcheck // ST1023/QF1011: ジェネリックメソッドの型推論に代入先の型が必要。
	result := pipe(func(got iter.Seq2[int, bool]) iter.Seq2[[]int, string] {
		calls++
		assertEqual(t, reads, 0)
		for k, v := range got {
			assertEqual(t, k, 7)
			assertEqual(t, v, true)
		}
		return func(yield func([]int, string) bool) { yield([]int{1}, "ok") }
	})
	assertEqual(t, calls, 1)
	assertEqual(t, reads, 1)
	assertEqual(t, collectPairs(seq.ToSeq2(result)), []pair[[]int, string]{{Key: []int{1}, Value: "ok"}})
	var zero seq.Seq2[int, bool]
	var nilCalls int
	nilResult := zero.Pipe2(func(got iter.Seq2[int, bool]) iter.Seq2[[]int, string] {
		nilCalls++
		if got != nil {
			t.Fatal("nil input changed")
		}
		return nil
	})
	assertEqual(t, nilCalls, 1)
	if nilResult != nil {
		t.Fatal("nil result changed")
	}
}
