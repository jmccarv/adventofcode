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
		fmt.Println(check(s.Text(), 4), check(s.Text(), 14))
	}
}

func check(line string, nrUniq int) int {
	nr, j := 0, 0
	t0 := time.Now()
	for i := nrUniq - 1; i < len(line) && nr == 0; i = j + nrUniq {
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
	t1 := time.Now()
	fmt.Println(nrUniq, t1.Sub(t0))
	return nr + 1
}
