package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func transformUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var total, mapped, filtered, filteredKeys, keys int64
	for k, v := range values {
		total += int64(v)
		keys += int64(k)
		mapped += int64(mapValue(v))
		if keepValue(v) {
			filtered += int64(v)
			filteredKeys += int64(k)
		}
	}
	data := &unitInput{
		values: values,
		source: source,
		pairs:  pairs,
	}
	return []unitCase{
		{"Map", "Evaluate", unitRun{data, runUnitMapEvaluateLibrary}, mapped},
		{"Filter", "Evaluate", unitRun{data, runUnitFilterEvaluateLibrary}, filtered},
		{"Map2", "Evaluate", unitRun{data, runUnitMap2EvaluateLibrary}, keys + mapped},
		{"Filter2", "Evaluate", unitRun{data, runUnitFilter2EvaluateLibrary}, filtered + filteredKeys},
		{"Keys", "Evaluate", unitRun{data, runUnitKeysEvaluateLibrary}, keys},
		{"Values", "Evaluate", unitRun{data, runUnitValuesEvaluateLibrary}, total},
	}
}

func transformUnitReferences(values []int) map[string]unitReference {
	data := &unitInput{
		values: values,
	}
	return map[string]unitReference{
		"Map":     {"MapLoop", unitRun{data, runUnitMapEvaluateGo}},
		"Filter":  {"FilterLoop", unitRun{data, runUnitFilterEvaluateGo}},
		"Map2":    {"Map2Loop", unitRun{data, runUnitMap2EvaluateGo}},
		"Filter2": {"Filter2Loop", unitRun{data, runUnitFilter2EvaluateGo}},
		"Keys":    {"KeyLoop", unitRun{data, runUnitKeysEvaluateGo}},
		"Values":  {"SliceRange", unitRun{data, runUnitValuesEvaluateGo}},
	}
}

func runUnitMapEvaluateLibrary(data *unitInput) int64 { return sumSeq(seq.Map(mapValue)(data.source)) }

func benchmarkUnitMapEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitMapEvaluateLibrary)
}

func TestUnitMapEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Map", "Evaluate", "Library", runUnitMapEvaluateLibrary)
}

func runUnitMapEvaluateGo(data *unitInput) int64 {
	var sum int64
	for _, v := range data.values {
		sum += int64(mapValue(v))
	}
	return sum
}

func benchmarkUnitMapEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitMapEvaluateGo)
}

func TestUnitMapEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Map", "Evaluate", "Go", runUnitMapEvaluateGo)
}

func runUnitFilterEvaluateLibrary(data *unitInput) int64 {
	return sumSeq(seq.Filter(keepValue)(data.source))
}

func benchmarkUnitFilterEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFilterEvaluateLibrary)
}

func TestUnitFilterEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Filter", "Evaluate", "Library", runUnitFilterEvaluateLibrary)
}

func runUnitFilterEvaluateGo(data *unitInput) int64 {
	var sum int64
	for _, v := range data.values {
		if keepValue(v) {
			sum += int64(v)
		}
	}
	return sum
}

func benchmarkUnitFilterEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFilterEvaluateGo)
}

func TestUnitFilterEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Filter", "Evaluate", "Go", runUnitFilterEvaluateGo)
}

func runUnitMap2EvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })(data.pairs))
}

func benchmarkUnitMap2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitMap2EvaluateLibrary)
}

func TestUnitMap2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Map2", "Evaluate", "Library", runUnitMap2EvaluateLibrary)
}

func runUnitMap2EvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.values {
		sum += int64(k) + int64(mapValue(v))
	}
	return sum
}

func benchmarkUnitMap2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitMap2EvaluateGo)
}

func TestUnitMap2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "Map2", "Evaluate", "Go", runUnitMap2EvaluateGo)
}

func runUnitFilter2EvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Filter2(func(k, v int) bool { return keepValue(v) })(data.pairs))
}

func benchmarkUnitFilter2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFilter2EvaluateLibrary)
}

func TestUnitFilter2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Filter2", "Evaluate", "Library", runUnitFilter2EvaluateLibrary)
}

func runUnitFilter2EvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.values {
		if keepValue(v) {
			sum += int64(k) + int64(v)
		}
	}
	return sum
}

func benchmarkUnitFilter2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFilter2EvaluateGo)
}

func TestUnitFilter2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "Filter2", "Evaluate", "Go", runUnitFilter2EvaluateGo)
}

func runUnitKeysEvaluateLibrary(data *unitInput) int64 { return sumSeq(seq.Keys(data.pairs)) }

func benchmarkUnitKeysEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitKeysEvaluateLibrary)
}

func TestUnitKeysEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Keys", "Evaluate", "Library", runUnitKeysEvaluateLibrary)
}

func runUnitKeysEvaluateGo(data *unitInput) int64 {
	var sum int64
	for k := range data.values {
		sum += int64(k)
	}
	return sum
}

func benchmarkUnitKeysEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitKeysEvaluateGo)
}

func TestUnitKeysEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Keys", "Evaluate", "Go", runUnitKeysEvaluateGo)
}

func runUnitValuesEvaluateLibrary(data *unitInput) int64 { return sumSeq(seq.Values(data.pairs)) }

func benchmarkUnitValuesEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitValuesEvaluateLibrary)
}

func TestUnitValuesEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Values", "Evaluate", "Library", runUnitValuesEvaluateLibrary)
}

func runUnitValuesEvaluateGo(data *unitInput) int64 { return rangeValues(data.values) }

func benchmarkUnitValuesEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitValuesEvaluateGo)
}

func TestUnitValuesEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Values", "Evaluate", "Go", runUnitValuesEvaluateGo)
}
