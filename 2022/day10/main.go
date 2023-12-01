package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"
)

type CRT struct{}

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
	m := machine{regX: 1}
	display.init()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var ins instruction
		nr, _ := fmt.Sscanf(s.Text(), "%s %d", &ins.op, &ins.arg)
		if nr < 1 {
			panic("Invalid input!")
		}

		m.tick()
		m.crt()
		m.cpu(ins)
	}

	tm.MoveCursor(1, 12)
	tm.Println("Part 1:", p1)
	tm.Flush()
}

func (d *CRT) init() {
	tm.Clear()

	tm.MoveCursor(1, 1)

	b := tm.NewBox(20, 4, 0)
	b.Border = "─ │ ┌ ┐ └ ┘"

	tm.Print(tm.Color(tm.MoveTo(b.String(), 10, 1), tm.RED))
	tm.Print(tm.Color(tm.MoveTo("CLK      X", 12, 2), tm.RED))
	tm.Print(tm.Color(tm.MoveTo("S      T", 14, 3), tm.RED))
}

func (m *machine) tick() {
	m.clock++
	tm.Print(tm.Color(tm.MoveTo(fmt.Sprintf("%03d", m.clock), 16, 2), tm.BLUE))
	tm.Print(tm.Color(tm.MoveTo(fmt.Sprintf("%03d", m.regX), 23, 2), tm.BLUE))
	tm.Flush()
}

func (m *machine) crt() {
	r := (m.clock - 1) / 40
	c := (m.clock - 1) % 40

	// X register is the midpoint of our 3px wide cursor
	if c >= m.regX-1 && c <= m.regX+1 {
		// The current pixel is lit
		tm.MoveCursor(c+1, r+5)
		tm.Print(tm.Color("█", tm.GREEN))
		tm.Flush()
		time.Sleep(time.Second / 25)
	}
}

func (m *machine) cpu(ins instruction) {
	m.part1()
	if ins.op == "addx" {
		m.tick()
		m.crt()
		m.part1()
		m.regX += ins.arg
	}
}

func (m machine) part1() {
	if (m.clock-20)%40 == 0 {
		s := m.regX * m.clock
		p1 += s

		tm.Print(tm.Color(tm.MoveTo(fmt.Sprintf("%4d", s), 16, 3), tm.BLUE))
		tm.Print(tm.Color(tm.MoveTo(fmt.Sprintf("%-5d", p1), 23, 3), tm.BLUE))
		tm.Flush()
	}
}
