package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourbasic/graph"
)

type location struct {
	r, c int
}

type node struct {
	nr int
	l  location
	h  int //height
}

var start, end location
var nrRows, nrCols int
var nodes [][]node

func main() {
	s := bufio.NewScanner(os.Stdin)
	nodeNr := 0
	for s.Scan() {
		r := make([]node, len(s.Text()))
		for c := 0; c < len(s.Text()); c++ {
			loc := location{r: len(nodes), c: c}
			h := s.Text()[c]
			switch h {
			case 'S':
				start = loc
				h = 'a'
			case 'E':
				end = loc
				h = 'z'
			}
			r[c] = node{h: int(h), l: loc, nr: nodeNr}
			nodeNr++
		}
		nodes = append(nodes, r)
	}
	nrRows = len(nodes)
	nrCols = len(nodes[0])

	part1()
}

func part1() {
	g := graph.New(nrRows * nrCols)

	// We can move from a node to another where the second
	// node's height is at most one more than ours
	add := func(a, b node) {
		if a.h >= b.h-1 {
			g.AddCost(a.nr, b.nr, int64(a.h-b.h+1))
		}
		if b.h >= a.h-1 {
			g.AddCost(b.nr, a.nr, int64(b.h-a.h+1))
		}
	}

	for r, cols := range nodes {
		for c, this := range cols {

			if c > 0 {
				add(nodes[r][c-1], this)
			}

			if r > 0 {
				add(nodes[r-1][c], this)
			}

		}
	}

	//fmt.Println(g)

	path, _ := graph.ShortestPath(g, nodes[start.r][start.c].nr, nodes[end.r][end.c].nr)
	//fmt.Println(cost)
	fmt.Println(len(path) - 1)
}
