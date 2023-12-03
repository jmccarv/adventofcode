package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var p1max map[string]int = map[string]int{"red": 12, "green": 13, "blue": 14}
	re := regexp.MustCompile(`(\d+)\s(\w+)`)
	var id, p1, p2 int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		id++
		mx := make(map[string]int)
		for _, m := range re.FindAllStringSubmatch(s.Text(), -1) {
			nr, _ := strconv.Atoi(m[1])
			mx[m[2]] = max(mx[m[2]], nr)
			fmt.Println(id, m[1], m[2], mx[m[2]])
		}
		p2tmp := 1
		good := true
		for color, nr := range mx {
			if nr > p1max[color] {
				good = false
			}
			p2tmp *= nr
		}
		if good {
			p1 += id
		}
		p2 += p2tmp
	}
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
}
