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
	var input [][]int
	for s.Scan() {
		r := make([]int, len(s.Text()))
		for c := 0; c < len(s.Text()); c++ {
			r[c], _ = strconv.Atoi(s.Text()[c : c+1])
		}
		input = append(input, r)
	}

	part1(input)
	part2(input)
}

func shortestPath(input [][]int) {
	g := graph.New(len(input) * len(input))
	costs := make(map[int]int)
	var pr []int
	nodeNr := 0
	for _, r := range input {
		//fmt.Println("new row, nodeNr:", nodeNr)
		for c, _ := range r {
			above := nodeNr - len(r)

			if c > 0 {
				g.AddCost(nodeNr-1, nodeNr, int64(r[c]))
				g.AddCost(nodeNr, nodeNr-1, int64(r[c-1]))
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

	_, cost := graph.ShortestPath(g, 0, nodeNr)
	fmt.Println(cost)
}

func part1(input [][]int) {
	shortestPath(input)
}

func part2(input [][]int) {
	size := len(input)
	for r, _ := range input {
		for dc := 1; dc < 5; dc++ {
			for c := 0; c < size; c++ {
				v := input[r][c+(dc-1)*size] + 1
				if v > 9 {
					v = 1
				}
				input[r] = append(input[r], v)
			}
		}
	}

	for dr := 1; dr < 5; dr++ {
		for r := 0; r < size; r++ {
			var nr []int
			for c, _ := range input[r] {
				v := input[r+(dr-1)*size][c] + 1
				if v > 9 {
					v = 1
				}
				nr = append(nr, v)
			}
			input = append(input, nr)
		}
	}

	shortestPath(input)
}
