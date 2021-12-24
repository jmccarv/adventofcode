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
	b   bounds
	lit map[point]bool // map of all lit pixels
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
	//fmt.Println(xform)
	//return

	r := 0
	for input.Scan() {
		if len(input.Text()) == 0 {
			// ignore blank lines
			continue
		}

		img.b.max.c = max(img.b.max.c, len(input.Text()))
		for c, p := range input.Text() {

			if p == '#' {
				img.light(point{r, c})
			}
		}
		r++
	}
	img.b.min = point{-2, -2}
	img.b.max.c += 2
	img.b.max.r = r + 1

	img.dump()
	fmt.Println(img)

	img.solve()
	img.dump()
	fmt.Println(img)

	img.solve()
	img.dump()
	fmt.Println(img)

	nr := 0
	for r := img.b.min.r; r <= img.b.max.r; r++ {
		for c := img.b.min.c; c <= img.b.max.c; c++ {
			if img.lit[point{r, c}] {
				nr++
			}
		}
	}
	fmt.Println(nr)
}

func (img *image) solve() {
	n := image{b: img.b, lit: make(map[point]bool)}

	for r := img.b.min.r - 1; r <= img.b.max.r+1; r++ {
		for c := img.b.min.c - 1; c <= img.b.max.c+1; c++ {
			if img.enhanced(point{r, c}) {
				n.light(point{r, c})
			}
		}
	}
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

func (i *image) dark(p point) {
	delete(i.lit, p)
}

func (i image) String() string {
	return fmt.Sprintf("min: %v max: %v lit: %d", i.b.min, i.b.max, len(i.lit))
}

func (i *image) dump() {
	fmt.Println()
	for r := i.b.min.r; r <= i.b.max.r; r++ {
		for c := i.b.min.c; c <= i.b.max.c; c++ {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxPoint(a, b point) point {
	a.r = max(a.r, b.r)
	a.c = max(a.c, b.c)
	return a
}

func minPoint(a, b point) point {
	a.r = min(a.r, b.r)
	a.c = min(a.c, b.c)
	return a
}
