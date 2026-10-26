package op

import "iter"

// Filter returns an Operator that yields only elements matching predicate.
func Filter[T any](predicate func(T) bool) Operator[T, T] {
	return func(src iter.Seq[T]) iter.Seq[T] {
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

// Filter2 returns an Operator2 that yields only key-value pairs matching predicate.
func Filter2[K, V any](predicate func(K, V) bool) Operator2[K, V, K, V] {
	return func(src iter.Seq2[K, V]) iter.Seq2[K, V] {
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
