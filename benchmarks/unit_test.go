package benchmarks

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedValues []int
var collectedMap map[int]int

type unitCase struct {
	api, phase string
	run        func() int64
	want       int64
}

// 入力・参照結果の準備は計測外。run には接続と評価を含める。
// source と pairs は再評価可能な入力であり、毎回同じ値を返す。
func unitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	standard := slices.Values(values)
	standardPairs := slices.All(values)
	mapping := make(map[int]int, len(values))
	var total, mapped, filtered, keys int64
	var taken, skipped int64
	for k, v := range values {
		mapping[k] = v
		total += int64(v)
		keys += int64(k)
		mapped += int64(mapValue(v))
		if keepValue(v) {
			filtered += int64(v)
		}
		if k < 8 {
			taken += int64(v)
		}
		if k >= 8 {
			skipped += int64(v)
		}
	}
	cases := []unitCase{
		{"FromSlice", "Evaluate", func() int64 { return sumSeq(seq.FromSlice(values)) }, total},
		{"FromSeq", "Evaluate", func() int64 { return sumSeq(seq.FromSeq(standard)) }, total},
		{"ToSeq", "Evaluate", func() int64 {
			var sum int64
			for v := range seq.ToSeq(source) {
				sum += int64(v)
			}
			return sum
		}, total},
		{"ToSlice", "Collect", func() int64 {
			collectedValues = seq.ToSlice(source)
			var sum int64
			for _, v := range collectedValues {
				sum += int64(v)
			}
			return sum
		}, total},
		{"Map", "Evaluate", func() int64 { return sumSeq(seq.Map(mapValue)(source)) }, mapped},
		{"Filter", "Evaluate", func() int64 { return sumSeq(seq.Filter(keepValue)(source)) }, filtered},
		{"Take.Operator", "Early8", func() int64 { return sumSeq(seq.Take(8).Operator(source)) }, taken},
		{"Skip.Operator", "Evaluate", func() int64 { return sumSeq(seq.Skip(8).Operator(source)) }, skipped},
		{"Seq.Pipe", "Evaluate", func() int64 { return sumSeq(source.Pipe(seq.Map(mapValue))) }, mapped},
		{"Seq.Pipe2", "Evaluate", func() int64 { return sumSeq2(source.Pipe2(seq.Zip(source).Operator)) }, total * 2},
		{"FromMap", "Evaluate", func() int64 { return sumSeq2(seq.FromMap(mapping)) }, keys + total},
		{"FromSeq2", "Evaluate", func() int64 { return sumSeq2(seq.FromSeq2(standardPairs)) }, keys + total},
		{"ToSeq2", "Evaluate", func() int64 {
			var sum int64
			for k, v := range seq.ToSeq2(pairs) {
				sum += int64(k) + int64(v)
			}
			return sum
		}, keys + total},
		{"ToMap", "Collect", func() int64 {
			collectedMap = seq.ToMap(pairs)
			var sum int64
			for k, v := range collectedMap {
				sum += int64(k) + int64(v)
			}
			return sum
		}, keys + total},
		{"Map2", "Evaluate", func() int64 { return sumSeq2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })(pairs)) }, keys + mapped},
		{"Filter2", "Evaluate", func() int64 { return sumSeq2(seq.Filter2(func(k, v int) bool { return keepValue(v) })(pairs)) }, filtered * 2},
		{"Take.Operator2", "Early8", func() int64 { return sumSeq2(seq.Take(8).Operator2(pairs)) }, taken * 2},
		{"Skip.Operator2", "Evaluate", func() int64 { return sumSeq2(seq.Skip(8).Operator2(pairs)) }, skipped * 2},
		{"Keys", "Evaluate", func() int64 { return sumSeq(seq.Keys(pairs)) }, keys},
		{"Values", "Evaluate", func() int64 { return sumSeq(seq.Values(pairs)) }, total},
		{"Seq2.Pipe", "Evaluate", func() int64 { return sumSeq(pairs.Pipe(seq.Values)) }, total},
		{"Seq2.Pipe2", "Evaluate", func() int64 {
			return sumSeq2(pairs.Pipe2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })))
		}, keys + mapped},
		{"FromChannel", "Evaluate", func() int64 {
			ch, done := producer(values, 0)
			sum := sumSeq(seq.FromChannel(context.Background(), ch))
			<-done
			return sum
		}, total},
		{"ToChannel", "Evaluate", func() int64 {
			var sum int64
			for v := range seq.ToChannel(context.Background(), source) {
				sum += int64(v)
			}
			return sum
		}, total},
		{"Cache", "Fill", func() int64 { s := seq.Cache(source); cachedConsumer = func() int64 { return sumSeq(s) }; return 0 }, 0},
		{"Cache2", "Fill", func() int64 { s := seq.Cache2(pairs); cachedConsumer = func() int64 { return sumSeq2(s) }; return 0 }, 0},
	}
	warm := seq.Cache(source)
	warm2 := seq.Cache2(pairs)
	cases = append(cases, unitCase{"Cache", "Warm", func() int64 { return sumSeq(warm) }, total}, unitCase{"Cache2", "Warm", func() int64 { return sumSeq2(warm2) }, total + keys})
	// 不足する左右の値を実際に通し、設定メソッドを評価費用と一緒に測る。
	short := seq.FromSlice(values[:len(values)/2])
	var shortTotal int64
	for _, v := range values[:len(values)/2] {
		shortTotal += int64(v)
	}
	missing := int64(len(values) - len(values)/2)
	cases = append(cases,
		unitCase{"Zip.Operator", "Evaluate", func() int64 { return sumSeq2(seq.Zip(source).Operator(source)) }, total * 2},
		unitCase{"Zip.AlignBy", "Evaluate", func() int64 { return sumSeq2(seq.Zip(short).AlignBy(seq.AlignLeft).Operator(source)) }, total + shortTotal},
		unitCase{"Zip.WithDefaultValue", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(short).AlignBy(seq.AlignLeft).WithDefaultValue(7).Operator(source))
		}, total + shortTotal + missing*7},
		unitCase{"Zip.WithDefaultKey", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(source).AlignBy(seq.AlignRight).WithDefaultKey(7).Operator(short))
		}, total + shortTotal + missing*7},
		unitCase{"ZipKey.Operator", "Evaluate", func() int64 { return sumSeq2(seq.Zip(source).WithDefaultKey(7).Operator(source)) }, total * 2},
		unitCase{"ZipKey.AlignBy", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(source).WithDefaultKey(7).AlignBy(seq.AlignRight).Operator(short))
		}, total + shortTotal + missing*7},
		unitCase{"ZipKey.WithDefaultValue", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(short).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultValue(9).Operator(source))
		}, total + shortTotal + missing*9},
		unitCase{"ZipKey.WithDefaultKey", "Evaluate", func() int64 {
			return sumSeq2(seq.Zip(source).WithDefaultKey(7).AlignBy(seq.AlignLongest).WithDefaultKey(9).Operator(short))
		}, total + shortTotal + missing*9},
	)
	return cases
}

func TestUnitResults(t *testing.T) {
	for _, n := range []int{0, 1, 64, 4096} {
		for _, tc := range unitCases(input(n)) {
			t.Run(fmt.Sprintf("%s/N=%d/phase=%s", tc.api, n, tc.phase), func(t *testing.T) {
				if got := tc.run(); got != tc.want {
					t.Fatalf("got %d; want %d", got, tc.want)
				}
				if tc.phase == "Fill" {
					want := int64(n * (n - 1) / 2)
					if tc.api == "Cache2" {
						want *= 2
					}
					if got := cachedConsumer(); got != want {
						t.Fatalf("filled cache: got %d; want %d", got, want)
					}
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
