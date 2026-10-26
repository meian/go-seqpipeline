package op_test

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

var withLength = op.MapToSeq2(func(v string) ([]int, string) { return []int{len(v)}, v })

func trackTransform[I, O any](calls *[]string, name string, transform func(I) O) func(I) O {
	return func(src I) O {
		*calls = append(*calls, name)
		return transform(src)
	}
}

func TestComposePipe(t *testing.T) {
	t.Parallel()
	var calls []string
	transform := op.Compose().Pipe(trackTransform(&calls, "map", op.Map(strconv.Itoa)))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.Values([]int{2, 2, 1}))
	assertEqual(t, calls, []string{"map"})
	assertEqual(t, slices.Collect(result), []string{"2", "2", "1"})
	assertEqual(t, slices.Collect(transform.Operator(nil)), []string(nil))
	assertEqual(t, slices.Collect(transform.Operator(slices.Values([]int{}))), []string(nil))

	explicit := op.Compose().Pipe[int](op.Take(2).Operator)
	assertEqual(t, slices.Collect(explicit.Operator(slices.Values([]int{3, 2, 1}))), []int{3, 2})
}

func TestComposePipe2(t *testing.T) {
	t.Parallel()
	var calls []string
	zip := op.Zip(slices.Values([]string{"a", "b"}))
	transform := op.Compose().Pipe2(trackTransform(&calls, "zip", zip.Operator[int]))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.Values([]int{2, 2, 1}))
	assertEqual(t, calls, []string{"zip"})
	want := []pair[int, string]{{Key: 2, Value: "a"}, {Key: 2, Value: "b"}}
	assertEqual(t, collectPairs(result), want)
	assertEqual(t, collectPairs(transform.Operator(nil)), []pair[int, string](nil))
	assertEqual(t, collectPairs(transform.Operator(slices.Values([]int{}))), []pair[int, string](nil))

	explicit := op.Compose().Pipe2[int](zip.Operator)
	assertEqual(t, collectPairs(explicit.Operator(slices.Values([]int{2, 2}))), want)
}

func TestCompose2Pipe(t *testing.T) {
	t.Parallel()
	var calls []string
	transform := op.Compose2().Pipe(trackTransform(&calls, "values", op.Values[int, string]))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.All([]string{"b", "b", "a"}))
	assertEqual(t, calls, []string{"values"})
	assertEqual(t, slices.Collect(result), []string{"b", "b", "a"})
	assertEqual(t, slices.Collect(transform.Operator(nil)), []string(nil))
	assertEqual(t, slices.Collect(transform.Operator(slices.All([]string{}))), []string(nil))

	explicit := op.Compose2().Pipe[int, string](op.Values)
	assertEqual(t, slices.Collect(explicit.Operator(slices.All([]string{"a"}))), []string{"a"})
}

func TestCompose2Pipe2(t *testing.T) {
	t.Parallel()
	var calls []string
	lengths := op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	transform := op.Compose2().Pipe2(trackTransform(&calls, "map", lengths))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.All([]string{"a", "bbb"}))
	assertEqual(t, calls, []string{"map"})
	want := []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(result), want)
	assertEqual(t, collectPairs(transform.Operator(nil)), []pair[string, int](nil))
	assertEqual(t, collectPairs(transform.Operator(slices.All([]string{}))), []pair[string, int](nil))

	explicit := op.Compose2().Pipe2[int, string](op.Take(1).Operator2)
	first := []pair[int, string]{{Key: 0, Value: "a"}}
	assertEqual(t, collectPairs(explicit.Operator(slices.All([]string{"a", "b"}))), first)
}

// 各合成メソッドで、構築中は変換関数を呼ばず、適用時に接続順に呼ぶことを確認する。
// 要素・キー・値を異なる型へ変換し、合成全体の入力型と途中の出力型の接続も検証する。
func TestComposedPipe(t *testing.T) {
	t.Parallel()
	var calls []string
	first := op.Compose().Pipe(trackTransform(&calls, "first", op.Map(strconv.Itoa)))
	next := op.Map(func(v string) []byte { return []byte(v + "!") })
	transform := first.Pipe(trackTransform(&calls, "next", next))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.Values([]int{2, 2, 1}))
	assertEqual(t, calls, []string{"first", "next"})
	assertEqual(t, slices.Collect(result), [][]byte{[]byte("2!"), []byte("2!"), []byte("1!")})
	assertEqual(t, slices.Collect(transform.Operator(nil)), [][]byte(nil))
}

func TestComposedPipe2(t *testing.T) {
	t.Parallel()
	var calls []string
	first := op.Compose().Pipe(trackTransform(&calls, "first", op.Map(strconv.Itoa)))
	transform := first.Pipe2(trackTransform(&calls, "next", withLength))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.Values([]int{2, 22}))
	assertEqual(t, calls, []string{"first", "next"})
	want := []pair[[]int, string]{{Key: []int{1}, Value: "2"}, {Key: []int{2}, Value: "22"}}
	assertEqual(t, collectPairs(result), want)
	assertEqual(t, collectPairs(transform.Operator(nil)), []pair[[]int, string](nil))
}

func TestComposedOperator(t *testing.T) {
	t.Parallel()
	calls := 0
	composed := op.Compose().Pipe(func(src iter.Seq[int]) iter.Seq[string] {
		calls++
		if src == nil {
			return nil
		}
		return op.Map(strconv.Itoa)(src)
	})
	var operator op.Operator[int, string] = composed.Operator
	assertEqual(t, calls, 0)
	if operator(nil) != nil {
		t.Fatal("expected nil result")
	}
	result := operator(slices.Values([]int{12, 3}))
	assertEqual(t, calls, 2)
	assertEqual(t, slices.Collect(result), []string{"12", "3"})
}

func TestComposedToSeq2Pipe(t *testing.T) {
	t.Parallel()
	var calls []string
	first := op.Compose().Pipe2(trackTransform(&calls, "first", withLength))
	transform := first.Pipe(trackTransform(&calls, "next", op.Keys[[]int, string]))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.Values([]string{"a", "bbb", "a"}))
	assertEqual(t, calls, []string{"first", "next"})
	assertEqual(t, slices.Collect(result), [][]int{{1}, {3}, {1}})
	assertEqual(t, slices.Collect(transform.Operator(nil)), [][]int(nil))
}

func TestComposedToSeq2Pipe2(t *testing.T) {
	t.Parallel()
	var calls []string
	first := op.Compose().Pipe2(trackTransform(&calls, "first", withLength))
	next := op.Map2(func(k []int, v string) (string, int) { return v, k[0] })
	transform := first.Pipe2(trackTransform(&calls, "next", next))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.Values([]string{"a", "bbb", "a"}))
	assertEqual(t, calls, []string{"first", "next"})
	want := []pair[string, int]{{Key: "a", Value: 1}, {Key: "bbb", Value: 3}, {Key: "a", Value: 1}}
	assertEqual(t, collectPairs(result), want)
	assertEqual(t, collectPairs(transform.Operator(nil)), []pair[string, int](nil))
}

func TestComposedToSeq2Operator(t *testing.T) {
	t.Parallel()
	calls := 0
	composed := op.Compose().Pipe2(func(src iter.Seq[string]) iter.Seq2[[]int, string] {
		calls++
		if src == nil {
			return nil
		}
		return withLength(src)
	})
	var operator op.ToSeq2Operator[string, []int, string] = composed.Operator
	assertEqual(t, calls, 0)
	if operator(nil) != nil {
		t.Fatal("expected nil result")
	}
	result := operator(slices.Values([]string{"a", "bbb"}))
	assertEqual(t, calls, 2)
	want := []pair[[]int, string]{{Key: []int{1}, Value: "a"}, {Key: []int{3}, Value: "bbb"}}
	assertEqual(t, collectPairs(result), want)
}

func TestComposedToSeqPipe(t *testing.T) {
	t.Parallel()
	var calls []string
	first := op.Compose2().Pipe(trackTransform(&calls, "first", op.Values[int, string]))
	transform := first.Pipe(trackTransform(&calls, "next", op.Map(func(v string) int { return len(v) })))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.All([]string{"a", "bbb", "a"}))
	assertEqual(t, calls, []string{"first", "next"})
	assertEqual(t, slices.Collect(result), []int{1, 3, 1})
	assertEqual(t, slices.Collect(transform.Operator(nil)), []int(nil))
}

func TestComposedToSeqPipe2(t *testing.T) {
	t.Parallel()
	var calls []string
	first := op.Compose2().Pipe(trackTransform(&calls, "first", op.Values[int, string]))
	transform := first.Pipe2(trackTransform(&calls, "next", withLength))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.All([]string{"a", "bbb"}))
	assertEqual(t, calls, []string{"first", "next"})
	want := []pair[[]int, string]{{Key: []int{1}, Value: "a"}, {Key: []int{3}, Value: "bbb"}}
	assertEqual(t, collectPairs(result), want)
	assertEqual(t, collectPairs(transform.Operator(nil)), []pair[[]int, string](nil))
}

func TestComposedToSeqOperator(t *testing.T) {
	t.Parallel()
	calls := 0
	composed := op.Compose2().Pipe(func(src iter.Seq2[int, string]) iter.Seq[string] {
		calls++
		if src == nil {
			return nil
		}
		return op.Values(src)
	})
	var operator op.ToSeqOperator[int, string, string] = composed.Operator
	assertEqual(t, calls, 0)
	if operator(nil) != nil {
		t.Fatal("expected nil result")
	}
	result := operator(slices.All([]string{"a", "bbb"}))
	assertEqual(t, calls, 2)
	assertEqual(t, slices.Collect(result), []string{"a", "bbb"})
}

func TestComposed2Pipe(t *testing.T) {
	t.Parallel()
	var calls []string
	lengths := op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	first := op.Compose2().Pipe2(trackTransform(&calls, "first", lengths))
	transform := first.Pipe(trackTransform(&calls, "next", op.Values[string, int]))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.All([]string{"a", "bbb", "a"}))
	assertEqual(t, calls, []string{"first", "next"})
	assertEqual(t, slices.Collect(result), []int{1, 3, 1})
	assertEqual(t, slices.Collect(transform.Operator(nil)), []int(nil))
}

func TestComposed2Pipe2(t *testing.T) {
	t.Parallel()
	var calls []string
	lengths := op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })
	first := op.Compose2().Pipe2(trackTransform(&calls, "first", lengths))
	next := op.Map2(func(k string, v int) ([]byte, string) { return []byte(k), strconv.Itoa(v) })
	transform := first.Pipe2(trackTransform(&calls, "next", next))
	assertEqual(t, calls, []string(nil))
	result := transform.Operator(slices.All([]string{"a", "bbb"}))
	assertEqual(t, calls, []string{"first", "next"})
	want := []pair[[]byte, string]{{Key: []byte("0"), Value: "1"}, {Key: []byte("1"), Value: "3"}}
	assertEqual(t, collectPairs(result), want)
	assertEqual(t, collectPairs(transform.Operator(nil)), []pair[[]byte, string](nil))
}

func TestComposed2Operator(t *testing.T) {
	t.Parallel()
	calls := 0
	composed := op.Compose2().Pipe2(func(src iter.Seq2[int, string]) iter.Seq2[string, int] {
		calls++
		if src == nil {
			return nil
		}
		return op.Map2(func(k int, v string) (string, int) { return strconv.Itoa(k), len(v) })(src)
	})
	var operator op.Operator2[int, string, string, int] = composed.Operator
	assertEqual(t, calls, 0)
	if operator(nil) != nil {
		t.Fatal("expected nil result")
	}
	result := operator(slices.All([]string{"a", "bbb"}))
	assertEqual(t, calls, 2)
	want := []pair[string, int]{{Key: "0", Value: 1}, {Key: "1", Value: 3}}
	assertEqual(t, collectPairs(result), want)
}

// 独自オペレーター間で nil をそのまま渡し、戻り値も保持することを確認する。
func TestCompositionCustomNil(t *testing.T) {
	t.Parallel()
	var first op.Operator[int, string] = func(src iter.Seq[int]) iter.Seq[string] {
		if src != nil {
			t.Fatal("expected nil input")
		}
		return nil
	}
	var next op.Operator[string, int] = func(src iter.Seq[string]) iter.Seq[int] {
		if src != nil {
			t.Fatal("expected nil intermediate")
		}
		return nil
	}
	if op.Compose().Pipe(first).Pipe(next).Operator(nil) != nil {
		t.Fatal("expected nil result")
	}
}
