// Solution for Advent of Code 2021 -- Day 4
// https://adventofcode.com/2021/day/4

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/huderlem/adventofcode2021/util"
)

type bingoSquare struct {
	value  int
	marked bool
}

type pos struct {
	x, y int
}

type bingoBoard struct {
	layout          [][]bingoSquare
	numberLocations map[int]pos
	won             bool
	id              int
}

func (b *bingoBoard) tryMarkNumber(n int) bool {
	if pos, ok := b.numberLocations[n]; ok {
		b.layout[pos.y][pos.x].marked = true
		return true
	}
	return false
}

func (b *bingoBoard) checkWin() bool {
	// Check row-wise victory.
	for y := 0; y < len(b.layout); y += 1 {
		won := true
		for x := 0; x < len(b.layout); x += 1 {
			won = won && b.layout[y][x].marked
		}
		if won {
			return true
		}
	}

	// Check column-wise victory.
	for x := 0; x < len(b.layout); x += 1 {
		won := true
		for y := 0; y < len(b.layout); y += 1 {
			won = won && b.layout[y][x].marked
		}
		if won {
			return true
		}
	}

	return false
}

func parseInput() ([]int, []bingoBoard) {
	var inputLines = util.ReadFileChunks("input.txt")
	numbers := []int{}
	for _, n := range strings.Split(inputLines[0][0], ",") {
		num, _ := strconv.Atoi(n)
		numbers = append(numbers, num)
	}

	boards := []bingoBoard{}
	for i, boardDef := range inputLines[1:] {
		layout := [][]bingoSquare{}
		locations := map[int]pos{}
		for y, row := range boardDef {
			nums := strings.Fields(row)
			curRow := []bingoSquare{}
			for x, n := range nums {
				num, _ := strconv.Atoi(n)
				curRow = append(curRow, bingoSquare{num, false})
				locations[num] = pos{x, y}
			}
			layout = append(layout, curRow)
		}
		boards = append(boards, bingoBoard{layout, locations, false, i})
	}
	return numbers, boards
}

func markBoards(boards []bingoBoard, calledNumber int) {
	for _, board := range boards {
		if !board.won {
			board.tryMarkNumber(calledNumber)
		}
	}
}

func checkNewBoardWins(boards []bingoBoard) []bingoBoard {
	newlyWinningBoards := []bingoBoard{}
	for i, board := range boards {
		if !board.won && board.checkWin() {
			boards[i].won = true
			newlyWinningBoards = append(newlyWinningBoards, board)
		}
	}
	return newlyWinningBoards
}

func sumUnmarkedNumbers(board bingoBoard) int {
	sum := 0
	for y := 0; y < len(board.layout); y += 1 {
		for x := 0; x < len(board.layout); x += 1 {
			if !board.layout[y][x].marked {
				sum += board.layout[y][x].value
			}
		}
	}
	return sum
}

func part1() int {
	var calledNumbers, boards = parseInput()
	for _, calledNumber := range calledNumbers {
		markBoards(boards, calledNumber)
		newBoardWins := checkNewBoardWins(boards)
		if len(newBoardWins) > 0 {
			sum := sumUnmarkedNumbers(newBoardWins[0])
			return sum * calledNumber
		}
	}
	return -1
}

func part2() int {
	var calledNumbers, boards = parseInput()
	numWinningBoards := 0
	for _, calledNumber := range calledNumbers {
		markBoards(boards, calledNumber)
		newBoardWins := checkNewBoardWins(boards)
		if len(newBoardWins) > 0 {
			numWinningBoards += len(newBoardWins)
			if numWinningBoards == len(boards) {
				sum := sumUnmarkedNumbers(newBoardWins[len(newBoardWins)-1])
				return sum * calledNumber
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
