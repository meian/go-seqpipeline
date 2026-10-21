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

// MapToSeq2 は各要素からキーと値の組を1つ生成する ToSeq2Operator を返す。
func MapToSeq2[I, K, V any](fn func(I) (K, V)) ToSeq2Operator[I, K, V] {
	return func(src iter.Seq[I]) iter.Seq2[K, V] {
		return func(yield func(K, V) bool) {
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

// MapToSeq は各キーと値の組から単一値を1つ生成する ToSeqOperator を返す。
func MapToSeq[K, V, O any](fn func(K, V) O) ToSeqOperator[K, V, O] {
	return func(src iter.Seq2[K, V]) iter.Seq[O] {
		return func(yield func(O) bool) {
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
