package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Side note, this method is much better/faster than what I used in part 1.
// The Part 1 code is terrible, but I just wanted to get to Part 2 quickly...

// A single point represents a 'cube' in our problem definition, so
// the word 'cube' can also be read as 'point'
type point [3]int // [x,y,z]

// A cuboid defines a cubish shaped region of points, inclusive.
// Where all points in this cuboid are considered to be turned 'on'
type cuboid struct {
	min point
	max point
}

// This is our game grid -- a list of all 'on' cuboid regions
type cuboids []cuboid

func main() {
	var grid cuboids
	input := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`[-0-9]+`)
	for input.Scan() {
		s := strings.SplitN(input.Text(), " ", 2)
		m := re.FindAllString(s[1], 6)
		if len(m) == 6 {
			w := [6]int{atoi(m[0]), atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4]), atoi(m[5])}
			r := cuboid{
				min: point{min(w[0], w[1]), min(w[2], w[3]), min(w[4], w[5])},
				max: point{max(w[0], w[1]), max(w[2], w[3]), max(w[4], w[5])},
			}

			switch s[0] {
			case "on":
				grid.add(r)
			case "off":
				grid.sub(r)
			}
		}
	}
	fmt.Println("Cuboids:", len(grid))
	fmt.Println("Part2:", grid.count())
}

// Return how many total points(cubes) are contained in all cuboids
// This is the number of points/cubes that are 'on'
func (cubes cuboids) count() (ret int) {
	for _, c := range cubes {
		ret += (1 + c.max[0] - c.min[0]) * (1 + c.max[1] - c.min[1]) * (1 + c.max[2] - c.min[2])
	}
	return
}

// Turn 'on' all points in the given cuboid by
// adding it to our list of all 'on' cubes
func (cubes *cuboids) add(r cuboid) {
	*cubes = append(*cubes, cubes.diff(r)...)
}

// given a cuboid, 'r', return a list of cuboids conained in r
// that are not already contained in the set 'cubes'
func (cubes *cuboids) diff(r cuboid) cuboids {
	newCuboids := cuboids{r}

	for _, c := range *cubes {
		var nc cuboids
		for _, x := range newCuboids {
			nc = append(nc, x.sub(c)...)
		}
		newCuboids = nc
	}
	return newCuboids
}

// Turn 'off' all cubes in the given cuboid by
// removing them from our list. This means having
// to intersect r with all full list to find out what cubes
// need to be removed.
func (cubes *cuboids) sub(r cuboid) {
	var newCuboids cuboids
	for _, c := range *cubes {
		newCuboids = append(newCuboids, c.sub(r)...)
	}
	*cubes = newCuboids
	return
}

// subtract x from r and return zero or more cuboids as a result
// This is easier to picture in 2D but basically for each face
// of 'x' we are treating that face as a plane to cut 'r'
// if it intersects with 'x'. We add the part of 'x' that was
// outside the cut plane to our list to return and continue to
// the next face.
func (r cuboid) sub(x cuboid) cuboids {

	// Quick check to see if they overlap, if they don't
	// there's no need to perform any processing, we just
	// return the original cuboid
	if !r.overlapsWith(x) {
		return []cuboid{r}
	}

	ret := make([]cuboid, 0, 6)

	for i := 0; i < 3; i++ {
		if x.min[i] >= r.min[i] {
			nr := r
			nr.max[i] = x.min[i] - 1
			if nr.max[i] >= nr.min[i] {
				ret = append(ret, nr)
			}
			r.min[i] = x.min[i]
		}
		if x.max[i] <= r.max[i] {
			nr := r
			nr.min[i] = x.max[i] + 1
			if nr.max[i] >= nr.min[i] {
				ret = append(ret, nr)
			}
			r.max[i] = x.max[i]
		}
	}
	return ret
}

func (r cuboid) overlapsWith(x cuboid) bool {
	for i := 0; i < 3; i++ {
		if r.max[i] < x.min[i] || r.min[i] > x.max[i] {
			return false
		}
	}
	return true
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
