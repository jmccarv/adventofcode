package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	cmd string
	val int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var lines []line

	for s.Scan() {
		f := strings.Fields(s.Text())
		d, err := strconv.Atoi(f[1])
		if err != nil {
			panic("Invalid input!")
		}

		lines = append(lines, line{f[0], d})
	}

	part1(lines)
	part2(lines)
}

func part1(lines []line) {
	var x, y int

	for _, l := range lines {
		switch l.cmd {
		case "forward":
			x += l.val
		case "down":
			y += l.val
		case "up":
			y -= l.val
		}
	}
	fmt.Println(x * y)
}

func part2(lines []line) {
	var x, y, a int

	for _, l := range lines {
		switch l.cmd {
		case "forward":
			x += l.val
			y += a * l.val
		case "down":
			a += l.val
		case "up":
			a -= l.val
		}
	}
	fmt.Println(x * y)
}
