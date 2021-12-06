package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}

	var fish [10]int
	for _, n := range strings.Split(s.Text(), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic("Invalid input")
		}
		fish[i]++
	}
	simulate(fish, 18)  // example
	simulate(fish, 80)  // part 1
	simulate(fish, 256) // part 2
}

func simulate(fish [10]int, days int) {
	for ; days > 0; days-- {
		fish[9] = fish[0]
		for i := 0; i < 9; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += fish[8]
	}

	nr := 0
	for i := 0; i < 9; i++ {
		nr += fish[i]
	}
	fmt.Println(nr)
}
