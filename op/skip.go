package op

import "iter"

// Skip は先頭から最大 n 個の要素を読み飛ばすための設定値を返す。
// n が 0 以下なら読み飛ばさず、すべての要素を返す。
// Pipe には Skip(n).Operator、Pipe2 には Skip(n).Operator2 を渡す。
func Skip(n int) skipOperator {
	return skipOperator{n: n}
}

type skipOperator struct {
	n int
}

// Operator は先頭の指定数を読み飛ばし、残りの要素を入力の順序と重複を保って返す。
// 指定数が 0 以下なら読み飛ばさず、入力件数以上なら空になる。
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

// Operator2 は先頭の指定数を読み飛ばし、残りのキーと値の組を入力の順序と重複を保って返す。
// 指定数が 0 以下なら読み飛ばさず、入力件数以上なら空になる。
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
