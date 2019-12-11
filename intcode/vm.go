package intcode

// VirtualMachine is an IntCode virtual machine.
type VirtualMachine struct {
	ip             int
	rb             int
	parameterModes uint8
	relativeModes  uint8
	Memory         []int
	Halted         bool
	Input          chan int
	Output         chan int
}

// NewVirtualMachine creates a new IntCode virtual machine, initialised to the given slice of memory.
// The caller is responsible for initialising the VM's I/O channels if required.
func NewVirtualMachine(memory []int) *VirtualMachine {
	vm := &VirtualMachine{
		ip:     0,
		Memory: memory,
		Halted: false,
	}

	return vm
}

// arg Returns the value of the given argument for the current instruction.
func (vm *VirtualMachine) arg(pos int) *int {
	mask := uint8(1) << uint8(pos)
	if vm.parameterModes&mask == mask {
		// Parameter mode - the value of the argument is just treated as an int
		for vm.ip+1+pos > len(vm.Memory) {
			vm.Memory = append(vm.Memory, make([]int, 1024)...)
		}

		return &vm.Memory[vm.ip+1+pos]
	} else if vm.relativeModes&mask == mask {
		// Relative mode - the value of the argument is treated as a memory offset from the relative base
		for vm.ip+1+pos > len(vm.Memory) || vm.rb+vm.Memory[vm.ip+1+pos] > len(vm.Memory) {
			vm.Memory = append(vm.Memory, make([]int, 1024)...)
		}

		return &vm.Memory[vm.rb+vm.Memory[vm.ip+1+pos]]
	} else {
		// Position mode - the value of the argument is treated as a memory offset from the start of the memory
		for vm.ip+1+pos > len(vm.Memory) || vm.Memory[vm.ip+1+pos] > len(vm.Memory) {
			vm.Memory = append(vm.Memory, make([]int, 1024)...)
		}

		return &vm.Memory[vm.Memory[vm.ip+1+pos]]
	}
}

// Run repeatedly executes instructions until the VM halts.
func (vm *VirtualMachine) Run() {
	for !vm.Halted {
		instruction := vm.Memory[vm.ip]
		opcode := instruction % 100

		vm.parameterModes = 0
		vm.relativeModes = 0
		mask := uint8(1)
		for i := instruction / 100; i > 0; i /= 10 {
			mode := i % 10
			if mode == 1 {
				vm.parameterModes = vm.parameterModes | mask
			} else if mode == 2 {
				vm.relativeModes = vm.relativeModes | mask
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

	// We may previously have expanded our own memory, reset that to zero.
	for i := len(memory); i < len(vm.Memory)-1; i++ {
		vm.Memory[i] = 0
	}

	vm.ip = 0
	vm.rb = 0
	vm.Halted = false
}
