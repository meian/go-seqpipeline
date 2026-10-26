package op

import "iter"

// Operator is a function type that transforms an iter.Seq[I] into an iter.Seq[O].
// It can be passed to seq.Seq.Pipe.
type Operator[I, O any] func(iter.Seq[I]) iter.Seq[O]

// ToSeq2Operator is a function type that transforms an iter.Seq[I] into an
// iter.Seq2[K, V]. It can be passed to seq.Seq.Pipe2.
type ToSeq2Operator[I, K, V any] func(iter.Seq[I]) iter.Seq2[K, V]

// ToSeqOperator is a function type that transforms an iter.Seq2[K, V] into an
// iter.Seq[O]. It can be passed to seq.Seq2.Pipe.
type ToSeqOperator[K, V, O any] func(iter.Seq2[K, V]) iter.Seq[O]

// Operator2 is a function type that transforms an iter.Seq2[K, V] into an
// iter.Seq2[K2, V2]. It can be passed to seq.Seq2.Pipe2.
type Operator2[K, V, K2, V2 any] func(iter.Seq2[K, V]) iter.Seq2[K2, V2]
