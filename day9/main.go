// Solution for Advent of Code 2021 -- Day 9
// https://adventofcode.com/2021/day/9

package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/huderlem/adventofcode2021/util"
)

type point struct {
	x, y int
}

func parseInput() [][]int {
	lines := util.ReadFileLines("input.txt")
	heightMap := make([][]int, len(lines[0]))
	for i := 0; i < len(lines[0]); i += 1 {
		heightMap[i] = make([]int, len(lines))
	}
	for x := 0; x < len(lines[0]); x += 1 {
		for y := 0; y < len(lines); y += 1 {
			height, _ := strconv.Atoi(string(lines[y][x]))
			heightMap[x][y] = height
		}
	}
	return heightMap
}

func inBounds(heightMap [][]int, x, y int) bool {
	width := len(heightMap)
	height := len(heightMap[0])
	return x >= 0 && x < width && y >= 0 && y < height
}

func getNeighbors(heightMap [][]int, x, y int) []point {
	neighbors := []point{}
	if inBounds(heightMap, x-1, y) {
		neighbors = append(neighbors, point{x - 1, y})
	}
	if inBounds(heightMap, x+1, y) {
		neighbors = append(neighbors, point{x + 1, y})
	}
	if inBounds(heightMap, x, y-1) {
		neighbors = append(neighbors, point{x, y - 1})
	}
	if inBounds(heightMap, x, y+1) {
		neighbors = append(neighbors, point{x, y + 1})
	}
	return neighbors
}

func isLowPoint(heightMap [][]int, x, y int) bool {
	p := heightMap[x][y]
	neighbors := getNeighbors(heightMap, x, y)
	for _, neighbor := range neighbors {
		if heightMap[neighbor.x][neighbor.y] <= p {
			return false
		}
	}
	return true
}

func findLowPoints(heightMap [][]int) []point {
	lowPoints := []point{}
	for x := 0; x < len(heightMap); x += 1 {
		for y := 0; y < len(heightMap[0]); y += 1 {
			if isLowPoint(heightMap, x, y) {
				lowPoints = append(lowPoints, point{x, y})
			}
		}
	}
	return lowPoints
}

func part1() int {
	heightMap := parseInput()
	lowPoints := findLowPoints(heightMap)
	riskSum := 0
	for _, p := range lowPoints {
		riskSum += heightMap[p.x][p.y] + 1
	}
	return riskSum
}

func getBasin(heightMap [][]int, origin point) map[point]struct{} {
	basinPoints := map[point]struct{}{
		origin: {},
	}
	// Do breadth-first search, with boundaries at height 9.
	unvisited := []point{origin}
	for len(unvisited) > 0 {
		p := unvisited[0]
		unvisited = unvisited[1:]
		neighbors := getNeighbors(heightMap, p.x, p.y)
		for _, neighbor := range neighbors {
			if _, ok := basinPoints[neighbor]; !ok && heightMap[neighbor.x][neighbor.y] < 9 {
				unvisited = append(unvisited, neighbor)
			}
		}
		basinPoints[p] = struct{}{}
	}
	return basinPoints
}

func part2() int {
	heightMap := parseInput()
	lowPoints := findLowPoints(heightMap)
	basins := []map[point]struct{}{}
	for _, p := range lowPoints {
		basins = append(basins, getBasin(heightMap, p))
	}
	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	return len(basins[0]) * len(basins[1]) * len(basins[2])
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
