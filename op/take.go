package op

import "iter"

// Take returns a configuration value for yielding up to n initial elements
// or key-value pairs.
// If n is zero or negative, it does not run upstream.
// Pass Take(n).Operator to Pipe or Take(n).Operator2 to Pipe2.
func Take(n int) takeOperator {
	return takeOperator{n: n}
}

type takeOperator struct {
	n int
}

// Operator yields at most the specified number of initial elements and stops
// upstream iteration at that limit.
// It does not run the input if the count is zero or negative.
func (t takeOperator) Operator[T any](src iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if src == nil || t.n <= 0 {
			return
		}
		remaining := t.n
		for v := range src {
			remaining--
			if !yield(v) || remaining == 0 {
				return
			}
		}
	}
}

// Operator2 yields at most the specified number of initial key-value pairs and
// stops upstream iteration at that limit.
// It does not run the input if the count is zero or negative.
func (t takeOperator) Operator2[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if src == nil || t.n <= 0 {
			return
		}
		remaining := t.n
		for k, v := range src {
			remaining--
			if !yield(k, v) || remaining == 0 {
				return
			}
		}
	}
}
