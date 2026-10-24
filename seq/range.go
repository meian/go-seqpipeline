package seq

import (
	"math"
	"reflect"
)

// Integer は Range で扱う整数型を表す。
// 基底型が整数型である定義型も含む。
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Range は start から step ずつ増減する整数を、stop まで生成する。
// step は int 型で、符号なし整数型にも負の値を指定できる。
// stop にちょうど達した場合は含み、飛び越した場合は含まない。
// start と stop が等しい場合は start を一度生成し、step の方向と stop の位置が合わない場合は空になる。
// step が 0 の場合は panic する。
// 次の値が stop を越えるか、整数型の範囲外になる場合は生成を終了する。
func Range[T Integer](start, stop T, step int) Seq[T] {
	if step == 0 {
		panic("seq.Range: step must not be zero")
	}

	switch any(start).(type) {
	case int:
		return rangeInt[T](int(start), int(stop), step)
	case int16:
		return rangeInt16[T](int16(start), int16(stop), step)
	case uint8:
		return rangeUint8[T](uint8(start), uint8(stop), step)
	}

	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int:
		return rangeInt[T](int(start), int(stop), step)
	case reflect.Int16:
		return rangeInt16[T](int16(start), int16(stop), step)
	case reflect.Uint8:
		return rangeUint8[T](uint8(start), uint8(stop), step)
	case reflect.Int8, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return rangeGeneric(start, stop, step)
	default:
		panic("seq.Range: unsupported integer type")
	}
}

// rangeInt は T の基底型が int の場合だけ呼び出す。
func rangeInt[T Integer](start, stop, step int) Seq[T] {
	return func(yield func(T) bool) {
		current, end, increment := start, stop, step
		// end - increment が int の範囲外なら、次の値は stop を越える。
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

// rangeInt16 は T の基底型が int16 の場合だけ呼び出す。
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

// rangeUint8 は T の基底型が uint8 の場合だけ呼び出す。
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

func rangeGeneric[T Integer](start, stop T, step int) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		// uint64 の差は符号付き型の全範囲にまたがる距離も表せる。
		if step > 0 {
			increment := uint64(step)
			for current <= stop {
				if !yield(current) {
					return
				}
				if uint64(stop)-uint64(current) < increment {
					return
				}
				current = T(uint64(current) + increment)
			}
			return
		}
		// step == MinInt でも絶対値を uint64 で表せる。
		decrement := uint64(-(step + 1)) + 1
		for current >= stop {
			if !yield(current) {
				return
			}
			if uint64(current)-uint64(stop) < decrement {
				return
			}
			current = T(uint64(current) - decrement)
		}
	}
}
