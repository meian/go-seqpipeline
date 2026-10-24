package seq_test

import (
	"math"
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
		{name: "ascending reaches stop", start: 1, stop: 7, step: 2, want: []int{1, 3, 5, 7}},
		{name: "descending", start: 8, stop: 1, step: -2, want: []int{8, 6, 4, 2}},
		{name: "descending reaches stop", start: 7, stop: 1, step: -2, want: []int{7, 5, 3, 1}},
		{name: "equal bounds", start: 3, stop: 3, step: 2, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqual(t, seq.ToSlice(seq.Range(tt.start, tt.stop, tt.step)), tt.want)
		})
	}
	t.Run("defined int", func(t *testing.T) {
		type definedInt int
		got := seq.ToSlice(seq.Range(definedInt(1), definedInt(6), 2))
		assertEqual(t, got, []definedInt{1, 3, 5})
	})
}

func TestRangeIntBounds(t *testing.T) {
	tests := []struct {
		name              string
		start, stop, step int
		want              []int
	}{
		{name: "ascending guard underflow", start: math.MinInt, stop: math.MinInt + 1, step: math.MaxInt, want: []int{math.MinInt}},
		{name: "ascending across bounds", start: math.MinInt, stop: math.MaxInt, step: math.MaxInt, want: []int{math.MinInt, -1, math.MaxInt - 1}},
		{name: "ascending reaches max", start: math.MaxInt - 1, stop: math.MaxInt, step: 1, want: []int{math.MaxInt - 1, math.MaxInt}},
		{name: "descending guard overflow", start: math.MaxInt, stop: math.MaxInt - 1, step: math.MinInt, want: []int{math.MaxInt}},
		{name: "descending across bounds", start: math.MaxInt, stop: math.MinInt, step: math.MinInt, want: []int{math.MaxInt, -1}},
		{name: "descending reaches min", start: math.MinInt + 1, stop: math.MinInt, step: -1, want: []int{math.MinInt + 1, math.MinInt}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqual(t, seq.ToSlice(seq.Range(tt.start, tt.stop, tt.step)), tt.want)
		})
	}
}

func TestRangeInt16(t *testing.T) {
	tests := []struct {
		name        string
		start, stop int16
		step        int
		want        []int16
	}{
		{name: "ascending", start: 1, stop: 8, step: 2, want: []int16{1, 3, 5, 7}},
		{name: "ascending reaches stop", start: 1, stop: 7, step: 2, want: []int16{1, 3, 5, 7}},
		{name: "descending", start: 8, stop: 1, step: -2, want: []int16{8, 6, 4, 2}},
		{name: "descending reaches stop", start: 7, stop: 1, step: -2, want: []int16{7, 5, 3, 1}},
		{name: "equal bounds", start: 3, stop: 3, step: -2, want: []int16{3}},
		{name: "ascending step exceeds distance", start: math.MinInt16, stop: math.MinInt16 + 1, step: math.MaxInt16, want: []int16{math.MinInt16}},
		{name: "ascending across bounds", start: math.MinInt16, stop: math.MaxInt16, step: math.MaxInt16, want: []int16{math.MinInt16, -1, math.MaxInt16 - 1}},
		{name: "ascending reaches max", start: math.MaxInt16 - 1, stop: math.MaxInt16, step: 1, want: []int16{math.MaxInt16 - 1, math.MaxInt16}},
		{name: "descending step exceeds distance", start: math.MaxInt16, stop: math.MaxInt16 - 1, step: math.MinInt16, want: []int16{math.MaxInt16}},
		{name: "descending across bounds", start: math.MaxInt16, stop: math.MinInt16, step: math.MinInt16, want: []int16{math.MaxInt16, -1}},
		{name: "descending reaches min", start: math.MinInt16 + 1, stop: math.MinInt16, step: -1, want: []int16{math.MinInt16 + 1, math.MinInt16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqual(t, seq.ToSlice(seq.Range(tt.start, tt.stop, tt.step)), tt.want)
		})
	}
	t.Run("defined int16", func(t *testing.T) {
		type definedInt16 int16
		got := seq.ToSlice(seq.Range(definedInt16(1), definedInt16(6), 2))
		assertEqual(t, got, []definedInt16{1, 3, 5})
	})
}

func TestRangeUint8(t *testing.T) {
	tests := []struct {
		name        string
		start, stop uint8
		step        int
		want        []uint8
	}{
		{name: "reaches stop", start: 1, stop: 7, step: 2, want: []uint8{1, 3, 5, 7}},
		{name: "passes stop", start: 1, stop: 8, step: 2, want: []uint8{1, 3, 5, 7}},
		{name: "guard underflow", start: 0, stop: 1, step: 255, want: []uint8{0}},
		{name: "reaches max", start: 254, stop: 255, step: 1, want: []uint8{254, 255}},
		{name: "passes max", start: 254, stop: 255, step: 2, want: []uint8{254}},
		{name: "descending reaches zero", start: 4, stop: 0, step: -2, want: []uint8{4, 2, 0}},
		{name: "equal bounds", start: 3, stop: 3, step: 2, want: []uint8{3}},
		{name: "start beyond stop", start: 8, stop: 1, step: 2, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqual(t, seq.ToSlice(seq.Range(tt.start, tt.stop, tt.step)), tt.want)
		})
	}
	t.Run("defined uint8", func(t *testing.T) {
		type definedUint8 uint8
		got := seq.ToSlice(seq.Range(definedUint8(255), definedUint8(253), -1))
		assertEqual(t, got, []definedUint8{255, 254, 253})
	})
}

func TestRangeOtherIntegers(t *testing.T) {
	assertEqual(t, seq.ToSlice(seq.Range(int8(-127), int8(-128), -1)), []int8{-127, -128})
	assertEqual(t, seq.ToSlice(seq.Range(uint16(8), uint16(2), -2)), []uint16{8, 6, 4, 2})
}
