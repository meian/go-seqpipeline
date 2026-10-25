package op

import "iter"

// Compose returns a starting value for composing operators with single-value input.
// The first operation passed to Pipe or Pipe2 determines the input and output
// types and returns a dedicated composition value.
// Pass its Operator method to seq.Seq.Pipe or seq.Seq.Pipe2.
// If type information is insufficient, specify it as in
// Compose().Pipe[int](Take(2).Operator).
func Compose() compose {
	return compose{}
}

type compose struct{}

// Pipe starts a composition with an operation that outputs single values.
func (compose) Pipe[I, O any](ope Operator[I, O]) composed[I, O] {
	return composed[I, O]{operator: ope}
}

// Pipe2 starts a composition with an operation that outputs key-value pairs.
func (compose) Pipe2[I, K, V any](ope ToSeq2Operator[I, K, V]) composedToSeq2[I, K, V] {
	return composedToSeq2[I, K, V]{operator: ope}
}

// Compose2 returns a starting value for composing operators with key-value input.
// The first operation passed to Pipe or Pipe2 determines the input and output
// types and returns a dedicated composition value.
// Pass its Operator method to seq.Seq2.Pipe or seq.Seq2.Pipe2.
// If type information is insufficient, specify it as in
// Compose2().Pipe2[int, string](Take(2).Operator2).
func Compose2() compose2 {
	return compose2{}
}

type compose2 struct{}

// Pipe starts a composition with an operation that outputs single values.
func (compose2) Pipe[K, V, O any](ope ToSeqOperator[K, V, O]) composedToSeq[K, V, O] {
	return composedToSeq[K, V, O]{operator: ope}
}

// Pipe2 starts a composition with an operation that outputs key-value pairs.
func (compose2) Pipe2[K, V, K2, V2 any](ope Operator2[K, V, K2, V2]) composed2[K, V, K2, V2] {
	return composed2[K, V, K2, V2]{operator: ope}
}

type composed[I, O any] struct {
	operator Operator[I, O]
}

// Pipe adds an operation that outputs single values, preserving the composition's input type.
func (c composed[I, O]) Pipe[U any](ope Operator[O, U]) composed[I, U] {
	return composed[I, U]{operator: func(src iter.Seq[I]) iter.Seq[U] {
		return ope(c.operator(src))
	}}
}

// Pipe2 adds an operation that outputs key-value pairs, preserving the composition's input type.
func (c composed[I, O]) Pipe2[K, V any](ope ToSeq2Operator[O, K, V]) composedToSeq2[I, K, V] {
	return composedToSeq2[I, K, V]{operator: func(src iter.Seq[I]) iter.Seq2[K, V] {
		return ope(c.operator(src))
	}}
}

// Operator applies the composed operations to input.
func (c composed[I, O]) Operator(src iter.Seq[I]) iter.Seq[O] {
	return c.operator(src)
}

type composedToSeq2[I, K, V any] struct {
	operator ToSeq2Operator[I, K, V]
}

// Pipe adds an operation that outputs single values, preserving the composition's input type.
func (c composedToSeq2[I, K, V]) Pipe[O any](ope ToSeqOperator[K, V, O]) composed[I, O] {
	return composed[I, O]{operator: func(src iter.Seq[I]) iter.Seq[O] {
		return ope(c.operator(src))
	}}
}

// Pipe2 adds an operation that outputs key-value pairs, preserving the composition's input type.
func (c composedToSeq2[I, K, V]) Pipe2[K2, V2 any](ope Operator2[K, V, K2, V2]) composedToSeq2[I, K2, V2] {
	return composedToSeq2[I, K2, V2]{operator: func(src iter.Seq[I]) iter.Seq2[K2, V2] {
		return ope(c.operator(src))
	}}
}

// Operator applies the composed operations to input.
func (c composedToSeq2[I, K, V]) Operator(src iter.Seq[I]) iter.Seq2[K, V] {
	return c.operator(src)
}

type composedToSeq[K, V, O any] struct {
	operator ToSeqOperator[K, V, O]
}

// Pipe adds an operation that outputs single values, preserving the composition's input type.
func (c composedToSeq[K, V, O]) Pipe[U any](ope Operator[O, U]) composedToSeq[K, V, U] {
	return composedToSeq[K, V, U]{operator: func(src iter.Seq2[K, V]) iter.Seq[U] {
		return ope(c.operator(src))
	}}
}

// Pipe2 adds an operation that outputs key-value pairs, preserving the composition's input type.
func (c composedToSeq[K, V, O]) Pipe2[K2, V2 any](ope ToSeq2Operator[O, K2, V2]) composed2[K, V, K2, V2] {
	return composed2[K, V, K2, V2]{operator: func(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
		return ope(c.operator(src))
	}}
}

// Operator applies the composed operations to input.
func (c composedToSeq[K, V, O]) Operator(src iter.Seq2[K, V]) iter.Seq[O] {
	return c.operator(src)
}

type composed2[K, V, K2, V2 any] struct {
	operator Operator2[K, V, K2, V2]
}

// Pipe adds an operation that outputs single values, preserving the composition's input type.
func (c composed2[K, V, K2, V2]) Pipe[O any](ope ToSeqOperator[K2, V2, O]) composedToSeq[K, V, O] {
	return composedToSeq[K, V, O]{operator: func(src iter.Seq2[K, V]) iter.Seq[O] {
		return ope(c.operator(src))
	}}
}

// Pipe2 adds an operation that outputs key-value pairs, preserving the composition's input type.
func (c composed2[K, V, K2, V2]) Pipe2[K3, V3 any](ope Operator2[K2, V2, K3, V3]) composed2[K, V, K3, V3] {
	return composed2[K, V, K3, V3]{operator: func(src iter.Seq2[K, V]) iter.Seq2[K3, V3] {
		return ope(c.operator(src))
	}}
}

// Operator applies the composed operations to input.
func (c composed2[K, V, K2, V2]) Operator(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
	return c.operator(src)
}
