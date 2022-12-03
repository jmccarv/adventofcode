package main

import (
	"bufio"
	"fmt"
	"os"
)

type compartment [52]bool
type rucksack [2]compartment

func idx(t byte) byte {
	if t > 'Z' {
		return t - byte('a')
	}
	return t - byte('A') + 26
}

func main() {
	var sacks []rucksack

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		var sack rucksack
		for _, t := range []byte(line[:len(line)/2]) {
			sack[0][idx(t)] = true
		}
		for _, t := range []byte(line[len(line)/2:]) {
			sack[1][idx(t)] = true
		}
		sacks = append(sacks, sack)
	}
	part1(sacks)
	part2(sacks)
}

func part1(sacks []rucksack) {
	sum := 0
	for _, sack := range sacks {
		for i := 0; i < 52; i++ {
			if sack[0][i] && sack[1][i] {
				sum += i + 1
			}
		}
	}
	fmt.Println(sum)
}

func part2(sacks []rucksack) {
	sum := 0
	for gidx := 0; gidx < len(sacks); gidx += 3 {
		var types [52]int
		for _, sack := range sacks[gidx : gidx+3] {
			for i := 0; i < 52; i++ {
				if sack[0][i] || sack[1][i] {
					types[i]++
				}
			}
		}
		for i := 0; i < 52; i++ {
			if types[i] == 3 {
				sum += i + 1
			}
		}
	}
	fmt.Println(sum)
}
