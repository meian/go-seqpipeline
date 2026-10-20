package op

import "iter"

// Cache は入力を呼び出し時に全件評価し、その結果を繰り返し反復できる iter.Seq を返す。
// 入力の順序と重複を保ち、途中終了後も保存した先頭から反復する。
// 上流の副作用は接続時に発生し、戻り値の反復では上流を再実行しない。
//
// 入力は有限である必要があり、保存には要素数に比例するメモリを使う。
// 無限入力は評価が完了しないため、必要なら Cache より前に Take を接続する。
// 要素は値として保存し、ポインター・スライス・map などの参照先はコピーしない。
// 戻り値は交互・並行に反復できるが、共有する参照先や下流のコールバックの同期は呼び出し側が担う。
func Cache[T any](src iter.Seq[T]) iter.Seq[T] {
	var values []T
	if src != nil {
		for v := range src {
			values = append(values, v)
		}
	}
	return func(yield func(T) bool) {
		for _, v := range values {
			if !yield(v) {
				return
			}
		}
	}
}

// Cache2 は入力を呼び出し時に全件評価し、その結果を繰り返し反復できる iter.Seq2 を返す。
// キーと値の組の順序と重複を保ち、キーに comparable を要求しない。
// 入力の有限性、保存方式、反復時の動作と同期の条件は [Cache] と同じ。
func Cache2[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
	type pair struct {
		key   K
		value V
	}
	var pairs []pair
	if src != nil {
		for k, v := range src {
			pairs = append(pairs, pair{key: k, value: v})
		}
	}
	return func(yield func(K, V) bool) {
		for _, p := range pairs {
			if !yield(p.key, p.value) {
				return
			}
		}
	}
}
