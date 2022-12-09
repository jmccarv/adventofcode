package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type point struct {
	x, y int
}

type knot struct {
	point
	visited map[point]struct{}
}

var directions = map[int]point{
	'U': point{0, -1},
	'D': point{0, 1},
	'L': point{-1, 0},
	'R': point{1, 0},
}

type rope [10]knot

var knots rope

func main() {
	t0 := time.Now()

	// Part 1 wants to know how many unique locations the second knot (first tail) visited
	knots[1].visited = map[point]struct{}{point{0, 0}: struct{}{}}

	// Part 2 is how many unique locations the final knot visited
	knots[len(knots)-1].visited = map[point]struct{}{point{0, 0}: struct{}{}}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var udrl, amt int
		nr, err := fmt.Sscanf(s.Text(), "%c %d", &udrl, &amt)
		if err != nil || nr != 2 {
			panic("Invalid input!")
		}

		//fmt.Println(s.Text())
		dir := directions[udrl]

		for z := 0; z < amt; z++ {
			// knots[0] is our 'Head' knot
			knots[0].point.add(dir)

			// Now successive knots follow the one before it so that they
			// are always neighboring the preceding knot.
			for i := 1; i < len(knots); i++ {
				knots[i].follow(knots[i-1])
			}
		}
	}

	fmt.Println(len(knots[1].visited))
	fmt.Println(len(knots[len(knots)-1].visited))
	fmt.Println("Total time ", time.Now().Sub(t0))
}

// Knot k follows knot h by a single step, abiding by the rules of the game
func (k *knot) follow(h knot) {
	if k.neighbors(h) {
		// No need to move if we're already neighbors
		return
	}

	ofs := sub(h.point, k.point)
	dir := point{x: sign(ofs.x), y: sign(ofs.y)}
	k.add(dir)

	if k.visited != nil {
		k.visited[k.point] = struct{}{}
	}
}

func (p *point) add(q point) {
	*p = add(*p, q)
}

func (p *point) sub(q point) {
	*p = sub(*p, q)
}

func (p point) min(q point) point {
	return point{x: min(p.x, q.x), y: min(p.y, q.y)}
}

func (p point) max(q point) point {
	return point{x: max(p.x, q.x), y: max(p.y, q.y)}
}

func (k knot) neighbors(h knot) bool {
	ofs := sub(h.point, k.point)
	return abs(ofs.x) <= 1 && abs(ofs.y) <= 1
}

func (k knot) String() string {
	return fmt.Sprintf("%v", k.point)
}

func add(p, q point) point {
	return point{x: p.x + q.x, y: p.y + q.y}
}

func sub(p, q point) point {
	return point{x: p.x - q.x, y: p.y - q.y}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -1
	}
	return 1
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

func (r rope) dump() {
	// Get the bouding box first
	var tl, br point
	for _, k := range r {
		tl = tl.min(k.point)
		br = br.max(k.point)
	}
	// Total size of our grid
	size := add(sub(br, tl), point{1, 1})

	// ofs will be what we translate all our x and y values by to get their location on our grid
	ofs := sub(point{}, tl)

	grid := make([][]byte, size.y)
	for r := 0; r < size.y; r++ {
		grid[r] = make([]byte, size.x)
		for c := 0; c < size.x; c++ {
			grid[r][c] = '.'
		}
	}

	grid[ofs.y][ofs.x] = 's' // starting point
	for i := len(r) - 1; i >= 0; i-- {
		// translate this one to it's location
		loc := add(ofs, r[i].point)
		grid[loc.y][loc.x] = byte('0' + i)
	}

	// And now we can print it
	for _, r := range grid {
		fmt.Println(string(r))
	}
	fmt.Println("p1", len(r[1].visited), "  p2", len(r[len(r)-1].visited))
	fmt.Println()
}
