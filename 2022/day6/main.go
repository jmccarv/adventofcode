package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		check(s.Text(), 14, check(s.Text(), 4, 0)-4)
	}
}

func check(line string, nrUniq, startOfs int) int {
	nr, j := 0, 0
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
	fmt.Println(nrUniq, nr+1, time.Now().Sub(t0))
	return nr + 1
}
