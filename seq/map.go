package seq

import "iter"

// Map は各要素に fn を適用する Operator を返す。
func Map[I, O any](fn func(I) O) mapOperator[I, O] {
	return func() func(iter.Seq[I]) iter.Seq[O] {
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
}

type mapOperator[I, O any] func() func(iter.Seq[I]) iter.Seq[O]

// Apply は変換を遅延評価する Iterator を返す。
// nil の入力は空のシーケンスとして扱う。
func (op mapOperator[I, O]) Apply(src iter.Seq[I]) iter.Seq[O] {
	return op()(src)
}

// Map2 は各キーと値に fn を適用する Operator を返す。
func Map2[K, V, K2, V2 any](fn func(K, V) (K2, V2)) map2Operator[K, V, K2, V2] {
	return func() func(iter.Seq2[K, V]) iter.Seq2[K2, V2] {
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
}

type map2Operator[K, V, K2, V2 any] func() func(iter.Seq2[K, V]) iter.Seq2[K2, V2]

// Apply はキーと値の変換を遅延評価する Iterator を返す。
// nil の入力は空のシーケンスとして扱う。
func (op map2Operator[K, V, K2, V2]) Apply(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
	return op()(src)
}
