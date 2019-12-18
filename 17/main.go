package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
	"strconv"
)

const (
	up     = '^'
	down   = 'v'
	left   = '<'
	right  = '>'
	borked = 'X'
)

func rotate(from, to rune) rune {
	if from == up && to == left || from == left && to == down || from == down && to == right || from == right && to == up {
		return 'L'
	} else {
		return 'R'
	}
}

func next(picture [][]rune, x, y int, direction rune) (nextDirection, turn rune) {
	if x > 0 && direction != right && picture[y][x-1] == '#' {
		nextDirection = left
	} else if x < len(picture[0])-1 && direction != left && picture[y][x+1] == '#' {
		nextDirection = right
	} else if y > 0 && direction != down && picture[y-1][x] == '#' {
		nextDirection = up
	} else if y < len(picture)-1 && direction != up && picture[y+1][x] == '#' {
		nextDirection = down
	} else {
		return borked, borked
	}

	return nextDirection, rotate(direction, nextDirection)
}

func follow(picture [][]rune, x, y int, direction rune) (int, int, int) {
	deltaX, deltaY := 0, 0
	switch direction {
	case up:
		deltaY = -1
	case down:
		deltaY = 1
	case left:
		deltaX = -1
	case right:
		deltaX = +1
	case borked:
		return x, y, borked
	}

	length := 0
	for x+deltaX >= 0 && x+deltaX <= len(picture[0])-1 && y+deltaY >= 0 && y+deltaY <= len(picture)-1 && picture[y+deltaY][x+deltaX] == '#' {
		x += deltaX
		y += deltaY
		length++
	}

	return x, y, length
}

func readPicture(memory []int) [][]rune {
	vm := intcode.NewVirtualMachine(memory)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	go vm.Run()

	var picture [][]rune
	var currentRow []rune
	for {
		char, more := <-vm.Output
		if !more {
			break
		}

		if char == '\n' && len(currentRow) > 0 {
			picture = append(picture, currentRow)
			currentRow = make([]rune, 0)
		} else {
			currentRow = append(currentRow, rune(char))
		}
	}
	return picture
}

func analysePicture(picture [][]rune) (sum int, robot rune, robotX, robotY int) {
	robot = borked
	for y, line := range picture {
		for x, r := range line {
			if r == '#' &&
				x > 0 && line[x-1] == '#' &&
				x < len(line)-1 && line[x+1] == '#' &&
				y > 0 && picture[y-1][x] == '#' &&
				y < len(picture)-1 && picture[y+1][x] == '#' {
				sum += x * y
			}

			if r == up || r == down || r == left || r == right {
				robotX = x
				robotY = y
				robot = r
			}
		}
	}
	return
}

func buildRoute(picture [][]rune, robot rune, robotX, robotY int) []string {
	var (
		length int
		turn   rune
		res    []string
	)
	for {
		robot, turn = next(picture, robotX, robotY, robot)
		robotX, robotY, length = follow(picture, robotX, robotY, robot)

		if robot == borked {
			break
		}

		res = append(res, string(turn))
		res = append(res, strconv.Itoa(length))
	}
	return res
}

func calculateDust(input []int, m, a, b, c string) int {
	vm := intcode.NewVirtualMachine(input)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	vm.Memory[0] = 2
	go vm.Run()

	go func() {
		for _, line := range []string{m, a, b, c, "n"} {
			for _, r := range line {
				vm.Input <- int(r)
			}
			vm.Input <- '\n'
		}
	}()

	return common.Last(vm.Output)
}

func main() {
	input := common.ReadCsvAsInts("17/input.txt")
	memory := make([]int, len(input))
	copy(memory, input)

	picture := readPicture(input)
	sum, robot, robotX, robotY := analysePicture(picture)
	route := buildRoute(picture, robot, robotX, robotY)
	m, a, b, c := compress(route)
	dust := calculateDust(memory, m, a, b, c)

	fmt.Println(sum)
	fmt.Println(dust)
}
