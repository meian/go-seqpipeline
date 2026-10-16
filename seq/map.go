package seq

// Map は各要素に fn を適用する Operator を返す。
// 変換関数は入力を評価せず、返すシーケンスの実行時に処理する。
// nil の入力は空のシーケンスとして扱う。
func Map[I, O any](fn func(I) O) Operator[I, O] {
	return func(src Seq[I]) Seq[O] {
		return func(yield func(O) bool) {
			if src == nil {
				return
			}
			for v := range src {
				if !yield(fn(v)) {
					return
				}
			}
		}
	}
}

// Map2 は各キーと値に fn を適用する Operator2 を返す。
// 変換関数は入力を評価せず、返すシーケンスの実行時に処理する。
// nil の入力は空のシーケンスとして扱う。
func Map2[K, V, K2, V2 any](fn func(K, V) (K2, V2)) Operator2[K, V, K2, V2] {
	return func(src Seq2[K, V]) Seq2[K2, V2] {
		return func(yield func(K2, V2) bool) {
			if src == nil {
				return
			}
			for k, v := range src {
				if !yield(fn(k, v)) {
					return
				}
			}
		}
	}
}
