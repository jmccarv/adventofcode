package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

// Elves carry rucksacks with two compartements
// Each compartment holds one or more items
// Each item has a type, denoted by a letter a-z or A-Z
// Each type has a priority assigned to it (see below)
// Items are split evenly between compartments

type rucksack [2]uint64

// Each line of input represents the items in one Elf's rucksack
// First half of the line are the items in the first compartment
func main() {
	var sacks []rucksack

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var sack rucksack

		line := s.Text()
		for _, t := range line[:len(line)/2] {
			sack[0] |= 1 << (int(t) - 'A')
		}
		for _, t := range line[len(line)/2:] {
			sack[1] |= 1 << (int(t) - 'A')
		}
		sacks = append(sacks, sack)
	}
	part1(sacks)
	part2(sacks)
}

// Priorities:
//   a-z => 1-26
//   A-Z => 27-52
func priority(r int) int {
	if r >= 'a' {
		return r - 'a' + 1
	}
	return r - 'A' + 26 + 1
}

func score(c uint64) int {
	return priority('A' + bits.TrailingZeros64(c))
}

func part1(sacks []rucksack) {
	sum := 0
	for _, sack := range sacks {
		sum += score(sack[0] & sack[1])
	}
	fmt.Println(sum)
}

func part2(sacks []rucksack) {
	sum := 0
	for gidx := 0; gidx < len(sacks); gidx += 3 {
		var c uint64 = 0xffffffffffffffff
		for _, sack := range sacks[gidx : gidx+3] {
			c &= sack[0] | sack[1]
		}
		sum += score(c)
	}
	fmt.Println(sum)
}
