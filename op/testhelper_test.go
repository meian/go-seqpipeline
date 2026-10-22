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

// 入力の開始・取得と、yield(false) による停止を記録する。
type inputTrace struct {
	starts, reads, stops int
}

func traceValues[T any](values []T, state *inputTrace) iter.Seq[T] {
	return func(yield func(T) bool) {
		state.starts++
		for _, v := range values {
			state.reads++
			if !yield(v) {
				state.stops++
				return
			}
		}
	}
}

func tracePairs[K, V any](pairs []pair[K, V], state *inputTrace) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		state.starts++
		for _, p := range pairs {
			state.reads++
			if !yield(p.Key, p.Value) {
				state.stops++
				return
			}
		}
	}
}
