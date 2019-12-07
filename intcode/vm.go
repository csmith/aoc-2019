package intcode

// VirtualMachine is an IntCode virtual machine.
type VirtualMachine struct {
	ip     int
	modes  uint8
	Memory []int
	Halted bool
	Input  chan int
	Output chan int
}

// NewVirtualMachine creates a new IntCode virtual machine, initialised
// to the given slice of memory.
func NewVirtualMachine(memory []int, hasIO bool) *VirtualMachine {
	vm := &VirtualMachine{
		ip:     0,
		Memory: memory,
		Halted: false,
	}

	if hasIO {
		vm.Input = make(chan int, 1)
		vm.Output = make(chan int, 1)
	}

	return vm
}

func (vm *VirtualMachine) arg(pos int) int {
	mask := uint8(1) << uint8(pos)
	if vm.modes&mask == mask {
		return vm.Memory[vm.ip+1+pos]
	} else {
		return vm.Memory[vm.Memory[vm.ip+1+pos]]
	}
}

// Run repeatedly executes instructions until the VM halts.
func (vm *VirtualMachine) Run() {
	for !vm.Halted {
		instruction := vm.Memory[vm.ip]
		opcode := instruction % 100

		vm.modes = 0
		mask := uint8(1)
		for i := instruction / 100; i > 0; i /= 10 {
			if i%10 == 1 {
				vm.modes = vm.modes | mask
			}
			mask = mask << 1
		}

		opcodes[opcode](vm)
	}
	if vm.Output != nil {
		close(vm.Output)
	}
}

// Reset resets the memory to the given slice, and all other state back to its original value.
func (vm *VirtualMachine) Reset(memory []int) {
	copy(vm.Memory, memory)
	vm.ip = 0
	vm.Halted = false
	if vm.Input != nil {
		vm.Input = make(chan int, 1)
		vm.Output = make(chan int, 1)
	}
}
