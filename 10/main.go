package main

import (
	"github.com/csmith/aoc-2019/common"
	"math"
	"sort"
)

type asteroid struct {
	x, y, visible int
}

func buildMap(input []string) []*asteroid {
	var res []*asteroid
	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				res = append(res, &asteroid{x: x, y: y})
			}
		}
	}
	return res
}

func checkAngles(asteroid1 *asteroid, others []*asteroid, countVisible bool) map[float64][]*asteroid {
	angles := make(map[float64][]*asteroid, len(others))
	for _, asteroid2 := range others {
		if asteroid2 == asteroid1 {
			continue
		}

		angle := math.Atan2(float64(asteroid2.x-asteroid1.x), float64(asteroid1.y-asteroid2.y))
		if angle < 0 {
			angle += math.Pi * 2
		}

		if countVisible && len(angles[angle]) == 0 {
			asteroid1.visible++
			asteroid2.visible++
		}

		angles[angle] = append(angles[angle], asteroid2)
	}
	return angles
}

func selectBest(asteroids []*asteroid) (best *asteroid) {
	for i, asteroid1 := range asteroids {
		checkAngles(asteroid1, asteroids[i+1:], true)
		if best == nil || asteroid1.visible > best.visible {
			best = asteroid1
		}
	}
	return
}

func closest(origin *asteroid, asteroids []*asteroid) (int, *asteroid) {
	var bestDistance = math.MaxFloat64
	var bestIndex = 0
	for j, target := range asteroids {
		distance := math.Abs(float64(target.x-origin.x)) + math.Abs(float64(target.y-origin.y))
		if distance < bestDistance {
			bestDistance = distance
			bestIndex = j
		}
	}
	return bestIndex, asteroids[bestIndex]
}

func sortedAngles(targets map[float64][]*asteroid) []float64 {
	angles := make([]float64, 0, len(targets))
	for k := range targets {
		angles = append(angles, k)
	}
	sort.Float64s(angles)
	return angles
}

func main() {
	var (
		input     = common.ReadFileAsStrings("10/input.txt")
		asteroids = buildMap(input)
		best      = selectBest(asteroids)
		targets   = checkAngles(best, asteroids, false)
		angles    = sortedAngles(targets)
	)

	var destroyed *asteroid
	if len(angles) >= 200 {
		// We don't complete a full loop, so the answer is just the nearest asteroid at the 200th angle
		_, destroyed = closest(best, targets[angles[199]])
	} else {
		// We loop once, actually simulate the whole thing.
		var i = 0
		for n := 0; n < 200; n++ {
			if len(targets[angles[i]]) == 1 {
				// There's a single target at this angle, skip the angle in the future
				destroyed = targets[angles[i]][0]
				angles = append(angles[:i], angles[i+1:]...)
			} else {
				// Multiple targets exists at this angle, remove the closest and move on to the next angle
				var nearestIndex int
				nearestIndex, destroyed = closest(best, targets[angles[i]])
				targets[angles[i]] = append(targets[angles[i]][:nearestIndex], targets[angles[i]][nearestIndex+1:]...)
				i++
			}

			i = i % len(angles)
		}
	}

	if destroyed == nil {
		panic("Universe doesn't make sense. Reboot and try again?")
	}

	println(best.visible)
	println(destroyed.x*100 + destroyed.y)
}
