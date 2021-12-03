// Solution for Advent of Code 2021 -- Day 3
// https://adventofcode.com/2021/day/3

package main

import (
	"fmt"
	"strconv"

	"github.com/huderlem/adventofcode2021/util"
)

var inputLines = util.ReadFileLines("input.txt")

// getBitCounts creates a slice whose elements indicate
// the relative number of 1 or 0 bits in the lines' columns.
// A value of "5" would mean there were 5 more 1's than 0's, whereas
// a value of "-3" would mean there were 3 more 0's than 1's.
func getBitCounts(lines []string) []int {
	bitCounts := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, bit := range line {
			if bit == '0' {
				bitCounts[i] -= 1
			} else {
				bitCounts[i] += 1
			}
		}
	}
	return bitCounts
}

func part1() int {
	bitCounts := getBitCounts(inputLines)
	gamma := 0
	epsilon := 0
	for i := 0; i < len(bitCounts); i += 1 {
		gamma <<= 1
		epsilon <<= 1
		if bitCounts[i] > 0 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	return gamma * epsilon
}

func filterMatchingLines(lines []string, bitIndex int, invertComparison bool) []string {
	bitCounts := getBitCounts(lines)
	expectedBit := '0'
	if invertComparison {
		if bitCounts[bitIndex] < 0 {
			expectedBit = '1'
		}
	} else {
		if bitCounts[bitIndex] >= 0 {
			expectedBit = '1'
		}
	}
	for i := 0; i < len(lines); {
		if lines[i][bitIndex] != byte(expectedBit) {
			// Remove the line from our slice, since it doesn't match our filtering criteria.
			lines = append(lines[:i], lines[i+1:]...)
		} else {
			i += 1
		}
	}
	return lines
}

func getFilteredRating(invertComparison bool) int {
	numBits := len(inputLines[0])
	remainingLines := make([]string, len(inputLines))
	copy(remainingLines, inputLines)
	// Filter out lines by bit index until exactly one remains.
	for bitIndex := 0; bitIndex < numBits && len(remainingLines) > 1; bitIndex += 1 {
		remainingLines = filterMatchingLines(remainingLines, bitIndex, invertComparison)
	}
	rating, _ := strconv.ParseInt(remainingLines[0], 2, 32)
	return int(rating)
}

func part2() int {
	oxygen := getFilteredRating(false)
	c02 := getFilteredRating(true)
	return oxygen * c02
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
