package op

import (
	"fmt"
	"iter"
)

// ZipAlign specifies how Zip determines the number of output pairs.
type ZipAlign int

const (
	// AlignShortest stops when the shorter input ends.
	AlignShortest ZipAlign = iota
	// AlignLeft matches the left input length, filling missing values.
	AlignLeft
	// AlignRight matches the right input length, filling missing keys.
	AlignRight
	// AlignLongest matches the longer input length, filling missing keys or values.
	AlignLongest
)

// Zip returns a configuration value for pairing elements of two sequences in
// iteration order.
// The connected sequence supplies keys, and the sequence passed to Zip supplies values.
// Pass Operator to seq.Seq.Pipe2; by default, iteration stops with the shorter input.
// AlignBy sets the stopping condition, while WithDefaultKey and WithDefaultValue
// set values used when one input ends first.
// Unspecified defaults are the respective types' zero values.
// AlignRight reads from the right first; other alignments read from the left first.
// AlignShortest may consume one extra left element when detecting that the right
// input has ended.
// Operator infers the key type from its input iter.Seq when applied.
func Zip[V any](values iter.Seq[V]) zipOperator[V] {
	return zipOperator[V]{values: values}
}

type zipOperator[V any] struct {
	values       iter.Seq[V]
	alignment    ZipAlign
	defaultValue V
}

// AlignBy selects which input determines the output length.
// AlignBy panics if alignment is not AlignShortest, AlignLeft, AlignRight, or AlignLongest.
func (z zipOperator[V]) AlignBy(alignment ZipAlign) zipOperator[V] {
	if alignment < AlignShortest || alignment > AlignLongest {
		panic(fmt.Sprintf("op.Zip: invalid alignment: %d", alignment))
	}
	z.alignment = alignment
	return z
}

// WithDefaultValue sets the value used when the value input ends first.
func (z zipOperator[V]) WithDefaultValue(defaultValue V) zipOperator[V] {
	z.defaultValue = defaultValue
	return z
}

// WithDefaultKey sets the key used when the key input ends first.
func (z zipOperator[V]) WithDefaultKey[K any](defaultKey K) zipKeyOperator[K, V] {
	return zipKeyOperator[K, V]{zip: z, defaultKey: defaultKey}
}

// Operator pairs keys from the left input with values from the configured right
// input in iteration order.
// It stops iteration of both inputs when evaluation ends, including on an early
// downstream stop or panic.
func (z zipOperator[V]) Operator[K any](src iter.Seq[K]) iter.Seq2[K, V] {
	var defaultKey K
	return zipSequences(src, z.values, z.alignment, defaultKey, z.defaultValue)
}

// Zip の時点では未確定なキー型を、WithDefaultKey で固定するために型を分ける。
type zipKeyOperator[K, V any] struct {
	zip        zipOperator[V]
	defaultKey K
}

// AlignBy selects which input determines the output length.
// When the alignment requires padding, missing keys use the configured default
// key and missing values use the configured default value.
// AlignBy panics if alignment is not AlignShortest, AlignLeft, AlignRight, or AlignLongest.
func (z zipKeyOperator[K, V]) AlignBy(alignment ZipAlign) zipKeyOperator[K, V] {
	z.zip = z.zip.AlignBy(alignment)
	return z
}

// WithDefaultValue sets the value used when the value input ends first.
// Missing keys continue to use the default key configured by WithDefaultKey.
func (z zipKeyOperator[K, V]) WithDefaultValue(defaultValue V) zipKeyOperator[K, V] {
	z.zip = z.zip.WithDefaultValue(defaultValue)
	return z
}

// WithDefaultKey changes the key used when the key input ends first.
func (z zipKeyOperator[K, V]) WithDefaultKey(defaultKey K) zipKeyOperator[K, V] {
	z.defaultKey = defaultKey
	return z
}

// Operator pairs keys from the left input with values from the configured right
// input in iteration order.
// It fills missing keys or values with the configured defaults.
// It stops iteration of both inputs when evaluation ends, including on an early
// downstream stop or panic.
func (z zipKeyOperator[K, V]) Operator(src iter.Seq[K]) iter.Seq2[K, V] {
	return zipSequences(src, z.zip.values, z.zip.alignment, z.defaultKey, z.zip.defaultValue)
}

// Pull を片側だけに使い、切り替えの負荷を抑える。
// AlignRight は右から、それ以外は左から読む順序を保つ。
func zipSequences[K, V any](left iter.Seq[K], right iter.Seq[V], alignment ZipAlign, defaultKey K, defaultValue V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		switch alignment {
		case AlignShortest:
			zipShortest(left, right, yield)
		case AlignLeft:
			zipLeft(left, right, defaultValue, yield)
		case AlignRight:
			zipRight(left, right, defaultKey, yield)
		default:
			zipLongest(left, right, defaultKey, defaultValue, yield)
		}
	}
}

func zipShortest[K, V any](left iter.Seq[K], right iter.Seq[V], yield func(K, V) bool) {
	if left == nil || right == nil {
		return
	}
	nextValue, stopValue := iter.Pull(right)
	defer stopValue()
	for k := range left {
		v, ok := nextValue()
		if !ok || !yield(k, v) {
			return
		}
	}
}

func zipLeft[K, V any](left iter.Seq[K], right iter.Seq[V], defaultValue V, yield func(K, V) bool) {
	if left == nil {
		return
	}
	values := right
	if values == nil {
		values = func(func(V) bool) {}
	}
	nextValue, stopValue := iter.Pull(values)
	defer stopValue()
	for k := range left {
		v, ok := nextValue()
		if !ok {
			v = defaultValue
		}
		if !yield(k, v) {
			return
		}
	}
}

func zipRight[K, V any](left iter.Seq[K], right iter.Seq[V], defaultKey K, yield func(K, V) bool) {
	if right == nil {
		return
	}
	keys := left
	if keys == nil {
		keys = func(func(K) bool) {}
	}
	nextKey, stopKey := iter.Pull(keys)
	defer stopKey()
	for v := range right {
		k, ok := nextKey()
		if !ok {
			k = defaultKey
		}
		if !yield(k, v) {
			return
		}
	}
}

func zipLongest[K, V any](left iter.Seq[K], right iter.Seq[V], defaultKey K, defaultValue V, yield func(K, V) bool) {
	keys := left
	if keys == nil {
		keys = func(func(K) bool) {}
	}
	values := right
	if values == nil {
		values = func(func(V) bool) {}
	}
	nextValue, stopValue := iter.Pull(values)
	defer stopValue()
	for k := range keys {
		v, ok := nextValue()
		if !ok {
			v = defaultValue
		}
		if !yield(k, v) {
			return
		}
	}
	for {
		v, ok := nextValue()
		if !ok || !yield(defaultKey, v) {
			return
		}
	}
}
