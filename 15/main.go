package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
)

type tile uint8
type direction int

const (
	unknown tile = iota
	wall    tile = iota
	open    tile = iota
	start   tile = iota
	oxygen  tile = iota
	filled  tile = iota

	north direction = 1
	south direction = 2
	west  direction = 3
	east  direction = 4
)

var directions = [4]direction{north, south, west, east}

type step struct {
	vm    *intcode.VirtualMachine
	x     int
	y     int
	steps int
}

func (s step) next(direction direction) (int, int) {
	switch direction {
	case north:
		return s.x, s.y - 1
	case south:
		return s.x, s.y + 1
	case west:
		return s.x - 1, s.y
	case east:
		return s.x + 1, s.y
	}
	panic(fmt.Sprintf("unknown direction: %v", direction))
}

func explore(vm *intcode.VirtualMachine) (grid [100][100]tile, stepsToOxygen int, oxygenX int, oxygenY int) {
	queue := make(chan step, 100)

	queue <- step{
		vm: vm,
		x:  50,
		y:  50,
	}
	grid[50][50] = start

	for {
		select {
		case s := <-queue:
			for _, d := range directions {
				xp, yp := s.next(d)

				if grid[yp][xp] != unknown {
					continue
				}

				vmp := s.vm.Clone()
				state := *vmp.RunForInput(int(d))

				if state == 0 {
					grid[yp][xp] = wall
				} else {
					if state == 2 {
						grid[yp][xp] = oxygen
						stepsToOxygen = s.steps + 1
						oxygenX = xp
						oxygenY = yp
					} else {
						grid[yp][xp] = open
					}

					queue <- step{
						vm:    vmp,
						x:     xp,
						y:     yp,
						steps: s.steps + 1,
					}
				}
			}
		default:
			return
		}
	}
}

func disperse(grid [100][100]tile, x, y int) (maxSteps int) {
	queue := make(chan step, 100)

	queue <- step{
		x: x,
		y: y,
	}

	for {
		select {
		case s := <-queue:
			maxSteps = common.Max(maxSteps, s.steps)
			for _, d := range directions {
				xp, yp := s.next(d)

				if grid[yp][xp] == filled || grid[yp][xp] == wall {
					continue
				}

				grid[yp][xp] = filled

				queue <- step{
					x:     xp,
					y:     yp,
					steps: s.steps + 1,
				}
			}
		default:
			return
		}
	}
}

func main() {
	input := common.ReadCsvAsInts("15/input.txt")
	memory := make([]int, len(input))
	copy(memory, input)
	vm := intcode.NewVirtualMachine(memory)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	go vm.Run()

	grid, stepsToOxygen, oxyX, oxyY := explore(vm)

	println(stepsToOxygen)
	println(disperse(grid, oxyX, oxyY))
}
