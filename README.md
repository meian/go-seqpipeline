# go-seqpipeline

English | [日本語](README.ja.md)  

A library for building type-safe pipelines through method chaining with `iter.Seq` / `iter.Seq2`.  

## Requirements and installation

Go 1.27.1 or later is required.  

```sh
go get github.com/meian/go-seqpipeline
```

## Usage

Use functions starting with `From` to wrap input, then connect operators with `Pipe` / `Pipe2`.  
Types are inferred from the inputs and outputs of the connected operations, and operations with incompatible types cannot be connected.  
Iterate over the result directly with `for range`, or collect it into a collection with `ToSlice` / `ToMap`.  

```go
package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/meian/go-seqpipeline/seq"
)

func main() {
	pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
		Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
		Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
		Pipe(seq.Take(1).Operator)

	// Collect the pipeline results into a slice
	fmt.Println(seq.ToSlice(pipeline)) // [333]

	// Iterate over the pipeline directly
	for s := range pipeline {
		fmt.Println(s) // 333
	}
	// Transform keys and values with Seq2
	pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
		Pipe2(seq.Map2(func(k int, v string) (string, int) {
			return strconv.Itoa(k), len(v)
		})).
		Pipe2(seq.Filter2(func(k string, v int) bool { return v > 1 })).
		Pipe2(seq.Take(1).Operator2)
	for k, v := range pairs {
		fmt.Println(k, v) // 1 3
	}
}
```

## Seq / Seq2

`Seq[T]` / `Seq2[K, V]` extend the function types from Go's standard [`iter`](https://pkg.go.dev/iter) package with `Pipe` / `Pipe2` methods for connecting transformations.  
Like standard iterators, the results can be iterated over directly with `for v := range pipeline` / `for k, v := range pipeline`.  
`Seq2[K, V]` places no constraints on `K` or `V`, but `FromMap` / `ToMap`, which convert to and from maps, require `K` to be `comparable`.  
A `Seq2` can also be converted into a single-value `Seq` to continue the pipeline.  

## Input and output

Functions starting with `From` convert collections, standard iterators, or channels into pipeline inputs.  
Functions starting with `To` collect results into collections or convert them to standard iterators or channels.  

| Operation | Seq | Seq2 |
| --- | --- | --- |
| Input from a collection | `FromSlice` | `FromMap` |
| Input from a standard iterator | `FromSeq` | `FromSeq2` |
| Collect into a collection | `ToSlice` | `ToMap` (keys must be `comparable`) |
| Convert to a standard iterator | `ToSeq` | `ToSeq2` |
| Input from / output to a channel | `FromChannel` / `ToChannel` | — |

Use `ToSeq` / `ToSeq2` when passing a pipeline to a function that accepts the standard `iter.Seq` / `iter.Seq2` types.  

### Channel conversions

`FromChannel(ctx, ch)` receives during iteration without starting a goroutine and ends when the input closes or the context is canceled.  
The caller owns the input channel; the conversion never closes it.  
Early termination downstream stops further receives, but the caller must manage stopping the sender.  
The input is consumable: re-evaluation receives remaining values, and multiple receivers divide the values among themselves.  
A nil channel waits until cancellation, while a closed, drained channel ends iteration immediately.  

`ToChannel(ctx, src)` starts evaluation in a dedicated goroutine when called and returns a new unbuffered receive-only channel.  
That goroutine closes the output when the input ends or the context is canceled.  
Nil or empty input closes the output without sending any values.  
If you stop receiving early, always cancel the context created with `context.WithCancel` or a similar function.  
Input may be evaluated ahead of the receiver to obtain the next value to send.  
Whether repeated or concurrent evaluation is safe depends on the input.  

Both conversions require a non-nil context and preserve value order and duplicates.  
Cancellation releases communication waits, but a concurrent communication may still complete.  
It cannot interrupt waits inside arbitrary iterators or callbacks; make those stoppable with the same context or another mechanism.  
See [ExampleFromChannel](seq/example_test.go#L271) and [ExampleToChannel](seq/example_test.go#L286) for usage.  

## Operators

In go-seqpipeline, an operator is an operation that transforms `Seq` / `Seq2` to produce a new sequence.  

### Built-in operators

The library provides operations such as transforming elements, filtering them, and limiting the number of results.  
For operations with separate functions for single values and key-value pairs, the library provides the pair version with a `2` suffix, as in `Filter` and `Filter2`.  

### Type inference

When types can be inferred from the arguments to an operator factory, pass its return value directly to `Pipe` / `Pipe2` according to the output type.  
Arguments such as callbacks carry the input and output type information, so explicit type arguments are unnecessary.  

Operations whose factory arguments carry no type information return a configuration value whose methods infer types from the input.  
Pass `.Operator` as a method value to `Pipe` for single values, or `.Operator2` to `Pipe2` for key-value pairs.  
The input determines the types when connected, so explicit type arguments are unnecessary.  

### Connecting operators

Both `Seq` and `Seq2` define `Pipe` / `Pipe2`, which accept operator functions that return `Seq` / `Seq2`, respectively.  
The function types `Operator`, `ToSeq2Operator`, `ToSeqOperator`, and `Operator2` represent the four combinations of input and output types.  

| Input | Method | Transformation function | Output |
| --- | --- | --- | --- |
| `Seq[I]` | `Pipe` | `Operator[I, O]` (`func(Seq[I]) Seq[O]`) | `Seq[O]` |
| `Seq[I]` | `Pipe2` | `ToSeq2Operator[I, K, V]` (`func(Seq[I]) Seq2[K, V]`) | `Seq2[K, V]` |
| `Seq2[K, V]` | `Pipe` | `ToSeqOperator[K, V, O]` (`func(Seq2[K, V]) Seq[O]`) | `Seq[O]` |
| `Seq2[K, V]` | `Pipe2` | `Operator2[K, V, K2, V2]` (`func(Seq2[K, V]) Seq2[K2, V2]`) | `Seq2[K2, V2]` |

`Pipe` / `Pipe2` call the supplied transformation function immediately and return its result unchanged, including nil.  
Input evaluation and stop propagation depend on the transformation function.  
Because all four operator types are function types, they can also be called directly as `op(src)`.  

### Implementing custom operators

Implement a transformation function using this library's `Seq` / `Seq2` types for input and output, and pass it directly to `Pipe` / `Pipe2` according to the output type.  
The function should return an iterator that processes input when executed, without evaluating input during the transformation call.  
You can also compose existing operators by calling `src.Pipe(...)` / `src.Pipe2(...)` inside the function.  
For operations whose factory arguments carry no type information, provide a configuration value with a transformation method that has type parameters, and pass its method value to `Pipe` / `Pipe2` according to the output type.  

### Executable examples

[Go Examples](seq/example_test.go) demonstrate each operator with expected output.  
Regular tests cover boundary values, nil, lazy evaluation, early termination, and re-evaluation.  

| Example | Seq | Seq2 |
| --- | --- | --- |
| Transform elements or keys and values | [ExampleMap](seq/example_test.go#L11) | [ExampleMap2](seq/example_test.go#L17) |
| Filter by a condition | [ExampleFilter](seq/example_test.go#L27) | [ExampleFilter2](seq/example_test.go#L33) |
| Limit results and infer types (.Operator / .Operator2) | [ExampleTake](seq/example_test.go#L41) | [ExampleTake_pairs](seq/example_test.go#L47) |
| Skip leading elements or pairs | [ExampleSkip](seq/example_test.go#L255) | [ExampleSkip](seq/example_test.go#L255) |
| Connect operations that change types | [ExampleSeq_Pipe](seq/example_test.go#L57) | [ExampleSeq2_Pipe2](seq/example_test.go#L199) |
| Connect between Seq and Seq2 | [ExampleSeq_Pipe2](seq/example_test.go#L179) | [ExampleSeq2_Pipe](seq/example_test.go#L191) |
| Extract keys / values from Seq2 | — | [ExampleKeys](seq/example_test.go#L144), [ExampleValues](seq/example_test.go#L153) |
| Directly call operators between Seq and Seq2 | [ExampleToSeq2Operator](seq/example_test.go#L211) | [ExampleToSeqOperator](seq/example_test.go#L222) |
| Connect and directly call function operators | [ExampleOperator](seq/example_test.go#L81) | [ExampleOperator2](seq/example_test.go#L128) |
| Extend with custom operators | [ExampleOperator_custom](seq/example_test.go#L111) | [ExampleOperator2](seq/example_test.go#L128) |

To run only the Examples, use the following command.  

```sh
go test ./... -run '^Example' -count=1
```

## Evaluation and usage semantics

### Lazy evaluation and early termination

Building a pipeline with built-in operators and calling `ToSeq` / `ToSeq2` do not evaluate elements.  
Input is processed in iteration order when `ToSlice` / `ToMap` is called or the iterator is executed.  

When a downstream `yield` returns `false` or a `range` loop is stopped with `break`, upstream iteration also stops.  
For example, `Take` does not execute the upstream iterator when `n <= 0`.  
Once the limit is reached, its operators exit the `range` loop with `return`, propagating the stop upstream so subsequent upstream processing also stops.  

`Skip(n)` skips up to the first `n` elements or pairs during iteration, preserving the order and duplicates of the rest.  
It skips nothing when `n <= 0` and returns an empty sequence when `n` is at least the input count.  

Custom operators must likewise process input when their returned iterator is executed and propagate downstream stops upstream.  

### Re-evaluation and input references

Whether a pipeline can be evaluated again depends on the original iterator.  
`Take` and `Skip` initialize their remaining counts for each evaluation.  
`FromSlice` / `FromMap` do not copy the input; they reference its elements at evaluation time.  
Slices are iterated in their original order, while map iteration order is not guaranteed.  
For duplicate keys, `ToMap` keeps the last value in iteration order.  

Built-in operations other than channel conversions do not use goroutines or channels.  
Callers are responsible for managing the reuse of stateful callbacks and concurrent access.  

### Nil and empty input

The zero values of `Seq` / `Seq2` and values converted from nil standard iterators are nil.  
Built-in pipeline transformations treat nil input iterators as empty.  
How a custom operator handles nil depends on its transformation function's implementation.  

`ToSeq` / `ToSeq2` return nil unchanged.  
Calling nil directly, iterating over it with `range`, or passing it to `ToSlice` / `ToMap` causes a panic.  
When collecting an empty non-nil iterator, `ToSlice` returns nil and `ToMap` returns an empty non-nil map.  

## License

[MIT License](LICENSE)  
