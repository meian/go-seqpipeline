# Usage guide

English | [日本語](guide.ja.md)  

This guide covers shared semantics and connections beyond the basic examples in the [README](../README.md).  
See the [README](../README.md#type-inference) for basic type inference and the GoDoc for [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op) for individual API contracts.  

## seq.Seq / seq.Seq2

`seq.Seq[T]` / `seq.Seq2[K, V]` extend the function types from the [`iter`](https://pkg.go.dev/iter) package with `Pipe` / `Pipe2` methods for connecting transformations.  
Like standard iterators, the results can be iterated over directly with `for v := range pipeline` / `for k, v := range pipeline`.  
`seq.Seq2[K, V]` places no constraints on `K` or `V`, but `FromMap` / `ToMap`, which convert to and from maps, require `K` to be `comparable`.  
A `seq.Seq2` can also be converted into a single-value `seq.Seq` to continue the pipeline.  

## Evaluation and usage semantics

### Evaluation and stopping

Built-in input conversions and operators are lazy, except for caching; collection or iteration starts processing.  
`Cache` / `Cache2` evaluate the entire input when called, including when connected with `Pipe` / `Pipe2`.  
Channel output starts evaluation in a goroutine when called.  
During `seq.Seq` / `seq.Seq2` iteration, a downstream `yield` returning `false` or a `range` loop ending with `break` propagates a stop upstream.  
Waits inside inputs or callbacks cannot be forcibly interrupted; callers must provide a way to stop them.  

### Input, order, and re-evaluation

Input and output conversions preserve iteration order and duplicates, except that map order is unspecified and `ToMap` keeps the last value for each key.  
`FromSlice` / `FromMap` reference input elements at evaluation time without copying them.  
Built-in operators other than caching initialize internal state for each evaluation; repeated or concurrent evaluation depends on the input and callbacks.  
Built-in operations other than channel conversions do not use goroutines or channels.  

### Reusing evaluated results

Use `Pipe(op.Cache)` / `Pipe2(op.Cache2)` to save results at a pipeline boundary and build multiple downstream pipelines from them.  
Caching evaluates upstream once during connection, preserves order and duplicates, and returns a non-nil sequence even for nil or empty input.  
Later iterations read the saved results from the beginning, including after an early stop; downstream operations run on every iteration.  

See [ExampleCache](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Cache) for executable usage.  

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
naturals.Pipe(op.Cache).Pipe(op.Take(3).Operator)
```

Putting `Take(3)` first stops upstream after three values and creates a reusable cache.  
For `seq.Seq2`, likewise connect `Take(3).Operator2` before `Cache2`.  

```go
cached := naturals.Pipe(op.Take(3).Operator).Pipe(op.Cache)
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
fmt.Println(seq.ToSlice(cached)) // [1 2 3]
```

Elements are saved by value without a deep copy: replacing original elements does not change the cache, but mutations through shared pointers, slices, or maps remain visible.  
Input iteration ends before the caching call returns, and the saved storage can be reclaimed when the cache is no longer used.  
Cached sequences support interleaved and concurrent iteration; callers must synchronize shared referenced data and downstream callbacks as needed.  

### Nil and empty input

The zero values of `seq.Seq` / `seq.Seq2` are nil, and conversions to and from standard iterators preserve nil.  
Built-in operators treat nil as empty, but direct calls, iteration, and passing nil to `ToSlice` / `ToMap` panic.  
For empty non-nil iterators, `ToSlice` returns nil and `ToMap` returns an empty non-nil map.  
Evaluation, stopping, and nil handling for custom operators depend on their implementation.  
See GoDoc and the tests for [seq](../seq) / [op](../op) for each API's boundary conditions.  

## Input and output

Functions starting with `From` convert collections, standard iterators, or channels into pipeline inputs.  
Functions starting with `To` collect results into collections or convert them to standard iterators or channels.  

| Operation | `seq.Seq` | `seq.Seq2` |
| --- | --- | --- |
| Input from a collection | `FromSlice` | `FromMap` |
| Input from a standard iterator | `FromSeq` | `FromSeq2` |
| Collect into a collection | `ToSlice` | `ToMap` (keys must be `comparable`) |
| Convert to a standard iterator | `ToSeq` | `ToSeq2` |
| Input from / output to a channel | `FromChannel` / `ToChannel` | — |

Use `ToSeq` / `ToSeq2` when passing a pipeline to a function that accepts the `iter.Seq` / `iter.Seq2` types.  

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

In go-seqpipeline, an operator is an operation that transforms `iter.Seq` / `iter.Seq2` to produce a new sequence.  

### Built-in operators

The library provides operations such as transforming elements, filtering them, and limiting the number of results.  
For operations with separate functions for single values and key-value pairs, the library provides the pair version with a `2` suffix, as in `Filter` and `Filter2`.  

### Connecting operators

Both `seq.Seq` and `seq.Seq2` define `Pipe` / `Pipe2`, which connect operators using standard iterators and return `seq.Seq` / `seq.Seq2`, respectively.  
The function types `op.Operator`, `op.ToSeq2Operator`, `op.ToSeqOperator`, and `op.Operator2` represent the four combinations of input and output types.  

| Input | Method | Transformation function | Output |
| --- | --- | --- | --- |
| `seq.Seq[I]` | `Pipe` | `op.Operator[I, O]` | `seq.Seq[O]` |
| `seq.Seq[I]` | `Pipe2` | `op.ToSeq2Operator[I, K, V]` | `seq.Seq2[K, V]` |
| `seq.Seq2[K, V]` | `Pipe` | `op.ToSeqOperator[K, V, O]` | `seq.Seq[O]` |
| `seq.Seq2[K, V]` | `Pipe2` | `op.Operator2[K, V, K2, V2]` | `seq.Seq2[K2, V2]` |

`Pipe` / `Pipe2` convert the input to a standard iterator, call the transformation immediately, and convert the result back to the library type.  
These type conversions preserve nil without evaluating or copying elements.  
Input evaluation and stop propagation depend on the transformation function.  
Use `ToSeq` / `ToSeq2` when passing library types as additional inputs to operators such as `Zip`.  

### Composing operators

Use `op.Compose()` for single-value input or `op.Compose2()` for key-value input to assemble a transformation before choosing its input.  
The first `Pipe` / `Pipe2` infers the input and output types from the transformation you pass.  
The empty start value has no input type yet and is not an operator or an identity transformation.  

```go
transform := op.Compose().
    Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
    Pipe(op.Map(strconv.Itoa)).
    Pipe(op.Take(2).Operator)

first := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform)
second := seq.FromSlice([]int{8, 9, 10, 12}).Pipe(transform)
fmt.Println(seq.ToSlice(first))  // [2 4]
fmt.Println(seq.ToSlice(second)) // [8 10]
```

All four operator function types have `Pipe` / `Pipe2` methods: `Pipe` adds a single-value output, and `Pipe2` adds a key-value output.  
Each step retains the original input type and connects the next transformation to the current output type, including transitions between single values and pairs.  
The result is an operator function type that can be passed directly to `seq.Seq` / `seq.Seq2` methods or called with a standard iterator, without a build step.  
You can also start from a named operator, such as `op.Map(strconv.Itoa).Pipe(op.Take(2).Operator)`.  

When the first transformation carries insufficient type information, specify the missing type arguments on the first method.  
Later steps can infer types from the current output.  

```go
firstTwo := op.Compose().Pipe[int](op.Take(2).Operator)
firstTwoPairs := op.Compose2().Pipe2[int, string](op.Take(2).Operator2)
values := op.Compose2().Pipe[int, string](op.Values)
```

Composition does not call the transformation functions or evaluate input.  
Applying the result to an input calls transformations in connection order; their evaluation and stopping semantics remain unchanged.  
`Cache` / `Cache2` create a cache at application time, independently for each application, and later iterations reuse that cache.  
Inputs, callbacks, and additional inputs such as the right side of `Zip` still follow their own rules for reuse and shared state.  

See [ExampleCompose](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose), [ExampleCompose2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose2), and [ExampleCompose2_values](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose2-Values) for executable examples.  

### Implementing custom operators

Implement a transformation function using `iter.Seq` / `iter.Seq2` for input and output, and pass it to `Pipe` / `Pipe2` according to the output type.  
The function should return an iterator that processes input when executed, without evaluating input during the transformation call.  
For operations whose factory arguments carry no type information, provide a configuration value with a transformation method that has type parameters, and pass its method value to `Pipe` / `Pipe2` according to the output type.  
To combine operators within a custom implementation, use composition in `op` while keeping input and output as standard iterators.  

### Executable examples

Go Examples for [seq](../seq/example_test.go) / [op](../op/example_test.go) demonstrate operators connected through `Pipe` / `Pipe2`.  

| Example | `seq.Seq` | `seq.Seq2` |
| --- | --- | --- |
| Transform elements or keys and values | [ExampleMap](../op/example_test.go#L13) | [ExampleMap2](../op/example_test.go#L20) |
| Filter by a condition | [ExampleFilter](../op/example_test.go#L31) | [ExampleFilter2](../op/example_test.go#L38) |
| Limit results and infer types (.Operator / .Operator2) | [ExampleTake](../op/example_test.go#L47) | [ExampleTake_pairs](../op/example_test.go#L53) |
| Skip leading elements or pairs | [ExampleSkip](../op/example_test.go#L208) | [ExampleSkip](../op/example_test.go#L208) |
| Connect operations that change types | [ExampleSeq_Pipe](../seq/example_test.go#L14) | [ExampleSeq2_Pipe2](../seq/example_test.go#L60) |
| Connect between `seq.Seq` and `seq.Seq2` | [ExampleSeq_Pipe2](../seq/example_test.go#L40) | [ExampleSeq2_Pipe](../seq/example_test.go#L52) |
| Extract keys / values from `seq.Seq2` | — | [ExampleKeys](../op/example_test.go#L131), [ExampleValues](../op/example_test.go#L140) |
| Connect function operators between `seq.Seq` and `seq.Seq2` | [ExampleToSeq2Operator](../op/example_test.go#L166) | [ExampleToSeqOperator](../op/example_test.go#L176) |
| Connect function operators within the same sequence type | [ExampleOperator](../op/example_test.go#L79) | [ExampleOperator2](../op/example_test.go#L121) |
| Extend with custom operators | [ExampleOperator_custom](../op/example_test.go#L105) | [ExampleOperator2](../op/example_test.go#L121) |

To run only the Examples, use the following command.  

```sh
go test ./... -run '^Example' -count=1
```
