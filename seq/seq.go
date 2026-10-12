package seq

import "iter"

// Seq は、iter.Seq に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込みオペレーター は空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq[T any] iter.Seq[T]

// Pipe は、この Seq に変換関数 op を適用した新しい Seq を返す。
// この時点では要素を評価しない。
func (s Seq[T]) Pipe[U any](op func(Seq[T]) Seq[U]) Seq[U] {
	return op(s)
}

// Seq2 は、iter.Seq2 に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込みオペレーター は空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq2[K, V any] iter.Seq2[K, V]

// Pipe は、この Seq2 に変換関数 op を適用した新しい Seq2 を返す。
// この時点では要素を評価しない。
func (s Seq2[K, V]) Pipe[K2, V2 any](op func(Seq2[K, V]) Seq2[K2, V2]) Seq2[K2, V2] {
	return op(s)
}
