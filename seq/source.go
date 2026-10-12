package seq

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice はスライスを、パイプラインを構築するための Seq に変換する。
// 要素は元の順序で反復する。
// スライスをコピーせず、評価時点の要素を参照する。
func FromSlice[T any](values []T) Seq[T] {
	return FromSeq(slices.Values(values))
}

// FromSeq は標準の iter.Seq を、パイプラインを構築するための Seq に変換する。
func FromSeq[T any](src iter.Seq[T]) Seq[T] {
	return Seq[T](src)
}

// FromMap は map を、パイプラインを構築するための Seq2 に変換する。
// キーと値を組として反復する。
// map をコピーせず、反復順序を保証しない。
func FromMap[K comparable, V any](values map[K]V) Seq2[K, V] {
	return FromSeq2(maps.All(values))
}

// FromSeq2 は標準の iter.Seq2 を、パイプラインを構築するための Seq2 に変換する。
func FromSeq2[K, V any](src iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](src)
}
