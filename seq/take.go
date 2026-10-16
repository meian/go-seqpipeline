package seq

// Take は先頭から最大 n 個の要素、またはキーと値の組を取得するための設定値を返す。
// n が 0 以下なら上流を実行しない。
// Pipe には Take(n).Operator、Pipe2 には Take(n).Operator2 を渡す。
func Take(n int) takeOperator {
	return takeOperator{n: n}
}

type takeOperator struct {
	n int
}

// Operator は各評価で先頭から指定数を取得する Seq を返す。
// 入力の型を推論させるため、Pipe には Operator をメソッド値として渡す。
// 入力は反復時に評価し、残り件数は評価ごとに初期化する。
// nil 入力は空として扱い、指定数が 0 以下なら入力を実行しない。
// 指定数に達するか下流が停止すると、上流の反復も停止する。
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

// Operator2 は各評価で先頭から指定数のキーと値の組を取得する Seq2 を返す。
// 入力の型を推論させるため、Pipe2 には Operator2 をメソッド値として渡す。
// 入力は反復時に評価し、残り件数は評価ごとに初期化する。
// nil 入力は空として扱い、指定数が 0 以下なら入力を実行しない。
// 指定数に達するか下流が停止すると、上流の反復も停止する。
func (t takeOperator) Operator2[K, V any](src Seq2[K, V]) Seq2[K, V] {
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
