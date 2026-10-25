package op

import (
	"iter"
	"slices"
	"sync"
)

// Cache は入力を初回反復時に全件評価し、その結果を繰り返し反復できる iter.Seq を返す。
// 入力の順序と重複を保ち、途中終了後も保存した先頭から反復する。
// 上流の副作用は初回反復時に発生し、以降の反復では上流を再実行しない。
// 上流が panic した場合、部分的な結果は公開せず、以降の反復でも同じ値で panic する。
//
// 入力は有限である必要があり、保存には要素数に比例するメモリを使う。
// 無限入力は評価が完了しないため、必要なら Cache より前に Take を接続する。
// 要素は値として保存し、ポインター・スライス・map などの参照先はコピーしない。
// 戻り値は交互・並行に反復できるが、共有する参照先や下流のコールバックの同期は呼び出し側が担う。
func Cache[T any](src iter.Seq[T]) iter.Seq[T] {
	collect := sync.OnceValue(func() []T {
		if src == nil {
			return nil
		}
		return slices.Collect(src)
	})
	return func(yield func(T) bool) {
		slices.Values(collect())(yield)
	}
}

// Cache2 は入力を初回反復時に全件評価し、その結果を繰り返し反復できる iter.Seq2 を返す。
// キーと値の組の順序と重複を保つ。
// 入力の有限性、保存方式、反復時の動作と同期の条件は [Cache] と同じ。
func Cache2[K, V any](src iter.Seq2[K, V]) iter.Seq2[K, V] {
	type pair struct {
		key   K
		value V
	}
	collect := sync.OnceValue(func() []pair {
		var pairs []pair
		if src != nil {
			for k, v := range src {
				pairs = append(pairs, pair{key: k, value: v})
			}
		}
		return pairs
	})
	return func(yield func(K, V) bool) {
		for _, p := range collect() {
			if !yield(p.key, p.value) {
				return
			}
		}
	}
}
