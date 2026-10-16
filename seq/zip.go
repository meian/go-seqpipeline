package seq

import (
	"fmt"
	"iter"
)

// ZipAlign は Zip の合成結果の出力件数の合わせ方を指定する。
type ZipAlign int

const (
	// AlignShortest は短い方の入力が終了した時点で終了する。
	AlignShortest ZipAlign = iota
	// AlignLeft は左の入力の件数に合わせ、不足する値を補う。
	AlignLeft
	// AlignRight は右の入力の件数に合わせ、不足するキーを補う。
	AlignRight
	// AlignLongest は長い方の入力の件数に合わせ、不足するキーや値を補う。
	AlignLongest
)

// Zip は元の Seq をキー、values を値として組み合わせた Seq2 を生成するための設定値を返す。
// Operator を Seq.Pipe2 に渡して使用し、既定では短い方の入力で終了する。
// AlignBy で終了条件、WithDefaultKey と WithDefaultValue で不足分の既定値を指定する。
// 未指定の既定値はそれぞれの型のゼロ値を使う。
// 各設定メソッドはコピーを返す。
// AlignRight では右から、それ以外では左から要素を取得する。
// AlignShortest では右の終了検出時に左を1要素余分に取得することがある。
// Operator の接続時に元の Seq からキー型を推論する。
func Zip[V any](values Seq[V]) zipOperator[V] {
	return zipOperator[V]{values: values}
}

type zipOperator[V any] struct {
	values       Seq[V]
	alignment    ZipAlign
	defaultValue V
}

// AlignBy は出力件数の基準を設定したコピーを返す。
// AlignShortest、AlignLeft、AlignRight、AlignLongest 以外の値は panic する。
func (z zipOperator[V]) AlignBy(alignment ZipAlign) zipOperator[V] {
	if alignment < AlignShortest || alignment > AlignLongest {
		panic(fmt.Sprintf("seq.Zip: invalid alignment: %d", alignment))
	}
	z.alignment = alignment
	return z
}

// WithDefaultValue は不足する値の既定値を設定したコピーを返す。
func (z zipOperator[V]) WithDefaultValue(defaultValue V) zipOperator[V] {
	z.defaultValue = defaultValue
	return z
}

// WithDefaultKey は不足するキーの既定値を設定したコピーを返す。
func (z zipOperator[V]) WithDefaultKey[K any](defaultKey K) zipKeyOperator[K, V] {
	return zipKeyOperator[K, V]{zip: z, defaultKey: defaultKey}
}

// Operator は左の入力 src と設定済みの右の入力を組み合わせる Seq2 を返す。
// この時点では入力を評価せず、nil の入力は空として扱う。
// 下流の途中終了や panic を含め、評価終了時には両入力の反復を停止する。
// 再評価できるかどうかは両方の元の Iterator に従う。
func (z zipOperator[V]) Operator[K any](src Seq[K]) Seq2[K, V] {
	var defaultKey K
	return zipSequences(src, z.values, z.alignment, defaultKey, z.defaultValue)
}

// zipKeyOperator は WithDefaultKey で確定したキー型と既定値を保持する。
// Zip の時点ではキー型が未確定なため、zipOperator とは分けている。
type zipKeyOperator[K, V any] struct {
	zip        zipOperator[V]
	defaultKey K
}

// AlignBy はキーの既定値を保ち、出力件数の基準を設定したコピーを返す。
// AlignShortest、AlignLeft、AlignRight、AlignLongest 以外の値は panic する。
func (z zipKeyOperator[K, V]) AlignBy(alignment ZipAlign) zipKeyOperator[K, V] {
	z.zip = z.zip.AlignBy(alignment)
	return z
}

// WithDefaultValue はキーの既定値を保ち、値の既定値を設定したコピーを返す。
func (z zipKeyOperator[K, V]) WithDefaultValue(defaultValue V) zipKeyOperator[K, V] {
	z.zip = z.zip.WithDefaultValue(defaultValue)
	return z
}

// WithDefaultKey はキーの既定値を置き換えたコピーを返す。
func (z zipKeyOperator[K, V]) WithDefaultKey(defaultKey K) zipKeyOperator[K, V] {
	z.defaultKey = defaultKey
	return z
}

// Operator は設定済みのキーと値の既定値を使って Seq2 を返す。
// この時点では入力を評価せず、nil の入力は空として扱う。
// 下流の途中終了や panic を含め、評価終了時には両入力の反復を停止する。
// 再評価できるかどうかは両方の元の Iterator に従う。
func (z zipKeyOperator[K, V]) Operator(src Seq[K]) Seq2[K, V] {
	return zipSequences(src, z.zip.values, z.zip.alignment, z.defaultKey, z.zip.defaultValue)
}

func zipSequences[K, V any](left Seq[K], right Seq[V], alignment ZipAlign, defaultKey K, defaultValue V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if left == nil && alignment <= AlignLeft {
			return
		}
		if right == nil && (alignment == AlignShortest || alignment == AlignRight) {
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
				if !hasKey && alignment <= AlignLeft {
					return
				}
				v, hasValue = nextValue()
				if !hasValue && (alignment == AlignShortest || !hasKey) {
					return
				}
			}
			if !hasKey {
				k = defaultKey
			}
			if !hasValue {
				v = defaultValue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
