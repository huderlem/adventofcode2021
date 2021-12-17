// Solution for Advent of Code 2021 -- Day 17
// https://adventofcode.com/2021/day/17

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

func parseInput() (int, int, int, int) {
	input := util.ReadFileString("input.txt")[15:]
	parts := strings.Split(input, ", y=")
	x := strings.Split(parts[0], "..")
	y := strings.Split(parts[1], "..")
	left, _ := strconv.Atoi(x[0])
	right, _ := strconv.Atoi(x[1])
	bottom, _ := strconv.Atoi(y[0])
	top, _ := strconv.Atoi(y[1])
	return left, right, bottom, top
}

func getYVelocityBounds(bottom int) (int, int) {
	return bottom, -bottom - 1
}

func getXVelocityBounds(left, right int) (int, int) {
	// Quadratic formula to solve n*n + n - 2*left = 0
	min := (-1 + math.Sqrt(1.0+8*float64(left))) / 2.0
	return int(math.Round(min + 0.5)), right
}

func part1() int {
	_, _, bottom, _ := parseInput()
	_, maxYVelocity := getYVelocityBounds(bottom)
	return maxYVelocity * (maxYVelocity + 1) / 2
}

type velocity struct {
	x, y int
}

func (v *velocity) update() {
	if v.x < 0 {
		v.x += 1
	} else if v.x > 0 {
		v.x -= 1
	}
	v.y -= 1
}

func simulate(v velocity, left, right, bottom, top int) bool {
	x, y := 0, 0
	for x <= right && y >= bottom {
		if x >= left && x <= right && y >= bottom && y <= top {
			return true
		}
		x += v.x
		y += v.y
		v.update()
	}
	return false
}

func part2() int {
	left, right, bottom, top := parseInput()
	minYVelocity, maxYVelocity := getYVelocityBounds(bottom)
	minXVelocity, maxXVelocity := getXVelocityBounds(left, right)
	result := map[velocity]struct{}{}
	for dx := minXVelocity; dx <= maxXVelocity; dx += 1 {
		for dy := minYVelocity; dy <= maxYVelocity; dy += 1 {
			v := velocity{dx, dy}
			if simulate(v, left, right, bottom, top) {
				result[v] = struct{}{}
			}
		}
	}
	return len(result)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
