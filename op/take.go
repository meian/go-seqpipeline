package op

import "iter"

// Take は先頭から最大 n 個の要素、またはキーと値の組を取得するための設定値を返す。
// n が 0 以下なら上流を実行しない。
// Pipe には Take(n).Operator、Pipe2 には Take(n).Operator2 を渡す。
func Take(n int) takeOperator {
	return takeOperator{n: n}
}

type takeOperator struct {
	n int
}

// Operator は先頭から最大で指定数の要素を取り出し、上流の反復を停止する。
// 指定数が 0 以下なら入力を実行しない。
func (t takeOperator) Operator[T any](src iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if src == nil || t.n <= 0 {
			return
		}
		remaining := t.n
		for v := range src {
			remaining--
			if !yield(v) || remaining == 0 {
				return
			}
		}
	}
}

// Operator2 は先頭から最大で指定数のキーと値の組を取り出し、上流の反復を停止する。
// 指定数が 0 以下なら入力を実行しない。
func (t takeOperator) Operator2[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if src == nil || t.n <= 0 {
			return
		}
		remaining := t.n
		for k, v := range src {
			remaining--
			if !yield(k, v) || remaining == 0 {
				return
			}
		}
	}
}
