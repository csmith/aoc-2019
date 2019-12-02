package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
	"sync"
)

func run(pool *sync.Pool, input []int, noun, verb int) int {
	vm := pool.Get().(*intcode.VirtualMachine)
	defer pool.Put(vm)

	vm.Reset(input)
	vm.Memory[1] = noun
	vm.Memory[2] = verb
	vm.Run()
	return vm.Memory[0]
}

func findOutput(pool *sync.Pool, input []int, target int) int {
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

	pool := &sync.Pool{
		New: func() interface{} {
			return intcode.NewVirtualMachine(make([]int, len(input)))
		},
	}

	fmt.Println(run(pool, input, 12, 2))
	fmt.Println(findOutput(pool, input, 19690720))
}
