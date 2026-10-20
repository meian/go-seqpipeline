package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

// 型変換の前後で Take を型引数なしに接続し、遅延評価と再評価ごとの件数制限を確認する。
func TestTakeTypeInferenceAcrossTypeChanges(t *testing.T) {
	reads := 0
	var pipeline seq.Seq[string] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(op.Take(10).Operator).
		Pipe(op.Map(func(v int) string {
			reads++
			return strconv.Itoa(v)
		})).
		Pipe(op.Filter(func(v string) bool { return len(v) > 1 })).
		Pipe(op.Take(2).Operator)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"22", "333"})
		assertEqual(t, reads, round*3)
	}
}

// 構築時には入力を評価せず、下流の途中終了が Map・Filter を経由して上流へ伝わる。
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
		Pipe(op.Map(func(v int) int {
			maps++
			return v * 2
		})).
		Pipe(op.Filter(func(v int) bool {
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

// Filter 後の Take が指定件数で停止し、下流が先に終了した場合も上流の取得を止める。
func TestTakeAfterFilterAndDownstreamStop(t *testing.T) {
	reads := 0
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 5, 6}).
		Pipe(op.Map(func(v int) int {
			reads++
			return v
		})).
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(op.Take(2).Operator)
	assertEqual(t, seq.ToSlice(pipeline), []int{2, 4})
	assertEqual(t, reads, 4)
	reads = 0
	for range seq.ToSeq(pipeline) {
		break
	}
	assertEqual(t, reads, 2)
}

// Seq のゼロ値・nil・空入力では、接続した Map・Filter のコールバックを呼び出さない。
func TestEmptyAndZeroValues(t *testing.T) {
	var zero seq.Seq[int]
	for _, src := range []seq.Seq[int]{
		zero,
		seq.FromSlice([]int{}),
		seq.FromSlice([]int(nil)),
		seq.FromSeq[int](nil),
	} {
		pipeline := src.Pipe(op.Map(func(int) int {
			t.Fatal("Map called")
			return 0
		})).
			Pipe(op.Filter(func(int) bool {
				t.Fatal("Filter called")
				return true
			})).
			Pipe(op.Take(1).Operator)
		got := seq.ToSlice(pipeline)
		assertEqual(t, got, []int(nil))
	}
}

// 再評価時に元のスライスの変更を反映し、Take の件数制限も毎回適用する。
func TestRepeatAndSourceChanges(t *testing.T) {
	values := []int{1, 2, 3}
	pipeline := seq.FromSlice(values).Pipe(op.Map(func(v int) int { return v * 10 })).
		Pipe(op.Filter(func(v int) bool { return v >= 20 })).Pipe(op.Take(2).Operator)
	assertEqual(t, seq.ToSlice(pipeline), []int{20, 30})
	values[0] = 4
	assertEqual(t, seq.ToSlice(pipeline), []int{40, 20})
	assertEqual(t, seq.ToSlice(pipeline), []int{40, 20})
}

// 反復状態を外部に持つ入力は、再評価しても巻き戻さない。
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
	pipeline := seq.FromSeq(src).Pipe(op.Take(1).Operator)
	assertEqual(t, seq.ToSlice(pipeline), []int{1})
	assertEqual(t, seq.ToSlice(pipeline), []int{2})
	assertEqual(t, seq.ToSlice(pipeline), []int{3})
	assertEqual(t, seq.ToSlice(pipeline), []int(nil))
}

// 名前付き関数型に代入したオペレーターを接続し、型推論と再評価を確認する。
func TestDirectTransformPipeline(t *testing.T) {
	var ope op.Operator[int, string] = op.Map(strconv.Itoa)
	var limit op.Operator[string, string] = op.Take(2).Operator
	pipeline := seq.FromSlice([]int{1, 2, 3}).Pipe(ope).Pipe(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1", "2"})
	}
}

// 独自の変換関数を Map・Take に接続し、遅延評価と再評価ごとの上流停止を確認する。
func TestTransformPipeline(t *testing.T) {
	ope := repeat
	calls := 0
	pipeline := seq.FromSlice([]int{1, 2, 3}).
		Pipe(op.Map(func(v int) int {
			calls++
			return v
		})).Pipe(ope).Pipe(op.Take(1).Operator)
	assertEqual(t, calls, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"1"})
		assertEqual(t, calls, round)
	}
}

// 組のままなら重複キーを保ち、ToMap で収集すると後の値を採用する。
func TestSeq2Pipeline(t *testing.T) {
	src := func(yield func(int, string) bool) {
		for _, p := range []pair[int, string]{{Key: 1, Value: "a"}, {Key: 1, Value: "bbb"}, {Key: 2, Value: "cc"}} {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
	var pipeline seq.Seq2[string, int] = seq.FromSeq2(src).
		Pipe2(op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })).
		Pipe2(op.Filter2(func(k string, v int) bool { return k == "1" && v > 0 }))
	want := []pair[string, int]{{Key: "1", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
	assertEqual(t, seq.ToMap(pipeline), map[string]int{"1": 3})
}

// 組のパイプラインでも遅延評価と下流からの停止伝播が働き、再評価で状態を共有しない。
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
		Pipe2(op.Map2(func(k, v int) (string, int) {
			maps++
			return strconv.Itoa(k), v
		})).
		Pipe2(op.Filter2(func(k string, v int) bool {
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

// Seq2 のゼロ値・nil・空入力では、接続した Map2・Filter2 のコールバックを呼び出さない。
func TestSeq2EmptyAndZero(t *testing.T) {
	var zero seq.Seq2[string, int]
	for _, src := range []seq.Seq2[string, int]{
		zero,
		seq.FromSeq2[string, int](nil),
		seq.FromMap(map[string]int{}),
		seq.FromMap(map[string]int(nil)),
	} {
		pipeline := src.Pipe2(op.Map2(func(string, int) (string, int) {
			t.Fatal("Map2 called")
			return "", 0
		})).
			Pipe2(op.Filter2(func(string, int) bool {
				t.Fatal("Filter2 called")
				return true
			}))
		got := collectPairs(seq.ToSeq2(pipeline))
		assertEqual(t, got, []pair[string, int](nil))
	}
}

// スライスをキーに持つ Seq2 を変換し、標準の iter.Seq2 として取り出せる。
func TestNonComparableKeys(t *testing.T) {
	src := func(yield func([]int, string) bool) { yield([]int{1, 2}, "value") }
	pipeline := seq.FromSeq2(src).Pipe2(op.Map2(func(k []int, v string) ([]int, int) {
		return k, len(v)
	}))
	got := collectPairs(seq.ToSeq2(pipeline))
	assertEqual(t, got, []pair[[]int, int]{{Key: []int{1, 2}, Value: 5}})
}

// 独自の組変換を直接適用・Pipe2 接続でき、接続時の遅延評価と下流の停止を保つ。
func TestExternalTransforms2(t *testing.T) {
	ope := func(src iter.Seq2[string, int]) iter.Seq2[string, string] {
		return func(yield func(string, string) bool) {
			for k, v := range src {
				if !yield(k, strconv.Itoa(v)) {
					return
				}
			}
		}
	}
	direct := ope(seq.ToSeq2(seq.FromMap(map[string]int{"1": 10})))
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
	pipeline := seq.FromSeq2(src).Pipe2(keyText).Pipe2(ope)
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

// 独自の組変換・Filter2・Take を名前付き関数型で接続し、型推論と再評価を確認する。
func TestDirectTransform2Pipeline(t *testing.T) {
	ope := keyText
	var filtered op.Operator2[string, int, string, int] = op.Filter2(func(k string, v int) bool {
		return v > 1
	})
	var limit op.Operator2[string, int, string, int] = op.Take(1).Operator2
	pipeline := seq.FromSeq2(slices.All([]int{1, 2, 3})).Pipe2(ope).Pipe2(filtered).Pipe2(limit)
	for round := 0; round < 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), []pair[string, int]{{Key: "1", Value: 2}})
	}
}

// 比較不能なキーを経由する型変換でも、Take の型推論・件数制限・再評価・途中終了が働く。
func TestTakeOperator2TypeInferenceAcrossTypeChanges(t *testing.T) {
	reads := 0
	var pipeline seq.Seq2[string, int] = seq.FromSeq2(slices.All([]int{1, 22, 333, 4444})).
		Pipe2(op.Take(10).Operator2).
		Pipe2(op.Map2(func(k, v int) ([]int, string) {
			reads++
			return []int{k}, strconv.Itoa(v)
		})).
		Pipe2(op.Filter2(func(k []int, v string) bool { return len(v) > 1 })).
		Pipe2(op.Take(2).Operator2).
		Pipe2(op.Map2(func(k []int, v string) (string, int) { return strconv.Itoa(k[0]), len(v) }))
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

// 比較不能なキーを含む4通りの接続で、下流の件数制限が上流まで届くことを確かめる。
func TestCrossTypePipeline(t *testing.T) {
	var toPairs op.ToSeq2Operator[string, []int, string] = withLength
	var toValues op.ToSeqOperator[[]int, string, string] = op.Values[[]int, string]
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
				Pipe(op.Map(strconv.Itoa)).
				Pipe2(toPairs).
				Pipe2(op.Filter2(func(k []int, v string) bool { return k[0] > 1 })).
				Pipe2(op.Take(2).Operator2).
				Pipe(toValues).
				Pipe(op.Take(n).Operator)
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

// 選別後の要素を Skip で読み飛ばし、Take の件数制限が再評価ごとに上流へ伝わる。
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
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(op.Skip(1).Operator).
		Pipe(op.Map(strconv.Itoa)).
		Pipe(op.Take(2).Operator)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"4", "6"})
		assertEqual(t, reads, round*6)
	}
}

// 選別後の組を Skip で読み飛ばし、型変換と Take を通して再評価できる。
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
		Pipe2(op.Filter2(func(k, v int) bool { return k%2 == 0 })).
		Pipe2(op.Skip(1).Operator2).
		Pipe2(op.Map2(func(k, v int) (string, int) { return strconv.Itoa(k), v })).
		Pipe2(op.Take(2).Operator2)
	assertEqual(t, reads, 0)
	want := []pair[string, int]{{Key: "4", Value: 40}, {Key: "6", Value: 60}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(pipeline)), want)
		assertEqual(t, reads, round*6)
	}
}

// 共通の設定を型変換の前後や別入力に使っても、残り件数を共有しない。
func TestTakeSharedConfigurationPipeline(t *testing.T) {
	take := op.Take(2)
	values := seq.FromSlice([]int{1, 22, 333}).
		Pipe(take.Operator).
		Pipe(op.Map(strconv.Itoa)).
		Pipe(take.Operator)
	pairs := seq.FromSeq2(slices.All([]int{1, 22, 333})).
		Pipe2(take.Operator2).
		Pipe2(op.Map2(func(k, v int) (string, string) {
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

// 変換後の組からキーを取り出し、単一値の選別・変換・件数制限へ接続する。
func TestKeysPipeline(t *testing.T) {
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(op.Map2(func(k int, v string) (int, int) { return k + 1, len(v) }))
	pipeline := pairs.Pipe(op.Keys).
		Pipe(op.Filter(func(k int) bool { return k > 1 })).
		Pipe(op.Map(strconv.Itoa)).
		Pipe(op.Take(1).Operator)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []string{"2"})
	}
}

// 選別後の組から値を取り出し、単一値の選別・変換・件数制限へ接続する。
func TestValuesPipeline(t *testing.T) {
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(op.Filter2(func(k int, v string) bool { return k > 0 }))
	pipeline := pairs.Pipe(op.Values).
		Pipe(op.Filter(func(v string) bool { return len(v) > 1 })).
		Pipe(op.Map(func(v string) int { return len(v) })).
		Pipe(op.Take(1).Operator)
	for i := 0; i < 2; i++ {
		assertEqual(t, seq.ToSlice(pipeline), []int{3})
	}
}

// 射影の前後の Skip と末尾の Take を接続し、評価・停止と Take(0) の上流非評価を確認する。
func TestProjectionWithSkip(t *testing.T) {
	for _, projection := range []struct {
		name string
		ope  op.ToSeqOperator[int, int, int]
		want []int
	}{
		{name: "keys", ope: op.Keys[int, int], want: []int{3, 4}},
		{name: "values", ope: op.Values[int, int], want: []int{30, 40}},
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
					skip := op.Skip(1)
					pipeline := src.Pipe2(skip.Operator2).
						Pipe(projection.ope).
						Pipe(skip.Operator).
						Pipe(op.Take(n).Operator)
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

// キャッシュ後の分岐だけを再評価し、上流の呼び出し数が増えないことを確かめる。
func TestCachePipeline(t *testing.T) {
	upstream, downstream := 0, 0
	var cache op.Operator[int, int] = op.Cache
	cached := seq.FromSlice([]int{1, 2, 3, 4}).
		Pipe(op.Map(func(v int) int { upstream++; return v * 10 })).
		Pipe(op.Take(3).Operator).
		Pipe(cache)
	assertEqual(t, upstream, 3)
	first := cached.Pipe(op.Map(func(v int) string {
		downstream++
		return strconv.Itoa(v)
	})).Pipe(op.Take(1).Operator)
	second := cached.Pipe(op.Filter(func(v int) bool { return v >= 20 }))
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(first), []string{"10"})
		assertEqual(t, seq.ToSlice(second), []int{20, 30})
		assertEqual(t, upstream, 3)
		assertEqual(t, downstream, round)
	}
	// 型引数なしの関数値を Pipe に直接渡せることも確認する。
	assertEqual(t, seq.ToSlice(cached.Pipe(op.Cache)), []int{10, 20, 30})
}

// 組のキャッシュ後に異なる変換へ分岐し、下流を再評価しても上流を再実行しない。
func TestCache2Pipeline(t *testing.T) {
	upstream, downstream := 0, 0
	var cache op.Operator2[int, int, int, int] = op.Cache2
	cached := seq.FromSeq2(slices.All([]int{1, 2, 3, 4})).
		Pipe2(op.Map2(func(k, v int) (int, int) { upstream++; return k, v * 10 })).
		Pipe2(op.Take(3).Operator2).
		Pipe2(cache)
	assertEqual(t, upstream, 3)
	first := cached.Pipe2(op.Map2(func(k, v int) (string, int) {
		downstream++
		return strconv.Itoa(k), v
	})).Pipe2(op.Take(1).Operator2)
	second := cached.Pipe(op.Values).Pipe(op.Filter(func(v int) bool { return v >= 20 }))
	want := []pair[string, int]{{Key: "0", Value: 10}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(first)), want)
		assertEqual(t, seq.ToSlice(second), []int{20, 30})
		assertEqual(t, upstream, 3)
		assertEqual(t, downstream, round)
	}
	assertEqual(t, seq.ToSlice(cached.Pipe2(op.Cache2).Pipe(op.Values)), []int{10, 20, 30})
}
