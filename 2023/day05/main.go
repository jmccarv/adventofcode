package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// A range is a half-open interval (like a slice in Go) which includes A, but not B
type Range struct {
	A, B int
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.A, r.B)
}

func (r Range) Len() int {
	return r.B - r.A
}

func (r1 Range) Overlaps(r2 Range) bool {
	if r1.A >= r2.A && r1.A < r2.B {
		return true
	}
	if r1.B > r2.A && r1.B <= r2.B {
		return true
	}
	return false
}

func (r1 Range) Contains(r2 Range) bool {
	if r1.A <= r2.A && r1.B >= r2.B {
		return true
	}
	return false
}

type rangeMap struct {
	Range
	destStart int
}

func (rm rangeMap) String() string {
	return fmt.Sprintf("%s->%d", rm.Range, rm.destStart)
}

var levels [][]rangeMap

func main() {
	var seeds []int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	for _, nr := range strings.Fields(s.Text())[1:] {
		x, _ := strconv.Atoi(nr)
		seeds = append(seeds, x)
	}

	var src, dest, length int
	var rMaps []rangeMap
	for s.Scan() {
		n, err := fmt.Sscan(s.Text(), &dest, &src, &length)
		if err == nil && n == 3 {
			rMaps = append(rMaps, rangeMap{Range{src, src + length}, dest})
		} else {
			if len(rMaps) > 0 {
				levels = append(levels, rMaps)
				rMaps = []rangeMap{}
			}
		}

	}
	if len(rMaps) > 0 {
		levels = append(levels, rMaps)
	}

	// part 1
	p1 := math.MaxInt
	for _, seed := range seeds {
		//fmt.Println("seed", seed, mapSeed(seed, levels))
		p1 = min(p1, mapSeed(seed, levels))
	}

	p2 := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		//fmt.Println("seed range", seeds[i], seeds[i+1])
		p2 = min(p2, mapRange(0, Range{seeds[i], seeds[i] + seeds[i+1]}))
	}
	//fmt.Println(maps)
	fmt.Println("Part 1", p1)
	fmt.Println("Part 2", p2)
}

func lookup(maps []rangeMap, src int) int {
	for _, m := range maps {
		//fmt.Println("lookup", src, i, start, rm.destStart[i], rm.length[i])
		if src >= m.A && src < m.B {
			return m.destStart + src - m.A
		}
	}
	return src
}

func mapSeed(src int, levels [][]rangeMap) int {
	for _, maps := range levels {
		src = lookup(maps, src)
	}
	return src
}

func mapRange(lvl int, src Range) int {
	for _, rm := range levels[lvl] {
		if rm.Range.Contains(src) {
			ofs := rm.destStart - rm.A
			s := Range{src.A + ofs, src.B + ofs}
			if lvl >= len(levels)-1 {
				return s.A
			}
			return mapRange(lvl+1, s)
		}
		if rm.A <= src.A && src.A < rm.B {
			// split the range into two
			return min(mapRange(lvl, Range{src.A, rm.B}), mapRange(lvl, Range{rm.B, src.B}))
		}
		if src.B > rm.A && src.B <= rm.B {
			// split
			return min(mapRange(lvl, Range{src.A, rm.A}), mapRange(lvl, Range{rm.A, src.B}))
		}
	}
	if lvl >= len(levels)-1 {
		return src.A
	}
	return mapRange(lvl+1, src)
}
