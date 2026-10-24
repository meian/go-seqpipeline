package seq_test

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name              string
		start, stop, step int
		want              []int
	}{
		{name: "ascending", start: 1, stop: 8, step: 2, want: []int{1, 3, 5, 7}},
		{name: "descending", start: 8, stop: 1, step: -2, want: []int{8, 6, 4, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqual(t, seq.ToSlice(seq.Range(tt.start, tt.stop, tt.step)), tt.want)
		})
	}
	t.Run("defined int", func(t *testing.T) {
		type definedInt int
		got := seq.ToSlice(seq.Range(definedInt(1), definedInt(6), definedInt(2)))
		assertEqual(t, got, []definedInt{1, 3, 5})
	})
}
