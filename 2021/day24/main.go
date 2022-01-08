package main

// I had to look this one up, and took inspiration from here:
// https://www.mattkeeter.com/blog/2021-12-27-brute/

//go:generate sh -c "go run gen/gen.go < input"

import (
	"fmt"
	"os"
	//"runtime/pprof"
	"sort"
	"sync"
	"time"
)

type registers [4]int

type state struct {
	regs     registers
	min, max int
}

type stateList []state

func main() {
	if len(os.Args) > 1 {
		interpret()
		return
	}

	/*
		f, _ := os.Create("cpu.prof")
		defer f.Close()
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
	*/
	solve()
}

func (a registers) Less(b registers) bool {
	for i := 0; i < 4; i++ {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	return false
}

func (l stateList) Less(a, b int) bool {
	for i := 0; i < 4; i++ {
		if l[a].regs[i] < l[b].regs[i] {
			return true
		}
		if l[a].regs[i] > l[b].regs[i] {
			return false
		}
	}
	return false
}

func (l stateList) Len() int {
	return len(l)
}

func (l stateList) Swap(a, b int) {
	l[a], l[b] = l[b], l[a]
}

// Note that this uses a lot of RAM. May be a better way, not sure.
// Initial version (no parallelism, etc) in about 1m31s on my puzzle input
// After some improvements and parallelization,  runs in about 26sec with better RAM usage characteristics
// And now with code generation, runs in about 12.5sec
func solve() {
	states := stateList{state{}}
	statesO := states
	ns := make([]stateList, 9)
	ns[0] = append(ns[0], state{})

	preSorts := getPreSorts()
	for i, block := range getBlocks() {
		fmt.Println()
		t1 := time.Now()
		fmt.Println("input #", i+1, " -- ")

		fmt.Println("  sorting...")
		var wg sync.WaitGroup
		wg.Add(9)
		for j := 0; j < 9; j++ {
			go func(ns stateList) {
				defer wg.Done()
				preSorts[i](ns)
				sort.Sort(ns)
			}(ns[j])
		}
		wg.Wait()
		fmt.Println("  sort done in", time.Now().Sub(t1))

		// Remove any duplicate states before splitting them
		// We do this at the same time we merge all our ns slices into the final states slice
		fmt.Println("      merging:", len(ns[0])*9)
		t1 = time.Now()
		next := state{regs: registers{99999999999999, 99999999999999, 99999999999999, 99999999999999}, min: 99999999999999}
		cur := next

		// Find our initial cur value -- the one with smallest registers set
		for _, x := range ns {
			if len(x) > 0 && x[0].regs.Less(cur.regs) {
				cur = x[0]
			}
		}

		// revert states back to an empty slice without having to reallocate anything
		states = statesO[0:0]
		done := false
		x := 0
		for !done {
			done = true
			states = append(states, cur)
			for i := 0; i < 9; i++ {
				if len(ns[i]) == 0 {
					continue
				}

				j := 0
				for j < len(ns[i]) && ns[i][j].regs == cur.regs {
					states[x].min = min(states[x].min, ns[i][j].min)
					states[x].max = max(states[x].max, ns[i][j].max)
					j++
				}
				if j < len(ns[i]) {
					if ns[i][j].regs.Less(next.regs) {
						next = ns[i][j]
					}
					ns[i] = ns[i][j:]
					done = false
				} else {
					ns[i] = ns[i][0:0]
				}
			}
			cur = next
			next = state{regs: registers{99999999999999, 99999999999999, 99999999999999, 99999999999999}, min: 9999999999999}
			x++
		}
		fmt.Println("  after merge:", len(states))
		fmt.Println("  merged in", time.Now().Sub(t1))

		// We now have a list of unique states to operate on
		// There are nine possible values we might input, so each of our existing
		// states split into 9 new states
		fmt.Println("  processing commands...")
		t1 = time.Now()
		wg = sync.WaitGroup{}
		wg.Add(9)
		for j := 0; j < 9; j++ {
			ns[j] = make([]state, len(states))
			go func(j int, ns []state) {
				defer wg.Done()
				block(j, states, ns)
			}(j+1, ns[j])
		}
		wg.Wait()
		fmt.Println("  processed in", time.Now().Sub(t1))
	}

	p1 := 0
	p2 := 99999999999999
	for _, sl := range ns {
		for _, s := range sl {
			if s.regs[z] == 0 {
				fmt.Println(s)
				p1 = max(p1, s.max)
				p2 = min(p2, s.min)
			}
		}
	}
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
