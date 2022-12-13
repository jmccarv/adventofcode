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

	nr, elp := run(p1monkeys, 20)
	fmt.Println("Part1:", nr, " in", elp)

	nr, elp = run(p2monkeys, 10000)
	fmt.Println("Part2:", nr, " in", elp)

	fmt.Println("Total Time:", time.Now().Sub(t0))
}

func run(monkeys []*monkey, rounds int) (int, time.Duration) {
	t0 := time.Now()
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			m.run()
		}
	}

	ins := make([]int, len(monkeys))
	for i, m := range monkeys {
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

func disp(monkeys []*monkey) {
	for _, m := range monkeys {
		fmt.Println(m)
	}
	fmt.Println()
}

func (m monkey) String() string {
	return fmt.Sprintf("%d: %v inspected=%d", m.nr, m.items, m.inspected)
}
