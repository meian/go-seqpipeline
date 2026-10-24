package seq

import (
	"math"
	"reflect"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Range[T Integer](start, stop T, step int) Seq[T] {
	if step == 0 {
		panic("seq.Range: step must not be zero")
	}

	switch any(start).(type) {
	case int:
		return rangeInt[T](int(start), int(stop), step)
	case int8:
		return rangeInt8[T](int8(start), int8(stop), step)
	case int16:
		return rangeInt16[T](int16(start), int16(stop), step)
	case int32:
		return rangeInt32[T](int32(start), int32(stop), step)
	case int64:
		return rangeInt64[T](int64(start), int64(stop), step)
	case uint:
		return rangeUint[T](uint(start), uint(stop), step)
	case uint8:
		return rangeUint8[T](uint8(start), uint8(stop), step)
	case uint16:
		return rangeUint16[T](uint16(start), uint16(stop), step)
	case uint32:
		return rangeUint32[T](uint32(start), uint32(stop), step)
	case uint64:
		return rangeUint64[T](uint64(start), uint64(stop), step)
	case uintptr:
		return rangeUintptr[T](uintptr(start), uintptr(stop), step)
	}

	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int:
		return rangeInt[T](int(start), int(stop), step)
	case reflect.Int8:
		return rangeInt8[T](int8(start), int8(stop), step)
	case reflect.Int16:
		return rangeInt16[T](int16(start), int16(stop), step)
	case reflect.Int32:
		return rangeInt32[T](int32(start), int32(stop), step)
	case reflect.Int64:
		return rangeInt64[T](int64(start), int64(stop), step)
	case reflect.Uint:
		return rangeUint[T](uint(start), uint(stop), step)
	case reflect.Uint8:
		return rangeUint8[T](uint8(start), uint8(stop), step)
	case reflect.Uint16:
		return rangeUint16[T](uint16(start), uint16(stop), step)
	case reflect.Uint32:
		return rangeUint32[T](uint32(start), uint32(stop), step)
	case reflect.Uint64:
		return rangeUint64[T](uint64(start), uint64(stop), step)
	case reflect.Uintptr:
		return rangeUintptr[T](uintptr(start), uintptr(stop), step)
	default:
		panic("seq.Range: unsupported integer type")
	}
}

func rangeInt[T Integer](start, stop, step int) Seq[T] {
	return func(yield func(T) bool) {
		current, end, increment := start, stop, step
		if increment > 0 {
			if current > end {
				return
			}
			if end < math.MinInt+increment {
				yield(T(current))
				return
			}
			guard := end - increment
			for current <= guard {
				if !yield(T(current)) {
					return
				}
				current += increment
			}
			if current <= end {
				yield(T(current))
			}
			return
		}
		if current < end {
			return
		}
		if end > math.MaxInt+increment {
			yield(T(current))
			return
		}
		guard := end - increment
		for current >= guard {
			if !yield(T(current)) {
				return
			}
			current += increment
		}
		if current >= end {
			yield(T(current))
		}
	}
}

func rangeInt16[T Integer](start, stop int16, step int) Seq[T] {
	return func(yield func(T) bool) {
		current, end, increment := int(start), int(stop), step
		if increment > 0 {
			if current > end {
				return
			}
			if increment > end-current {
				yield(T(current))
				return
			}
			guard := end - increment
			for current <= guard {
				if !yield(T(current)) {
					return
				}
				current += increment
			}
			if current <= end {
				yield(T(current))
			}
			return
		}
		if current < end {
			return
		}
		if increment < end-current {
			yield(T(current))
			return
		}
		guard := end - increment
		for current >= guard {
			if !yield(T(current)) {
				return
			}
			current += increment
		}
		if current >= end {
			yield(T(current))
		}
	}
}

func rangeUint8[T Integer](start, stop uint8, step int) Seq[T] {
	return func(yield func(T) bool) {
		current, end, increment := int(start), int(stop), step
		if increment > 0 {
			if current > end {
				return
			}
			if increment > end-current {
				yield(T(current))
				return
			}
			guard := end - increment
			for current <= guard {
				if !yield(T(current)) {
					return
				}
				current += increment
			}
			if current <= end {
				yield(T(current))
			}
			return
		}
		if current < end {
			return
		}
		if increment < end-current {
			yield(T(current))
			return
		}
		guard := end - increment
		for current >= guard {
			if !yield(T(current)) {
				return
			}
			current += increment
		}
		if current >= end {
			yield(T(current))
		}
	}
}

func rangeInt8[T Integer](start, stop int8, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = int8(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = int8(uint64(current) - decrement)
		}
	}
}

func rangeInt32[T Integer](start, stop int32, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = int32(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = int32(uint64(current) - decrement)
		}
	}
}

func rangeInt64[T Integer](start, stop int64, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = int64(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = int64(uint64(current) - decrement)
		}
	}
}

func rangeUint[T Integer](start, stop uint, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = uint(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = uint(uint64(current) - decrement)
		}
	}
}

func rangeUint16[T Integer](start, stop uint16, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = uint16(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = uint16(uint64(current) - decrement)
		}
	}
}

func rangeUint32[T Integer](start, stop uint32, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = uint32(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = uint32(uint64(current) - decrement)
		}
	}
}

func rangeUint64[T Integer](start, stop uint64, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if stop-current < increment {
					return
				}
				current += increment
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if current-stop < decrement {
				return
			}
			current -= decrement
		}
	}
}

func rangeUintptr[T Integer](start, stop uintptr, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(T(current)) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = uintptr(uint64(current) + increment)
			}
			return
		}
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(T(current)) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = uintptr(uint64(current) - decrement)
		}
	}
}
