package op_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/op"
)

var zipChecksum int64

// 同じ長さの出力について、補完割合ごとの Zip の構築・全件評価を測る。
// 入力イテレーターは再利用し、スライスの準備は計測から除く。
func BenchmarkZipPadding(b *testing.B) {
	const n = 4096
	for _, alignment := range []struct {
		name  string
		align op.ZipAlign
	}{
		{"left", op.AlignLeft},
		{"right", op.AlignRight},
		{"longest", op.AlignLongest},
	} {
		for _, size := range []struct {
			name string
			n    int
		}{
			{"nil", -1}, {"empty", 0}, {"one", 1}, {"half", n / 2}, {"equal", n},
		} {
			b.Run(fmt.Sprintf("align=%s/short=%s/N=%d", alignment.name, size.name, n), func(b *testing.B) {
				values := make([]int, n)
				for i := range values {
					values[i] = i + 1
				}
				left := slices.Values(values)
				var right iter.Seq[int]
				if size.n >= 0 {
					right = slices.Values(values[:size.n])
				}
				if alignment.align == op.AlignRight {
					left, right = right, left
				}
				b.ReportAllocs()
				for b.Loop() {
					output := op.Zip(right).AlignBy(alignment.align).
						WithDefaultKey(-1).WithDefaultValue(-2).Operator(left)
					var sum int64
					for k, v := range output {
						sum += int64(k) + int64(v)
					}
					zipChecksum = sum
				}
			})
		}
	}
}
