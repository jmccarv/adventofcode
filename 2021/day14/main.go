package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

type polymer struct {
	chain map[string]int
	rules map[string]string
	elem  map[byte]int
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

	var lo, hi int
	for _, x := range p.elem {
		if lo == 0 || x < lo {
			lo = x
		}
		hi = max(hi, x)
	}
	fmt.Println("Part1", hi-lo)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func (p polymer) addChain(s string) {
	for i := 1; i < len(s); i++ {
		p.chain[s[i-1:i+1]]++
		p.elem[s[i]]++
	}
	p.elem[s[0]]++
}

func (p polymer) addRule(pair, insert string) {
	p.rules[pair] = insert
}

func (p polymer) step() {
	n := make(map[string]int)
	for k, v := range p.chain {
		n[k] = v
	}

	//fmt.Println("Step:", p)
	for k, count := range n {
		for i := 0; i < count; i++ {
			ins, ok := p.rules[k]
			if !ok {
				continue
			}

			if p.chain[k] < 2 {
				delete(p.chain, k)
			} else {
				p.chain[k]--
			}

			p.elem[ins[0]]++

			p.chain[string(k[0])+ins]++
			p.chain[ins+string(k[1])]++
			//fmt.Printf("%s -> %s %s -> %s\n", k, string(k[0])+ins, ins+string(k[1]), p)
		}
	}
}

func (p polymer) String() string {
	nr := 0
	for _, i := range p.elem {
		nr += i
	}

	return fmt.Sprintf("%v %v len=%d", p.chain, p.elem, nr)
}
