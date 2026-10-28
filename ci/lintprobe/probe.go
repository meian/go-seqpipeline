package lintprobe

func Narrow(value int64) int8 {
	return int8(value)
}

func Warning(value int) int {
	var result int = value + 1
	return result
}
