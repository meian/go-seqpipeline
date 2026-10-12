package seq_test

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func ExampleMap() {
	op := seq.Map(func(v int) string { return strconv.Itoa(v) })
	fmt.Printf("%q\n", seq.ToSlice(op(seq.FromSlice([]int{2, 1}))))
	// Output: ["2" "1"]
}

func ExampleMap2() {
	op := seq.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	for k, v := range op(seq.FromSeq2(slices.All([]string{"a", "bbb"}))) {
		fmt.Printf("%q: %d\n", k, v)
	}
	// Output:
	// "0": 1
	// "1": 3
}

func ExampleFilter() {
	op := seq.Filter(func(v int) bool { return v%2 == 0 })
	fmt.Println(seq.ToSlice(op(seq.FromSlice([]int{3, 2, 4, 1}))))
	// Output: [2 4]
}

func ExampleFilter2() {
	op := seq.Filter2(func(k int, v string) bool { return k > 0 && len(v) > 1 })
	for k, v := range op(seq.FromSeq2(slices.All([]string{"a", "bbb", "c"}))) {
		fmt.Printf("%d: %q\n", k, v)
	}
	// Output: 1: "bbb"
}

func ExampleTake() {
	pipeline := seq.FromSlice([]int{3, 2, 1}).Pipe(seq.Take(2).Operator)
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [3 2]
}

func ExampleTake2() {
	pipeline := seq.FromSeq2(slices.All([]int{3, 2, 1})).Pipe(seq.Take2(2).Operator)
	for k, v := range pipeline {
		fmt.Printf("%d: %d\n", k, v)
	}
	// Output:
	// 0: 3
	// 1: 2
}

func ExampleSeq_Pipe() {
	var pipeline seq.Seq[int] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Map(func(v string) int { return len(v) }))
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [3 4]
}

func ExampleMap2_directCall() {
	src := func(yield func(int, int) bool) {
		if !yield(1, 10) {
			return
		}
		yield(2, 20)
	}
	mapped := seq.Map2(func(k, v int) (string, int) { return strconv.Itoa(k), v })(src)
	filtered := seq.Filter2(func(k string, v int) bool { return v > 10 })(mapped)
	for k, v := range filtered {
		fmt.Printf("%q: %d\n", k, v)
	}
	// Output: "2": 20
}

func ExampleOperator() {
	op := seq.Operator[int, string](func(src seq.Seq[int]) seq.Seq[string] {
		return func(yield func(string) bool) {
			for v := range src {
				if !yield(strconv.Itoa(v)) {
					return
				}
			}
		}
	})
	src := seq.FromSlice([]int{1, 2})
	fmt.Printf("%q\n", seq.ToSlice(src.Pipe(op)))
	fmt.Printf("%q\n", seq.ToSlice(op(src)))
	// Output:
	// ["1" "2"]
	// ["1" "2"]
}

func repeat(src seq.Seq[int]) seq.Seq[string] {
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
	var op seq.Operator[int, string] = repeat
	pipeline := seq.FromSlice([]int{1, 2}).Pipe(op).Pipe(seq.Take(3).Operator)
	fmt.Printf("%q\n", seq.ToSlice(pipeline))
	// Output: ["1" "1" "2"]
}

func keyText(src seq.Seq2[int, int]) seq.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for k, v := range src {
			if !yield(strconv.Itoa(k), v) {
				return
			}
		}
	}
}

func ExampleOperator2() {
	var op seq.Operator2[int, int, string, int] = keyText
	src := seq.FromSeq2(slices.All([]int{4, 5}))
	for k, v := range src.Pipe(op) {
		fmt.Printf("%q: %d\n", k, v)
	}
	for k, v := range op(src) {
		fmt.Printf("%q: %d\n", k, v)
	}
	// Output:
	// "0": 4
	// "1": 5
	// "0": 4
	// "1": 5
}
