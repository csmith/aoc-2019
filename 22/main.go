package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
)

func main() {
	input := common.ReadFileAsStrings("22/input.txt")
	deckSize := 10007
	position := 2019
	for _, line := range input {
		var n int
		if "deal into new stack" == line {
			position = deckSize - 1 - position
		} else if _, err := fmt.Sscanf(line, "cut %d", &n); err == nil {
			n = (n + deckSize) % deckSize
			if position < n {
				position += deckSize - n
			} else {
				position -= n
			}
		} else if _, err := fmt.Sscanf(line, "deal with increment %d", &n); err == nil {
			position = (position * n) % deckSize
		} else {
			panic(fmt.Sprintf("Unrecognised line: %s", line))
		}
	}
	println(position)
}
