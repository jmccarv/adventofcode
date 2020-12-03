package main

import (
	"bufio"
	"fmt"
	"os"
)

type slope struct {
	right int
	down  int
	trees int
}

func main() {
	var lines []string
	slopes := []slope{
		{1, 1, 0},
		{3, 1, 0},
		{5, 1, 0},
		{7, 1, 0},
		{1, 2, 0},
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	ans := 1
	for _, s := range slopes {
		ans *= s.count(lines)
		fmt.Println(s)
	}

	fmt.Println("Product:", ans)
}

func (s slope) String() string {
	return fmt.Sprintf("r%dd%d %d trees", s.right, s.down, s.trees)
}

func (s *slope) count(lines []string) int {
	var x int

	s.trees = 0
	for y := s.down; y < len(lines); y += s.down {
		x += s.right
		line := lines[y]

		if x >= len(line) {
			x %= len(line)
		}

		if line[x] == '#' {
			s.trees++
		}
	}

	return s.trees
}
