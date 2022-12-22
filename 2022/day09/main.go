package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jmccarv/adventofcode/util/math"
	s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

type knot struct {
	s2d.Point
	visited map[s2d.Point]struct{}
}

type rope [10]knot

var (
	directions = map[int]s2d.Point{
		'U': s2d.Point{0, -1},
		'D': s2d.Point{0, 1},
		'L': s2d.Point{-1, 0},
		'R': s2d.Point{1, 0},
	}

	knots rope
)

func main() {
	t0 := time.Now()

	// Part 1 wants to know how many unique locations the second knot (first tail) visited
	knots[1].visited = map[s2d.Point]struct{}{s2d.Point{0, 0}: struct{}{}}

	// Part 2 is how many unique locations the final knot visited
	knots[len(knots)-1].visited = map[s2d.Point]struct{}{s2d.Point{0, 0}: struct{}{}}

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
			knots[0].Point = knots[0].Add(dir)

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

	ofs := h.Sub(k.Point)
	dir := s2d.Point{X: math.Sign(ofs.X), Y: math.Sign(ofs.Y)}
	k.Point = k.Add(dir)

	if k.visited != nil {
		k.visited[k.Point] = struct{}{}
	}
}

func (k knot) neighbors(h knot) bool {
	ofs := h.Sub(k.Point)
	return math.Abs(ofs.X) <= 1 && math.Abs(ofs.Y) <= 1
}

func (k knot) String() string {
	return fmt.Sprintf("%v", k.Point)
}

func (r rope) dump() {
	// Get the bouding box first
	var tl, br s2d.Point
	for _, k := range r {
		tl = tl.Min(k.Point)
		br = br.Max(k.Point)
	}
	// Total size of our grid
	size := br.Sub(tl).Add(s2d.Point{1, 1})

	// ofs will be what we translate all our x and y values by to get their location on our grid
	ofs := s2d.Point{}.Sub(tl)

	grid := make([][]byte, size.Y)
	for r := 0; r < size.Y; r++ {
		grid[r] = make([]byte, size.X)
		for c := 0; c < size.X; c++ {
			grid[r][c] = '.'
		}
	}

	grid[ofs.Y][ofs.X] = 's' // starting point
	for i := len(r) - 1; i >= 0; i-- {
		// translate this one to it's location
		loc := ofs.Add(r[i].Point)
		grid[loc.Y][loc.X] = byte('0' + i)
	}

	// And now we can print it
	for _, r := range grid {
		fmt.Println(string(r))
	}
	fmt.Println("p1", len(r[1].visited), "  p2", len(r[len(r)-1].visited))
	fmt.Println()
}
