package common

import (
	"strconv"
)

// MustAtoi converts the input string to an integer and panics on failure.
func MustAtoi(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return res
}
