package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type region struct {
	min point
	max point
	on  bool
}

type cuboids []region

type point struct {
	x, y, z int
}

func main() {
	var grid cuboids
	input := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`[-0-9]+`)
	for input.Scan() {
		s := strings.SplitN(input.Text(), " ", 2)
		m := re.FindAllString(s[1], 6)
		if len(m) == 6 {
			r := region{
				min: point{atoi(m[0]), atoi(m[2]), atoi(m[4])},
				max: point{atoi(m[1]), atoi(m[3]), atoi(m[5])},
				on:  s[0] == "on",
			}
			if r.min.x < -50 || r.min.y < -50 || r.min.z < -50 || r.max.x > 50 || r.max.y > 50 || r.max.z > 50 {
				continue
			}
			grid = append(grid, r)
		}
	}
	fmt.Println(grid)
	part1(grid)
}

func part1(cuboids []region) {
	grid := make(map[point]struct{})

	for _, c := range cuboids {
		for _, p := range c.points() {
			if c.on {
				grid[p] = struct{}{}
			} else {
				delete(grid, p)
			}
		}
	}
	fmt.Println(len(grid))
}

func (r region) points() []point {
	fmt.Println(r, r.count())
	ret := make([]point, r.count())
	for x := r.min.x; x <= r.max.x; x++ {
		for y := r.min.y; y <= r.max.y; y++ {
			for z := r.min.z; z <= r.max.z; z++ {
				ret = append(ret, point{x, y, z})
			}
		}
	}
	return ret
}

func (r region) count() int {
	return (r.max.x - r.min.x) * (r.max.y - r.min.y) * (r.max.z - r.min.z)
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
