// Package op provides operators that transform iter.Seq and iter.Seq2 values.
// Connect them through the Pipe or Pipe2 methods in package seq.
//
// Built-in operators treat nil input as empty and return a non-nil iterator,
// even when the input is empty.
// [Cache] and [Cache2] evaluate the entire input on the first iteration.
// Other built-in operators evaluate input during iteration, propagate downstream
// stops upstream, and reset their internal state for each evaluation.
// Whether evaluation can be repeated or run concurrently depends on the input
// and callbacks.
// For custom operators, the connected transformation function determines
// evaluation timing, stop propagation, and nil handling.
//
// Use Pipe or Pipe2 in package seq to connect operations to an input.
// Build reusable operator compositions with the dedicated values returned by
// [Compose] or [Compose2], add operations through their Pipe or Pipe2 methods,
// and pass their Operator method to Pipe or Pipe2 in package seq.
// Composition does not call transformation functions; applying the composition
// to input calls them in connection order.
// If the composition includes Cache or Cache2, each application creates an
// independent cache that saves its input on the first iteration.
// Sharing rules for state referenced by connected operations, such as inputs
// and callbacks, remain unchanged.
package op
