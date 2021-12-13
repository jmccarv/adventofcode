package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type path []*node

type node struct {
	name        string
	isSmallCave bool
}

type graph struct {
	nodes map[string]*node
	edges map[node][]*node
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	g := newGraph()
	for s.Scan() {
		x := strings.Split(s.Text(), "-")
		if len(x) != 2 {
			panic("Invalid input")
		}
		g.addEdge(x[0], x[1])
	}
	fmt.Println(g)
	part1(g)
}

func part1(g graph) {
	var descend func(*node, path)

	var found []path

	descend = func(from *node, p path) {
		if from.isSmallCave && p.hasNode(*from) {
			return
		}

		p = append(p, from)

		if from.name == "end" {
			found = append(found, p)
			return
		}

		for _, x := range g.edges[*from] {
			descend(x, p.copy())
		}
	}

	descend(g.nodes["start"], path{})

	for _, p := range found {
		fmt.Println(p)
	}
	fmt.Println(len(found))
}

func (p path) hasNode(n node) bool {
	for _, x := range p {
		if x.name == n.name {
			return true
		}
	}
	return false
}

func (p path) copy() path {
	x := make([]*node, len(p))
	copy(x, p)
	return x
}

func (p path) String() string {
	ret := ""
	for _, n := range p {
		ret += "," + n.name
	}
	if len(ret) > 0 {
		ret = ret[1:]
	}
	return ret
}

func newGraph() graph {
	return graph{
		nodes: make(map[string]*node),
		edges: make(map[node][]*node),
	}
}

func (g *graph) addEdge(v, w string) {
	n1, ok := g.nodes[v]
	if !ok {
		n1 = newNode(v)
		g.nodes[v] = n1
	}

	n2, ok := g.nodes[w]
	if !ok {
		n2 = newNode(w)
		g.nodes[w] = n2
	}

	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func newNode(name string) *node {
	n := node{name: name, isSmallCave: name == strings.ToLower(name)}
	return &n
}

func (n *node) String() string {
	f := "*"
	if n.isSmallCave {
		f = "_"
	}
	return fmt.Sprintf("%s%s", f, n.name)
}

func (g graph) String() string {
	ret := ""
	seen := make(map[node]bool)

	for _, n := range g.nodes {
		if _, ok := seen[*n]; ok {
			continue
		}
		seen[*n] = true

		for _, x := range g.edges[*n] {
			if _, ok := seen[*x]; ok {
				continue
			}
			if len(ret) > 0 {
				ret += "\n"
			}
			ret += fmt.Sprintf("%s=%s", n, x)
		}
	}
	return ret
}
