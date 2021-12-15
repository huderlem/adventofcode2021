// Solution for Advent of Code 2021 -- Day 15
// https://adventofcode.com/2021/day/15

package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"

	"github.com/huderlem/adventofcode2021/util"
)

type point struct {
	x, y int
}

type node struct {
	pos            point
	risk           int
	gScore, fScore int
	prev           point
	index          int
}

func parseInput() (map[point]*node, int, int) {
	lines := util.ReadFileLines("input.txt")
	nodes := map[point]*node{}
	width := len(lines[0])
	height := len(lines)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			risk, _ := strconv.Atoi(string(lines[y][x]))
			p := point{x, y}
			nodes[p] = &node{
				pos:    p,
				risk:   risk,
				gScore: math.MaxInt64,
				fScore: math.MaxInt64,
			}
		}
	}
	return nodes, width, height
}

func (n node) manhattanDistance(dest point) int {
	return int(math.Abs(float64(dest.x-n.pos.x))) + int(math.Abs(float64(dest.y-n.pos.y)))
}

func (n node) getNeighbors(width, height int) []point {
	neighbors := []point{}
	if n.pos.x > 0 {
		neighbors = append(neighbors, point{n.pos.x - 1, n.pos.y})
	}
	if n.pos.y > 0 {
		neighbors = append(neighbors, point{n.pos.x, n.pos.y - 1})
	}
	if n.pos.x < width-1 {
		neighbors = append(neighbors, point{n.pos.x + 1, n.pos.y})
	}
	if n.pos.y < height-1 {
		neighbors = append(neighbors, point{n.pos.x, n.pos.y + 1})
	}
	return neighbors
}

func buildPath(nodes map[point]*node, origin, dest point) []*node {
	path := []*node{}
	curNode := nodes[dest]
	for curNode.pos != origin {
		path = append([]*node{curNode}, path...)
		curNode = nodes[curNode.prev]
	}
	return path
}

func findPath(nodes map[point]*node, origin, dest point, width, height int) []*node {
	// Straightforward A* pathfinding algorithm.
	nodes[origin].fScore = 0
	candidates := PriorityQueue{nodes[origin]}
	heap.Init(&candidates)
	for {
		curNode := heap.Pop(&candidates).(*node)
		if curNode.pos == dest {
			return buildPath(nodes, origin, dest)
		}
		neighbors := curNode.getNeighbors(width, height)
		for _, neighborPos := range neighbors {
			neighbor := nodes[neighborPos]
			gScore := curNode.gScore + neighbor.risk
			if gScore < neighbor.gScore {
				neighbor.gScore = gScore
				neighbor.fScore = gScore + neighbor.manhattanDistance(dest)
				neighbor.prev = curNode.pos
				found := false
				for _, n := range candidates {
					if n.pos == neighbor.pos {
						found = true
						break
					}
				}
				if !found {
					heap.Push(&candidates, neighbor)
				} else {
					candidates.update(neighbor, neighbor.fScore)
				}
			}
		}
	}
}

func sumRisk(nodes []*node) int {
	sum := 0
	for _, n := range nodes {
		sum += n.risk
	}
	return sum
}

func part1() int {
	nodes, width, height := parseInput()
	origin, dest := point{0, 0}, point{width - 1, height - 1}
	path := findPath(nodes, origin, dest, width, height)
	return sumRisk(path)
}

func expandMap(nodes map[point]*node, width, height, factor int) (map[point]*node, int, int) {
	for i := 0; i < factor; i += 1 {
		for j := 0; j < factor; j += 1 {
			if i == 0 && j == 0 {
				continue
			}
			diff := i + j
			for x := 0; x < width; x += 1 {
				for y := 0; y < height; y += 1 {
					p := point{
						x: x + i*width,
						y: y + j*height,
					}
					nodes[p] = &node{
						pos:    p,
						risk:   ((nodes[point{x, y}].risk + diff - 1) % 9) + 1,
						gScore: math.MaxInt64,
						fScore: math.MaxInt64,
					}
				}
			}
		}
	}
	return nodes, width * factor, height * factor
}

func part2() int {
	nodes, width, height := parseInput()
	nodes, width, height = expandMap(nodes, width, height, 5)
	origin, dest := point{0, 0}, point{width - 1, height - 1}
	path := findPath(nodes, origin, dest, width, height)
	return sumRisk(path)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
