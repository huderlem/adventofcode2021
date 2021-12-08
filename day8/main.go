// Solution for Advent of Code 2021 -- Day 8
// https://adventofcode.com/2021/day/8

package main

import (
	"fmt"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

type data struct {
	patterns []string
	outputs  []string
}

func parseInput() []data {
	lines := util.ReadFileLines("input.txt")
	d := []data{}
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		d = append(d, data{
			patterns: strings.Split(parts[0], " "),
			outputs:  strings.Split(parts[1], " "),
		})
	}
	return d
}

func part1() int {
	data := parseInput()
	count := 0
	for _, d := range data {
		for _, o := range d.outputs {
			if len(o) == 2 || len(o) == 3 || len(o) == 4 || len(o) == 7 {
				count += 1
			}
		}
	}
	return count
}

func getKnownPatterns(patterns []string) (map[rune]struct{}, map[rune]struct{}) {
	one := map[rune]struct{}{}
	four := map[rune]struct{}{}
	for _, p := range patterns {
		if len(p) == 2 {
			for _, r := range p {
				one[r] = struct{}{}
			}
		} else if len(p) == 4 {
			for _, r := range p {
				four[r] = struct{}{}
			}
		}
	}

	return one, four
}

func countLetterSimilarities(pattern string, known map[rune]struct{}) (int, int) {
	numShared := 0
	numDifferent := 0
	for _, r := range pattern {
		if _, ok := known[r]; ok {
			numShared += 1
		} else {
			numDifferent += 1
		}
	}
	return numShared, numDifferent
}

// Key comprises of the number of different/shared lines on the clock
// compared to digits 1 and 4.
// {1 difference}/{1 shared},{4 difference}/{4 shared}
var fiveLettersMap = map[string]int{
	"4/1,3/2": 2,
	"3/2,2/3": 3,
	"4/1,2/3": 5,
}
var sixLettersMap = map[string]int{
	"4/2,3/3": 0,
	"5/1,3/3": 6,
	"4/2,2/4": 9,
}

// These are the digits that have unique segment lengths.
var knownLengthMap = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

func determineOutputs(patterns []string, outputs []string) []int {
	outputValues := []int{}
	oneLetters, fourLetters := getKnownPatterns(patterns)
	for _, output := range outputs {
		switch len(output) {
		case 5, 6:
			// By comparing the number of shared and differing letters to the known
			// one and four digits, we can determine exactly what each 5- and 6-letter
			// pattern is. This is because all of the 5- or 6-letter digits have unique
			// numbers of different/shared segments with one and four.
			numSharedLettersOne, numDifferentLettersOne := countLetterSimilarities(output, oneLetters)
			numSharedLettersFour, numDifferentLettersFour := countLetterSimilarities(output, fourLetters)
			key := fmt.Sprintf("%d/%d,%d/%d", numDifferentLettersOne, numSharedLettersOne, numDifferentLettersFour, numSharedLettersFour)
			if len(output) == 5 {
				outputValues = append(outputValues, fiveLettersMap[key])
			} else {
				outputValues = append(outputValues, sixLettersMap[key])
			}
		default:
			outputValues = append(outputValues, knownLengthMap[len(output)])
		}
	}
	return outputValues
}

func part2() int {
	data := parseInput()
	sum := 0
	for _, d := range data {
		outputValues := determineOutputs(d.patterns, d.outputs)
		num := 0
		for _, value := range outputValues {
			num *= 10
			num += value
		}
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
