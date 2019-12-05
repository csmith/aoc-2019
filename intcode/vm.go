package intcode

import (
	"fmt"
	"math"
)

// VirtualMachine is an IntCode virtual machine.
type VirtualMachine struct {
	ip      int
	opcodes map[int]OpcodeFunc
	Memory  []int
	Halted  bool
	Input   chan int
	Output  chan int
}

// NewVirtualMachine creates a new IntCode virtual machine, initialised
// to the given slice of memory.
func NewVirtualMachine(memory []int) *VirtualMachine {
	return &VirtualMachine{
		ip:     0,
		Memory: memory,
		Halted: false,
		Input:  make(chan int, 100),
		Output: make(chan int, 100),
		opcodes: map[int]OpcodeFunc{
			1:  AddOpcode,
			2:  MulOpcode,
			3:  ReadOpCode,
			4:  WriteOpCode,
			5:  JumpIfTrueOpCode,
			6:  JumpIfFalseOpCode,
			7:  LessThanOpCode,
			8:  EqualsOpCode,
			99: HaltOpcode,
		},
	}
}

// Run repeatedly executes instructions until the VM halts.
func (vm *VirtualMachine) Run() {
	for !vm.Halted {
		instruction := vm.Memory[vm.ip]
		opcode := instruction % 100

		vm.opcodes[opcode](vm, func(pos int) int {
			mode := (instruction / int(math.Pow10(2+pos))) % 10
			switch mode {
			case 0:
				return vm.Memory[vm.Memory[vm.ip+1+pos]]
			case 1:
				return vm.Memory[vm.ip+1+pos]
			default:
				panic(fmt.Sprintf("Unknown parameter mode: %d", mode))
			}
		})
	}
	close(vm.Input)
	close(vm.Output)
}

// Reset resets the memory to the given slice, and all other state back to its original value.
func (vm *VirtualMachine) Reset(memory []int) {
	copy(vm.Memory, memory)
	vm.ip = 0
	vm.Halted = false
	vm.Input = make(chan int, 100)
	vm.Output = make(chan int, 100)
}
