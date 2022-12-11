package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	t0 := time.Now()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		check(s.Text(), 14, check(s.Text(), 4, 0)-4)
	}
	fmt.Println("Total time ", time.Now().Sub(t0))
}

func check(line string, nrUniq, startOfs int) int {
	var nr, j int
	t0 := time.Now()
	for i := startOfs + nrUniq - 1; i < len(line) && nr == 0; i = j + nrUniq {
		var seen [123]bool // ASCII 'z' == 122
		nr = i
		for j = i; j >= i-(nrUniq-1); j-- {
			if seen[line[j]] {
				nr = 0
				break
			}
			seen[line[j]] = true
		}
	}
	fmt.Printf("%2d %4d %v\n", nrUniq, nr+1, time.Now().Sub(t0))
	return nr + 1
}
