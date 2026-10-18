package benchmarks

import (
	"fmt"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

type pair struct{ key, value int }

var output []pair

var checksum int64

func input(n int) []int {
	values := make([]int, n)
	for i := range values {
		values[i] = i
	}
	return values
}

func mapValue(v int) int { return (v + 1) & 65535 }

func keepValue(v int) bool { return v%3 != 0 }

func sumSeq(s seq.Seq[int]) int64 {
	var sum int64
	for v := range s {
		sum += int64(v)
	}
	return sum
}

func sumSeq2(s seq.Seq2[int, int]) int64 {
	var sum int64
	for k, v := range s {
		sum += int64(k) + int64(v)
	}
	return sum
}

type unitCase struct {
	api, phase string
	run        func() int64
	want       int64
}

// 機能別のケースを集約し、公開済みの Unit/api=... という名前を保つ。
// 入力生成とケースの準備は共通の測定ループより前に行う。
func unitCases(values []int) []unitCase {
	var cases []unitCase
	cases = append(cases, conversionUnitCases(values)...)
	cases = append(cases, transformUnitCases(values)...)
	cases = append(cases, limitUnitCases(values)...)
	cases = append(cases, pipeUnitCases(values)...)
	cases = append(cases, channelUnitCases(values)...)
	cases = append(cases, cacheUnitCases(values)...)
	cases = append(cases, zipUnitCases(values)...)
	return cases
}

func checkUnitCases(t *testing.T, factory func([]int) []unitCase) {
	t.Helper()
	for _, n := range []int{0, 1, 64, 4096} {
		for _, tc := range factory(input(n)) {
			t.Run(fmt.Sprintf("%s/N=%d/phase=%s", tc.api, n, tc.phase), func(t *testing.T) {
				if got := tc.run(); got != tc.want {
					t.Fatalf("got %d; want %d", got, tc.want)
				}
			})
		}
	}
}

func BenchmarkUnit(b *testing.B) {
	for _, n := range []int{64, 4096} {
		for _, tc := range unitCases(input(n)) {
			b.Run(fmt.Sprintf("api=%s/N=%d/phase=%s", tc.api, n, tc.phase), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					checksum = tc.run()
				}
			})
		}
	}
}
