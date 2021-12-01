package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	var nums sort.IntSlice
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		x, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal("Invalid input", s.Text())
		}
		nums = append(nums, x)
	}
	nums = append(nums, 0)
	nums.Sort()
	nums = append(nums, nums[len(nums)-1]+3)

	dist := make(map[int]int)
	for i := 1; i < len(nums); i++ {
		dist[abs(nums[i]-nums[i-1])]++
	}
	fmt.Println(dist)
	fmt.Println("part1:", dist[1]*dist[3])

	fmt.Println("part2:", part2(nums))
	//nr := part1(preamble, nums)
	//fmt.Println("part1:", nr)
	//fmt.Println("part2:", part2(nr, nums))
}

type node int64

func (n node) ID() int64 {
	return int64(n)
}

type graph struct {
	first int64
	last  int64
	dag   *simple.DirectedGraph
}

func part2(nums []int) int {
	var prev []node
	var graphs []graph

	g := graph{first: int64(nums[0]), dag: simple.NewDirectedGraph()}

	for _, n := range nums {
		for len(prev) > 0 && int(prev[0]) < n-3 {
			prev = shift(prev)
		}
		fmt.Println(n, prev)

		if len(prev) == 0 || int(prev[len(prev)-1])+3 == n {
			if g.dag.Edges().Len() > 0 {
				graphs = append(graphs, g)
			}
			g = graph{first: int64(n), dag: simple.NewDirectedGraph()}
			// Add nodes from @prev?
			if len(prev) > 0 {
				g.first = int64(prev[0])
			}
			for _, p := range prev {
				g.dag.AddNode(p)
			}
		}

		g.dag.AddNode(node(n))
		g.last = int64(n)

		for _, p := range prev {
			g.dag.SetEdge(g.dag.NewEdge(p, node(n)))
		}

		prev = append(prev, node(n))
	}
	if g.dag.Edges().Len() > 0 {
		graphs = append(graphs, g)
	}

	fmt.Println("nr graphs:", len(graphs))

	ret := 1
	for _, g = range graphs {
		fmt.Println()
		fmt.Printf("%+v\n", g)
		fmt.Printf("%+v\n", g.dag)
		ap := path.DijkstraAllPaths(g.dag)
		//ap, _ := path.FloydWarshall(g.dag)
		p, _ := ap.AllBetween(g.first, g.last)
		fmt.Println(p)
		ret *= len(p)
	}

	return ret
}

func shift(x []node) []node {
	if len(x) > 0 {
		copy(x, x[1:])
		x = x[:len(x)-1]
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
