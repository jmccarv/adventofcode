package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type box struct {
	tl point
	br point
}

type grid struct {
	box
	points map[point]byte
}

func main() {
	cave := grid{points: make(map[point]byte), box: box{tl: point{x: math.MaxInt}}}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var p1, p2 point
		points := strings.Split(s.Text(), "->")
		for _, p := range points {
			fmt.Sscanf(p, "%d,%d", &p2.x, &p2.y)
			if !p1.equals(point{0, 0}) {
				cave.line(p1, p2)
			}
			cave.tl = cave.tl.min(p2)
			cave.br = cave.br.max(p2)
			p1 = p2
		}
	}
	//cave.dump()
	nr := 0
	for cave.sand() {
		//cave.dump()
		nr++
	}
	fmt.Println("Part1", nr)
}

func (g grid) inBounds(p point) bool {
	return p.y >= 0 && p.y <= g.br.y && p.x >= g.tl.x && p.x <= g.br.x
}

// returns true if the sand settled and we should continue
// returns false if the sand fell off into the void
func (g grid) sand() bool {
	find := func(p point) point {
		for p.y++; g.inBounds(p); p.y++ {
			if !g.occupied(p) {
				continue
			}
			p.x--
			if !g.occupied(p) {
				continue
			}
			p.x += 2
			if !g.occupied(p) {
				continue
			}
			return p.sub(point{1, 1})
		}
		return p
	}

	p := point{x: 500}
	p = find(p)
	if g.occupied(p.add(point{y: 1})) {
		g.points[p] = 'o'
		return true
	}
	return false
}

func (g grid) line(p1, p2 point) {
	var dir point // direction to move to get from p1 to p2
	dir.x = cmp(p2.x, p1.x)
	dir.y = cmp(p2.y, p1.y)

	for ; !p1.equals(p2); p1 = p1.add(dir) {
		g.points[p1] = '#'
	}
	g.points[p1] = '#'
}

func (g grid) occupied(p point) bool {
	_, ok := g.points[p]
	return ok
}

func (p1 point) add(p2 point) point {
	p1.x += p2.x
	p1.y += p2.y
	return p1
}

func (p1 point) min(p2 point) point {
	p1.x = min(p1.x, p2.x)
	p1.y = min(p1.y, p2.y)
	return p1
}

func (p1 point) max(p2 point) point {
	p1.x = max(p1.x, p2.x)
	p1.y = max(p1.y, p2.y)
	return p1
}

func (p1 point) equals(p2 point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cmp(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func (p1 point) sub(p2 point) point {
	p1.x -= p2.x
	p1.y -= p2.y
	return p1
}

func (g grid) dump() {
	fmt.Printf("%v - %v\n", g.tl, g.br)
	ofs := g.br.sub(g.tl)
	for y := 0; y <= ofs.y; y++ {
		for x := 0; x <= ofs.x; x++ {
			if p, ok := g.points[point{x: g.tl.x + x, y: g.tl.y + y}]; ok {
				fmt.Printf("%c", p)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
