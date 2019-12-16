package intcode

// opcodeFunc is a function that describes an opcode implemented in the VM.
type opcodeFunc = func(vm *VirtualMachine, arg1, arg2, arg3 *int)

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

var opcodeArity = [100]int{
	1:  3,
	2:  3,
	3:  1,
	4:  1,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
	9:  1,
	99: 0,
}

// addOpcode takes the values specified by args 1 and 2, adds them together, and stores at the memory address given
// by arg 3.
func addOpcode(vm *VirtualMachine, arg1, arg2, arg3 *int) {
	*arg3 = *arg1 + *arg2
	vm.ip += 4
}

// mulOpcode takes the values specified by args 1 and 2, multiplies them together, and stores at the memory address
// given by arg 3.
func mulOpcode(vm *VirtualMachine, arg1, arg2, arg3 *int) {
	*arg3 = *arg1 * *arg2
	vm.ip += 4
}

// readOpCode reads a value from the input stream and stores it at the memory address given by arg 1.
func readOpCode(vm *VirtualMachine, arg1, _, _ *int) {
	*arg1 = <-vm.Input
	vm.ip += 2
}

// writeOpCode writes the value specified by the first argument to the output stream.
func writeOpCode(vm *VirtualMachine, arg1, _, _ *int) {
	vm.Output <- *arg1
	vm.ip += 2
}

// jumpIfTrueOpCode checks if the first argument is not zero, and if so jumps to the second argument.
func jumpIfTrueOpCode(vm *VirtualMachine, arg1, arg2, _ *int) {
	if *arg1 != 0 {
		vm.ip = *arg2
	} else {
		vm.ip += 3
	}
}

// jumpIfFalseOpCode checks if the first argument is zero, and if so jumps to the second argument.
func jumpIfFalseOpCode(vm *VirtualMachine, arg1, arg2, _ *int) {
	if *arg1 == 0 {
		vm.ip = *arg2
	} else {
		vm.ip += 3
	}
}

// lessThanOpCode checks if the first argument is less than the second, and stores the result at the address given
// by the third argument.
func lessThanOpCode(vm *VirtualMachine, arg1, arg2, arg3 *int) {
	if *arg1 < *arg2 {
		*arg3 = 1
	} else {
		*arg3 = 0
	}
	vm.ip += 4
}

// equalsOpCode checks if the first argument is equal to the second, and stores the result at the address given
// by the third argument.
func equalsOpCode(vm *VirtualMachine, arg1, arg2, arg3 *int) {
	if *arg1 == *arg2 {
		*arg3 = 1
	} else {
		*arg3 = 0
	}
	vm.ip += 4
}

// relativeBaseOffsetOpCode increases the relative base by the given argument.
func relativeBaseOffsetOpCode(vm *VirtualMachine, arg1, _, _ *int) {
	vm.rb += *arg1
	vm.ip += 2
}

// haltOpcode halts the VM and takes no arguments.
func haltOpcode(vm *VirtualMachine, _, _, _ *int) {
	vm.Halted = true
}
