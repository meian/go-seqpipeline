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
			iterator := op.Take(tc.n).Operator(src)
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

func TestTakeOperatorEmptyAndNilInput(t *testing.T) {
	for _, src := range []iter.Seq[int]{nil, func(func(int) bool) {}} {
		iterator := op.Take(1).Operator(src)
		for round := 0; round < 2; round++ {
			assertEqual(t, slices.Collect(iterator), []int(nil))
		}
	}
}

// 組の順序とキーの重複を保ち、遅延評価・件数制限・上流停止が再評価ごとに働く。
// 指定数が 0 以下なら入力を開始しない。
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
			pipeline := op.Take(tc.n).Operator2(src)
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

func TestTakeOperator2EmptyAndNilInput(t *testing.T) {
	for _, src := range []iter.Seq2[int, string]{nil, func(func(int, string) bool) {}} {
		iterator := op.Take(1).Operator2(src)
		for round := 0; round < 2; round++ {
			assertEqual(t, collectPairs(iterator), []pair[int, string](nil))
		}
	}
}

// 下流が指定件数より先に終了した場合も、組を生成する上流の反復を停止する。
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
	pipeline := op.Take(3).Operator2(src)
	for k, v := range pipeline {
		assertEqual(t, k, 1)
		assertEqual(t, v, 10)
		break
	}
	assertEqual(t, []int{reads, stops}, []int{1, 1})
}

// 指定件数に達する前の下流停止が、再評価のたびに上流へ伝わる。
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
	iterator := op.Take(3).Operator(src)
	for round := 1; round <= 2; round++ {
		for v := range iterator {
			assertEqual(t, v, 1)
			break
		}
		assertEqual(t, []int{reads, stops}, []int{round, round})
	}
}

// 同じ設定を異なる型と入力に使っても、評価時の残り件数を共有しない。
func TestTakeConfigurationReuse(t *testing.T) {
	take := op.Take(2)
	var ope op.Operator[int, int] = take.Operator
	var ope2 op.Operator2[[]int, string, []int, string] = take.Operator2
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
	first, second := ope(values), take.Operator(values)
	paired := ope2(pairs)
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
