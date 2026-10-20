package op

import "iter"

// Keys は各組のキーを入力の反復順序と重複を保った iter.Seq として返す。
func Keys[K, V any](src iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		if src == nil {
			return
		}
		for k := range src {
			if !yield(k) {
				return
			}
		}
	}
}

// Values は各組の値を入力の反復順序と重複を保った iter.Seq として返す。
func Values[K, V any](src iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		if src == nil {
			return
		}
		for _, v := range src {
			if !yield(v) {
				return
			}
		}
	}
}
