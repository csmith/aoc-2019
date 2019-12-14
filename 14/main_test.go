package main

import (
	"github.com/csmith/aoc-2019/common"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recipes := parseInput("input.txt")
		used := make(map[chemical]int, len(recipes))
		spare := make(map[chemical]int, len(recipes))

		recipes.produce("FUEL", 1, used, spare)
		orePerFuel := used["ORE"]

		last := 0
		for used["ORE"] < 1000000000000 {
			last = used["FUEL"]
			recipes.produce("FUEL", common.Max(1, (1000000000000-used["ORE"])/orePerFuel), used, spare)
		}

		_ = last
	}
}
