package main

import (
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
)

func patch(input []int, width int) {
	// Insert two quarters
	input[0] = 2

	// Find the bottom line of the play area and change it to wall
	for i := len(input) - 1; i > 0; i-- {
		if input[i] == 1 && input[i-width] == 1 {
			allEmpty := true
			for j := i - width + 1; j < i-1; j++ {
				if input[j] != 0 {
					allEmpty = false
					break
				}
			}

			if allEmpty {
				for j := i - width + 1; j < i-1; j++ {
					input[j] = 1
				}
				break
			}
		}
	}
}

func feedInput(input chan int) {
	for {
		input <- 0
	}
}

func readOutput(output chan int, handler func(int, int, int)) {
	for {
		x, more := <-output
		if !more {
			break
		}
		y := <-output
		tile := <-output
		handler(x, y, tile)
	}
}

func countBlocks(screen [60][80]int) (blocks int) {
	for _, line := range screen {
		for _, pixel := range line {
			if pixel == 2 {
				blocks++
			}
		}
	}
	return
}

func run(input []int, handler func(int, int, int)) {
	memory := make([]int, len(input))
	copy(memory, input)
	vm := intcode.NewVirtualMachine(memory)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 6)

	go feedInput(vm.Input)
	go vm.Run()

	readOutput(vm.Output, handler)
}

func main() {
	input := common.ReadCsvAsInts("13/input.txt")
	memory := make([]int, len(input))
	copy(memory, input)

	width := 0
	var screen [60][80]int
	run(memory, func(x int, y int, tile int) {
		screen[y][x] = tile
		if x > width {
			width = x
		}
	})

	println(countBlocks(screen))

	patch(input, width)

	score := 0
	run(input, func(x int, y int, tile int) {
		if x == -1 {
			score = tile
		}
	})

	println(score)
}
