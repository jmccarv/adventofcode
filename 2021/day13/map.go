package main

// This version uses a map to track just the points that are in use instead
// of the slice version that kept a giant 2D slice of all points. This one
// performs less processing and runs faster than the slice version.

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

type page struct {
	points map[point]struct{}
	limit  point
}

func main() {
	var folds []fold
	p := newPage()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		f := strings.Split(s.Text(), ",")

		if len(f) == 2 {
			x, _ := strconv.Atoi(f[0])
			y, _ := strconv.Atoi(f[1])
			p.add(point{x, y})
			continue
		}

		f = strings.Split(s.Text(), "=")
		if len(f) == 2 {
			axis := f[0][len(f[0])-1]
			a, _ := strconv.Atoi(f[1])
			folds = append(folds, fold{axis, a})
		}
	}

	solve(p, folds)
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

func (p *page) add(t point) {
	p.points[t] = struct{}{}
	p.limit.x = max(p.limit.x, t.x)
	p.limit.y = max(p.limit.y, t.y)
}

func newPage() page {
	p := page{points: make(map[point]struct{})}
	return p
}

func (p page) count() (nr int) {
	return len(p.points)
}

func (p *page) fold(f fold) {
	var calc func(t point) (point, bool)

	switch f.axis {
	case 'x':
		calc = func(t point) (point, bool) {
			if t.x > f.where {
				t.x = f.where - (t.x - f.where)
				return t, true
			}
			return t, false
		}
		p.limit.x = f.where - 1
	case 'y':
		calc = func(t point) (point, bool) {
			if t.y > f.where {
				t.y = f.where - (t.y - f.where)
				return t, true
			}
			return t, false
		}
		p.limit.y = f.where - 1
	}

	for t := range p.points {
		if n, ok := calc(t); ok {
			p.points[n] = struct{}{}
			delete(p.points, t)
		}
	}
}

func (p page) String() (s string) {
	for y := 0; y <= p.limit.y; y++ {
		for x := 0; x <= p.limit.x; x++ {
			if _, ok := p.points[point{x, y}]; ok {
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
