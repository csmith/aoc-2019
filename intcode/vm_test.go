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
