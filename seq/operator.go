package seq

// Operator は Seq[I] を Seq[O] に変換する関数型である。
// 関数として直接呼び出すか、Seq.Pipe に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator[I, O any] func(Seq[I]) Seq[O]
