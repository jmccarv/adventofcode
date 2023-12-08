package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	loc string
	l   *node
	r   *node
}

var directions string

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	directions = s.Text()
	s.Scan()

	nodes := make(map[string]*node)
	for s.Scan() {
		var loc, l, r string
		fmt.Sscanf(s.Text(), "%s = (%s %s)", &loc, &l, &r)
		l = l[0:3]
		r = r[0:3]
		ln, ok := nodes[l]
		if !ok {
			ln = &node{loc: l}
			nodes[l] = ln
		}
		rn, ok := nodes[r]
		if !ok {
			rn = &node{loc: r}
			nodes[r] = rn
		}
		n, ok := nodes[loc]
		if !ok {
			n = &node{loc: loc}
			nodes[loc] = n
		}
		n.l, n.r = ln, rn
	}

	part1(nodes)
	part2(nodes)
}

func (n *node) move(direction byte) *node {
	switch direction {
	case 'L':
		return n.l
	case 'R':
		return n.r
	}
	panic(fmt.Sprintf("Invalid direction '%c'", direction))

}

func part1(nodes map[string]*node) {
	var res, i int
	node, ok := nodes["AAA"]
	if !ok {
		fmt.Println("Invalid input for part 1; skipping")
		return
	}
	for {
		res++
		node = node.move(directions[i])
		if node.loc == "ZZZ" {
			break
		}
		i = (i + 1) % len(directions)
	}
	println("Part1", res)
}

// gcd and lcd functions stolen from the internets
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(nums ...int) int {
	if len(nums) < 2 {
		return 0
	}
	a, b := nums[0], nums[1]
	nums = nums[2:]

	ret := a * b / gcd(a, b)
	for _, i := range nums {
		ret = lcm(ret, i)
	}
	return ret
}

// While based on the description of the problem I don't see how
// the LCM of each path could be guaranteed to be the anser, it
// does work for the input set provided so ...
func part2(nodes map[string]*node) {
	var counts []int

	count := func(node *node) int {
		var cnt, i int
		for {
			cnt++
			node = node.move(directions[i])
			if node.loc[2] == 'Z' {
				break
			}
			i = (i + 1) % len(directions)
		}
		return cnt
	}

	for loc, node := range nodes {
		if loc[2] == 'A' {
			counts = append(counts, count(node))
		}
	}
	fmt.Println("Part2", lcm(counts...))
}
