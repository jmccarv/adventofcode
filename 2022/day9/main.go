package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type point struct {
	x, y int
}

type knot struct {
	point         // where this knot currently sits
	ofs     point // offset from the 'head' knot
	visited map[point]struct{}
}

var directions = map[int]point{
	'U': point{0, -1},
	'D': point{0, 1},
	'L': point{-1, 0},
	'R': point{1, 0},
}

func main() {
	t0 := time.Now()

	tail := knot{visited: map[point]struct{}{point{0, 0}: struct{}{}}}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var udrl int
		var amt int
		nr, err := fmt.Sscanf(s.Text(), "%c %d", &udrl, &amt)
		if err != nil || nr != 2 {
			panic("Invalid input!")
		}
		tail.move(directions[udrl], amt)
	}

	fmt.Println(len(tail.visited))
	fmt.Println("Total time ", time.Now().Sub(t0))
}

// move head and update tail to follow head
// dir is the direction head moves and amt is number of spaces it moves in that direction
func (k *knot) move(dir point, amt int) {
	// Move the 'head' knot by adjusting our offset
	vec := point{x: dir.x * amt, y: dir.y * amt}
	k.ofs.add(vec)

	// Now the tail moves so it is always neighboring head
	// For the first step, tail may move in both axes
	// After the first move, it may only move in the axis that head moved in
	moveDir := point{x: sign(k.ofs.x), y: sign(k.ofs.y)}
	for !k.isNeighbor() {
		k.add(moveDir)
		k.ofs.sub(moveDir)
		k.visited[k.point] = struct{}{}
		moveDir = point{x: moveDir.x * abs(dir.x), y: moveDir.y * abs(dir.y)}
	}
}

func (p *point) add(q point) {
	p.x += q.x
	p.y += q.y
}

func (p *point) sub(q point) {
	p.x -= q.x
	p.y -= q.y
}

func (k knot) isNeighbor() bool {
	return abs(k.ofs.x) <= 1 && abs(k.ofs.y) <= 1
}

func (k knot) String() string {
	return fmt.Sprintf("%v %v", k.point, k.ofs)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -1
	}
	return 1
}
