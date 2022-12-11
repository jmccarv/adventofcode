package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Elves are assigned a range of sections to clean
type assignment struct {
	min, max int
}

func main() {
	nrContain := 0
	nrOverlap := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		pairs := strings.Split(s.Text(), ",")
		a1, a2 := parse(pairs[0]), parse(pairs[1])

		// part 1
		if a1.contains(a2) || a2.contains(a1) {
			nrContain++
		}

		// part 2
		if a1.overlaps(a2) {
			nrOverlap++
		}
	}
	fmt.Println(nrContain)
	fmt.Println(nrOverlap)
}

func parse(assn string) (ret assignment) {
	minmax := strings.Split(assn, "-")
	ret.min, _ = strconv.Atoi(minmax[0])
	ret.max, _ = strconv.Atoi(minmax[1])
	return
}

func (a assignment) contains(b assignment) bool {
	if a.min <= b.min && a.max >= b.max {
		return true
	}
	return false
}

func (a assignment) overlaps(b assignment) bool {
	if a.min <= b.max && a.max >= b.min {
		return true
	}
	return false
}
