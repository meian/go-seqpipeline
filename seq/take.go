package seq

import "iter"

// Take は先頭から最大 n 個の要素を保持する Operator を返す。
// n が 0 以下なら上流を実行しない。
// Pipe に Take(n).Type を渡すと、入力から要素型を推論する。
func Take(n int) takeOperator {
	return takeOperator{n: n}
}

type takeOperator struct {
	n int
}

// Apply は各評価で先頭から指定数を取得する Iterator を返す。
// nil の入力は空のシーケンスとして扱う。
func (t takeOperator) Apply[T any](src iter.Seq[T]) iter.Seq[T] {
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

// Type は指定した要素型の Iterator 変換関数を返す。
// Pipe には Type を呼び出さずにメソッド値として渡し、要素型を推論させる。
func (t takeOperator) Type[T any]() func(iter.Seq[T]) iter.Seq[T] {
	return t.Apply[T]
}

// Take2 は先頭から最大 n 個のキーと値の組を保持する Operator を返す。
// n が 0 以下なら上流を実行しない。
// Pipe に Take2(n).Type を渡すと、入力からキーと値の型を推論する。
func Take2(n int) take2Operator {
	return take2Operator{n: n}
}

type take2Operator struct {
	n int
}

// Apply は各評価で先頭から指定数のキーと値の組を取得する Iterator を返す。
// nil の入力は空のシーケンスとして扱う。
func (t take2Operator) Apply[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
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

// Type は指定したキーと値の型の Iterator 変換関数を返す。
// Pipe には Type を呼び出さずにメソッド値として渡し、キーと値の型を推論させる。
func (t take2Operator) Type[K, V any]() func(iter.Seq2[K, V]) iter.Seq2[K, V] {
	return t.Apply[K, V]
}
