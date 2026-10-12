package seq

// Take は先頭から最大 n 個の要素を取得するための設定値を返す。
// n が 0 以下なら上流を実行しない。
// Pipe に渡す際は、入力から要素型を推論させるために Take(n).Operator を渡す。
func Take(n int) takeOperator {
	return takeOperator{n: n}
}

type takeOperator struct {
	n int
}

// Operator は各評価で先頭から指定数を取得する Seq を返す。
// 入力の型を推論させるため、Pipe には Operator をメソッド値として渡す。
// nil の入力は空のシーケンスとして扱う。
func (t takeOperator) Operator[T any](src Seq[T]) Seq[T] {
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

// Take2 は先頭から最大 n 個のキーと値の組を取得するための設定値を返す。
// n が 0 以下なら上流を実行しない。
// Pipe に渡す際は、入力からキーと値の型を推論させるために Take2(n).Operator を渡す。
func Take2(n int) take2Operator {
	return take2Operator{n: n}
}

type take2Operator struct {
	n int
}

// Operator は各評価で先頭から指定数のキーと値の組を取得する Seq2 を返す。
// 入力の型を推論させるため、Pipe には Operator をメソッド値として渡す。
// nil の入力は空のシーケンスとして扱う。
func (t take2Operator) Operator[K, V any](src Seq2[K, V]) Seq2[K, V] {
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
