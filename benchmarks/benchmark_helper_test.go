package benchmarks

type pair struct{ key, value int }

var output []pair
var checksum int64

func input(n int) []int {
	values := make([]int, n)
	for i := range values {
		values[i] = i
	}
	return values
}

func mapValue(v int) int   { return (v + 1) & 65535 }
func keepValue(v int) bool { return v%3 != 0 }
