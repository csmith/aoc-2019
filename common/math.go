package common

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	y := x >> 63
	return (x ^ y) - y
}

// Max returns the highest of the two given ints.
func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// Min returns the lowest of the two given ints.
func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
