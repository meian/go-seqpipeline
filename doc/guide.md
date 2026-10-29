# Usage guide

English | [日本語](guide.ja.md)  

This guide builds on the basic [README](../README.md) examples to show how to create pipelines for your own data.  
See the GoDoc for [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op) for individual API contracts.  

## Roles of seq, op, Seq, and Seq2

`seq` provides input and result conversions and the `seq.Seq` / `seq.Seq2` types for connecting operations.  
`op` provides operators such as transforming elements, filtering them, and limiting the number of results.  
An operator transforms `iter.Seq` / `iter.Seq2` to produce a new sequence; connect it through the `seq` methods `Pipe` / `Pipe2`.  

`seq.Seq[T]` / `seq.Seq2[K, V]` extend the function types from the [`iter`](https://pkg.go.dev/iter) package with `Pipe` / `Pipe2` methods for connecting transformations.  
Like standard iterators, the results can be iterated over directly with `for v := range pipeline` / `for k, v := range pipeline`.  
`seq.Seq2[K, V]` places no constraints on `K` or `V`, but `FromMap` / `ToMap`, which convert to and from maps, require `K` to be `comparable`.  
A `seq.Seq2` can also be converted into a single-value `seq.Seq` to continue the pipeline.  

For operations with separate functions for single values and key-value pairs, the pair version has a `2` suffix, as in `Filter` and `Filter2`.  

## Input and result collection

Functions starting with `From` create pipeline inputs from collections, standard iterators, channels, or integer ranges.  
Functions starting with `To` collect results into collections or convert them to standard iterators or channels.  

| Operation | `seq.Seq` | `seq.Seq2` |
| --- | --- | --- |
| Input from a collection | `FromSlice` | `FromMap` |
| Input from an integer range | `FromRange` | — |
| Input from a standard iterator | `FromSeq` | `FromSeq2` |
| Collect into a collection | `ToSlice` | `ToMap` (keys must be `comparable`) |
| Convert to a standard iterator | `ToSeq` | `ToSeq2` |
| Input from / output to a channel | `FromChannel` / `ToChannel` | — |

`FromRange(start, stop, step)` builds an integer sequence using an `int64` step and includes the end value when reached.  
A zero `step` panics at the call, while a `step` pointing away from the end produces an empty sequence.  
See the [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#FromRange) for boundary details.  

Use `ToSeq` / `ToSeq2` when passing a pipeline to a function that accepts the `iter.Seq` / `iter.Seq2` types.  

Compiler optimizations such as inlining depend on how inputs are produced, where sequence values are stored, and how operators connect through `Pipe` / `Pipe2`; the same operation can take more time or allocate more memory.  
For example, collecting with `ToSlice` / `ToMap` can be slower when reading a sequence from a global variable than when passing a locally created sequence directly or reusing it through a local variable.  
Generating and iterating with `FromSlice` / `FromMap` can also be affected, so measure the full operation with your Go version and actual inputs and connections.  

## Connecting operations with Pipe / Pipe2

For either `seq.Seq` or `seq.Seq2` input, use `Pipe` for single-value output and `Pipe2` for key-value output.  
Each subsequent operation accepts the preceding output type, including when element types change.  

```go
pipeline := seq.FromSlice([]int{1, 22, 333}).
    Pipe(op.Map(strconv.Itoa)).
    Pipe(op.Take(2).Operator)
fmt.Println(seq.ToSlice(pipeline)) // [1 22]
```

Pass factory results such as `Map` / `Filter` directly when their callbacks carry type information.  
For operations such as `Take` / `Skip`, whose factory arguments do not determine types, pass `.Operator` for single values or `.Operator2` for pairs to infer types at connection time.  

`Map` / `Map2` transform values without changing the sequence shape.  
Use `MapToSeq2` to produce one key-value pair per value, and `MapToSeq` to produce one value per pair.  
Both infer types from the callback and connect through `Pipe2` / `Pipe`, including in `Compose` / `Compose2`.  
Keys and values need not be `comparable`; order and duplicates are preserved.  
See [ExampleMapToSeq2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-MapToSeq2) / [ExampleMapToSeq](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-MapToSeq) for usage.  

The next example uses `Pipe2` to turn single values into key-value pairs, then `Pipe` to extract a single-value sequence.  
The additional input to `Zip` accepts `iter.Seq`; use `ToSeq` when passing a library sequence as that input.  
Configuration method names depend on the API: `Zip` uses `.Operator` with `Pipe2` even though it outputs pairs.  

```go
pairs := seq.FromSlice([]int{1, 2, 3}).
    Pipe2(op.Zip(slices.Values([]string{"a", "b", "c"})).Operator).
    Pipe2(op.Take(2).Operator2)
values := pairs.Pipe(op.Values)
fmt.Println(seq.ToSlice(values)) // [a b]
```

See [ExampleSeq_Pipe](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq.Pipe) / [ExampleSeq2_Pipe2](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq2.Pipe2) for changing element types, and [ExampleSeq_Pipe2](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq.Pipe2) / [ExampleSeq2_Pipe](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-Seq2.Pipe) for transitions between single values and pairs.  
The [op Example index](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#pkg-examples) covers transformation, filtering, limits, projection, and Zip.  

### Connection function types

Four function types represent the combinations of input and output types.  

| Input | Method | Transformation function | Output |
| --- | --- | --- | --- |
| `seq.Seq[I]` | `Pipe` | `op.Operator[I, O]` | `seq.Seq[O]` |
| `seq.Seq[I]` | `Pipe2` | `op.ToSeq2Operator[I, K, V]` | `seq.Seq2[K, V]` |
| `seq.Seq2[K, V]` | `Pipe` | `op.ToSeqOperator[K, V, O]` | `seq.Seq[O]` |
| `seq.Seq2[K, V]` | `Pipe2` | `op.Operator2[K, V, K2, V2]` | `seq.Seq2[K2, V2]` |

`Pipe` / `Pipe2` convert the input to a standard iterator, call the transformation immediately, and convert the result back to the library type.  
These type conversions preserve nil without evaluating or copying elements.  
Input evaluation and stop propagation depend on the transformation function.  

## Evaluation and stopping basics

Built-in input conversions and operators are lazy; collection or iteration starts processing.  
`Cache` / `Cache2` evaluate the entire input on the first iteration.  
Channel output starts evaluation in a goroutine when called.  
During `seq.Seq` / `seq.Seq2` iteration, a downstream `yield` returning `false` or a `range` loop ending with `break` propagates a stop upstream.  
Waits inside inputs or callbacks cannot be forcibly interrupted; callers must provide a way to stop them.  

## Reusing evaluated results with Cache

Use `Pipe(op.Cache)` / `Pipe2(op.Cache2)` to save results at a pipeline boundary and build multiple downstream pipelines from them.  
Caching evaluates upstream once on the first iteration before yielding any values, preserves order and duplicates, and returns a non-nil sequence even for nil or empty input.  
Later iterations read the saved results from the beginning, including after an early stop; downstream operations run on every iteration.  

To fix the results when this code runs, collect them explicitly and create a sequence from the saved slice.  

```go
source := seq.FromSlice([]int{1, 2, 3}).Pipe(op.Map(func(v int) int { return v * 2 }))
values := slices.Collect(seq.ToSeq(source))
saved := seq.FromSeq(slices.Values(values))
fmt.Println(seq.ToSlice(saved)) // [2 4 6]
```

`Collect` evaluates `source` immediately; iterating `saved` does not rerun it.  
Keep `values` unchanged if the saved results must stay fixed.  

See [ExampleCache](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Cache) / [ExampleCache2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Cache2) for executable usage.  

Input must be finite, and memory use grows with the number of saved elements or pairs.  
For an unbounded input, limit it before caching; a downstream limit cannot stop the initial evaluation.  

For example, consider an input that generates values indefinitely.  

```go
naturals := seq.FromSeq(func(yield func(int) bool) {
    for n := 1; ; n++ {
        if !yield(n) {
            return
        }
    }
})
```

The following pipeline can be connected, but its iteration never yields a value because `Cache` tries to evaluate the entire infinite input first.  
The downstream `Take(3)` cannot limit that evaluation.  

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

Changes to the input before the first iteration are reflected in the cache; later replacements are not.  
Elements are saved by value without a deep copy, so mutations through shared pointers, slices, or maps remain visible.  
If upstream panics during the first iteration, subsequent iterations panic with the same value without retrying upstream or yielding a partial cache.  
Upstream iteration finishes before the first cached value is yielded, and the saved storage can be reclaimed when the cache is no longer used.  
Cached sequences support interleaved and concurrent iteration, including the first iteration; callers must synchronize shared referenced data and downstream callbacks as needed.  

## Channel conversions

Channel conversions take a context to cancel send and receive waits.  
A value racing with cancellation may still be received or sent.  
See [ExampleFromChannel](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-FromChannel) for input and [ExampleToChannel](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#example-ToChannel) for output with cancellation.  

### FromChannel input

`FromChannel(ctx, ch)` receives during iteration and does not close the input channel.  
The sender is responsible for closing it.  
After the input is closed, iteration consumes any buffered values before ending.  
Stopping downstream or canceling the context does not stop external senders automatically.  
Provide a way to signal them to stop so they do not remain blocked on sends.  

### ToChannel output

`ToChannel(ctx, src)` starts evaluation in a goroutine when called and returns an unbuffered, receive-only channel.  
The goroutine closes the output when input iteration ends; receivers must not close it.  
If you stop receiving early, use a cancelable context and cancel it to release a blocked send.  
Cancellation cannot forcibly interrupt waits inside inputs or callbacks, which may need a separate stopping mechanism.  
To also interrupt upstream channel waits, pass the same context upstream.  

### Contexts that cannot be canceled

With `context.Background()` or another context whose `Done()` is nil, cancellation cannot release communication waits.  
A `FromChannel → processing → ToChannel` pipeline can finish normally without cancellation if the sender closes the input and the receiver drains the output.  
A nil channel, a nonterminating input, or stopping reception early may leave waits or goroutines running, so provide a way for processing to finish.  

## Composing operators with Compose

Use `seq.Seq` / `seq.Seq2` methods to connect operations to an input.  
To build a preset from existing operators before applying it to an input, use `op.Compose()` for single-value input or `op.Compose2()` for key-value input.  
The first `Pipe` / `Pipe2` infers the input and output types from the transformation you pass.  
The empty start value has no input type yet; the first connection returns a dedicated composition value.  

```go
transform := op.Compose().
    Pipe(op.Filter(func(v int) bool { return v%2 == 0 })).
    Pipe(op.Map(strconv.Itoa)).
    Pipe(op.Take(2).Operator)

first := seq.FromSlice([]int{1, 2, 3, 4, 6}).Pipe(transform.Operator)
second := seq.FromSlice([]int{8, 9, 10, 12}).Pipe(transform.Operator)
fmt.Println(seq.ToSlice(first))  // [2 4]
fmt.Println(seq.ToSlice(second)) // [8 10]
```

Composition values provide `Pipe` / `Pipe2`: `Pipe` adds a single-value output, and `Pipe2` adds a key-value output.  
Each step retains the original input type and connects the next transformation to the current output type, including transitions between single values and pairs.  
Pass the composition's `.Operator` method value to `seq.Seq` / `seq.Seq2` methods according to the final output type.  
It also matches the corresponding `op.Operator`, `op.ToSeq2Operator`, `op.ToSeqOperator`, or `op.Operator2` function type, so you can use it as a transformation inside another composition.  

When the first transformation carries insufficient type information, specify the missing type arguments on the first method.  
Later steps can infer types from the current output.  

```go
firstTwo := op.Compose().Pipe[int](op.Take(2).Operator)
firstTwoPairs := op.Compose2().Pipe2[int, string](op.Take(2).Operator2)
values := op.Compose2().Pipe[int, string](op.Values)
```

Composition does not call the transformation functions or evaluate input.  
Applying the result to an input calls transformations in connection order; their evaluation and stopping semantics remain unchanged.  
`Cache` / `Cache2` create a cache independently for each application; the first iteration fills it and later iterations reuse it.  
Inputs, callbacks, and additional inputs such as the right side of `Zip` still follow their own rules for reuse and shared state.  

See [ExampleCompose](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose), [ExampleCompose2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose2), and [ExampleCompose2_values](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Compose2-Values) for executable examples.  

## Implementing custom operators

Implement a custom operator when composing existing operators cannot express the operation you need.  

Implement a transformation function using `iter.Seq` / `iter.Seq2` for input and output, and pass it to `Pipe` / `Pipe2` according to the output type.  
The function should return an iterator that processes input when iterated, without evaluating input during the transformation call.  
For operations whose factory arguments carry no type information, provide a configuration value with a transformation method that has type parameters, and pass its method value to `Pipe` / `Pipe2` according to the output type.  
To combine operators within a custom implementation, use composition in `op` while keeping input and output as standard iterators.  

See [ExampleOperator_custom](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Operator-Custom) / [ExampleOperator2](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-Operator2) for usage, and [ExampleToSeq2Operator](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-ToSeq2Operator) / [ExampleToSeqOperator](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#example-ToSeqOperator) for function types that cross between single values and pairs.  

## Shared semantics in detail

### Input, order, and re-evaluation

Input and output conversions preserve iteration order and duplicates, except that map order is unspecified and `ToMap` keeps the last value for each key.  
`FromSlice` / `FromMap` reference input elements at evaluation time without copying them.  
Built-in operators other than caching initialize internal state for each evaluation; repeated or concurrent evaluation depends on the input and callbacks.  
Built-in operations other than channel conversions do not use goroutines or channels.  

### Nil and empty input

The zero values of `seq.Seq` / `seq.Seq2` are nil, and conversions to and from standard iterators preserve nil.  
Built-in operators treat nil as empty, but direct calls, iteration, and passing nil to `ToSlice` / `ToMap` panic.  
For empty non-nil iterators, `ToSlice` returns nil and `ToMap` returns an empty non-nil map.  
Evaluation, stopping, and nil handling for custom operators depend on their implementation.  
See GoDoc and the tests for [seq](../seq) / [op](../op) for each API's boundary conditions.  

### Example index

- [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#pkg-examples): input, output, and `Pipe` / `Pipe2` ([source](../seq/example_test.go))
- [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#pkg-examples): built-in, composed, and custom operators ([source](../op/example_test.go))
