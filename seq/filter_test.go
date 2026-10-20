package seq_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFilterNilInput(t *testing.T) {
	iterator := seq.Filter(func(int) bool {
		t.Fatal("Filter callback called")
		return true
	})(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestFilter2NilInput(t *testing.T) {
	iterator := seq.Filter2(func(int, string) bool {
		t.Fatal("Filter2 callback called")
		return true
	})(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

// TestFilter は条件ごとの選別・順序・重複・遅延評価・再評価・停止伝播を単独で検証する。
func TestFilter(t *testing.T) {
	for _, tc := range []struct {
		name      string
		matches   func(int) bool
		stopReads int
		want      []int
	}{
		{name: "some", matches: func(v int) bool { return v > 1 }, want: []int{3, 3}, stopReads: 2},
		{name: "all", matches: func(int) bool { return true }, want: []int{1, 3, 3}, stopReads: 1},
		{name: "none", matches: func(int) bool { return false }, stopReads: 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads, calls := 0, 0
			src := iter.Seq[int](func(yield func(int) bool) {
				for _, v := range []int{1, 3, 3} {
					reads++
					if !yield(v) {
						return
					}
				}
			})
			filtered := seq.Filter(func(v int) bool {
				calls++
				return tc.matches(v)
			})(src)
			assertEqual(t, []int{reads, calls}, []int{0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, slices.Collect(filtered), tc.want)
				assertEqual(t, []int{reads, calls}, []int{3 * round, 3 * round})
			}
			reads, calls = 0, 0
			filtered(func(v int) bool {
				assertEqual(t, v, tc.want[0])
				return false
			})
			assertEqual(t, []int{reads, calls}, []int{tc.stopReads, tc.stopReads})
		})
	}
}

// TestFilter2 は両引数による選別・順序・重複・遅延評価・再評価・停止伝播を単独で検証する。
func TestFilter2(t *testing.T) {
	for _, tc := range []struct {
		name      string
		matches   func(int, string) bool
		stopReads int
		want      []pair[int, string]
	}{
		{
			name:      "some",
			matches:   func(k int, v string) bool { return k > 1 && len(v) > 1 },
			want:      []pair[int, string]{{3, "cc"}, {3, "cc"}},
			stopReads: 3,
		},
		{
			name:      "all",
			matches:   func(int, string) bool { return true },
			want:      []pair[int, string]{{1, "aa"}, {3, "b"}, {3, "cc"}, {3, "cc"}},
			stopReads: 1,
		},
		{name: "none", matches: func(int, string) bool { return false }, stopReads: 4},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads, calls := 0, 0
			src := iter.Seq2[int, string](func(yield func(int, string) bool) {
				for _, p := range []pair[int, string]{{1, "aa"}, {3, "b"}, {3, "cc"}, {3, "cc"}} {
					reads++
					if !yield(p.Key, p.Value) {
						return
					}
				}
			})
			filtered := seq.Filter2(func(k int, v string) bool {
				calls++
				return tc.matches(k, v)
			})(src)
			assertEqual(t, []int{reads, calls}, []int{0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, collectPairs(filtered), tc.want)
				assertEqual(t, []int{reads, calls}, []int{4 * round, 4 * round})
			}
			reads, calls = 0, 0
			filtered(func(k int, v string) bool {
				assertEqual(t, pair[int, string]{k, v}, tc.want[0])
				return false
			})
			assertEqual(t, []int{reads, calls}, []int{tc.stopReads, tc.stopReads})
		})
	}
}
