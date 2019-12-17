package main

import (
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
	"testing"
)

func Benchmark_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := common.ReadCsvAsInts("input.txt")
		memory := make([]int, len(input))
		copy(memory, input)

		vm := intcode.NewVirtualMachine(memory)
		vm.Input = make(chan int, 1)
		vm.Output = make(chan int, 1)
		vm.Input <- 1
		go vm.Run()
		_ = common.Last(vm.Output)
	}
}

func Benchmark_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := common.ReadCsvAsInts("input.txt")
		memory := make([]int, len(input))
		copy(memory, input)

		vm := intcode.NewVirtualMachine(memory)
		vm.Input = make(chan int, 1)
		vm.Output = make(chan int, 1)
		vm.Input <- 2
		go vm.Run()
		_ = common.Last(vm.Output)
	}
}
