// Solution for Advent of Code 2021 -- Day 1
// https://adventofcode.com/2021/day/1

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2021/util"
)

var depths = util.ReadFileInts("input.txt")

func part1() int {
	count := 0
	for i := 1; i < len(depths); i += 1 {
		if depths[i] > depths[i-1] {
			count += 1
		}
	}
	return count
}

func part2() int {
	count := 0
	for i := 3; i < len(depths); i += 1 {
		if depths[i] > depths[i-3] {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
