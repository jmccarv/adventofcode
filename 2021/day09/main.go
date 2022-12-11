package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type location struct {
	value   byte
	visited bool
}

type coordinate struct {
	r, c int
}

type game struct {
	m          [][]location
	low        []coordinate
	coords     []coordinate
	rows, cols int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	var g game
	for s.Scan() {
		g.addRow(s.Text())
	}
	// part 1
	fmt.Println(g.riskLevel())

	// part 2
	var basins []int
	for _, c := range g.lowPoints() {
		basins = append(basins, g.basinAt(c))
	}
	if len(basins) < 3 {
		panic("Invalid input")
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	num := 1
	for i := 0; i < 3; i++ {
		num *= basins[i]
	}
	fmt.Println(num)
}

func (g *game) lowPoints() []coordinate {
	if g.low != nil {
		return g.low
	}

	g.low = make([]coordinate, 0)
	for _, c := range g.coordinates() {
		me := g.m[c.r][c.c]
		func() {
			for _, n := range g.neighborsOf(c) {
				if me.value >= g.at(n).value {
					return
				}
			}
			g.low = append(g.low, c)
		}()
	}
	return g.low
}

func (g *game) neighborsOf(c coordinate) []coordinate {
	return []coordinate{
		coordinate{c.r - 1, c.c},
		coordinate{c.r + 1, c.c},
		coordinate{c.r, c.c - 1},
		coordinate{c.r, c.c + 1},
	}
}

func (g *game) coordinates() []coordinate {
	if g.coords != nil {
		return g.coords
	}
	g.coords = make([]coordinate, 0, (g.rows)*(g.cols))
	for r := 1; r <= g.rows; r++ {
		for c := 1; c <= g.cols; c++ {
			g.coords = append(g.coords, coordinate{r, c})
		}
	}
	return g.coords
}

func (g *game) riskLevel() int {
	ret := 0
	for _, c := range g.lowPoints() {
		ret += int(g.at(c).value) + 1
	}
	return ret
}

// returns the size of the basin starting at c
// basin size is the number of locations in the basin, including the starting point
func (g *game) basinAt(c coordinate) int {
	//fmt.Println("Enter: ", c)
	g.at(c).visited = true
	ret := 1

	for _, n := range g.neighborsOf(c) {
		if g.at(n).visited || g.at(n).value < g.at(c).value {
			continue
		}
		//fmt.Println("check: ", x)
		ret += g.basinAt(n)
	}
	return ret
}

func (g *game) addRow(line string) {
	l := len(line)
	border := location{10, true}

	if g.rows == 0 {
		g.cols = l
		top := make([]location, l+2)
		bot := make([]location, l+2)
		for i := range top {
			top[i] = border
			bot[i] = border
		}
		g.m = append(g.m, top)
		g.m = append(g.m, bot)

	} else if l != g.cols {
		panic("Invalid input")
	}

	g.rows++

	g.m = append(g.m, g.m[g.rows])
	g.m[g.rows] = make([]location, l+2)
	g.m[g.rows][0] = border
	g.m[g.rows][l+1] = border
	for c, num := range line {
		loc := location{byte(num - '0'), false}
		if loc.value == 9 {
			loc.visited = true
		}
		g.m[g.rows][c+1] = loc
	}
}

func (g *game) at(rc coordinate) *location {
	return &g.m[rc.r][rc.c]
}
