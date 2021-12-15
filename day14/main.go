// Solution for Advent of Code 2021 -- Day 14
// https://adventofcode.com/2021/day/14

package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

func parseInput() (map[string]int, rune, rune, map[string]rune) {
	chunks := util.ReadFileChunks("input.txt")
	template := map[string]int{}
	rules := map[string]rune{}
	t := chunks[0][0]
	for i := 0; i < len(t)-1; i += 1 {
		pair := string(t[i]) + string(t[i+1])
		if _, ok := template[pair]; !ok {
			template[pair] = 1
		} else {
			template[pair] += 1
		}
	}
	for _, line := range chunks[1] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = rune(parts[1][0])
	}

	return template, rune(t[0]), rune(t[len(t)-1]), rules
}

func applyRules(template map[string]int, rules map[string]rune) map[string]int {
	newTemplate := map[string]int{}
	for pair, num := range template {
		t1 := string(pair[0]) + string(rules[pair])
		t2 := string(rules[pair]) + string(pair[1])
		if _, ok := newTemplate[t1]; !ok {
			newTemplate[t1] = num
		} else {
			newTemplate[t1] += num
		}
		if _, ok := newTemplate[t2]; !ok {
			newTemplate[t2] = num
		} else {
			newTemplate[t2] += num
		}
	}
	return newTemplate
}

func countElements(template map[string]int, first, last rune) map[rune]int {
	counts := map[rune]int{}
	for pair, n := range template {
		t1, t2 := rune(pair[0]), rune(pair[1])
		if _, ok := counts[t1]; !ok {
			counts[t1] = n
		} else {
			counts[t1] += n
		}
		if _, ok := counts[t2]; !ok {
			counts[t2] = n
		} else {
			counts[t2] += n
		}
	}
	for r := range counts {
		counts[r] /= 2
	}
	counts[first] += 1
	counts[last] += 1
	return counts
}

func getMinAndMaxElementCounts(template map[string]int, first, last rune) (int, int) {
	counts := countElements(template, first, last)
	min, max := math.MaxInt64, math.MinInt64
	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	return min, max
}

func simulate(n int) int {
	template, first, last, rules := parseInput()
	for i := 0; i < n; i += 1 {
		template = applyRules(template, rules)
	}
	minCount, maxCount := getMinAndMaxElementCounts(template, first, last)
	return maxCount - minCount
}

func part1() int {
	return simulate(10)
}

func part2() int {
	return simulate(40)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
