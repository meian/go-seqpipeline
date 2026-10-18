package benchmarks

import (
	"fmt"
	"iter"
	"reflect"
	"runtime"
	"strings"
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
	run        unitRun
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
	run  unitRun
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
	if !ok || reference.run.fn == nil || reference.name == "" {
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

// 入力の準備は測定ループの外で行い、測定とテストは同じ名前付き処理を呼ぶ。
type unitInput struct {
	values        []int
	source        seq.Seq[int]
	pairs         seq.Seq2[int, int]
	standard      iter.Seq[int]
	standardPairs iter.Seq2[int, int]
	mapping       map[int]int
	short         seq.Seq[int]
	shortValues   []int
	warm          seq.Seq[int]
	warm2         seq.Seq2[int, int]
	warmValues    []int
	warmPairs     []pair
	end           int
}

type unitRun struct {
	data *unitInput
	fn   func(*unitInput) int64
}

func (run unitRun) call() int64 { return run.fn(run.data) }

func resolveUnitOperation(values []int, api, phase, impl string) (unitRun, int64) {
	for _, tc := range unitCases(values) {
		if tc.api != api || tc.phase != phase {
			continue
		}
		if impl == "Library" {
			return tc.run, tc.want
		}
		return referenceFor(tc, unitReferences(values)).run, tc.want
	}
	panic("missing unit operation: " + api + "/" + phase + "/" + impl)
}

func benchmarkUnitOperation(b *testing.B, data *unitInput, run func(*unitInput) int64) {
	b.ReportAllocs()
	for b.Loop() {
		checksum = run(data)
	}
}

func testUnitOperation(t *testing.T, api, phase, impl string, run func(*unitInput) int64) {
	t.Helper()
	inputs := [][]int{input(0), input(1), input(64), input(4096), {17, -2, 0, 9, 4, 8, 1, -7, 22, 5, 3}}
	for i, values := range inputs {
		t.Run(fmt.Sprintf("input=%d", i), func(t *testing.T) {
			operation, want := resolveUnitOperation(values, api, phase, impl)
			if got := run(operation.data); got != want {
				t.Fatalf("got %d; want %d", got, want)
			}
			if phase == "Fill" && (api == "Cache" || api == "Cache2") {
				expected := rangeValues(values)
				if api == "Cache2" {
					for k := range values {
						expected += int64(k)
					}
				}
				if got := cachedConsumer(); got != expected {
					t.Fatalf("filled cache: got %d; want %d", got, expected)
				}
			}
		})
	}
}

func TestUnitComparisonCoverage(t *testing.T) {
	for _, n := range []int{64, 4096} {
		references := unitReferences(input(n))
		used := make(map[string]bool)
		for _, tc := range unitCases(input(n)) {
			reference := referenceFor(tc, references)
			for impl, run := range map[string]unitRun{"Go": reference.run, "Library": tc.run} {
				stem := strings.ReplaceAll(tc.api, ".", "") + tc.phase + impl
				name := runtime.FuncForPC(reflect.ValueOf(run.fn).Pointer()).Name()
				if !strings.HasSuffix(name, ".runUnit"+stem) {
					t.Fatalf("unexpected shared function: %s", name)
				}
				entry := unitBenchmarkEntries[tc.api+"/"+tc.phase+"/"+impl]
				if entry == nil {
					t.Fatalf("missing benchmark entry: %s", stem)
				}
				entryName := runtime.FuncForPC(reflect.ValueOf(entry).Pointer()).Name()
				if !strings.HasSuffix(entryName, ".benchmarkUnit"+stem) {
					t.Fatalf("unexpected benchmark entry: %s", entryName)
				}
			}
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
				name      string
				operation unitRun
			}{{"Go", reference.run}, {"Library", tc.run}} {
				entry := unitBenchmarkEntries[tc.api+"/"+tc.phase+"/"+impl.name]
				b.Run(fmt.Sprintf("api=%s/N=%d/phase=%s/ref=%s/impl=%s", tc.api, n, tc.phase, reference.name, impl.name), func(b *testing.B) { entry(b, impl.operation.data) })
			}
		}
	}
}

var unitBenchmarkEntries = map[string]func(*testing.B, *unitInput){
	"FromSlice/Evaluate/Library":               benchmarkUnitFromSliceEvaluateLibrary,
	"FromSeq/Evaluate/Library":                 benchmarkUnitFromSeqEvaluateLibrary,
	"ToSeq/Evaluate/Library":                   benchmarkUnitToSeqEvaluateLibrary,
	"ToSlice/Collect/Library":                  benchmarkUnitToSliceCollectLibrary,
	"FromMap/Evaluate/Library":                 benchmarkUnitFromMapEvaluateLibrary,
	"FromSeq2/Evaluate/Library":                benchmarkUnitFromSeq2EvaluateLibrary,
	"ToSeq2/Evaluate/Library":                  benchmarkUnitToSeq2EvaluateLibrary,
	"ToMap/Collect/Library":                    benchmarkUnitToMapCollectLibrary,
	"FromSlice/Evaluate/Go":                    benchmarkUnitFromSliceEvaluateGo,
	"FromSeq/Evaluate/Go":                      benchmarkUnitFromSeqEvaluateGo,
	"ToSeq/Evaluate/Go":                        benchmarkUnitToSeqEvaluateGo,
	"ToSlice/Collect/Go":                       benchmarkUnitToSliceCollectGo,
	"FromMap/Evaluate/Go":                      benchmarkUnitFromMapEvaluateGo,
	"FromSeq2/Evaluate/Go":                     benchmarkUnitFromSeq2EvaluateGo,
	"ToSeq2/Evaluate/Go":                       benchmarkUnitToSeq2EvaluateGo,
	"ToMap/Collect/Go":                         benchmarkUnitToMapCollectGo,
	"Map/Evaluate/Library":                     benchmarkUnitMapEvaluateLibrary,
	"Filter/Evaluate/Library":                  benchmarkUnitFilterEvaluateLibrary,
	"Map2/Evaluate/Library":                    benchmarkUnitMap2EvaluateLibrary,
	"Filter2/Evaluate/Library":                 benchmarkUnitFilter2EvaluateLibrary,
	"Keys/Evaluate/Library":                    benchmarkUnitKeysEvaluateLibrary,
	"Values/Evaluate/Library":                  benchmarkUnitValuesEvaluateLibrary,
	"Map/Evaluate/Go":                          benchmarkUnitMapEvaluateGo,
	"Filter/Evaluate/Go":                       benchmarkUnitFilterEvaluateGo,
	"Map2/Evaluate/Go":                         benchmarkUnitMap2EvaluateGo,
	"Filter2/Evaluate/Go":                      benchmarkUnitFilter2EvaluateGo,
	"Keys/Evaluate/Go":                         benchmarkUnitKeysEvaluateGo,
	"Values/Evaluate/Go":                       benchmarkUnitValuesEvaluateGo,
	"Take.Operator/Early8/Library":             benchmarkUnitTakeOperatorEarly8Library,
	"Skip.Operator/Evaluate/Library":           benchmarkUnitSkipOperatorEvaluateLibrary,
	"Take.Operator2/Early8/Library":            benchmarkUnitTakeOperator2Early8Library,
	"Skip.Operator2/Evaluate/Library":          benchmarkUnitSkipOperator2EvaluateLibrary,
	"Take.Operator/Early8/Go":                  benchmarkUnitTakeOperatorEarly8Go,
	"Skip.Operator/Evaluate/Go":                benchmarkUnitSkipOperatorEvaluateGo,
	"Take.Operator2/Early8/Go":                 benchmarkUnitTakeOperator2Early8Go,
	"Skip.Operator2/Evaluate/Go":               benchmarkUnitSkipOperator2EvaluateGo,
	"Seq.Pipe/Evaluate/Library":                benchmarkUnitSeqPipeEvaluateLibrary,
	"Seq.Pipe2/Evaluate/Library":               benchmarkUnitSeqPipe2EvaluateLibrary,
	"Seq2.Pipe/Evaluate/Library":               benchmarkUnitSeq2PipeEvaluateLibrary,
	"Seq2.Pipe2/Evaluate/Library":              benchmarkUnitSeq2Pipe2EvaluateLibrary,
	"Seq.Pipe/Evaluate/Go":                     benchmarkUnitSeqPipeEvaluateGo,
	"Seq.Pipe2/Evaluate/Go":                    benchmarkUnitSeqPipe2EvaluateGo,
	"Seq2.Pipe/Evaluate/Go":                    benchmarkUnitSeq2PipeEvaluateGo,
	"Seq2.Pipe2/Evaluate/Go":                   benchmarkUnitSeq2Pipe2EvaluateGo,
	"FromChannel/Evaluate/Library":             benchmarkUnitFromChannelEvaluateLibrary,
	"ToChannel/Evaluate/Library":               benchmarkUnitToChannelEvaluateLibrary,
	"FromChannel/Evaluate/Go":                  benchmarkUnitFromChannelEvaluateGo,
	"ToChannel/Evaluate/Go":                    benchmarkUnitToChannelEvaluateGo,
	"Cache/Fill/Library":                       benchmarkUnitCacheFillLibrary,
	"Cache/Warm/Library":                       benchmarkUnitCacheWarmLibrary,
	"Cache2/Fill/Library":                      benchmarkUnitCache2FillLibrary,
	"Cache2/Warm/Library":                      benchmarkUnitCache2WarmLibrary,
	"Cache/Fill/Go":                            benchmarkUnitCacheFillGo,
	"Cache/Warm/Go":                            benchmarkUnitCacheWarmGo,
	"Cache2/Fill/Go":                           benchmarkUnitCache2FillGo,
	"Cache2/Warm/Go":                           benchmarkUnitCache2WarmGo,
	"Zip.Operator/Evaluate/Library":            benchmarkUnitZipOperatorEvaluateLibrary,
	"Zip.AlignBy/Evaluate/Library":             benchmarkUnitZipAlignByEvaluateLibrary,
	"Zip.WithDefaultValue/Evaluate/Library":    benchmarkUnitZipWithDefaultValueEvaluateLibrary,
	"Zip.WithDefaultKey/Evaluate/Library":      benchmarkUnitZipWithDefaultKeyEvaluateLibrary,
	"ZipKey.Operator/Evaluate/Library":         benchmarkUnitZipKeyOperatorEvaluateLibrary,
	"ZipKey.AlignBy/Evaluate/Library":          benchmarkUnitZipKeyAlignByEvaluateLibrary,
	"ZipKey.WithDefaultValue/Evaluate/Library": benchmarkUnitZipKeyWithDefaultValueEvaluateLibrary,
	"ZipKey.WithDefaultKey/Evaluate/Library":   benchmarkUnitZipKeyWithDefaultKeyEvaluateLibrary,
	"Zip.Operator/Evaluate/Go":                 benchmarkUnitZipOperatorEvaluateGo,
	"Zip.AlignBy/Evaluate/Go":                  benchmarkUnitZipAlignByEvaluateGo,
	"Zip.WithDefaultValue/Evaluate/Go":         benchmarkUnitZipWithDefaultValueEvaluateGo,
	"Zip.WithDefaultKey/Evaluate/Go":           benchmarkUnitZipWithDefaultKeyEvaluateGo,
	"ZipKey.Operator/Evaluate/Go":              benchmarkUnitZipKeyOperatorEvaluateGo,
	"ZipKey.AlignBy/Evaluate/Go":               benchmarkUnitZipKeyAlignByEvaluateGo,
	"ZipKey.WithDefaultValue/Evaluate/Go":      benchmarkUnitZipKeyWithDefaultValueEvaluateGo,
	"ZipKey.WithDefaultKey/Evaluate/Go":        benchmarkUnitZipKeyWithDefaultKeyEvaluateGo,
}
