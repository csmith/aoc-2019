package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
)

func countOrbits(satellites map[string][]string, start string, depth int) (res int) {
	for _, satellite := range satellites[start] {
		res += depth + 1 + countOrbits(satellites, satellite, depth+1)
	}
	return
}

func routeToCenter(orbits map[string]string, start string) (res []string) {
	next := start
	for next != "COM" {
		next = orbits[next]
		res = append(res, next)
	}
	return
}

func countToCenter(orbits map[string]string, start string) (res map[string]int) {
	res = make(map[string]int)

	steps := 0
	next := start
	for next != "COM" {
		next = orbits[next]
		res[next] = steps
		steps++
	}
	return
}

func shortestPath(orbits map[string]string, from, to string) int {
	route := routeToCenter(orbits, from)
	steps := countToCenter(orbits, to)

	for i, body := range route {
		if count, ok := steps[body]; ok {
			return count + i
		}
	}
	panic(fmt.Sprintf("No path found between %s and %s.", from, to))
}

func main() {
	lines := common.ReadFileAsStrings("06/input.txt")
	satellites := make(map[string][]string, len(lines))
	orbits := make(map[string]string, len(lines))
	for _, line := range lines {
		around := line[0:3]
		body := line[4:7]
		satellites[around] = append(satellites[around], body)
		orbits[body] = around
	}

	println(countOrbits(satellites, "COM", 0))
	println(shortestPath(orbits, "YOU", "SAN"))
}
