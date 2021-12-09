package main

import (
	"bufio"
	"fmt"
	"os"
)

type heightMap [][]byte

type game struct {
	m          heightMap
	rows, cols int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	var g game
	for s.Scan() {
		g.addRow(s.Text())
	}
	//fmt.Println(g)
	fmt.Println(g.riskLevel())
}

func (g *game) riskLevel() int {
	ret := 0
	for r := 1; r <= g.rows; r++ {
		for c := 1; c <= g.cols; c++ {
			n := g.m[r][c]
			if n < g.m[r-1][c] && n < g.m[r+1][c] && n < g.m[r][c-1] && n < g.m[r][c+1] {
				ret += int(n) + 1
			}
		}
	}
	return ret
}

func (g *game) addRow(line string) {
	l := len(line)

	if g.rows == 0 {
		g.cols = l
		top := make([]byte, l+2)
		bot := make([]byte, l+2)
		for i := range top {
			top[i] = 10
			bot[i] = 10
		}
		g.m = append(g.m, top)
		g.m = append(g.m, bot)

	} else if l != g.cols {
		panic("Invalid input")
	}

	g.rows++

	g.m = append(g.m, g.m[g.rows])
	g.m[g.rows] = make([]byte, l+2)
	g.m[g.rows][0] = 10
	g.m[g.rows][l+1] = 10
	for c, num := range line {
		g.m[g.rows][c+1] = byte(num - '0')
	}
}
