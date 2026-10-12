package seq

// Operator は Seq の要素型 I を O に変換する。
// Apply は入力を評価せず、返す Seq の実行時に処理する。
// 下流の yield が false を返したら、上流の反復も停止すること。
type Operator[I, O any] interface {
	Apply(Seq[I]) Seq[O]
}

// OperatorFunc は Seq 変換関数を Pipe に渡せる Operator に変換する。
func OperatorFunc[I, O any](fn func(Seq[I]) Seq[O]) operatorFunc[I, O] {
	return operatorFunc[I, O](fn)
}

// operatorFunc は Pipe に直接渡せる Seq 変換関数型である。
type operatorFunc[I, O any] func(Seq[I]) Seq[O]

// Apply は変換関数を呼び出して Seq を返す。
func (f operatorFunc[I, O]) Apply(src Seq[I]) Seq[O] {
	return f(src)
}

// Operator2 は Seq2 のキー型 K と値型 V を K2 と V2 に変換する。
// Apply の遅延評価と停止伝播の契約は Operator と同じである。
type Operator2[K, V, K2, V2 any] interface {
	Apply(Seq2[K, V]) Seq2[K2, V2]
}

// Operator2Func は Seq2 変換関数を Seq2 の Pipe に渡せる Operator に変換する。
func Operator2Func[K, V, K2, V2 any](fn func(Seq2[K, V]) Seq2[K2, V2]) operator2Func[K, V, K2, V2] {
	return operator2Func[K, V, K2, V2](fn)
}

// operator2Func は Seq2 の Pipe に直接渡せる Seq2 変換関数型である。
type operator2Func[K, V, K2, V2 any] func(Seq2[K, V]) Seq2[K2, V2]

// Apply は変換関数を呼び出して Seq2 を返す。
func (f operator2Func[K, V, K2, V2]) Apply(src Seq2[K, V]) Seq2[K2, V2] {
	return f(src)
}
