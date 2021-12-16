// Solution for Advent of Code 2021 -- Day 16
// https://adventofcode.com/2021/day/16

package main

import (
	"encoding/hex"
	"fmt"
	"math"

	"github.com/huderlem/adventofcode2021/util"
)

type bitReader struct {
	pos   uint64
	bytes []byte
}

func newBitReader(bytes []byte) bitReader {
	return bitReader{
		pos:   0,
		bytes: bytes,
	}
}

func (b *bitReader) next() int {
	byteIndex := b.pos / 8
	bitIndex := b.pos % 8
	masked := b.bytes[byteIndex] & (1 << (7 - bitIndex))
	b.pos += 1
	if masked == 0 {
		return 0
	}
	return 1
}

func (b *bitReader) readInt(numBits int) uint64 {
	var val uint64
	for i := 0; i < numBits; i += 1 {
		val <<= 1
		if b.next() != 0 {
			val |= 1
		}
	}
	return val
}

func (b *bitReader) readLiteralValue() uint64 {
	var value uint64
	for {
		chunk := b.readInt(5)
		value <<= 4
		value |= chunk & 15
		if chunk>>4 == 0 {
			return value
		}
	}
}

func parseInput() []byte {
	input := util.ReadFileString("input.txt")
	result, _ := hex.DecodeString(input)
	return result
}

func (b *bitReader) parsePacket(versionSum *int) uint64 {
	version := b.readInt(3)
	*versionSum += int(version)
	typeID := b.readInt(3)
	if typeID == 4 {
		literal := b.readLiteralValue()
		return literal
	}
	lengthType := b.next()
	values := []uint64{}
	if lengthType == 0 {
		length := b.readInt(15)
		startPos := b.pos
		for b.pos != startPos+length {
			values = append(values, b.parsePacket(versionSum))
		}
	} else {
		numSubPackets := b.readInt(11)
		for i := 0; i < int(numSubPackets); i += 1 {
			values = append(values, b.parsePacket(versionSum))
		}
	}

	var result uint64
	switch typeID {
	case 0:
		for _, v := range values {
			result += v
		}
	case 1:
		result = 1
		for _, v := range values {
			result *= v
		}
	case 2:
		result = math.MaxUint64
		for _, v := range values {
			if v < result {
				result = v
			}
		}
	case 3:
		result = 0
		for _, v := range values {
			if v > result {
				result = v
			}
		}
	case 5:
		if values[0] > values[1] {
			result = 1
		}
	case 6:
		if values[0] < values[1] {
			result = 1
		}
	case 7:
		if values[0] == values[1] {
			result = 1
		}
	}
	return result
}

func part1() int {
	bytes := parseInput()
	reader := newBitReader(bytes)
	versionSum := 0
	reader.parsePacket(&versionSum)
	return versionSum
}

func part2() uint64 {
	bytes := parseInput()
	reader := newBitReader(bytes)
	versionSum := 0
	return reader.parsePacket(&versionSum)
}

func main() {
	fmt.Println("Part 1 Answer:", part1())
	fmt.Println("Part 2 Answer:", part2())
}
