package seq

import "iter"

// Seq は、iter.Seq に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込みオペレーター は空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq[T any] iter.Seq[T]

// Pipe は、この Seq に Operator op を適用した新しい Seq を返す。
// op はその場で呼び出すが、入力の評価と停止伝播は op の実装に従う。
func (s Seq[T]) Pipe[U any](op Operator[T, U]) Seq[U] {
	return op(s)
}

// Pipe2 は、この Seq に Operator12 op を適用した新しい Seq2 を返す。
// op はその場で呼び出すが、入力の評価と停止伝播は op の実装に従う。
func (s Seq[T]) Pipe2[K, V any](op Operator12[T, K, V]) Seq2[K, V] {
	return op(s)
}

// Seq2 は、iter.Seq2 に Pipeline 操作を追加した関数型である。
// 再評価できるかどうかは元の Iterator に従う。
// ゼロ値は nil であり、組み込みオペレーター は空のシーケンスとして扱う。
// nil を直接呼び出したり、range で反復したり、収集関数に渡したりすると panic する。
type Seq2[K, V any] iter.Seq2[K, V]

// Pipe は、この Seq2 に Operator21 op を適用した新しい Seq を返す。
// op はその場で呼び出すが、入力の評価と停止伝播は op の実装に従う。
func (s Seq2[K, V]) Pipe[U any](op Operator21[K, V, U]) Seq[U] {
	return op(s)
}

// Pipe2 は、この Seq2 に Operator2 op を適用した新しい Seq2 を返す。
// op はその場で呼び出すが、入力の評価と停止伝播は op の実装に従う。
func (s Seq2[K, V]) Pipe2[K2, V2 any](op Operator2[K, V, K2, V2]) Seq2[K2, V2] {
	return op(s)
}
