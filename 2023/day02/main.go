package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var p1max map[string]int = map[string]int{"red": 12, "green": 13, "blue": 14}
	var id, p1, p2 int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var p2max map[string]int = map[string]int{"red": 0, "green": 0, "blue": 0}
		game := strings.Split(s.Text(), ":")
		fmt.Sscanf(game[0], "Game %d", &id)
		good := true
		for _, c := range strings.FieldsFunc(game[1], func(r rune) bool { return r == ',' || r == ';' }) {
			var nr int
			var color string
			fmt.Sscanf(c, "%d %s", &nr, &color)
			//fmt.Println("Game", id, c, nr, color)
			p2max[color] = max(p2max[color], nr)
			if nr > p1max[color] {
				good = false
			}
		}
		//fmt.Println(id, good, p2max)
		if good {
			p1 += id
		}
		pwr := 1
		for _, v := range p2max {
			pwr *= v
		}
		p2 += pwr
	}
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
}
