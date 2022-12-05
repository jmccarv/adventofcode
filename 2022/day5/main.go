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
	// While we're at it, crate a copy of each stack to use for part 2
	var p2stacks hold = make([]stack, 0, len(stacks))
	for _, s := range stacks {
		s.reverse()

		var t stack = make([]byte, len(s))
		copy(t, s)
		p2stacks = append(p2stacks, t)
	}

	// Now read and execute the instructions to move crates
	re := regexp.MustCompile(`move \d+ from \d+ to \d+`)
	for s.Scan() {
		if !re.MatchString(s.Text()) {
			continue
		}

		var nr, from, to int
		fmt.Sscanf(s.Text(), "move %d from %d to %d", &nr, &from, &to)

		// part 1
		for i := 0; i < nr; i++ {
			stacks[to-1].push(stacks[from-1].pop(1))
		}

		// part2
		p2stacks[to-1].push(p2stacks[from-1].pop(nr))
	}

	stacks.dispTop()
	p2stacks.dispTop()
}

func (s stack) reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (s *stack) pop(nr int) (b []byte) {
	b = (*s)[len(*s)-nr:]
	*s = (*s)[:len(*s)-nr]
	return b
}

func (s *stack) push(b []byte) {
	*s = append(*s, b...)
}

func (h hold) dispTop() {
	for _, s := range h {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Println()
}

func (h hold) dump() {
	for _, s := range h {
		fmt.Println(string(s))
	}
}
