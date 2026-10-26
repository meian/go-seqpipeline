package op

import "iter"

// Skip returns a configuration value for skipping up to n initial elements.
// If n is zero or negative, it skips none and yields all elements.
// Pass Skip(n).Operator to Pipe or Skip(n).Operator2 to Pipe2.
func Skip(n int) skipOperator {
	return skipOperator{n: n}
}

type skipOperator struct {
	n int
}

// Operator skips the specified number of initial elements and yields the rest,
// preserving input order and duplicates.
// It skips none if the count is zero or negative and yields nothing if the
// count is at least the input length.
func (s skipOperator) Operator[T any](src iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if src == nil {
			return
		}
		remaining := s.n
		for v := range src {
			if remaining > 0 {
				remaining--
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Operator2 skips the specified number of initial key-value pairs and yields
// the rest, preserving input order and duplicates.
// It skips none if the count is zero or negative and yields nothing if the
// count is at least the input length.
func (s skipOperator) Operator2[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if src == nil {
			return
		}
		remaining := s.n
		for k, v := range src {
			if remaining > 0 {
				remaining--
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
