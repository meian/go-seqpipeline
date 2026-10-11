package seq

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice はスライスを元の順序で反復する Seq を返す。
// スライスをコピーせず、評価時点の要素を参照する。
func FromSlice[T any](values []T) Seq[T] {
	return FromSeq(slices.Values(values))
}

// FromSeq は Iterator を Seq に変換する。
// nil はそのまま保持する。
func FromSeq[T any](src iter.Seq[T]) Seq[T] {
	return Seq[T](src)
}

// FromSeq2 は Iterator を Seq2 に変換する。
// nil はそのまま保持する。
func FromSeq2[K, V any](src iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](src)
}

// FromMap は map のキーと値を反復する Seq2 を返す。
// map をコピーせず、反復順序を保証しない。
func FromMap[K comparable, V any](values map[K]V) Seq2[K, V] {
	return FromSeq2(maps.All(values))
}
