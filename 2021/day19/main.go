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

type rotate struct {
	xdeg, zdeg int
}

type beaconPair struct {
	p1       point
	p2       point
	distance float64
}

type cloud struct {
	id      int
	origin  point
	beacons []point
	locked  byte
}

func main() {
	var clouds []*cloud
	s := &cloud{}
	id := 0

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if strings.Index(input.Text(), "scanner") > -1 {
			if len(s.beacons) > 0 {
				clouds = append(clouds, s)
				id++
				s = &cloud{id: id}
			}
			continue
		}

		b := strings.Split(input.Text(), ",")
		if len(b) == 3 {
			s.beacons = append(s.beacons, point{atoi(b[0]), atoi(b[1]), atoi(b[2])})
		}
	}
	if len(s.beacons) > 0 {
		clouds = append(clouds, s)
	}

	part1(clouds)
}

func test() {
	p := point{5, 6, -4}
	for x := 0; x < 4; x++ {
		px := p.rotate('x', x*90)
		fmt.Println(px, px.rotate('z', 90), px.rotate('z', 180), px.rotate('z', 270), px.rotate('y', 90), px.rotate('y', 270))
	}
}

func part1(clouds []*cloud) {
	var found, remain []*cloud
	clouds[0].locked = 1

	/*
		fmt.Println(clouds)
		if clouds[0].detect(clouds[1]) {
			fmt.Println("found", clouds[0], clouds[1])
			m := make(map[point]int)
			for _, p := range clouds[0].beacons {
				m[p]++
			}
			for _, p := range clouds[1].beacons {
				m[p]++
			}
			fmt.Println(len(m))
			for k, v := range m {
				if v > 1 {
					x := point{-clouds[1].origin.x, -clouds[1].origin.y, -clouds[1].origin.z}
					fmt.Println(k, k.translate(x))
				}
			}
		}
		return
	*/

	found = append(found, clouds[0])
	remain = append(remain, clouds[1:]...)
	done := false

	for !done {
		done = true
		for _, s1 := range found {
			var r []*cloud
			for _, s2 := range remain {
				if s1.detect(s2) {
					done = false
					found = append(found, s2)
					fmt.Println("Found:", s2)
				} else {
					r = append(r, s2)
				}
			}
			remain = r
		}
	}
	if len(remain) > 0 {
		panic("Failed to lock all scanners :(")
	}

	m := make(map[point]int)
	for _, s := range found {
		for _, p := range s.beacons {
			m[p]++
		}
	}
	fmt.Println(len(m))

	//fmt.Println(nr)
}

func (s *cloud) detect(s2 *cloud) (found bool) {
	overlap := 0
	for _, b := range s.beacons {
		// Translate this to the origin 0,0,0
		xlate_origin := point{-b.x, -b.y, -b.z}
		xlate_back := b

		s.translate(xlate_origin)
		b = point{}

	out:
		for _, b2 := range s2.beacons {
			// Let's pretend this beacon (b2) is the same as b and see what we find
			xlate2 := point{-b2.x, -b2.y, -b2.z}
			s2.translate(xlate2)
			b2 = point{}

			// Now go through all possible rotations and count how many beacons are in common between s and s2
			// I think I'm doing too much work here but my head hurts and this is working
			for x := 0; x < 4; x++ {
				s2.rotate('x', 90)
				if overlap = s.overlapping(s2); overlap > 11 {
					break out
				}
				for z := 0; z < 4; z++ {
					s2.rotate('z', 90)
					if overlap = s.overlapping(s2); overlap > 11 {
						break out
					}

					for y := 0; y < 4; y++ {
						s2.rotate('y', 90)
						if overlap = s.overlapping(s2); overlap > 11 {
							break out
						}
					}
				}
				/*
					s2.rotate('y', 90)
					if overlap = s.overlapping(s2); overlap > 11 {
						break out
					}
					s2.rotate('y', 180)
					if overlap = s.overlapping(s2); overlap > 11 {
						break out
					}
				*/
			}
		}
		s.translate(xlate_back)

		if overlap > 11 {
			s2.translate(xlate_back)
			s2.locked = 1
			found = true
			break
		}
	}
	return
}

func (s *cloud) overlapping(s2 *cloud) (nr int) {
	m := make(map[point]struct{})
	for _, b := range s.beacons {
		m[b] = struct{}{}
	}
	for _, b := range s2.beacons {
		if _, ok := m[b]; ok {
			nr++
		}
	}
	return
}

func (s *cloud) translate(by point) {
	s.origin = s.origin.translate(by)
	for i, _ := range s.beacons {
		s.beacons[i] = s.beacons[i].translate(by)
	}
}

func (s *cloud) rotate(around byte, deg int) {
	for i, _ := range s.beacons {
		s.beacons[i] = s.beacons[i].rotate(around, deg)
	}
	s.origin = s.origin.rotate(around, deg)
}

func (c *cloud) String() string {
	return fmt.Sprintf("{id: %d  origin: %v  nrpt: %d  lck: %d}", c.id, c.origin, len(c.beacons), c.locked)
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
	candidates := clouds[0].beacons

	fmt.Println("candidate len", len(candidates))
	for _, s := range clouds[1:] {
		candidates = s.find(candidates)
		fmt.Println("candidate len", len(candidates))
		//break
	}
}
*/

/*
func (s cloud) find(beacons []point) (found []point) {
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
