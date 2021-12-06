// Solution for Advent of Code 2021 -- Day 6
// https://adventofcode.com/2021/day/6

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

func parseInput() ([]int, []int) {
	nums := util.ReadFileLines("input.txt")[0]
	matureGroups := []int{}
	for i := 0; i < 7; i += 1 {
		matureGroups = append(matureGroups, 0)
	}
	freshGroups := []int{}
	for i := 0; i < 9; i += 1 {
		freshGroups = append(freshGroups, 0)
	}
	for _, n := range strings.Split(nums, ",") {
		num, _ := strconv.Atoi(n)
		matureGroups[num] += 1
	}
	return freshGroups, matureGroups
}

func simulateDay(freshGroups, matureGroups []int, day int) {
	freshIndex := day % 9
	matureIndex := day % 7
	numNewMatureFish := freshGroups[freshIndex]
	numNewFreshFish := matureGroups[matureIndex]
	matureGroups[matureIndex] += numNewMatureFish
	freshGroups[freshIndex] += numNewFreshFish
}

func countFish(freshGroups, matureGroups []int) int {
	sum := 0
	for _, n := range freshGroups {
		sum += n
	}
	for _, n := range matureGroups {
		sum += n
	}
	return sum
}

func simulate(freshGroups, matureGroups []int, numDays int) {
	for i := 0; i < numDays; i += 1 {
		simulateDay(freshGroups, matureGroups, i)
	}
}

func part1() int {
	var freshGroups, matureGroups = parseInput()
	simulate(freshGroups, matureGroups, 80)
	return countFish(freshGroups, matureGroups)
}

func part2() int {
	var freshGroups, matureGroups = parseInput()
	simulate(freshGroups, matureGroups, 256)
	return countFish(freshGroups, matureGroups)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
