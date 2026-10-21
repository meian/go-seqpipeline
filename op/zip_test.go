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

// 既定では短い入力で終了し、同じ設定を別のキー型へ接続できることを確認する。
func TestZip(t *testing.T) {
	left := slices.Values([]int{3, 1, 3})
	right := slices.Values([]string{"a", "b"})
	config := op.Zip(right)
	var ope op.ToSeq2Operator[int, int, string] = config.Operator
	assertEqual(t, collectPairs(ope(left)), []pair[int, string]{{3, "a"}, {1, "b"}})
	// 同じ設定から異なるキー型を推論できる。
	stringKeys := config.Operator(slices.Values([]string{"x"}))
	assertEqual(t, maps.Collect(stringKeys), map[string]string{"x": "a"})
}

// 入力の長さ・空・nil を組み合わせ、4つの終了条件ごとの出力順序と補完を確認する。
func TestZipAlignBy(t *testing.T) {
	type outputs struct {
		shortest, left, right, longest []pair[int, string]
	}
	both := []pair[int, string]{{3, "a"}, {1, "b"}, {3, "c"}}
	first := []pair[int, string]{{3, "a"}}
	paddedKeys := []pair[int, string]{{3, "a"}, {0, "b"}, {0, "c"}}
	paddedValues := []pair[int, string]{{3, "a"}, {1, ""}, {3, ""}}
	onlyKeys := []pair[int, string]{{3, ""}}
	onlyValues := []pair[int, string]{{0, "a"}}
	for _, tc := range []struct {
		name   string
		keys   iter.Seq[int]
		values iter.Seq[string]
		want   outputs
	}{
		{
			name:   "equal",
			keys:   slices.Values([]int{3, 1, 3}),
			values: slices.Values([]string{"a", "b", "c"}),
			want:   outputs{shortest: both, left: both, right: both, longest: both},
		},
		{
			name:   "keys longer",
			keys:   slices.Values([]int{3, 1, 3}),
			values: slices.Values([]string{"a"}),
			want:   outputs{shortest: first, left: paddedValues, right: first, longest: paddedValues},
		},
		{
			name:   "values longer",
			keys:   slices.Values([]int{3}),
			values: slices.Values([]string{"a", "b", "c"}),
			want:   outputs{shortest: first, left: first, right: paddedKeys, longest: paddedKeys},
		},
		{
			name:   "empty keys",
			keys:   slices.Values([]int{}),
			values: slices.Values([]string{"a"}),
			want:   outputs{right: onlyValues, longest: onlyValues},
		},
		{
			name:   "empty values",
			keys:   slices.Values([]int{3}),
			values: slices.Values([]string{}),
			want:   outputs{left: onlyKeys, longest: onlyKeys},
		},
		{
			name:   "both empty",
			keys:   slices.Values([]int{}),
			values: slices.Values([]string{}),
			want:   outputs{},
		},
		{
			name:   "nil keys",
			keys:   nil,
			values: slices.Values([]string{"a"}),
			want:   outputs{right: onlyValues, longest: onlyValues},
		},
		{
			name:   "nil values",
			keys:   slices.Values([]int{3}),
			values: nil,
			want:   outputs{left: onlyKeys, longest: onlyKeys},
		},
		{
			name:   "both nil",
			keys:   nil,
			values: nil,
			want:   outputs{},
		},
		{
			name:   "nil and empty",
			keys:   nil,
			values: slices.Values([]string{}),
			want:   outputs{},
		},
		{
			name:   "empty and nil",
			keys:   slices.Values([]int{}),
			values: nil,
			want:   outputs{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, align := range []struct {
				name  string
				value op.ZipAlign
				want  []pair[int, string]
			}{
				{"shortest", op.AlignShortest, tc.want.shortest},
				{"left", op.AlignLeft, tc.want.left},
				{"right", op.AlignRight, tc.want.right},
				{"longest", op.AlignLongest, tc.want.longest},
			} {
				t.Run(align.name, func(t *testing.T) {
					base := op.Zip(tc.values).AlignBy(align.value)
					for _, implementation := range []struct {
						name string
						ope  op.ToSeq2Operator[int, int, string]
					}{
						{"inferred key", base.Operator[int]}, {"fixed key", base.WithDefaultKey(0).Operator},
					} {
						t.Run(implementation.name, func(t *testing.T) {
							output := implementation.ope(tc.keys)
							if output == nil {
								t.Fatal("Operator must return a non-nil iterator")
							}
							assertEqual(t, collectPairs(output), align.want)
						})
					}
				})
			}
		})
	}
}

// 既定値の指定順序にかかわらず、不足分だけを補完し、入力中のゼロ値を保持する。
func TestZipDefaults(t *testing.T) {
	for _, tc := range []struct {
		name      string
		configure func(iter.Seq[string]) op.ToSeq2Operator[int, int, string]
		key       int
		value     string
	}{
		{"none", func(values iter.Seq[string]) op.ToSeq2Operator[int, int, string] {
			return op.Zip(values).AlignBy(op.AlignLongest).Operator
		}, 0, ""},
		{"key only", func(values iter.Seq[string]) op.ToSeq2Operator[int, int, string] {
			return op.Zip(values).AlignBy(op.AlignLongest).WithDefaultKey(-1).Operator
		}, -1, ""},
		{"value only", func(values iter.Seq[string]) op.ToSeq2Operator[int, int, string] {
			return op.Zip(values).AlignBy(op.AlignLongest).WithDefaultValue("missing").Operator
		}, 0, "missing"},
		{"key then value", func(values iter.Seq[string]) op.ToSeq2Operator[int, int, string] {
			return op.Zip(values).AlignBy(op.AlignLongest).WithDefaultKey(-1).WithDefaultValue("missing").Operator
		}, -1, "missing"},
		{"value then key", func(values iter.Seq[string]) op.ToSeq2Operator[int, int, string] {
			return op.Zip(values).AlignBy(op.AlignLongest).WithDefaultValue("missing").WithDefaultKey(-1).Operator
		}, -1, "missing"},
		{"align last", func(values iter.Seq[string]) op.ToSeq2Operator[int, int, string] {
			return op.Zip(values).WithDefaultKey(-1).WithDefaultValue("missing").AlignBy(op.AlignLongest).Operator
		}, -1, "missing"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range []struct {
				name   string
				keys   []int
				values []string
				want   []pair[int, string]
			}{
				{
					name:   "missing keys",
					keys:   []int{0},
					values: []string{"", "b", "c"},
					want:   []pair[int, string]{{0, ""}, {tc.key, "b"}, {tc.key, "c"}},
				},
				{
					name:   "missing values",
					keys:   []int{0, 2, 3},
					values: []string{""},
					want:   []pair[int, string]{{0, ""}, {2, tc.value}, {3, tc.value}},
				},
			} {
				t.Run(input.name, func(t *testing.T) {
					output := tc.configure(slices.Values(input.values))(slices.Values(input.keys))
					assertEqual(t, collectPairs(output), input.want)
				})
			}
		})
	}
}

// 設定変更が元の設定を変えず、再設定した項目だけが置き換わることを確認する。
func TestZipConfigurationCopies(t *testing.T) {
	right := slices.Values([]string{"a"})
	left := slices.Values([]int{1, 2})
	base := op.Zip(right)
	longest := base.AlignBy(op.AlignLongest)
	padded := longest.WithDefaultValue("missing")
	typed := padded.WithDefaultKey(-1)
	changed := typed.AlignBy(op.AlignRight).WithDefaultValue("other").WithDefaultKey(-2)
	assertEqual(t, maps.Collect(base.Operator(left)), map[int]string{1: "a"})
	assertEqual(t, maps.Collect(longest.Operator(left)), map[int]string{1: "a", 2: ""})
	assertEqual(t, maps.Collect(padded.Operator(left)), map[int]string{1: "a", 2: "missing"})
	assertEqual(t, maps.Collect(typed.Operator(left)), map[int]string{1: "a", 2: "missing"})
	assertEqual(t, maps.Collect(changed.Operator(left)), map[int]string{1: "a"})
	assertEqual(t, maps.Collect(typed.Operator(nil)), map[int]string{-1: "a"})
	assertEqual(t, maps.Collect(changed.Operator(nil)), map[int]string{-2: "a"})
	assertEqual(t, maps.Collect(changed.AlignBy(op.AlignLeft).Operator(left)), map[int]string{1: "a", 2: "other"})
	assertEqual(t, maps.Collect(typed.WithDefaultKey(0).Operator(nil)), map[int]string{0: "a"})
	assertEqual(t, maps.Collect(typed.WithDefaultValue("").Operator(left)), map[int]string{1: "a", 2: ""})
}

// 両設定型で不正な終了条件を拒否し、panic のメッセージに指定値を含める。
func TestZipAlignByInvalid(t *testing.T) {
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

// trackedZipInput は停止時のリソース解放と取得した要素数を記録する。
func trackedZipInput[T any](values []T, started, visited, stopped *int) iter.Seq[T] {
	return func(yield func(T) bool) {
		*started++
		defer func() { *stopped++ }()
		for _, v := range values {
			*visited++
			if !yield(v) {
				return
			}
		}
	}
}

// 両設定型と全終了条件で遅延評価・再評価を確認し、途中終了や panic でも開始した入力を解放する。
func TestZipEvaluation(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
		t.Run(fmt.Sprint(alignment), func(t *testing.T) {
			t.Run("inferred key", func(t *testing.T) {
				testZipEvaluation(t, func(k iter.Seq[int], v iter.Seq[string]) iter.Seq2[int, string] {
					return op.Zip(v).AlignBy(alignment).Operator(k)
				})
			})
			t.Run("fixed key", func(t *testing.T) {
				testZipEvaluation(t, func(k iter.Seq[int], v iter.Seq[string]) iter.Seq2[int, string] {
					return op.Zip(v).WithDefaultKey(-1).AlignBy(alignment).Operator(k)
				})
			})
		})
	}
}

func testZipEvaluation(t *testing.T, zip func(iter.Seq[int], iter.Seq[string]) iter.Seq2[int, string]) {
	t.Helper()
	for _, tc := range []struct {
		name    string
		consume func(*testing.T, iter.Seq2[int, string])
		reads   int
	}{
		{
			name: "full",
			consume: func(t *testing.T, src iter.Seq2[int, string]) {
				want := []pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}
				assertEqual(t, collectPairs(src), want)
			},
			reads: 3,
		},
		{
			name: "stop",
			consume: func(t *testing.T, src iter.Seq2[int, string]) {
				var got []pair[int, string]
				for k, v := range src {
					got = append(got, pair[int, string]{k, v})
					break
				}
				assertEqual(t, got, []pair[int, string]{{1, "a"}})
			},
			reads: 1,
		},
		{
			name: "panic",
			consume: func(t *testing.T, src iter.Seq2[int, string]) {
				defer func() {
					assertEqual(t, recover(), any("downstream"))
				}()
				src(func(int, string) bool { panic("downstream") })
			},
			reads: 1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var started, visited, stopped [2]int
			keys := trackedZipInput([]int{1, 2, 3}, &started[0], &visited[0], &stopped[0])
			values := trackedZipInput([]string{"a", "b", "c"}, &started[1], &visited[1], &stopped[1])
			zipped := zip(keys, values)
			assertEqual(t, started, [2]int{})
			for round := 1; round <= 2; round++ {
				tc.consume(t, zipped)
				assertEqual(t, started, [2]int{round, round})
				assertEqual(t, stopped, [2]int{round, round})
				assertEqual(t, visited, [2]int{tc.reads * round, tc.reads * round})
			}
		})
	}
}

// 比較不能なキー・値を扱い、同じ出力の並行評価が反復状態を共有しないことを確認する。
func TestZipNonComparableAndConcurrentEvaluation(t *testing.T) {
	left := slices.Values([][]int{{1}, {2}})
	right := slices.Values([]map[string]int{{"a": 1}})
	var wg sync.WaitGroup
	for _, all := range []bool{false, true} {
		output := op.Zip(right).Operator(left)
		want := []pair[[]int, map[string]int]{{[]int{1}, map[string]int{"a": 1}}}
		if all {
			output = op.Zip(right).AlignBy(op.AlignLongest).Operator(left)
			want = append(want, pair[[]int, map[string]int]{[]int{2}, nil})
		}
		for i := 0; i < 2; i++ {
			wg.Go(func() { assertEqual(t, collectPairs(output), want) })
		}
	}
	wg.Wait()
}

// 反復状態を外部に持つ両入力を消費した後は、Zip を再評価しても出力が空になる。
func TestZipSingleUseInputs(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
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
		output := op.Zip(right).AlignBy(alignment).Operator(left)
		assertEqual(t, collectPairs(output), []pair[int, string]{{1, "a"}, {2, "b"}})
		assertEqual(t, collectPairs(output), []pair[int, string](nil))
	}
}

// 左右どちらの入力の panic も呼び出し元へ伝わり、開始した両入力を解放する。
func TestZipInputPanic(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
		for _, panicIndex := range []int{0, 1} {
			var started, stopped [2]int
			input := func(index int) iter.Seq[int] {
				return func(yield func(int) bool) {
					started[index]++
					defer func() { stopped[index]++ }()
					if !yield(1) {
						return
					}
					if index == panicIndex {
						panic("input")
					}
					yield(2)
				}
			}
			output := op.Zip(input(1)).AlignBy(alignment).Operator(input(0))
			func() {
				defer func() {
					if got := recover(); got != "input" {
						t.Fatalf("panic = %v", got)
					}
				}()
				collectPairs(output)
			}()
			assertEqual(t, started, [2]int{1, 1})
			assertEqual(t, stopped, started)
		}
	}
}

// 不等長・空入力で、出力件数と左右の取得件数、開始した入力の解放を確認する。
func TestZipInputTermination(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
		for _, lengths := range [][2]int{{0, 3}, {3, 0}, {1, 3}, {3, 1}, {2, 2}} {
			var started, visited, stopped [2]int
			left := trackedZipInput(make([]int, lengths[0]), &started[0], &visited[0], &stopped[0])
			right := trackedZipInput(make([]string, lengths[1]), &started[1], &visited[1], &stopped[1])
			got := collectPairs(op.Zip(right).AlignBy(alignment).Operator(left))
			count := min(lengths[0], lengths[1])
			// Shortest では右の終了検出までに左を1要素余分に取得することがある。
			reads := [2]int{min(lengths[0], lengths[1]+1), count}
			switch alignment {
			case op.AlignLeft:
				count = lengths[0]
				reads = [2]int{count, min(count, lengths[1])}
			case op.AlignRight:
				count = lengths[1]
				reads = [2]int{min(count, lengths[0]), count}
			case op.AlignLongest:
				count = max(lengths[0], lengths[1])
				reads = lengths
			}
			assertEqual(t, len(got), count)
			assertEqual(t, visited, reads)
			assertEqual(t, stopped, started)
		}
	}
}

// 不等長入力で、左右の取得・終了検出・出力の順序を終了条件ごとに確認する。
func TestZipInputOrder(t *testing.T) {
	for _, tc := range []struct {
		name      string
		alignment op.ZipAlign
		left      []int
		right     []int
		want      []string
	}{
		{
			name:      "shortest",
			alignment: op.AlignShortest,
			left:      []int{1, 2},
			right:     []int{10},
			want:      []string{"L1", "R10", "out:1,10", "L2", "R:end"},
		},
		{
			name:      "left",
			alignment: op.AlignLeft,
			left:      []int{1, 2},
			right:     []int{10},
			want:      []string{"L1", "R10", "out:1,10", "L2", "R:end", "out:2,-2", "L:end"},
		},
		{
			name:      "right",
			alignment: op.AlignRight,
			left:      []int{1},
			right:     []int{10, 20},
			want:      []string{"R10", "L1", "out:1,10", "R20", "L:end", "out:-1,20", "R:end"},
		},
		{
			name:      "longest/right remainder",
			alignment: op.AlignLongest,
			left:      []int{1},
			right:     []int{10, 20},
			want:      []string{"L1", "R10", "out:1,10", "L:end", "R20", "out:-1,20", "R:end"},
		},
		{
			name:      "longest/left remainder",
			alignment: op.AlignLongest,
			left:      []int{1, 2},
			right:     []int{10},
			want:      []string{"L1", "R10", "out:1,10", "L2", "R:end", "out:2,-2", "L:end"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var events []string
			input := func(name string, values []int) iter.Seq[int] {
				return func(yield func(int) bool) {
					for _, v := range values {
						events = append(events, fmt.Sprint(name, v))
						if !yield(v) {
							return
						}
					}
					events = append(events, name+":end")
				}
			}
			config := op.Zip(input("R", tc.right)).AlignBy(tc.alignment).
				WithDefaultValue(-2).WithDefaultKey(-1)
			for k, v := range config.Operator(input("L", tc.left)) {
				events = append(events, fmt.Sprintf("out:%d,%d", k, v))
			}
			assertEqual(t, events, tc.want)
		})
	}
}

// 不足分の補完中でも下流の途中終了を伝え、余分な取得をせずに入力を解放する。
func TestZipStopAfterPadding(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignLeft, op.AlignRight, op.AlignLongest} {
		for _, lengths := range [][2]int{{0, 3}, {3, 0}, {1, 3}, {3, 1}} {
			var started, visited, stopped [2]int
			left := trackedZipInput(make([]int, lengths[0]), &started[0], &visited[0], &stopped[0])
			right := trackedZipInput(make([]string, lengths[1]), &started[1], &visited[1], &stopped[1])
			config := op.Zip(right).
				AlignBy(alignment).
				WithDefaultValue("missing").
				WithDefaultKey(-1)
			var got []pair[int, string]
			for k, v := range config.Operator(left) {
				got = append(got, pair[int, string]{k, v})
				if len(got) == 2 {
					break
				}
			}
			count := min(2, max(lengths[0], lengths[1]))
			if alignment == op.AlignLeft {
				count = min(2, lengths[0])
			}
			if alignment == op.AlignRight {
				count = min(2, lengths[1])
			}
			assertEqual(t, len(got), count)
			assertEqual(t, visited, [2]int{min(lengths[0], count), min(lengths[1], count)})
			assertEqual(t, stopped, started)
		}
	}
}

// 有限入力を終了基準にすると、相手が無限入力でも停止して解放する。
func TestZipFiniteAndInfinite(t *testing.T) {
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

// スライスや map を補完値として使い、キー・値の不足分を補えることを確認する。
func TestZipNonComparableDefaults(t *testing.T) {
	left := slices.Values([][]int{{1}})
	right := slices.Values([]map[string]int{{"a": 1}, {"b": 2}})
	config := op.Zip(right).
		WithDefaultKey([]int{-1}).
		WithDefaultValue(map[string]int{"missing": -1}).
		AlignBy(op.AlignLongest)
	assertEqual(t, collectPairs(config.Operator(left)), []pair[[]int, map[string]int]{
		{[]int{1}, map[string]int{"a": 1}}, {[]int{-1}, map[string]int{"b": 2}},
	})
	emptyRight := op.Zip(iter.Seq[map[string]int](nil)).
		WithDefaultKey([]int{-1}).
		WithDefaultValue(map[string]int{"missing": -1}).
		AlignBy(op.AlignLeft)
	output := emptyRight.Operator(left)
	want := []pair[[]int, map[string]int]{
		{[]int{1}, map[string]int{"missing": -1}},
	}
	assertEqual(t, collectPairs(output), want)
}

// nil 入力で出力が空と決まる終了条件では、もう片方の入力を評価しない。
func TestZipNilDoesNotEvaluateOther(t *testing.T) {
	unexpected := iter.Seq[int](func(func(int) bool) { t.Fatal("unexpected input evaluation") })
	empty := iter.Seq[int](nil)
	for _, output := range []iter.Seq2[int, int]{
		op.Zip(unexpected).Operator(empty),
		op.Zip(empty).Operator(unexpected),
		op.Zip(unexpected).AlignBy(op.AlignLeft).Operator(empty),
		op.Zip(empty).AlignBy(op.AlignRight).Operator(unexpected),
	} {
		assertEqual(t, collectPairs(output), []pair[int, int](nil))
	}
}
