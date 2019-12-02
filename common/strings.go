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

// Atoi converts a slice of strings to a slice of integers.
// It panics if any element couldn't be converted.
func Atoi(input []string) []int {
	var result []int
	for _, element := range input {
		result = append(result, MustAtoi(element))
	}
	return result
}
