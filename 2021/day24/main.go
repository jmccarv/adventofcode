package main

// I had to look this one up, and took inspiration from here:
// https://www.mattkeeter.com/blog/2021-12-27-brute/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type operation int
type register int

type registers [4]int

const (
	w register = iota
	x
	y
	z
)

const (
	inp operation = iota
	add
	mul
	div
	mod
	eql
)

type operand struct {
	literal bool
	lval    int
	reg     register
}

type instruction struct {
	op   operation
	args [2]operand
}

type program struct {
	ins   []instruction
	text  []string
	regs  registers
	input []int
	pc    int
}

var regMap map[string]register
var opMap map[string]operation

type state struct {
	regs     registers
	min, max int
}

func init() {
	//regMap := make(map[rune], register, 4)
	regMap = map[string]register{
		"w": w,
		"x": x,
		"y": y,
		"z": z,
	}

	opMap = map[string]operation{
		"inp": inp,
		"add": add,
		"mul": mul,
		"div": div,
		"mod": mod,
		"eql": eql,
	}
}

func main() {
	prg := program{}
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		s := strings.Fields(input.Text())
		prg.parseInstruction(s)
	}

	if len(os.Args) > 1 {
		fmt.Println(prg.run(os.Args[1]))
		fmt.Println(prg.regs)
	} else {
		f, _ := os.Create("cpu.prof")
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
		solve(prg)
	}
}

type stateList []state

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

type codeBlock []instruction

// Note that this uses a lot of RAM. May be a better way, not sure.
// Initial version (no parallelism, etc) in about 1m31s on my puzzle input
// Now runs in about 26sec with better RAM usage characteristics
func solve(prg program) {
	var blocks []codeBlock
	var block codeBlock
	for _, i := range prg.ins {
		switch i.op {
		case inp:
			if len(block) > 0 {
				blocks = append(blocks, block)
			}
			block = codeBlock{}
		}
		block = append(block, i)
	}
	if len(block) > 0 {
		blocks = append(blocks, block)
	}

	nrinp := 0
	states := stateList{state{}}
	statesO := states
	ns := make([]stateList, 9)
	ns[0] = append(ns[0], state{})

	for _, block = range blocks {
		// first instruction is an inp
		if block[0].op != inp {
			panic("Invalid code block!")
		}

		ra := block[0].args[0].reg
		nrinp++
		fmt.Println()
		t1 := time.Now()
		fmt.Println("input #", nrinp, " -- ")

		fmt.Println("  sorting...")
		var wg sync.WaitGroup
		wg.Add(9)
		for j := 0; j < 9; j++ {
			go func(ns stateList) {
				defer wg.Done()
				// Inp is about to change register r, so we can set it to
				// zero in every state and then deduplicate our list
				for i := range ns {
					ns[i].regs[ra] = 0
				}

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
				for i := range states {
					ns[i] = states[i]
					ns[i].regs[ra] = j
					ns[i].max = ns[i].max*10 + j
					ns[i].min = ns[i].min*10 + j
				}
				for _, i := range block[1:] {
					ra := i.args[0].reg
					ob := i.args[1]
					switch i.op {
					case add:
						for st := range ns {
							ns[st].regs[ra] += ns[st].regs.val(ob)
						}
					case mul:
						for st := range ns {
							ns[st].regs[ra] *= ns[st].regs.val(ob)
						}
					case div:
						for st := range ns {
							ns[st].regs[ra] /= ns[st].regs.val(ob)
						}
					case mod:
						for st := range ns {
							ns[st].regs[ra] %= ns[st].regs.val(ob)
						}
					case eql:
						for st := range ns {
							if ns[st].regs[ra] == ns[st].regs.val(ob) {
								ns[st].regs[ra] = 1
							} else {
								ns[st].regs[ra] = 0
							}
						}
					}
				}
			}(j+1, ns[j])
		}
		wg.Wait()
		fmt.Println("  processed in", time.Now().Sub(t1))
		//states = ns
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

func (r registers) val(o operand) int {
	if o.literal {
		return o.lval
	}
	return r[o.reg]
}

func (p *program) parseInstruction(s []string) {
	p.text = append(p.text, fmt.Sprintf("%v", s))
	var r1, r2 register
	var l2 int
	var ok1, ok2 bool

	if r1, ok1 = regMap[s[1]]; !ok1 {
		panic(fmt.Sprintf("Invalid register '%s' for %v'", s[1], s))
	}

	if len(s) > 2 {
		if r2, ok2 = regMap[s[2]]; !ok2 {
			l2 = atoi(s[2])
		}
	}

	i := instruction{op: opMap[s[0]]}
	i.args[0].reg = r1
	if ok2 {
		i.args[1].reg = r2
	} else {
		i.args[1].literal = true
		i.args[1].lval = l2
	}
	p.ins = append(p.ins, i)
}

func (p *program) nextInput() (int, error) {
	if len(p.input) == 0 {
		return 0, errors.New("Not enough input values")
	}

	ret := p.input[0]
	if ret == 0 {
		return 0, errors.New("Input cannot be zero")
	}

	if len(p.input) > 0 {
		p.input = p.input[1:]
	} else {
		p.input = []int{}
	}

	return ret, nil
}

func (p *program) val(o operand) int {
	if o.literal {
		return o.lval
	} else {
		return p.regs[o.reg]
	}
}

// returns true if no exception was encountered
func (p *program) run(inNum string) (ok bool, err error) {
	p.pc = 0
	p.regs = registers{0, 0, 0, 0}
	p.input = make([]int, 0, len(inNum))

	for _, c := range []byte(inNum) {
		if c == '0' {
			return false, nil
		}
		p.input = append(p.input, int(c-'0'))
	}

	for _, i := range p.ins {
		r := i.args[0].reg
		v := p.val(i.args[1])
		switch i.op {
		case inp:
			p.regs[r], err = p.nextInput()
			if err != nil {
				return
			}
		case add:
			p.regs[r] += v
		case mul:
			p.regs[r] *= v
		case div:
			if v == 0 {
				return false, errors.New("Divide by zero")
			}
			p.regs[r] /= v
		case mod:
			if p.regs[r] < 0 {
				return false, errors.New("Cannot mod 0")
			}
			if v <= 0 {
				return false, errors.New("Cannot mod by 0")
			}
			p.regs[r] %= v
		case eql:
			if p.regs[r] == v {
				p.regs[r] = 1
			} else {
				p.regs[r] = 0
			}
		}
		//fmt.Println(p)
		p.pc++
	}
	return true, nil
}

func (p *program) String() string {
	text := ""
	if p.pc < len(p.text) {
		text = p.text[p.pc]
	}
	return fmt.Sprintf("pc: %d (%s) %v", p.pc, text, p.regs)
}

func (r registers) String() string {
	return fmt.Sprintf("R[w:%d x:%d y:%d z:%d]", r[w], r[x], r[y], r[z])
}

func atoi(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Invalid integer '%s'", s))
	}
	return ret
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
