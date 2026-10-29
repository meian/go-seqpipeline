package op_test

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"sync"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// 既定の終了条件と、同じ設定から異なるキー型を推論できることを確認する。
func TestZip(t *testing.T) {
	t.Parallel()
	left := slices.Values([]int{3, 3, 1})
	right := slices.Values([]string{"a", "b"})
	config := op.Zip(right)
	ope := config.Operator[int]
	assertEqual(t, collectPairs(ope(left)), []pair[int, string]{{3, "a"}, {3, "b"}})
	// 同じ設定から異なるキー型を推論できる。
	stringKeys := config.Operator(slices.Values([]string{"x"}))
	assertEqual(t, maps.Collect(stringKeys), map[string]string{"x": "a"})
}

// 設定コピー・再指定と入力中のゼロ値の保持を確認する。
func TestZipConfigurationCopies(t *testing.T) {
	t.Parallel()
	right := slices.Values([]string{""})
	left := slices.Values([]int{0, 2})
	base := op.Zip(right)
	longest := base.AlignBy(op.AlignLongest)
	padded := longest.WithDefaultValue("missing")
	typed := padded.WithDefaultKey(-1)
	changed := typed.AlignBy(op.AlignRight).WithDefaultValue("other").WithDefaultKey(-2)
	assertEqual(t, maps.Collect(base.Operator(left)), map[int]string{0: ""})
	assertEqual(t, maps.Collect(longest.Operator(left)), map[int]string{0: "", 2: ""})
	assertEqual(t, maps.Collect(padded.Operator(left)), map[int]string{0: "", 2: "missing"})
	assertEqual(t, maps.Collect(typed.Operator(left)), map[int]string{0: "", 2: "missing"})
	assertEqual(t, maps.Collect(changed.Operator(left)), map[int]string{0: ""})
	assertEqual(t, maps.Collect(typed.Operator(nil)), map[int]string{-1: ""})
	assertEqual(t, maps.Collect(changed.Operator(nil)), map[int]string{-2: ""})
	assertEqual(t, maps.Collect(changed.AlignBy(op.AlignLeft).Operator(left)), map[int]string{0: "", 2: "other"})
	assertEqual(t, maps.Collect(typed.WithDefaultKey(0).Operator(nil)), map[int]string{0: ""})
	assertEqual(t, maps.Collect(typed.WithDefaultValue("").Operator(left)), map[int]string{0: "", 2: ""})
	// キーを先に指定する順序でも、設定値と入力のゼロ値を保持する。
	keyFirst := base.WithDefaultKey(-1).WithDefaultValue("missing").AlignBy(op.AlignLongest)
	assertEqual(t, maps.Collect(keyFirst.Operator(left)), maps.Collect(typed.Operator(left)))
	assertEqual(t, maps.Collect(keyFirst.Operator(nil)), map[int]string{-1: ""})
}

func TestZipAlignByInvalid(t *testing.T) {
	t.Parallel()
	for _, alignment := range []op.ZipAlign{-1, op.AlignLongest + 1} {
		for _, configure := range []func(){
			func() { op.Zip(iter.Seq[string](nil)).AlignBy(alignment) },
			func() { op.Zip(iter.Seq[string](nil)).WithDefaultKey(0).AlignBy(alignment) },
		} {
			func() {
				defer func() {
					want := fmt.Sprintf("op.Zip: invalid alignment: %d", alignment)
					assertEqual(t, recover(), any(want))
				}()
				configure()
			}()
		}
	}
}

// 不等長の左右×4条件で出力・取得数・解放を検証し、順序が異なる経路はイベントも確認する。
func TestZipAlignBy(t *testing.T) {
	t.Parallel()
	first := []pair[int, int]{{1, 10}}
	leftPadded := []pair[int, int]{{1, 10}, {2, -2}, {3, -2}}
	rightPadded := []pair[int, int]{{1, 10}, {-1, 20}, {-1, 30}}
	for _, tc := range []struct {
		name    string
		align   op.ZipAlign
		lengths [2]int
		want    []pair[int, int]
		reads   [2]int
		events  []string
	}{
		{"equal", op.AlignShortest, [2]int{2, 2}, []pair[int, int]{{1, 10}, {2, 20}}, [2]int{2, 2}, nil},
		{"shortest/left", op.AlignShortest, [2]int{3, 1}, first, [2]int{2, 1},
			[]string{"L1", "R10", "out:1,10", "L2", "R:end"}},
		{"shortest/right", op.AlignShortest, [2]int{1, 3}, first, [2]int{1, 1}, nil},
		{"left/left", op.AlignLeft, [2]int{3, 1}, leftPadded, [2]int{3, 1},
			[]string{"L1", "R10", "out:1,10", "L2", "R:end", "out:2,-2", "L3", "out:3,-2", "L:end"}},
		{"left/right", op.AlignLeft, [2]int{1, 3}, first, [2]int{1, 1}, nil},
		{"right/left", op.AlignRight, [2]int{3, 1}, first, [2]int{1, 1}, nil},
		{"right/right", op.AlignRight, [2]int{1, 3}, rightPadded, [2]int{1, 3},
			[]string{"R10", "L1", "out:1,10", "R20", "L:end", "out:-1,20", "R30", "out:-1,30", "R:end"}},
		{"longest/left", op.AlignLongest, [2]int{3, 1}, leftPadded, [2]int{3, 1},
			[]string{"L1", "R10", "out:1,10", "L2", "R:end", "out:2,-2", "L3", "out:3,-2", "L:end"}},
		{"longest/right", op.AlignLongest, [2]int{1, 3}, rightPadded, [2]int{1, 3},
			[]string{"L1", "R10", "out:1,10", "L:end", "R20", "out:-1,20", "R30", "out:-1,30", "R:end"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var state zipTrace
			left := state.input(0, []int{1, 2, 3}[:tc.lengths[0]])
			right := state.input(1, []int{10, 20, 30}[:tc.lengths[1]])
			output := op.Zip(right).AlignBy(tc.align).WithDefaultKey(-1).WithDefaultValue(-2).Operator(left)
			var got []pair[int, int]
			for k, v := range output {
				got = append(got, pair[int, int]{k, v})
				state.events = append(state.events, fmt.Sprintf("out:%d,%d", k, v))
			}
			assertEqual(t, got, tc.want)
			assertEqual(t, state.reads, tc.reads)
			assertEqual(t, state.stopped, [2]int{1, 1})
			if tc.events != nil {
				assertEqual(t, state.events, tc.events)
			}
		})
	}
}

// 空は評価して判定し、nil で空と確定する場合は他方も開始しない。補完は型のゼロ値を使う。
func TestZipEmpty(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		align op.ZipAlign
		empty int
		want  []pair[int, int]
		reads [2]int
	}{
		{op.AlignShortest, 0, nil, [2]int{}},
		{op.AlignShortest, 1, nil, [2]int{1, 0}},
		{op.AlignLeft, 0, nil, [2]int{}},
		{op.AlignLeft, 1, []pair[int, int]{{1, 0}}, [2]int{1, 0}},
		{op.AlignRight, 0, []pair[int, int]{{0, 1}}, [2]int{0, 1}},
		{op.AlignRight, 1, nil, [2]int{}},
		{op.AlignLongest, 0, []pair[int, int]{{0, 1}}, [2]int{0, 1}},
		{op.AlignLongest, 1, []pair[int, int]{{1, 0}}, [2]int{1, 0}},
	} {
		for _, isNil := range []bool{false, true} {
			t.Run(fmt.Sprintf("align=%d/empty=%d/nil=%t", tc.align, tc.empty, isNil), func(t *testing.T) {
				t.Parallel()
				var state zipTrace
				inputs := [2]iter.Seq[int]{state.input(0, []int{1}), state.input(1, []int{1})}
				inputs[tc.empty] = state.input(tc.empty, nil)
				if isNil {
					inputs[tc.empty] = nil
				}
				output := op.Zip(inputs[1]).AlignBy(tc.align).Operator(inputs[0])
				if output == nil {
					t.Fatal("Operator must return a non-nil iterator")
				}
				assertEqual(t, collectPairs(output), tc.want)
				assertEqual(t, state.stopped, state.started)
				if isNil && tc.want == nil {
					assertEqual(t, state.started, [2]int{})
				} else {
					assertEqual(t, state.reads, tc.reads)
				}
				if !isNil {
					assertEqual(t, state.started[tc.empty], 1)
				}
			})
		}
	}
	assertEqual(t, collectPairs(op.Zip(iter.Seq[int](nil)).Operator(iter.Seq[int](nil))), []pair[int, int](nil))
	assertEqual(t, collectPairs(op.Zip(slices.Values([]int{})).Operator(slices.Values([]int{}))), []pair[int, int](nil))
}

// 全件評価の再実行と、左右の駆動ループ・Longest の右残余での停止と panic を確認する。
func TestZipEvaluation(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name        string
		align       op.ZipAlign
		left, right []int
		want        []pair[int, int]
		reads       [2]int
	}{
		{"paired/left", op.AlignShortest, []int{1, 2, 3}, []int{10, 20, 30},
			[]pair[int, int]{{1, 10}, {2, 20}}, [2]int{2, 2}},
		{"paired/longest", op.AlignLongest, []int{1, 2, 3}, []int{10, 20, 30},
			[]pair[int, int]{{1, 10}, {2, 20}}, [2]int{2, 2}},
		{"paired/right", op.AlignRight, []int{1, 2, 3}, []int{10, 20, 30},
			[]pair[int, int]{{1, 10}, {2, 20}}, [2]int{2, 2}},
		{"left", op.AlignLeft, []int{1, 2, 3}, []int{10}, []pair[int, int]{{1, 10}, {2, 0}}, [2]int{2, 1}},
		{"right", op.AlignRight, []int{1}, []int{10, 20, 30}, []pair[int, int]{{1, 10}, {0, 20}}, [2]int{1, 2}},
		{"remainder", op.AlignLongest, []int{1}, []int{10, 20, 30}, []pair[int, int]{{1, 10}, {0, 20}}, [2]int{1, 2}},
	} {
		for _, panics := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/panic=%t", tc.name, panics), func(t *testing.T) {
				t.Parallel()
				var state zipTrace
				output := op.Zip(state.input(1, tc.right)).AlignBy(tc.align).Operator(state.input(0, tc.left))
				assertEqual(t, state.started, [2]int{})
				var got []pair[int, int]
				func() {
					defer func() {
						var want any
						if panics {
							want = "downstream"
						}
						assertEqual(t, recover(), want)
					}()
					for k, v := range output {
						got = append(got, pair[int, int]{k, v})
						if len(got) == 2 {
							if panics {
								panic("downstream")
							}
							break
						}
					}
				}()
				assertEqual(t, got, tc.want)
				assertEqual(t, state.reads, tc.reads)
				assertEqual(t, state.stopped, [2]int{1, 1})
			})
		}
	}
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
		t.Run(fmt.Sprintf("reevaluate/align=%d", alignment), func(t *testing.T) {
			t.Parallel()
			var state zipTrace
			config := op.Zip(state.input(1, []int{10})).AlignBy(alignment)
			output := config.Operator(state.input(0, []int{1}))
			assertEqual(t, state.started, [2]int{})
			for i := 1; i <= 2; i++ {
				assertEqual(t, collectPairs(output), []pair[int, int]{{1, 10}})
				assertEqual(t, state.started, [2]int{i, i})
				assertEqual(t, state.reads, state.started)
				assertEqual(t, state.stopped, state.started)
			}
		})
	}
}

// 直接反復・Pull の両側と Longest の右残余で、入力 panic の伝播と解放を確認する。
func TestZipInputPanic(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		align      op.ZipAlign
		panicIndex int
		lengths    [2]int
	}{
		{op.AlignShortest, 0, [2]int{2, 2}}, {op.AlignShortest, 1, [2]int{2, 2}},
		{op.AlignLeft, 0, [2]int{2, 2}}, {op.AlignLeft, 1, [2]int{2, 2}},
		{op.AlignRight, 0, [2]int{2, 2}}, {op.AlignRight, 1, [2]int{2, 2}},
		{op.AlignLongest, 0, [2]int{2, 2}}, {op.AlignLongest, 1, [2]int{2, 2}},
		{op.AlignLongest, 1, [2]int{1, 2}},
	} {
		t.Run(fmt.Sprintf("align=%d/panic=%d", tc.align, tc.panicIndex), func(t *testing.T) {
			t.Parallel()
			var state zipTrace
			input := func(index int) iter.Seq[int] {
				return func(yield func(int) bool) {
					for v := range state.input(index, []int{1, 2}[:tc.lengths[index]]) {
						if index == tc.panicIndex && v == 2 {
							panic("input")
						}
						if !yield(v) {
							return
						}
					}
				}
			}
			func() {
				defer func() { assertEqual(t, recover(), any("input")) }()
				collectPairs(op.Zip(input(1)).AlignBy(tc.align).Operator(input(0)))
			}()
			assertEqual(t, state.started, [2]int{1, 1})
			assertEqual(t, state.stopped, state.started)
		})
	}
}

// 比較不能な入力と補完値を使い、同じ出力の並行評価でも状態が混ざらないことを確認する。
func TestZipNonComparableAndConcurrentEvaluation(t *testing.T) {
	t.Parallel()
	left := slices.Values([][]int{{1}})
	right := slices.Values([]map[string]int{{"a": 1}, {"b": 2}})
	config := op.Zip(right).WithDefaultKey([]int{-1}).WithDefaultValue(map[string]int{"missing": -1})
	output := config.AlignBy(op.AlignLongest).Operator(left)
	want := []pair[[]int, map[string]int]{
		{[]int{1}, map[string]int{"a": 1}}, {[]int{-1}, map[string]int{"b": 2}},
	}
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Go(func() { assertEqual(t, collectPairs(output), want) })
	}
	wg.Wait()
	missingValues := op.Zip(iter.Seq[map[string]int](nil)).WithDefaultValue(map[string]int{"missing": -1})
	assertEqual(t, collectPairs(missingValues.AlignBy(op.AlignLeft).Operator(left)),
		[]pair[[]int, map[string]int]{{[]int{1}, map[string]int{"missing": -1}}})
}

// 入力ごとの開始・取得・解放と、左右の取得順序を記録する。
type zipTrace struct {
	started, reads, stopped [2]int
	events                  []string
}

func (s *zipTrace) input(index int, values []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		s.started[index]++
		defer func() { s.stopped[index]++ }()
		name := []string{"L", "R"}[index]
		for _, v := range values {
			s.reads[index]++
			s.events = append(s.events, fmt.Sprint(name, v))
			if !yield(v) {
				return
			}
		}
		s.events = append(s.events, name+":end")
	}
}

// 一度しか読めない入力は、Zip を再評価しても復元されない。
func TestZipSingleUseInputs(t *testing.T) {
	t.Parallel()
	keys, values := []int{1, 2}, []string{"a", "b"}
	left := iter.Seq[int](func(yield func(int) bool) {
		for len(keys) > 0 {
			v := keys[0]
			keys = keys[1:]
			if !yield(v) {
				return
			}
		}
	})
	right := iter.Seq[string](func(yield func(string) bool) {
		for len(values) > 0 {
			v := values[0]
			values = values[1:]
			if !yield(v) {
				return
			}
		}
	})
	output := op.Zip(right).Operator(left)
	assertEqual(t, collectPairs(output), []pair[int, string]{{1, "a"}, {2, "b"}})
	assertEqual(t, collectPairs(output), []pair[int, string](nil))
}

// 有限側を基準に、無限入力も停止して解放する。
func TestZipFiniteAndInfinite(t *testing.T) {
	t.Parallel()
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight} {
		var reads, stops int
		infinite := iter.Seq[int](func(yield func(int) bool) {
			defer func() { stops++ }()
			for i := 0; ; i++ {
				reads++
				if reads > 4 {
					t.Fatal("did not stop")
				}
				if !yield(i) {
					return
				}
			}
		})
		finite := slices.Values([]int{10, 20})
		left, right := finite, infinite
		if alignment == op.AlignRight {
			left, right = infinite, finite
		}
		output := op.Zip(right).AlignBy(alignment).Operator(left)
		got := collectPairs(output)
		assertEqual(t, len(got), 2)
		assertEqual(t, reads, 2)
		assertEqual(t, stops, 1)
	}
}
