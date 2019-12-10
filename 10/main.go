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

func angleBetween(asteroid1, asteroid2 *asteroid) float64 {
	if asteroid1.y == asteroid2.y {
		if asteroid2.x > asteroid1.x {
			return math.Pi / 2
		} else {
			return 3 * math.Pi / 2
		}
	} else {
		angle := math.Atan(float64(asteroid2.x-asteroid1.x) / float64(asteroid1.y-asteroid2.y))
		if asteroid1.y < asteroid2.y {
			angle += math.Pi
		}
		if angle < 0 {
			angle += math.Pi * 2
		}
		return angle
	}
}

func checkAngles(asteroid1 *asteroid, others []*asteroid, countVisible bool) map[float64][]*asteroid {
	angles := make(map[float64][]*asteroid)
	for _, asteroid2 := range others {
		if asteroid2 == asteroid1 {
			continue
		}

		angle := angleBetween(asteroid1, asteroid2)

		if len(angles[angle]) == 0 && countVisible {
			asteroid1.visible++
			asteroid2.visible++
		}

		angles[angle] = append(angles[angle], asteroid2)
	}
	return angles
}

func main() {
	var (
		input     = common.ReadFileAsStrings("10/input.txt")
		asteroids = buildMap(input)
		best      *asteroid
	)

	for i, asteroid1 := range asteroids {
		checkAngles(asteroid1, asteroids[i+1:], true)
		if best == nil || asteroid1.visible > best.visible {
			best = asteroid1
		}
	}

	if best == nil {
		panic("No asteroids found?")
	}

	println(best.visible)

	targets := checkAngles(best, asteroids, false)
	angles := make([]float64, 0, len(targets))
	for k := range targets {
		angles = append(angles, k)
	}
	sort.Float64s(angles)

	var destroyed *asteroid
	var i = 0
	for n := 0; n < 200; n++ {
		if len(targets[angles[i]]) == 1 {
			// There's a single target at this angle, skip the angle in the future
			destroyed = targets[angles[i]][0]
			angles = append(angles[:i], angles[i+1:]...)
		} else {
			// Multiple targets exists at this angle, remove the closest and move on to the next angle
			var bestDistance = math.MaxFloat64
			var bestTarget = 0
			for j, target := range targets[angles[i]] {
				distance := math.Abs(float64(target.x-best.x)) + math.Abs(float64(target.y-best.y))
				if distance < bestDistance {
					bestDistance = distance
					bestTarget = j
				}
			}

			destroyed = targets[angles[i]][bestTarget]
			targets[angles[i]] = append(targets[angles[i]][:bestTarget], targets[angles[i]][bestTarget+1:]...)
			i++
		}

		i = i % len(angles)
	}

	if destroyed == nil {
		panic("Universe doesn't make sense. Reboot and try again?")
	}

	println(destroyed.x*100 + destroyed.y)
}
