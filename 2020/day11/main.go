package main

import (
	"bufio"
	"fmt"
	//"log"
	"bytes"
	"os"
)

// strings are read only and relatively expensive to modify
// since we'll be modifying the grid quite a bit I'm going to
// use byte slices instead of strings.
type grid [][]byte

type location struct {
	r int
	c int
}

var directions = []location{
	location{-1, -1},
	location{0, -1},
	location{1, -1},
	location{-1, 0},
	location{1, 0},
	location{-1, 1},
	location{0, 1},
	location{1, 1},
}

func main() {
	// Start with a completely empty line at the top of the grid
	var g grid = [][]byte{[]byte{}}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		g = append(g, []byte(" "+s.Text()+" "))
	}
	l := len(g[1])
	g[0] = bytes.Repeat([]byte{' '}, l+1)
	g = append(g, bytes.Repeat([]byte{' '}, l+1))

	fmt.Println("part1:", partx(g, 4, func(g grid, l location) int { return g.occupiedNeighbors(l) }))
	fmt.Println("part2:", partx(g, 5, func(g grid, l location) int { return g.occupiedLineOfSight(l) }))
}

func partx(g grid, threshold int, occupancy func(g grid, l location) int) int {
	var m bool
	for {
		g, m = g.mutate(threshold, occupancy)
		if !m {
			return g.occupied()
		}
	}
	return 0
}

// Returns the new grid and whether it differs from the old grid
func (g grid) mutate(maxOccupied int, occupancy func(g grid, l location) int) (grid, bool) {
	mutated := false
	new := g.copy()

	for r := 1; r < len(g)-1; r++ {
		for c := 1; c < len(g[r])-1; c++ {
			l := location{r, c}
			x := g.at(l)

			if x == '.' {
				continue
			}

			o := occupancy(g, l)

			switch x {
			case 'L':
				if o == 0 {
					new[r][c] = '#'
					mutated = true
				}
			case '#':
				if o >= maxOccupied {
					new[r][c] = 'L'
					mutated = true
				}
			}
		}
	}

	return new, mutated
}

func (g grid) copy() grid {
	var n grid
	for _, r := range g {
		n = append(n, append([]byte{}, r...))
	}
	return n
}

func (g grid) occupiedNeighbors(l location) int {
	var ret int
	for _, d := range directions {
		if g.at(l.add(d)) == '#' {
			ret++
		}
	}
	return ret
}

func (g grid) occupiedLineOfSight(origin location) int {
	var ret int

	for _, d := range directions {
		l := origin
	out:
		for g.validLocation(l) {
			l = l.add(d)
			switch g.at(l) {
			case '#':
				ret++
				break out
			case 'L':
				break out
			}
		}
	}

	return ret
}

func (g grid) at(l location) byte {
	return g[l.r][l.c]
}

func (g grid) validLocation(l location) bool {
	return g.at(l) != ' '
}

func (g grid) occupied() int {
	count := 0
	for r := 1; r < len(g)-1; r++ {
		count += bytes.Count(g[r], []byte{'#'})
	}
	return count
}

func (g grid) String() string {
	var ret string
	for r := 1; r < len(g)-1; r++ {
		for c := 1; c < len(g[r])-1; c++ {
			ret += string(g[r][c])
		}
		ret += "\n"
	}
	return ret
}

func (l location) add(o location) location {
	return location{l.r + o.r, l.c + o.c}
}
