package op

import "iter"

// Map returns an Operator that transforms each element with fn.
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

// Map2 returns an Operator2 that transforms each key-value pair with fn.
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

// MapToSeq2 returns a ToSeq2Operator that produces one key-value pair per element.
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

// MapToSeq returns a ToSeqOperator that produces one value per key-value pair.
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
