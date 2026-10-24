package seq_test

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestRange(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(1, 3, 1)), []int{1, 2, 3})
	})
	t.Run("int8", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(int8(3), int8(1), -1)), []int8{3, 2, 1})
	})
	t.Run("int16", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(int16(3), int16(1), -1)), []int16{3, 2, 1})
	})
	t.Run("int32", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(int32(1), int32(3), 1)), []int32{1, 2, 3})
	})
	t.Run("int64", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(int64(3), int64(1), -1)), []int64{3, 2, 1})
	})
	t.Run("uint", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(uint(3), uint(1), -1)), []uint{3, 2, 1})
	})
	t.Run("uint8", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(uint8(3), uint8(1), -1)), []uint8{3, 2, 1})
	})
	t.Run("uint16", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(uint16(3), uint16(1), -1)), []uint16{3, 2, 1})
	})
	t.Run("uint32", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(uint32(3), uint32(1), -1)), []uint32{3, 2, 1})
	})
	t.Run("uint64", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(uint64(3), uint64(1), -1)), []uint64{3, 2, 1})
	})
	t.Run("uintptr", func(t *testing.T) {
		assertEqual(t, seq.ToSlice(seq.Range(uintptr(3), uintptr(1), -1)), []uintptr{3, 2, 1})
	})
	t.Run("defined int", func(t *testing.T) {
		type definedInt int
		assertEqual(t, seq.ToSlice(seq.Range(definedInt(1), definedInt(3), 1)), []definedInt{1, 2, 3})
	})
	t.Run("defined int16", func(t *testing.T) {
		type definedInt16 int16
		assertEqual(t, seq.ToSlice(seq.Range(definedInt16(3), definedInt16(1), -1)), []definedInt16{3, 2, 1})
	})
	t.Run("defined uint8", func(t *testing.T) {
		type definedUint8 uint8
		assertEqual(t, seq.ToSlice(seq.Range(definedUint8(3), definedUint8(1), -1)), []definedUint8{3, 2, 1})
	})
}
