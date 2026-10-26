package seq_test

import (
	"math"
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFromRange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		start, stop int
		step        int64
		want        []int
	}{
		{name: "ascending endpoint", start: 1, stop: 7, step: 2, want: []int{1, 3, 5, 7}},
		{name: "ascending overshoot", start: 1, stop: 8, step: 3, want: []int{1, 4, 7}},
		{name: "descending endpoint", start: 7, stop: 1, step: -2, want: []int{7, 5, 3, 1}},
		{name: "descending overshoot", start: 8, stop: 1, step: -3, want: []int{8, 5, 2}},
		{name: "equal ascending", start: 4, stop: 4, step: 2, want: []int{4}},
		{name: "equal descending", start: 4, stop: 4, step: -2, want: []int{4}},
		{name: "away ascending", start: 7, stop: 1, step: 2},
		{name: "away descending", start: 1, stop: 7, step: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			source := seq.FromRange(tt.start, tt.stop, tt.step)
			if source == nil {
				t.Fatal("FromRange returned a nil Seq")
			}
			assertEqual(t, seq.ToSlice(source), tt.want)
		})
	}
}

func TestFromRangeZeroStep(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		start, stop int
	}{
		{name: "equal", start: 1, stop: 1},
		{name: "ascending", start: 1, stop: 3},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if recover() == nil {
					t.Fatal("FromRange with step 0 must panic at the call")
				}
			}()
			_ = seq.FromRange(tt.start, tt.stop, 0)
		})
	}
}

func TestFromRangeIntegerKinds(t *testing.T) {
	t.Parallel()

	assertEqual(t, seq.ToSlice(seq.FromRange(int8(-2), int8(2), 2)), []int8{-2, 0, 2})
	assertEqual(t, seq.ToSlice(seq.FromRange(int16(3), int16(-1), -2)), []int16{3, 1, -1})
	assertEqual(t, seq.ToSlice(seq.FromRange(int32(-2), int32(2), 2)), []int32{-2, 0, 2})
	assertEqual(t, seq.ToSlice(seq.FromRange(int64(3), int64(-1), -2)), []int64{3, 1, -1})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint(4), uint(0), -2)), []uint{4, 2, 0})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint8(4), uint8(0), -2)), []uint8{4, 2, 0})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint16(0), uint16(4), 2)), []uint16{0, 2, 4})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint32(4), uint32(0), -2)), []uint32{4, 2, 0})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint64(0), uint64(4), 2)), []uint64{0, 2, 4})
	assertEqual(t, seq.ToSlice(seq.FromRange(uintptr(0), uintptr(4), 2)), []uintptr{0, 2, 4})
}

func TestFromRangeDefinedTypes(t *testing.T) {
	t.Parallel()

	type signed int16
	type unsigned uint32
	type pointer uintptr

	assertEqual(t, seq.ToSlice(seq.FromRange(signed(-2), signed(2), 2)), []signed{-2, 0, 2})
	assertEqual(t, seq.ToSlice(seq.FromRange(unsigned(4), unsigned(0), -2)), []unsigned{4, 2, 0})
	assertEqual(t, seq.ToSlice(seq.FromRange(pointer(1), pointer(5), 2)), []pointer{1, 3, 5})
}

// 狭い要素型を超える step と各型の端点でも、値を切り詰めたり折り返したりしない。
func TestFromRangeBounds(t *testing.T) {
	t.Parallel()

	assertEqual(t, seq.ToSlice(seq.FromRange(int8(math.MinInt8), int8(math.MaxInt8), 256)), []int8{math.MinInt8})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint8(math.MaxUint8), uint8(0), -256)), []uint8{math.MaxUint8})
	assertEqual(t, seq.ToSlice(seq.FromRange(int(math.MaxInt-1), int(math.MaxInt), 1)), []int{math.MaxInt - 1, math.MaxInt})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint(math.MaxUint-1), uint(math.MaxUint), 1)), []uint{math.MaxUint - 1, math.MaxUint})
	assertEqual(t, seq.ToSlice(seq.FromRange(int64(math.MaxInt64-1), int64(math.MaxInt64), 1)), []int64{math.MaxInt64 - 1, math.MaxInt64})
	assertEqual(t, seq.ToSlice(seq.FromRange(int64(math.MinInt64+1), int64(math.MinInt64), -1)), []int64{math.MinInt64 + 1, math.MinInt64})
	// 境界の等号と、その直前で初回のみ返す場合を両方向で確認する。
	min, max := int64(math.MinInt64), int64(math.MaxInt64)
	assertEqual(t, seq.ToSlice(seq.FromRange(min, -1, max)), []int64{min, -1})
	assertEqual(t, seq.ToSlice(seq.FromRange(min, -2, max)), []int64{min})
	assertEqual(t, seq.ToSlice(seq.FromRange(min, max, max)), []int64{min, -1, max - 1})
	assertEqual(t, seq.ToSlice(seq.FromRange(max, -1, min)), []int64{max, -1})
	assertEqual(t, seq.ToSlice(seq.FromRange(max, 0, min)), []int64{max})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint64(math.MaxUint64-1), uint64(math.MaxUint64), 1)), []uint64{math.MaxUint64 - 1, math.MaxUint64})
	assertEqual(t, seq.ToSlice(seq.FromRange(uint64(1), uint64(0), -1)), []uint64{1, 0})

	bigStep := int64(1) << 40
	assertEqual(t, seq.ToSlice(seq.FromRange(int64(0), bigStep, bigStep)), []int64{0, bigStep})
	minStep := int64(math.MinInt64)
	assertEqual(t, seq.ToSlice(seq.FromRange(int64(0), minStep, minStep)), []int64{0, minStep})
	unsignedStart := uint64(math.MaxInt64) + 1
	assertEqual(t, seq.ToSlice(seq.FromRange(unsignedStart, uint64(0), minStep)), []uint64{unsignedStart, 0})
}

// 全幅の残距離を扱い、下流停止後も次の反復は初期値から始まる。
func TestFromRangeStopAndRepeat(t *testing.T) {
	t.Parallel()

	source := seq.FromRange(int64(math.MinInt64), int64(math.MaxInt64), 1)
	for i := 0; i < 2; i++ {
		var got []int64
		source(func(v int64) bool {
			if len(got) == 2 {
				t.Fatal("FromRange yielded after false")
			}
			got = append(got, v)
			return len(got) < 2
		})
		assertEqual(t, got, []int64{math.MinInt64, math.MinInt64 + 1})
	}
}
