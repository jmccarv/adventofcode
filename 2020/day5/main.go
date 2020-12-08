package main

import (
	"bufio"
	"fmt"
	"os"
)

func bspish(code string, up rune) int {
	val := 0
	Δ := 1 << (len(code) - 1)
	for _, c := range code {
		if c == up {
			val += Δ
		}
		Δ >>= 1
	}
	return val
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	maxID := 0
	minID := 128*8 + 8
	var occupied [128*8 + 8]bool

	for s.Scan() {
		bpass := s.Text()

		rows := bpass[0:7]
		cols := bpass[7:]

		row := bspish(rows, 'B')
		col := bspish(cols, 'R')

		id := row<<3 | col
		occupied[id] = true

		if id > maxID {
			maxID = id
		} else if id < minID {
			minID = id
		}

		fmt.Printf("%s  row: %3d  col: %3d  id: %d\n", s.Text(), row, col, id)
	}
	fmt.Println("Max ID:", maxID)

	myID := 0
	for i := minID + 1; i < maxID; i++ {
		if !occupied[i] {
			myID = i
			break
		}
	}
	fmt.Println(" My ID:", myID)
}
