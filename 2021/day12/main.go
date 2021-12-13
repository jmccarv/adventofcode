package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	name        string
	isSmallCave bool
}

type graph struct {
	nodes map[string]*node
	edges map[node][]*node
}

type path struct {
	nodes []*node
	twice bool
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
	//fmt.Println(g)
	solve(g, okPart1)
	solve(g, okPart2)
}

func okPart1(n *node, p *path) bool {
	return !p.hasNode(*n)
}

func okPart2(n *node, p *path) bool {
	if !p.hasNode(*n) {
		return true
	}

	if !p.twice && n.name != "start" {
		p.twice = true
		return true
	}

	return false
}

func solve(g graph, okToVisit func(*node, *path) bool) {
	var descend func(*node, *path)
	found := 0

	descend = func(from *node, p *path) {
		if from.isSmallCave && !okToVisit(from, p) {
			return
		}

		p.add(from)
		//fmt.Printf("%s\n", p)

		if from.name == "end" {
			found++
			return
		}

		for _, x := range g.edges[*from] {
			descend(x, p.copy())
		}
	}

	descend(g.nodes["start"], &path{})

	fmt.Println(found)
}

func (p *path) nrNodes() int {
	return len(p.nodes)
}

func (p *path) add(n *node) {
	p.nodes = append(p.nodes, n)
}

func (p *path) hasNode(n node) bool {
	for _, x := range p.nodes {
		if x.name == n.name {
			return true
		}
	}
	return false
}

func (p *path) copy() *path {
	x := path{twice: p.twice}
	x.nodes = make([]*node, len(p.nodes))
	copy(x.nodes, p.nodes)
	return &x
}

func (p path) String() string {
	ret := ""
	for _, n := range p.nodes {
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
