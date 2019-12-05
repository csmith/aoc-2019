package intcode

// ArgFunc provides the value of an argument for an opcode
type ArgFunc = func(pos int) int

// OpcodeFunc is a function that describes an opcode implemented in the VM.
type OpcodeFunc = func(vm *VirtualMachine, param ArgFunc)

// AddOpcode takes the values specified by args 1 and 2, adds them together, and stores at the memory address given
// by arg 3.
func AddOpcode(vm *VirtualMachine, arg ArgFunc) {
	vm.Memory[vm.Memory[vm.ip+3]] = arg(0) + arg(1)
	vm.ip += 4
}

// MulOpcode takes the values specified by args 1 and 2, multiplies them together, and stores at the memory address
// given by arg 3.
func MulOpcode(vm *VirtualMachine, arg ArgFunc) {
	vm.Memory[vm.Memory[vm.ip+3]] = arg(0) * arg(1)
	vm.ip += 4
}

// ReadOpCode reads a value from the input stream and stores it at the memory address given by arg 1.
func ReadOpCode(vm *VirtualMachine, arg ArgFunc) {
	vm.Memory[vm.Memory[vm.ip+1]] = <-vm.Input
	vm.ip += 2
}

// WriteOpCode writes the value specified by the first argument to the output stream.
func WriteOpCode(vm *VirtualMachine, arg ArgFunc) {
	vm.Output <- arg(0)
	vm.ip += 2
}

// JumpIfTrueOpCode checks if the first argument is not zero, and if so jumps to the second argument.
func JumpIfTrueOpCode(vm *VirtualMachine, arg ArgFunc) {
	if arg(0) != 0 {
		vm.ip = arg(1)
	} else {
		vm.ip += 3
	}
}

// JumpIfFalseOpCode checks if the first argument is zero, and if so jumps to the second argument.
func JumpIfFalseOpCode(vm *VirtualMachine, arg ArgFunc) {
	if arg(0) == 0 {
		vm.ip = arg(1)
	} else {
		vm.ip += 3
	}
}

// LessThanOpCode checks if the first argument is less than the second, and stores the result at the address given
// by the third argument.
func LessThanOpCode(vm *VirtualMachine, arg ArgFunc) {
	if arg(0) < arg(1) {
		vm.Memory[vm.Memory[vm.ip+3]] = 1
	} else {
		vm.Memory[vm.Memory[vm.ip+3]] = 0
	}
	vm.ip += 4
}

// EqualsOpCode checks if the first argument is equal to the second, and stores the result at the address given
// by the third argument.
func EqualsOpCode(vm *VirtualMachine, arg ArgFunc) {
	if arg(0) == arg(1) {
		vm.Memory[vm.Memory[vm.ip+3]] = 1
	} else {
		vm.Memory[vm.Memory[vm.ip+3]] = 0
	}
	vm.ip += 4
}

// HaltOpcode halts the VM and takes no arguments.
func HaltOpcode(vm *VirtualMachine, arg ArgFunc) {
	vm.Halted = true
}
