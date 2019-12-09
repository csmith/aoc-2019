package main

import (
	"github.com/csmith/aoc-2019/common"
	"runtime/debug"
	"testing"
)

func Benchmark(b *testing.B) {
	debug.SetGCPercent(-1)
	for i := 0; i < b.N; i++ {
		input := common.ReadCsvAsInts("input.txt")
		memoryBanks := make([]int, len(input)*5)
		_ = maxOutput(memoryBanks, input, []int{0, 1, 2, 3, 4}, false)
		_ = maxOutput(memoryBanks, input, []int{5, 6, 7, 8, 9}, true)
	}
}
