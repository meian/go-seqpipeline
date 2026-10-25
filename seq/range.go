package seq

import (
	"math"
	"reflect"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// FromRange は start から stop に向かう整数列を step ごとに生成する。
// uintptr を含む組み込み整数型と、それらを基底型とする定義型を受け付ける。
// step の符号で昇順・降順を決める。
// 到達した stop は含み、飛び越す前や整数型の範囲を超える前に終了する。
// start と stop が等しい場合は 1 件返し、step の方向が逆の場合は空の非 nil の Seq を返す。
// step が 0 の場合は呼び出し時に panic する。
// step は int64 で、表せない増減値は指定できない。
// 反復のたびに start から始める。
func FromRange[T integer](start, stop T, step int64) Seq[T] {
	if step == 0 {
		panic("seq: FromRange step must not be zero")
	}

	// Kind による分岐では T の制約が狭まらないため、両経路とも integer を使う。
	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromSignedRange[T](int64(start), int64(stop), step)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fromUnsignedRange[T](uint64(start), uint64(stop), step)
	default:
		panic("seq: unsupported integer kind")
	}
}

func fromSignedRange[T integer](start, stop int64, step int64) Seq[T] {
	// stop - step が int64 の範囲外になる場合、次の値は生成できない。
	if step > 0 {
		return func(yield func(T) bool) {
			if start > stop {
				return
			}
			if stop < math.MinInt64+step {
				yield(T(start))
				return
			}
			guard := stop - step
			current := start
			for current <= guard {
				if !yield(T(current)) {
					return
				}
				current += step
			}
			if current <= stop {
				yield(T(current))
			}
		}
	}

	return func(yield func(T) bool) {
		if start < stop {
			return
		}
		if stop > math.MaxInt64+step {
			yield(T(start))
			return
		}
		guard := stop - step
		current := start
		for current >= guard {
			if !yield(T(current)) {
				return
			}
			current += step
		}
		if current >= stop {
			yield(T(current))
		}
	}
}

func fromUnsignedRange[T integer](start, stop uint64, step int64) Seq[T] {
	if step > 0 {
		increment := uint64(step)
		return func(yield func(T) bool) {
			if start > stop {
				return
			}
			current := start
			for {
				if !yield(T(current)) {
					return
				}
				if stop-current < increment {
					return
				}
				current += increment
			}
		}
	}

	decrement := uint64(-(step + 1)) + 1
	return func(yield func(T) bool) {
		if start < stop {
			return
		}
		current := start
		for {
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
