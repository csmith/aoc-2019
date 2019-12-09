package intcode

// opcodeFunc is a function that describes an opcode implemented in the VM.
type opcodeFunc = func(vm *VirtualMachine)

var opcodes = [100]opcodeFunc{
	1:  addOpcode,
	2:  mulOpcode,
	3:  readOpCode,
	4:  writeOpCode,
	5:  jumpIfTrueOpCode,
	6:  jumpIfFalseOpCode,
	7:  lessThanOpCode,
	8:  equalsOpCode,
	9:  relativeBaseOffsetOpCode,
	99: haltOpcode,
}

// addOpcode takes the values specified by args 1 and 2, adds them together, and stores at the memory address given
// by arg 3.
func addOpcode(vm *VirtualMachine) {
	*vm.arg(2) = *vm.arg(0) + *vm.arg(1)
	vm.ip += 4
}

// mulOpcode takes the values specified by args 1 and 2, multiplies them together, and stores at the memory address
// given by arg 3.
func mulOpcode(vm *VirtualMachine) {
	*vm.arg(2) = *vm.arg(0) * *vm.arg(1)
	vm.ip += 4
}

// readOpCode reads a value from the input stream and stores it at the memory address given by arg 1.
func readOpCode(vm *VirtualMachine) {
	*vm.arg(0) = <-vm.Input
	vm.ip += 2
}

// writeOpCode writes the value specified by the first argument to the output stream.
func writeOpCode(vm *VirtualMachine) {
	vm.Output <- *vm.arg(0)
	vm.ip += 2
}

// jumpIfTrueOpCode checks if the first argument is not zero, and if so jumps to the second argument.
func jumpIfTrueOpCode(vm *VirtualMachine) {
	if *vm.arg(0) != 0 {
		vm.ip = *vm.arg(1)
	} else {
		vm.ip += 3
	}
}

// jumpIfFalseOpCode checks if the first argument is zero, and if so jumps to the second argument.
func jumpIfFalseOpCode(vm *VirtualMachine) {
	if *vm.arg(0) == 0 {
		vm.ip = *vm.arg(1)
	} else {
		vm.ip += 3
	}
}

// lessThanOpCode checks if the first argument is less than the second, and stores the result at the address given
// by the third argument.
func lessThanOpCode(vm *VirtualMachine) {
	if *vm.arg(0) < *vm.arg(1) {
		*vm.arg(2) = 1
	} else {
		*vm.arg(2) = 0
	}
	vm.ip += 4
}

// equalsOpCode checks if the first argument is equal to the second, and stores the result at the address given
// by the third argument.
func equalsOpCode(vm *VirtualMachine) {
	if *vm.arg(0) == *vm.arg(1) {
		*vm.arg(2) = 1
	} else {
		*vm.arg(2) = 0
	}
	vm.ip += 4
}

// relativeBaseOffsetOpCode increases the relative base by the given argument.
func relativeBaseOffsetOpCode(vm *VirtualMachine) {
	vm.rb += *vm.arg(0)
	vm.ip += 2
}

// haltOpcode halts the VM and takes no arguments.
func haltOpcode(vm *VirtualMachine) {
	vm.Halted = true
}
