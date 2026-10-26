package op_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// 型変換後も入力の順序と重複を保ち、反復時にだけ入力と変換関数を実行する。
// 再評価と下流の途中終了でも、取得件数と変換回数が一致することを確認する。
func TestMap(t *testing.T) {
	t.Parallel()
	var state inputTrace
	calls := 0
	src := traceValues([]int{3, 1, 3}, &state)
	mapped := op.Map(func(v int) string {
		calls++
		return strconv.Itoa(v)
	})(src)
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	for i := 1; i <= 2; i++ {
		assertEqual(t, slices.Collect(mapped), []string{"3", "1", "3"})
		assertEqual(t, []int{state.reads, calls}, []int{3 * i, 3 * i})
	}
	state.reads, calls = 0, 0
	mapped(func(v string) bool {
		assertEqual(t, v, "3")
		return false
	})
	assertEqual(t, []int{state.reads, calls}, []int{1, 1})
}

func TestMapNilInput(t *testing.T) {
	t.Parallel()
	iterator := op.Map(func(int) int {
		t.Fatal("Map callback called")
		return 0
	})(nil)
	for range iterator {
		t.Fatal("nil input yielded a value")
	}
}

// キーと値の型変換後も組の順序とキーの重複を保ち、反復時にだけ入力と変換関数を実行する。
// 再評価と下流の途中終了でも、取得件数と変換回数が一致することを確認する。
func TestMap2(t *testing.T) {
	t.Parallel()
	var state inputTrace
	calls := 0
	src := tracePairs([]pair[int, string]{{3, "aaa"}, {1, "b"}, {3, "cc"}}, &state)
	mapped := op.Map2(func(k int, v string) (string, int) {
		calls++
		return strconv.Itoa(k), len(v)
	})(src)
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	want := []pair[string, int]{{"3", 3}, {"1", 1}, {"3", 2}}
	for i := 1; i <= 2; i++ {
		assertEqual(t, collectPairs(mapped), want)
		assertEqual(t, []int{state.reads, calls}, []int{3 * i, 3 * i})
	}
	state.reads, calls = 0, 0
	mapped(func(k string, v int) bool {
		assertEqual(t, pair[string, int]{k, v}, want[0])
		return false
	})
	assertEqual(t, []int{state.reads, calls}, []int{1, 1})
}

func TestMap2NilInput(t *testing.T) {
	t.Parallel()
	iterator := op.Map2(func(int, string) (int, string) {
		t.Fatal("Map2 callback called")
		return 0, ""
	})(nil)
	for range iterator {
		t.Fatal("nil input yielded a value")
	}
}

// 型変換と比較不能なキー・値を扱い、順序・重複と評価回数を確認する。
// 再評価後に途中終了しても、余分な入力やコールバックを実行しない。
func TestMapToSeq2(t *testing.T) {
	t.Parallel()
	var state inputTrace
	calls := 0
	src := traceValues([]int{3, 1, 3}, &state)
	var ope op.ToSeq2Operator[int, []int, []string] = op.MapToSeq2(func(v int) ([]int, []string) {
		calls++
		return []int{v}, []string{strconv.Itoa(v)}
	})
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	mapped := ope(src)
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	want := []pair[[]int, []string]{
		{Key: []int{3}, Value: []string{"3"}},
		{Key: []int{1}, Value: []string{"1"}},
		{Key: []int{3}, Value: []string{"3"}},
	}
	for i := 1; i <= 2; i++ {
		assertEqual(t, collectPairs(mapped), want)
		assertEqual(t, []int{state.reads, calls}, []int{3 * i, 3 * i})
	}
	state.reads, calls = 0, 0
	for k, v := range mapped {
		assertEqual(t, pair[[]int, []string]{Key: k, Value: v}, want[0])
		break
	}
	assertEqual(t, []int{state.reads, calls}, []int{1, 1})
	assertEqual(t, collectPairs(mapped), want)
	assertEqual(t, []int{state.reads, calls}, []int{4, 4})
}

func TestMapToSeq2EmptyInput(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		src  iter.Seq[int]
	}{
		{name: "nil"},
		{name: "empty", src: slices.Values([]int{})},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mapped := op.MapToSeq2(func(int) (string, int) {
				t.Fatal("MapToSeq2 callback called")
				return "", 0
			})(tc.src)
			if mapped == nil {
				t.Fatal("expected non-nil iterator")
			}
			assertEqual(t, collectPairs(mapped), []pair[string, int](nil))
		})
	}
}

// 比較不能なキー・値を単一値へ変換し、順序・重複と評価回数を確認する。
// 再評価後に途中終了しても、余分な入力やコールバックを実行しない。
func TestMapToSeq(t *testing.T) {
	t.Parallel()
	var state inputTrace
	calls := 0
	src := tracePairs([]pair[[]int, []string]{
		{[]int{3}, []string{"3"}}, {[]int{1}, []string{"1"}}, {[]int{3}, []string{"3"}},
	}, &state)
	var ope op.ToSeqOperator[[]int, []string, string] = op.MapToSeq(func(k []int, v []string) string {
		calls++
		return strconv.Itoa(k[0]) + ":" + v[0]
	})
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	mapped := ope(src)
	assertEqual(t, []int{state.reads, calls}, []int{0, 0})
	want := []string{"3:3", "1:1", "3:3"}
	for i := 1; i <= 2; i++ {
		assertEqual(t, slices.Collect(mapped), want)
		assertEqual(t, []int{state.reads, calls}, []int{3 * i, 3 * i})
	}
	state.reads, calls = 0, 0
	for v := range mapped {
		assertEqual(t, v, want[0])
		break
	}
	assertEqual(t, []int{state.reads, calls}, []int{1, 1})
	assertEqual(t, slices.Collect(mapped), want)
	assertEqual(t, []int{state.reads, calls}, []int{4, 4})
}

func TestMapToSeqEmptyInput(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name string
		src  iter.Seq2[int, string]
	}{
		{name: "nil"},
		{name: "empty", src: slices.All([]string{})},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mapped := op.MapToSeq(func(int, string) string {
				t.Fatal("MapToSeq callback called")
				return ""
			})(tc.src)
			if mapped == nil {
				t.Fatal("expected non-nil iterator")
			}
			assertEqual(t, slices.Collect(mapped), []string(nil))
		})
	}
}
