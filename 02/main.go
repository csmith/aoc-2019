package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"sync"
)

func run(pool sync.Pool, input []int, noun, verb int) int {
	state := pool.Get().([]int)
	defer pool.Put(state)
	copy(state, input)
	state[1] = noun
	state[2] = verb

	var ip = 0
	for {
		var instr = state[ip]
		if instr == 1 {
			state[state[ip+3]] = state[state[ip+1]] + state[state[ip+2]]
		} else if instr == 2 {
			state[state[ip+3]] = state[state[ip+1]] * state[state[ip+2]]
		} else if instr == 99 {
			return state[0]
		}
		ip += 4
	}
}

func findOutput(pool sync.Pool, input []int, target int) int {
	res := make(chan int)

	for n := 0; n < 100; n++ {
		go func(n int) {
			for v := 0; v < 100; v++ {
				if run(pool, input, n, v) == target {
					res <- 100*n + v
				}
			}
		}(n)
	}

	return <-res
}

func main() {
	input := common.Atoi(common.ReadCommas("02/input.txt"))

	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, len(input))
		},
	}

	fmt.Println(run(pool, input, 12, 2))
	fmt.Println(findOutput(pool, input, 19690720))
}
