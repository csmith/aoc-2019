package intcode

// OpcodeFunc is a function that describes an opcode implemented in the VM.
type OpcodeFunc = func(vm *VirtualMachine, args []int)

// AddOpcode takes the values from the memory addresses given by args 1 and 2, adds them together,
// and stores at the memory address given by arg 3.
func AddOpcode(vm *VirtualMachine, args []int) {
	vm.Memory[args[2]] = vm.Memory[args[0]] + vm.Memory[args[1]]
	vm.ip += 4
}

// MulOpcode takes the values from the memory addresses given by args 1 and 2, muliplies them together,
// and stores at the memory address given by arg 3.
func MulOpcode(vm *VirtualMachine, args []int) {
	vm.Memory[args[2]] = vm.Memory[args[0]] * vm.Memory[args[1]]
	vm.ip += 4
}

// HaltOpcode halts the VM and takes no arguments.
func HaltOpcode(vm *VirtualMachine, args []int) {
	vm.Halted = true
}
