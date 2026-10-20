package op

import "iter"

// Operator は単一値のシーケンスを変換し、seq.Seq.Pipe に接続する関数型である。
type Operator[I, O any] func(iter.Seq[I]) iter.Seq[O]

// ToSeq2Operator は単一値をキーと値の組へ変換し、seq.Seq.Pipe2 に接続する関数型である。
type ToSeq2Operator[I, K, V any] func(iter.Seq[I]) iter.Seq2[K, V]

// ToSeqOperator はキーと値の組を単一値へ変換し、seq.Seq2.Pipe に接続する関数型である。
type ToSeqOperator[K, V, O any] func(iter.Seq2[K, V]) iter.Seq[O]

// Operator2 はキーと値の組のシーケンスを変換し、seq.Seq2.Pipe2 に接続する関数型である。
type Operator2[K, V, K2, V2 any] func(iter.Seq2[K, V]) iter.Seq2[K2, V2]
