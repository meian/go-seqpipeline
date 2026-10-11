// Package seq は標準の iter を利用した遅延評価の Pipeline を提供する。
package seq

import "iter"

// Seq は、iter.Seq に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込み Operator は空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq[T any] iter.Seq[T]

// Pipe は、この Seq に factory が返す変換関数を適用した新しい Seq を返す。
// この時点では要素を評価しない。
func (s Seq[T]) Pipe[U any](factory func() func(iter.Seq[T]) iter.Seq[U]) Seq[U] {
	return FromSeq(factory()(ToSeq(s)))
}

// Seq2 は、iter.Seq2 に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込み Operator は空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq2[K, V any] iter.Seq2[K, V]

// Pipe は、この Seq2 に factory が返す変換関数を適用した新しい Seq2 を返す。
// この時点では要素を評価しない。
func (s Seq2[K, V]) Pipe[K2, V2 any](factory func() func(iter.Seq2[K, V]) iter.Seq2[K2, V2]) Seq2[K2, V2] {
	return FromSeq2(factory()(ToSeq2(s)))
}
