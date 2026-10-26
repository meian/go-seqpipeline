package seq_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

// 型引数なしで両方向を接続し、途中終了と再評価が変換をまたいで上流へ伝わることを確認する。
func TestMapAcrossShapes(t *testing.T) {
	t.Parallel()

	reads, pairCalls, valueCalls := 0, 0, 0
	src := iter.Seq[string](func(yield func(string) bool) {
		for _, v := range []string{"Go", "iter", "Go"} {
			reads++
			if !yield(v) {
				return
			}
		}
	})
	pipeline := seq.FromSeq(src).
		Pipe2(op.MapToSeq2(func(s string) ([]int, []byte) {
			pairCalls++
			return []int{len(s)}, []byte(s)
		})).
		Pipe(op.MapToSeq(func(length []int, text []byte) string {
			valueCalls++
			return fmt.Sprintf("%s:%d", text, length[0])
		}))
	assertEqual(t, []int{reads, pairCalls, valueCalls}, []int{0, 0, 0})
	want := []string{"Go:2", "iter:4", "Go:2"}
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), want)
		assertEqual(t, []int{reads, pairCalls, valueCalls}, []int{round * 3, round * 3, round * 3})
	}
	reads, pairCalls, valueCalls = 0, 0, 0
	for v := range pipeline {
		assertEqual(t, v, want[0])
		break
	}
	assertEqual(t, []int{reads, pairCalls, valueCalls}, []int{1, 1, 1})
	assertEqual(t, seq.ToSlice(pipeline.Pipe(op.Take(2).Operator)), want[:2])
	assertEqual(t, []int{reads, pairCalls, valueCalls}, []int{3, 3, 3})
}

// どちらの入力形からも新 API で合成を開始し、往復前後の出力型に応じて seq に接続する。
func TestMapComposeAcrossShapes(t *testing.T) {
	t.Parallel()

	toPairs := op.MapToSeq2(func(s string) ([]int, []byte) { return []int{len(s)}, []byte(s) })
	toValue := op.MapToSeq(func(length []int, text []byte) string {
		return fmt.Sprintf("%s:%d", text, length[0])
	})
	fromValues := op.Compose().Pipe2(toPairs)
	valueRoundTrip := fromValues.Pipe(toValue)
	src := seq.FromSlice([]string{"Go", "iter", "Go"})
	pairs := src.Pipe2(fromValues.Operator)
	pairWant := []pair[[]int, []byte]{
		{Key: []int{2}, Value: []byte("Go")},
		{Key: []int{4}, Value: []byte("iter")},
		{Key: []int{2}, Value: []byte("Go")},
	}
	assertEqual(t, collectPairs(seq.ToSeq2(pairs)), pairWant)
	valueWant := []string{"Go:2", "iter:4", "Go:2"}
	assertEqual(t, seq.ToSlice(src.Pipe(valueRoundTrip.Operator)), valueWant)

	fromPairs := op.Compose2().Pipe(toValue)
	pairRoundTrip := fromPairs.Pipe2(toPairs)
	assertEqual(t, seq.ToSlice(pairs.Pipe(fromPairs.Operator)), valueWant)
	roundTripWant := []pair[[]int, []byte]{
		{Key: []int{4}, Value: []byte("Go:2")},
		{Key: []int{6}, Value: []byte("iter:4")},
		{Key: []int{4}, Value: []byte("Go:2")},
	}
	assertEqual(t, collectPairs(seq.ToSeq2(pairs.Pipe2(pairRoundTrip.Operator))), roundTripWant)

	other := seq.FromSeq2(slices.All([]string{"a", "bbb"})).
		Pipe(op.MapToSeq(func(index int, s string) string { return fmt.Sprintf("%d:%s", index, s) })).
		Pipe2(fromValues.Operator)
	otherWant := []pair[[]int, []byte]{
		{Key: []int{3}, Value: []byte("0:a")},
		{Key: []int{5}, Value: []byte("1:bbb")},
	}
	assertEqual(t, collectPairs(seq.ToSeq2(other)), otherWant)
}
