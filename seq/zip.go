package seq

import "iter"

// ZipAlign は Zip の出力件数をどちらの入力に合わせるかを表す。
type ZipAlign int

const (
	// AlignShortest は短い方の入力が終了した時点で終了する。
	// Zip の既定値である。
	AlignShortest ZipAlign = iota
	// AlignLeft は左の入力の件数に合わせ、不足する値を補う。
	AlignLeft
	// AlignRight は右の入力の件数に合わせ、不足するキーを補う。
	AlignRight
	// AlignLongest は長い方の入力の件数に合わせ、不足するキーや値を補う。
	AlignLongest
)

// Zip は右の入力 other と組み合わせる設定値を返す。
// 左の要素をキー、右の要素を値とし、既定では短い方の入力で終了する。
// AlignBy で終了条件、WithKey と WithValue で不足分の既定値を指定する。
// 未指定の既定値はそれぞれの型のゼロ値であり、AlignShortest では補完しない。
// 各設定メソッドはコピーを返す。
// Seq.Pipe2 に Zip(other).Operator を渡すと、左の入力からキー型を推論する。
func Zip[V any](other Seq[V]) zipOperator[V] {
	return zipOperator[V]{other: other}
}

type zipOperator[V any] struct {
	other     Seq[V]
	alignment ZipAlign
	value     V
}

// AlignBy は出力件数の基準を設定したコピーを返す。
// AlignShortest、AlignLeft、AlignRight、AlignLongest 以外の値は panic する。
func (z zipOperator[V]) AlignBy(alignment ZipAlign) zipOperator[V] {
	if alignment < AlignShortest || alignment > AlignLongest {
		panic("seq.Zip: invalid alignment")
	}
	z.alignment = alignment
	return z
}

// WithValue は不足する値の既定値を設定したコピーを返す。
func (z zipOperator[V]) WithValue(value V) zipOperator[V] {
	z.value = value
	return z
}

// WithKey は不足するキーの既定値を設定したコピーを返す。
// キー型は key から確定し、Operator に渡す左の入力も同じ型でなければならない。
func (z zipOperator[V]) WithKey[K any](key K) zipKeyOperator[K, V] {
	return zipKeyOperator[K, V]{config: z, key: key}
}

// Operator は左の入力 src と設定済みの右の入力を組み合わせる Seq2 を返す。
// この時点では入力を評価せず、nil の入力は空として扱う。
// AlignRight では右から、それ以外では左から要素を取得する。
// AlignShortest では右の終了検出時に左を1要素余分に取得することがある。
// 下流の途中終了や panic を含め、評価終了時には両入力の反復を停止する。
// 再評価できるかどうかは両方の元の Iterator に従う。
func (z zipOperator[V]) Operator[K any](src Seq[K]) Seq2[K, V] {
	var key K
	return zipSequences(src, z.other, z.alignment, key, z.value)
}

type zipKeyOperator[K, V any] struct {
	config zipOperator[V]
	key    K
}

// AlignBy はキーの既定値を保ち、出力件数の基準を設定したコピーを返す。
// AlignShortest、AlignLeft、AlignRight、AlignLongest 以外の値は panic する。
func (z zipKeyOperator[K, V]) AlignBy(alignment ZipAlign) zipKeyOperator[K, V] {
	z.config = z.config.AlignBy(alignment)
	return z
}

// WithValue はキーの既定値を保ち、値の既定値を設定したコピーを返す。
func (z zipKeyOperator[K, V]) WithValue(value V) zipKeyOperator[K, V] {
	z.config = z.config.WithValue(value)
	return z
}

// WithKey はキーの既定値を置き換えたコピーを返す。
func (z zipKeyOperator[K, V]) WithKey(key K) zipKeyOperator[K, V] {
	z.key = key
	return z
}

// Operator は設定済みのキーと値の既定値を使って Seq2 を返す。
// この時点では入力を評価せず、nil の入力は空として扱う。
// AlignRight では右から、それ以外では左から要素を取得する。
// AlignShortest では右の終了検出時に左を1要素余分に取得することがある。
// 下流の途中終了や panic を含め、評価終了時には両入力の反復を停止する。
// 再評価できるかどうかは両方の元の Iterator に従う。
func (z zipKeyOperator[K, V]) Operator(src Seq[K]) Seq2[K, V] {
	return zipSequences(src, z.config.other, z.config.alignment, z.key, z.config.value)
}

func zipSequences[K, V any](left Seq[K], right Seq[V], alignment ZipAlign, key K, value V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if alignment == AlignShortest && (left == nil || right == nil) || alignment == AlignLeft && left == nil || alignment == AlignRight && right == nil {
			return
		}
		keys := iter.Seq[K](left)
		if keys == nil {
			keys = func(func(K) bool) {}
		}
		values := iter.Seq[V](right)
		if values == nil {
			values = func(func(V) bool) {}
		}
		nextKey, stopKey := iter.Pull(keys)
		defer stopKey()
		nextValue, stopValue := iter.Pull(values)
		defer stopValue()
		for {
			var k K
			var v V
			var hasKey, hasValue bool
			if alignment == AlignRight {
				v, hasValue = nextValue()
				if !hasValue {
					return
				}
				k, hasKey = nextKey()
			} else {
				k, hasKey = nextKey()
				if !hasKey && (alignment == AlignShortest || alignment == AlignLeft) {
					return
				}
				v, hasValue = nextValue()
				if !hasValue && (alignment == AlignShortest || !hasKey) {
					return
				}
			}
			if !hasKey {
				k = key
			}
			if !hasValue {
				v = value
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
