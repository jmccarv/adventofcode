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
	var p1 int
	re := regexp.MustCompile(`:\s+((?:\d+\s+)+)\|\s+((?:\d+\s*)+)$`)
	s := bufio.NewScanner(os.Stdin)

	//var copies = make(map[int]int)
	//id := 0
	for s.Scan() {
		//id++
		var winning = make(map[int]struct{})
		matches := re.FindAllStringSubmatch(s.Text(), -1)
		for _, nr := range strings.Fields(matches[0][1]) {
			val, _ := strconv.Atoi(nr)
			winning[val] = struct{}{}
		}
		nrWinning := 0
		for _, nr := range strings.Fields(matches[0][2]) {
			val, _ := strconv.Atoi(nr)
			if _, ok := winning[val]; ok {
				nrWinning++
			}
		}
		if nrWinning > 0 {
			p1 += int(math.Pow(2, float64(nrWinning-1)))
		}
	}
	fmt.Println("Part 1", p1)
}
