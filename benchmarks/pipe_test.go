package benchmarks

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func pipeUnitCases(values []int) []unitCase {
	source := seq.FromSlice(values)
	pairs := seq.FromSeq2(slices.All(values))
	var mapped, total, keys int64
	for k, v := range values {
		total += int64(v)
		keys += int64(k)
		mapped += int64(mapValue(v))
	}
	data := &unitInput{
		values: values,
		source: source,
		pairs:  pairs,
	}
	return []unitCase{
		{"Seq.Pipe", "Evaluate", unitRun{data, runUnitSeqPipeEvaluateLibrary}, mapped},
		{"Seq.Pipe2", "Evaluate", unitRun{data, runUnitSeqPipe2EvaluateLibrary}, total * 2},
		{"Seq2.Pipe", "Evaluate", unitRun{data, runUnitSeq2PipeEvaluateLibrary}, total},
		{"Seq2.Pipe2", "Evaluate", unitRun{data, runUnitSeq2Pipe2EvaluateLibrary}, keys + mapped},
	}
}

func pipeUnitReferences(values []int) map[string]unitReference {
	data := &unitInput{
		values: values,
	}
	return map[string]unitReference{
		"Seq.Pipe":   {"MapLoop", unitRun{data, runUnitSeqPipeEvaluateGo}},
		"Seq.Pipe2":  {"ZipIndex", unitRun{data, runUnitSeqPipe2EvaluateGo}},
		"Seq2.Pipe":  {"SliceRange", unitRun{data, runUnitSeq2PipeEvaluateGo}},
		"Seq2.Pipe2": {"Map2Loop", unitRun{data, runUnitSeq2Pipe2EvaluateGo}},
	}
}

func runUnitSeqPipeEvaluateLibrary(data *unitInput) int64 {
	return sumSeq(data.source.Pipe(seq.Map(mapValue)))
}

func benchmarkUnitSeqPipeEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeqPipeEvaluateLibrary)
}

func TestUnitSeqPipeEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Seq.Pipe", "Evaluate", "Library", runUnitSeqPipeEvaluateLibrary)
}

func runUnitSeqPipeEvaluateGo(data *unitInput) int64 {
	var sum int64
	for _, v := range data.values {
		sum += int64(mapValue(v))
	}
	return sum
}

func benchmarkUnitSeqPipeEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeqPipeEvaluateGo)
}

func TestUnitSeqPipeEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Seq.Pipe", "Evaluate", "Go", runUnitSeqPipeEvaluateGo)
}

func runUnitSeqPipe2EvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(data.source.Pipe2(seq.Zip(data.source).Operator))
}

func benchmarkUnitSeqPipe2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeqPipe2EvaluateLibrary)
}

func TestUnitSeqPipe2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Seq.Pipe2", "Evaluate", "Library", runUnitSeqPipe2EvaluateLibrary)
}

func runUnitSeqPipe2EvaluateGo(data *unitInput) int64 { return zipIndex(data.values, data.values) }

func benchmarkUnitSeqPipe2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeqPipe2EvaluateGo)
}

func TestUnitSeqPipe2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "Seq.Pipe2", "Evaluate", "Go", runUnitSeqPipe2EvaluateGo)
}

func runUnitSeq2PipeEvaluateLibrary(data *unitInput) int64 {
	return sumSeq(data.pairs.Pipe(seq.Values))
}

func benchmarkUnitSeq2PipeEvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeq2PipeEvaluateLibrary)
}

func TestUnitSeq2PipeEvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Seq2.Pipe", "Evaluate", "Library", runUnitSeq2PipeEvaluateLibrary)
}

func runUnitSeq2PipeEvaluateGo(data *unitInput) int64 { return rangeValues(data.values) }

func benchmarkUnitSeq2PipeEvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeq2PipeEvaluateGo)
}

func TestUnitSeq2PipeEvaluateGo(t *testing.T) {
	testUnitOperation(t, "Seq2.Pipe", "Evaluate", "Go", runUnitSeq2PipeEvaluateGo)
}

func runUnitSeq2Pipe2EvaluateLibrary(data *unitInput) int64 {
	return sumSeq2(data.pairs.Pipe2(seq.Map2(func(k, v int) (int, int) { return k, mapValue(v) })))
}

func benchmarkUnitSeq2Pipe2EvaluateLibrary(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeq2Pipe2EvaluateLibrary)
}

func TestUnitSeq2Pipe2EvaluateLibrary(t *testing.T) {
	testUnitOperation(t, "Seq2.Pipe2", "Evaluate", "Library", runUnitSeq2Pipe2EvaluateLibrary)
}

func runUnitSeq2Pipe2EvaluateGo(data *unitInput) int64 {
	var sum int64
	for k, v := range data.values {
		sum += int64(k) + int64(mapValue(v))
	}
	return sum
}

func benchmarkUnitSeq2Pipe2EvaluateGo(b *testing.B, data *unitInput) {
	benchmarkUnitOperation(b, data, runUnitSeq2Pipe2EvaluateGo)
}

func TestUnitSeq2Pipe2EvaluateGo(t *testing.T) {
	testUnitOperation(t, "Seq2.Pipe2", "Evaluate", "Go", runUnitSeq2Pipe2EvaluateGo)
}
