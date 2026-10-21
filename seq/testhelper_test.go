package seq_test

import (
	"iter"
	"reflect"
	"testing"

	"github.com/meian/go-seqpipeline/op"
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

var withLength = op.MapToSeq2(func(v string) ([]int, string) { return []int{len(v)}, v })
