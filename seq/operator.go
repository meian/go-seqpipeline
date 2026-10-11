package seq

import "iter"

// Operator は Seq の要素型 I を O に変換する。
// Apply は入力を評価せず、返す Iterator の実行時に処理する。
// 下流の yield が false を返したら、上流の反復も停止すること。
type Operator[I, O any] interface {
	Apply(iter.Seq[I]) iter.Seq[O]
}

// Operator2 は Seq2 のキー型 K と値型 V を K2 と V2 に変換する。
// Apply の遅延評価と停止伝播の契約は Operator と同じである。
type Operator2[K, V, K2, V2 any] interface {
	Apply(iter.Seq2[K, V]) iter.Seq2[K2, V2]
}

// OperatorFunc は Iterator 変換関数を Pipe に渡せる Operator に変換する。
func OperatorFunc[I, O any](fn func(iter.Seq[I]) iter.Seq[O]) operatorFunc[I, O] {
	return func() func(iter.Seq[I]) iter.Seq[O] { return fn }
}

type operatorFunc[I, O any] func() func(iter.Seq[I]) iter.Seq[O]

// Apply は変換関数を呼び出して Iterator を返す。
func (f operatorFunc[I, O]) Apply(src iter.Seq[I]) iter.Seq[O] {
	return f()(src)
}

// Operator2Func は Iterator 変換関数を Seq2 の Pipe に渡せる Operator に変換する。
func Operator2Func[K, V, K2, V2 any](fn func(iter.Seq2[K, V]) iter.Seq2[K2, V2]) operator2Func[K, V, K2, V2] {
	return func() func(iter.Seq2[K, V]) iter.Seq2[K2, V2] { return fn }
}

type operator2Func[K, V, K2, V2 any] func() func(iter.Seq2[K, V]) iter.Seq2[K2, V2]

// Apply は変換関数を呼び出して Iterator を返す。
func (f operator2Func[K, V, K2, V2]) Apply(src iter.Seq2[K, V]) iter.Seq2[K2, V2] {
	return f()(src)
}
