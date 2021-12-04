package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

type cell struct {
	val    int
	r      int
	c      int
	marked bool
}

type player struct {
	board       [boardSize][boardSize]cell
	winningCell *cell
	filled      bool
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var draws []int

	// the first line of input is expected to be a comma separated list of
	// bingo numbers to draw.
	if !s.Scan() {
		panic("Invalid input")
	}
	for _, t := range strings.Split(s.Text(), ",") {
		i, err := strconv.Atoi(t)
		if err != nil {
			panic("Invalid input")
		}
		draws = append(draws, i)
	}
	if len(draws) < 2 {
		panic("Invalid input")
	}
	fmt.Println(draws)

	// Now we are expecting to find puzzle boards separated by blank lines
	// Each board consits of rows of space separated numbers
	// Each player gets a board and then the game starts!
	var players []*player
	var r int
	p := &player{}

	for s.Scan() {
		nums := strings.Fields(s.Text())
		if len(nums) == 0 {
			if p.filled {
				players = append(players, p)
				p = &player{}
				r = 0
			}
			continue
		} else if len(nums) != boardSize {
			panic("Invalid input!")
		}

		for c, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic("Invalid input!")
			}
			p.board[r][c] = cell{val: i, r: r, c: c}
			p.filled = true
		}

		r++
	}
	if p.filled {
		players = append(players, p)
	}

	part1(draws, players)
}

func part1(draws []int, players []*player) {
out:
	for _, nr := range draws {
		for _, p := range players {
			if p.checkDraw(nr) {
				// found a winner
				fmt.Println(p)
				break out
			}
		}
	}
}

// return true if it was a winning draw
func (p *player) checkDraw(nr int) bool {
	// First try to find the number on our board
	for r := range p.board {
		for c := range p.board[r] {
			if p.board[r][c].val == nr {
				// found the number
				p.board[r][c].marked = true

				// is this a winning draw?
				if p.checkRow(r) || p.checkCol(c) || p.checkDiag() {
					p.winningCell = &p.board[r][c]
					return true
				}
				return false
			}
		}
	}
	return false
}

func (p *player) checkRow(r int) bool {
	for c := 0; c < boardSize; c++ {
		if !p.board[r][c].marked {
			return false
		}
	}
	return true
}

func (p *player) checkCol(c int) bool {
	for r := 0; r < boardSize; r++ {
		if !p.board[r][c].marked {
			return false
		}
	}
	return true
}

func (p *player) checkDiag() bool {
	found := true
out:
	for x := 0; x < boardSize; x++ {
		if !p.board[x][x].marked {
			found = false
			break out
		}
	}
	if found {
		return true
	}

	for x := boardSize - 1; x >= 0; x-- {
		if !p.board[x][x].marked {
			return false
		}
	}
	return true
}

func (p *player) String() string {
	var s string

	if !p.filled {
		return s
	}

	if p.winningCell != nil {
		s = "Winning: " + p.winningCell.String() + "\n"
		s += fmt.Sprintf("Score: %v\n", p.score())
	}

	for _, cell := range p.board {
		for _, c := range cell {
			s += c.String()
		}
		s += "\n"
	}

	return s
}

func (p *player) score() int {
	s := 0

	for r, _ := range p.board {
		for c, _ := range p.board[r] {
			if !p.board[r][c].marked {
				s += p.board[r][c].val
			}
		}
	}
	s *= p.winningCell.val
	return s
}

func (c cell) String() string {
	var s string
	m := " "
	if c.marked {
		m = "*"
	}
	s = fmt.Sprintf("%s%2d ", m, c.val)
	return s
}
