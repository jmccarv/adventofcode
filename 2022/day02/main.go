package main

import (
	"bufio"
	"fmt"
	"os"
)

var scores = map[string]int{
	"AX": 3, "AY": 6, "AZ": 0,
	"BX": 0, "BY": 3, "BZ": 6,
	"CX": 6, "CY": 0, "CZ": 3,
}

var xform = map[string]string{
	"AX": "AZ", "AY": "AX", "AZ": "AY",
	"BX": "BX", "BY": "BY", "BZ": "BZ",
	"CX": "CY", "CY": "CZ", "CZ": "CX",
}

func main() {
	var games []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		games = append(games, string(line[0])+string(line[2]))
	}
	run(games)

	for i := range games {
		games[i] = xform[games[i]]
	}
	run(games)
}

func run(games []string) {
	total := 0
	for _, g := range games {
		total += scores[g] + 3 - ('Z' - int(g[1]))
	}
	fmt.Println(total)
}
