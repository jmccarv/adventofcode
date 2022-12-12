package main

import (
	"fmt"
	"sort"
	"time"
)

// Support 'go generate' without using the Makefile
//go:generate sh -c "go run gen/gen.go < input"

func main() {
	t0 := time.Now()

	nr, elp := run(p1monkies, 20)
	fmt.Println("Part1:", nr, " in", elp)

	nr, elp = run(p2monkies, 10000)
	fmt.Println("Part2:", nr, " in", elp)

	fmt.Println("Main: ", time.Now().Sub(t0))
}

func run(monkies []*monkey, rounds int) (int, time.Duration) {
	t0 := time.Now()
	for i := 0; i < rounds; i++ {
		for _, m := range monkies {
			m.run()
		}
	}

	ins := make([]int, len(monkies))
	for i, m := range monkies {
		ins[i] = m.inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ins)))

	return ins[0] * ins[1], time.Now().Sub(t0)
}

func (m *monkey) run() {
	for _, i := range m.items {
		m.op(i)
	}
	m.inspected += len(m.items)
	m.items = []int{}
}

func disp(monkies []*monkey) {
	for _, m := range monkies {
		fmt.Println(m)
	}
	fmt.Println()
}

func (m monkey) String() string {
	return fmt.Sprintf("%d: %v inspected=%d", m.nr, m.items, m.inspected)
}
