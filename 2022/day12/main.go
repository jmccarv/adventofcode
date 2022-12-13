package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

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

var start, end, nrCols int
var nodes []node

func main() {
	t0 := time.Now()
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

	g := genGraph(nodes)

	solve(g)
	fmt.Println("Total Runtime", time.Now().Sub(t0))
}

func genGraph(nodes []node) *graph.Mutable {
	g := graph.New(len(nodes) * nrCols)

	// We can move from a node to another where the second
	// node's height is at most one more than ours
	// But we're going to be searching from the end to the beginning so reverse the edges
	add := func(a, b node) {
		if a.h >= b.h-1 {
			g.AddCost(b.nr, a.nr, 0)
		}
		if b.h >= a.h-1 {
			g.AddCost(a.nr, b.nr, 0)
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

func solve(g *graph.Mutable) {
	// We do a breadth first search starting at the end
	// The first 'a' we come to is the shortest path for part 2
	// When we find the start node, we've got the length for part 1
	var p1, p2 int

	// distance traveled to get to any node from the end node
	dist := make([]int, g.Order())

	// v is the node coming from and w is the node we're going to
	graph.BFS(g, end, func(v, w int, _ int64) {
		dist[w] = dist[v] + 1
		if nodes[w].h == 'a' && p2 == 0 {
			// This is the shortest route from any 'a' to the end node
			p2 = dist[w]
		}
		if w == start {
			p1 = dist[w]
		}
	})

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
