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
	nr  int
	loc location
	h   int //height
}

var start, end int
var p2start []int
var nrRows, nrCols int
var nodes []node

func main() {
	s := bufio.NewScanner(os.Stdin)
	nodeNr := 0
	row := 0
	for s.Scan() {
		nrCols = len(s.Text())
		for c := 0; c < len(s.Text()); c++ {
			h := s.Text()[c]
			switch h {
			case 'S':
				start = nodeNr
				p2start = append(p2start, nodeNr)
				h = 'a'
			case 'E':
				end = nodeNr
				h = 'z'
			}
			nodes = append(nodes, node{h: int(h), nr: nodeNr, loc: location{r: row, c: c}})
			nodeNr++
		}
		row++
	}
	nrRows = len(nodes)

	g := genGraph(nodes)

	part1(g)
}

func genGraph(nodes []node) *graph.Mutable {
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

	for i, n := range nodes {
		if n.loc.c > 0 {
			add(nodes[i-1], n)
		}

		if n.loc.r > 0 {
			add(nodes[i-nrCols], n)
		}
	}

	return g
}

func part1(g *graph.Mutable) {
	path, _ := graph.ShortestPath(g, start, end)
	fmt.Println(len(path) - 1)
}
