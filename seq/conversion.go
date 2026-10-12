package seq

import (
	"iter"
	"maps"
	"slices"
)

// ToSlice は Seq を評価し、要素を反復順にスライスへ収集する。
// 空のシーケンスでは nil を返す。
// nil の入力では slices.Collect と同様に panic する。
func ToSlice[T any](s Seq[T]) []T {
	return slices.Collect(ToSeq(s))
}

// ToSeq は Seq を標準の iter.Seq に変換する。
// この時点では要素を評価しない。
func ToSeq[T any](s Seq[T]) iter.Seq[T] {
	return iter.Seq[T](s)
}

// ToMap は Seq2 を評価し、キーと値を新しい map へ収集する。
// 同じキーでは後の値を採用する。
// 空のシーケンスでは空の非 nil の map を返す。
// nil の入力では maps.Collect と同様に panic する。
func ToMap[K comparable, V any](s Seq2[K, V]) map[K]V {
	return maps.Collect(ToSeq2(s))
}

// ToSeq2 は Seq2 を標準の iter.Seq2 に変換する。
// この時点では要素を評価しない。
func ToSeq2[K, V any](s Seq2[K, V]) iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}
