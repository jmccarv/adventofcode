package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cost struct {
	nrCrabs int
	fuel    int
}

type bucket struct {
	pos     int
	nrCrabs int
	left    cost
	right   cost
	total   cost
}

type buckets []*bucket

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}

	bmap := make(map[int]*bucket)
	for _, n := range strings.Split(s.Text(), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic("Invalid input")
		}
		if b, ok := bmap[i]; ok {
			b.addCrab()
		} else {
			bmap[i] = &bucket{pos: i, nrCrabs: 1}
		}
	}

	var blist buckets
	for _, b := range bmap {
		blist = append(blist, b)
	}
	sort.Slice(blist, func(i, j int) bool {
		return blist[i].pos < blist[j].pos
	})

	part1(blist)
}

func part1(blist buckets) {
	var minCost int
	var winner *bucket

	prev := blist[0]
	for i := 1; i < len(blist); i++ {
		nr := prev.nrCrabs + prev.left.nrCrabs
		blist[i].left = cost{
			nrCrabs: nr,
			fuel:    prev.left.fuel + nr*(blist[i].pos-prev.pos),
		}
		prev = blist[i]
	}

	prev = blist[len(blist)-1]
	for i := len(blist) - 2; i >= 0; i-- {
		nr := prev.nrCrabs + prev.right.nrCrabs
		blist[i].right = cost{
			nrCrabs: nr,
			fuel:    prev.right.fuel + nr*(prev.pos-blist[i].pos),
		}
		blist[i].total = blist[i].left.Add(blist[i].right)
		if minCost == 0 || blist[i].total.fuel < minCost {
			minCost = blist[i].total.fuel
			winner = blist[i]
		}
		prev = blist[i]
	}

	fmt.Println(blist)
	fmt.Println(winner)
	fmt.Println("required fuel: ", winner.total.fuel)
}

func (c cost) Add(d cost) cost {
	return cost{
		nrCrabs: c.nrCrabs + d.nrCrabs,
		fuel:    c.fuel + d.fuel,
	}
}

func (b *bucket) addCrab() {
	b.nrCrabs++
}

func (b *bucket) String() string {
	return fmt.Sprintf("%+v", *b)
}

func (l buckets) String() string {
	s := ""
	for _, b := range l {
		s += fmt.Sprintf("%+v\n", b)
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
