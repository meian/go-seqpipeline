package seq

import "reflect"

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
		return rangeInt(start, stop, step)
	}

	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int:
		return rangeInt(start, stop, step)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return rangeGeneric(start, stop, step)
	default:
		panic("seq.Range: unsupported integer type")
	}
}

func rangeInt[T Integer](start, stop, step T) Seq[T] {
	return func(yield func(T) bool) {
		current, end, increment := int(start), int(stop), int(step)
		if increment > 0 {
			for current < end {
				if !yield(T(current)) {
					return
				}
				next := current + increment
				if next <= current {
					return
				}
				current = next
			}
			return
		}
		for current > end {
			if !yield(T(current)) {
				return
			}
			next := current + increment
			if next >= current {
				return
			}
			current = next
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
