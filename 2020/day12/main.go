package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type command struct {
	cmd byte
	arg int
}

type position struct {
	x int
	y int
}

var compass_dir = map[byte]int{
	'N': 0,
	'E': 90,
	'S': 180,
	'W': 270,
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var commands []command

	for s.Scan() {
		cmd, arg := s.Text()[0], s.Text()[1:]
		n, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal("Invalid input: ", s.Text())
		}
		commands = append(commands, command{cmd, n})
	}

	fmt.Println("part1:", part1(commands))
	fmt.Println("part1:", part2(commands))
}

func part1(commands []command) int {
	var pos position
	facing := compass_dir['E']

	fmt.Println(pos, facing)
	for _, c := range commands {
		switch c.cmd {
		case 'F':
			pos = pos.move(facing, c.arg)
		case 'N', 'E', 'S', 'W':
			pos = pos.move(compass_dir[c.cmd], c.arg)
		case 'R':
			facing = (facing + c.arg) % 360
		case 'L':
			facing = (facing + (360 - c.arg)) % 360
		}
		fmt.Println(c, pos, facing)
	}
	return pos.distance()
}

// 21952 too low
func part2(commands []command) int {
	ship := position{}       // relative to origin (0,0)
	waypt := position{10, 1} // relative to ship's position

	for _, c := range commands {
		switch c.cmd {
		case 'F':
			ship = ship.vector(waypt, c.arg)
		case 'N', 'E', 'S', 'W':
			waypt = waypt.move(compass_dir[c.cmd], c.arg)
		case 'R':
			waypt = waypt.rotate(c.arg)
		case 'L':
			waypt = waypt.rotate(360 - c.arg)
		}
		fmt.Println(c, ship, waypt)
	}
	return ship.distance()
}

func (c command) String() string {
	return fmt.Sprintf("%c %3d", c.cmd, c.arg)
}

func (p position) move(deg int, amt int) position {
	p.x += amt * int(math.Sin(rad(float64(deg))))
	p.y += amt * int(math.Cos(rad(float64(deg))))
	return p
}

func (p position) vector(pt position, amt int) position {
	p.x += pt.x * amt
	p.y += pt.y * amt
	return p
}

func (p position) rotate(deg int) position {
	var ret position

	s := int(math.Sin(rad(float64(deg))))
	c := int(math.Cos(rad(float64(deg))))

	ret.x = p.x*c + p.y*s
	ret.y = p.y*c - p.x*s
	return ret
}

func (p position) distance() int {
	return abs(p.x) + abs(p.y)
}

func (p position) String() string {
	return fmt.Sprintf("[%5d %5d]", p.x, p.y)
}

func rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
