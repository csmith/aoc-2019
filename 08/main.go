package main

import (
	"fmt"
	"github.com/csmith/aoc-2019/common"
	"math"
)

const width = 25
const height = 6

func checkChecksum(layer string, fewestZeros *int, checksum *int) {
	counts := []int{0, 0, 0}
	for _, r := range layer {
		counts[r-'0']++
	}

	if counts[0] < *fewestZeros {
		*fewestZeros = counts[0]
		*checksum = counts[1] * counts[2]
	}
}

func main() {
	var (
		input       = common.ReadFileAsStrings("08/input.txt")[0]
		fewestZeros = math.MaxInt64
		checksum    = 0
		pixels      [width * height]rune
	)

	for i := len(input) - width*height; i >= 0; i -= width * height {
		layer := input[i : i+width*height]
		checkChecksum(layer, &fewestZeros, &checksum)
		for i, r := range layer {
			if r == '0' {
				pixels[i] = ' '
			} else if r == '1' {
				pixels[i] = '█'
			}
		}
	}

	fmt.Printf("%d\n", checksum)
	for i, p := range pixels {
		fmt.Printf("%c", p)
		if i%width == width-1 {
			fmt.Printf("\n")
		}
	}
}
