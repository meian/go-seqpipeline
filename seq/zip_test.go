package seq_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestZip(t *testing.T) {
	left := seq.FromSlice([]int{3, 1, 3})
	right := seq.FromSlice([]string{"a", "b"})
	config := seq.Zip(right)
	var op seq.ToSeq2Operator[int, int, string] = config.Operator
	assertEqual(t, collectPairs(seq.ToSeq2(op(left))), []pair[int, string]{{3, "a"}, {1, "b"}})
	// 同じ設定から異なるキー型を推論できる。
	assertEqual(t, seq.ToMap(seq.FromSlice([]string{"x"}).Pipe2(config.Operator)), map[string]string{"x": "a"})
}

func TestZipAlignBy(t *testing.T) {
	both := []pair[int, string]{{3, "a"}, {1, "b"}, {3, "c"}}
	first := []pair[int, string]{{3, "a"}}
	paddedKeys := []pair[int, string]{{3, "a"}, {0, "b"}, {0, "c"}}
	paddedValues := []pair[int, string]{{3, "a"}, {1, ""}, {3, ""}}
	onlyKeys := []pair[int, string]{{3, ""}}
	onlyValues := []pair[int, string]{{0, "a"}}
	for _, tc := range []struct {
		name   string
		keys   seq.Seq[int]
		values seq.Seq[string]
		want   [4][]pair[int, string]
	}{
		{"equal", seq.FromSlice([]int{3, 1, 3}), seq.FromSlice([]string{"a", "b", "c"}), [4][]pair[int, string]{both, both, both, both}},
		{"keys longer", seq.FromSlice([]int{3, 1, 3}), seq.FromSlice([]string{"a"}), [4][]pair[int, string]{first, paddedValues, first, paddedValues}},
		{"values longer", seq.FromSlice([]int{3}), seq.FromSlice([]string{"a", "b", "c"}), [4][]pair[int, string]{first, first, paddedKeys, paddedKeys}},
		{"empty keys", seq.FromSlice([]int{}), seq.FromSlice([]string{"a"}), [4][]pair[int, string]{nil, nil, onlyValues, onlyValues}},
		{"empty values", seq.FromSlice([]int{3}), seq.FromSlice([]string{}), [4][]pair[int, string]{nil, onlyKeys, nil, onlyKeys}},
		{"both empty", seq.FromSlice([]int{}), seq.FromSlice([]string{}), [4][]pair[int, string]{}},
		{"nil keys", nil, seq.FromSlice([]string{"a"}), [4][]pair[int, string]{nil, nil, onlyValues, onlyValues}},
		{"nil values", seq.FromSlice([]int{3}), nil, [4][]pair[int, string]{nil, onlyKeys, nil, onlyKeys}},
		{"both nil", nil, nil, [4][]pair[int, string]{}},
		{"nil and empty", nil, seq.FromSlice([]string{}), [4][]pair[int, string]{}},
		{"empty and nil", seq.FromSlice([]int{}), nil, [4][]pair[int, string]{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for i, align := range []struct {
				name  string
				value seq.ZipAlign
			}{
				{"shortest", seq.AlignShortest}, {"left", seq.AlignLeft}, {"right", seq.AlignRight}, {"longest", seq.AlignLongest},
			} {
				t.Run(align.name, func(t *testing.T) {
					base := seq.Zip(tc.values).AlignBy(align.value)
					for _, implementation := range []struct {
						name string
						op   seq.ToSeq2Operator[int, int, string]
					}{
						{"inferred key", base.Operator[int]}, {"fixed key", base.WithKey(0).Operator},
					} {
						t.Run(implementation.name, func(t *testing.T) {
							output := implementation.op(tc.keys)
							if output == nil {
								t.Fatal("Operator must return a non-nil iterator")
							}
							assertEqual(t, collectPairs(seq.ToSeq2(output)), tc.want[i])
						})
					}
				})
			}
		})
	}
}

func TestZipDefaults(t *testing.T) {
	for _, tc := range []struct {
		name      string
		configure func(seq.Seq[string]) seq.ToSeq2Operator[int, int, string]
		key       int
		value     string
	}{
		{"none", func(values seq.Seq[string]) seq.ToSeq2Operator[int, int, string] {
			return seq.Zip(values).AlignBy(seq.AlignLongest).Operator
		}, 0, ""},
		{"key only", func(values seq.Seq[string]) seq.ToSeq2Operator[int, int, string] {
			return seq.Zip(values).AlignBy(seq.AlignLongest).WithKey(-1).Operator
		}, -1, ""},
		{"value only", func(values seq.Seq[string]) seq.ToSeq2Operator[int, int, string] {
			return seq.Zip(values).AlignBy(seq.AlignLongest).WithValue("missing").Operator
		}, 0, "missing"},
		{"key then value", func(values seq.Seq[string]) seq.ToSeq2Operator[int, int, string] {
			return seq.Zip(values).AlignBy(seq.AlignLongest).WithKey(-1).WithValue("missing").Operator
		}, -1, "missing"},
		{"value then key", func(values seq.Seq[string]) seq.ToSeq2Operator[int, int, string] {
			return seq.Zip(values).AlignBy(seq.AlignLongest).WithValue("missing").WithKey(-1).Operator
		}, -1, "missing"},
		{"align last", func(values seq.Seq[string]) seq.ToSeq2Operator[int, int, string] {
			return seq.Zip(values).WithKey(-1).WithValue("missing").AlignBy(seq.AlignLongest).Operator
		}, -1, "missing"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, input := range []struct {
				name   string
				keys   []int
				values []string
				want   []pair[int, string]
			}{
				{"missing keys", []int{0}, []string{"", "b", "c"}, []pair[int, string]{{0, ""}, {tc.key, "b"}, {tc.key, "c"}}},
				{"missing values", []int{0, 2, 3}, []string{""}, []pair[int, string]{{0, ""}, {2, tc.value}, {3, tc.value}}},
			} {
				t.Run(input.name, func(t *testing.T) {
					output := tc.configure(seq.FromSlice(input.values))(seq.FromSlice(input.keys))
					assertEqual(t, collectPairs(seq.ToSeq2(output)), input.want)
				})
			}
		})
	}
}

func TestZipConfigurationCopies(t *testing.T) {
	right := seq.FromSlice([]string{"a"})
	left := seq.FromSlice([]int{1, 2})
	base := seq.Zip(right)
	longest := base.AlignBy(seq.AlignLongest)
	padded := longest.WithValue("missing")
	typed := padded.WithKey(-1)
	changed := typed.AlignBy(seq.AlignRight).WithValue("other").WithKey(-2)
	assertEqual(t, seq.ToMap(base.Operator(left)), map[int]string{1: "a"})
	assertEqual(t, seq.ToMap(longest.Operator(left)), map[int]string{1: "a", 2: ""})
	assertEqual(t, seq.ToMap(padded.Operator(left)), map[int]string{1: "a", 2: "missing"})
	assertEqual(t, seq.ToMap(typed.Operator(left)), map[int]string{1: "a", 2: "missing"})
	assertEqual(t, seq.ToMap(changed.Operator(left)), map[int]string{1: "a"})
	assertEqual(t, seq.ToMap(typed.Operator(nil)), map[int]string{-1: "a"})
	assertEqual(t, seq.ToMap(changed.Operator(nil)), map[int]string{-2: "a"})
	assertEqual(t, seq.ToMap(changed.AlignBy(seq.AlignLeft).Operator(left)), map[int]string{1: "a", 2: "other"})
	assertEqual(t, seq.ToMap(typed.WithKey(0).Operator(nil)), map[int]string{0: "a"})
	assertEqual(t, seq.ToMap(typed.WithValue("").Operator(left)), map[int]string{1: "a", 2: ""})
}

func TestZipAlignByInvalid(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{-1, seq.AlignLongest + 1} {
		for _, configure := range []func(){
			func() { seq.Zip(seq.Seq[string](nil)).AlignBy(alignment) },
			func() { seq.Zip(seq.Seq[string](nil)).WithKey(0).AlignBy(alignment) },
		} {
			func() {
				defer func() { assertEqual(t, recover(), any(fmt.Sprintf("seq.Zip: invalid alignment %d", alignment))) }()
				configure()
			}()
		}
	}
}

// trackedZipInput は停止時のリソース解放と取得した要素数を記録する。
func trackedZipInput[T any](values []T, started, visited, stopped *int) seq.Seq[T] {
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

func TestZipEvaluation(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignShortest, seq.AlignLeft, seq.AlignRight, seq.AlignLongest} {
		t.Run(fmt.Sprint(alignment), func(t *testing.T) {
			testZipEvaluation(t, func(k seq.Seq[int], v seq.Seq[string]) seq.Seq2[int, string] {
				return seq.Zip(v).AlignBy(alignment).Operator(k)
			})
			testZipEvaluation(t, func(k seq.Seq[int], v seq.Seq[string]) seq.Seq2[int, string] {
				return seq.Zip(v).WithKey(-1).AlignBy(alignment).Operator(k)
			})
		})
	}
}

func testZipEvaluation(t *testing.T, zip func(seq.Seq[int], seq.Seq[string]) seq.Seq2[int, string]) {
	t.Helper()
	for _, mode := range []string{"full", "yield", "break", "Take", "TakeZero", "panic"} {
		t.Run(mode, func(t *testing.T) {
			var started, visited, stopped [2]int
			keys := trackedZipInput([]int{1, 2, 3}, &started[0], &visited[0], &stopped[0])
			values := trackedZipInput([]string{"a", "b", "c"}, &started[1], &visited[1], &stopped[1])
			zipped := zip(keys, values)
			assertEqual(t, started, [2]int{})
			for i := 1; i <= 2; i++ {
				switch mode {
				case "full":
					assertEqual(t, collectPairs(seq.ToSeq2(zipped)), []pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}})
				case "yield":
					zipped(func(int, string) bool { return false })
				case "break":
					for range zipped {
						break
					}
				case "Take":
					assertEqual(t, seq.ToMap(zipped.Pipe2(seq.Take2(1).Operator)), map[int]string{1: "a"})
				case "TakeZero":
					assertEqual(t, seq.ToMap(zipped.Pipe2(seq.Take2(0).Operator)), map[int]string{})
				case "panic":
					func() {
						defer func() {
							if got := recover(); got != "downstream" {
								t.Fatalf("panic = %v", got)
							}
						}()
						zipped(func(int, string) bool { panic("downstream") })
					}()
				}
				n := i
				if mode == "TakeZero" {
					n = 0
				}
				assertEqual(t, started, [2]int{n, n})
				assertEqual(t, stopped, [2]int{n, n})
				if mode == "full" {
					n *= 3
				}
				assertEqual(t, visited, [2]int{n, n})
			}
		})
	}
}

func TestZipInfiniteInputs(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignShortest, seq.AlignLeft, seq.AlignRight, seq.AlignLongest} {
		var visited, stopped [2]int
		source := func(index int) seq.Seq[int] {
			return func(yield func(int) bool) {
				defer func() { stopped[index]++ }()
				for i := 0; ; i++ {
					visited[index]++
					if !yield(i) {
						return
					}
				}
			}
		}
		output := source(0).Pipe2(seq.Zip(source(1)).AlignBy(alignment).Operator)
		assertEqual(t, len(collectPairs(seq.ToSeq2(output.Pipe2(seq.Take2(3).Operator)))), 3)
		assertEqual(t, visited, [2]int{3, 3})
		assertEqual(t, stopped, [2]int{1, 1})
	}
}

func TestZipNonComparableAndConcurrentEvaluation(t *testing.T) {
	left := seq.FromSlice([][]int{{1}, {2}})
	right := seq.FromSlice([]map[string]int{{"a": 1}})
	var wg sync.WaitGroup
	for _, all := range []bool{false, true} {
		output := left.Pipe2(seq.Zip(right).Operator)
		want := []pair[[]int, map[string]int]{{[]int{1}, map[string]int{"a": 1}}}
		if all {
			output = left.Pipe2(seq.Zip(right).AlignBy(seq.AlignLongest).Operator)
			want = append(want, pair[[]int, map[string]int]{[]int{2}, nil})
		}
		for i := 0; i < 2; i++ {
			wg.Go(func() { assertEqual(t, collectPairs(seq.ToSeq2(output)), want) })
		}
	}
	wg.Wait()
}

func TestZipSingleUseInputs(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignShortest, seq.AlignLeft, seq.AlignRight, seq.AlignLongest} {
		keys, values := []int{1, 2}, []string{"a", "b"}
		left := seq.Seq[int](func(yield func(int) bool) {
			for len(keys) > 0 {
				v := keys[0]
				keys = keys[1:]
				if !yield(v) {
					return
				}
			}
		})
		right := seq.Seq[string](func(yield func(string) bool) {
			for len(values) > 0 {
				v := values[0]
				values = values[1:]
				if !yield(v) {
					return
				}
			}
		})
		output := left.Pipe2(seq.Zip(right).AlignBy(alignment).Operator)
		assertEqual(t, collectPairs(seq.ToSeq2(output)), []pair[int, string]{{1, "a"}, {2, "b"}})
		assertEqual(t, collectPairs(seq.ToSeq2(output)), []pair[int, string](nil))
	}
}

func TestZipInputPanic(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignShortest, seq.AlignLeft, seq.AlignRight, seq.AlignLongest} {
		for _, panicIndex := range []int{0, 1} {
			var started, stopped [2]int
			input := func(index int) seq.Seq[int] {
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
			output := input(0).Pipe2(seq.Zip(input(1)).AlignBy(alignment).Operator)
			func() {
				defer func() {
					if got := recover(); got != "input" {
						t.Fatalf("panic = %v", got)
					}
				}()
				collectPairs(seq.ToSeq2(output))
			}()
			assertEqual(t, started, [2]int{1, 1})
			assertEqual(t, stopped, started)
		}
	}
}

func TestZipInputTermination(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignShortest, seq.AlignLeft, seq.AlignRight, seq.AlignLongest} {
		for _, lengths := range [][2]int{{0, 3}, {3, 0}, {1, 3}, {3, 1}, {2, 2}} {
			var started, visited, stopped [2]int
			left := trackedZipInput(make([]int, lengths[0]), &started[0], &visited[0], &stopped[0])
			right := trackedZipInput(make([]string, lengths[1]), &started[1], &visited[1], &stopped[1])
			got := collectPairs(seq.ToSeq2(seq.Zip(right).AlignBy(alignment).Operator(left)))
			count := min(lengths[0], lengths[1])
			reads := [2]int{min(lengths[0], lengths[1]+1), count}
			switch alignment {
			case seq.AlignLeft:
				count = lengths[0]
				reads = [2]int{count, min(count, lengths[1])}
			case seq.AlignRight:
				count = lengths[1]
				reads = [2]int{min(count, lengths[0]), count}
			case seq.AlignLongest:
				count = max(lengths[0], lengths[1])
				reads = lengths
			}
			assertEqual(t, len(got), count)
			assertEqual(t, visited, reads)
			assertEqual(t, stopped, started)
		}
	}
}

func TestZipStopAfterPadding(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignLeft, seq.AlignRight, seq.AlignLongest} {
		for _, lengths := range [][2]int{{0, 3}, {3, 0}, {1, 3}, {3, 1}} {
			var started, visited, stopped [2]int
			left := trackedZipInput(make([]int, lengths[0]), &started[0], &visited[0], &stopped[0])
			right := trackedZipInput(make([]string, lengths[1]), &started[1], &visited[1], &stopped[1])
			output := left.Pipe2(seq.Zip(right).AlignBy(alignment).WithValue("missing").WithKey(-1).Operator).Pipe2(seq.Take2(2).Operator)
			count := min(2, max(lengths[0], lengths[1]))
			if alignment == seq.AlignLeft {
				count = min(2, lengths[0])
			}
			if alignment == seq.AlignRight {
				count = min(2, lengths[1])
			}
			assertEqual(t, len(collectPairs(seq.ToSeq2(output))), count)
			assertEqual(t, visited, [2]int{min(lengths[0], count), min(lengths[1], count)})
			assertEqual(t, stopped, started)
		}
	}
}

func TestZipFiniteAndInfinite(t *testing.T) {
	for _, alignment := range []seq.ZipAlign{seq.AlignShortest, seq.AlignLeft, seq.AlignRight} {
		var reads, stops int
		infinite := seq.Seq[int](func(yield func(int) bool) {
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
		finite := seq.FromSlice([]int{10, 20})
		left, right := finite, infinite
		if alignment == seq.AlignRight {
			left, right = infinite, finite
		}
		assertEqual(t, len(collectPairs(seq.ToSeq2(left.Pipe2(seq.Zip(right).AlignBy(alignment).Operator)))), 2)
		assertEqual(t, reads, 2)
		assertEqual(t, stops, 1)
	}
}

func TestZipNonComparableDefaults(t *testing.T) {
	left := seq.FromSlice([][]int{{1}})
	right := seq.FromSlice([]map[string]int{{"a": 1}, {"b": 2}})
	config := seq.Zip(right).WithKey([]int{-1}).WithValue(map[string]int{"missing": -1}).AlignBy(seq.AlignLongest)
	assertEqual(t, collectPairs(seq.ToSeq2(left.Pipe2(config.Operator))), []pair[[]int, map[string]int]{
		{[]int{1}, map[string]int{"a": 1}}, {[]int{-1}, map[string]int{"b": 2}},
	})
	assertEqual(t, collectPairs(seq.ToSeq2(seq.Zip(seq.Seq[map[string]int](nil)).WithKey([]int{-1}).WithValue(map[string]int{"missing": -1}).AlignBy(seq.AlignLeft).Operator(left))), []pair[[]int, map[string]int]{{[]int{1}, map[string]int{"missing": -1}}})
}

func TestZipNilDoesNotEvaluateOther(t *testing.T) {
	unexpected := seq.Seq[int](func(func(int) bool) { t.Fatal("unexpected input evaluation") })
	empty := seq.Seq[int](nil)
	for _, output := range []seq.Seq2[int, int]{
		seq.Zip(unexpected).Operator(empty), seq.Zip(empty).Operator(unexpected),
		seq.Zip(unexpected).AlignBy(seq.AlignLeft).Operator(empty), seq.Zip(empty).AlignBy(seq.AlignRight).Operator(unexpected),
	} {
		assertEqual(t, collectPairs(seq.ToSeq2(output)), []pair[int, int](nil))
	}
}
