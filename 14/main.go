package main

import (
	"github.com/csmith/aoc-2019/common"
	"math"
	"strings"
)

type chemical string
type recipeBook map[chemical]*recipe

type amount struct {
	chemical chemical
	quantity int
}

type recipe struct {
	produces *amount
	requires []*amount
}

func parseAmount(text string) *amount {
	parts := strings.Split(text, " ")
	return &amount{
		quantity: common.MustAtoi(parts[0]),
		chemical: chemical(parts[1]),
	}
}

func parseInput(file string) recipeBook {
	recipes := make(map[chemical]*recipe)
	for _, line := range common.ReadFileAsStrings(file) {
		parts := strings.Split(line, " => ")
		ingredients := strings.Split(parts[0], ", ")
		r := &recipe{
			produces: parseAmount(parts[1]),
			requires: make([]*amount, 0, len(ingredients)),
		}
		for _, in := range ingredients {
			r.requires = append(r.requires, parseAmount(in))
		}
		recipes[r.produces.chemical] = r
	}
	return recipes
}

func (rb recipeBook) produce(target *amount, used map[chemical]int, spare map[chemical]int) {
	recipe := rb[target.chemical]
	needed := target.quantity

	if recipe == nil {
		used[target.chemical] += needed
		return
	}

	free := spare[target.chemical]
	if free >= needed {
		used[target.chemical] += needed
		spare[target.chemical] -= needed
		return
	} else {
		used[target.chemical] += free
		spare[target.chemical] = 0
		needed -= free
	}

	runs := int(math.Ceil(float64(needed) / float64(recipe.produces.quantity)))
	for _, i := range recipe.requires {
		rb.produce(&amount{
			chemical: i.chemical,
			quantity: i.quantity * runs,
		}, used, spare)
	}

	used[target.chemical] += needed
	spare[target.chemical] += (recipe.produces.quantity * runs) - needed
}

func main() {
	recipes := parseInput("14/input.txt")
	used := make(map[chemical]int, len(recipes))
	spare := make(map[chemical]int, len(recipes))

	recipes.produce(&amount{
		chemical: "FUEL",
		quantity: 1,
	}, used, spare)

	orePerFuel := used["ORE"]
	println(used["ORE"])

	last := 0
	for used["ORE"] < 1000000000000 {
		last = used["FUEL"]
		recipes.produce(&amount{
			chemical: "FUEL",
			quantity: common.Max(1, (1000000000000-used["ORE"])/orePerFuel),
		}, used, spare)
	}

	println(last)
}
