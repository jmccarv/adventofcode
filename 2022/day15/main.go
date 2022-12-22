package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"

	sm "github.com/jmccarv/adventofcode/util/math"
	s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

type sensor struct {
	s2d.Point
	beacon s2d.Point
}

var (
	p1y, p2max int
)

func usage() {
	fmt.Println("usage: main example|input")
	os.Exit(1)
}

func main() {
	t0 := time.Now()
	if len(os.Args) != 2 {
		usage()
	}
	switch os.Args[1] {
	case "example":
		p1y = 10
		p2max = 20
	case "input":
		p1y = 2000000
		p2max = 4000000
	default:
		usage()
	}

	input := bufio.NewScanner(os.Stdin)

	var sensors []sensor

	for input.Scan() {
		var s sensor
		nr, err := fmt.Sscanf(input.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &s.beacon.X, &s.beacon.Y)
		if nr != 4 {
			panic("Bad input or something " + err.Error())
		}

		sensors = append(sensors, s)
	}

	part1(sensors)
	part2(sensors)
	fmt.Println("Total Time", time.Now().Sub(t0))
}

func findCoveredInRow(y int, sensors []sensor) []s2d.Box {
	var lines []s2d.Box
	for _, s := range sensors {
		sbDist := s.MhDistance(s.beacon)
		if sm.Abs(s.Y-y) <= sbDist {
			//Sensor overlaps

			// lenx = length to the left and right centered at s.X, that the sensor can see on row y
			lenx := sbDist - sm.Abs(s.Y-y)
			l := s2d.Box{s2d.Point{s.X - lenx, y}, s2d.Point{s.X + lenx, y}}
			//fmt.Println(s, s.beacon, sbDist, lenx, l)

			lines = append(lines, l)
		}
	}

	// Sort lines from left (-x) to right (+x)
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].TL.X < lines[j].TL.X
	})

	var ret []s2d.Box
	cur := lines[0]
	for _, l := range lines[1:] {
		if cur.Overlaps(l) {
			//fmt.Println("Overlap:", cur, l)
			cur.BR = cur.BR.Max(l.BR)
		} else {
			//fmt.Println("CUR:", cur, cur.NrPoints())
			ret = append(ret, cur)
			cur = l
		}
	}
	//fmt.Println("CUR:", cur, cur.NrPoints())
	ret = append(ret, cur)

	return ret
}

func part1(sensors []sensor) {
	t0 := time.Now()
	beaconsAtY := make(map[s2d.Point]struct{})
	for _, s := range sensors {
		if s.beacon.Y == p1y {
			beaconsAtY[s.beacon] = struct{}{}
		}
	}
	sum := 0
	for _, l := range findCoveredInRow(p1y, sensors) {
		sum += l.NrPoints()
	}

	fmt.Println("Part1", sum-len(beaconsAtY), time.Now().Sub(t0))
}

func part2(sensors []sensor) {
	t0 := time.Now()
	find := func(lines []s2d.Box, y int) (s2d.Point, bool) {
		prev := s2d.Box{s2d.Point{-1, y}, s2d.Point{-1, y}}
		for _, l := range lines {
			if prev.BR.X+1 < l.TL.X {
				return prev.BR.Add(s2d.Point{1, 0}), true
			}
			prev = l
		}

		if prev.BR.X < p2max {
			return prev.BR.Add(s2d.Point{1, 0}), true
		}

		return s2d.Point{}, false
	}

	for y := 0; y <= p2max; y++ {
		lines := findCoveredInRow(y, sensors)
		if p, ok := find(lines, y); ok {
			fmt.Println("Part 2", p.X*4000000+p.Y, p, time.Now().Sub(t0))
			break
		}
	}
}
