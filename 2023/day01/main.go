package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	var p1, p2 int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		p1 += decode(s.Text(), p1Digit)
		p2 += decode(s.Text(), p2Digit)
	}
	fmt.Println("part1", p1)
	fmt.Println("part2", p2)
}

func decode(text string, digit func(s string) int) int {
	var l, r int
	for i := range text {
		if d := digit(text[i:]); d > 0 {
			l = d
			break
		}
	}
	for i := len(text) - 1; i >= 0; i-- {
		if d := digit(text[i:]); d > 0 {
			r = d
			break
		}
	}
	return l*10 + r
}

func p1Digit(text string) int {
	if text[0] >= '0' && text[0] <= '9' {
		return int(text[0] - '0')
	}
	return -1
}

func p2Digit(text string) int {
	if n := p1Digit(text); n >= 0 {
		return n
	}
	for i, num := range nums {
		if strings.Index(text, num) == 0 {
			return i + 1
		}
	}
	return -1
}
