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

type unitReference struct {
	name string
	run  func() int64
}

func unitReferences(values []int) map[string]unitReference {
	references := make(map[string]unitReference)
	for _, family := range []map[string]unitReference{
		conversionUnitReferences(values), transformUnitReferences(values), limitUnitReferences(values),
		pipeUnitReferences(values), channelUnitReferences(values), cacheUnitReferences(values), zipUnitReferences(values),
	} {
		for api, reference := range family {
			if _, exists := references[api]; exists {
				panic("duplicate unit reference: " + api)
			}
			references[api] = reference
		}
	}
	return references
}

func referenceFor(tc unitCase, references map[string]unitReference) unitReference {
	reference, ok := references[tc.api+"/"+tc.phase]
	if !ok {
		reference, ok = references[tc.api]
	}
	if !ok || reference.run == nil || reference.name == "" {
		panic("missing unit reference: " + tc.api + "/" + tc.phase)
	}
	return reference
}

func rangeValues(values []int) int64 {
	var sum int64
	for _, v := range values {
		sum += int64(v)
	}
	return sum
}

func checkUnitCases(t *testing.T, factory func([]int) []unitCase, references func([]int) map[string]unitReference) {
	t.Helper()
	inputs := [][]int{input(0), input(1), input(64), input(4096), {17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}}
	for i, values := range inputs {
		plain := references(values)
		for _, tc := range factory(values) {
			reference := referenceFor(tc, plain)
			t.Run(fmt.Sprintf("%s/input=%d/phase=%s", tc.api, i, tc.phase), func(t *testing.T) {
				for _, impl := range []struct {
					name string
					run  func() int64
				}{{"Go", reference.run}, {"Library", tc.run}} {
					if got := impl.run(); got != tc.want {
						t.Fatalf("%s got %d; want %d", impl.name, got, tc.want)
					}
				}
			})
		}
	}
}

func TestUnitComparisonCoverage(t *testing.T) {
	for _, n := range []int{64, 4096} {
		references := unitReferences(input(n))
		used := make(map[string]bool)
		for _, tc := range unitCases(input(n)) {
			_ = referenceFor(tc, references)
			key := tc.api
			if _, ok := references[key+"/"+tc.phase]; ok {
				key += "/" + tc.phase
			}
			if used[key] {
				t.Fatalf("duplicate unit comparison: %s", key)
			}
			used[key] = true
		}
		if len(used) != len(references) {
			t.Fatalf("unused references: used %d of %d", len(used), len(references))
		}
	}
}

func BenchmarkUnit(b *testing.B) {
	for _, n := range []int{64, 4096} {
		values := input(n)
		references := unitReferences(values)
		for _, tc := range unitCases(values) {
			reference := referenceFor(tc, references)
			for _, impl := range []struct {
				name string
				run  func() int64
			}{{"Go", reference.run}, {"Library", tc.run}} {
				b.Run(fmt.Sprintf("api=%s/N=%d/phase=%s/ref=%s/impl=%s", tc.api, n, tc.phase, reference.name, impl.name), func(b *testing.B) {
					b.ReportAllocs()
					for b.Loop() {
						checksum = impl.run()
					}
				})
			}
		}
	}
}
