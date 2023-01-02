package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// We are offloading a ship, its cargo hold contains stacks of crates
// Crates are labeled with a letter
// We are given instructions for how to move crates from one stack to another
// Crates are moved one at a time

type stack []byte
type hold []stack

type move struct {
	nr   string
	from string
	to   string
}

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
			if line[1] != ' ' {
				stacks[sidx] = append(stacks[sidx], line[1])
			}
			sidx++
			if len(line) < 4 {
				break
			}
		}
	}

	for _, s := range stacks {
		// Reverse the stacks because the ends of the arrays should be the top of the stacks
		s.reverse()
	}

	// Now generate C struct with stacks set to their initial configuration
	doStack("p1", stacks)
	doStack("p2", stacks)

	fmt.Printf("#define NR_STACKS %d\n", len(stacks))

	// Now the move instructions
	var nr, from, to []string
	for s.Scan() {
		var m move
		n, _ := fmt.Sscanf(s.Text(), "move %s from %s to %s", &m.nr, &m.from, &m.to)
		if n < 3 {
			continue
		}
		nr = append(nr, m.nr)
		from = append(from, m.from)
		to = append(to, m.to)
	}

	fmt.Printf(`
struct moves_t {
	unsigned char nr[%d];
	unsigned char from[%d];
	unsigned char to[%d];
};

struct moves_t moves = {
`, len(nr), len(nr), len(nr))

	fmt.Printf("{%s},\n", strings.Join(nr, ","))
	fmt.Printf("{%s-1},\n", strings.Join(from, "-1,"))
	fmt.Printf("{%s-1},\n", strings.Join(to, "-1,"))
	fmt.Println("};")
	fmt.Printf("const unsigned int nr_moves = %d;\n", len(nr))
}

func doStack(p string, stacks hold) {
	fmt.Printf("unsigned char %sstacks[%d][50] = {\n", p, len(stacks))
	for _, s := range stacks {
		fmt.Printf("{'%s'},\n", strings.Join(strings.Split(string(s), ""), "','"))
	}
	fmt.Printf("};\n")

	fmt.Printf("unsigned char *%ssp[%d] = {\n", p, len(stacks))
	//x := ""
	for i, s := range stacks {
		//x += fmt.Sprintf(",%d", len(s)-1)
		fmt.Printf("(unsigned char *)&(%sstacks[%d]) + %d,\n", p, i, len(s))
	}
	//fmt.Printf("%s};\n", x[1:])
	fmt.Printf("};\n")
}

func (s stack) reverse() {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
