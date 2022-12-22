package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	sm "github.com/jmccarv/adventofcode/util/math"
	s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

type grid struct {
	s2d.Box
	floor  int
	points map[s2d.Point]byte
}

func main() {
	cave := grid{points: make(map[s2d.Point]byte), Box: s2d.Box{TL: s2d.Point{X: math.MaxInt}}}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var p1, p2 s2d.Point
		points := strings.Split(s.Text(), "->")
		for _, p := range points {
			fmt.Sscanf(p, "%d,%d", &p2.X, &p2.Y)
			if !p1.Equals(s2d.Point{0, 0}) {
				cave.line(p1, p2)
			}
			cave.TL = cave.TL.Min(p2)
			cave.BR = cave.BR.Max(p2)
			p1 = p2
		}
	}
	//cave.dump()
	nr := 0
	for cave.sand() {
		nr++
	}
	fmt.Println("Part1", nr)

	cave.floor = cave.BR.Y + 2
	fmt.Println("floor", cave.floor)
	for !cave.occupied(s2d.Point{X: 500}) && cave.sand() {
		nr++
		/*
			if nr%100 == 0 {
				fmt.Println(nr)
				cave.dump()
			}
		*/
	}
	fmt.Println("Part2", nr)
}

func (g grid) inBounds(p s2d.Point) bool {
	if g.floor > 0 {
		return p.Y >= 0 && p.Y <= g.floor
	}
	return p.Y >= 0 && p.Y <= g.BR.Y && p.X >= g.TL.X && p.X <= g.BR.X
}

// returns true if the sand settled and we should continue
// returns false if the sand fell off into the void
func (g grid) sand() bool {
	find := func(p s2d.Point) s2d.Point {
		for p.Y++; g.inBounds(p); p.Y++ {
			if !g.occupied(p) {
				continue
			}
			p.X--
			if !g.occupied(p) {
				continue
			}
			p.X += 2
			if !g.occupied(p) {
				continue
			}
			return p.Sub(s2d.Point{1, 1})
		}
		return p
	}

	p := s2d.Point{X: 500}
	p = find(p)
	if g.occupied(p.Add(s2d.Point{Y: 1})) {
		g.addSand(p)
		return true
	}
	return false
}

func (g grid) addSand(p s2d.Point) {
	g.points[p] = 'o'
	g.TL = g.TL.Min(p)
	g.BR = g.BR.Max(p)
}

func (g grid) line(p1, p2 s2d.Point) {
	dir := p1.DirectionTo(p2) // direction to move to get from p1 to p2

	for ; !p1.Equals(p2); p1 = p1.Add(dir) {
		g.points[p1] = '#'
	}
	g.points[p1] = '#'
}

func (g grid) occupied(p s2d.Point) bool {
	if p.Y == g.floor {
		return true
	}
	_, ok := g.points[p]
	return ok
}

func (g grid) dump() {
	fmt.Printf("%v - %v\n", g.TL, g.BR)
	ofs := g.BR.Sub(g.TL)
	for y := 0; y <= sm.Max(ofs.Y, g.floor); y++ {
		for x := 0; x <= ofs.X; x++ {
			if p, ok := g.points[s2d.Point{X: g.TL.X + x, Y: g.TL.Y + y}]; ok {
				fmt.Printf("%c", p)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
