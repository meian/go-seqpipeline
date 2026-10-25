package seq

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice converts a slice to a Seq for building a pipeline.
// Elements are iterated in their original order.
// The slice is not copied; iteration reads its elements at evaluation time.
func FromSlice[T any](values []T) Seq[T] {
	return FromSeq(slices.Values(values))
}

// ToSlice evaluates a Seq and collects its elements in iteration order.
// It returns nil for an empty sequence.
func ToSlice[T any](s Seq[T]) []T {
	return slices.Collect(ToSeq(s))
}

// FromSeq converts an iter.Seq to a Seq for building a pipeline.
func FromSeq[T any](src iter.Seq[T]) Seq[T] {
	return Seq[T](src)
}

// ToSeq converts a Seq to an iter.Seq.
func ToSeq[T any](s Seq[T]) iter.Seq[T] {
	return iter.Seq[T](s)
}

// FromMap converts a map to a Seq2 for building a pipeline.
// The map is not copied, and iteration order is unspecified.
func FromMap[K comparable, V any](values map[K]V) Seq2[K, V] {
	return FromSeq2(maps.All(values))
}

// ToMap evaluates a Seq2 and collects its key-value pairs into a new map.
// For duplicate keys, the last value wins.
// It returns an empty, non-nil map for an empty sequence.
func ToMap[K comparable, V any](s Seq2[K, V]) map[K]V {
	return maps.Collect(ToSeq2(s))
}

// FromSeq2 converts an iter.Seq2 to a Seq2 for building a pipeline.
func FromSeq2[K, V any](src iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](src)
}

// ToSeq2 converts a Seq2 to an iter.Seq2.
func ToSeq2[K, V any](s Seq2[K, V]) iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}
