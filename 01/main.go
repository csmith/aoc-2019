package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"strconv"
)

func fuel(mass int) int {
	return (mass / 3) - 2
}

func fuelRequirements(input []string) (int, int) {
	var (
		simpleTotal = 0
		recursiveTotal = 0
	)

	for _, line := range input {
		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

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
	input := common.ReadLines("01/input.txt")
	part1, part2 := fuelRequirements(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
