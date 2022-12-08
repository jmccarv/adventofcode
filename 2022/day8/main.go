package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	U = 0
	D = 1
	L = 2
	R = 3
)

type loc struct {
	r, c int
}

type direction struct {
	r, c int // intended to be used as a direction, i.e. -1 <= val <= 1
}

type tree struct {
	loc
	h   int      //height
	vis [4]*tree // tree furthest away that can be seen from this tree
}

type grove struct {
	t    [][]*tree
	size int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var trees [][]*tree

	r := 0
	for s.Scan() {
		row := make([]*tree, len(s.Text()))
		for c, h := range []byte(s.Text()) {
			row[c] = &tree{loc: loc{r: r, c: c}, h: int(h - '0')}
			row[c].vis[U], row[c].vis[D], row[c].vis[L], row[c].vis[R] = row[c], row[c], row[c], row[c]
		}
		trees = append(trees, row)
		r++
	}
	g := grove{t: trees, size: len(trees[0])}

	for _, t := range g.innerTrees() {
		t.calcVisFrom(g)
	}

	//g.dump()
	part1(g)
	part2(g)
}

func part1(g grove) {
	check := func(t *tree, didx int) bool {
		return g.isEdge(t.vis[didx]) && t.vis[didx].h < t.h
	}
	vis := 0
	for _, t := range g.innerTrees() {
		if check(t, U) || check(t, D) || check(t, L) || check(t, R) {
			vis++
		}
	}
	fmt.Println(vis + (g.size-1)*4)
}

func part2(g grove) {
	res := 0

	dist := func(t *tree, didx int) int {
		v := t.vis[didx]
		switch didx {
		case U:
			return t.r - v.r
		case D:
			return v.r - t.r
		case L:
			return t.c - v.c
		case R:
			return v.c - t.c
		}
		return 0
	}

	for _, t := range g.innerTrees() {
		score := dist(t, U) * dist(t, D) * dist(t, L) * dist(t, R)
		if score > res {
			res = score
		}
	}
	fmt.Println(res)
}

func (t *tree) calcVisFrom(g grove) {
	calc := func(didx int) {
		var dir direction
		switch didx {
		case U:
			dir = direction{-1, 0}
		case D:
			dir = direction{1, 0}
		case L:
			dir = direction{0, -1}
		case R:
			dir = direction{0, 1}
		}
		x := g.move(t, dir)
		for ; ; x = g.move(x, dir) {
			if g.isEdge(x) || x.h >= t.h {
				break
			}
		}
		t.vis[didx] = x
	}

	calc(U)
	calc(D)
	calc(L)
	calc(R)
}

func (g grove) isEdge(t *tree) bool {
	return t.r <= 0 || t.c <= 0 || t.r >= g.size-1 || t.c >= g.size-1
}

func (g grove) move(from *tree, d direction) *tree {
	return g.t[from.r+d.r][from.c+d.c]
}

func (g grove) innerTrees() []*tree {
	var ret []*tree
	for r := 1; r < g.size-1; r++ {
		for c := 1; c < g.size-1; c++ {
			ret = append(ret, g.t[r][c])
		}
	}
	return ret
}

func (g grove) dump() {
	for _, row := range g.t {
		for _, t := range row {
			fmt.Printf("%v ", t)
		}
		fmt.Println()
	}
}

func (t *tree) String() string {
	ret := fmt.Sprintf("[%2d,%2d,%d U:%2d D:%2d L:%2d R:%2d]", t.r, t.c, t.h, t.vis[U].r, t.vis[D].r, t.vis[L].c, t.vis[R].c)
	return ret
}
