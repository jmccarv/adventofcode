package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yourbasic/graph"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	costs := make(map[int]int)
	var pr []int

	g := graph.New(100 * 100)
	nodeNr := 0
	for s.Scan() {
		r := make([]int, len(s.Text()))
		//fmt.Println("new row, nodeNr:", nodeNr)
		for c := 0; c < len(s.Text()); c++ {
			r[c], _ = strconv.Atoi(s.Text()[c : c+1])
			above := nodeNr - len(r)

			if c > 0 {
				g.AddCost(nodeNr-1, nodeNr, int64(r[c]))
				g.AddCost(nodeNr, nodeNr-1, int64(r[c-1]))
				//g.AddBothCost(nodeNr-1, nodeNr, int64(r[c]))
			}
			if len(pr) > 0 {
				g.AddCost(above, nodeNr, int64(r[c]))
				g.AddCost(nodeNr, above, int64(pr[c]))
			}
			costs[nodeNr] = r[c]
			nodeNr++
		}
		pr = r
	}
	nodeNr--

	//fmt.Println(costs)
	//path, cost := graph.ShortestPath(g, 0, nodeNr)
	_, cost := graph.ShortestPath(g, 0, nodeNr)
	fmt.Println(cost)
	//fmt.Println(path)
	/*
		c := 0
		for i := 1; i < len(path); i++ {
			c += costs[path[i]]
			fmt.Println(i, path[i], costs[path[i]], c)
		}
		fmt.Println(c)
	*/

	//fmt.Println(graph.ShortestPath(g, 0, nodeNr))

	//fmt.Println(g)

}
