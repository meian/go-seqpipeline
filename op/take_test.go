package op_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// 遅延評価と指定件数での上流停止を確認し、再評価のたびに件数制限が働く。
// 指定数が 0 以下なら入力を開始しない。
func TestTakeOperator(t *testing.T) {
	t.Parallel()
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
			t.Parallel()
			var state inputTrace
			src := traceValues([]int{1, 2, 3}, &state)
			iterator := op.Take(tc.n).Operator(src)
			assertEqual(t, []int{state.reads, state.starts, state.stops}, []int{0, 0, 0})
			for i := 1; i <= 2; i++ {
				assertEqual(t, slices.Collect(iterator), tc.want)
				gotCounts := []int{state.reads, state.starts, state.stops}
				wantCounts := []int{tc.reads * i, tc.starts * i, tc.stops * i}
				assertEqual(t, gotCounts, wantCounts)
			}
		})
	}
}

func TestTakeOperatorEmptyAndNilInput(t *testing.T) {
	t.Parallel()
	for _, src := range []iter.Seq[int]{nil, func(func(int) bool) {}} {
		iterator := op.Take(1).Operator(src)
		assertEqual(t, slices.Collect(iterator), []int(nil))
	}
}

// 指定件数に達する前の下流停止が、再評価のたびに上流へ伝わる。
func TestTakeOperatorDownstreamStop(t *testing.T) {
	t.Parallel()
	var state inputTrace
	src := traceValues([]int{1, 2, 3}, &state)
	iterator := op.Take(3).Operator(src)
	for i := 1; i <= 2; i++ {
		for v := range iterator {
			assertEqual(t, v, 1)
			break
		}
		assertEqual(t, []int{state.reads, state.stops}, []int{i, i})
	}
}

// 組の順序とキーの重複を保ち、遅延評価・件数制限・上流停止が再評価ごとに働く。
// 指定数が 0 以下なら入力を開始しない。
func TestTakeOperator2(t *testing.T) {
	t.Parallel()
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
			t.Parallel()
			var state inputTrace
			src := tracePairs(pairs, &state)
			iterator := op.Take(tc.n).Operator2(src)
			assertEqual(t, []int{state.reads, state.starts, state.stops}, []int{0, 0, 0})
			for i := 1; i <= 2; i++ {
				assertEqual(t, collectPairs(iterator), tc.want)
				gotCounts := []int{state.reads, state.starts, state.stops}
				wantCounts := []int{tc.reads * i, tc.starts * i, tc.stops * i}
				assertEqual(t, gotCounts, wantCounts)
			}
		})
	}
}

func TestTakeOperator2EmptyAndNilInput(t *testing.T) {
	t.Parallel()
	for _, src := range []iter.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
		iterator := op.Take(1).Operator2(src)
		assertEqual(t, collectPairs(iterator), []pair[int, string](nil))
	}
}

// 下流が指定件数より先に終了した場合も、組を生成する上流の反復を停止する。
func TestTakeOperator2DownstreamStop(t *testing.T) {
	t.Parallel()
	var state inputTrace
	src := tracePairs([]pair[int, int]{{1, 10}, {2, 20}, {3, 30}}, &state)
	pipeline := op.Take(3).Operator2(src)
	for k, v := range pipeline {
		assertEqual(t, k, 1)
		assertEqual(t, v, 10)
		break
	}
	assertEqual(t, []int{state.reads, state.stops}, []int{1, 1})
}

// 同じ設定を異なる型と入力に使っても、評価時の残り件数を共有しない。
func TestTakeConfigurationReuse(t *testing.T) {
	t.Parallel()
	take := op.Take(2)
	ope := take.Operator[int]
	ope2 := take.Operator2[[]int, string]
	values := slices.Values([]int{3, 3, 1})
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
	first, second := ope(values), take.Operator(values)
	paired := ope2(pairs)
	for i := 0; i < 2; i++ {
		assertEqual(t, slices.Collect(first), []int{3, 3})
		assertEqual(t, slices.Collect(second), []int{3, 3})
		want := []pair[[]int, string]{
			{Key: []int{1}, Value: "a"},
			{Key: []int{1}, Value: "a"},
		}
		assertEqual(t, collectPairs(paired), want)
	}
}
