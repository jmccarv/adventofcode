package main

// I had to look this one up, and took inspiration from here:
// https://www.mattkeeter.com/blog/2021-12-27-brute/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
		solve(prg)
	}
}

// Note that this uses a lot of RAM. May be a better way, not sure.
// Ran in about 1m31s on my puzzle input
func solve(prg program) {
	nrinp := 0
	states := []state{state{}}

	for _, i := range prg.ins {
		ra := i.args[0].reg
		ob := i.args[1]
		switch i.op {
		case inp:
			nrinp++
			fmt.Println()
			fmt.Println("input #", nrinp)
			// Inp is about to change register r, so we can set it to
			// zero in every state and then deduplicate our list
			for i := range states {
				states[i].regs[ra] = 0
			}

			// Sort so we can do a single pass to deduplicate the list
			sort.Slice(states, func(a, b int) bool {
				return states[a].regs.less(states[b].regs)
			})

			// Remove any duplicate states before splitting them
			fmt.Println("Before", len(states))
			i := 0
			for j := 1; j < len(states); j++ {
				if states[i].regs == states[j].regs {
					states[i].min = min(states[i].min, states[j].min)
					states[i].max = max(states[i].max, states[j].max)
					continue
				}
				i++
				states[i] = states[j]
			}
			states = states[0 : i+1]
			fmt.Println(" After", len(states))

			// There are nine possible values we might input, so each of our existing
			// states split into 9 new states
			//ns := make([]state, len(states) * 9)
			newStates := make([]state, 0, len(states)*9)
			for i := range states {
				for j := 1; j <= 9; j++ {
					s := states[i]
					s.regs[ra] = j
					s.max = s.max*10 + j
					s.min = s.min*10 + j
					newStates = append(newStates, s)
				}
			}
			states = newStates
		case add:
			for i := range states {
				states[i].regs[ra] += states[i].regs.val(ob)
			}
		case mul:
			for i := range states {
				states[i].regs[ra] *= states[i].regs.val(ob)
			}
		case div:
			for i := range states {
				states[i].regs[ra] /= states[i].regs.val(ob)
			}
		case mod:
			for i := range states {
				states[i].regs[ra] %= states[i].regs.val(ob)
			}
		case eql:
			for i := range states {
				if states[i].regs[ra] == states[i].regs.val(ob) {
					states[i].regs[ra] = 1
				} else {
					states[i].regs[ra] = 0
				}
			}
		}
		prg.pc++
	}

	p1 := 0
	p2 := 99999999999999
	for _, s := range states {
		if s.regs[z] == 0 {
			fmt.Println(s)
			p1 = max(p1, s.max)
			p2 = min(p2, s.min)
		}
	}
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	return
}

func (a registers) less(b registers) bool {
	for i := range a {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	return false
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
