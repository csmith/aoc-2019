package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"math"
	"strings"
)

type line struct {
	from  int64
	to    int64
	perp  int64
	steps int64
}

// readWire reads the instructions for a wire and populates the given slices with the horizontal and vertical lines
// that make up the wire's path.
func readWire(wire string, horizontal *[]line, vertical *[]line) {
	moves := strings.Split(wire, ",")
	x := int64(0)
	y := int64(0)
	steps := int64(0)

	for _, move := range moves {
		length := int64(common.MustAtoi(move[1:]))
		switch move[0] {
		case 'U':
			*vertical = append(*vertical, line{from: y, to: y - length, perp: x, steps: steps})
			y -= length
		case 'D':
			*vertical = append(*vertical, line{from: y, to: y + length, perp: x, steps: steps})
			y += length
		case 'L':
			*horizontal = append(*horizontal, line{from: x, to: x - length, perp: y, steps: steps})
			x -= length
		case 'R':
			*horizontal = append(*horizontal, line{from: x, to: x + length, perp: y, steps: steps})
			x += length
		}
		steps += length
	}
}

// checkCrosses checks if any of the given sets of perpendicular lines cross, and returns the
// smallest Manhattan distance to the origin, and the smallest number of combined steps, of
// those crosses.
func checkCrosses(horizontal *[]line, vertical *[]line) (int64, int64) {
	var (
		bestDistance int64 = math.MaxInt64
		bestSteps    int64 = math.MaxInt64
	)

	for _, h := range *horizontal {
		for _, v := range *vertical {
			var steps int64 = math.MaxInt64
			if h.from <= v.perp && v.perp <= h.to {
				// If the horizontal line goes left-to-right
				if v.from <= h.perp && h.perp <= v.to {
					// If the vertical line goes top-to-bottom
					steps = (h.steps + v.perp - h.from) + (v.steps + h.perp - v.from)
				} else if v.to <= h.perp && h.perp <= v.from {
					// If the vertical line goes bottom-to-top
					steps = (h.steps + v.perp - h.from) + (v.steps + v.from - h.perp)
				} else {
					continue
				}
			} else if h.to <= v.perp && v.perp <= h.from {
				// If the horizontal line goes right-to-left
				if v.from <= h.perp && h.perp <= v.to {
					// If the vertical line goes top-to-bottom
					steps = (h.steps + h.from - v.perp) + (v.steps + h.perp - v.from)
				} else if v.to <= h.perp && h.perp <= v.from {
					// If the vertical line goes bottom-to-top
					steps = (h.steps + h.from - v.perp) + (v.steps + v.from - h.perp)
				} else {
					continue
				}
			} else {
				continue
			}

			distance := common.Abs(v.perp) + common.Abs(h.perp)
			if distance < bestDistance {
				bestDistance = distance
			}
			if steps < bestSteps {
				bestSteps = steps
			}
		}
	}

	return bestDistance, bestSteps
}

func traceWires() (int64, int64) {
	var (
		horiz1 []line
		horiz2 []line
		vert1  []line
		vert2  []line
	)

	wires := common.ReadFileAsStrings("03/input.txt")
	readWire(wires[0], &horiz1, &vert1)
	readWire(wires[1], &horiz2, &vert2)

	d1, s1 := checkCrosses(&horiz1, &vert2)
	d2, s2 := checkCrosses(&horiz2, &vert1)
	return min(d1, d2), min(s1, s2)
}

func main() {
	part1, part2 := traceWires()
	fmt.Println(part1)
	fmt.Println(part2)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}
