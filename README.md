# go-seqpipeline

English | [日本語](README.ja.md)  

A library for building type-safe pipelines through method chaining with `iter.Seq` / `iter.Seq2`.  

## Requirements and installation

Go 1.27.1 or later is required.  

```sh
go get github.com/meian/go-seqpipeline
```

## Usage

Use functions starting with `From` to create a `seq.Seq` / `seq.Seq2`, then connect operators with `Pipe` / `Pipe2`.  
`seq.Seq[T]` / `seq.Seq2[K, V]` extend the `iter.Seq` / `iter.Seq2` types with connection methods.  
For either input type, `Pipe` returns `seq.Seq` and `Pipe2` returns `seq.Seq2`.  
Operators transform standard iterators; `Pipe` / `Pipe2` infer their input and output types.  
Iterate directly with `for range`, or collect results with `ToSlice` / `ToMap`.  

### Single-value pipeline

```go
pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
    Pipe(op.Map(func(v int) string { return strconv.Itoa(v) })).
    Pipe(op.Filter(func(v string) bool { return len(v) > 2 })).
    Pipe(op.Take(1).Operator)
fmt.Println(seq.ToSlice(pipeline)) // [333]
```

### Key-value pipeline

```go
pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
    Pipe2(op.Map2(func(k int, v string) (string, int) {
        return strconv.Itoa(k), len(v)
    })).
    Pipe2(op.Filter2(func(k string, v int) bool { return v > 1 })).
    Pipe2(op.Take(1).Operator2)
for k, v := range pairs {
    fmt.Println(k, v) // 1 3
}
```

### Type inference

When arguments such as callbacks carry type information, pass the factory result directly to `Pipe` / `Pipe2`.  
Operations whose factory arguments carry no type information return a configuration value whose methods infer types from the input.  
Pass `.Operator` to `Pipe` for single values, or `.Operator2` to `Pipe2` for pairs; explicit type arguments are unnecessary.  

## Usage notes

`Cache` / `Cache2` evaluate the entire input when applied to it; limit infinite inputs before caching.  
See [caching precautions](doc/guide.md#reusing-evaluated-results) for details.  
When stopping channel consumption early, cancel the context passed to the conversion.  
See [channel conversions](doc/guide.md#channel-conversions) for communication waits and cancellation races.  

## Further reading

- [Usage guide](doc/guide.md): evaluation, stopping, re-evaluation, nil, conversions, connection types, composition, and custom operators
- GoDoc: [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op) for individual API contracts and boundary conditions
- Go Examples: [seq](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#pkg-examples) / [op](https://pkg.go.dev/github.com/meian/go-seqpipeline/op#pkg-examples) ([seq source](seq/example_test.go), [op source](op/example_test.go))
- [Performance summary](benchmarks/REPORT.md): measured results and conditions

## License

[MIT License](LICENSE)  
