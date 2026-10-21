package op_test

import (
	"fmt"
	"iter"
	"slices"
	"strconv"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func ExampleMap() {
	pipeline := seq.FromSlice([]int{2, 1}).
		Pipe(op.Map(func(v int) string { return strconv.Itoa(v) }))
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["2" "1"]
}

func ExampleMap2() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb"})).
		Pipe2(op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) }))
	for k, v := range pipeline {
		fmt.Printf("%q: %d\n", k, v)
	}
	// Output:
	// "0": 1
	// "1": 3
}

func ExampleMapToSeq2() {
	pipeline := seq.FromSlice([]string{"Go", "iter"}).
		Pipe2(op.MapToSeq2(func(s string) (string, int) { return s, len(s) }))
	for k, v := range pipeline {
		fmt.Println(k, v)
	}
	// Output:
	// Go 2
	// iter 4
}

func ExampleMapToSeq() {
	pipeline := seq.FromSeq2(slices.All([]string{"Go", "iter"})).
		Pipe(op.MapToSeq(func(index int, s string) string {
			return fmt.Sprintf("%d:%s", index, s)
		}))
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [0:Go 1:iter]
}

func ExampleFilter() {
	pipeline := seq.FromSlice([]int{3, 2, 4, 1}).
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 }))
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [2 4]
}

func ExampleFilter2() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "c"})).
		Pipe2(op.Filter2(func(k int, v string) bool { return k > 0 && len(v) > 1 }))
	for k, v := range pipeline {
		fmt.Printf("%d: %q\n", k, v)
	}
	// Output: 1: "bbb"
}

func ExampleTake() {
	pipeline := seq.FromSlice([]int{3, 2, 1}).Pipe(op.Take(2).Operator)
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [3 2]
}

func ExampleTake_pairs() {
	pipeline := seq.FromSeq2(slices.All([]int{3, 2, 1})).Pipe2(op.Take(2).Operator2)
	for k, v := range pipeline {
		fmt.Printf("%d: %d\n", k, v)
	}
	// Output:
	// 0: 3
	// 1: 2
}

func ExampleMap2_filter() {
	src := func(yield func(int, int) bool) {
		if !yield(1, 10) {
			return
		}
		yield(2, 20)
	}
	pipeline := seq.FromSeq2(src).
		Pipe2(op.Map2(func(k, v int) (string, int) { return strconv.Itoa(k), v })).
		Pipe2(op.Filter2(func(k string, v int) bool { return v > 10 }))
	for k, v := range pipeline {
		fmt.Printf("%q: %d\n", k, v)
	}
	// Output: "2": 20
}

func ExampleOperator() {
	pipeline := seq.FromSlice([]int{1, 2}).Pipe(func(src iter.Seq[int]) iter.Seq[string] {
		return func(yield func(string) bool) {
			for v := range src {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["1" "2"]
}

func repeat(src iter.Seq[int]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for v := range src {
			for i := 0; i < 2; i++ {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	}
}

func ExampleOperator_custom() {
	pipeline := seq.FromSlice([]int{1, 2}).Pipe(repeat).Pipe(op.Take(3).Operator)
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["1" "1" "2"]
}

func keyText(src iter.Seq2[int, int]) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range src {
			if !yield(strconv.Itoa(k), v) {
				return
			}
		}
	}
}

func ExampleOperator2() {
	pipeline := seq.FromSeq2(slices.All([]int{4, 5})).Pipe2(keyText)
	for k, v := range pipeline {
		fmt.Printf("%q: %d\n", k, v)
	}
	// Output:
	// "0": 4
	// "1": 5
}

func ExampleKeys() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })).
		Pipe(op.Keys).
		Pipe(op.Take(2).Operator)
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["0" "1"]
}

func ExampleValues() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(op.Filter2(func(k int, v string) bool { return len(v) > 1 })).
		Pipe(op.Values).
		Pipe(op.Take(1).Operator)
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["bbb"]
}

func ExampleToSeq2Operator() {
	pipeline := seq.FromSlice([]string{"a", "bbb"}).
		Pipe2(op.MapToSeq2(func(v string) ([]int, string) { return []int{len(v)}, v }))
	for k, v := range pipeline {
		fmt.Println(k, v)
	}
	// Output:
	// [1] a
	// [3] bbb
}

func ExampleToSeqOperator() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb"})).Pipe(op.Values)
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [a bbb]
}

func ExampleZip() {
	left := seq.FromSlice([]int{1, 2, 3})
	right := seq.FromSlice([]string{"a", "b"})
	for k, v := range left.Pipe2(op.Zip(seq.ToSeq(right)).Operator) {
		fmt.Println(k, v)
	}
	// Output:
	// 1 a
	// 2 b
}

func ExampleZip_alignBy() {
	left := seq.FromSlice([]int{1, 2})
	right := seq.FromSlice([]string{"a"})
	alignLeft := op.Zip(seq.ToSeq(right)).AlignBy(op.AlignLeft).WithDefaultValue("missing")
	alignRight := op.Zip(seq.ToSeq(right)).WithDefaultKey(-1).AlignBy(op.AlignRight)
	alignLongest := op.Zip(seq.ToSeq(right)).AlignBy(op.AlignLongest)
	fmt.Println(seq.ToMap(left.Pipe2(alignLeft.Operator)))
	fmt.Println(seq.ToMap(seq.Seq[int](nil).Pipe2(alignRight.Operator)))
	fmt.Println(seq.ToMap(left.Pipe2(alignLongest.Operator)))
	// Output:
	// map[1:a 2:missing]
	// map[-1:a]
	// map[1:a 2:]
}

func ExampleSkip() {
	skip := op.Skip(1)
	values := seq.FromSlice([]int{10, 20, 30}).Pipe(skip.Operator)
	fmt.Println(seq.ToSlice(values))

	pairs := seq.FromSeq2(slices.All([]string{"a", "b", "c"})).Pipe2(skip.Operator2)
	for k, v := range pairs {
		fmt.Println(k, v)
	}
	// Output:
	// [20 30]
	// 1 b
	// 2 c
}

func ExampleCache() {
	calls := 0
	cached := seq.FromSlice([]string{"a", "bbb", "cc"}).
		Pipe(op.Map(func(v string) int {
			calls++
			return len(v)
		})).
		Pipe(op.Cache)
	fmt.Println("cached:", calls)
	first := cached.Pipe(op.Take(1).Operator)
	longer := cached.Pipe(op.Filter(func(v int) bool { return v > 1 }))
	fmt.Println(seq.ToSlice(first))
	fmt.Println(seq.ToSlice(longer))
	fmt.Println(seq.ToSlice(first))
	fmt.Println("upstream calls:", calls)
	// Output:
	// cached: 3
	// [1]
	// [3 2]
	// [1]
	// upstream calls: 3
}

func ExampleCache2() {
	calls := 0
	cached := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(op.Map2(func(k int, v string) (int, int) {
			calls++
			return k, len(v)
		})).
		Pipe2(op.Cache2)
	fmt.Println("cached:", calls)
	first := cached.Pipe2(op.Take(1).Operator2)
	for k, v := range first {
		fmt.Println(k, v)
	}
	fmt.Println(seq.ToSlice(cached.Pipe(op.Values)))
	for k, v := range first {
		fmt.Println(k, v)
	}
	fmt.Println("upstream calls:", calls)
	// Output:
	// cached: 3
	// 0 1
	// [1 3 2]
	// 0 1
	// upstream calls: 3
}

func ExampleCompose() {
	transform := op.Compose().
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(op.Map(strconv.Itoa)).
		Pipe(op.Take(2).Operator)
	first := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform.Operator)
	second := seq.FromSlice([]int{8, 9, 10, 12}).Pipe(transform.Operator)
	fmt.Printf("%q\n", seq.ToSlice(first))
	fmt.Printf("%q\n", seq.ToSlice(second))
	// Output:
	// ["2" "4"]
	// ["8" "10"]
}

func ExampleCompose_explicit() {
	transform := op.Compose().
		Pipe[int](op.Take(2).Operator).
		Pipe(op.Map(strconv.Itoa))
	fmt.Printf("%q\n", seq.ToSlice(seq.FromSlice([]int{3, 2, 1}).Pipe(transform.Operator)))
	// Output: ["3" "2"]
}

func ExampleCompose2() {
	transform := op.Compose2().
		Pipe2(op.Filter2(func(k int, v string) bool { return len(v) > 1 })).
		Pipe2(op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })).
		Pipe2(op.Take(2).Operator2)
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).Pipe2(transform.Operator)
	for k, v := range pipeline {
		fmt.Println(k, v)
	}
	// Output:
	// 1 3
	// 2 2
}

func ExampleCompose2_values() {
	transform := op.Compose2().
		Pipe[int, string](op.Values).
		Pipe(op.Map(func(v string) string { return v + "!" })).
		Pipe2(op.MapToSeq2(func(v string) ([]int, string) { return []int{len(v)}, v }))
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb"})).Pipe2(transform.Operator)
	for k, v := range pipeline {
		fmt.Println(k, v)
	}
	// Output:
	// [2] a!
	// [4] bbb!
}

func ExampleCompose_nested() {
	firstEven := op.Compose().
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(op.Take(2).Operator)
	transform := op.Compose().Pipe(firstEven.Operator).Pipe(op.Map(strconv.Itoa))
	pipeline := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform.Operator)
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["2" "4"]
}
