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

// Range は start から step ずつ増減する整数を、stop の手前まで生成する。
// stop は含まず、step の方向と stop の位置が合わない場合は空になる。
// step が 0 の場合は panic する。
// 加算で値が進まないか、整数のオーバーフローで方向が逆転した場合は生成を終了する。
func Range[T Integer](start, stop, step T) Seq[T] {
	if step == 0 {
		panic("seq.Range: step must not be zero")
	}

	switch any(start).(type) {
	case int:
		return rangeInt[T](int(start), int(stop), int(step))
	}

	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int:
		return rangeInt[T](int(start), int(stop), int(step))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
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
			if current >= end {
				return
			}
			if end < math.MinInt+increment {
				yield(T(current))
				return
			}
			guard := end - increment
			for current < guard {
				if !yield(T(current)) {
					return
				}
				current += increment
			}
			if current < end {
				yield(T(current))
			}
			return
		}
		if current <= end {
			return
		}
		if end > math.MaxInt+increment {
			yield(T(current))
			return
		}
		guard := end - increment
		for current > guard {
			if !yield(T(current)) {
				return
			}
			current += increment
		}
		if current > end {
			yield(T(current))
		}
	}
}

func rangeGeneric[T Integer](start, stop, step T) Seq[T] {
	return func(yield func(T) bool) {
		current := start
		if step > 0 {
			for current < stop {
				if !yield(current) {
					return
				}
				next := current + step
				if next <= current {
					return
				}
				current = next
			}
			return
		}
		if step < 0 {
			for current > stop {
				if !yield(current) {
					return
				}
				next := current + step
				if next >= current {
					return
				}
				current = next
			}
		}
	}
}
