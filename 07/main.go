package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
)

func copyOf(memory []int) []int {
	res := make([]int, len(memory))
	copy(res, memory)
	return res
}

func runPipeline(program []int, ps []int, feedback bool) int {
	// Create a series of VMs for our amplifiers
	vms := make([]*intcode.VirtualMachine, len(ps))
	for i := 0; i < len(ps); i++ {
		vms[i] = intcode.NewVirtualMachine(copyOf(program), true)
	}

	// Link all the inputs and outputs
	for i, vm := range vms {
		if i > 0 {
			vm.Input = vms[i-1].Output
		} else if feedback {
			vm.Input = vms[len(vms)-1].Output
		}
	}

	// Seed the phase settings
	for i, vm := range vms {
		vm.Input <- ps[i]
	}

	// Background run all but the last VM
	for _, vm := range vms[:len(vms)-1] {
		go vm.Run()
	}

	// Kick off the first VM and then wait for the last VM to finish
	vms[0].Input <- 0
	vms[len(vms)-1].Run()

	// Snag the left over value from the last VM's output
	return <-vms[len(vms)-1].Output
}

func maxOutput(input []int, ps []int, feedback bool) int {
	max := 0
	for _, p := range common.Permutations(ps) {
		val := runPipeline(input, p, feedback)
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	input := common.ReadCsvAsInts("07/input.txt")
	fmt.Println(maxOutput(input, []int{0, 1, 2, 3, 4}, false))
	fmt.Println(maxOutput(input, []int{5, 6, 7, 8, 9}, true))
}
