package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	n string
	l string
	r string
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	directions := s.Text()
	s.Scan()

	nodes := make(map[string]node)
	for s.Scan() {
		var loc, l, r string
		fmt.Sscanf(s.Text(), "%s = (%s %s)", &loc, &l, &r)
		nodes[loc] = node{loc, l[0:3], r[0:3]}
	}
	fmt.Printf("%#v\n", nodes)

	p1 := 0
	i := 0
	node := nodes["AAA"]
	fmt.Printf("Start: %#v\n", node)
	for {
		p1++
		fmt.Printf("%s %c", node.n, directions[i])
		switch directions[i] {
		case 'L':
			fmt.Printf(" -> %s", node.l)
			node = nodes[node.l]
		case 'R':
			fmt.Printf(" -> %s", node.r)
			node = nodes[node.r]
		}
		fmt.Printf(" = %#v\n", node)
		if node.n == "ZZZ" {
			break
		}
		i = (i + 1) % len(directions)
	}
	println("Part1", p1)
}
