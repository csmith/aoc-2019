package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
	"math"
)

type point struct {
	x int
	y int
}

var dirs = [4]point{
	0: {y: -1},
	1: {x: +1},
	2: {y: +1},
	3: {x: -1},
}

func run(memory []int, startWhite bool) (int, map[point]bool) {
	input := make([]int, len(memory))
	copy(input, memory)
	vm := intcode.NewVirtualMachine(input)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	go vm.Run()

	grid := make(map[point]bool, 50)
	paintCount := 0

	pos := point{x: 0, y: 0}
	dir := 0
	colour := startWhite

	for {
		if colour {
			vm.Input <- 1
		} else {
			vm.Input <- 0
		}

		paint, more := <-vm.Output
		if !more {
			break
		}
		turn := <-vm.Output

		if _, ok := grid[pos]; !ok {
			paintCount++
		}
		grid[pos] = paint == 1

		dir = (4 + dir + 2*turn - 1) % 4
		pos.x += dirs[dir].x
		pos.y += dirs[dir].y
		colour = grid[pos]
	}

	return paintCount, grid
}

func main() {
	input := common.ReadCsvAsInts("11/input.txt")
	paintCount, _ := run(input, false)
	_, grid := run(input, true)

	fmt.Println(paintCount)

	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	for pos := range grid {
		minX, maxX = common.Min(minX, pos.x), common.Max(maxX, pos.x)
		minY, maxY = common.Min(minY, pos.y), common.Max(maxY, pos.y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[point{x: x, y: y}] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
