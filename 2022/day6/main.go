package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Printf("%d %d\n", check(s.Text(), 4), check(s.Text(), 14))
	}
}

func check(line string, nrUniq int) int {
	nrUniq--
	nr := 0
	for i := nrUniq; i < len(line) && nr == 0; i++ {
		var seen [123]bool // ASCII 'z' == 122
		nr = i
		for j := i - nrUniq; j <= i; j++ {
			if seen[line[j]] {
				nr = 0
				break
			}
			seen[line[j]] = true
		}
	}
	return nr + 1
}
