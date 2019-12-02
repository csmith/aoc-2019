package intcode

// VirtualMachine is an IntCode virtual machine.
type VirtualMachine struct {
	ip      int
	opcodes map[int]OpcodeFunc
	Memory  []int
	Halted  bool
}

// NewVirtualMachine creates a new IntCode virtual machine, initialised
// to the given slice of memory.
func NewVirtualMachine(memory []int) *VirtualMachine {
	return &VirtualMachine{
		ip:     0,
		Memory: memory,
		Halted: false,
		opcodes: map[int]OpcodeFunc{
			1:  AddOpcode,
			2:  MulOpcode,
			99: HaltOpcode,
		},
	}
}

// Run repeatedly executes instructions until the VM halts.
func (vm *VirtualMachine) Run() {
	for !vm.Halted {
		vm.opcodes[vm.Memory[vm.ip]](vm, vm.Memory[vm.ip+1:])
	}
}

// Reset resets the memory to the given slice, and all other state back to its original value.
func (vm *VirtualMachine) Reset(memory []int) {
	copy(vm.Memory, memory)
	vm.ip = 0
	vm.Halted = false
}
