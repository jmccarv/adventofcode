package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	sm "github.com/jmccarv/adventofcode/util/math"
	s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

// 4032909 too low :(
// 4287538 too low :(

type line struct {
	l, r s2d.Point
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	y := 2000000 // part 1
	//y := 10 // part 1 example
	beaconsAtY := make(map[s2d.Point]struct{})
	var lines []s2d.Box
	for s.Scan() {
		var sensor, beacon s2d.Point
		nr, err := fmt.Sscanf(s.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.X, &sensor.Y, &beacon.X, &beacon.Y)
		if nr != 4 {
			panic("Bad input or something " + err.Error())
		}

		// part 1
		if beacon.Y == y {
			beaconsAtY[beacon] = struct{}{}
		}
		sbDist := sensor.MhDistance(beacon)
		if sm.Abs(sensor.Y-y) <= sbDist {
			//Sensor overlaps

			// lenx = length to the left and right centered at sensor.X, that the sensor can see on row y
			lenx := sbDist - sm.Abs(sensor.Y-y)
			line := s2d.Box{s2d.Point{sensor.X - lenx, y}, s2d.Point{sensor.X + lenx, y}}
			fmt.Println(sensor, beacon, sbDist, lenx, line)

			lines = append(lines, line)
		}
	}
	// Sort lines from left (-x) to right (+x)
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].TL.X < lines[j].TL.X
	})

	cur := lines[0]
	sum := 0
	for _, l := range lines[1:] {
		if cur.Overlaps(l) {
			fmt.Println("Overlap:", cur, l)
			cur.BR = cur.BR.Max(l.BR)
		} else {
			fmt.Println("CUR:", cur, cur.NrPoints())
			sum += cur.NrPoints()
			cur = l
		}
	}
	fmt.Println("CUR:", cur, cur.NrPoints())
	sum += cur.NrPoints()

	fmt.Println(lines)
	fmt.Println("Part1", sum-len(beaconsAtY))
}
