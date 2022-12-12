package main

import (
	"fmt"
	"sort"
)

//go:generate sh -c "go run gen/gen.go < input"

func main() {
	disp()

	for i := 0; i < 20; i++ {
		for _, m := range monkies {
			m.run()
		}
		fmt.Println("Round ", i+1)
		disp()
	}

	part1()
}

func part1() {
	ins := make([]int, len(monkies))
	for i, m := range monkies {
		ins[i] = m.inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ins)))
	fmt.Println("Part 1", ins[0]*ins[1])
}

func (m *monkey) run() {
	for _, i := range m.itemQueue {
		m.op(i)
	}
	m.inspected += len(m.itemQueue)
	m.itemQueue = []int{}
}

func disp() {
	for _, m := range monkies {
		fmt.Println(m)
	}
	fmt.Println()
}

func (m monkey) String() string {
	return fmt.Sprintf("%d: %v inspected=%d", m.nr, m.itemQueue, m.inspected)
}
