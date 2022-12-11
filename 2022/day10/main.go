package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type CRT [6][]byte

type instruction struct {
	op  string
	arg int
}

type machine struct {
	clock int
	regX  int
}

var display CRT
var p1 int

func main() {
	t0 := time.Now()

	m := machine{regX: 1}
	display.init()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var ins instruction
		nr, _ := fmt.Sscanf(s.Text(), "%s %d", &ins.op, &ins.arg)
		if nr < 1 {
			panic("Invalid input!")
		}

		m.clock++
		m.crt()
		m.cpu(ins)
	}

	fmt.Println("Part 1", p1)
	display.show()
	fmt.Println("Total Time", time.Now().Sub(t0))
}

func (d *CRT) init() {
	for r := 0; r < 6; r++ {
		d[r] = make([]byte, 40)
		for i := 0; i < 40; i++ {
			d[r][i] = '.'
		}
	}
}

func (d CRT) show() {
	for r := 0; r < 6; r++ {
		fmt.Println(string(d[r]))
	}
}

func (m *machine) crt() {
	r := (m.clock - 1) / 40
	c := (m.clock - 1) % 40

	// X register is the midpoint of our 3px wide cursor
	if c >= m.regX-1 && c <= m.regX+1 {
		// The current pixel is lit
		display[r][c] = '#'
	}
}

func (m *machine) cpu(ins instruction) {
	m.part1()
	if ins.op == "addx" {
		m.clock++
		m.crt()
		m.part1()
		m.regX += ins.arg
	}
}

func (m machine) part1() {
	if m.clock == 20 || (m.clock-20)%40 == 0 {
		p1 += m.regX * m.clock
	}
}
