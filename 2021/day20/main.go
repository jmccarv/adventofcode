package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	r, c int
}

type bounds struct {
	min point
	max point
}

type image struct {
	inner bounds
	outer bounds
	lit   map[point]bool // map of all lit pixels
}

var xform [512]bool

func main() {
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		panic("Invalid input")
	}
	ea := input.Text()
	if len(ea) != 512 {
		panic("Invalid input")
	}

	img := image{lit: make(map[point]bool)}

	for i, c := range input.Text() {
		if c == '#' {
			xform[i] = true
		}
	}

	r := 0
	for input.Scan() {
		if len(input.Text()) == 0 {
			// ignore blank lines
			continue
		}

		img.inner.max.c = max(img.inner.max.c, len(input.Text())-1)
		for c, p := range input.Text() {

			if p == '#' {
				img.light(point{r, c})
			}
		}
		r++
	}
	img.inner.max.r = r - 1
	img.outer.min = point{-100, -100}
	img.outer.max.r = img.inner.max.r + 100
	img.outer.max.c = img.inner.max.c + 100

	// part 1
	//img.dump()
	fmt.Println(img)
	for i := 0; i < 2; i++ {
		img.iterate()
		fmt.Println(img)
	}
	part1 := img.nrLit()

	// part 2
	for i := 0; i < 48; i++ {
		img.iterate()
	}
	fmt.Println(img)
	part2 := img.nrLit()
	fmt.Println("part1", part1)
	fmt.Println("part2", part2)
}

func (img *image) iterate() {
	n := *img
	n.lit = make(map[point]bool)

	for r := n.outer.min.r - 1; r <= n.outer.max.r+1; r++ {
		for c := n.outer.min.c - 1; c <= n.outer.max.c+1; c++ {
			if img.enhanced(point{r, c}) {
				n.light(point{r, c})
			}
		}
	}

	n.inner.min.r--
	n.inner.min.c--
	n.inner.max.r++
	n.inner.max.c++
	*img = n
}

func (i *image) light(p point) {
	i.lit[p] = true
}

func (i *image) enhanced(p point) bool {
	idx := 0
	for r := p.r - 1; r <= p.r+1; r++ {
		for c := p.c - 1; c <= p.c+1; c++ {
			idx <<= 1
			if i.lit[point{r, c}] {
				idx |= 1
			}
		}
	}
	return xform[idx]
}

func (i image) nrLit() (nr int) {
	for r := i.inner.min.r; r <= i.inner.max.r; r++ {
		for c := i.inner.min.c; c <= i.inner.max.c; c++ {
			if i.lit[point{r, c}] {
				nr++
			}
		}
	}
	return
}

func (i image) String() string {
	return fmt.Sprintf("inner.min:%v inner.max:%v outer.min:%v outer.max:%v lit:%d", i.inner.min, i.inner.max, i.outer.min, i.outer.max, i.nrLit())
}

func (i *image) dump() {
	fmt.Println()
	for r := i.inner.min.r; r <= i.inner.max.r; r++ {
		for c := i.inner.min.c; c <= i.inner.max.c; c++ {
			if i.lit[point{r, c}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
