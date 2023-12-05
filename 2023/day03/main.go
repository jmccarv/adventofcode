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
}

func (c GBCell) Equals(d GBCell) bool {
	return c.Loc.Equals(d.Loc) && c.Val == d.Val
}

// game boards are expected to be square
type GameBoard struct {
	Max   s2d.Point
	Empty rune
	Cells [][]GBCell
}

func ReadGameBoard(r io.Reader, empty rune) *GameBoard {
	gb := GameBoard{Empty: empty}
	s := bufio.NewScanner(r)
	y, x := 0, 0
	for s.Scan() {
		x = 0
		var cells []GBCell
		for _, r := range s.Text() {
			cells = append(cells, GBCell{Loc: s2d.Point{x, y}, Val: r})
			x++
		}
		gb.Cells = append(gb.Cells, cells)
		y++
	}
	gb.Max = s2d.Point{x - 1, y - 1}
	return &gb
}

func (gb *GameBoard) Rows() int {
	return gb.Max.Y + 1
}

func (gb *GameBoard) Cols() int {
	return gb.Max.X + 1
}

func (gb *GameBoard) Bounds() s2d.Box {
	return s2d.Box{TL: s2d.Point{0, 0}, BR: gb.Max}
}

func (gb *GameBoard) String() string {
	ret := fmt.Sprintf("Rows: %d Cols: %d Max: %+v Empty '%v'\n", gb.Rows(), gb.Cols(), gb.Max, gb.Empty)
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
	ret := fmt.Sprintf("Rows: %d Cols: %d Max: %+v Empty '%v'\n", gb.Rows(), gb.Cols(), gb.Max, gb.Empty)
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
	n := GameBoard{
		Max:   max,
		Empty: gb.Empty,
	}
	for y := 0; y < n.Rows(); y++ {
		n.Cells = append(n.Cells, gb.Cells[from.TL.Y+y][from.TL.X:from.TL.X+n.Cols()])
	}

	return &n
}

func (gb *GameBoard) NeighborSlice(of s2d.Bounder) *GameBoard {
	b := of.Bounds()
	return gb.Slice(s2d.Box{
		TL: b.TL.Sub(s2d.Point{1, 1}),
		BR: b.BR.Add(s2d.Point{1, 1}),
	})
}

func (gb *GameBoard) NeighboringCells(of s2d.Bounder) []GBCell {
	b := of.Bounds()
	var ret []GBCell
	for _, cell := range gb.NeighborSlice(b).Flat() {
		if !b.Contains(cell.Loc) {
			ret = append(ret, cell)
		}
	}
	return ret
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

type Number struct {
	s2d.Box
	val int
}

func main() {
	var nr Number
	var p1, p2 int
	var nums []Number
	gb := ReadGameBoard(os.Stdin, '.')

	// Locate numbers first
	for _, cell := range gb.Flat() {
		if cell.Val < '0' || cell.Val > '9' {
			if nr.val > 0 {
				nums = append(nums, nr)
			}
			nr = Number{}
			continue
		}

		if nr.val == 0 {
			nr.TL = cell.Loc
		}
		nr.BR = cell.Loc
		nr.val = nr.val*10 + int(cell.Val-'0')
	}
	if nr.val > 0 {
		nums = append(nums, nr)
	}

	gears := make(map[GBCell][]int) // We need these for part 2
	for _, nr := range nums {
		p1Num := 0
		for _, cell := range gb.NeighboringCells(nr.Box) {
			if !isSym(cell.Val) {
				continue
			}
			p1Num = nr.val
			if cell.Val == '*' {
				gears[cell] = append(gears[cell], nr.val)
			}
		}
		p1 += p1Num
	}

	for _, g := range gears {
		if len(g) == 2 {
			p2 += g[0] * g[1]
		}
	}

	fmt.Println("Part1", p1)
	fmt.Println("Part2", p2)
}
