package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	U = 0
	D = 1
	L = 2
	R = 3
)

type loc struct {
	r, c int
}
type dir struct {
	r, c int // intended to be used as a direction, i.e. -1 <= val <= 1
}
type slope struct {
	pn   int // positive/negative
	from loc // location of tree this slope ends at
}
type tree struct {
	loc
	h   int //height
	vis bool
	//s [4]slope
}

type grove struct {
	t    [][]tree
	size int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var trees [][]tree

	r := 0
	for s.Scan() {
		row := make([]tree, len(s.Text()))
		for c, h := range []byte(s.Text()) {
			row[c] = tree{loc: loc{r: r, c: c}, h: int(h - '0')}
		}
		trees = append(trees, row)
		r++
	}
	g := grove{t: trees, size: len(trees[0])}
	g.dump()
	part1(g)
}

func part1(g grove) {
	count := func(g grove) {
		for c := 1; c < g.size-1; c++ {
			h := g.t[0][c]
			fmt.Println("h ", h)

			for r := 1; r < g.size-1; r++ {
				fmt.Println(h.h, g.t[r][c])
				if g.t[r][c].h > h.h {
					g.t[r][c].vis = true
					h = g.t[r][c]
				}
			}

			h = g.t[g.size-1][c]
			fmt.Println("h ", h)
			for r := g.size - 1; r > 0; r-- {
				fmt.Println(h.h, g.t[r][c])
				if g.t[r][c].h > h.h {
					g.t[r][c].vis = true
					h = g.t[r][c]
				}
			}
		}

		fmt.Println()
		for r := 1; r < g.size-1; r++ {
			h := g.t[r][0]
			fmt.Println("h ", h)

			for c := 1; c < g.size-1; c++ {
				fmt.Println(h.h, g.t[r][c])
				if g.t[r][c].h > h.h {
					g.t[r][c].vis = true
					h = g.t[r][c]
				}
			}

			h = g.t[r][g.size-1]
			fmt.Println("h ", h)
			for c := g.size - 1; c > 0; c-- {
				fmt.Println(h.h, g.t[r][c])
				if g.t[r][c].h > h.h {
					g.t[r][c].vis = true
					h = g.t[r][c]
				}
			}
		}
	}

	count(g)
	vis := 0
	for r := 1; r < g.size-1; r++ {
		for c := 1; c < g.size-1; c++ {
			if g.t[r][c].vis {
				vis++
			}
		}
	}
	fmt.Println(vis + (g.size-1)*4)
}

/*
func (g grove) visibileFrom(to tree, from loc) bool {
	dir := to.dirFrom(from)

	for ; !to.sameLoc(from); from = from.add(dir) {
		fmt.Println("from ", from, "  to ", to)
	}

	return true
}

func (a loc) dirFrom(b loc) loc {
	var dir loc
	if a.r < b.r {
		dir.r = -1
	} else if a.r > b.r {
		dir.r = 1
	}

	if a.c < b.c {
		dir.c = -1
	} else if a.c > b.c {
		dir.c = 1
	}
	return dir
}

func (a loc) sameLoc(b loc) bool {
	if a.r == b.r && a.c == b.c {
		return true
	}
	return false
}

func (a loc) add(b loc) loc {
	a.r += b.r
	a.c += b.c
	return a
}
*/

func (g grove) dump() {
	for _, row := range g.t {
		for _, t := range row {
			fmt.Printf("%d", t.h)
		}
		fmt.Println()
	}
}
