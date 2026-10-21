package seq_test

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

func ExampleSeq_Pipe() {
	var pipeline seq.Seq[int] = seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(op.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(op.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(op.Map(func(v string) int { return len(v) }))
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [3 4]
}

func ExampleSeq_Pipe2() {
	pairs := seq.FromSlice([]string{"a", "bbb", "cc"}).
		Pipe2(op.MapToSeq2(func(v string) ([]int, string) { return []int{len(v)}, v })).
		Pipe2(op.Take(2).Operator2)
	for k, v := range pairs {
		fmt.Println(k, v)
	}
	// Output:
	// [1] a
	// [3] bbb
}

func ExampleSeq2_Pipe() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe(op.Values).
		Pipe(op.Take(2).Operator)
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [a bbb]
}

func ExampleSeq2_Pipe2() {
	pipeline := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })).
		Pipe2(op.Take(2).Operator2)
	for k, v := range pipeline {
		fmt.Println(k, v)
	}
	// Output:
	// 0 1
	// 1 3
}

func ExampleFromChannel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	input := make(chan int, 4)
	for _, v := range []int{1, 2, 2, 3} {
		input <- v
	}
	close(input)
	pipeline := seq.FromChannel(ctx, input).
		Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
		Pipe(op.Map(func(v int) int { return v * 10 }))
	fmt.Println(seq.ToSlice(pipeline))
	// Output: [20 20]
}

func ExampleToChannel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 途中で受信をやめた場合も送信待ちを解除する。
	pipeline := seq.FromSlice([]int{1, 2, 2, 3}).
		Pipe(op.Map(func(v int) int { return v * 10 }))
	for v := range seq.ToChannel(ctx, pipeline) {
		fmt.Println(v)
		if v == 20 {
			break
		}
	}
	// Output:
	// 10
	// 20
}
