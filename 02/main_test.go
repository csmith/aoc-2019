package main

import (
	"github.com/csmith/aoc-2019/common"
	"github.com/csmith/aoc-2019/intcode"
	"runtime/debug"
	"sync"
	"testing"
)

func findOutputMany(pool *sync.Pool, input []int, target int) int {
	res := make(chan int)

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			go func(n, v int) {
				if run(pool, input, n, v) == target {
					res <- 100*n + v
				}
			}(n, v)
		}

	}

	return <-res
}

func findOutputLinear(pool *sync.Pool, input []int, target int) int {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			if run(pool, input, n, v) == target {
				return 100*n + v
			}
		}
	}
	panic("Should've returned")
}

func Benchmark_goroutine_per_pair(b *testing.B) {
	debug.SetGCPercent(-1)
	for i := 0; i < b.N; i++ {
		input := common.ReadCsvAsInts("input.txt")

		pool := &sync.Pool{
			New: func() interface{} {
				return intcode.NewVirtualMachine(make([]int, len(input)))
			},
		}

		_ = run(pool, input, 12, 2)
		_ = findOutputMany(pool, input, 19690720)
	}
}

func Benchmark_goroutine_per_noun(b *testing.B) {
	debug.SetGCPercent(-1)
	for i := 0; i < b.N; i++ {
		input := common.ReadCsvAsInts("input.txt")

		pool := &sync.Pool{
			New: func() interface{} {
				return intcode.NewVirtualMachine(make([]int, len(input)))
			},
		}

		_ = run(pool, input, 12, 2)
		_ = findOutput(pool, input, 19690720)
	}
}

func Benchmark_no_goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := common.ReadCsvAsInts("input.txt")

		pool := &sync.Pool{
			New: func() interface{} {
				return intcode.NewVirtualMachine(make([]int, len(input)))
			},
		}

		_ = run(pool, input, 12, 2)
		_ = findOutputLinear(pool, input, 19690720)
	}
}
