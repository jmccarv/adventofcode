package main

import (
	"bufio"
	"fmt"
	"os"
)

// Elves carry rucksacks with two compartements
// Each compartment holds one or more items
// Each item has a type, denoted by a letter a-z or A-Z
// Each type has a priority assigned to it (see below)
// Items are split evenly between compartments

type compartment [53]bool // compartment index is the type's priority
type rucksack [2]compartment

// Priorities:
//   a-z => 1-26
//   A-Z => 27-52
func priority(r rune) int {
	if r >= 'a' {
		return int(r) - 'a' + 1
	}
	return int(r) - 'A' + 26 + 1
}

// Each line of input represents the items in one Elf's rucksack
// First half of the line are the items in the first compartment
func main() {
	var sacks []rucksack

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var sack rucksack

		line := s.Text()
		for _, t := range line[:len(line)/2] {
			sack[0][priority(t)] = true
		}
		for _, t := range line[len(line)/2:] {
			sack[1][priority(t)] = true
		}
		sacks = append(sacks, sack)
	}
	part1(sacks)
	part2(sacks)
}

func part1(sacks []rucksack) {
	sum := 0
	for _, sack := range sacks {
		for i := 1; i <= 52; i++ {
			if sack[0][i] && sack[1][i] {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}

func part2(sacks []rucksack) {
	sum := 0
	for gidx := 0; gidx < len(sacks); gidx += 3 {
		var types [53]int
		for _, sack := range sacks[gidx : gidx+3] {
			for i := 1; i <= 52; i++ {
				if sack[0][i] || sack[1][i] {
					types[i]++
				}
			}
		}
		for i := 1; i <= 52; i++ {
			if types[i] == 3 {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
