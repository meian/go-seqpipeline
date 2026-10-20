package seq

import "iter"

// Operator は iter.Seq[I] を iter.Seq[O] に変換する関数型である。
// 関数として直接呼び出すか、seq.Seq.Pipe に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator[I, O any] func(iter.Seq[I]) iter.Seq[O]

// ToSeq2Operator は iter.Seq[I] を iter.Seq2[K, V] に変換する関数型である。
// 関数として直接呼び出すか、seq.Seq.Pipe2 に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type ToSeq2Operator[I, K, V any] func(iter.Seq[I]) iter.Seq2[K, V]

// ToSeqOperator は iter.Seq2[K, V] を iter.Seq[O] に変換する関数型である。
// 関数として直接呼び出すか、seq.Seq2.Pipe に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type ToSeqOperator[K, V, O any] func(iter.Seq2[K, V]) iter.Seq[O]

// Operator2 は iter.Seq2[K, V] を iter.Seq2[K2, V2] に変換する関数型である。
// 関数として直接呼び出すか、seq.Seq2.Pipe2 に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator2[K, V, K2, V2 any] func(iter.Seq2[K, V]) iter.Seq2[K2, V2]
