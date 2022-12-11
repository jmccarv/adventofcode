package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ages [10]int

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}

	var fish ages // each element is the count of fish at that timer value
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

func (fish ages) count() int {
	nr := 0
	for i := 0; i < 9; i++ {
		nr += fish[i]
	}
	return nr
}

func simulate(fish ages, days int) {
	for d := 1; d <= days; d++ {
		// All fish at timer 0 will spawn a new fish with its timer=9 (will be corrected to 8 next)
		fish[9] = fish[0]

		// Now decrement all fish's timers by 1
		for i := 0; i < 9; i++ {
			fish[i] = fish[i+1]
		}

		// Finally, all fish that were at timer=0 reset to timer=6 and since they
		// spawned new fish at timer=8 we can just use the count from timer=8
		fish[6] += fish[8]
		//fmt.Printf("%d %d\n", d, fish.count())
	}
	fmt.Printf("%d %d\n", days, fish.count())
}
