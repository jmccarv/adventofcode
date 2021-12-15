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
	var pr []int
	node := -1
	for _, r := range input {
		for c, _ := range r {
			node++
			above := node - len(r)

			if c > 0 {
				g.AddCost(node-1, node, int64(r[c]))
				g.AddCost(node, node-1, int64(r[c-1]))
			}
			if above >= 0 {
				g.AddCost(above, node, int64(r[c]))
				g.AddCost(node, above, int64(pr[c]))
			}
		}
		pr = r
	}

	_, cost := graph.ShortestPath(g, 0, node)
	fmt.Println(cost)
}

func part1(input [][]int) {
	shortestPath(input)
}

func part2(input [][]int) {
	size := len(input)

	// Extend to the right 4x
	for r, _ := range input {
		for dc := 1; dc < 5; dc++ {
			for c := 0; c < size; c++ {
				v := input[r][c+(dc-1)*size]%9 + 1
				input[r] = append(input[r], v)
			}
		}
	}

	// Extend down 4x
	for dr := 1; dr < 5; dr++ {
		for r := 0; r < size; r++ {
			var nr []int
			for c, _ := range input[r] {
				v := input[r+(dr-1)*size][c]%9 + 1
				nr = append(nr, v)
			}
			input = append(input, nr)
		}
	}

	shortestPath(input)
}
