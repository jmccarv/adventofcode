package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// We are offloading a ship, its cargo hold contains stacks of crates
// Crates are labeled with a letter
// We are given instructions for how to move crates from one stack to another
// Crates are moved one at a time

type stack []byte
type hold []stack

func main() {
	var stacks hold
	s := bufio.NewScanner(os.Stdin)

	// Read the starting crate stack configuration
	for s.Scan() {
		line := s.Text()
		if line[1] == '1' {
			break
		}

		sidx := 0
		for ; len(line) > 1; line = line[4:] {
			if sidx > len(stacks)-1 {
				stacks = append(stacks, stack{})
			}
			fmt.Printf("%c ", line[1])
			if line[1] != ' ' {
				stacks[sidx] = append(stacks[sidx], line[1])
			}
			sidx++
			if len(line) < 4 {
				break
			}
		}
		fmt.Println()
	}
	// Reverse the stacks because the ends of the arrays should be the top of the stacks
	// And by using append() while reading, we were essentially pushing them in reverse order
	for _, s := range stacks {
		s.reverse()
	}

	// Now read the instructions for how to move crates
	re := regexp.MustCompile(`move \d+ from \d+ to \d+`)
	for s.Scan() {
		var nr, from, to int
		if !re.MatchString(s.Text()) {
			continue
		}
		n, _ := fmt.Sscanf(s.Text(), "move %d from %d to %d", &nr, &from, &to)
		if n > 0 {
			fmt.Printf("Move %d from %d to %d\n", nr, from, to)
			for i := 0; i < nr; i++ {
				stacks[to-1].push(stacks[from-1].pop())
			}
		}
	}
	for _, s := range stacks {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()
}

func (s stack) reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (s *stack) pop() (b byte) {
	if len(*s) == 0 {
		return
	}

	b = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return b
}

func (s *stack) push(b byte) {
	*s = append(*s, b)
}
