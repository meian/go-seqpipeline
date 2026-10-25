package op

import "iter"

// Operator transforms a single-value sequence for connection to seq.Seq.Pipe.
type Operator[I, O any] func(iter.Seq[I]) iter.Seq[O]

// ToSeq2Operator transforms single values into key-value pairs for seq.Seq.Pipe2.
type ToSeq2Operator[I, K, V any] func(iter.Seq[I]) iter.Seq2[K, V]

// ToSeqOperator transforms key-value pairs into single values for seq.Seq2.Pipe.
type ToSeqOperator[K, V, O any] func(iter.Seq2[K, V]) iter.Seq[O]

// Operator2 transforms a key-value sequence for connection to seq.Seq2.Pipe2.
type Operator2[K, V, K2, V2 any] func(iter.Seq2[K, V]) iter.Seq2[K2, V2]
