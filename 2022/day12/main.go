package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
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
var p2start []int
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
			case 'a':
				p2start = append(p2start, nodeNr)
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

	g := genGraph(nodes)

	part1(g)
	part2(g)
	fmt.Println("Total Runtime", time.Now().Sub(t0))
}

func genGraph(nodes []node) *graph.Mutable {
	g := graph.New(len(nodes) * nrCols)

	// We can move from a node to another where the second
	// node's height is at most one more than ours
	add := func(a, b node) {
		if a.h >= b.h-1 {
			g.AddCost(a.nr, b.nr, 1)
		}
		if b.h >= a.h-1 {
			g.AddCost(b.nr, a.nr, 1)
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
	t0 := time.Now()
	path, _ := graph.ShortestPath(g, start, end)
	fmt.Println("Part 1:", len(path)-1, " in", time.Now().Sub(t0))
}

func part2(g *graph.Mutable) {
	t0 := time.Now()
	ch := make(chan int)
	workers := make(chan struct{}, runtime.NumCPU())
	var res int = math.MaxInt

	go func() {
		for _, s := range p2start {
			workers <- struct{}{}
			go func(s int) {
				var res int = math.MaxInt
				path, _ := graph.ShortestPath(g, s, end)
				if len(path) > 0 && len(path)-1 < res {
					res = len(path) - 1
				}
				ch <- res
				<-workers
			}(s)
		}
	}()

	for i := 0; i < len(p2start); i++ {
		r := <-ch
		if r < res {
			res = r
		}
	}
	fmt.Println("Part 2:", res, " in", time.Now().Sub(t0))
}
