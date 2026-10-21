package op_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

// Take の件数制限を両入力へ伝え、再評価でも余分な取得をせずに入力を解放する。
// Take(0) では、どちらの Zip 設定型でも両入力を開始しない。
func TestZipTake(t *testing.T) {
	for _, alignment := range []op.ZipAlign{op.AlignShortest, op.AlignLeft, op.AlignRight, op.AlignLongest} {
		for _, fixedKey := range []bool{false, true} {
			for _, tc := range []struct {
				name   string
				n      int
				want   []pair[int, int]
				starts int
			}{
				{name: "zero"},
				{name: "one", n: 1, want: []pair[int, int]{{0, 0}}, starts: 1},
				{name: "prefix", n: 3, want: []pair[int, int]{{0, 0}, {1, 1}, {2, 2}}, starts: 1},
			} {
				name := fmt.Sprintf("align=%d/fixed_key=%t/%s", alignment, fixedKey, tc.name)
				t.Run(name, func(t *testing.T) {
					var started, visited, stopped [2]int
					input := func(index int) iter.Seq[int] {
						return func(yield func(int) bool) {
							started[index]++
							defer func() { stopped[index]++ }()
							for v := 0; v < 5; v++ {
								visited[index]++
								if !yield(v) {
									return
								}
							}
						}
					}
					config := op.Zip(input(1)).AlignBy(alignment)
					zipped := config.Operator(input(0))
					if fixedKey {
						zipped = config.WithDefaultKey(-1).Operator(input(0))
					}
					output := op.Take(tc.n).Operator2(zipped)
					assertEqual(t, started, [2]int{})
					for round := 1; round <= 2; round++ {
						assertEqual(t, collectPairs(output), tc.want)
						assertEqual(t, started, [2]int{tc.starts * round, tc.starts * round})
						assertEqual(t, visited, [2]int{tc.n * round, tc.n * round})
						assertEqual(t, stopped, started)
					}
				})
			}
		}
	}
}
