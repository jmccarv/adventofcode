package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation func()
type register int

type instruction func() error

type registers [4]int

const (
	w register = iota
	x
	y
	z
)

type program struct {
	ins   []instruction
	text  []string
	regs  registers
	input []int
	pc    int
}

var regMap map[string]register

func init() {
	//regMap := make(map[rune], register, 4)
	regMap = map[string]register{
		"w": w,
		"x": x,
		"y": y,
		"z": z,
	}
}

func main() {
	prg := &program{}
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		s := strings.Fields(input.Text())
		prg.parseInstruction(s)
	}

	if len(os.Args) > 1 {
		fmt.Println(prg.run(os.Args[1]))
		fmt.Println(prg)
	} else {
		part1(prg)
	}

	//if program.run() {
	//}
}

func part1(prg *program) {
	j := 99999999999999
	for i := 99999999999999; i > 9999999999999; i-- {
		if !prg.run(fmt.Sprintf("%d", i)) {
			continue
		}
		if j > i+1000000 {
			fmt.Println(i, prg)
			j = i
		}
		if prg.regs[z] == 0 {
			fmt.Println(i)
			return
		}
	}
	return
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

	switch s[0] {
	case "inp":
		p.ins = append(p.ins, func() (err error) {
			p.regs[r1], err = p.nextInput()
			return
		})
	case "add":
		if ok2 {
			p.ins = append(p.ins, func() (err error) {
				p.regs[r1] += p.regs[r2]
				return
			})
		} else {
			p.ins = append(p.ins, func() (err error) {
				p.regs[r1] += l2
				return
			})
		}
	case "mul":
		if ok2 {
			p.ins = append(p.ins, func() (err error) {
				p.regs[r1] *= p.regs[r2]
				return
			})
		} else {
			if l2 == 0 {
				p.ins = append(p.ins, func() (err error) {
					p.regs[r1] = 0
					return
				})
			} else {
				p.ins = append(p.ins, func() (err error) {
					p.regs[r1] *= l2
					return
				})
			}
		}
	case "div":
		if ok2 {
			p.ins = append(p.ins, func() (err error) {
				if p.regs[r2] == 0 {
					return errors.New("Divide by zero")
				}
				p.regs[r1] /= p.regs[r2]
				return
			})
		} else {
			if l2 == 0 {
				panic("Divide by zero")
			}
			if l2 == 1 {
				// div by 1 is a noop
				p.ins = append(p.ins, func() (err error) { return })
			} else {
				p.ins = append(p.ins, func() (err error) {
					p.regs[r1] /= l2
					return
				})
			}
		}
	case "mod":
		if ok2 {
			p.ins = append(p.ins, func() (err error) {
				if p.regs[r1] < 0 {
					return errors.New("Cannod mod 0")
				}
				if p.regs[r2] <= 0 {
					return errors.New("Cannot mod by 0")
				}
				p.regs[r1] %= p.regs[r2]
				return
			})
		} else {
			if l2 <= 0 {
				panic("Mod by zero")
			}
			p.ins = append(p.ins, func() (err error) {
				if p.regs[r1] < 0 {
					return errors.New("Cannod mod 0")
				}
				p.regs[r1] %= l2
				return
			})
		}
	case "eql":
		if ok2 {
			p.ins = append(p.ins, func() (err error) {
				if p.regs[r1] == p.regs[r2] {
					p.regs[r1] = 1
				} else {
					p.regs[r1] = 0
				}
				return
			})
		} else {
			p.ins = append(p.ins, func() (err error) {
				if p.regs[r1] == l2 {
					p.regs[r1] = 1
				} else {
					p.regs[r1] = 0
				}
				return
			})
		}
	}
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

// returns true if no exception was encountered
func (p *program) run(inNum string) bool {
	p.pc = 0
	p.regs = registers{0, 0, 0, 0}
	p.input = make([]int, 0, len(inNum))

	for _, c := range []byte(inNum) {
		if c == '0' {
			return false
		}
		p.input = append(p.input, int(c-'0'))
	}

	for _, i := range p.ins {
		if err := i(); err != nil {
			fmt.Printf("%v at %s", err, p)
			return false
		}
		//fmt.Println(p)
		p.pc++

	}
	return true
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
