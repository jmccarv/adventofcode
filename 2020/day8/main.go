package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type memCell struct {
	op      string
	arg     int
	visited bool
}

type registers struct {
	pc  int // "program counter" -- offset into memory containing next instruction to execute
	acc int // accumulator -- initialized to zero
}

type machine struct {
	mem  []memCell
	regs registers
	ops  map[string]func(arg int)
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	m := newMachine()

	for s.Scan() {
		f := strings.Fields(s.Text())
		if len(f) != 2 {
			log.Fatal("Invalid input", s.Text())
		}

		arg, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatal("Invalid arg", f[0], " in", s.Text())
		}

		// place our new instruction into memory
		m.addInstruction(f[0], arg)
	}

	m.run()
	fmt.Println("Part1:", m)

	part2(m)
	fmt.Printf("Part2: %+v\n", m.regs)
}

func part2(m *machine) {
	// oh boy ... brute force here we go

	orig := make([]memCell, len(m.mem))
	copy(orig, m.mem)

	for i, c := range orig {
		// restore memory from our copy
		copy(m.mem, orig)

		// Check if this is an instruction we are interested in
		switch c.op {
		case "nop":
			m.mem[i].op = "jmp"
		case "jmp":
			m.mem[i].op = "nop"

		default: // not interested
			continue
		}

		// We've modified the code, try it out
		m.run()
		if m.regs.pc == len(m.mem) {
			// hooray, we fixed it!
			return
		}
	}
}

func newMachine() *machine {
	m := &machine{}

	m.ops = map[string]func(arg int){
		"nop": func(arg int) {},
		"acc": func(arg int) { m.regs.acc += arg },
		"jmp": func(arg int) { m.regs.pc += arg - 1 },
	}

	return m
}

func (m *machine) addInstruction(op string, arg int) {
	m.mem = append(m.mem, memCell{op: op, arg: arg})
}

func (m *machine) init() {
	m.regs = registers{}
	for i := range m.mem {
		m.mem[i].visited = false
	}
}

// run the program in memory in our machine
// returns if a loop is detected or if execution would go outside memory
func (m *machine) run() {
	m.init()

	for {
		if m.regs.pc >= len(m.mem) || m.regs.pc < 0 {
			// fell off the end of memory
			break
		}

		// fetch next instruction
		instr := &m.mem[m.regs.pc]

		// increment program counter to next instruction to fetch
		m.regs.pc++

		if instr.visited {
			// we're calling this an infinite loop, stop here
			break
		}
		instr.visited = true

		// execute instruction
		m.ops[instr.op](instr.arg)
	}
}

func (m *machine) String() string {
	return fmt.Sprintf("%+v", m.regs)
}
