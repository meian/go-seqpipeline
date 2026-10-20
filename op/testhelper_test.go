package op_test

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

// pair はテストで比較するキーと値の組を保持する。
type pair[K, V any] struct {
	Key   K
	Value V
}

// collectPairs は反復順序と重複を保ち、キーと値の組をスライスへ収集する。
func collectPairs[K, V any](src iter.Seq2[K, V]) []pair[K, V] {
	var got []pair[K, V]
	for k, v := range src {
		got = append(got, pair[K, V]{Key: k, Value: v})
	}
	return got
}
