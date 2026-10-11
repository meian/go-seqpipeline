package seq_test

import (
	"iter"
	"reflect"
	"testing"
)

func assertEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}

// pair と collectPairs は、Iterator の順序と重複を検証するテスト専用の補助型・関数である。
type pair[K, V any] struct {
	Key   K
	Value V
}

func collectPairs[K, V any](src iter.Seq2[K, V]) []pair[K, V] {
	var got []pair[K, V]
	for k, v := range src {
		got = append(got, pair[K, V]{Key: k, Value: v})
	}
	return got
}
