package seq_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

// 同じ合成を別入力へ接続し、遅延評価・停止伝播と反復ごとの Take の状態を確認する。
func TestComposeReuse(t *testing.T) {
	reads := 0
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { reads++; return v })).
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(op.Map(strconv.Itoa)).
		Pipe(op.Take(2).Operator)
	first := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform.Operator)
	second := seq.FromSlice([]int{8, 9, 10, 12}).Pipe(transform.Operator)
	assertEqual(t, reads, 0)
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(first), []string{"2", "4"})
		assertEqual(t, seq.ToSlice(second), []string{"8", "10"})
		assertEqual(t, reads, round*7)
	}
	reads = 0
	for range first {
		break
	}
	assertEqual(t, reads, 2)
	assertEqual(t, seq.ToSlice(first), []string{"2", "4"})
	assertEqual(t, reads, 6)
}

// 組の合成も入力ごと・反復ごとに件数制限を独立させ、キーを保持して停止する。
func TestCompose2Reuse(t *testing.T) {
	reads := 0
	transform := op.Compose2().
		Pipe2(op.Map2(func(k, v int) (int, int) { reads++; return k, v })).
		Pipe2(op.Filter2(func(k, v int) bool { return v%2 == 0 })).
		Pipe2(op.Map2(func(k, v int) (string, string) { return strconv.Itoa(k), strconv.Itoa(v) })).
		Pipe2(op.Take(2).Operator2)
	first := seq.FromSeq2(slices.All([]int{1, 2, 3, 4, 6})).Pipe2(transform.Operator)
	second := seq.FromSeq2(slices.All([]int{8, 9, 10, 12})).Pipe2(transform.Operator)
	assertEqual(t, reads, 0)
	firstWant := []pair[string, string]{{Key: "1", Value: "2"}, {Key: "3", Value: "4"}}
	secondWant := []pair[string, string]{{Key: "0", Value: "8"}, {Key: "2", Value: "10"}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(first)), firstWant)
		assertEqual(t, collectPairs(seq.ToSeq2(second)), secondWant)
		assertEqual(t, reads, round*7)
	}
	reads = 0
	for range first {
		break
	}
	assertEqual(t, reads, 2)
}

// 単一値と組を往復しても、合成全体の入力型と最終出力型で seq に接続できる。
func TestComposeAcrossShapes(t *testing.T) {
	fromValues := op.Compose().
		Pipe(op.Map(strconv.Itoa)).
		Pipe2(withLength[string]).
		Pipe2(op.Map2(func(k []int, v string) (string, int) { return v, k[0] })).
		Pipe(op.Values).
		Pipe(op.Map(strconv.Itoa))
	fromPairs := op.Compose2().
		Pipe(op.Values[int, string]).
		Pipe2(withLength[string]).
		Pipe2(op.Map2(func(k []int, v string) (string, int) { return v, k[0] }))
	assertEqual(t, seq.ToSlice(seq.FromSlice([]int{1, 22, 1}).Pipe(fromValues.Operator)), []string{"1", "2", "1"})
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "a"})).Pipe2(fromPairs.Operator)
	want := []pair[string, int]{{Key: "a", Value: 1}, {Key: "bbb", Value: 3}, {Key: "a", Value: 1}}
	assertEqual(t, collectPairs(seq.ToSeq2(pairs)), want)

	toPairs := op.Compose().Pipe2(withLength[string])
	toValues := op.Compose2().Pipe(op.Keys[[]int, string]).
		Pipe(op.Map(func(length []int) int { return length[0] }))
	pairPipeline := seq.FromSlice([]string{"a", "bbb"}).Pipe2(toPairs.Operator)
	pairWant := []pair[[]int, string]{{Key: []int{1}, Value: "a"}, {Key: []int{3}, Value: "bbb"}}
	assertEqual(t, collectPairs(seq.ToSeq2(pairPipeline)), pairWant)
	assertEqual(t, seq.ToSlice(pairPipeline.Pipe(toValues.Operator)), []int{1, 3})
}

// 合成値から分岐・入れ子を作っても、元の処理と各枝の状態を変更しない。
func TestComposeBranches(t *testing.T) {
	base := op.Compose().Pipe(op.Filter(func(v int) bool { return v%2 == 0 }))
	first := base.Pipe(op.Take(1).Operator)
	firstTwo := base.Pipe(op.Take(2).Operator)
	text := op.Compose().Pipe(first.Operator).Pipe(op.Map(strconv.Itoa))
	src := seq.FromSlice([]int{1, 2, 3, 4, 6})
	assertEqual(t, seq.ToSlice(src.Pipe(base.Operator)), []int{2, 4, 6})
	assertEqual(t, seq.ToSlice(src.Pipe(first.Operator)), []int{2})
	assertEqual(t, seq.ToSlice(src.Pipe(firstTwo.Operator)), []int{2, 4})
	assertEqual(t, seq.ToSlice(src.Pipe(text.Operator)), []string{"2"})
	other := seq.FromSlice([]int{5, 8, 10}).Pipe(text.Operator)
	assertEqual(t, seq.ToSlice(other), []string{"8"})
	assertEqual(t, seq.ToSlice(src.Pipe(firstTwo.Operator)), []int{2, 4})
}

// キャッシュは合成の構築時には作らず、適用時に入力ごとに作る。
// キャッシュより後の処理は反復ごとに実行し、元入力の置換や別入力の適用の影響を受けない。
func TestComposeCacheTiming(t *testing.T) {
	upstream, downstream := 0, 0
	transform := op.Compose().
		Pipe(op.Map(func(v int) int { upstream++; return v * 10 })).
		Pipe(op.Cache).
		Pipe(op.Skip(1).Operator).
		Pipe(op.Map(func(v int) string { downstream++; return strconv.Itoa(v) })).
		Pipe(op.Take(1).Operator)
	assertEqual(t, upstream, 0)
	values := []int{1, 2, 3}
	first := seq.FromSlice(values).Pipe(transform.Operator)
	assertEqual(t, upstream, 3)
	assertEqual(t, downstream, 0)
	second := seq.FromSlice([]int{4, 5}).Pipe(transform.Operator)
	assertEqual(t, upstream, 5)
	values[1] = 99
	for round := 1; round <= 2; round++ {
		assertEqual(t, seq.ToSlice(first), []string{"20"})
		assertEqual(t, seq.ToSlice(second), []string{"50"})
		assertEqual(t, upstream, 5)
		assertEqual(t, downstream, round*2)
	}
}

func TestCompose2CacheTiming(t *testing.T) {
	upstream, downstream := 0, 0
	transform := op.Compose2().
		Pipe2(op.Map2(func(k int, v string) (int, int) { upstream++; return k, len(v) })).
		Pipe2(op.Cache2).
		Pipe2(op.Skip(1).Operator2).
		Pipe2(op.Map2(func(k, v int) (string, int) { downstream++; return strconv.Itoa(k), v })).
		Pipe2(op.Take(1).Operator2)
	assertEqual(t, upstream, 0)
	values := []string{"a", "bbb", "cc"}
	first := seq.FromSeq2(slices.All(values)).Pipe2(transform.Operator)
	assertEqual(t, upstream, 3)
	assertEqual(t, downstream, 0)
	second := seq.FromSeq2(slices.All([]string{"d", "ee"})).Pipe2(transform.Operator)
	assertEqual(t, upstream, 5)
	values[1] = "changed"
	firstWant := []pair[string, int]{{Key: "1", Value: 3}}
	secondWant := []pair[string, int]{{Key: "1", Value: 2}}
	for round := 1; round <= 2; round++ {
		assertEqual(t, collectPairs(seq.ToSeq2(first)), firstWant)
		assertEqual(t, collectPairs(seq.ToSeq2(second)), secondWant)
		assertEqual(t, upstream, 5)
		assertEqual(t, downstream, round*2)
	}
}

func TestComposeTakeZero(t *testing.T) {
	values := seq.Seq[int](func(func(int) bool) { t.Fatal("unexpected value evaluation") })
	pairs := seq.Seq2[int, string](func(func(int, string) bool) { t.Fatal("unexpected pair evaluation") })
	fromValues := op.Compose().Pipe[int](op.Take(0).Operator).Pipe(op.Cache)
	fromPairs := op.Compose2().Pipe2[int, string](op.Take(0).Operator2).Pipe2(op.Cache2)
	assertEqual(t, seq.ToSlice(values.Pipe(fromValues.Operator)), []int(nil))
	assertEqual(t, collectPairs(seq.ToSeq2(pairs.Pipe2(fromPairs.Operator))), []pair[int, string](nil))
}

// nil と空入力を往復する合成に通し、組み込み処理が非 nil の空イテレーターを返す契約を確認する。
func TestComposeEmptyInputs(t *testing.T) {
	fromValues := op.Compose().Pipe2[string](withLength[string]).Pipe(op.Values)
	fromPairs := op.Compose2().Pipe(op.Values[int, string]).Pipe2(withLength[string])
	for _, src := range []iter.Seq[string]{nil, slices.Values([]string{})} {
		result := seq.FromSeq(src).Pipe(fromValues.Operator)
		if result == nil {
			t.Fatal("expected non-nil value iterator")
		}
		assertEqual(t, seq.ToSlice(result), []string(nil))
	}
	for _, src := range []iter.Seq2[int, string]{nil, slices.All([]string{})} {
		result := seq.FromSeq2(src).Pipe2(fromPairs.Operator)
		if result == nil {
			t.Fatal("expected non-nil pair iterator")
		}
		assertEqual(t, collectPairs(seq.ToSeq2(result)), []pair[[]int, string](nil))
	}
}

// Zip の Pull を合成・再利用しても、停止と panic のどちらでも両入力を解放する。
func TestComposeZipRelease(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
		for _, panics := range []bool{false, true} {
			name := "alignment=" + strconv.Itoa(int(alignment)) + "/panic=" + strconv.FormatBool(panics)
			t.Run(name, func(t *testing.T) {
				leftStarts, leftStops, rightStarts, rightStops := 0, 0, 0, 0
				left := iter.Seq[int](func(yield func(int) bool) {
					leftStarts++
					defer func() { leftStops++ }()
					for v := 0; v < 3; v++ {
						if !yield(v) {
							return
						}
					}
				})
				right := iter.Seq[string](func(yield func(string) bool) {
					rightStarts++
					defer func() { rightStops++ }()
					for _, v := range []string{"a", "b", "c"} {
						if !yield(v) {
							return
						}
					}
				})
				transform := op.Compose().
					Pipe2[int](op.Zip(right).AlignBy(alignment).Operator).
					Pipe(op.Values)
				pipeline := seq.FromSeq(left).Pipe(transform.Operator)
				assertEqual(t, leftStarts, 0)
				assertEqual(t, rightStarts, 0)
				for round := 1; round <= 2; round++ {
					func() {
						defer func() {
							got := recover()
							if panics {
								assertEqual(t, got, any("stop"))
							} else if got != nil {
								t.Fatalf("unexpected panic: %v", got)
							}
						}()
						for range pipeline {
							if panics {
								panic("stop")
							}
							break
						}
					}()
					assertEqual(t, leftStarts, round)
					assertEqual(t, leftStops, round)
					assertEqual(t, rightStarts, round)
					assertEqual(t, rightStops, round)
				}
			})
		}
	}
}
