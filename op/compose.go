package op

import "iter"

// Compose は単一値を入力とする合成の開始値を返す。
// 最初の Pipe / Pipe2 に渡す処理から入出力の型を決める。
// 開始値自体は Operator ではなく、最初の接続後に Operator 関数型を得る。
// 型情報が不足する場合は Compose().Pipe[int](Take(2).Operator) のように指定する。
func Compose() composeStart {
	return composeStart{}
}

type composeStart struct{}

// Pipe は単一値を出力する処理から合成を開始する。
func (composeStart) Pipe[I, O any](ope Operator[I, O]) Operator[I, O] {
	return ope
}

// Pipe2 はキーと値の組を出力する処理から合成を開始する。
func (composeStart) Pipe2[I, K, V any](ope ToSeq2Operator[I, K, V]) ToSeq2Operator[I, K, V] {
	return ope
}

// Compose2 はキーと値の組を入力とする合成の開始値を返す。
// 開始値自体は Operator ではなく、最初の Pipe / Pipe2 に渡す処理から入出力の型を決める。
// 型情報が不足する場合は Compose2().Pipe2[int, string](Take(2).Operator2) のように指定する。
func Compose2() composeStart2 {
	return composeStart2{}
}

type composeStart2 struct{}

// Pipe は単一値を出力する処理から合成を開始する。
func (composeStart2) Pipe[K, V, O any](ope ToSeqOperator[K, V, O]) ToSeqOperator[K, V, O] {
	return ope
}

// Pipe2 はキーと値の組を出力する処理から合成を開始する。
func (composeStart2) Pipe2[K, V, K2, V2 any](ope Operator2[K, V, K2, V2]) Operator2[K, V, K2, V2] {
	return ope
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (ope Operator[I, O]) Pipe[U any](next Operator[O, U]) Operator[I, U] {
	return func(src iter.Seq[I]) iter.Seq[U] {
		return next(ope(src))
	}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (ope Operator[I, O]) Pipe2[K, V any](next ToSeq2Operator[O, K, V]) ToSeq2Operator[I, K, V] {
	return func(src iter.Seq[I]) iter.Seq2[K, V] {
		return next(ope(src))
	}
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (ope ToSeq2Operator[I, K, V]) Pipe[O any](next ToSeqOperator[K, V, O]) Operator[I, O] {
	return func(src iter.Seq[I]) iter.Seq[O] {
		return next(ope(src))
	}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (ope ToSeq2Operator[I, K, V]) Pipe2[K2, V2 any](next Operator2[K, V, K2, V2]) ToSeq2Operator[I, K2, V2] {
	return func(src iter.Seq[I]) iter.Seq2[K2, V2] {
		return next(ope(src))
	}
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (ope ToSeqOperator[K, V, O]) Pipe[U any](next Operator[O, U]) ToSeqOperator[K, V, U] {
	return func(src iter.Seq2[K, V]) iter.Seq[U] {
		return next(ope(src))
	}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (ope ToSeqOperator[K, V, O]) Pipe2[K2, V2 any](next ToSeq2Operator[O, K2, V2]) Operator2[K, V, K2, V2] {
	return func(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
		return next(ope(src))
	}
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (ope Operator2[K, V, K2, V2]) Pipe[O any](next ToSeqOperator[K2, V2, O]) ToSeqOperator[K, V, O] {
	return func(src iter.Seq2[K, V]) iter.Seq[O] {
		return next(ope(src))
	}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (ope Operator2[K, V, K2, V2]) Pipe2[K3, V3 any](next Operator2[K2, V2, K3, V3]) Operator2[K, V, K3, V3] {
	return func(src iter.Seq2[K, V]) iter.Seq2[K3, V3] {
		return next(ope(src))
	}
}
