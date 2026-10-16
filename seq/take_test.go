package seq_test

import (
	"slices"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestTake(t *testing.T) {
	for _, tc := range []struct {
		name   string
		n      int
		want   []int
		reads  int
		starts int
		stops  int
	}{
		{name: "negative", n: -1},
		{name: "zero", n: 0},
		{name: "one", n: 1, want: []int{1}, reads: 1, starts: 1, stops: 1},
		{name: "prefix", n: 2, want: []int{1, 2}, reads: 2, starts: 1, stops: 1},
		{name: "all", n: 3, want: []int{1, 2, 3}, reads: 3, starts: 1, stops: 1},
		{name: "larger", n: 4, want: []int{1, 2, 3}, reads: 3, starts: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads, starts, stops := 0, 0, 0
			src := func(yield func(int) bool) {
				starts++
				for i := 1; i <= 3; i++ {
					reads++
					if !yield(i) {
						stops++
						return
					}
				}
			}
			pipeline := seq.FromSeq(src).Pipe(seq.Take(tc.n).Operator)
			for round := 1; round <= 2; round++ {
				assertEqual(t, seq.ToSlice(pipeline), tc.want)
				gotCounts := []int{reads, starts, stops}
				wantCounts := []int{tc.reads * round, tc.starts * round, tc.stops * round}
				assertEqual(t, gotCounts, wantCounts)
			}
		})
	}
}

func TestTakeNilInput(t *testing.T) {
	iterator := seq.Take(1).Operator[int](nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestTake2(t *testing.T) {
	pairs := []pair[int, int]{
		{Key: 1, Value: 10},
		{Key: 0, Value: 20},
		{Key: 1, Value: 30},
	}
	for _, tc := range []struct {
		name   string
		n      int
		want   []pair[int, int]
		reads  int
		starts int
		stops  int
	}{
		{name: "negative", n: -1},
		{name: "zero", n: 0},
		{name: "one", n: 1, want: pairs[:1], reads: 1, starts: 1, stops: 1},
		{name: "prefix", n: 2, want: pairs[:2], reads: 2, starts: 1, stops: 1},
		{name: "all", n: 3, want: pairs[:3], reads: 3, starts: 1, stops: 1},
		{name: "larger", n: 4, want: pairs[:3], reads: 3, starts: 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			reads, starts, stops := 0, 0, 0
			src := func(yield func(int, int) bool) {
				starts++
				for i := 1; i <= 3; i++ {
					reads++
					if !yield(i%2, i*10) {
						stops++
						return
					}
				}
			}
			pipeline := seq.FromSeq2(src).Pipe2(seq.Take2(tc.n).Operator)
			iterator := seq.ToSeq2(pipeline)
			assertEqual(t, []int{reads, starts, stops}, []int{0, 0, 0})
			for round := 1; round <= 2; round++ {
				assertEqual(t, collectPairs(iterator), tc.want)
				gotCounts := []int{reads, starts, stops}
				wantCounts := []int{tc.reads * round, tc.starts * round, tc.stops * round}
				assertEqual(t, gotCounts, wantCounts)
			}
		})
	}
}

func TestTake2EmptyAndNilInput(t *testing.T) {
	for _, src := range []seq.Seq2[int, string]{nil, seq.FromSeq2(slices.All([]string{}))} {
		iterator := seq.Take2(1).Operator(src)
		for round := 0; round < 2; round++ {
			assertEqual(t, collectPairs(seq.ToSeq2(iterator)), []pair[int, string](nil))
		}
	}
}

func TestTake2DownstreamStop(t *testing.T) {
	reads, stops := 0, 0
	src := func(yield func(int, int) bool) {
		for i := 1; i <= 3; i++ {
			reads++
			if !yield(i, i*10) {
				stops++
				return
			}
		}
	}
	pipeline := seq.FromSeq2(src).Pipe2(seq.Take2(3).Operator)
	for k, v := range seq.ToSeq2(pipeline) {
		assertEqual(t, k, 1)
		assertEqual(t, v, 10)
		break
	}
	assertEqual(t, []int{reads, stops}, []int{1, 1})
}
