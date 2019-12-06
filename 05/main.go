package main

import (
	"fmt"
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
	input := common.ReadCsvAsInts("05/input.txt")
	memory := make([]int, len(input))
	copy(memory, input)

	vm := intcode.NewVirtualMachine(memory, true)
	vm.Input <- 1
	go vm.Run()
	fmt.Println(last(vm.Output))

	vm.Reset(input)
	vm.Input <- 5
	go vm.Run()
	fmt.Println(last(vm.Output))
}
