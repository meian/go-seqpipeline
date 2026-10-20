package op

import "iter"

// Map は各要素を指定した関数で変換する Operator を返す。
func Map[I, O any](fn func(I) O) Operator[I, O] {
	return func(src iter.Seq[I]) iter.Seq[O] {
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

// Map2 は各キーと値の組を指定した関数で変換する Operator2 を返す。
func Map2[K, V, K2, V2 any](fn func(K, V) (K2, V2)) Operator2[K, V, K2, V2] {
	return func(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
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
