package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold struct {
	axis  byte
	where int
}

type page [][]bool

func main() {
	var m point
	var points []point
	var folds []fold

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		f := strings.Split(s.Text(), ",")
		if len(f) == 2 {
			x, _ := strconv.Atoi(f[0])
			y, _ := strconv.Atoi(f[1])
			points = append(points, point{x, y})
			m.x = max(m.x, x)
			m.y = max(m.y, y)
			continue
		}

		f = strings.Split(s.Text(), "=")
		if len(f) == 2 {
			axis := f[0][len(f[0])-1]
			a, _ := strconv.Atoi(f[1])
			folds = append(folds, fold{axis, a})
		}
	}

	solve(newPage(points, m), folds)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(p page, folds []fold) {
	p.fold(folds[0])
	fmt.Println("Part1", p.count())

	for _, f := range folds[1:] {
		p.fold(f)
	}
	fmt.Println(p)
}

func newPage(points []point, bounds point) page {
	p := make([][]bool, bounds.y+1)
	for y := 0; y <= bounds.y; y++ {
		p[y] = make([]bool, bounds.x+1)
	}

	for _, t := range points {
		p[t.y][t.x] = true
	}

	return p
}

func (p page) count() (nr int) {
	for _, y := range p {
		for _, set := range y {
			if set {
				nr++
			}
		}
	}
	return
}

func (p *page) fold(f fold) {
	switch f.axis {
	case 'x':
		p.foldX(f.where)
	case 'y':
		p.foldY(f.where)
	}
}

func (pp *page) foldX(w int) {
	p := *pp
	for y := 0; y < len(p); y++ {
		for x := w + 1; x < len(p[y]); x++ {
			if p[y][x] {
				p[y][w-(x-w)] = true
			}
		}
		p[y] = p[y][0:w]
	}
}

func (pp *page) foldY(w int) {
	p := *pp
	for y := w + 1; y < len(p); y++ {
		for x := 0; x < len(p[0]); x++ {
			if p[y][x] {
				p[w-(y-w)][x] = true
			}
		}
	}
	*pp = p[0:w]
}

func (p page) String() (s string) {
	for _, y := range p {
		for _, set := range y {
			if set {
				s += "#"
			} else {
				s += " "
			}
		}
		s += "\n"
	}

	if len(s) > 0 {
		s = s[:len(s)]
	}

	return
}
