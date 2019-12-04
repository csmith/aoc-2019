package main

import (
	"github.com/csmith/aoc-2019/common"
	"strings"
)

func matchesRules(input int) (hasAnyRun bool, hasTwoRun bool, next int) {
	last := 10
	run := 0

	for input > 0 {
		digit := input % 10

		if digit > last {
			// Simple optimisation: if we hit a digit that's larger than the previous (running right-to-left) then
			// we can predict the next possible number that might match. e.g. 1234111 => 1234444.
			for input < 100000 {
				input = 10*input + digit
			}
			return false, false, input
		} else if digit == last {
			run++
		} else {
			hasTwoRun = hasTwoRun || run == 1
			hasAnyRun = hasAnyRun || run > 0
			last = digit
			run = 0
		}

		input = input / 10
	}

	hasTwoRun = hasTwoRun || run == 1
	hasAnyRun = hasAnyRun || run > 0
	return
}

func main() {
	var (
		input = strings.Split(common.ReadFileAsStrings("04/input.txt")[0], "-")
		from  = common.MustAtoi(input[0])
		to    = common.MustAtoi(input[1])
		part1 = 0
		part2 = 0
	)

	for i := from; i <= to; i++ {
		p1, p2, n := matchesRules(i)
		if p1 {
			part1++
		}
		if p2 {
			part2++
		}
		if n > 0 {
			i = n - 1
		}
	}

	println(part1)
	println(part2)
}
