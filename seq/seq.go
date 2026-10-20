package seq

import (
	"iter"

	"github.com/meian/go-seqpipeline/op"
)

// Seq は単一値のシーケンスに Pipe / Pipe2 によるオペレーター接続を提供する。
type Seq[T any] iter.Seq[T]

// Pipe は、このシーケンスにオペレーターを適用し、単一値のシーケンスを返す。
func (s Seq[T]) Pipe[U any](ope op.Operator[T, U]) Seq[U] {
	return Seq[U](ope(iter.Seq[T](s)))
}

// Pipe2 は、このシーケンスにオペレーターを適用し、キーと値の組のシーケンスを返す。
func (s Seq[T]) Pipe2[K, V any](ope op.ToSeq2Operator[T, K, V]) Seq2[K, V] {
	return Seq2[K, V](ope(iter.Seq[T](s)))
}

// Seq2 はキーと値の組のシーケンスに Pipe / Pipe2 によるオペレーター接続を提供する。
type Seq2[K, V any] iter.Seq2[K, V]

// Pipe は、このシーケンスにオペレーターを適用し、単一値のシーケンスを返す。
func (s Seq2[K, V]) Pipe[U any](ope op.ToSeqOperator[K, V, U]) Seq[U] {
	return Seq[U](ope(iter.Seq2[K, V](s)))
}

// Pipe2 は、このシーケンスにオペレーターを適用し、キーと値の組のシーケンスを返す。
func (s Seq2[K, V]) Pipe2[K2, V2 any](ope op.Operator2[K, V, K2, V2]) Seq2[K2, V2] {
	return Seq2[K2, V2](ope(iter.Seq2[K, V](s)))
}
