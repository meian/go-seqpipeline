package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestTakeTypeInferenceAcrossTypeChanges(t *testing.T) {
	reads := 0
	var pipeline seq.Seq[string] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Take(10).Type).
		Pipe(seq.Map(func(v int) string { reads++; return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 1 })).
		Pipe(seq.Take(2).Type)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"22", "333"})
		assertEqual(t, reads, round*3)
	}
}

func TestSeq2Pipeline(t *testing.T) {
	src := func(yield func(int, string) bool) {
		for _, p := range []pair[int, string]{{Key: 1, Value: "a"}, {Key: 1, Value: "bbb"}, {Key: 2, Value: "cc"}} {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
	var pipeline seq.Seq2[string, int] = seq.FromSeq2(src).
		Pipe(seq.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })).
		Pipe(seq.Filter2(func(k string, v int) bool { return k == "1" && v > 0 }))
	want := []pair[string, int]{{Key: "1", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	assertEqual(t, seq.ToMap(pipeline), map[string]int{"1": 3})
}

func TestSeq2LazyAndStop(t *testing.T) {
	starts, reads, maps, filters, stops := 0, 0, 0, 0, 0
	src := func(yield func(int, int) bool) {
		starts++
		for i := 1; i <= 10; i++ {
			reads++
			if !yield(i, i*10) {
				stops++
				return
			}
		}
	}
	pipeline := seq.FromSeq2(src).
		Pipe(seq.Map2(func(k, v int) (string, int) { maps++; return strconv.Itoa(k), v })).
		Pipe(seq.Filter2(func(k string, v int) bool { filters++; return v%20 == 0 }))
	iterator := seq.ToSeq2(pipeline)
	assertEqual(t, []int{starts, reads, maps, filters, stops}, []int{0, 0, 0, 0, 0})
	for round := 1; round <= 2; round++ {
		for k, v := range iterator {
			assertEqual(t, k, "2")
			assertEqual(t, v, 20)
			break
		}
		assertEqual(t, []int{starts, reads, maps, filters, stops}, []int{round, 2 * round, 2 * round, 2 * round, round})
	}
}

func TestSeq2EmptyAndZero(t *testing.T) {
	var zero seq.Seq2[string, int]
	for _, src := range []seq.Seq2[string, int]{zero, seq.FromSeq2[string, int](nil), seq.FromMap(map[string]int{}), seq.FromMap(map[string]int(nil))} {
		pipeline := src.Pipe(seq.Map2(func(string, int) (string, int) { t.Fatal("Map2 called"); return "", 0 })).
			Pipe(seq.Filter2(func(string, int) bool { t.Fatal("Filter2 called"); return true }))
		got := collectPairs(seq.ToSeq2(pipeline))
		assertEqual(t, got, []pair[string, int](nil))
	}
}

func TestNonComparableKeys(t *testing.T) {
	src := func(yield func([]int, string) bool) { yield([]int{1, 2}, "value") }
	got := collectPairs(seq.ToSeq2(seq.FromSeq2(src).Pipe(seq.Map2(func(k []int, v string) ([]int, int) { return k, len(v) }))))
	assertEqual(t, got, []pair[[]int, int]{{Key: []int{1, 2}, Value: 5}})
}

func TestExternalOperator2AndFunc(t *testing.T) {
	op := seq.Operator2Func[string, int, string, string](func(src iter.Seq2[string, int]) iter.Seq2[string, string] {
		return func(yield func(string, string) bool) {
			for k, v := range src {
				if !yield(k, strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	var _ seq.Operator2[string, int, string, string] = op
	assertEqual(t, collectPairs(op.Apply(seq.ToSeq2(seq.FromMap(map[string]int{"1": 10})))), []pair[string, string]{{Key: "1", Value: "10"}})
	reads := 0
	src := func(yield func(int, int) bool) {
		for i := 1; i <= 3; i++ {
			reads++
			if !yield(i, i*10) {
				return
			}
		}
	}
	pipeline := seq.FromSeq2(src).Pipe(seq.Operator2Func(keyTextOperator{}.Apply)).Pipe(op)
	assertEqual(t, reads, 0)
	for k, v := range seq.ToSeq2(pipeline) {
		assertEqual(t, k, "1")
		assertEqual(t, v, "10")
		break
	}
	assertEqual(t, reads, 1)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), []pair[string, string]{{Key: "1", Value: "10"}, {Key: "2", Value: "20"}, {Key: "3", Value: "30"}})
}

func TestMap2Filter2Apply(t *testing.T) {
	src := func(yield func(int, int) bool) {
		if !yield(1, 10) {
			return
		}
		yield(2, 20)
	}
	mapped := seq.Map2(func(k, v int) (string, int) { return strconv.Itoa(k), v }).Apply(src)
	filtered := seq.Filter2(func(k string, v int) bool { return v > 10 }).Apply(mapped)
	assertEqual(t, collectPairs(filtered), []pair[string, int]{{Key: "2", Value: 20}})
}

func TestPipelineTypeInference(t *testing.T) {
	// 型引数なしで int → string → int を接続できることをコンパイルで確認する。
	var pipeline seq.Seq[int] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Map(func(v string) int { return len(v) }))
	assertEqual(t, seq.ToSlice(pipeline), []int{3, 4})
}

func TestLazyEvaluationAndStopPropagation(t *testing.T) {
	starts, reads, maps, filters, stops := 0, 0, 0, 0, 0
	src := func(yield func(int) bool) {
		starts++
		for i := 1; i <= 10; i++ {
			reads++
			if !yield(i) {
				stops++
				return
			}
		}
	}
	pipeline := seq.FromSeq(src).
		Pipe(seq.Map(func(v int) int { maps++; return v * 2 })).
		Pipe(seq.Filter(func(v int) bool { filters++; return v%4 == 0 }))
	iterator := seq.ToSeq(pipeline)
	assertEqual(t, []int{starts, reads, maps, filters, stops}, []int{0, 0, 0, 0, 0})
	var got []int
	for v := range iterator {
		got = append(got, v)
		break
	}
	assertEqual(t, got, []int{4})
	assertEqual(t, []int{starts, reads, maps, filters, stops}, []int{1, 2, 2, 2, 1})
}

func TestTakeAfterFilterAndDownstreamStop(t *testing.T) {
	reads := 0
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 5, 6}).
		Pipe(seq.Map(func(v int) int { reads++; return v })).
		Pipe(seq.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(seq.Take(2).Type)
	assertEqual(t, seq.ToSlice(pipeline), []int{2, 4})
	assertEqual(t, reads, 4)
	reads = 0
	for range seq.ToSeq(pipeline) {
		break
	}
	assertEqual(t, reads, 2)
}

func TestEmptyAndZeroValues(t *testing.T) {
	var zero seq.Seq[int]
	for _, src := range []seq.Seq[int]{zero, seq.FromSlice([]int{}), seq.FromSlice([]int(nil)), seq.FromSeq[int](nil)} {
		pipeline := src.Pipe(seq.Map(func(int) int { t.Fatal("Map called"); return 0 })).
			Pipe(seq.Filter(func(int) bool { t.Fatal("Filter called"); return true })).
			Pipe(seq.Take(1).Type)
		got := seq.ToSlice(pipeline)
		assertEqual(t, got, []int(nil))
	}
}

func TestRepeatAndSourceChanges(t *testing.T) {
	values := []int{1, 2, 3}
	pipeline := seq.FromSlice(values).Pipe(seq.Map(func(v int) int { return v * 10 })).
		Pipe(seq.Filter(func(v int) bool { return v >= 20 })).Pipe(seq.Take(2).Type)
	assertEqual(t, seq.ToSlice(pipeline), []int{20, 30})
	values[0] = 4
	assertEqual(t, seq.ToSlice(pipeline), []int{40, 20})
	assertEqual(t, seq.ToSlice(pipeline), []int{40, 20})
}

func TestSingleUseSource(t *testing.T) {
	remaining := []int{1, 2, 3}
	src := func(yield func(int) bool) {
		for len(remaining) > 0 {
			v := remaining[0]
			remaining = remaining[1:]
			if !yield(v) {
				return
			}
		}
	}
	pipeline := seq.FromSeq(src).Pipe(seq.Take(1).Type)
	assertEqual(t, seq.ToSlice(pipeline), []int{1})
	assertEqual(t, seq.ToSlice(pipeline), []int{2})
	assertEqual(t, seq.ToSlice(pipeline), []int{3})
	assertEqual(t, seq.ToSlice(pipeline), []int(nil))
}

func TestExternalOperator(t *testing.T) {
	var op seq.Operator[int, string] = repeatOperator{}
	got := seq.ToSlice(seq.FromSlice([]int{1, 2}).Pipe(op.Apply).Pipe(seq.Take(3).Type))
	assertEqual(t, got, []string{"1", "1", "2"})
}

func TestDirectTransformPipeline(t *testing.T) {
	var transform func(iter.Seq[int]) iter.Seq[string] = seq.Map(strconv.Itoa)
	var limit func(iter.Seq[string]) iter.Seq[string] = seq.Take(2).Type
	pipeline := seq.FromSlice([]int{1, 2, 3}).Pipe(transform).Pipe(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1", "2"})
	}
}

func TestDirectTransform2Pipeline(t *testing.T) {
	var op seq.Operator2[int, int, string, int] = keyTextOperator{}
	var transform func(iter.Seq2[string, int]) iter.Seq2[string, int] = seq.Filter2(func(k string, v int) bool {
		return v > 1
	})
	var limit func(iter.Seq2[string, int]) iter.Seq2[string, int] = seq.Take2(1).Type
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3})).Pipe(op.Apply).Pipe(transform).Pipe(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), []pair[string, int]{{Key: "1", Value: 2}})
	}
}

func TestOperatorFuncPipeline(t *testing.T) {
	op := seq.OperatorFunc(repeatOperator{}.Apply)
	calls := 0
	pipeline := seq.FromSlice([]int{1, 2, 3}).
		Pipe(seq.Map(func(v int) int { calls++; return v })).Pipe(op).Pipe(seq.Take(1).Type)
	assertEqual(t, calls, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1"})
		assertEqual(t, calls, round)
	}
}

func TestTake2TypeInferenceAcrossTypeChanges(t *testing.T) {
	reads := 0
	var pipeline seq.Seq2[string, int] = seq.FromSeq2(slices.All([]int{1, 22, 333, 4444})).
		Pipe(seq.Take2(10).Type).
		Pipe(seq.Map2(func(k, v int) ([]int, string) {
			reads++
			return []int{k}, strconv.Itoa(v)
		})).
		Pipe(seq.Filter2(func(k []int, v string) bool { return len(v) > 1 })).
		Pipe(seq.Take2(2).Type).
		Pipe(seq.Map2(func(k []int, v string) (string, int) { return strconv.Itoa(k[0]), len(v) }))
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToMap(pipeline), map[string]int{"1": 2, "2": 3})
		assertEqual(t, reads, round*3)
	}
	for range seq.ToSeq2(pipeline) {
		break
	}
	assertEqual(t, reads, 8)
}
