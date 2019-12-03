package common

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	y := x >> 63
	return (x ^ y) - y
}
