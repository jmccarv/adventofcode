package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type die interface {
	roll() int
	nrRolled() int
}

type deterministicDie struct {
	nrRolls int
}

type player struct {
	nr    int
	pos   int
	score int
}

func main() {
	var players []player
	input := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`Player *(\d) starting position: *(\d+)`)
	for input.Scan() {
		m := re.FindStringSubmatch(input.Text())
		if len(m) != 3 {
			panic("Invalid input")
		}
		players = append(players, player{atoi(m[1]), atoi(m[2]) - 1, 0})
	}

	p := 1
	d := die(&deterministicDie{})
	fmt.Println(players)
	for players[0].score < 1000 && players[1].score < 1000 {
		p = (p + 1) % 2
		players[p].takeTurn(d)
		//fmt.Println(players)
	}
	fmt.Println(players)

	fmt.Println(min(players[0].score, players[1].score) * d.nrRolled())
}

func (p *player) position() int {
	return p.pos + 1
}

func (p *player) takeTurn(d die) {
	pd := 0
	for i := 0; i < 3; i++ {
		pd += d.roll()
	}
	p.pos = (p.pos + pd) % 10
	//fmt.Println("pd", pd, "pos", p.pos)
	p.score += p.pos + 1
}

func (d *deterministicDie) roll() (pips int) {
	pips = (d.nrRolls + 1) % 100
	d.nrRolls++
	return
}

func (d *deterministicDie) nrRolled() int {
	return d.nrRolls
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		panic("Invalid input")
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
