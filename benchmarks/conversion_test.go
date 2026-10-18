package benchmarks

import (
	"maps"
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

var collectedValues []int

var collectedMap map[int]int

func conversionUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	standard := slices.Values(values)
	standardPairs := slices.All(values)
	mapping := make(map[int]int, len(values))
	var total, keys int64
	for k, v := range values {
		mapping[k] = v
		total += int64(v)
		keys += int64(k)
	}
	data := &unitInput{
		values:        values,
		source:        source,
		pairs:         pairs,
		standard:      standard,
		standardPairs: standardPairs,
		mapping:       mapping,
	}
	return []unitCase{
		{"FromSlice", "Evaluate", unitRun{data, runUnitFromSliceEvaluateLibrary}, total},
		{"FromSeq", "Evaluate", unitRun{data, runUnitFromSeqEvaluateLibrary}, total},
		{"ToSeq", "Evaluate", unitRun{data, runUnitToSeqEvaluateLibrary}, total},
		{"ToSlice", "Collect", unitRun{data, runUnitToSliceCollectLibrary}, total},
		{"FromMap", "Evaluate", unitRun{data, runUnitFromMapEvaluateLibrary}, keys + total},
		{"FromSeq2", "Evaluate", unitRun{data, runUnitFromSeq2EvaluateLibrary}, keys + total},
		{"ToSeq2", "Evaluate", unitRun{data, runUnitToSeq2EvaluateLibrary}, keys + total},
		{"ToMap", "Collect", unitRun{data, runUnitToMapCollectLibrary}, keys + total},
	}
}

func conversionUnitReferences(values []int) map[string]unitReference {
	standard := slices.Values(values)
	standardPairs := slices.All(values)
	mapping := make(map[int]int, len(values))
	for k, v := range values {
		mapping[k] = v
	}
	data := &unitInput{
		values:        values,
		standard:      standard,
		standardPairs: standardPairs,
		mapping:       mapping,
	}

	return map[string]unitReference{
		"FromSlice": {"SliceRange", unitRun{data, runUnitFromSliceEvaluateGo}},
		"FromSeq":   {"IteratorRange", unitRun{data, runUnitFromSeqEvaluateGo}},
		"ToSeq":     {"IteratorRange", unitRun{data, runUnitToSeqEvaluateGo}},
		"ToSlice":   {"SlicesCollect", unitRun{data, runUnitToSliceCollectGo}},
		"FromMap":   {"MapRange", unitRun{data, runUnitFromMapEvaluateGo}},
		"FromSeq2":  {"Iterator2Range", unitRun{data, runUnitFromSeq2EvaluateGo}},
		"ToSeq2":    {"Iterator2Range", unitRun{data, runUnitToSeq2EvaluateGo}},
		"ToMap":     {"MapsCollect", unitRun{data, runUnitToMapCollectGo}},
	}
}

func runUnitFromSliceEvaluateLibrary(data *unitInput) int64 {
	return sumSeq(seq.FromSlice(data.values))
}

func benchmarkUnitFromSliceEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromSliceEvaluateLibrary)
}

func TestUnitFromSliceEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "FromSlice", "Evaluate", "Library", runUnitFromSliceEvaluateLibrary)
}

func runUnitFromSliceEvaluateGo(data *unitInput) int64 { return rangeValues(data.values) }

func benchmarkUnitFromSliceEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromSliceEvaluateGo)
}

func TestUnitFromSliceEvaluateGo(t *testing.T) {
	testUnitOperation(t, "FromSlice", "Evaluate", "Go", runUnitFromSliceEvaluateGo)
}

func runUnitFromSeqEvaluateLibrary(data *unitInput) int64 { return sumSeq(seq.FromSeq(data.standard)) }

func benchmarkUnitFromSeqEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromSeqEvaluateLibrary)
}

func TestUnitFromSeqEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "FromSeq", "Evaluate", "Library", runUnitFromSeqEvaluateLibrary)
}

func runUnitFromSeqEvaluateGo(data *unitInput) int64 {
	var sum int64
	for v := range data.standard {
		sum += int64(v)
	}
	return sum
}

func benchmarkUnitFromSeqEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromSeqEvaluateGo)
}

func TestUnitFromSeqEvaluateGo(t *testing.T) {
	testUnitOperation(t, "FromSeq", "Evaluate", "Go", runUnitFromSeqEvaluateGo)
}

func runUnitToSeqEvaluateLibrary(data *unitInput) int64 {
	var sum int64
	for v := range seq.ToSeq(data.source) {
		sum += int64(v)
	}
	return sum
}

func benchmarkUnitToSeqEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToSeqEvaluateLibrary)
}

func TestUnitToSeqEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ToSeq", "Evaluate", "Library", runUnitToSeqEvaluateLibrary)
}

func runUnitToSeqEvaluateGo(data *unitInput) int64 {
	var sum int64
	for v := range data.standard {
		sum += int64(v)
	}
	return sum
}

func benchmarkUnitToSeqEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToSeqEvaluateGo)
}

func TestUnitToSeqEvaluateGo(t *testing.T) {
	testUnitOperation(t, "ToSeq", "Evaluate", "Go", runUnitToSeqEvaluateGo)
}

func runUnitToSliceCollectLibrary(data *unitInput) int64 {
	collectedValues = seq.ToSlice(data.source)
	var sum int64
	for _, v := range collectedValues {
		sum += int64(v)
	}
	return sum
}

func benchmarkUnitToSliceCollectLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToSliceCollectLibrary)
}

func TestUnitToSliceCollectLibrary(t *testing.T) {
	testUnitOperation(t, "ToSlice", "Collect", "Library", runUnitToSliceCollectLibrary)
}

func runUnitToSliceCollectGo(data *unitInput) int64 {
	collectedValues = slices.Collect(data.standard)
	return rangeValues(collectedValues)
}

func benchmarkUnitToSliceCollectGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToSliceCollectGo)
}

func TestUnitToSliceCollectGo(t *testing.T) {
	testUnitOperation(t, "ToSlice", "Collect", "Go", runUnitToSliceCollectGo)
}

func runUnitFromMapEvaluateLibrary(data *unitInput) int64 { return sumSeq2(seq.FromMap(data.mapping)) }

func benchmarkUnitFromMapEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromMapEvaluateLibrary)
}

func TestUnitFromMapEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "FromMap", "Evaluate", "Library", runUnitFromMapEvaluateLibrary)
}

func runUnitFromMapEvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.mapping {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitFromMapEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromMapEvaluateGo)
}

func TestUnitFromMapEvaluateGo(t *testing.T) {
	testUnitOperation(t, "FromMap", "Evaluate", "Go", runUnitFromMapEvaluateGo)
}

func runUnitFromSeq2EvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.FromSeq2(data.standardPairs))
}

func benchmarkUnitFromSeq2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromSeq2EvaluateLibrary)
}

func TestUnitFromSeq2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "FromSeq2", "Evaluate", "Library", runUnitFromSeq2EvaluateLibrary)
}

func runUnitFromSeq2EvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.standardPairs {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitFromSeq2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitFromSeq2EvaluateGo)
}

func TestUnitFromSeq2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "FromSeq2", "Evaluate", "Go", runUnitFromSeq2EvaluateGo)
}

func runUnitToSeq2EvaluateLibrary(data *unitInput) int64 {
	var sum int64
	for k, v := range seq.ToSeq2(data.pairs) {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitToSeq2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToSeq2EvaluateLibrary)
}

func TestUnitToSeq2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "ToSeq2", "Evaluate", "Library", runUnitToSeq2EvaluateLibrary)
}

func runUnitToSeq2EvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.standardPairs {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitToSeq2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToSeq2EvaluateGo)
}

func TestUnitToSeq2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "ToSeq2", "Evaluate", "Go", runUnitToSeq2EvaluateGo)
}

func runUnitToMapCollectLibrary(data *unitInput) int64 {
	collectedMap = seq.ToMap(data.pairs)
	var sum int64
	for k, v := range collectedMap {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitToMapCollectLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToMapCollectLibrary)
}

func TestUnitToMapCollectLibrary(t *testing.T) {
	testUnitOperation(t, "ToMap", "Collect", "Library", runUnitToMapCollectLibrary)
}

func runUnitToMapCollectGo(data *unitInput) int64 {
	collectedMap = maps.Collect(data.standardPairs)
	var sum int64
	for k, v := range collectedMap {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitToMapCollectGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitToMapCollectGo)
}

func TestUnitToMapCollectGo(t *testing.T) {
	testUnitOperation(t, "ToMap", "Collect", "Go", runUnitToMapCollectGo)
}
