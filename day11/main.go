// Solution for Advent of Code 2021 -- Day 11
// https://adventofcode.com/2021/day/11

package main

import (
	"fmt"
	"strconv"

	"github.com/huderlem/adventofcode2021/util"
)

type point struct {
	x, y int
}

func parseInput() (map[point]int, int, int) {
	grid := map[point]int{}
	lines := util.ReadFileLines("input.txt")
	for y, l := range lines {
		for x, n := range l {
			num, _ := strconv.Atoi(string(n))
			grid[point{x, y}] = num
		}
	}
	return grid, len(lines[0]), len(lines)
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func getNeighbors(x, y, width, height int) []point {
	candidates := []point{
		{x + 1, y},
		{x + 1, y + 1},
		{x, y + 1},
		{x - 1, y + 1},
		{x - 1, y},
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
	}
	neighbors := []point{}
	for _, c := range candidates {
		if inBounds(c.x, c.y, width, height) {
			neighbors = append(neighbors, c)
		}
	}
	return neighbors
}

func flash(grid map[point]int, p point, width, height int, numFlashes *int) {
	*numFlashes += 1
	neighbors := getNeighbors(p.x, p.y, width, height)
	for _, n := range neighbors {
		grid[n] += 1
		if grid[n] == 10 {
			flash(grid, n, width, height, numFlashes)
		}
	}
}

func simulateStep(grid map[point]int, width, height int) int {
	numFlashes := 0
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			p := point{x, y}
			grid[p] += 1
			if grid[p] == 10 {
				flash(grid, p, width, height, &numFlashes)
			}
		}
	}
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			p := point{x, y}
			if grid[p] > 9 {
				grid[p] = 0
			}
		}
	}
	return numFlashes
}

func part1() int {
	grid, width, height := parseInput()
	numFlashes := 0
	for i := 0; i < 100; i += 1 {
		numFlashes += simulateStep(grid, width, height)
	}
	return numFlashes
}

func part2() int {
	grid, width, height := parseInput()
	numPoints := width * height
	for i := 0; ; i += 1 {
		numFlashes := simulateStep(grid, width, height)
		if numFlashes == numPoints {
			return i + 1
		}
	}
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
