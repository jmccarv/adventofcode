package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type beaconPair struct {
	p1       point
	p2       point
	distance float64
}

type scanner struct {
	origin  point
	beacons []point
}

func main() {
	var scanners []scanner
	var s scanner

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if strings.Index(input.Text(), "scanner") > -1 {
			if len(s.beacons) > 0 {
				scanners = append(scanners, s)
				s = scanner{}
			}
			continue
		}

		b := strings.Split(input.Text(), ",")
		if len(b) == 3 {
			s.beacons = append(s.beacons, point{atoi(b[0]), atoi(b[1]), atoi(b[2])})
		}
	}
	if len(s.beacons) > 0 {
		scanners = append(scanners, s)
	}

	/*
		fmt.Println(scanners)
		for i := 1; i < len(scanners[0].beacons)-1; i++ {
			b := scanners[0].beacons
			fmt.Println(b[i], b[i+1], distance(b[i], b[i+1]))
		}
	*/
	part1(scanners)
}

func test() {
	p := point{5, 6, -4}
	for x := 0; x < 4; x++ {
		px := p.rotate('x', x*90)
		fmt.Println(px, px.rotate('z', 90), px.rotate('z', 180), px.rotate('z', 270), px.rotate('y', 90), px.rotate('y', 270))
	}
}

func part1(scanners []scanner) {
	for _, s1 := range scanners {
		for _, s2 := range scanners[1:] {
			fmt.Println("here")
			s1.findOverlap(s2)
			return
		}
	}
	//for
}

func (s scanner) findOverlap(s2 scanner) []point {
	for i1, b := range s.beacons {
		// Translate this to the origin 0,0,0
		xlate1 := point{-b.x, -b.y, -b.z}
		s.translate(xlate1)
		b = point{}

		for i2, b2 := range s2.beacons {
			// Let's pretend this beacon (b2) is the same as b and see what we find
			xlate2 := point{-b2.x, -b2.y, -b2.z}
			s2.translate(xlate2)
			b2 = point{}

			// Now go through all possible rotations and count how many beacons are in common between s and s2
			nr := 0
			for x := 0; x < 4; x++ {
				s2.rotate('x', 90)
				nr = s.countOverlap(s2)
				if nr > 11 {
					break
				}
				for z := 1; z < 4; z++ {
					s2.rotate('z', 90)
					nr = s.countOverlap(s2)
					if nr > 11 {
						break
					}
				}
				s2.rotate('y', 90)
				nr = s.countOverlap(s2)
				if nr > 11 {
					break
				}
				s2.rotate('y', 180)
				nr = s.countOverlap(s2)
				if nr > 11 {
					break
				}
			}
			if nr > 11 {
				fmt.Println(i1, i2, nr, s.origin, s2.origin)
				break
			}
		}
	}
	return []point{}
}

func (s scanner) countOverlap(s2 scanner) int {
	ret := 0
	m := make(map[point]struct{})
	for _, b := range s.beacons {
		m[b] = struct{}{}
	}

	for _, b := range s2.beacons {
		if _, ok := m[b]; ok {
			ret++
		}
	}
	return ret
}

func (s *scanner) translate(by point) {
	s.origin = s.origin.translate(by)
	for i, _ := range s.beacons {
		s.beacons[i] = s.beacons[i].translate(by)
	}
}

func (s *scanner) rotate(around byte, deg int) {
	for i, _ := range s.beacons {
		s.beacons[i] = s.beacons[i].rotate(around, deg)
	}
	s.origin = s.origin.rotate(around, deg)
}

func (p point) translate(by point) point {
	p.x += by.x
	p.y += by.y
	p.z += by.z
	return p
}

func (p point) String() string {
	return fmt.Sprintf("{%4d %4d %4d}", p.x, p.y, p.z)
}

func (p point) rotate(around byte, deg int) point {
	if deg == 0 {
		return p
	}

	sin, cos := math.Sincos(float64(deg) * math.Pi / 180)
	x := float64(p.x)
	y := float64(p.y)
	z := float64(p.z)

	switch around {
	case 'x':
		p.y = int(math.Round(y*cos - z*sin))
		p.z = int(math.Round(y*sin + z*cos))
	case 'y':
		p.x = int(math.Round(x*cos + z*sin))
		p.z = int(math.Round(z*cos - x*sin))
	case 'z':
		p.x = int(math.Round(x*cos - y*sin))
		p.y = int(math.Round(x*sin + y*cos))
	}
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func atoi(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func atof(s string) float64 {
	x, _ := strconv.ParseFloat(s, 64)
	return x
}

// distance between two points, A & B:
// AB = √( (x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2 )
/*
func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2) + math.Pow(p2.z-p1.z, 2))
}
*/

/*
func part1()
	candidates := scanners[0].beacons

	fmt.Println("candidate len", len(candidates))
	for _, s := range scanners[1:] {
		candidates = s.find(candidates)
		fmt.Println("candidate len", len(candidates))
		//break
	}
}
*/

/*
func (s scanner) find(beacons []point) (found []point) {
	bp := make(map[float64]beaconPair)
	for i := 0; i < len(beacons)-1; i++ {
		b1 := beacons[i]
		for j := i + 1; j < len(beacons); j++ {
			b2 := beacons[j]
			p := beaconPair{b1, b2, distance(b1, b2)}
			if x, ok := bp[p.distance]; ok {
				fmt.Println("Whoops, duplicate lenghts")
				fmt.Println(x)
				fmt.Println(p)
			}
			bp[p.distance] = p
		}
	}
	fmt.Println()
	fmt.Println(bp)
	fmt.Println()

	for i := 0; i < len(s.beacons)-1; i++ {
		b1 := s.beacons[i]
		for j := i + 1; j < len(s.beacons); j++ {
			b2 := s.beacons[j]
			p := beaconPair{b1, b2, distance(b1, b2)}
			if f, ok := bp[p.distance]; ok {
				fmt.Println("Found one", f)
				found = append(found, f.p1, f.p2)
			}
		}
	}

	return
}
*/
