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

		m.clock++
		m.crt()
		m.cpu(ins)
	}

	tm.Println()
	tm.Println("Part 1:", p1)
	tm.Flush()
}

func (d *CRT) init() {
	tm.Clear()
}

func (m *machine) crt() {
	r := (m.clock - 1) / 40
	c := (m.clock - 1) % 40

	// X register is the midpoint of our 3px wide cursor
	if c >= m.regX-1 && c <= m.regX+1 {
		// The current pixel is lit
		tm.MoveCursor(c+1, r+1)
		tm.Print("█")
		tm.Flush()
		time.Sleep(time.Second / 50)
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
