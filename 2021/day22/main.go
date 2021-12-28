package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cuboid struct {
	min point
	max point
}

type cuboids []cuboid

type point [3]int // [x,y,z]
/*
type point struct {
	x, y, z int
}
*/

func main() {
	var grid cuboids
	input := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`[-0-9]+`)
	//in:
	for input.Scan() {
		s := strings.SplitN(input.Text(), " ", 2)
		m := re.FindAllString(s[1], 6)
		if len(m) == 6 {
			w := [6]int{atoi(m[0]), atoi(m[1]), atoi(m[2]), atoi(m[3]), atoi(m[4]), atoi(m[5])}
			r := cuboid{
				min: point{min(w[0], w[1]), min(w[2], w[3]), min(w[4], w[5])},
				max: point{max(w[0], w[1]), max(w[2], w[3]), max(w[4], w[5])},
			}

			// part1
			/*
				for i := 0; i < 3; i++ {
					if r.min[i] < -50 || r.max[i] > 50 {
						fmt.Println(i, r.min[i], r.max[i], input.Text())
						continue in
					}
				}
			*/

			switch s[0] {
			case "on":
				grid.add(r)
			case "off":
				grid.sub(r)
			}
		}
	}
	//fmt.Println(grid)
	fmt.Println(grid.count())
}

func (cubes cuboids) count() (ret int) {
	for _, c := range cubes {
		ret += (1 + c.max[0] - c.min[0]) * (1 + c.max[1] - c.min[1]) * (1 + c.max[2] - c.min[2])
	}
	return
}

func (cubes *cuboids) add(r cuboid) {
	*cubes = append(*cubes, cubes.diff(r)...)
	//cubes.verify()
}

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

func (cubes *cuboids) sub(r cuboid) {
	var newCuboids cuboids
	for _, c := range *cubes {
		newCuboids = append(newCuboids, c.sub(r)...)
	}
	*cubes = newCuboids
	cubes.verify()
	return
}

func (cubes cuboids) verify() {
	if len(cubes) < 2 {
		return
	}
	for i, c1 := range cubes[1:] {
		for _, c2 := range cubes[0:i] {
			if c1.overlapsWith(c2) {
				fmt.Println("*** Overlap ***", c1, c2)
			}
		}
	}
}

// subtract x from r and return zero or more cuboids as a result
func (r cuboid) sub(x cuboid) cuboids {

	// Quick check to see if they overlap, if they don't
	// there's no need to perform any processing, we just
	// return the original cuboid
	//fmt.Println("sub: ", x, "from", r)
	if !r.overlapsWith(x) {
		return []cuboid{r}
	}

	ret := make([]cuboid, 0, 6)

	for i := 0; i < 3; i++ {
		if x.min[i] >= r.min[i] {
			nr := r
			nr.max[i] = x.min[i] - 1
			ret = append(ret, nr)
			r.min[i] = x.min[i]
		}
		if x.max[i] <= r.max[i] {
			nr := r
			nr.min[i] = x.max[i] + 1
			ret = append(ret, nr)
			r.max[i] = x.max[i]
		}
		//fmt.Println("i", i, "r", r, "ret", ret)
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
