package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	maxID := 0
	minID := 128*8 + 8
	var occupied [128*8 + 8]bool

	for s.Scan() {
		bpass := s.Text()

		rows := bpass[0:7]
		cols := bpass[7:]

		row := 0
		Δ := 64
		for _, v := range rows {
			if v == 'B' {
				row += Δ
			}
			Δ >>= 1
		}

		col := 0
		Δ = 4
		for _, v := range cols {
			if v == 'R' {
				col += Δ
			}
			Δ >>= 1
		}

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
