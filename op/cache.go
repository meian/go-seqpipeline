package op

import (
	"iter"
	"slices"
	"sync"
)

// Cache returns an iter.Seq that evaluates all input on its first iteration
// and can then be iterated repeatedly over the saved results.
// It preserves input order and duplicates and restarts from the saved beginning
// even after an early stop.
// Upstream side effects occur on the first iteration; later iterations do not
// rerun upstream.
// If upstream panics, Cache yields no partial results and panics with the same
// value on subsequent iterations.
//
// Input must be finite, and storage grows with the number of elements.
// Evaluation of infinite input never completes; connect [Take] before Cache
// when a limit is needed.
// Elements are saved by value; data referenced by pointers, slices, or maps
// is not copied.
// The result supports interleaved and concurrent iteration, but callers must
// synchronize shared referenced data and downstream callbacks.
func Cache[T any](src iter.Seq[T]) iter.Seq[T] {
	collect := sync.OnceValue(func() []T {
		if src == nil {
			return nil
		}
		return slices.Collect(src)
	})
	return func(yield func(T) bool) {
		slices.Values(collect())(yield)
	}
}

// Cache2 returns an iter.Seq2 that evaluates all input on its first iteration
// and can then be iterated repeatedly over the saved results.
// It preserves key-value pair order and duplicates.
// The input finiteness, storage, iteration, and synchronization conditions are
// the same as for [Cache].
func Cache2[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
	type pair struct {
		key   K
		value V
	}
	collect := sync.OnceValue(func() []pair {
		var pairs []pair
		if src != nil {
			for k, v := range src {
				pairs = append(pairs, pair{key: k, value: v})
			}
		}
		return pairs
	})
	return func(yield func(K, V) bool) {
		for _, p := range collect() {
			if !yield(p.key, p.value) {
				return
			}
		}
	}
}
