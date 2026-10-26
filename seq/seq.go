package seq

import (
	"iter"

	"github.com/meian/go-seqpipeline/op"
)

// Seq is an iterator over a sequence of values, with Pipe and Pipe2 methods
// for connecting operators.
type Seq[T any] iter.Seq[T]

// Pipe applies an operator to the sequence and returns a Seq over the resulting values.
func (s Seq[T]) Pipe[U any](ope op.Operator[T, U]) Seq[U] {
	return Seq[U](ope(iter.Seq[T](s)))
}

// Pipe2 applies an operator to the sequence and returns a Seq2 over the resulting key-value pairs.
func (s Seq[T]) Pipe2[K, V any](ope op.ToSeq2Operator[T, K, V]) Seq2[K, V] {
	return Seq2[K, V](ope(iter.Seq[T](s)))
}

// Seq2 is an iterator over a sequence of key-value pairs, with Pipe and Pipe2
// methods for connecting operators.
type Seq2[K, V any] iter.Seq2[K, V]

// Pipe applies an operator to the sequence and returns a Seq over the resulting values.
func (s Seq2[K, V]) Pipe[U any](ope op.ToSeqOperator[K, V, U]) Seq[U] {
	return Seq[U](ope(iter.Seq2[K, V](s)))
}

// Pipe2 applies an operator to the sequence and returns a Seq2 over the resulting key-value pairs.
func (s Seq2[K, V]) Pipe2[K2, V2 any](ope op.Operator2[K, V, K2, V2]) Seq2[K2, V2] {
	return Seq2[K2, V2](ope(iter.Seq2[K, V](s)))
}
