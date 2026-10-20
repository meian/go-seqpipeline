package op_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

func TestMapNilInput(t *testing.T) {
	iterator := op.Map(func(int) int {
		t.Fatal("Map callback called")
		return 0
	})(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestMap2NilInput(t *testing.T) {
	iterator := op.Map2(func(int, string) (int, string) {
		t.Fatal("Map2 callback called")
		return 0, ""
	})(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

// 型変換後も入力の順序と重複を保ち、反復時にだけ入力と変換関数を実行する。
// 再評価と下流の途中終了でも、取得件数と変換回数が一致することを確認する。
func TestMap(t *testing.T) {
	reads, calls := 0, 0
	src := iter.Seq[int](func(yield func(int) bool) {
		for _, v := range []int{3, 1, 3} {
			reads++
			if !yield(v) {
				return
			}
		}
	})
	mapped := op.Map(func(v int) string {
		calls++
		return strconv.Itoa(v)
	})(src)
	assertEqual(t, []int{reads, calls}, []int{0, 0})
	for round := 1; round <= 2; round++ {
		assertEqual(t, slices.Collect(mapped), []string{"3", "1", "3"})
		assertEqual(t, []int{reads, calls}, []int{3 * round, 3 * round})
	}
	reads, calls = 0, 0
	mapped(func(v string) bool {
		assertEqual(t, v, "3")
		return false
	})
	assertEqual(t, []int{reads, calls}, []int{1, 1})
}

// キーと値の型変換後も組の順序とキーの重複を保ち、反復時にだけ入力と変換関数を実行する。
// 再評価と下流の途中終了でも、取得件数と変換回数が一致することを確認する。
func TestMap2(t *testing.T) {
	reads, calls := 0, 0
	src := iter.Seq2[int, string](func(yield func(int, string) bool) {
		for _, p := range []pair[int, string]{{3, "aaa"}, {1, "b"}, {3, "cc"}} {
			reads++
			if !yield(p.Key, p.Value) {
				return
			}
		}
	})
	mapped := op.Map2(func(k int, v string) (string, int) {
		calls++
		return strconv.Itoa(k), len(v)
	})(src)
	assertEqual(t, []int{reads, calls}, []int{0, 0})
	want := []pair[string, int]{{"3", 3}, {"1", 1}, {"3", 2}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(mapped), want)
		assertEqual(t, []int{reads, calls}, []int{3 * round, 3 * round})
	}
	reads, calls = 0, 0
	mapped(func(k string, v int) bool {
		assertEqual(t, pair[string, int]{k, v}, want[0])
		return false
	})
	assertEqual(t, []int{reads, calls}, []int{1, 1})
}
