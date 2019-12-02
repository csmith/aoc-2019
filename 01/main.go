package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
)

func fuel(mass int) int {
	return (mass / 3) - 2
}

func fuelRequirements(input []int) (int, int) {
	var (
		simpleTotal    = 0
		recursiveTotal = 0
	)

	for _, mass := range input {
		required := fuel(mass)
		simpleTotal += required
		for required > 0 {
			recursiveTotal += required
			required = fuel(required)
		}
	}
	return simpleTotal, recursiveTotal
}

func main() {
	input := common.Atoi(common.ReadLines("01/input.txt"))
	part1, part2 := fuelRequirements(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
