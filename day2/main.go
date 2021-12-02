// Solution for Advent of Code 2021 -- Day 2
// https://adventofcode.com/2021/day/2

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

type command struct {
	direction string
	amount    int
}

type position struct {
	x     int
	depth int
	aim   int
}

func (p *position) applyCommandPart1(c command) {
	switch c.direction {
	case "forward":
		p.x += c.amount
	case "up":
		p.depth -= c.amount
	case "down":
		p.depth += c.amount
	}
}

func (p *position) applyCommandPart2(c command) {
	switch c.direction {
	case "forward":
		p.x += c.amount
		p.depth += p.aim * c.amount
	case "up":
		p.aim -= c.amount
	case "down":
		p.aim += c.amount
	}
}

func parseInput() []command {
	var lines = util.ReadFileLines("input.txt")
	commands := []command{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		commands = append(commands, command{direction, amount})
	}
	return commands
}

var commands = parseInput()

func part1() int {
	position := position{}
	for _, c := range commands {
		position.applyCommandPart1(c)
	}
	return position.x * position.depth
}

func part2() int {
	position := position{}
	for _, c := range commands {
		position.applyCommandPart2(c)
	}
	return position.x * position.depth
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
