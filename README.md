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

## Evaluation and usage semantics

### Evaluation and stopping

Built-in input conversions and operators are lazy; collection or iteration starts processing.  
Channel output starts evaluation in a goroutine when called.  
During `Seq` / `Seq2` iteration, a downstream `yield` returning `false` or a `range` loop ending with `break` propagates a stop upstream.  
Waits inside inputs or callbacks cannot be forcibly interrupted; callers must provide a way to stop them.  

### Input, order, and re-evaluation

Input and output conversions preserve iteration order and duplicates, except that map order is unspecified and `ToMap` keeps the last value for each key.  
`FromSlice` / `FromMap` reference input elements at evaluation time without copying them.  
Built-in operators initialize internal state for each evaluation; repeated or concurrent evaluation depends on the input and callbacks.  
Built-in operations other than channel conversions do not use goroutines or channels.  

### Nil and empty input

The zero values of `Seq` / `Seq2` are nil, and conversions to and from standard iterators preserve nil.  
Built-in operators treat nil as empty, but direct calls, iteration, and passing nil to `ToSlice` / `ToMap` panic.  
For empty non-nil iterators, `ToSlice` returns nil and `ToMap` returns an empty non-nil map.  
Evaluation, stopping, and nil handling for custom operators depend on their implementation.  
See [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) and [tests](seq) for each API's boundary conditions.  

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

`FromChannel(ctx, ch)` receives during iteration and ends when the input closes.  
It never closes the input; callers manage stopping the sender.  
Re-evaluation consumes remaining values, and multiple receivers divide values among themselves.  
`ToChannel(ctx, src)` returns an unbuffered receive-only channel and closes it when evaluation ends.  

Both conversions accept a non-nil context to release communication waits.  
Always cancel when you stop receiving output early, and share the stop signal with upstream processing.  
See [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) for nil and cancellation race semantics, and [ExampleFromChannel](seq/example_test.go#L271) / [ExampleToChannel](seq/example_test.go#L286) for usage.  

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

## License

[MIT License](LICENSE)  
