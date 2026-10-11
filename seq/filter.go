package seq

import "iter"

// Filter は predicate が true を返す要素を保持する Operator を返す。
func Filter[T any](predicate func(T) bool) filterOperator[T] {
	return func() func(iter.Seq[T]) iter.Seq[T] {
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
}

type filterOperator[T any] func() func(iter.Seq[T]) iter.Seq[T]

// Apply は選別を遅延評価する Iterator を返す。
// nil の入力は空のシーケンスとして扱う。
func (op filterOperator[T]) Apply(src iter.Seq[T]) iter.Seq[T] {
	return op()(src)
}

// Filter2 は predicate が true を返すキーと値を保持する Operator を返す。
func Filter2[K, V any](predicate func(K, V) bool) filter2Operator[K, V] {
	return func() func(iter.Seq2[K, V]) iter.Seq2[K, V] {
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
}

type filter2Operator[K, V any] func() func(iter.Seq2[K, V]) iter.Seq2[K, V]

// Apply は選別を遅延評価する Iterator を返す。
// nil の入力は空のシーケンスとして扱う。
func (op filter2Operator[K, V]) Apply(src iter.Seq2[K, V]) iter.Seq2[K, V] {
	return op()(src)
}
