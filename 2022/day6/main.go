package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	// Read the starting crate stack configuration
	for s.Scan() {
		line := s.Text()
		nr := 0
		for i := 3; i < len(line); i++ {
			var seen [256]bool
			nr = i
			for j := i - 3; j <= i; j++ {
				if seen[line[j]] {
					nr = 0
					break
				}
				seen[line[j]] = true
			}
			if nr > 0 {
				break
			}
		}
		fmt.Println(nr + 1)
	}
}
