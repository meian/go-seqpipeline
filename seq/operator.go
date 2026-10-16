package seq

// Operator は Seq[I] を Seq[O] に変換する関数型である。
// 関数として直接呼び出すか、Seq.Pipe に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator[I, O any] func(Seq[I]) Seq[O]

// Operator12 は Seq[I] を Seq2[K, V] に変換する関数型である。
// 関数として直接呼び出すか、Seq.Pipe2 に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator12[I, K, V any] func(Seq[I]) Seq2[K, V]

// Operator21 は Seq2[K, V] を Seq[O] に変換する関数型である。
// 関数として直接呼び出すか、Seq2.Pipe に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator21[K, V, O any] func(Seq2[K, V]) Seq[O]

// Operator2 は Seq2[K, V] を Seq2[K2, V2] に変換する関数型である。
// 関数として直接呼び出すか、Seq2.Pipe2 に渡して使用する。
// 入力の評価タイミングと nil の扱いは、変換関数の実装に従う。
type Operator2[K, V, K2, V2 any] func(Seq2[K, V]) Seq2[K2, V2]
