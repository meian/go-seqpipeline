package seq

// Keys は各組のキーを入力の反復順序と重複を保った Seq として返す。
// この時点では入力を評価せず、下流の途中終了は上流へ伝播する。
// nil の入力は空のシーケンスとして扱う。
// 再評価できるかどうかは元の Iterator に従う。
func (s Seq2[K, V]) Keys() Seq[K] {
	return func(yield func(K) bool) {
		if s == nil {
			return
		}
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}

// Values は各組の値を入力の反復順序と重複を保った Seq として返す。
// この時点では入力を評価せず、下流の途中終了は上流へ伝播する。
// nil の入力は空のシーケンスとして扱う。
// 再評価できるかどうかは元の Iterator に従う。
func (s Seq2[K, V]) Values() Seq[V] {
	return func(yield func(V) bool) {
		if s == nil {
			return
		}
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
