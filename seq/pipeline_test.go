package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

// TestTakeTypeInferenceAcrossTypeChanges は、型変換の前後で Take の型を推論でき、遅延評価・件数制限・再評価が成立することを確認する。
func TestTakeTypeInferenceAcrossTypeChanges(t *testing.T) {
	reads := 0
	var pipeline seq.Seq[string] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Take(10).Operator).
		Pipe(seq.Map(func(v int) string {
			reads++
			return strconv.Itoa(v)
		})).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 1 })).
		Pipe(seq.Take(2).Operator)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"22", "333"})
		assertEqual(t, reads, round*3)
	}
}

// TestLazyEvaluationAndStopPropagation は、seq.Seq の構築時に上流を実行せず、下流の途中終了が Map と Filter を経由して上流へ伝播することを確認する。
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
		Pipe(seq.Map(func(v int) int {
			maps++
			return v * 2
		})).
		Pipe(seq.Filter(func(v int) bool {
			filters++
			return v%4 == 0
		}))
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
		Pipe(seq.Map(func(v int) int {
			reads++
			return v
		})).
		Pipe(seq.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(seq.Take(2).Operator)
	assertEqual(t, seq.ToSlice(pipeline), []int{2, 4})
	assertEqual(t, reads, 4)
	reads = 0
	for range seq.ToSeq(pipeline) {
		break
	}
	assertEqual(t, reads, 2)
}

// TestEmptyAndZeroValues は、seq.Seq のゼロ値・nil・空のスライスでは Map と Filter を呼び出さず、空として評価することを確認する。
func TestEmptyAndZeroValues(t *testing.T) {
	var zero seq.Seq[int]
	for _, src := range []seq.Seq[int]{
		zero,
		seq.FromSlice([]int{}),
		seq.FromSlice([]int(nil)),
		seq.FromSeq[int](nil),
	} {
		pipeline := src.Pipe(seq.Map(func(int) int {
			t.Fatal("Map called")
			return 0
		})).
			Pipe(seq.Filter(func(int) bool {
				t.Fatal("Filter called")
				return true
			})).
			Pipe(seq.Take(1).Operator)
		got := seq.ToSlice(pipeline)
		assertEqual(t, got, []int(nil))
	}
}

// TestRepeatAndSourceChanges は、パイプラインの再評価時に元のスライスの変更を反映し、Take の件数制限も評価ごとに適用することを確認する。
func TestRepeatAndSourceChanges(t *testing.T) {
	values := []int{1, 2, 3}
	pipeline := seq.FromSlice(values).Pipe(seq.Map(func(v int) int { return v * 10 })).
		Pipe(seq.Filter(func(v int) bool { return v >= 20 })).Pipe(seq.Take(2).Operator)
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
	pipeline := seq.FromSeq(src).Pipe(seq.Take(1).Operator)
	assertEqual(t, seq.ToSlice(pipeline), []int{1})
	assertEqual(t, seq.ToSlice(pipeline), []int{2})
	assertEqual(t, seq.ToSlice(pipeline), []int{3})
	assertEqual(t, seq.ToSlice(pipeline), []int(nil))
}

// TestDirectTransformPipeline は、Map と Take.Operator を Operator へ代入して直接接続し、型推論と再評価を確認する。
func TestDirectTransformPipeline(t *testing.T) {
	var transform seq.Operator[int, string] = seq.Map(strconv.Itoa)
	var limit seq.Operator[string, string] = seq.Take(2).Operator
	pipeline := seq.FromSlice([]int{1, 2, 3}).Pipe(transform).Pipe(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1", "2"})
	}
}

// TestTransformPipeline は、独自の変換関数を Map と Take に接続し、遅延評価・件数制限による停止・再評価を確認する。
func TestTransformPipeline(t *testing.T) {
	op := repeat
	calls := 0
	pipeline := seq.FromSlice([]int{1, 2, 3}).
		Pipe(seq.Map(func(v int) int {
			calls++
			return v
		})).Pipe(op).Pipe(seq.Take(1).Operator)
	assertEqual(t, calls, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1"})
		assertEqual(t, calls, round)
	}
}

// TestSeq2Pipeline は、seq.Seq2 の Map2 と Filter2 を接続し、再評価と重複キーの保持、ToMap の後勝ちを確認する。
func TestSeq2Pipeline(t *testing.T) {
	src := func(yield func(int, string) bool) {
		for _, p := range []pair[int, string]{{Key: 1, Value: "a"}, {Key: 1, Value: "bbb"}, {Key: 2, Value: "cc"}} {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
	var pipeline seq.Seq2[string, int] = seq.FromSeq2(src).
		Pipe2(seq.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })).
		Pipe2(seq.Filter2(func(k string, v int) bool { return k == "1" && v > 0 }))
	want := []pair[string, int]{{Key: "1", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	assertEqual(t, seq.ToMap(pipeline), map[string]int{"1": 3})
}

// TestSeq2LazyAndStop は、seq.Seq2 の構築時に上流を実行せず、途中終了が上流へ伝播し、再評価でも同じ動作になることを確認する。
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
		Pipe2(seq.Map2(func(k, v int) (string, int) {
			maps++
			return strconv.Itoa(k), v
		})).
		Pipe2(seq.Filter2(func(k string, v int) bool {
			filters++
			return v%20 == 0
		}))
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

// TestSeq2EmptyAndZero は、seq.Seq2 のゼロ値・nil・空の map では Map2 と Filter2 を呼び出さず、空として評価することを確認する。
func TestSeq2EmptyAndZero(t *testing.T) {
	var zero seq.Seq2[string, int]
	for _, src := range []seq.Seq2[string, int]{
		zero,
		seq.FromSeq2[string, int](nil),
		seq.FromMap(map[string]int{}),
		seq.FromMap(map[string]int(nil)),
	} {
		pipeline := src.Pipe2(seq.Map2(func(string, int) (string, int) {
			t.Fatal("Map2 called")
			return "", 0
		})).
			Pipe2(seq.Filter2(func(string, int) bool {
				t.Fatal("Filter2 called")
				return true
			}))
		got := collectPairs(seq.ToSeq2(pipeline))
		assertEqual(t, got, []pair[string, int](nil))
	}
}

// TestNonComparableKeys は、比較できないキーを持つ seq.Seq2 も、Map2 と iter.Seq2 への変換で扱えることを確認する。
func TestNonComparableKeys(t *testing.T) {
	src := func(yield func([]int, string) bool) { yield([]int{1, 2}, "value") }
	pipeline := seq.FromSeq2(src).Pipe2(seq.Map2(func(k []int, v string) ([]int, int) {
		return k, len(v)
	}))
	got := collectPairs(seq.ToSeq2(pipeline))
	assertEqual(t, got, []pair[[]int, int]{{Key: []int{1, 2}, Value: 5}})
}

// TestExternalTransforms2 は、独自の iter.Seq2 変換関数を接続し、遅延評価・途中終了・再評価が成立することを確認する。
func TestExternalTransforms2(t *testing.T) {
	op := func(src iter.Seq2[string, int]) iter.Seq2[string, string] {
		return func(yield func(string, string) bool) {
			for k, v := range src {
				if !yield(k, strconv.Itoa(v)) {
					return
				}
			}
		}
	}
	direct := op(seq.ToSeq2(seq.FromMap(map[string]int{"1": 10})))
	assertEqual(t, collectPairs(direct), []pair[string, string]{{Key: "1", Value: "10"}})
	reads := 0
	src := func(yield func(int, int) bool) {
		for i := 1; i <= 3; i++ {
			reads++
			if !yield(i, i*10) {
				return
			}
		}
	}
	pipeline := seq.FromSeq2(src).Pipe2(keyText).Pipe2(op)
	assertEqual(t, reads, 0)
	for k, v := range seq.ToSeq2(pipeline) {
		assertEqual(t, k, "1")
		assertEqual(t, v, "10")
		break
	}
	assertEqual(t, reads, 1)
	want := []pair[string, string]{
		{Key: "1", Value: "10"},
		{Key: "2", Value: "20"},
		{Key: "3", Value: "30"},
	}
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
}

// TestDirectTransform2Pipeline は、独自の iter.Seq2 変換関数と Filter2、Take.Operator2 を Operator2 として接続し、型推論と再評価を確認する。
func TestDirectTransform2Pipeline(t *testing.T) {
	op := keyText
	var transform seq.Operator2[string, int, string, int] = seq.Filter2(func(k string, v int) bool {
		return v > 1
	})
	var limit seq.Operator2[string, int, string, int] = seq.Take(1).Operator2
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3})).Pipe2(op).Pipe2(transform).Pipe2(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), []pair[string, int]{{Key: "1", Value: 2}})
	}
}

// TestTakeOperator2TypeInferenceAcrossTypeChanges は、複数の型変換の前後で Take の型を推論でき、遅延評価・件数制限・再評価・途中終了が成立することを確認する。
func TestTakeOperator2TypeInferenceAcrossTypeChanges(t *testing.T) {
	reads := 0
	var pipeline seq.Seq2[string, int] = seq.FromSeq2(slices.All([]int{1, 22, 333, 4444})).
		Pipe2(seq.Take(10).Operator2).
		Pipe2(seq.Map2(func(k, v int) ([]int, string) {
			reads++
			return []int{k}, strconv.Itoa(v)
		})).
		Pipe2(seq.Filter2(func(k []int, v string) bool { return len(v) > 1 })).
		Pipe2(seq.Take(2).Operator2).
		Pipe2(seq.Map2(func(k []int, v string) (string, int) { return strconv.Itoa(k[0]), len(v) }))
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

// TestCrossTypePipeline は、4通りの接続、型推論、遅延評価、停止伝播、再評価を確認する。
func TestCrossTypePipeline(t *testing.T) {
	var toPairs seq.ToSeq2Operator[string, []int, string] = withLength
	var toValues seq.ToSeqOperator[[]int, string, string] = seq.Values[[]int, string]
	for _, n := range []int{0, 1, 2} {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			starts, reads, stops := 0, 0, 0
			src := seq.Seq[int](func(yield func(int) bool) {
				starts++
				for _, v := range []int{1, 22, 333} {
					reads++
					if !yield(v) {
						stops++
						return
					}
				}
			})
			var pipeline seq.Seq[string] = src.
				Pipe(seq.Map(strconv.Itoa)).
				Pipe2(toPairs).
				Pipe2(seq.Filter2(func(k []int, v string) bool { return k[0] > 1 })).
				Pipe2(seq.Take(2).Operator2).
				Pipe(toValues).
				Pipe(seq.Take(n).Operator)
			assertEqual(t, []int{starts, reads, stops}, []int{0, 0, 0})
			want := []string{"22", "333"}[:n]
			if n == 0 {
				want = nil
			}
			for round := 1; round <= 2; round++ {
				assertEqual(t, seq.ToSlice(pipeline), want)
				if n == 0 {
					assertEqual(t, []int{starts, reads, stops}, []int{0, 0, 0})
				} else {
					assertEqual(t, []int{starts, reads, stops}, []int{round, (n + 1) * round, round})
				}
			}
			if n > 0 {
				for range pipeline {
					break
				}
				assertEqual(t, []int{starts, reads, stops}, []int{3, 2*(n+1) + 2, 3})
			}
		})
	}
}

// TestSkipAfterFilter は、選別後の要素を読み飛ばし、Take の停止が上流まで伝播することを確認する。
func TestSkipAfterFilter(t *testing.T) {
	reads := 0
	src := seq.Seq[int](func(yield func(int) bool) {
		for i := 1; i <= 10; i++ {
			reads++
			if !yield(i) {
				return
			}
		}
	})
	var pipeline seq.Seq[string] = src.
		Pipe(seq.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(seq.Skip(1).Operator).
		Pipe(seq.Map(strconv.Itoa)).
		Pipe(seq.Take(2).Operator)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"4", "6"})
		assertEqual(t, reads, round*6)
	}
}

// TestSkip2AfterFilter は、選別後の組を読み飛ばし、型変換と Take に接続できることを確認する。
func TestSkip2AfterFilter(t *testing.T) {
	reads := 0
	src := seq.Seq2[int, int](func(yield func(int, int) bool) {
		for i := 1; i <= 10; i++ {
			reads++
			if !yield(i, i*10) {
				return
			}
		}
	})
	var pipeline seq.Seq2[string, int] = src.
		Pipe2(seq.Filter2(func(k, v int) bool { return k%2 == 0 })).
		Pipe2(seq.Skip(1).Operator2).
		Pipe2(seq.Map2(func(k, v int) (string, int) { return strconv.Itoa(k), v })).
		Pipe2(seq.Take(2).Operator2)
	assertEqual(t, reads, 0)
	want := []pair[string, int]{{Key: "4", Value: 40}, {Key: "6", Value: 60}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
		assertEqual(t, reads, round*6)
	}
}

// TestTakeSharedConfigurationPipeline は、共通設定を型変換の前後と seq.Seq・seq.Seq2 間で接続し、再評価時も独立して制限することを確認する。
func TestTakeSharedConfigurationPipeline(t *testing.T) {
	take := seq.Take(2)
	values := seq.FromSlice([]int{1, 22, 333}).
		Pipe(take.Operator).
		Pipe(seq.Map(strconv.Itoa)).
		Pipe(take.Operator)
	pairs := seq.FromSeq2(slices.All([]int{1, 22, 333})).
		Pipe2(take.Operator2).
		Pipe2(seq.Map2(func(k, v int) (string, string) {
			return strconv.Itoa(k), strconv.Itoa(v)
		})).
		Pipe2(take.Operator2)
	other := seq.FromSlice([]int{4, 5, 6}).Pipe(take.Operator)
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(values), []string{"1", "22"})
		want := []pair[string, string]{{Key: "0", Value: "1"}, {Key: "1", Value: "22"}}
		assertEqual(t, collectPairs(seq.ToSeq2(pairs)), want)
		assertEqual(t, seq.ToSlice(other), []int{4, 5})
	}
}

// TestKeysPipeline は、組の変換後にキーを取り出し、単一値のパイプラインを続ける。
func TestKeysPipeline(t *testing.T) {
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(seq.Map2(func(k int, v string) (int, int) { return k + 1, len(v) }))
	pipeline := pairs.Pipe(seq.Keys).
		Pipe(seq.Filter(func(k int) bool { return k > 1 })).
		Pipe(seq.Map(strconv.Itoa)).
		Pipe(seq.Take(1).Operator)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"2"})
	}
}

// TestValuesPipeline は、組の選別後に値を取り出し、単一値のパイプラインを続ける。
func TestValuesPipeline(t *testing.T) {
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(seq.Filter2(func(k int, v string) bool { return k > 0 }))
	pipeline := pairs.Pipe(seq.Values).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 1 })).
		Pipe(seq.Map(func(v string) int { return len(v) })).
		Pipe(seq.Take(1).Operator)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []int{3})
	}
}

// TestProjectionWithSkip は、seq.Seq2 と seq.Seq の両側の Skip、射影、Take を接続して評価と停止を確認する。
func TestProjectionWithSkip(t *testing.T) {
	for _, projection := range []struct {
		name string
		op   seq.ToSeqOperator[int, int, int]
		want []int
	}{
		{name: "keys", op: seq.Keys[int, int], want: []int{3, 4}},
		{name: "values", op: seq.Values[int, int], want: []int{30, 40}},
	} {
		t.Run(projection.name, func(t *testing.T) {
			for _, n := range []int{0, 1, 2} {
				t.Run(strconv.Itoa(n), func(t *testing.T) {
					starts, reads, stops := 0, 0, 0
					src := seq.Seq2[int, int](func(yield func(int, int) bool) {
						starts++
						defer func() { stops++ }()
						for i := 1; i <= 5; i++ {
							reads++
							if !yield(i, i*10) {
								return
							}
						}
					})
					skip := seq.Skip(1)
					pipeline := src.Pipe2(skip.Operator2).
						Pipe(projection.op).
						Pipe(skip.Operator).
						Pipe(seq.Take(n).Operator)
					assertEqual(t, []int{starts, reads, stops}, []int{0, 0, 0})
					want := projection.want[:n]
					if n == 0 {
						want = nil
					}
					for round := 1; round <= 2; round++ {
						assertEqual(t, seq.ToSlice(pipeline), want)
						if n == 0 {
							assertEqual(t, []int{starts, reads, stops}, []int{0, 0, 0})
						} else {
							assertEqual(t, []int{starts, reads, stops}, []int{round, (n + 2) * round, round})
						}
					}
				})
			}
		})
	}
}

// TestCachePipeline は上流を一度だけ評価し、分岐した下流は反復ごとに実行することを確認する。
func TestCachePipeline(t *testing.T) {
	upstream, downstream := 0, 0
	var cache seq.Operator[int, int] = seq.Cache
	cached := seq.FromSlice([]int{1, 2, 3, 4}).
		Pipe(seq.Map(func(v int) int { upstream++; return v * 10 })).
		Pipe(seq.Take(3).Operator).
		Pipe(cache)
	assertEqual(t, upstream, 3)
	first := cached.Pipe(seq.Map(func(v int) string {
		downstream++
		return strconv.Itoa(v)
	})).Pipe(seq.Take(1).Operator)
	second := cached.Pipe(seq.Filter(func(v int) bool { return v >= 20 }))
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(first), []string{"10"})
		assertEqual(t, seq.ToSlice(second), []int{20, 30})
		assertEqual(t, upstream, 3)
		assertEqual(t, downstream, round)
	}
	// 型引数なしの関数値を Pipe に直接渡せることも確認する。
	assertEqual(t, seq.ToSlice(cached.Pipe(seq.Cache)), []int{10, 20, 30})
}

// TestCache2Pipeline は組のキャッシュを分岐し、上流の再計算と下流の再評価を区別する。
func TestCache2Pipeline(t *testing.T) {
	upstream, downstream := 0, 0
	var cache seq.Operator2[int, int, int, int] = seq.Cache2
	cached := seq.FromSeq2(slices.All([]int{1, 2, 3, 4})).
		Pipe2(seq.Map2(func(k, v int) (int, int) { upstream++; return k, v * 10 })).
		Pipe2(seq.Take(3).Operator2).
		Pipe2(cache)
	assertEqual(t, upstream, 3)
	first := cached.Pipe2(seq.Map2(func(k, v int) (string, int) {
		downstream++
		return strconv.Itoa(k), v
	})).Pipe2(seq.Take(1).Operator2)
	second := cached.Pipe(seq.Values).Pipe(seq.Filter(func(v int) bool { return v >= 20 }))
	want := []pair[string, int]{{Key: "0", Value: 10}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(first)), want)
		assertEqual(t, seq.ToSlice(second), []int{20, 30})
		assertEqual(t, upstream, 3)
		assertEqual(t, downstream, round)
	}
	assertEqual(t, seq.ToSlice(cached.Pipe2(seq.Cache2).Pipe(seq.Values)), []int{10, 20, 30})
}
