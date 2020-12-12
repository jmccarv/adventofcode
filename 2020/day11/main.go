package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	floor        = '.'
	emptySeat    = 'L'
	occupiedSeat = '#'
)

// strings are read only and relatively expensive to modify.
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
	// We'll eventually pad the perimiter of the grid with spaces as an
	// easy way to know when we've gone out of bounds.
	var g grid = [][]byte{[]byte{}}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		g = append(g, []byte(" "+s.Text()+" "))
	}
	l := len(g[1])
	g[0] = bytes.Repeat([]byte{' '}, l+1)
	g = append(g, bytes.Repeat([]byte{' '}, l+1))

	fmt.Println("part1:", run(g, 4, func(g grid, l location) int { return g.occupiedNeighbors(l) }))
	fmt.Println("part2:", run(g, 5, func(g grid, l location) int { return g.occupiedLineOfSight(l) }))
}

func run(g grid, threshold int, occupancy func(g grid, l location) int) int {
	var m bool
	for {
		g, m = g.mutate(threshold, occupancy)
		if !m {
			return g.occupied()
		}
	}
	return 0
}

// The callback is used to count the number of occupied seats visible from a given location
// An ooccupied seat that can see n occupied seats where n > threshold will be emptied
//
// Returns the new grid and whether it differs from the old grid
func (g grid) mutate(threshold int, occupancy func(g grid, l location) int) (grid, bool) {
	mutated := false
	new := g.copy()

	for r := 1; r < len(g)-1; r++ {
		for c := 1; c < len(g[r])-1; c++ {
			l := location{r, c}
			x := g.at(l)

			if x == floor {
				// floor -- ignore, no need to count occupancy
				continue
			}

			// find the number of occupied seats visible from this location
			o := occupancy(g, l)

			switch x {
			case emptySeat:
				// empty seats are filled if there are no visible occupied seats
				if o == 0 {
					new.set(l, occupiedSeat)
					mutated = true
				}
			case occupiedSeat:
				// occupied seats are emptied if the number of visible occupied
				// seats meets or exceeds the threshold
				if o >= threshold {
					new.set(l, emptySeat)
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

// Return the number of adjacent seats that are occupied
func (g grid) occupiedNeighbors(l location) int {
	var ret int
	for _, d := range directions {
		if g.at(l.add(d)) == occupiedSeat {
			ret++
		}
	}
	return ret
}

// Used for part 2 of the puzzle -- looks in each direction for the
// first visible seat and returns the number of those that are occupied.
func (g grid) occupiedLineOfSight(origin location) int {
	var ret int

	for _, d := range directions {
		l := origin
	out:
		for g.validLocation(l) {
			l = l.add(d)
			switch g.at(l) {
			case occupiedSeat:
				ret++
				break out
			case emptySeat:
				break out
			}
		}
	}
	return ret
}

func (g grid) at(l location) byte {
	return g[l.r][l.c]
}

func (g grid) set(l location, to byte) {
	g[l.r][l.c] = to
}

func (g grid) validLocation(l location) bool {
	return g.at(l) != ' '
}

// returns the number of occupied seats in the grid.
func (g grid) occupied() int {
	count := 0
	for r := 1; r < len(g)-1; r++ {
		count += bytes.Count(g[r], []byte{occupiedSeat})
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
