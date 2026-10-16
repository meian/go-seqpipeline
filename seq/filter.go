package seq

// Filter は条件に一致する要素だけを取り出す Operator を返す。
// 変換関数は入力を評価せず、返すシーケンスの実行時に処理する。
// nil の入力は空のシーケンスとして扱う。
func Filter[T any](predicate func(T) bool) Operator[T, T] {
	return func(src Seq[T]) Seq[T] {
		return func(yield func(T) bool) {
			if src == nil {
				return
			}
			for v := range src {
				if predicate(v) && !yield(v) {
					return
				}
			}
		}
	}
}

// Filter2 は条件に一致するキーと値の組だけを取り出す Operator2 を返す。
// 変換関数は入力を評価せず、返すシーケンスの実行時に処理する。
// nil の入力は空のシーケンスとして扱う。
func Filter2[K, V any](predicate func(K, V) bool) Operator2[K, V, K, V] {
	return func(src Seq2[K, V]) Seq2[K, V] {
		return func(yield func(K, V) bool) {
			if src == nil {
				return
			}
			for k, v := range src {
				if predicate(k, v) && !yield(k, v) {
					return
				}
			}
		}
	}
}
