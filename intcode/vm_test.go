package intcode

import (
	"reflect"
	"testing"
)

func TestDayTwoSamples(t *testing.T) {
	tables := []struct {
		given    []int
		expected []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, table := range tables {
		vm := NewVirtualMachine(table.given)
		vm.Run()
		if !reflect.DeepEqual(table.expected, vm.Memory) {
			t.Errorf("Evaluation of %v was incorrect, got: %v, want: %v.", table.given, vm.Memory, table.expected)
		}
	}
}

func TestDayFiveSamples(t *testing.T) {
	tables := []struct {
		given  []int
		input  []int
		output []int
	}{
		// Reads then outputs a number
		{[]int{3, 0, 4, 0, 99}, []int{123}, []int{123}},
		// Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
		{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8}, []int{1}},
		{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{7}, []int{0}},
		// Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
		{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{7}, []int{1}},
		{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8}, []int{0}},
		// Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
		{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{8}, []int{1}},
		{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{9}, []int{0}},
		// Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
		{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{0}, []int{1}},
		{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{10}, []int{0}},
	}

	for _, table := range tables {
		vm := NewVirtualMachine(table.given)
		vm.Input = make(chan int, 1)
		vm.Output = make(chan int, 1)

		for _, v := range table.input {
			vm.Input <- v
		}

		vm.Run()

		for _, v := range table.output {
			actual := <-vm.Output
			if !reflect.DeepEqual(v, actual) {
				t.Errorf("Wrong output value received for %v, got: %v, want: %v.", table.given, actual, v)
			}
		}
	}
}
