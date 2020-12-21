package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nums []int
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		for _, x := range strings.Split(s.Text(), ",") {
			n, err := strconv.Atoi(x)
			if err != nil {
				log.Fatalf("Invalid number '%s' in input: %s", x, s.Text())
			}
			nums = append(nums, n)
		}
	}

	fmt.Println("part1:", part1(2020, nums))
	fmt.Println("part2:", part1(30000000, nums))
}

type ent struct {
	last int
	prev int
}

type history map[int]ent

func part1(turns int, nums []int) int {
	var spoke int
	hist := history(make(map[int]ent))

	i := 1
	for _, n := range nums {
		spoke = hist.speak(i, n)
		i++
	}
	fmt.Println(spoke, hist)

	for ; i <= turns; i++ {
		if hist[spoke].prev > 0 {
			spoke = hist.speak(i, hist[spoke].last-hist[spoke].prev)
		} else {
			spoke = hist.speak(i, 0)
		}
	}

	return spoke
}

func (h history) speak(turn, num int) int {
	if x, ok := h[num]; ok {
		x.prev = x.last
		x.last = turn
		h[num] = x
	} else {
		h[num] = ent{last: turn}
	}
	return num
}
