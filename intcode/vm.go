package intcode

type parameterMode int

const (
	position  parameterMode = 0
	immediate parameterMode = 1
	relative  parameterMode = 2
)

// VirtualMachine is an IntCode virtual machine.
type VirtualMachine struct {
	ip     int
	rb     int
	Memory []int
	Halted bool
	Input  chan int
	Output chan int
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

// Clone returns a deep copy of this VM, with newly allocated memory and I/O channels.
func (vm *VirtualMachine) Clone() *VirtualMachine {
	memory := make([]int, len(vm.Memory))
	copy(memory, vm.Memory)

	return &VirtualMachine{
		ip:     vm.ip,
		rb:     vm.rb,
		Memory: memory,
		Halted: vm.Halted,
		Input:  make(chan int, 1),
		Output: make(chan int, 1),
	}
}

// Run repeatedly executes instructions until the VM halts.
func (vm *VirtualMachine) Run() {
	var args [3]*int
	for !vm.Halted {
		instruction := vm.Memory[vm.ip]
		opcode := instruction % 100

		param := instruction / 100
		for i := 0; i < opcodeArity[opcode]; i++ {

			switch parameterMode(param % 10) {

			// The argument is the actual value
			case immediate:
				args[i] = &vm.Memory[vm.ip+1+i]

			// The argument is a memory reference
			case position:
				for vm.Memory[vm.ip+1+i] >= len(vm.Memory) {
					vm.Memory = append(vm.Memory, make([]int, 128)...)
				}
				args[i] = &vm.Memory[vm.Memory[vm.ip+1+i]]

			// The argument is a memory reference offset by the relative base
			case relative:
				for vm.rb+vm.Memory[vm.ip+1+i] >= len(vm.Memory) {
					vm.Memory = append(vm.Memory, make([]int, 128)...)
				}
				args[i] = &vm.Memory[vm.rb+vm.Memory[vm.ip+1+i]]

			}

			param /= 10
		}

		opcodes[opcode](vm, args[0], args[1], args[2])
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
