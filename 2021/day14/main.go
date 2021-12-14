package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type polymer struct {
	// Count of each element pair in the chain, where pairs overlap.
	// ABBBC is four pairs: [AB:1 BB:2 BC:1]
	pairs map[string]int

	// Maps element pairs to an element to be inserted between the pair,
	// creating two new elements. See puzzle instructions :)
	rules map[string]string

	// Tracks the count of each element in the resulting chain
	elems map[byte]int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	p := newPolymer()

	for s.Scan() {
		f := strings.Fields(s.Text())

		if len(f) == 1 && len(f[0]) > 0 {
			p.addChain(s.Text())
		} else if len(f) == 3 {
			p.addRule(f[0], f[2])
		}
	}

	for i := 0; i < 10; i++ {
		p.step()
	}
	fmt.Println("Part1 (10 steps)", p.score())

	for i := 0; i < 30; i++ {
		p.step()
	}
	fmt.Println("Part2 (40 steps)", p.score())
	//fmt.Println(p)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newPolymer() polymer {
	return polymer{
		make(map[string]int),
		make(map[string]string),
		make(map[byte]int),
	}
}

func (p polymer) score() int {
	var lo, hi int
	for _, x := range p.elems {
		if lo == 0 || x < lo {
			lo = x
		}
		hi = max(hi, x)
	}
	return hi - lo
}

func (p polymer) addChain(s string) {
	for i := 1; i < len(s); i++ {
		p.pairs[s[i-1:i+1]]++
		p.elems[s[i]]++
	}
	p.elems[s[0]]++
}

func (p polymer) addRule(pair, insert string) {
	p.rules[pair] = insert
}

func (p polymer) step() {
	n := make(map[string]int)
	for k, v := range p.pairs {
		n[k] = v
	}

	//fmt.Println("Step:", p)
	for k, count := range n {
		// we have count k pairs so we are doing this same
		// operation count times.

		ins, ok := p.rules[k] // should never fail with valid input
		if !ok {
			continue
		}
		// ins is the new element we're inserting between pair k,
		// which will replace pair k with two new pairs instead

		// replacing pair k with two new ones so update our count
		if p.pairs[k] <= count {
			delete(p.pairs, k)
		} else {
			p.pairs[k] -= count
		}

		// now add in the new pairs, we're adding count of them
		p.pairs[string(k[0])+ins] += count
		p.pairs[ins+string(k[1])] += count

		// and finally we've added count new ins elements to the final chain; count them
		p.elems[ins[0]] += count
	}
}

func (p polymer) String() string {
	nr := 0
	for _, i := range p.elems {
		nr += i
	}
	return fmt.Sprintf("%v %v len=%d", p.pairs, p.elems, nr)
}
