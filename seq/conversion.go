package seq

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice はスライスを、パイプラインを構築するための seq.Seq に変換する。
// 要素は元の順序で反復する。
// スライスをコピーせず、評価時点の要素を参照する。
func FromSlice[T any](values []T) Seq[T] {
	return FromSeq(slices.Values(values))
}

// ToSlice は seq.Seq を評価し、要素を反復順にスライスへ収集する。
// 空のシーケンスでは nil を返す。
// nil の入力では slices.Collect と同様に panic する。
func ToSlice[T any](s Seq[T]) []T {
	return slices.Collect(ToSeq(s))
}

// FromSeq は標準の iter.Seq を、パイプラインを構築するための seq.Seq に変換する。
func FromSeq[T any](src iter.Seq[T]) Seq[T] {
	return Seq[T](src)
}

// ToSeq は seq.Seq を標準の iter.Seq に変換する。
// この時点では要素を評価しない。
func ToSeq[T any](s Seq[T]) iter.Seq[T] {
	return iter.Seq[T](s)
}

// FromMap は map を、パイプラインを構築するための seq.Seq2 に変換する。
// キーと値を組として反復する。
// map をコピーせず、反復順序を保証しない。
func FromMap[K comparable, V any](values map[K]V) Seq2[K, V] {
	return FromSeq2(maps.All(values))
}

// ToMap は seq.Seq2 を評価し、キーと値を新しい map へ収集する。
// 同じキーでは後の値を採用する。
// 空のシーケンスでは空の非 nil の map を返す。
// nil の入力では maps.Collect と同様に panic する。
func ToMap[K comparable, V any](s Seq2[K, V]) map[K]V {
	return maps.Collect(ToSeq2(s))
}

// FromSeq2 は標準の iter.Seq2 を、パイプラインを構築するための seq.Seq2 に変換する。
func FromSeq2[K, V any](src iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](src)
}

// ToSeq2 は seq.Seq2 を標準の iter.Seq2 に変換する。
// この時点では要素を評価しない。
func ToSeq2[K, V any](s Seq2[K, V]) iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}
