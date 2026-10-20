package op

import "iter"

// Compose は単一値を入力とする合成の開始値を返す。
// 最初の Pipe / Pipe2 に渡す処理から入出力の型を決め、専用の合成値を返す。
// 合成値の Operator を seq.Seq.Pipe / Pipe2 に渡して使用する。
// 型情報が不足する場合は Compose().Pipe[int](Take(2).Operator) のように指定する。
func Compose() compose {
	return compose{}
}

type compose struct{}

// Pipe は単一値を出力する処理から合成を開始する。
func (compose) Pipe[I, O any](ope Operator[I, O]) composed[I, O] {
	return composed[I, O]{operator: ope}
}

// Pipe2 はキーと値の組を出力する処理から合成を開始する。
func (compose) Pipe2[I, K, V any](ope ToSeq2Operator[I, K, V]) composedToSeq2[I, K, V] {
	return composedToSeq2[I, K, V]{operator: ope}
}

// Compose2 はキーと値の組を入力とする合成の開始値を返す。
// 最初の Pipe / Pipe2 に渡す処理から入出力の型を決め、専用の合成値を返す。
// 合成値の Operator を seq.Seq2.Pipe / Pipe2 に渡して使用する。
// 型情報が不足する場合は Compose2().Pipe2[int, string](Take(2).Operator2) のように指定する。
func Compose2() compose2 {
	return compose2{}
}

type compose2 struct{}

// Pipe は単一値を出力する処理から合成を開始する。
func (compose2) Pipe[K, V, O any](ope ToSeqOperator[K, V, O]) composedToSeq[K, V, O] {
	return composedToSeq[K, V, O]{operator: ope}
}

// Pipe2 はキーと値の組を出力する処理から合成を開始する。
func (compose2) Pipe2[K, V, K2, V2 any](ope Operator2[K, V, K2, V2]) composed2[K, V, K2, V2] {
	return composed2[K, V, K2, V2]{operator: ope}
}

type composed[I, O any] struct {
	operator Operator[I, O]
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (c composed[I, O]) Pipe[U any](next Operator[O, U]) composed[I, U] {
	return composed[I, U]{operator: func(src iter.Seq[I]) iter.Seq[U] {
		return next(c.operator(src))
	}}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (c composed[I, O]) Pipe2[K, V any](next ToSeq2Operator[O, K, V]) composedToSeq2[I, K, V] {
	return composedToSeq2[I, K, V]{operator: func(src iter.Seq[I]) iter.Seq2[K, V] {
		return next(c.operator(src))
	}}
}

// Operator は合成した処理を入力へ適用する。
func (c composed[I, O]) Operator(src iter.Seq[I]) iter.Seq[O] {
	return c.operator(src)
}

type composedToSeq2[I, K, V any] struct {
	operator ToSeq2Operator[I, K, V]
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (c composedToSeq2[I, K, V]) Pipe[O any](next ToSeqOperator[K, V, O]) composed[I, O] {
	return composed[I, O]{operator: func(src iter.Seq[I]) iter.Seq[O] {
		return next(c.operator(src))
	}}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (c composedToSeq2[I, K, V]) Pipe2[K2, V2 any](next Operator2[K, V, K2, V2]) composedToSeq2[I, K2, V2] {
	return composedToSeq2[I, K2, V2]{operator: func(src iter.Seq[I]) iter.Seq2[K2, V2] {
		return next(c.operator(src))
	}}
}

// Operator は合成した処理を入力へ適用する。
func (c composedToSeq2[I, K, V]) Operator(src iter.Seq[I]) iter.Seq2[K, V] {
	return c.operator(src)
}

type composedToSeq[K, V, O any] struct {
	operator ToSeqOperator[K, V, O]
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (c composedToSeq[K, V, O]) Pipe[U any](next Operator[O, U]) composedToSeq[K, V, U] {
	return composedToSeq[K, V, U]{operator: func(src iter.Seq2[K, V]) iter.Seq[U] {
		return next(c.operator(src))
	}}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (c composedToSeq[K, V, O]) Pipe2[K2, V2 any](next ToSeq2Operator[O, K2, V2]) composed2[K, V, K2, V2] {
	return composed2[K, V, K2, V2]{operator: func(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
		return next(c.operator(src))
	}}
}

// Operator は合成した処理を入力へ適用する。
func (c composedToSeq[K, V, O]) Operator(src iter.Seq2[K, V]) iter.Seq[O] {
	return c.operator(src)
}

type composed2[K, V, K2, V2 any] struct {
	operator Operator2[K, V, K2, V2]
}

// Pipe は合成全体の入力型を保ち、単一値を出力する処理を追加する。
func (c composed2[K, V, K2, V2]) Pipe[O any](next ToSeqOperator[K2, V2, O]) composedToSeq[K, V, O] {
	return composedToSeq[K, V, O]{operator: func(src iter.Seq2[K, V]) iter.Seq[O] {
		return next(c.operator(src))
	}}
}

// Pipe2 は合成全体の入力型を保ち、キーと値の組を出力する処理を追加する。
func (c composed2[K, V, K2, V2]) Pipe2[K3, V3 any](next Operator2[K2, V2, K3, V3]) composed2[K, V, K3, V3] {
	return composed2[K, V, K3, V3]{operator: func(src iter.Seq2[K, V]) iter.Seq2[K3, V3] {
		return next(c.operator(src))
	}}
}

// Operator は合成した処理を入力へ適用する。
func (c composed2[K, V, K2, V2]) Operator(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
	return c.operator(src)
}
