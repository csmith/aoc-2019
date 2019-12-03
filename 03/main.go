package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"math"
	"strings"
)

var directions = map[byte]common.Point{
	'L': {X: -1, Y: 0},
	'R': {X: +1, Y: 0},
	'U': {X: 0, Y: -1},
	'D': {X: 0, Y: +1},
}

// buildMap constructs a map of points that one or more wires hit, to a slice
// of ints representing the number of steps each wire takes to reach that point.
// The number of steps is 0 if the point is the origin or if the wire doesn't
// reach that point.
func buildMap(origin common.Point, wires []string) map[common.Point][]int64 {
	points := map[common.Point][]int64{}

	for n, wire := range wires {
		var (
			pos         = origin
			steps int64 = 0
			moves       = strings.Split(wire, ",")
		)

		for _, move := range moves {
			dir := directions[move[0]]
			length := common.MustAtoi(move[1:])

			for i := 0; i < length; i++ {
				pos = pos.Plus(dir)
				steps++

				val, ok := points[pos]

				if !ok {
					points[pos] = make([]int64, len(wires))
					val = points[pos]
				}
				val[n] = steps
			}
		}
	}

	return points
}

// combinedSteps computes the total number of steps each wire had to take, given a slice of measurements. If not all
// wires have measurement, the second return parameter will be false.
func combinedSteps(steps []int64) (int64, bool) {
	var res int64 = 0
	for _, distance := range steps {
		if distance == 0 {
			return 0, false
		} else {
			res += distance
		}
	}
	return res, true
}

func main() {
	wires := common.ReadFileAsStrings("03/input.txt")
	origin := common.Point{X: 0, Y: 0}
	points := buildMap(origin, wires)

	var (
		bestDistance int64 = math.MaxInt64
		bestSteps    int64 = math.MaxInt64
	)

	for pos, v := range points {
		if combinedSteps, hitAll := combinedSteps(v); hitAll {
			if distance := pos.Manhattan(origin); distance < bestDistance {
				bestDistance = distance
			}

			if combinedSteps < bestSteps {
				bestSteps = combinedSteps
			}
		}
	}

	fmt.Println(bestDistance)
	fmt.Println(bestSteps)
}
