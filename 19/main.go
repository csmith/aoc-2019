package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
)

const shipSize = 100

func main() {
	input := common.ReadCsvAsInts("19/input.txt")
	vm := intcode.NewVirtualMachine(make([]int, len(input)))

	var endPositions [shipSize]int
	endIndex := 0

	sum := 0
	lastStart := 0
	lastWidth := 0
	for y := 0; ; y++ {
		first := 0

		for x := lastStart; ; x++ {
			vm.Reset(input)
			if *vm.RunForInput(x, y) == 1 {
				if first == 0 {
					first = x
					if y > 50 {
						x += lastWidth
					}
				}
				if x < 50 && y < 50 {
					sum++
				}
			} else if first > 0 || y < 50 && x > 50 {
				// For y < 50 if we hit x > 50 without finding the beam, assume it's not there. (The beam isn't
				// continuous at the start, there's a couple of blanks).

				if x-first >= shipSize-1 {
					previousEnd := endPositions[endIndex]

					if previousEnd-shipSize >= first-1 {
						fmt.Println(sum)
						fmt.Println((previousEnd-shipSize)*10000 + (y - 100))
						return
					}

					endPositions[endIndex] = x
					endIndex = (endIndex + 1) % shipSize
				}

				lastWidth = x - first - 1
				break
			}
		}

		lastStart = first
	}
}
