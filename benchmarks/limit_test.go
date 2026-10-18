package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func limitUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var taken, skipped, takenKeys, skippedKeys int64
	for k, v := range values {
		if k < 8 {
			taken += int64(v)
			takenKeys += int64(k)
		}
		if k >= 8 {
			skipped += int64(v)
			skippedKeys += int64(k)
		}
	}
	data := &unitInput{
		values: values,
		source: source,
		pairs:  pairs,
	}
	return []unitCase{
		{"Take.Operator", "Early8", unitRun{data, runUnitTakeOperatorEarly8Library}, taken},
		{"Skip.Operator", "Evaluate", unitRun{data, runUnitSkipOperatorEvaluateLibrary}, skipped},
		{"Take.Operator2", "Early8", unitRun{data, runUnitTakeOperator2Early8Library}, taken + takenKeys},
		{"Skip.Operator2", "Evaluate", unitRun{data, runUnitSkipOperator2EvaluateLibrary}, skipped + skippedKeys},
	}
}

func limitUnitReferences(values []int) map[string]unitReference {
	end := min(8, len(values))
	data := &unitInput{
		values: values,
		end:    end,
	}
	return map[string]unitReference{
		"Take.Operator":  {"TakeSlice", unitRun{data, runUnitTakeOperatorEarly8Go}},
		"Skip.Operator":  {"SkipSlice", unitRun{data, runUnitSkipOperatorEvaluateGo}},
		"Take.Operator2": {"Take2Slice", unitRun{data, runUnitTakeOperator2Early8Go}},
		"Skip.Operator2": {"Skip2Slice", unitRun{data, runUnitSkipOperator2EvaluateGo}},
	}
}

func runUnitTakeOperatorEarly8Library(data *unitInput) int64 {
	return sumSeq(seq.Take(8).Operator(data.source))
}

func benchmarkUnitTakeOperatorEarly8Library(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitTakeOperatorEarly8Library)
}

func TestUnitTakeOperatorEarly8Library(t *testing.T) {
	testUnitOperation(t, "Take.Operator", "Early8", "Library", runUnitTakeOperatorEarly8Library)
}

func runUnitTakeOperatorEarly8Go(data *unitInput) int64 { return rangeValues(data.values[:data.end]) }

func benchmarkUnitTakeOperatorEarly8Go(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitTakeOperatorEarly8Go)
}

func TestUnitTakeOperatorEarly8Go(t *testing.T) {
	testUnitOperation(t, "Take.Operator", "Early8", "Go", runUnitTakeOperatorEarly8Go)
}

func runUnitSkipOperatorEvaluateLibrary(data *unitInput) int64 {
	return sumSeq(seq.Skip(8).Operator(data.source))
}

func benchmarkUnitSkipOperatorEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSkipOperatorEvaluateLibrary)
}

func TestUnitSkipOperatorEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Skip.Operator", "Evaluate", "Library", runUnitSkipOperatorEvaluateLibrary)
}

func runUnitSkipOperatorEvaluateGo(data *unitInput) int64 { return rangeValues(data.values[data.end:]) }

func benchmarkUnitSkipOperatorEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSkipOperatorEvaluateGo)
}

func TestUnitSkipOperatorEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Skip.Operator", "Evaluate", "Go", runUnitSkipOperatorEvaluateGo)
}

func runUnitTakeOperator2Early8Library(data *unitInput) int64 {
	return sumSeq2(seq.Take(8).Operator2(data.pairs))
}

func benchmarkUnitTakeOperator2Early8Library(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitTakeOperator2Early8Library)
}

func TestUnitTakeOperator2Early8Library(t *testing.T) {
	testUnitOperation(t, "Take.Operator2", "Early8", "Library", runUnitTakeOperator2Early8Library)
}

func runUnitTakeOperator2Early8Go(data *unitInput) int64 {
	var sum int64
	for k, v := range data.values[:data.end] {
		sum += int64(k) + int64(v)
	}
	return sum
}

func benchmarkUnitTakeOperator2Early8Go(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitTakeOperator2Early8Go)
}

func TestUnitTakeOperator2Early8Go(t *testing.T) {
	testUnitOperation(t, "Take.Operator2", "Early8", "Go", runUnitTakeOperator2Early8Go)
}

func runUnitSkipOperator2EvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(seq.Skip(8).Operator2(data.pairs))
}

func benchmarkUnitSkipOperator2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSkipOperator2EvaluateLibrary)
}

func TestUnitSkipOperator2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Skip.Operator2", "Evaluate", "Library", runUnitSkipOperator2EvaluateLibrary)
}

func runUnitSkipOperator2EvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.values[data.end:] {
		sum += int64(k+data.end) + int64(v)
	}
	return sum
}

func benchmarkUnitSkipOperator2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSkipOperator2EvaluateGo)
}

func TestUnitSkipOperator2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "Skip.Operator2", "Evaluate", "Go", runUnitSkipOperator2EvaluateGo)
}
