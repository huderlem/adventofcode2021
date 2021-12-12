// Solution for Advent of Code 2021 -- Day 12
// https://adventofcode.com/2021/day/12

package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/huderlem/adventofcode2021/util"
)

type allowFunc func(node string, path []string) bool

func parseInput() map[string][]string {
	nodes := map[string][]string{}
	lines := util.ReadFileLines("input.txt")
	for _, l := range lines {
		parts := strings.Split(l, "-")
		n1, n2 := parts[0], parts[1]
		if _, ok := nodes[n1]; !ok {
			nodes[n1] = []string{}
		}
		if _, ok := nodes[n2]; !ok {
			nodes[n2] = []string{}
		}
		nodes[n1] = append(nodes[n1], n2)
		nodes[n2] = append(nodes[n2], n1)
	}
	return nodes
}

func isSmallCave(name string) bool {
	return unicode.IsLower(rune(name[0]))
}

func part1AllowRevisit(node string, path []string) bool {
	if !isSmallCave(node) {
		return true
	}
	for _, n := range path {
		if n == node {
			return false
		}
	}
	return true
}

func part2AllowRevisit(node string, path []string) bool {
	if !isSmallCave(node) {
		return true
	}
	if node == "start" {
		return false
	}
	visited := map[string]struct{}{}
	visitedSmallCaveTwice := false
	for _, n := range path {
		if isSmallCave(n) {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
			} else {
				visitedSmallCaveTwice = true
			}
		}
	}
	if _, ok := visited[node]; ok && visitedSmallCaveTwice {
		return false
	}
	return true
}

func extendPath(curPath []string, node string) []string {
	path := make([]string, len(curPath))
	copy(path, curPath)
	return append(path, node)
}

func getAllPaths(nodes map[string][]string, curNode string, curPath []string, result [][]string, allowFunc allowFunc) [][]string {
	if curNode == "end" {
		return append(result, curPath)
	}
	neighbors := nodes[curNode]
	for _, neighbor := range neighbors {
		if !allowFunc(neighbor, curPath) {
			continue
		}
		path := extendPath(curPath, neighbor)
		result = getAllPaths(nodes, neighbor, path, result, allowFunc)
	}
	return result
}

func part1() int {
	nodes := parseInput()
	paths := getAllPaths(nodes, "start", []string{"start"}, [][]string{}, part1AllowRevisit)
	return len(paths)
}

func part2() int {
	nodes := parseInput()
	paths := getAllPaths(nodes, "start", []string{"start"}, [][]string{}, part2AllowRevisit)
	return len(paths)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
