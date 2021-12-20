package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type point struct {
	x, y int
}

type box struct {
	tl point
	br point
}

type hit struct {
	v0 int // initial velocity
	v1 int // final velocity
	t  int // time of hit
	s  int // net distance traveled from origin
}

// distance traveled (s): s = (v0 + v1) * t/2
// where v0 = initial velocity
//       v1 = final velocity
//       t  = time
func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}
	re := regexp.MustCompile(`[-0-9]+`)
	m := re.FindAllString(s.Text(), 4)
	if len(m) != 4 {
		panic("Invalid input")
	}

	// target is defined by the top-left and bottom-right points of a box
	// y is negative, so max gets the top of the box
	target := box{
		point{min(atoi(m[0]), atoi(m[1])), max(atoi(m[2]), atoi(m[3]))},
		point{max(atoi(m[0]), atoi(m[1])), min(atoi(m[2]), atoi(m[3]))},
	}

	xhits, yhits, xtmap := findHits(target)
	maxY := findMaxY(target, xhits, yhits, xtmap)
	fmt.Println("Part1", maxY)

	part2(target, xhits, yhits, xtmap)
}

func part2(target box, xhits, yhits []hit, xtmap map[int][]hit) {
	// find how many initial velocities hit the target
	h := make(map[point]bool)
	for _, y := range yhits {
		for _, x := range xtmap[y.t] {
			h[point{x.v0, y.v0}] = true
		}
	}
	fmt.Println("Part2", len(h))
}

func findHits(target box) ([]hit, []hit, map[int][]hit) {
	// Find all possible x,t and y,t pairs that are hits

	// Y
	// min(v0) is negative -- br.y
	// max(v0) value would be the (-br.y)-1 (a positive point), because when the
	// probe comes down, it'll reach y=0 with a velocity of -v0; it's next step
	// will have velocity of -v0-1
	var maxt int
	var yhits []hit
	for v0 := target.br.y; v0 <= (-target.br.y)-1; v0++ {
		t := 1
		s, v1 := yTraveled(v0, t)
		for s >= target.br.y {
			if target.hit(point{target.tl.x, s}) {
				yhits = append(yhits, hit{v0, v1, t, s})
			}
			t++
			s, v1 = yTraveled(v0, t)
			maxt = max(maxt, t)
		}
	}

	// X
	// There is a minimum initial velocity needed to reach the target.
	// where v1 = 0 and s = target.tl.x but I don't know how to calculate
	// it so we'll just try them all up to our known max v0
	// And since the probe will eventually come to a stop, we want to
	// keep finding hits up through our max(t) we found above
	var xhits []hit
	for v0 := 1; v0 <= target.br.x; v0++ {
		for t := 1; t <= maxt; t++ {
			s, v1 := xTraveled(v0, t)
			if target.hit(point{s, target.tl.y}) {
				xhits = append(xhits, hit{v0, v1, t, s})
			}
		}
	}

	// Map xhits by time so we can intersect with Y hits
	xmap := make(map[int][]hit)
	for _, hit := range xhits {
		xmap[hit.t] = append(xmap[hit.t], hit)
	}

	return xhits, yhits, xmap
}

func findMaxY(target box, xhits, yhits []hit, xtmap map[int][]hit) int {
	// Now find the y hits with the highest v0 that hit at the same
	// time as x. if that makes sense.
	sort.SliceStable(yhits, func(i, j int) bool { return yhits[i].v0 > yhits[i].v1 })
	for _, y := range yhits {
		if _, ok := xtmap[y.t]; ok {
			s, _ := yTraveled(y.v0, y.v0)
			return s
		}
	}

	return 0
}

func xTraveled(v0, t1 int) (s, v1 int) {
	v1 = v0
	for t := 0; t < t1 && v1 > 0; t++ {
		s += v1
		v1--
	}
	return
}

func yTraveled(v0, t1 int) (s, v1 int) {
	v1 = v0
	for t := 0; t < t1; t++ {
		s += v1
		v1--
	}
	return
}

func (b box) hit(p point) bool {
	if p.x < b.tl.x || p.x > b.br.x {
		return false
	}
	if p.y > b.tl.y || p.y < b.br.y {
		return false
	}
	return true
}

func atoi(a string) int {
	x, _ := strconv.Atoi(a)
	return x
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
