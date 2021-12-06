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

	var fish [10]int // each element is the count of fish at that timer value
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
		// All fish at timer 0 will spawn a new fish with its timer=9 (will be corrected to 8 next)
		fish[9] = fish[0]

		// Now decrement all fish's counters by 1
		for i := 0; i < 9; i++ {
			fish[i] = fish[i+1]
		}

		// Finally, all fish that were at timer=0 reset to timer=6 and since they
		// spawned new fish at timer=8 we can just use the count from timer=8
		fish[6] += fish[8]
	}

	nr := 0
	for i := 0; i < 9; i++ {
		nr += fish[i]
	}
	fmt.Println(nr)
}
