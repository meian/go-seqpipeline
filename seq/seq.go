package seq

import (
	"iter"

	"github.com/meian/go-seqpipeline/op"
)

// Seq は、iter.Seq に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込みオペレーターは空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、ToSlice に渡したりすると panic する。
type Seq[T any] iter.Seq[T]

// Pipe は、このシーケンスにオペレーターを適用し、単一値のシーケンスを返す。
// transform はその場で呼び出すが、入力の評価と停止伝播は transform の実装に従う。
func (s Seq[T]) Pipe[U any](transform op.Operator[T, U]) Seq[U] {
	return Seq[U](transform(iter.Seq[T](s)))
}

// Pipe2 は、このシーケンスにオペレーターを適用し、キーと値の組のシーケンスを返す。
// transform はその場で呼び出すが、入力の評価と停止伝播は transform の実装に従う。
func (s Seq[T]) Pipe2[K, V any](transform op.ToSeq2Operator[T, K, V]) Seq2[K, V] {
	return Seq2[K, V](transform(iter.Seq[T](s)))
}

// Seq2 は、iter.Seq2 に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込みオペレーターは空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq2[K, V any] iter.Seq2[K, V]

// Pipe は、このシーケンスにオペレーターを適用し、単一値のシーケンスを返す。
// transform はその場で呼び出すが、入力の評価と停止伝播は transform の実装に従う。
func (s Seq2[K, V]) Pipe[U any](transform op.ToSeqOperator[K, V, U]) Seq[U] {
	return Seq[U](transform(iter.Seq2[K, V](s)))
}

// Pipe2 は、このシーケンスにオペレーターを適用し、キーと値の組のシーケンスを返す。
// transform はその場で呼び出すが、入力の評価と停止伝播は transform の実装に従う。
func (s Seq2[K, V]) Pipe2[K2, V2 any](transform op.Operator2[K, V, K2, V2]) Seq2[K2, V2] {
	return Seq2[K2, V2](transform(iter.Seq2[K, V](s)))
}
