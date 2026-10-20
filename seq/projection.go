package seq

import "iter"

// Keys は各組のキーを入力の反復順序と重複を保った iter.Seq として返す。
// Seq2.Pipe に渡すか、関数として直接呼び出して使用する。
// 入力は反復時に評価し、下流の途中終了は上流へ伝播する。
// nil の入力は空のシーケンスとして扱う。
// 再評価できるかどうかは元の Iterator に従う。
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
// Seq2.Pipe に渡すか、関数として直接呼び出して使用する。
// 入力は反復時に評価し、下流の途中終了は上流へ伝播する。
// nil の入力は空のシーケンスとして扱う。
// 再評価できるかどうかは元の Iterator に従う。
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
