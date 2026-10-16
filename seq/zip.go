package seq

import "iter"

// ZipDefaults は ZipAll で不足するキーと値を補う設定値である。
// ゼロ値はキーと値それぞれの型のゼロ値を使う。
type ZipDefaults[K, V any] struct {
	key   K
	value V
}

// WithKey は不足するキーの既定値を設定したコピーを返す。
// 値側の設定は保持する。
func (d ZipDefaults[K, V]) WithKey(key K) ZipDefaults[K, V] {
	d.key = key
	return d
}

// WithValue は不足する値の既定値を設定したコピーを返す。
// キー側の設定は保持する。
func (d ZipDefaults[K, V]) WithValue(value V) ZipDefaults[K, V] {
	d.value = value
	return d
}

// Zip は入力の反復順序に沿って、この Seq の要素をキー、other の要素を値とする Seq2 を返す。
// 短い方の入力が終了した時点で終了し、nil の入力は空として扱う。
// この時点では入力を評価しない。
// 下流の途中終了や入力の終了時には両方の反復を停止し、入力の defer を実行する。
// キー側を先に取得するため、値側の終了を検出する際にキー側を1要素余分に取得することがある。
// 再評価できるかどうかは元の Iterator に従う。
func (s Seq[K]) Zip[V any](other Seq[V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if s == nil || other == nil {
			return
		}
		nextKey, stopKey := iter.Pull(iter.Seq[K](s))
		defer stopKey()
		nextValue, stopValue := iter.Pull(iter.Seq[V](other))
		defer stopValue()
		for {
			k, ok := nextKey()
			if !ok {
				return
			}
			v, ok := nextValue()
			if !ok || !yield(k, v) {
				return
			}
		}
	}
}

// ZipAll は入力の反復順序に沿って、この Seq の要素をキー、other の要素を値とする Seq2 を返す。
// 長い方の入力が終了するまで反復し、nil の入力は空として扱う。
// 不足分は defaults で補い、省略時はそれぞれの型のゼロ値を使う。
// defaults を複数渡した場合は最後の設定値を使う。
// この時点では入力を評価せず、キー側、値側の順に取得する。
// 下流の途中終了や入力の終了時には両方の反復を停止し、入力の defer を実行する。
// 再評価できるかどうかは元の Iterator に従う。
func (s Seq[K]) ZipAll[V any](other Seq[V], defaults ...ZipDefaults[K, V]) Seq2[K, V] {
	var d ZipDefaults[K, V]
	if len(defaults) > 0 {
		d = defaults[len(defaults)-1]
	}
	return func(yield func(K, V) bool) {
		keys := iter.Seq[K](s)
		if keys == nil {
			keys = func(func(K) bool) {}
		}
		values := iter.Seq[V](other)
		if values == nil {
			values = func(func(V) bool) {}
		}
		nextKey, stopKey := iter.Pull(keys)
		defer stopKey()
		nextValue, stopValue := iter.Pull(values)
		defer stopValue()
		for {
			k, hasKey := nextKey()
			v, hasValue := nextValue()
			if !hasKey && !hasValue {
				return
			}
			if !hasKey {
				k = d.key
			}
			if !hasValue {
				v = d.value
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
