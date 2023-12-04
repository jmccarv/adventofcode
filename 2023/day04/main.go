package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var p1, p2 int
	re := regexp.MustCompile(`:\s+((?:\d+\s+)+)\|\s+((?:\d+\s*)+)$`)
	s := bufio.NewScanner(os.Stdin)

	var copies = make(map[int]int)
	id := 0
	for s.Scan() {
		id++
		var winning = make(map[int]struct{})
		matches := re.FindStringSubmatch(s.Text())
		for _, nr := range strings.Fields(matches[1]) {
			val, _ := strconv.Atoi(nr)
			winning[val] = struct{}{}
		}
		wins := 0
		for _, nr := range strings.Fields(matches[2]) {
			val, _ := strconv.Atoi(nr)
			if _, ok := winning[val]; ok {
				wins++
			}
		}
		p2 += copies[id] + 1
		if wins > 0 {
			p1 += int(math.Pow(2, float64(wins-1)))
			for i := id + 1; i < id+1+wins; i++ {
				copies[i] += copies[id] + 1
			}
		}
	}
	fmt.Println("Part 1", p1)
	fmt.Println("Part 2", p2)
}
