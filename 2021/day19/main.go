package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type cloud struct {
	id       int
	origin   point
	beacons  []point
	searched bool
}

type rotation struct {
	around byte
	deg    int
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

	solve(clouds)
}

func solve(clouds []*cloud) {
	var found, remain []*cloud

	found = append(found, clouds[0])
	remain = append(remain, clouds[1:]...)
	done := false

	for !done {
		done = true
		for _, s1 := range found {
			if s1.searched {
				continue
			}
			s1.searched = true

			var r []*cloud
			for _, s2 := range remain {
				if s1.detect(s2) {
					done = false
					found = append(found, s2)
					fmt.Println("Found:", s2, " From", s1)
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
	fmt.Println("Part1", len(m))

	dist := 0
	for i := 0; i < len(clouds)-1; i++ {
		for j := i + 1; j < len(clouds); j++ {
			dist = max(dist, distance(clouds[i].origin, clouds[j].origin))
		}
	}
	fmt.Println("Part2", dist)
}

func (s *cloud) detect(s2 *cloud) (found bool) {
	rotations := [4]rotation{
		rotation{'x', 90},
		rotation{'y', 90},
		rotation{'y', 90},
		rotation{'y', 90},
	}

	overlap := 0
	for _, b := range s.beacons {
		// Translate this to the origin 0,0,0
		xlate_origin := point{-b.x, -b.y, -b.z}
		xlate_back := b

		s.translate(xlate_origin)
		b = point{}

		// Now go through all possible rotations and count how many beacons are in common between s and s2
		// If we have at 12 or more we've matched s1 to s2 and have located s2's scanner's origin
	out:
		for p := 0; p < 2; p++ {
			for q := 0; q < 3; q++ {
				for _, r := range rotations {
					s2.rotate(r)
					for _, b2 := range s2.beacons {
						s2.translate(point{-b2.x, -b2.y, -b2.z})
						if overlap = s.overlapping(s2); overlap > 11 {
							break out
						}
					}
				}
			}
			s2.rotate(rotations[0])
			s2.rotate(rotations[1])
			s2.rotate(rotations[0])
		}

		// return s to it's original orientation
		s.translate(xlate_back)

		if overlap > 11 {
			// Orient s2 the same as s1
			s2.translate(xlate_back)

			// s2 is now locked in the same orientation / origin as s1
			found = true
			return
		}
	}
	return
}

func (s *cloud) overlapping(s2 *cloud) (nr int) {
	m := make(map[point]struct{}, len(s.beacons))
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

func (s *cloud) rotate(r rotation) {
	for i, _ := range s.beacons {
		s.beacons[i] = s.beacons[i].rotate(r)
	}
	s.origin = s.origin.rotate(r)
}

func (c *cloud) String() string {
	return fmt.Sprintf("{id: %2d origin: %v nrpt: %d}", c.id, c.origin, len(c.beacons))
}

func (p point) translate(by point) point {
	p.x += by.x
	p.y += by.y
	p.z += by.z
	return p
}

func (p point) String() string {
	return fmt.Sprintf("%5d,%5d,%5d", p.x, p.y, p.z)
}

// Minimal rotation function that should be faster than the full blown floating
// point version below
func (p point) rotate(r rotation) point {
	switch r.around {
	case 'x':
		return point{p.x, -p.z, p.y}
	case 'y':
		return point{p.z, p.y, -p.x}
	default:
		panic("Invalid axis of rotation")
	}
	return point{}
}

/*
func (p point) rotate(r rotation) point {
	if r.deg == 0 {
		return p
	}

	sin, cos := math.Sincos(float64(r.deg) * math.Pi / 180)
	x := float64(p.x)
	y := float64(p.y)
	z := float64(p.z)

	switch r.around {
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
*/

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

// distance between two points, A & B:
// AB = √( (x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2 )
/*
func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2) + math.Pow(p2.z-p1.z, 2))
}
*/

// Manhattan distance
func distance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
