package seq_test

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFilter(t *testing.T) {
	op := seq.Filter(func(v int) bool { return v%2 == 0 })
	assertEqual(t, slices.Collect(op.Apply(slices.Values([]int{3, 2, 4, 1}))), []int{2, 4})
}

func TestFilter2(t *testing.T) {
	op := seq.Filter2(func(k int, v string) bool { return k > 0 && len(v) > 1 })
	got := collectPairs(op.Apply(slices.All([]string{"a", "bbb", "c"})))
	assertEqual(t, got, []pair[int, string]{{Key: 1, Value: "bbb"}})
}

func TestFilterNilInput(t *testing.T) {
	iterator := seq.Filter(func(int) bool { t.Fatal("Filter callback called"); return true }).Apply(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestFilter2NilInput(t *testing.T) {
	iterator := seq.Filter2(func(int, string) bool { t.Fatal("Filter2 callback called"); return true }).Apply(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}
