package seq_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestMap(t *testing.T) {
	op := seq.Map(func(v int) string { return strconv.Itoa(v) })
	assertEqual(t, slices.Collect(op.Apply(slices.Values([]int{2, 1}))), []string{"2", "1"})
}

func TestMap2(t *testing.T) {
	op := seq.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	got := collectPairs(op.Apply(slices.All([]string{"a", "bbb"})))
	assertEqual(t, got, []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}})
}

func TestMapNilInput(t *testing.T) {
	iterator := seq.Map(func(int) int { t.Fatal("Map callback called"); return 0 }).Apply(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestMap2NilInput(t *testing.T) {
	iterator := seq.Map2(func(int, string) (int, string) { t.Fatal("Map2 callback called"); return 0, "" }).Apply(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}
