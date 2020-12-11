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

	fmt.Println("part1:", part1(g))
}

func part1(g grid) int {
	var m bool
	for {
		g, m = g.mutate()
		if !m {
			return g.occupied()
		}
	}
	return 0
}

// Returns the new grid and whether it differs from the old grid
func (g grid) mutate() (grid, bool) {
	mutated := false
	new := g.copy()

	for r := 1; r < len(g)-1; r++ {
		for c := 1; c < len(g[r])-1; c++ {
			x := g[r][c]
			o := g.occupiedNeighbors(r, c)

			switch x {
			case '.', ' ':
				break
			case 'L':
				if o == 0 {
					new[r][c] = '#'
					mutated = true
				}
			case '#':
				if o >= 4 {
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

func (g grid) occupiedNeighbors(rx, cx int) int {
	ret := 0
	for r := rx - 1; r < rx+2; r++ {
		for c := cx - 1; c < cx+2; c++ {
			if r == rx && c == cx {
				continue
			}
			if g[r][c] == '#' {
				ret++
			}
		}
	}
	return ret
}

func (g grid) occupied() int {
	count := 0
	for r := 1; r < len(g)-1; r++ {
		count += bytes.Count(g[r], []byte{'#'})
	}
	return count
}
