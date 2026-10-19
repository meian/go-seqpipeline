# Usage guide

English | [日本語](guide.ja.md)  

This guide covers shared semantics and connections beyond the basic examples in the [README](../README.md).  
See the [README](../README.md#type-inference) for basic type inference and [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) for individual API contracts.  

## Seq / Seq2

`Seq[T]` / `Seq2[K, V]` extend the function types from Go's standard [`iter`](https://pkg.go.dev/iter) package with `Pipe` / `Pipe2` methods for connecting transformations.  
Like standard iterators, the results can be iterated over directly with `for v := range pipeline` / `for k, v := range pipeline`.  
`Seq2[K, V]` places no constraints on `K` or `V`, but `FromMap` / `ToMap`, which convert to and from maps, require `K` to be `comparable`.  
A `Seq2` can also be converted into a single-value `Seq` to continue the pipeline.  

## Evaluation and usage semantics

### Evaluation and stopping

Built-in input conversions and operators are lazy, except for caching; collection or iteration starts processing.  
`Cache` / `Cache2` evaluate the entire input when called, including when connected with `Pipe` / `Pipe2`.  
Channel output starts evaluation in a goroutine when called.  
During `Seq` / `Seq2` iteration, a downstream `yield` returning `false` or a `range` loop ending with `break` propagates a stop upstream.  
Waits inside inputs or callbacks cannot be forcibly interrupted; callers must provide a way to stop them.  

### Input, order, and re-evaluation

Input and output conversions preserve iteration order and duplicates, except that map order is unspecified and `ToMap` keeps the last value for each key.  
`FromSlice` / `FromMap` reference input elements at evaluation time without copying them.  
Built-in operators other than caching initialize internal state for each evaluation; repeated or concurrent evaluation depends on the input and callbacks.  
Built-in operations other than channel conversions do not use goroutines or channels.  

### Reusing evaluated results

Use `Pipe(seq.Cache)` / `Pipe2(seq.Cache2)` to save results at a pipeline boundary and build multiple downstream pipelines from them.  
Caching evaluates upstream once during connection, preserves order and duplicates, and returns a non-nil sequence even for nil or empty input.  
Later iterations read the saved results from the beginning, including after an early stop; downstream operations run on every iteration.  

See [ExampleCache](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Cache) for executable usage.  

Input must be finite, and memory use grows with the number of saved elements or pairs.  
For an unbounded input, limit it before caching; a downstream limit cannot stop the initial evaluation.  

For example, consider an input that generates values indefinitely.  

```go
naturals := seq.Seq[int](func(yield func(int) bool) {
    for n := 1; ; n++ {
        if !yield(n) {
            return
        }
    }
})
```

The following connection never returns because `Cache` tries to evaluate the entire infinite input.  
The downstream `Take(3)` is never reached.  

```go
naturals.Pipe(seq.Cache).Pipe(seq.Take(3).Operator)
```

Putting `Take(3)` first stops upstream after three values and creates a reusable cache.  
For `Seq2`, likewise connect `Take(3).Operator2` before `Cache2`.  

```go
cached := naturals.Pipe(seq.Take(3).Operator).Pipe(seq.Cache)
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
```

Elements are saved by value without a deep copy: replacing original elements does not change the cache, but mutations through shared pointers, slices, or maps remain visible.  
Input iteration ends before the caching call returns, and the saved storage can be reclaimed when the cache is no longer used.  
Cached sequences support interleaved and concurrent iteration; callers must synchronize shared referenced data and downstream callbacks as needed.  

### Nil and empty input

The zero values of `Seq` / `Seq2` are nil, and conversions to and from standard iterators preserve nil.  
Built-in operators treat nil as empty, but direct calls, iteration, and passing nil to `ToSlice` / `ToMap` panic.  
For empty non-nil iterators, `ToSlice` returns nil and `ToMap` returns an empty non-nil map.  
Evaluation, stopping, and nil handling for custom operators depend on their implementation.  
See [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) and [tests](../seq) for each API's boundary conditions.  

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

Functions that convert to and from channels take a context to support early cancellation, including channel communication waits.  
To stop processing early, cancel the context passed to the conversion.  

See [ExampleToChannel](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-ToChannel) for usage with cancellation.  

Cancellation releases channel communication waits, but cannot forcibly interrupt waits inside inputs or callbacks.  
However, if a send races with cancellation, the send may be selected, so you may receive the next value depending on timing.  

`FromChannel(ctx, ch)` receives during iteration without closing the input.  
`ToChannel(ctx, src)` returns an unbuffered channel and closes it when evaluation ends.  
To also interrupt upstream communication waits, pass the same context upstream.  

## Operators

In go-seqpipeline, an operator is an operation that transforms `Seq` / `Seq2` to produce a new sequence.  

### Built-in operators

The library provides operations such as transforming elements, filtering them, and limiting the number of results.  
For operations with separate functions for single values and key-value pairs, the library provides the pair version with a `2` suffix, as in `Filter` and `Filter2`.  

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

[Go Examples](../seq/example_test.go) demonstrate each operator with expected output.  

| Example | Seq | Seq2 |
| --- | --- | --- |
| Transform elements or keys and values | [ExampleMap](../seq/example_test.go#L12) | [ExampleMap2](../seq/example_test.go#L18) |
| Filter by a condition | [ExampleFilter](../seq/example_test.go#L28) | [ExampleFilter2](../seq/example_test.go#L34) |
| Limit results and infer types (.Operator / .Operator2) | [ExampleTake](../seq/example_test.go#L42) | [ExampleTake_pairs](../seq/example_test.go#L48) |
| Skip leading elements or pairs | [ExampleSkip](../seq/example_test.go#L256) | [ExampleSkip](../seq/example_test.go#L256) |
| Connect operations that change types | [ExampleSeq_Pipe](../seq/example_test.go#L58) | [ExampleSeq2_Pipe2](../seq/example_test.go#L200) |
| Connect between Seq and Seq2 | [ExampleSeq_Pipe2](../seq/example_test.go#L180) | [ExampleSeq2_Pipe](../seq/example_test.go#L192) |
| Extract keys / values from Seq2 | — | [ExampleKeys](../seq/example_test.go#L145), [ExampleValues](../seq/example_test.go#L154) |
| Directly call operators between Seq and Seq2 | [ExampleToSeq2Operator](../seq/example_test.go#L212) | [ExampleToSeqOperator](../seq/example_test.go#L223) |
| Connect and directly call function operators | [ExampleOperator](../seq/example_test.go#L82) | [ExampleOperator2](../seq/example_test.go#L129) |
| Extend with custom operators | [ExampleOperator_custom](../seq/example_test.go#L112) | [ExampleOperator2](../seq/example_test.go#L129) |

To run only the Examples, use the following command.  

```sh
go test ./... -run '^Example' -count=1
```
