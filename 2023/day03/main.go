package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

type GBCell struct {
	Loc s2d.Point
	Val rune
}

func (c GBCell) String() string {
	return string(c.Val)
}

func (c GBCell) Dump() string {
	sym := " "
	if isSym(c.Val) {
		sym = "+"
	}
	return fmt.Sprintf("[%d,%d %s%s]", c.Loc.X, c.Loc.Y, string(c.Val), sym)
	//return fmt.Sprintf("%+v %s(%s)  ", c.Loc, string(c.Val), sym)
}

func (c GBCell) Equals(d GBCell) bool {
	return c.Loc.Equals(d.Loc) && c.Val == d.Val
}

// game boards are expected to be square
type GameBoard struct {
	Rows, Cols int
	Max        s2d.Point
	Empty      rune
	Cells      [][]GBCell
}

func ReadGameBoard(r io.Reader, empty rune) *GameBoard {
	gb := GameBoard{Empty: empty}
	s := bufio.NewScanner(r)
	for s.Scan() {
		x := 0
		var cells []GBCell
		for _, r := range s.Text() {
			cells = append(cells, GBCell{Loc: s2d.Point{x, gb.Rows}, Val: r})
			x++
		}
		gb.Cells = append(gb.Cells, cells)
		gb.Rows++
	}
	gb.Cols = len(gb.Cells[0])
	gb.Max = s2d.Point{X: gb.Cols - 1, Y: gb.Rows - 1}
	return &gb
}

func (gb *GameBoard) String() string {
	ret := fmt.Sprintf("Rows: %d Cols: %d Max: %+v Empty '%v'\n", gb.Rows, gb.Cols, gb.Max, gb.Empty)
	for _, r := range gb.Cells {
		if ret != "" {
			ret += "\n"
		}
		for _, c := range r {
			ret += c.String()
		}
	}
	return ret
}

func (gb *GameBoard) Dump() string {
	ret := fmt.Sprintf("Rows: %d Cols: %d Max: %+v Empty '%v'\n", gb.Rows, gb.Cols, gb.Max, gb.Empty)
	for _, r := range gb.Cells {
		if ret != "" {
			ret += "\n"
		}
		for _, c := range r {
			ret += c.Dump()
		}
	}
	return ret
}

func (gb *GameBoard) Mirror() *GameBoard {
	for i, _ := range gb.Cells {
		slices.Reverse(gb.Cells[i])
		for j, _ := range gb.Cells[i] {
			gb.Cells[i][j].Loc.X = len(gb.Cells[i]) - gb.Cells[i][j].Loc.X - 1
		}
	}
	return gb
}

func (gb *GameBoard) Slice(from s2d.Box) *GameBoard {
	from.TL = from.TL.Max(s2d.Point{0, 0})
	from.BR = from.BR.Min(gb.Max)
	max := from.BR.Sub(from.TL)
	//fmt.Println("from", from, "max", max)
	n := GameBoard{
		Rows:  max.Y + 1,
		Cols:  max.X + 1,
		Max:   max,
		Empty: gb.Empty,
	}
	//fmt.Printf("%+v\n", n)
	for y := 0; y < n.Rows; y++ {
		//fmt.Println(y, len(gb.Cells[y]), from.TL.Y, from.TL.X)
		n.Cells = append(n.Cells, gb.Cells[from.TL.Y+y][from.TL.X:from.TL.X+n.Cols])
	}

	return &n
}

func (gb *GameBoard) Neighbors(of s2d.Point) *GameBoard {
	return gb.Slice(s2d.Box{
		TL: of.Sub(s2d.Point{1, 1}),
		BR: of.Add(s2d.Point{1, 1}),
	})
}

func (gb *GameBoard) Flat() []GBCell {
	var ret []GBCell
	for _, r := range gb.Cells {
		ret = append(ret, r...)
	}
	return ret
}

func isSym(r rune) bool {
	return strings.IndexRune(`#$%&*+-/=@`, r) > -1
}

func main() {
	var nr, p1, p2 int
	var keep bool
	gears := make(map[GBCell][]int)
	stars := make(map[GBCell]struct{})
	gb := ReadGameBoard(os.Stdin, '.')
	for _, cell := range gb.Flat() {
		if cell.Val < '0' || cell.Val > '9' {
			if cell.Val == '*' {
				stars[cell] = struct{}{}
			}
			if nr > 0 && keep {
				p1 += nr
				for cell := range stars {
					gears[cell] = append(gears[cell], nr)
				}
			}
			clear(stars)
			nr, keep = 0, false
			continue
		}
		for _, n := range gb.Neighbors(cell.Loc).Flat() {
			if isSym(n.Val) {
				keep = true
				if n.Val == '*' {
					stars[n] = struct{}{}
				}
			}
		}
		nr = nr*10 + int(cell.Val-'0')
	}
	if nr > 0 && keep {
		p1 += nr
		for cell := range stars {
			gears[cell] = append(gears[cell], nr)
		}
	}

	for _, g := range gears {
		if len(g) == 2 {
			p2 += g[0] * g[1]
		}
	}

	fmt.Println("Part1", p1)
	fmt.Println("Part2", p2)
}
