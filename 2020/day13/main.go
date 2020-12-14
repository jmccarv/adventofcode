package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	id   int
	wait int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	first, err := strconv.Atoi(getLine(s))
	if err != nil {
		log.Fatal("Invalid input")
	}

	var b bus
	for _, n := range strings.Split(getLine(s), ",") {
		id, err := strconv.Atoi(n)
		if err != nil {
			continue
		}

		x := id - first%id
		if b.id == 0 || x < b.wait {
			b.id = id
			b.wait = x
		}
	}
	fmt.Println("part1:", b, b.id*b.wait)

}

func getLine(s *bufio.Scanner) string {
	if !s.Scan() {
		log.Fatal("Invalid input")
	}
	return s.Text()
}
