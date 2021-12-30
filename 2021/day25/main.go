package main

import (
	"bufio"
	"fmt"
	"os"
)

type grid [][]byte

func main() {
	var game grid
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		game = append(game, []byte(input.Text()))
	}

	nr := 1
	for game.move() {
		nr++
	}
	fmt.Println(nr)
}

func (g grid) move() (moved bool) {
	h := g.copy()
	for r, l := range h {
		for c, v := range l {
			if v != '>' {
				continue
			}
			if loc, ok := h.tryEast(r, c); ok {
				g[r][c] = '.'
				g[r][loc] = '>'
				moved = true
			}
		}
	}

	h = g.copy()
	for r, l := range h {
		for c, v := range l {
			if v != 'v' {
				continue
			}
			if loc, ok := h.trySouth(r, c); ok {
				g[r][c] = '.'
				g[loc][c] = 'v'
				moved = true
			}
		}
	}
	return
}

func (g grid) String() (ret string) {
	for i, l := range g {
		if i > 0 {
			ret += "\n"
		}
		ret += string(l)
	}
	return
}

func (g grid) tryEast(r, c int) (int, bool) {
	c++
	if c > len(g[r])-1 {
		c = 0
	}
	return c, g[r][c] == '.'

}

func (g grid) trySouth(r, c int) (int, bool) {
	r++
	if r > len(g)-1 {
		r = 0
	}
	return r, g[r][c] == '.'
}

func (g grid) copy() grid {
	h := make([][]byte, 0, len(g))
	for _, l := range g {
		m := make([]byte, len(l))
		copy(m, l)
		h = append(h, m)
	}
	return h
}
