// Package seq provides sequences and input/output conversions built on iter.
// Connect operators from package op through [Seq.Pipe], [Seq.Pipe2],
// [Seq2.Pipe], or [Seq2.Pipe2] to build a pipeline.
// Pipe and Pipe2 call the operator when connecting it; the connected operator
// determines when input is evaluated and how stopping propagates.
// Whether evaluation can be repeated or run concurrently depends on the input
// and callbacks.
//
// [FromSeq], [FromSeq2], [ToSeq], and [ToSeq2] preserve nil without evaluating
// or copying elements.
// The zero values of [Seq] and [Seq2] are nil, which built-in operators treat
// as empty input.
// Iterating a nil sequence or passing it to [ToSlice] or [ToMap] panics.
// [ToChannel] treats a nil sequence as empty input.
package seq
