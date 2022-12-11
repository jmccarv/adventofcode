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

// Each line of input represents the items in one Elf's rucksack
// First half of the line are the items in the first compartment
func main() {
	s := bufio.NewScanner(os.Stdin)
	nr, p1sum, p2sum := 0, 0, 0
	var c uint64 = 0xffffffffffffffff
	for s.Scan() {
		var sack [2]uint64
		nr++

		line := s.Text()
		for _, t := range line[:len(line)/2] {
			sack[0] |= 1 << (int(t) - 'A')
		}
		for _, t := range line[len(line)/2:] {
			sack[1] |= 1 << (int(t) - 'A')
		}
		// part 1
		p1sum += score(sack[0] & sack[1])

		// part 2
		c &= sack[0] | sack[1]
		if nr%3 == 0 {
			p2sum += score(c)
			c = 0xffffffffffffffff
		}
	}
	fmt.Println(p1sum)
	fmt.Println(p2sum)
}

// Priorities:
//
//	a-z => 1-26
//	A-Z => 27-52
func priority(r int) int {
	if r >= 'a' {
		return r - 'a' + 1
	}
	return r - 'A' + 26 + 1
}

func score(c uint64) int {
	return priority('A' + bits.TrailingZeros64(c))
}
