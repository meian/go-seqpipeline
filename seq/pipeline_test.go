package seq_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

// TestTakeTypeInferenceAcrossTypeChanges は、型変換の前後で Take の型を推論でき、遅延評価・件数制限・再評価が成立することを確認する。
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

// TestPipelineTypeInference は、型引数を指定せずに int から string、再び int へ変換し、途中の Filter も接続できることを確認する。
func TestPipelineTypeInference(t *testing.T) {
	var pipeline seq.Seq[int] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Map(func(v string) int { return len(v) }))
	assertEqual(t, seq.ToSlice(pipeline), []int{3, 4})
}

// TestLazyEvaluationAndStopPropagation は、Seq の構築時に上流を実行せず、下流の途中終了が Map と Filter を経由して上流へ伝播することを確認する。
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

// TestTakeAfterFilterAndDownstreamStop は、Filter 後の Take が指定件数で停止し、下流が先に終了した場合も上流の読み取りを止めることを確認する。
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

// TestEmptyAndZeroValues は、Seq のゼロ値・nil・空のスライスでは Map と Filter を呼び出さず、空として評価することを確認する。
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

// TestRepeatAndSourceChanges は、パイプラインの再評価時に元のスライスの変更を反映し、Take の件数制限も評価ごとに適用することを確認する。
func TestRepeatAndSourceChanges(t *testing.T) {
	values := []int{1, 2, 3}
	pipeline := seq.FromSlice(values).Pipe(seq.Map(func(v int) int { return v * 10 })).
		Pipe(seq.Filter(func(v int) bool { return v >= 20 })).Pipe(seq.Take(2).Type)
	assertEqual(t, seq.ToSlice(pipeline), []int{20, 30})
	values[0] = 4
	assertEqual(t, seq.ToSlice(pipeline), []int{40, 20})
	assertEqual(t, seq.ToSlice(pipeline), []int{40, 20})
}

// TestSingleUseSource は、状態を消費する上流では、Take を再評価するたびに続きの要素を取得し、最後は空になることを確認する。
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

// TestExternalOperator は、独自 Operator を Apply 経由で接続し、一つの入力から複数の出力を生成した後に Take で制限できることを確認する。
func TestExternalOperator(t *testing.T) {
	var op seq.Operator[int, string] = repeatOperator{}
	got := seq.ToSlice(seq.FromSlice([]int{1, 2}).Pipe(op.Apply).Pipe(seq.Take(3).Type))
	assertEqual(t, got, []string{"1", "1", "2"})
}

// TestDirectTransformPipeline は、Map と Take.Type を Seq 変換関数へ代入して直接接続し、型推論と再評価を確認する。
func TestDirectTransformPipeline(t *testing.T) {
	var transform func(seq.Seq[int]) seq.Seq[string] = seq.Map(strconv.Itoa)
	var limit func(seq.Seq[string]) seq.Seq[string] = seq.Take(2).Type
	pipeline := seq.FromSlice([]int{1, 2, 3}).Pipe(transform).Pipe(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1", "2"})
	}
}

// TestOperatorFuncPipeline は、OperatorFunc を Map と Take に接続し、遅延評価・件数制限による停止・再評価を確認する。
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

// TestSeq2Pipeline は、Seq2 の Map2 と Filter2 を接続し、再評価と重複キーの保持、ToMap の後勝ちを確認する。
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

// TestSeq2LazyAndStop は、Seq2 の構築時に上流を実行せず、途中終了が上流へ伝播し、再評価でも同じ動作になることを確認する。
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

// TestSeq2EmptyAndZero は、Seq2 のゼロ値・nil・空の map では Map2 と Filter2 を呼び出さず、空として評価することを確認する。
func TestSeq2EmptyAndZero(t *testing.T) {
	var zero seq.Seq2[string, int]
	for _, src := range []seq.Seq2[string, int]{zero, seq.FromSeq2[string, int](nil), seq.FromMap(map[string]int{}), seq.FromMap(map[string]int(nil))} {
		pipeline := src.Pipe(seq.Map2(func(string, int) (string, int) { t.Fatal("Map2 called"); return "", 0 })).
			Pipe(seq.Filter2(func(string, int) bool { t.Fatal("Filter2 called"); return true }))
		got := collectPairs(seq.ToSeq2(pipeline))
		assertEqual(t, got, []pair[string, int](nil))
	}
}

// TestNonComparableKeys は、比較できないキーを持つ Seq2 も、Map2 と標準の iter.Seq2 への変換で扱えることを確認する。
func TestNonComparableKeys(t *testing.T) {
	src := func(yield func([]int, string) bool) { yield([]int{1, 2}, "value") }
	got := collectPairs(seq.ToSeq2(seq.FromSeq2(src).Pipe(seq.Map2(func(k []int, v string) ([]int, int) { return k, len(v) }))))
	assertEqual(t, got, []pair[[]int, int]{{Key: []int{1, 2}, Value: 5}})
}

// TestExternalOperator2AndFunc は、独自 Operator2 と関数型アダプターを接続し、遅延評価・途中終了・再評価が成立することを確認する。
func TestExternalOperator2AndFunc(t *testing.T) {
	op := seq.Operator2Func[string, int, string, string](func(src seq.Seq2[string, int]) seq.Seq2[string, string] {
		return func(yield func(string, string) bool) {
			for k, v := range src {
				if !yield(k, strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	var _ seq.Operator2[string, int, string, string] = op
	assertEqual(t, collectPairs(seq.ToSeq2(op.Apply(seq.FromMap(map[string]int{"1": 10})))), []pair[string, string]{{Key: "1", Value: "10"}})
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

// TestMap2Filter2Apply は、Map2 と Filter2 の Apply を直接接続し、変換後の型で選別できることを確認する。
func TestMap2Filter2Apply(t *testing.T) {
	src := func(yield func(int, int) bool) {
		if !yield(1, 10) {
			return
		}
		yield(2, 20)
	}
	mapped := seq.Map2(func(k, v int) (string, int) { return strconv.Itoa(k), v }).Apply(src)
	filtered := seq.Filter2(func(k string, v int) bool { return v > 10 }).Apply(mapped)
	assertEqual(t, collectPairs(seq.ToSeq2(filtered)), []pair[string, int]{{Key: "2", Value: 20}})
}

// TestDirectTransform2Pipeline は、独自 Operator2 の Apply と Filter2、Take2.Type を Seq2 変換関数として接続し、型推論と再評価を確認する。
func TestDirectTransform2Pipeline(t *testing.T) {
	var op seq.Operator2[int, int, string, int] = keyTextOperator{}
	var transform func(seq.Seq2[string, int]) seq.Seq2[string, int] = seq.Filter2(func(k string, v int) bool {
		return v > 1
	})
	var limit func(seq.Seq2[string, int]) seq.Seq2[string, int] = seq.Take2(1).Type
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3})).Pipe(op.Apply).Pipe(transform).Pipe(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), []pair[string, int]{{Key: "1", Value: 2}})
	}
}

// TestTake2TypeInferenceAcrossTypeChanges は、複数の型変換の前後で Take2 の型を推論でき、遅延評価・件数制限・再評価・途中終了が成立することを確認する。
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
