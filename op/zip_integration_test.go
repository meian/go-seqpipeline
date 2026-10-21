package op_test

import (
	"fmt"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// Take から Zip への停止伝播と、0 件指定時の非評価だけを連携テストで確認する。
func TestZipTake(t *testing.T) {
	for _, n := range []int{0, 3} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			var state zipTrace
			input := []int{0, 1, 2, 3, 4}
			zipped := op.Zip(state.input(1, input)).Operator(state.input(0, input))
			output := op.Take(n).Operator2(zipped)
			assertEqual(t, state.started, [2]int{})
			var want []pair[int, int]
			if n > 0 {
				want = []pair[int, int]{{0, 0}, {1, 1}, {2, 2}}
			}
			assertEqual(t, collectPairs(output), want)
			assertEqual(t, state.reads, [2]int{n, n})
			assertEqual(t, state.started, [2]int{min(n, 1), min(n, 1)})
			assertEqual(t, state.stopped, state.started)
		})
	}
}
