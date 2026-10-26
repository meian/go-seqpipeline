package op

import "iter"

// Keys returns an iter.Seq of keys, preserving input order and duplicates.
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

// Values returns an iter.Seq of values, preserving input order and duplicates.
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
