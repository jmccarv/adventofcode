package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bucket struct {
	nrCrabs        int
	distanceWalked int
	fuelUsed       int
}

type location struct {
	pos         int
	nativeCrabs bucket   // the crabs that started life here
	left        []bucket // crabs that migrated from the left
	right       []bucket // crabs that migrated from the right
}

type locations []location

func main() {
	var first, last int
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}

	lmap := make(map[int]*location)
	for _, n := range strings.Split(s.Text(), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic("Invalid input")
		}
		if _, ok := lmap[i]; !ok {
			lmap[i] = &location{pos: i, left: make([]bucket, 0), right: make([]bucket, 0)}
		}
		lmap[i].nativeCrabs.nrCrabs++
		first = min(first, i)
		last = max(last, i)
	}

	part2(lmap, first, last)
}

func (from location) migrateTo(to location) []bucket {
	ret := make([]bucket, 0)

	lr := from.left
	distance := to.pos - from.pos
	if distance < 0 {
		distance = -distance
		lr = from.right
	}

	ret = append(ret, from.nativeCrabs)
	ret[0].fuelUsed = ret[0].nrCrabs * calcFuelCost(ret[0].distanceWalked+distance)
	ret[0].distanceWalked += distance

	for i, b := range lr {
		ret = append(ret, b)
		ret[i+1].fuelUsed = ret[i+1].nrCrabs * calcFuelCost(ret[i+1].distanceWalked+distance)
		ret[i+1].distanceWalked += distance
	}

	return ret
}

func calcFuelCost(n int) int {
	return (n * (n + 1)) / 2
}

func part2(lmap map[int]*location, first, last int) {
	nrLocs := last - first + 1
	llist := locations(make([]location, nrLocs))

	if loc, ok := lmap[0]; ok {
		llist[0] = *loc
	}
	prev := llist[0]
	for i := 1; i < nrLocs; i++ {
		loc := location{pos: i}
		if _, ok := lmap[i]; ok {
			loc = *lmap[i]
		}
		loc.left = prev.migrateTo(loc)
		llist[i] = loc
		prev = loc
	}

	prev = llist[len(llist)-1]
	for i := len(llist) - 2; i >= 0; i-- {
		llist[i].right = prev.migrateTo(llist[i])
		prev = llist[i]
	}

	var winner location
	var minFuel int
	for _, l := range llist {
		f := l.cost()
		if minFuel == 0 || f < minFuel {
			minFuel = f
			winner = l
		}
	}

	//fmt.Println(llist)
	fmt.Println(winner.pos, minFuel)
}

func (l location) cost() int {
	ret := 0
	for _, b := range l.left {
		ret += b.fuelUsed
	}
	for _, b := range l.right {
		ret += b.fuelUsed
	}
	return ret
}

func (b *bucket) String() string {
	return fmt.Sprintf("%+v", *b)
}

func (locs locations) String() string {
	s := ""
	for _, l := range locs {
		s += fmt.Sprintf("pos:%d native:%+v left:%d right:%d cost:%d\n", l.pos, l.nativeCrabs, len(l.left), len(l.right), l.cost())
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
