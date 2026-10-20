package seq_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

// TestTakeOperator は、遅延評価・件数制限・上流停止・再評価・0 以下での非評価を確認する。
func TestTakeOperator(t *testing.T) {
	for _, tc := range []struct {
		name   string
		n      int
		want   []int
		reads  int
		starts int
		stops  int
	}{
		{name: "negative", n: -1},
		{name: "zero", n: 0},
		{name: "one", n: 1, want: []int{1}, reads: 1, starts: 1, stops: 1},
		{name: "prefix", n: 2, want: []int{1, 2}, reads: 2, starts: 1, stops: 1},
		{name: "all", n: 3, want: []int{1, 2, 3}, reads: 3, starts: 1, stops: 1},
		{name: "larger", n: 4, want: []int{1, 2, 3}, reads: 3, starts: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads, starts, stops := 0, 0, 0
			src := iter.Seq[int](func(yield func(int) bool) {
				starts++
				for i := 1; i <= 3; i++ {
					reads++
					if !yield(i) {
						stops++
						return
					}
				}
			})
			iterator := seq.Take(tc.n).Operator(src)
			assertEqual(t, []int{reads, starts, stops}, []int{0, 0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, slices.Collect(iterator), tc.want)
				gotCounts := []int{reads, starts, stops}
				wantCounts := []int{tc.reads * round, tc.starts * round, tc.stops * round}
				assertEqual(t, gotCounts, wantCounts)
			}
		})
	}
}

// TestTakeOperatorEmptyAndNilInput は、nil と空入力を再評価しても要素を返さないことを確認する。
func TestTakeOperatorEmptyAndNilInput(t *testing.T) {
	for _, src := range []iter.Seq[int]{nil, func(func(int) bool) {}} {
		iterator := seq.Take(1).Operator(src)
		for round := 0; round < 2; round++ {
			assertEqual(t, slices.Collect(iterator), []int(nil))
		}
	}
}

// TestTakeOperator2 は、組の順序と重複、遅延評価・件数制限・停止・再評価・0 以下での非評価を確認する。
func TestTakeOperator2(t *testing.T) {
	pairs := []pair[int, int]{
		{Key: 1, Value: 10},
		{Key: 0, Value: 20},
		{Key: 1, Value: 30},
	}
	for _, tc := range []struct {
		name   string
		n      int
		want   []pair[int, int]
		reads  int
		starts int
		stops  int
	}{
		{name: "negative", n: -1},
		{name: "zero", n: 0},
		{name: "one", n: 1, want: pairs[:1], reads: 1, starts: 1, stops: 1},
		{name: "prefix", n: 2, want: pairs[:2], reads: 2, starts: 1, stops: 1},
		{name: "all", n: 3, want: pairs[:3], reads: 3, starts: 1, stops: 1},
		{name: "larger", n: 4, want: pairs[:3], reads: 3, starts: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads, starts, stops := 0, 0, 0
			src := iter.Seq2[int, int](func(yield func(int, int) bool) {
				starts++
				for i := 1; i <= 3; i++ {
					reads++
					if !yield(i%2, i*10) {
						stops++
						return
					}
				}
			})
			pipeline := seq.Take(tc.n).Operator2(src)
			iterator := pipeline
			assertEqual(t, []int{reads, starts, stops}, []int{0, 0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, collectPairs(iterator), tc.want)
				gotCounts := []int{reads, starts, stops}
				wantCounts := []int{tc.reads * round, tc.starts * round, tc.stops * round}
				assertEqual(t, gotCounts, wantCounts)
			}
		})
	}
}

// TestTakeOperator2EmptyAndNilInput は、nil と空入力を再評価しても組を返さないことを確認する。
func TestTakeOperator2EmptyAndNilInput(t *testing.T) {
	for _, src := range []iter.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
		iterator := seq.Take(1).Operator2(src)
		for round := 0; round < 2; round++ {
			assertEqual(t, collectPairs(iterator), []pair[int, string](nil))
		}
	}
}

// TestTakeOperator2DownstreamStop は、下流の途中終了が上流へ伝播することを確認する。
func TestTakeOperator2DownstreamStop(t *testing.T) {
	reads, stops := 0, 0
	src := iter.Seq2[int, int](func(yield func(int, int) bool) {
		for i := 1; i <= 3; i++ {
			reads++
			if !yield(i, i*10) {
				stops++
				return
			}
		}
	})
	pipeline := seq.Take(3).Operator2(src)
	for k, v := range pipeline {
		assertEqual(t, k, 1)
		assertEqual(t, v, 10)
		break
	}
	assertEqual(t, []int{reads, stops}, []int{1, 1})
}

// TestTakeOperatorDownstreamStop は、指定数より先に下流が停止した場合も上流を停止することを確認する。
func TestTakeOperatorDownstreamStop(t *testing.T) {
	reads, stops := 0, 0
	src := iter.Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			reads++
			if !yield(i) {
				stops++
				return
			}
		}
	})
	iterator := seq.Take(3).Operator(src)
	for round := 1; round <= 2; round++ {
		for v := range iterator {
			assertEqual(t, v, 1)
			break
		}
		assertEqual(t, []int{reads, stops}, []int{round, round})
	}
}

// TestTakeConfigurationReuse は、同じ設定値を異なる型と複数の入力に適用しても残り件数を共有しないことを確認する。
func TestTakeConfigurationReuse(t *testing.T) {
	take := seq.Take(2)
	var op seq.Operator[int, int] = take.Operator
	var op2 seq.Operator2[[]int, string, []int, string] = take.Operator2
	values := iter.Seq[int](func(yield func(int) bool) {
		for _, v := range []int{3, 3, 1} {
			if !yield(v) {
				return
			}
		}
	})
	pairs := iter.Seq2[[]int, string](func(yield func([]int, string) bool) {
		for _, p := range []pair[[]int, string]{
			{Key: []int{1}, Value: "a"},
			{Key: []int{1}, Value: "a"},
			{Key: []int{2}, Value: "b"},
		} {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	first, second := op(values), take.Operator(values)
	paired := op2(pairs)
	for round := 0; round < 2; round++ {
		assertEqual(t, slices.Collect(first), []int{3, 3})
		assertEqual(t, slices.Collect(second), []int{3, 3})
		want := []pair[[]int, string]{
			{Key: []int{1}, Value: "a"},
			{Key: []int{1}, Value: "a"},
		}
		assertEqual(t, collectPairs(paired), want)
	}
}
