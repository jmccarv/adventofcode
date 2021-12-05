package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	p1    point
	p2    point
	slope float64
}

type ventField [][]int

func main() {
	s := bufio.NewScanner(os.Stdin)
	var maxX, maxY int
	var lines []line

	for s.Scan() {
		coords := strings.Split(s.Text(), " -> ")
		if len(coords) != 2 {
			panic("Invalid input")
		}

		l := lineFromStr(coords[0], coords[1])
		m := l.max()
		maxX = max(maxX, m.x)
		maxY = max(maxY, m.y)
		lines = append(lines, l)
	}

	field := newVentField(maxX, maxY)

	fmt.Println(maxX, maxY)
	part1(field, lines)
}

func part1(field ventField, lines []line) {
	for _, l := range lines {
		if !(l.horizontal() || l.vertical()) {
			// Only interested in horizontal or vertical lines
			continue
		}
		field.drawLine(l)
	}
	fmt.Println(len(field.overlapping()))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func newVentField(mx, my int) ventField {
	field := ventField(make([][]int, my+1))
	for y := 0; y <= my; y++ {
		field[y] = make([]int, mx+1)
	}
	return field
}

func lineFromStr(x, y string) line {
	l := line{p1: pointFromStr(x), p2: pointFromStr(y)}
	l.slope = (float64(l.p2.y) - float64(l.p1.y)) / (float64(l.p2.x) - float64(l.p1.x))
	return l
}

func pointFromStr(xyStr string) point {
	var p point
	var err error

	xy := strings.Split(xyStr, ",")
	if len(xy) != 2 {
		panic("Invalid input")
	}

	p.x, err = strconv.Atoi(xy[0])
	if err != nil {
		panic("Invalid input")
	}

	p.y, err = strconv.Atoi(xy[1])
	if err != nil {
		panic("Invalid niput")
	}

	return p
}

func (f ventField) String() string {
	s := ""
	for _, r := range f {
		for _, c := range r {
			if c > 0 {
				s += fmt.Sprintf("%d", c)
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (f ventField) drawLine(l line) {
	var move func(p point) point

	/*
		dx := min(1, max(-1, l.p2.x-l.p1.x))
		dy := min(1, max(-1, l.p2.y-l.p1.y))

		move = func(p point) point {
			p.x += dx
			p.y += dy
			return p
		}
	*/

	if l.horizontal() {
		dir := 1
		if l.p1.x > l.p2.x {
			dir = -1
		}
		move = func(p point) point {
			p.x += dir
			return p
		}
	} else if l.vertical() {
		dir := 1
		if l.p1.y > l.p2.y {
			dir = -1
		}
		move = func(p point) point {
			p.y += dir
			return p
		}
	} else {

	}

	p := l.p1
	//fmt.Println("Drawing:", l)
	for p != l.p2 {
		f[p.y][p.x]++
		p = move(p)
	}
	f[p.y][p.x]++

	// don't want to do this for the real input, way too big!
	//fmt.Println(f)
}

func (f ventField) overlapping() []point {
	var ret []point

	for y, r := range f {
		for x, c := range r {
			if c > 1 {
				ret = append(ret, point{x, y})
			}
		}
	}

	return ret
}

func (l line) max() point {
	var p point
	p.x = max(l.p1.x, l.p2.x)
	p.y = max(l.p1.y, l.p2.y)
	return p
}

func (l line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d m=%v", l.p1.x, l.p1.y, l.p2.x, l.p2.y, l.slope)
}

func (l line) horizontal() bool {
	if l.p1.y == l.p2.y {
		return true
	}
	return false
}

func (l line) vertical() bool {
	if l.p1.x == l.p2.x {
		return true
	}
	return false
}
