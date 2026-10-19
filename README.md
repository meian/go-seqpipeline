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
`Seq[T]` / `Seq2[K, V]` extend the standard `iter.Seq` / `iter.Seq2` types with connection methods.  
For either input type, `Pipe` returns `Seq` and `Pipe2` returns `Seq2`.  
Types are inferred from connected inputs and outputs; incompatible operations cannot be connected.  
Iterate directly with `for range`, or collect results with `ToSlice` / `ToMap`.  

### Single-value pipeline

Use this example with imports for `fmt`, `strconv`, and `github.com/meian/go-seqpipeline/seq`.  

```go
pipeline := seq.FromSlice([]int{1, 22, 333, 4444}).
    Pipe(seq.Map(func(v int) string { return strconv.Itoa(v) })).
    Pipe(seq.Filter(func(v string) bool { return len(v) > 2 })).
    Pipe(seq.Take(1).Operator)
fmt.Println(seq.ToSlice(pipeline)) // [333]
```

### Key-value pipeline

Use this example with imports for `fmt`, `slices`, `strconv`, and `github.com/meian/go-seqpipeline/seq`.  

```go
pairs := seq.FromSeq2(slices.All([]string{"a", "bbb", "cc"})).
    Pipe2(seq.Map2(func(k int, v string) (string, int) {
        return strconv.Itoa(k), len(v)
    })).
    Pipe2(seq.Filter2(func(k string, v int) bool { return v > 1 })).
    Pipe2(seq.Take(1).Operator2)
for k, v := range pairs {
    fmt.Println(k, v) // 1 3
}
```

### Type inference

When arguments such as callbacks carry type information, pass the factory result directly to `Pipe` / `Pipe2`.  
Operations whose factory arguments carry no type information return a configuration value whose methods infer types from the input.  
Pass `.Operator` to `Pipe` for single values, or `.Operator2` to `Pipe2` for pairs; explicit type arguments are unnecessary.  

## Usage notes

`Cache` / `Cache2` evaluate the entire input during connection; limit infinite inputs before caching.  
See [caching precautions](doc/guide.md#reusing-evaluated-results) for details.  
When stopping channel consumption early, cancel the context passed to the conversion.  
See [channel conversions](doc/guide.md#channel-conversions) for communication waits and cancellation races.  

## Further reading

- [Usage guide](doc/guide.md): evaluation, stopping, re-evaluation, nil, conversions, connection types, and custom operators
- [GoDoc](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq): individual API contracts and boundary conditions
- [Go Examples](https://pkg.go.dev/github.com/meian/go-seqpipeline/seq#pkg-examples): executable usage with expected output ([source](seq/example_test.go))
- [Performance summary](benchmarks/REPORT.md): measured results and conditions

## License

[MIT License](LICENSE)  
