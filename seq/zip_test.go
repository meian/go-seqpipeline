package seq_test

import (
	"sync"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestSeqZip(t *testing.T) {
	for _, tc := range []struct {
		name   string
		keys   []int
		values []string
		want   []pair[int, string]
	}{
		{"equal", []int{2, 1, 2}, []string{"a", "b", "c"}, []pair[int, string]{{2, "a"}, {1, "b"}, {2, "c"}}},
		{"short keys", []int{1}, []string{"a", "b"}, []pair[int, string]{{1, "a"}}},
		{"short values", []int{1, 2}, []string{"a"}, []pair[int, string]{{1, "a"}}},
		{"empty keys", []int{}, []string{"a"}, nil},
		{"empty values", []int{1}, []string{}, nil},
		{"both empty", []int{}, []string{}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := seq.FromSlice(tc.keys).Zip(seq.FromSlice(tc.values))
			assertEqual(t, collectPairs(seq.ToSeq2(got)), tc.want)
		})
	}
	var keys seq.Seq[int]
	var values seq.Seq[string]
	for _, output := range []seq.Seq2[int, string]{keys.Zip(values), keys.Zip(seq.FromSlice([]string{"a"})), seq.FromSlice([]int{1}).Zip(values)} {
		if output == nil {
			t.Fatal("Zip must return a non-nil iterator")
		}
		assertEqual(t, collectPairs(seq.ToSeq2(output)), []pair[int, string](nil))
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

func TestSeqZipEvaluation(t *testing.T) {
	testZipEvaluation(t, func(k seq.Seq[int], v seq.Seq[string]) seq.Seq2[int, string] { return k.Zip(v) })
}

func TestSeqZipAllEvaluation(t *testing.T) {
	testZipEvaluation(t, func(k seq.Seq[int], v seq.Seq[string]) seq.Seq2[int, string] { return k.ZipAll(v) })
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
					assertEqual(t, seq.ToMap(zipped.Pipe(seq.Take2(1).Operator)), map[int]string{1: "a"})
				case "TakeZero":
					assertEqual(t, seq.ToMap(zipped.Pipe(seq.Take2(0).Operator)), map[int]string{})
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

func TestSeqZipAll(t *testing.T) {
	for _, keys := range [][]int{nil, {}, {1}, {1, 2, 1}} {
		for _, values := range [][]string{nil, {}, {"a"}, {"a", "b", "c"}} {
			for _, config := range []struct {
				key      int
				value    string
				defaults []seq.ZipDefaults[int, string]
			}{
				{},
				{-1, "", []seq.ZipDefaults[int, string]{seq.ZipDefaults[int, string]{}.WithKey(-1)}},
				{0, "missing", []seq.ZipDefaults[int, string]{seq.ZipDefaults[int, string]{}.WithValue("missing")}},
				{-1, "missing", []seq.ZipDefaults[int, string]{seq.ZipDefaults[int, string]{}.WithKey(-1).WithValue("missing")}},
				{-1, "missing", []seq.ZipDefaults[int, string]{seq.ZipDefaults[int, string]{}.WithValue("missing").WithKey(-1)}},
				{0, "missing", []seq.ZipDefaults[int, string]{seq.ZipDefaults[int, string]{}.WithKey(-1), seq.ZipDefaults[int, string]{}.WithValue("missing")}},
			} {
				var left seq.Seq[int]
				if keys != nil {
					left = seq.FromSlice(keys)
				}
				var right seq.Seq[string]
				if values != nil {
					right = seq.FromSlice(values)
				}
				var want []pair[int, string]
				for i := 0; i < max(len(keys), len(values)); i++ {
					k, v := config.key, config.value
					if i < len(keys) {
						k = keys[i]
					}
					if i < len(values) {
						v = values[i]
					}
					want = append(want, pair[int, string]{k, v})
				}
				output := left.ZipAll(right, config.defaults...)
				if output == nil {
					t.Fatal("ZipAll must return a non-nil iterator")
				}
				assertEqual(t, collectPairs(seq.ToSeq2(output)), want)
			}
		}
	}
}

func TestZipDefaultsWithKey(t *testing.T) {
	base := seq.ZipDefaults[int, string]{}.WithValue("missing")
	configured := base.WithKey(-1)
	empty := seq.Seq[int](nil)
	right := seq.FromSlice([]string{"a"})
	assertEqual(t, collectPairs(seq.ToSeq2(empty.ZipAll(right, configured))), []pair[int, string]{{-1, "a"}})
	assertEqual(t, collectPairs(seq.ToSeq2(empty.ZipAll(right, base))), []pair[int, string]{{0, "a"}})
	left := seq.FromSlice([]int{1})
	assertEqual(t, collectPairs(seq.ToSeq2(left.ZipAll(seq.Seq[string](nil), configured))), []pair[int, string]{{1, "missing"}})
}

func TestZipDefaultsWithValue(t *testing.T) {
	base := seq.ZipDefaults[int, string]{}.WithKey(-1)
	configured := base.WithValue("missing")
	empty := seq.Seq[string](nil)
	left := seq.FromSlice([]int{1})
	assertEqual(t, collectPairs(seq.ToSeq2(left.ZipAll(empty, configured))), []pair[int, string]{{1, "missing"}})
	assertEqual(t, collectPairs(seq.ToSeq2(left.ZipAll(empty, base))), []pair[int, string]{{1, ""}})
	right := seq.FromSlice([]string{"a"})
	assertEqual(t, collectPairs(seq.ToSeq2(seq.Seq[int](nil).ZipAll(right, configured))), []pair[int, string]{{-1, "a"}})
}

func TestZipInputTermination(t *testing.T) {
	for _, all := range []bool{false, true} {
		for _, lengths := range [][2]int{{0, 3}, {3, 0}, {1, 3}, {3, 1}} {
			var started, visited, stopped [2]int
			left := trackedZipInput(make([]int, lengths[0]), &started[0], &visited[0], &stopped[0])
			right := trackedZipInput(make([]string, lengths[1]), &started[1], &visited[1], &stopped[1])
			output := left.Zip(right)
			if all {
				output = left.ZipAll(right)
			}
			got := collectPairs(seq.ToSeq2(output))
			count := min(lengths[0], lengths[1])
			expectedVisited := [2]int{min(lengths[0], lengths[1]+1), count}
			if all {
				count = max(lengths[0], lengths[1])
				expectedVisited = lengths
			}
			assertEqual(t, len(got), count)
			assertEqual(t, visited, expectedVisited)
			assertEqual(t, stopped, started)
			if all {
				assertEqual(t, started, [2]int{1, 1})
			}
		}
	}
}

func TestZipInfiniteInputs(t *testing.T) {
	for _, all := range []bool{false, true} {
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
		output := source(0).Zip(source(1))
		if all {
			output = source(0).ZipAll(source(1))
		}
		assertEqual(t, len(collectPairs(seq.ToSeq2(output.Pipe(seq.Take2(3).Operator)))), 3)
		assertEqual(t, visited, [2]int{3, 3})
		assertEqual(t, stopped, [2]int{1, 1})
	}
}

func TestZipNonComparableAndConcurrentEvaluation(t *testing.T) {
	left := seq.FromSlice([][]int{{1}, {2}})
	right := seq.FromSlice([]map[string]int{{"a": 1}})
	var wg sync.WaitGroup
	for _, all := range []bool{false, true} {
		output := left.Zip(right)
		want := []pair[[]int, map[string]int]{{[]int{1}, map[string]int{"a": 1}}}
		if all {
			output = left.ZipAll(right)
			want = append(want, pair[[]int, map[string]int]{[]int{2}, nil})
		}
		for i := 0; i < 2; i++ {
			wg.Go(func() { assertEqual(t, collectPairs(seq.ToSeq2(output)), want) })
		}
	}
	wg.Wait()
}

func TestZipSingleUseInputs(t *testing.T) {
	for _, all := range []bool{false, true} {
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
		output := left.Zip(right)
		if all {
			output = left.ZipAll(right)
		}
		assertEqual(t, collectPairs(seq.ToSeq2(output)), []pair[int, string]{{1, "a"}, {2, "b"}})
		assertEqual(t, collectPairs(seq.ToSeq2(output)), []pair[int, string](nil))
	}
}

func TestZipInputPanic(t *testing.T) {
	for _, all := range []bool{false, true} {
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
			output := input(0).Zip(input(1))
			if all {
				output = input(0).ZipAll(input(1))
			}
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

func TestSeqZipAllStopAfterPadding(t *testing.T) {
	for _, lengths := range [][2]int{{0, 3}, {3, 0}, {1, 3}, {3, 1}} {
		var started, visited, stopped [2]int
		left := trackedZipInput(make([]int, lengths[0]), &started[0], &visited[0], &stopped[0])
		right := trackedZipInput(make([]string, lengths[1]), &started[1], &visited[1], &stopped[1])
		assertEqual(t, len(collectPairs(seq.ToSeq2(left.ZipAll(right).Pipe(seq.Take2(2).Operator)))), 2)
		assertEqual(t, visited, [2]int{min(lengths[0], 2), min(lengths[1], 2)})
		assertEqual(t, stopped, [2]int{1, 1})
	}
}
