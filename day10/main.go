// Solution for Advent of Code 2021 -- Day 10
// https://adventofcode.com/2021/day/10

package main

import (
	"fmt"
	"sort"

	"github.com/huderlem/adventofcode2021/util"
)

type syntaxResult int

const (
	syntaxValid syntaxResult = iota
	syntaxCorrupt
	syntaxUnfinished
)

var openChars = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}
var closeChars = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}
var corruptScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}
var unfinishedScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func checkSyntax(line string) (syntaxResult, int, []rune) {
	unclosed := []rune{}
	for i, c := range line {
		if _, ok := openChars[c]; ok {
			unclosed = append(unclosed, c)
		} else {
			if unclosed[len(unclosed)-1] == closeChars[c] {
				unclosed = unclosed[:len(unclosed)-1]
			} else {
				return syntaxCorrupt, i, unclosed
			}
		}
	}
	if len(unclosed) > 0 {
		return syntaxUnfinished, len(line), unclosed
	}
	return syntaxValid, 0, unclosed
}

func part1() int {
	lines := util.ReadFileLines("input.txt")
	score := 0
	for _, l := range lines {
		result, index, _ := checkSyntax(l)
		if result == syntaxCorrupt {
			score += corruptScores[rune(l[index])]
		}
	}
	return score
}

func getScoreUnclosed(unclosed []rune) int {
	score := 0
	for i := len(unclosed) - 1; i >= 0; i -= 1 {
		score *= 5
		score += unfinishedScores[unclosed[i]]
	}
	return score
}

func part2() int {
	lines := util.ReadFileLines("input.txt")
	scores := []int{}
	for _, l := range lines {
		result, _, unclosed := checkSyntax(l)
		if result == syntaxUnfinished {
			scores = append(scores, getScoreUnclosed(unclosed))
		}
	}
	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	return scores[len(scores)/2]
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
