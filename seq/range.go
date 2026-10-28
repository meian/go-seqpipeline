package seq

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

	return func(yield func(T) bool) {
		remaining, ok := rangeSteps(start, stop, step)
		if !ok {
			return
		}
		current := start
		for {
			if !yield(current) || remaining == 0 {
				return
			}
			// 加算結果のビット列は符号の有無で変わらないため、uint64 で足して T に戻せる。
			// 残り回数で終点までの更新に限るため、次の値は T の範囲に収まる。
			current = T(uint64(current) + uint64(step))
			remaining--
		}
	}
}

// 要素数ではなく遷移回数を使い、2^64 要素の範囲でも uint64 に収める。
func rangeSteps[T integer](start, stop T, step int64) (uint64, bool) {
	// 向きを T で確認した後の uint64 の差は、負数を含んでも元の整数間の距離になる。
	if step > 0 {
		if start > stop {
			return 0, false
		}
		return (uint64(stop) - uint64(start)) / uint64(step), true
	}
	if start < stop {
		return 0, false
	}
	decrement := uint64(-(step + 1)) + 1
	return (uint64(start) - uint64(stop)) / decrement, true
}
