package main

import (
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
)

func last(channel <-chan int) (res int) {
	for {
		o, more := <-channel
		if more {
			res = o
		} else {
			return
		}
	}
}

func main() {
	input := common.ReadCsvAsInts("09/input.txt")
	memory := make([]int, len(input))
	copy(memory, input)

	vm := intcode.NewVirtualMachine(memory)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	vm.Input <- 1
	go vm.Run()
	println(last(vm.Output))

	vm.Reset(input)
	vm.Input = make(chan int, 1)
	vm.Output = make(chan int, 1)
	vm.Input <- 2
	go vm.Run()
	println(last(vm.Output))
}
