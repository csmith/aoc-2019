package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"sync"
)

func readInput(file string) (pos [3][]int64, vel [3][]int64) {
	for _, line := range common.ReadFileAsStrings(file) {
		var x, y, z int64
		if _, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z); err != nil {
			panic(fmt.Sprintf("unable to parse line '%v': %v", line, err))
		}

		pos[0] = append(pos[0], x)
		pos[1] = append(pos[1], y)
		pos[2] = append(pos[2], z)
		vel[0] = append(vel[0], 0)
		vel[1] = append(vel[1], 0)
		vel[2] = append(vel[2], 0)
	}
	return
}

func attract(p1, p2 int64, dp1, dp2 *int64) {
	if p1 < p2 {
		*dp1++
		*dp2--
	} else if p1 > p2 {
		*dp1--
		*dp2++
	}
}

func step(pos, vel []int64) {
	for x := 0; x < len(pos); x++ {
		for y := x + 1; y < len(pos); y++ {
			attract(pos[x], pos[y], &vel[x], &vel[y])
		}
	}

	for i, v := range vel {
		pos[i] += v
	}
}

func static(vel []int64) bool {
	for _, v := range vel {
		if v != 0 {
			return false
		}
	}
	return true
}

func energy(channel chan []int64, moons int) int64 {
	sums := make([]int64, moons*2)
	for i := 0; i < 3; i++ {
		row := <-channel
		for n, v := range row {
			sums[n] += common.Abs(v)
		}
	}

	energy := int64(0)
	for i := 0; i < len(sums)/2; i++ {
		energy += sums[i] * sums[len(sums)/2+i]
	}
	return energy
}

// Sweeping assumptions/notes on how this mess works:
//
// 1) the movement of the moons will be parabolic - they'll accelerate towards each other,
//    then gradually slow down to a complete stop and reverse path going back to the starting point.
//    I'm not sure I can prove that's the case in general, but it does seem to hold true.
//
// 2) all the moons will hit the inflection point of the parabola at the same step. I think this is reasonable
//    as each force is mirrored, so bodies accelerating towards one another will end up with the same total velocity
//    when they cross
//
// 3) the middle of the parabola occurs after step 1000 (otherwise the code would need a fiddly bit of
//    state tracking to ensure it returned a value for part 1).
//
// 4) because the axes are independent, they can be simulated in parallel, and their individual parabolic inflection
//    points found. This gives us the loop count for each axis (2x the number of steps to reach the inflection point),
//    and we can find the first time all three axis's loops intersect by finding the lowest common multiple of those
//    three values.
//
// e.g.
//
// <------------------- position on axis ------------------>
//
//                      (somewhere)               __,..--'""      |
//                       step 1000       _,..--'""         ^      |
//                           v   _,..-'""                start    |
//                        _,..-'"                       vel = 0   |
//                  _,.-'"                                        |
//             _.-'"                                              |
//         _.-"                                                   |
//      .-'                                                       |
//    .'                                                          |
//   /                                                            |
//  ;  } inflection point                                       steps
//  ;  } vel = 0, step = n                                        |
//   \                                                            |
//    `.                                                          |
//      `-.                                                       |
//         "-._                                                   |
//             "`-._                                              |
//                  "`-.,_                                        |
//                        "`-..,_                                 |
//                               ""`-..,_                         |
//                                       ""`--..,_                |
//                                                ""`--..,__      V
//                                                         ^
//                                               back to original position
//                                                  vel = 0, step = 2n
func main() {
	pos, vel := readInput("12/input.txt")

	part1wg, part1chan := &sync.WaitGroup{}, make(chan []int64, 3)
	part2wg, part2chan := &sync.WaitGroup{}, make(chan int64, 3)
	for i, ps := range pos {
		part1wg.Add(1)
		part2wg.Add(1)

		go func(pos, vel []int64) {
			for n := int64(0); ; n++ {
				step(pos, vel)

				if n+1 == 1000 {
					part1chan <- append(pos, vel...)
					part1wg.Done()
				}

				if static(vel) {
					part2chan <- 2 * (n + 1)
					part2wg.Done()
					return
				}
			}
		}(ps, vel[i])
	}

	part1wg.Wait()
	println(energy(part1chan, len(pos[0])))

	part2wg.Wait()
	println(common.LCM(<-part2chan, <-part2chan, <-part2chan))
}
