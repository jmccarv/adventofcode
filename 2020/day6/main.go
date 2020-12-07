package main

import (
	"bufio"
	"fmt"
	"os"
)

type group map[rune]bool

func main() {
	count := 0
	chGroup := getGroups()
	for g := range chGroup {
		count += len(g)
	}
	fmt.Println("Total count:", count)
}

func getGroups() chan group {
	ch := make(chan group)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		g := make(group)

		for s.Scan() {
			line := s.Text()
			if len(line) == 0 && len(g) > 0 {
				ch <- g
				g = make(group)
				continue
			}

			for _, r := range line {
				g[r] = true
			}
		}

		if len(g) > 0 {
			ch <- g
		}

		close(ch)
	}()

	return ch
}
