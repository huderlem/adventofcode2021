// Solution for Advent of Code 2021 -- Day 7
// https://adventofcode.com/2021/day/7

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

func parseInput() ([]int, int, int) {
	var input = util.ReadFileString("input.txt")
	values := []int{}
	for _, num := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(num)
		values = append(values, n)
	}
	minValue := math.MaxInt32
	maxValue := math.MinInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
		if v > maxValue {
			maxValue = v
		}
	}
	return values, minValue, maxValue
}

var inputValues, minInputValue, maxInputValue = parseInput()

func sumDistance(values []int, position int) int {
	sum := 0
	for _, v := range values {
		distance := util.Abs(v - position)
		sum += distance
	}
	return sum
}

func part1() int {
	minSum := math.MaxInt32
	for i := minInputValue; i < maxInputValue; i += 1 {
		distanceSum := sumDistance(inputValues, i)
		if distanceSum < minSum {
			minSum = distanceSum
		}
	}
	return minSum
}

func sumVariableFuel(values []int, position int) int {
	sum := 0
	for _, v := range values {
		distance := util.Abs(v - position)
		// Since the fuel increases by 1 every step, it can be modeled
		// with a simple closed expression.
		sum += ((distance * distance) + distance) / 2
	}
	return sum
}

func part2() int {
	minSum := math.MaxInt32
	for i := minInputValue; i < maxInputValue; i += 1 {
		distanceSum := sumVariableFuel(inputValues, i)
		if distanceSum < minSum {
			minSum = distanceSum
		}
	}
	return minSum
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
