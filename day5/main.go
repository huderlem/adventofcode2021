// Solution for Advent of Code 2021 -- Day 5
// https://adventofcode.com/2021/day/5

package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/huderlem/adventofcode2021/util"
)

type point struct {
	x, y int
}

type line struct {
	p1, p2 point
}

func (l *line) isStraight() bool {
	return l.p1.x == l.p2.x || l.p1.y == l.p2.y
}

func parseInput() []line {
	var inputLines = util.ReadFileLines("input.txt")
	lines := []line{}
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	for _, l := range inputLines {
		match := r.FindStringSubmatch(l)
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])
		lines = append(lines, line{point{x1, y1}, point{x2, y2}})
	}
	return lines
}

var lines = parseInput()

func incrementPointCount(points map[point]int, p point, numOverlaps *int) {
	if _, ok := points[p]; !ok {
		points[p] = 1
	} else {
		points[p] += 1
		if points[p] == 2 {
			*numOverlaps += 1
		}
	}
}

func markHorizontalLine(l line, points map[point]int, numOverlaps *int) {
	delta := 1
	if l.p1.y > l.p2.y {
		delta = -1
	}
	for y := l.p1.y; ; y += delta {
		incrementPointCount(points, point{l.p1.x, y}, numOverlaps)
		if y == l.p2.y {
			break
		}
	}
}

func markVerticalLine(l line, points map[point]int, numOverlaps *int) {
	delta := 1
	if l.p1.x > l.p2.x {
		delta = -1
	}
	for x := l.p1.x; ; x += delta {
		incrementPointCount(points, point{x, l.p1.y}, numOverlaps)
		if x == l.p2.x {
			break
		}
	}
}

func markDiagonalLine(l line, points map[point]int, numOverlaps *int) {
	deltaX := 1
	if l.p1.x > l.p2.x {
		deltaX = -1
	}
	deltaY := 1
	if l.p1.y > l.p2.y {
		deltaY = -1
	}
	x := l.p1.x
	y := l.p1.y
	for {
		incrementPointCount(points, point{x, y}, numOverlaps)
		if x == l.p2.x && y == l.p2.y {
			break
		}
		x += deltaX
		y += deltaY
	}
}

func part1() int {
	points := map[point]int{}
	numOverlaps := 0
	for _, l := range lines {
		if l.isStraight() {
			if l.p1.x == l.p2.x {
				markHorizontalLine(l, points, &numOverlaps)
			} else {
				markVerticalLine(l, points, &numOverlaps)
			}
		}
	}
	return numOverlaps
}

func part2() int {
	points := map[point]int{}
	numOverlaps := 0
	for _, l := range lines {
		if l.isStraight() {
			if l.p1.x == l.p2.x {
				markHorizontalLine(l, points, &numOverlaps)
			} else {
				markVerticalLine(l, points, &numOverlaps)
			}
		} else {
			markDiagonalLine(l, points, &numOverlaps)
		}
	}
	return numOverlaps
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
