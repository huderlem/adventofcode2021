// Solution for Advent of Code 2021 -- Day 12
// https://adventofcode.com/2021/day/12

package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/huderlem/adventofcode2021/util"
)

type nodeKind int

const (
	invalid nodeKind = iota
	literal
	pair
)

type node struct {
	kind   nodeKind
	depth  int
	left   *node
	right  *node
	parent *node
	value  int
}

func (n *node) String() string {
	switch n.kind {
	default:
		return "INVALID"
	case literal:
		return fmt.Sprintf("%d", n.value)
	case pair:
		return fmt.Sprintf("[%s,%s]", n.left, n.right)
	}
}

func (n *node) deepCopy() *node {
	if n.kind == literal {
		return &node{
			kind:  literal,
			depth: n.depth,
			value: n.value,
		}
	} else {
		newNode := &node{
			kind:  pair,
			depth: n.depth,
		}
		leftCopy := n.left.deepCopy()
		leftCopy.parent = newNode
		rightCopy := n.right.deepCopy()
		rightCopy.parent = newNode
		newNode.left = leftCopy
		newNode.right = rightCopy
		return newNode
	}
}

func (n *node) performSplit() {
	n.kind = pair
	n.left = &node{
		kind:   literal,
		value:  n.value / 2,
		parent: n,
		depth:  n.depth + 1,
	}
	n.right = &node{
		kind:   literal,
		value:  int(math.Ceil(float64(n.value) / 2.0)),
		parent: n,
		depth:  n.depth + 1,
	}
}

func (n *node) findLeftLiteral() *node {
	if n.parent == nil {
		return nil
	}
	if n.parent.left != n {
		return n.parent.left.findRightMostLiteral()
	}
	return n.parent.findLeftLiteral()
}

func (n *node) findLeftMostLiteral() *node {
	if n.kind == literal {
		return n
	}
	return n.left.findLeftMostLiteral()
}

func (n *node) findRightMostLiteral() *node {
	if n.kind == literal {
		return n
	}
	return n.right.findRightMostLiteral()
}

func (n *node) findRightLiteral() *node {
	if n.parent == nil {
		return nil
	}
	if n.parent.right != n {
		return n.parent.right.findLeftMostLiteral()
	}
	return n.parent.findRightLiteral()
}

func (n *node) performExplode() {
	left := n.findLeftLiteral()
	if left != nil {
		left.value += n.left.value
	}
	right := n.findRightLiteral()
	if right != nil {
		right.value += n.right.value
	}
	n.kind = literal
	n.value = 0
	n.left = nil
	n.right = nil
}

func (n *node) increaseDepth() {
	n.depth += 1
	if n.kind == pair {
		n.left.increaseDepth()
		n.right.increaseDepth()
	}
}

func (n *node) add(other *node) *node {
	if other == nil {
		return n
	}
	newNode := &node{
		kind:  pair,
		depth: -1,
		left:  n,
		right: other,
	}
	n.parent = newNode
	other.parent = newNode
	newNode.increaseDepth()
	return newNode
}

func (n *node) magnitude() int {
	if n.kind == literal {
		return n.value
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *node) processExplode() bool {
	switch n.kind {
	case literal:
		return false
	case pair:
		if n.depth >= 4 {
			n.performExplode()
			return true
		}
		return n.left.processExplode() || n.right.processExplode()
	}
	panic("Invalid node kind!")
}

func (n *node) processSplit() bool {
	switch n.kind {
	case literal:
		if n.value >= 10 {
			n.performSplit()
			return true
		}
		return false
	case pair:
		return n.left.processSplit() || n.right.processSplit()
	}
	panic("Invalid node kind!")
}

func (n *node) reduce() {
	for {
		if !n.processExplode() && !n.processSplit() {
			return
		}
	}
}

func parseSnailNode(l string, curIndex *int, parent *node) *node {
	newNode := &node{}
	if parent == nil {
		newNode.depth = 0
	} else {
		newNode.depth = parent.depth + 1
	}

	r := rune(l[*curIndex])
	*curIndex += 1
	if r == '[' {
		newNode.kind = pair
		newNode.parent = parent
		newNode.left = parseSnailNode(l, curIndex, newNode)
		newNode.right = parseSnailNode(l, curIndex, newNode)
		// Skip over closing bracket character.
		*curIndex += 1
		return newNode
	} else {
		n, _ := strconv.Atoi(string(r))
		newNode.kind = literal
		newNode.parent = parent
		newNode.value = n
		// Skip over comma character.
		*curIndex += 1
	}
	return newNode
}

func parseInput() []*node {
	lines := util.ReadFileLines("input.txt")
	snailNums := []*node{}
	for _, l := range lines {
		i := 0
		snailNums = append(snailNums, parseSnailNode(l, &i, nil))
	}
	return snailNums
}

func part1() int {
	snailNums := parseInput()
	n := snailNums[0]
	for i := 1; i < len(snailNums); i += 1 {
		n.reduce()
		n = n.add(snailNums[i])
	}
	n.reduce()
	return n.magnitude()
}

func part2() int {
	snailNums := parseInput()
	max := 0
	for i := 0; i < len(snailNums); i += 1 {
		for j := 0; j < len(snailNums); j += 1 {
			n := snailNums[i].deepCopy().add(snailNums[j].deepCopy())
			n.reduce()
			magnitude := n.magnitude()
			if magnitude > max {
				max = magnitude
			}
		}
	}
	return max
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
