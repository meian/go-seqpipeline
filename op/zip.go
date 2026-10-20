package op

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

// Zip は2つのシーケンスの要素を反復順にキーと値の組にする設定値を返す。
// 接続先のシーケンスをキー、指定したシーケンスを値として使う。
// Operator を seq.Seq.Pipe2 に渡して使用し、既定では短い方の入力で終了する。
// AlignBy で終了条件、WithDefaultKey と WithDefaultValue で不足分の既定値を指定する。
// 未指定の既定値はそれぞれの型のゼロ値を使う。
// AlignRight では右から、それ以外では左から要素を取得する。
// AlignShortest では右の終了検出時に左を1要素余分に取得することがある。
// Operator の適用時に入力の iter.Seq からキー型を推論する。
func Zip[V any](values iter.Seq[V]) zipOperator[V] {
	return zipOperator[V]{values: values}
}

type zipOperator[V any] struct {
	values       iter.Seq[V]
	alignment    ZipAlign
	defaultValue V
}

// AlignBy は出力件数をどちらの入力に合わせるかを設定する。
// AlignShortest、AlignLeft、AlignRight、AlignLongest 以外の値は panic する。
func (z zipOperator[V]) AlignBy(alignment ZipAlign) zipOperator[V] {
	if alignment < AlignShortest || alignment > AlignLongest {
		panic(fmt.Sprintf("op.Zip: invalid alignment: %d", alignment))
	}
	z.alignment = alignment
	return z
}

// WithDefaultValue は値側の入力が先に終了したときの補完値を設定する。
func (z zipOperator[V]) WithDefaultValue(defaultValue V) zipOperator[V] {
	z.defaultValue = defaultValue
	return z
}

// WithDefaultKey はキー側の入力が先に終了したときの補完値を設定する。
func (z zipOperator[V]) WithDefaultKey[K any](defaultKey K) zipKeyOperator[K, V] {
	return zipKeyOperator[K, V]{zip: z, defaultKey: defaultKey}
}

// Operator は左の入力をキー、設定済みの右の入力を値として、反復順に組み合わせる。
// 下流の途中終了や panic を含め、評価終了時には両入力の反復を停止する。
func (z zipOperator[V]) Operator[K any](src iter.Seq[K]) iter.Seq2[K, V] {
	var defaultKey K
	return zipSequences(src, z.values, z.alignment, defaultKey, z.defaultValue)
}

// Zip の時点では未確定なキー型を、WithDefaultKey で固定するために型を分ける。
type zipKeyOperator[K, V any] struct {
	zip        zipOperator[V]
	defaultKey K
}

// AlignBy はキーの既定値を保ち、出力件数をどちらの入力に合わせるかを設定する。
// AlignShortest、AlignLeft、AlignRight、AlignLongest 以外の値は panic する。
func (z zipKeyOperator[K, V]) AlignBy(alignment ZipAlign) zipKeyOperator[K, V] {
	z.zip = z.zip.AlignBy(alignment)
	return z
}

// WithDefaultValue はキーの既定値を保ち、値側の入力が先に終了したときの補完値を設定する。
func (z zipKeyOperator[K, V]) WithDefaultValue(defaultValue V) zipKeyOperator[K, V] {
	z.zip = z.zip.WithDefaultValue(defaultValue)
	return z
}

// WithDefaultKey はキー側の入力が先に終了したときの補完値を変更する。
func (z zipKeyOperator[K, V]) WithDefaultKey(defaultKey K) zipKeyOperator[K, V] {
	z.defaultKey = defaultKey
	return z
}

// Operator は左の入力をキー、設定済みの右の入力を値として、反復順に組み合わせる。
// 不足するキーや値には設定済みの既定値を使う。
// 下流の途中終了や panic を含め、評価終了時には両入力の反復を停止する。
func (z zipKeyOperator[K, V]) Operator(src iter.Seq[K]) iter.Seq2[K, V] {
	return zipSequences(src, z.zip.values, z.zip.alignment, z.defaultKey, z.zip.defaultValue)
}

func zipSequences[K, V any](left iter.Seq[K], right iter.Seq[V], alignment ZipAlign, defaultKey K, defaultValue V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if left == nil && alignment <= AlignLeft {
			return
		}
		if right == nil && (alignment == AlignShortest || alignment == AlignRight) {
			return
		}
		keys := left
		if keys == nil {
			keys = func(func(K) bool) {}
		}
		values := right
		if values == nil {
			values = func(func(V) bool) {}
		}
		// Pull を片側だけに使い、切り替えの負荷を抑える。
		// AlignRight は右から、それ以外は左から読む順序を保つ。
		if alignment == AlignRight {
			nextKey, stopKey := iter.Pull(keys)
			defer stopKey()
			for v := range values {
				k, ok := nextKey()
				if !ok {
					k = defaultKey
				}
				if !yield(k, v) {
					return
				}
			}
			return
		}

		nextValue, stopValue := iter.Pull(values)
		defer stopValue()
		for k := range keys {
			v, ok := nextValue()
			if !ok {
				if alignment == AlignShortest {
					return
				}
				v = defaultValue
			}
			if !yield(k, v) {
				return
			}
		}
		if alignment == AlignLongest {
			for {
				v, ok := nextValue()
				if !ok || !yield(defaultKey, v) {
					return
				}
			}
		}
	}
}
