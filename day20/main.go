// Solution for Advent of Code 2021 -- Day 20
// https://adventofcode.com/2021/day/20

package main

import (
	"fmt"

	"github.com/huderlem/adventofcode2021/util"
)

type point struct {
	x, y int
}

func parseInput() (string, map[point]struct{}, int) {
	chunks := util.ReadFileChunks("input.txt")
	algorithm := chunks[0][0]
	pixels := map[point]struct{}{}
	for y, line := range chunks[1] {
		for x, c := range line {
			if c == '#' {
				pixels[point{x, y}] = struct{}{}
			}
		}
	}
	return algorithm, pixels, len(chunks[1][0])
}

var offsets = []point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{0, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func getPixelAlgorithmIndex(x, y int, pixels map[point]struct{}) int {
	index := 0
	for _, o := range offsets {
		index <<= 1
		if _, ok := pixels[point{x + o.x, y + o.y}]; ok {
			index |= 1
		}
	}
	return index
}

func processPixel(x, y int, pixels map[point]struct{}, algorithm string) bool {
	index := getPixelAlgorithmIndex(x, y, pixels)
	return rune(algorithm[index]) == '#'
}

func processImage(pixels map[point]struct{}, algorithm string, iteration, width, n int) map[point]struct{} {
	newPixels := map[point]struct{}{}
	// Crude, but it does the job for a small number of iterations.
	// Simply simulate a larger area around the original image, rather
	// than cleverly keeping track of the infinite solid grid outside of
	// the image's possible boundaries.
	start := -2 - n*2 - iteration
	end := width + 1 + n*2 + iteration
	for x := start; x <= end; x += 1 {
		for y := start; y <= end; y += 1 {
			if processPixel(x, y, pixels, algorithm) {
				newPixels[point{x, y}] = struct{}{}
			}
		}
	}
	return newPixels
}

func countPixels(pixels map[point]struct{}, iteration, width int) int {
	start := -2 - iteration
	end := width + 1 + iteration
	count := 0
	for y := start; y <= end; y += 1 {
		for x := start; x <= end; x += 1 {
			if _, ok := pixels[point{x, y}]; ok {
				count += 1
			}
		}
	}
	return count
}

func simulate(n int) int {
	algorithm, pixels, width := parseInput()
	for i := 0; i < n; i += 1 {
		pixels = processImage(pixels, algorithm, i, width, n)
	}
	return countPixels(pixels, n, width)
}

func part1() int {
	return simulate(2)
}

func part2() int {
	return simulate(50)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
