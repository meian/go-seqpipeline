package seq

import (
	"math"
	"reflect"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// FromRange generates an integer sequence from start toward stop in steps of step.
// It accepts built-in integer types, including uintptr, and defined types with
// those underlying types.
// The sign of step determines the direction.
// The sequence includes stop if reached and ends before overshooting it or
// exceeding the integer type's range.
// When start equals stop, it yields one value; when step points away from stop,
// it returns an empty, non-nil Seq.
// FromRange panics if step is zero.
// Step has type int64, so increments or decrements outside its range cannot be
// specified.
// Each iteration begins at start.
func FromRange[T integer](start, stop T, step int64) Seq[T] {
	if step == 0 {
		panic("seq: FromRange step must not be zero")
	}

	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromSignedRange[T](int64(start), int64(stop), step)
	}
	return fromUnsignedRange[T](uint64(start), uint64(stop), step)
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
