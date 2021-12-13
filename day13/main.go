// Solution for Advent of Code 2021 -- Day 13
// https://adventofcode.com/2021/day/13

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

type point struct {
	x, y int
}

type foldInstruction struct {
	axis  rune
	value int
}

func parseInput() (map[point]struct{}, []foldInstruction) {
	chunks := util.ReadFileChunks("input.txt")
	points := map[point]struct{}{}
	for _, line := range chunks[0] {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points[point{x, y}] = struct{}{}
	}
	instructions := []foldInstruction{}
	for _, line := range chunks[1] {
		axis := rune(line[11])
		value, _ := strconv.Atoi(line[13:])
		instructions = append(instructions, foldInstruction{axis, value})
	}

	return points, instructions
}

func foldPoint(p point, axis rune, value int) point {
	if axis == 'x' {
		x := p.x
		if x > value {
			x = value - (x - value)
		}
		return point{x, p.y}
	} else {
		y := p.y
		if y > value {
			y = value - (y - value)
		}
		return point{p.x, y}
	}
}

func fold(points map[point]struct{}, instruction foldInstruction) map[point]struct{} {
	foldedPoints := map[point]struct{}{}
	for p := range points {
		foldedPoint := foldPoint(p, instruction.axis, instruction.value)
		if foldedPoint.x >= 0 && foldedPoint.y >= 0 {
			foldedPoints[foldedPoint] = struct{}{}
		}
	}
	return foldedPoints
}

func part1() int {
	points, instructions := parseInput()
	points = fold(points, instructions[0])
	return len(points)
}

func printPoints(points map[point]struct{}) {
	maxX := 0
	maxY := 0
	for p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := 0; y <= maxY; y += 1 {
		for x := 0; x <= maxX; x += 1 {
			if _, ok := points[point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func part2() {
	points, instructions := parseInput()
	for _, instruction := range instructions {
		points = fold(points, instruction)
	}
	printPoints(points)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:")
	part2()
}
