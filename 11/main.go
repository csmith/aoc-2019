package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
)

func run(memory []int, startWhite bool) (int, [500][500]bool) {
	input := make([]int, len(memory))
	copy(input, memory)
	vm := intcode.NewVirtualMachine(input)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	go vm.Run()

	var grid [500][500]bool
	var painted [500][500]bool

	paintCount := 0
	x := 250
	y := 250
	dir := 0
	grid[y][x] = startWhite

	for {
		if grid[y][x] {
			vm.Input <- 1
		} else {
			vm.Input <- 0
		}

		paint, more := <-vm.Output
		if !more {
			break
		}
		turn := <-vm.Output

		grid[y][x] = paint == 1
		if !painted[y][x] {
			paintCount++
			painted[y][x] = true
		}

		if turn == 1 {
			dir = (dir + 1) % 4
		} else {
			dir = (dir + 3) % 4
		}

		switch dir {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return paintCount, grid
}

func main() {
	input := common.ReadCsvAsInts("11/input.txt")
	paintCount, _ := run(input, false)
	_, grid := run(input, true)

	println(paintCount)
	for _, line := range grid {
		for _, b := range line {
			if b {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
