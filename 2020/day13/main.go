package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type bus struct {
	id   uint64
	wait uint64
	ofs  uint64
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	first, err := strconv.Atoi(getLine(s))
	if err != nil {
		log.Fatal("Invalid input")
	}

	var nextBus bus
	var busses []bus
	for i, n := range strings.Split(getLine(s), ",") {
		id, err := strconv.Atoi(n)
		if err != nil {
			continue
		}

		b := bus{id: uint64(id), wait: uint64(id - first%id), ofs: uint64(i)}

		if nextBus.id == 0 || b.wait < nextBus.wait {
			nextBus = b
		}

		busses = append(busses, b)
	}
	fmt.Println("part1:", nextBus, nextBus.id*nextBus.wait)

	fmt.Println("part2:", part2(busses))
}

func getLine(s *bufio.Scanner) string {
	if !s.Scan() {
		log.Fatal("Invalid input")
	}
	return s.Text()
}

const batch = 100000

func part2(busses []bus) uint64 {
	sort.Slice(busses, func(i, j int) bool { return busses[i].id > busses[j].id })

	work := make(chan uint64)
	ret := make(chan uint64)

	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go hunt(busses, work, ret)
	}

	low := busses[0].id*2 - uint64(busses[0].ofs)
	var found uint64

out:
	for id := low; id > low-1; id += busses[0].id * batch {
		select {
		case work <- id:
		case found = <-ret:
			break out
		}
	}
	close(work)

	if found == 0 {
		return 0
	}

	for i := 1; i < runtime.GOMAXPROCS(0); i++ {
		x := <-ret
		if x > 0 && x < found {
			found = x
		}
	}

	return found
}

func hunt(busses []bus, work, ret chan uint64) {
	for start := range work {
		for id := start; id < id+batch; id += busses[0].id {
			good := true
			for _, b := range busses[1:] {
				if (id+b.ofs)%b.id != 0 {
					good = false
					break
				}
			}
			if good {
				ret <- id
				return
			}
		}
	}
	ret <- 0
}
