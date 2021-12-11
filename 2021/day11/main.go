package main

import (
	"bufio"
	"fmt"
	"os"
)

type octopus struct {
	energy  int
	flashed bool
}

type gameGrid [10][10]octopus

func main() {
	s := bufio.NewScanner(os.Stdin)

	var grid gameGrid
	r := 0
	for s.Scan() {
		for c, nr := range s.Text() {
			grid[r][c].energy = int(nr - '0')
		}
		r++
	}

	p1 := 0
	p2 := 0
	for i := 0; p2 == 0; i++ {
		nr, all := grid.step()
		if i < 100 {
			p1 += nr
		}
		if all {
			p2 = i + 1
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func (g *gameGrid) step() (int, bool) {
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			g[r][c].energy++
			if g[r][c].energy > 9 {
				g.flash(r, c)
			}
		}
	}

	flashed := 0
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			if g[r][c].flashed {
				flashed++
				g[r][c].flashed = false
				g[r][c].energy = 0
			}
		}
	}

	return flashed, flashed == 100
}

func (g *gameGrid) flash(r, c int) {
	if g[r][c].flashed {
		return
	}

	g[r][c].flashed = true

	g.neighbors(r, c, func(r, c int) {
		if g[r][c].flashed {
			return
		}

		g[r][c].energy++
		if g[r][c].energy > 9 {
			g.flash(r, c)
		}
	})
}

func (g *gameGrid) neighbors(r, c int, cb func(int, int)) {
	for dr := -1; dr < 2; dr++ {
		for dc := -1; dc < 2; dc++ {
			rr := r + dr
			cc := c + dc
			switch {
			case rr < 0 || rr > 9 || cc < 0 || cc > 9:
				continue
			case rr == r && cc == c:
				continue
			}
			cb(rr, cc)
		}
	}
}

func (g gameGrid) String() (ret string) {
	for _, row := range g {
		for _, o := range row {
			f := " "
			if o.flashed {
				f = "*"
			}
			ret += fmt.Sprintf("%s%2d ", f, o.energy)
		}
		ret += "\n"
	}
	return
}
