package main

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	nrRespondents int
	answers       map[rune]int
}

func main() {
	p1Count := 0
	p2Count := 0

	for g := range getGroups() {
		for _, a := range g.answers {
			if a == g.nrRespondents {
				p2Count++
			}
		}
		p1Count += len(g.answers)
	}

	fmt.Println("Part 1 count:", p1Count)
	fmt.Println("Part 2 count:", p2Count)
}

func getGroups() chan group {
	ch := make(chan group)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		g := group{0, make(map[rune]int)}

		for s.Scan() {
			line := s.Text()
			if len(line) == 0 && g.nrRespondents > 0 {
				ch <- g
				g = group{0, make(map[rune]int)}
				continue
			}

			for _, r := range line {
				g.answers[r]++
			}
			g.nrRespondents++
		}

		if g.nrRespondents > 0 {
			ch <- g
		}

		close(ch)
	}()

	return ch
}
