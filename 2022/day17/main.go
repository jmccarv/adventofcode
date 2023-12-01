package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	sm "github.com/jmccarv/adventofcode/util/math"
	s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

var rocks = []s2d.Shape{
	s2d.NewShapeFromString([]string{"####"}, '#'),
	s2d.NewShapeFromString([]string{
		" #",
		"###",
		" #"}, '#'),
	s2d.NewShapeFromString([]string{
		"  #",
		"  #",
		"###"}, '#'),
	s2d.NewShapeFromString([]string{"#", "#", "#", "#"}, '#'),
	s2d.NewShapeFromString([]string{"##", "##"}, '#'),
}

type historyKey struct {
	shape int
	jet   int
}

type history struct {
	rockNr int
	height int
}

var hist = make(map[historyKey][]history)

func main() {
	t0 := time.Now()

	nextRock := nextFunc(rocks)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	nextJet := nextFunc([]byte(input.Text()))

	// shaft is 7 units wide (0 - 6)
	// each shape starts at x:2 and y is top of highest (most negative) shape-3 (or floor-3 if no shape present)
	//
	// We're going to call the floor y=-1 to make height calculation easier
	shaft := s2d.ShapeField{Bounds: s2d.Box{TL: s2d.Point{0, math.MinInt}, BR: s2d.Point{6, -1}}}
	height := shaft.Bounds.BR.Y + 1

	for rockNr := 1; rockNr <= 2022; rockNr++ {
		ri, r := nextRock()
		//fmt.Println("ROCK", r)
		r.Loc = s2d.Point{2, height - r.Size.Y - 3}

		rock := shaft.AddShape(r)
		first := true
		//fmt.Println("hight:", height, " sizeY:", rock.Size.Y, "locY:", rock.Loc.Y)

		// Now we alternate between being pushed by a jet and falling one space
		// until we can fall no farther
		//fmt.Println(shaft.DumpWindow(s2d.Box{s2d.Point{0, rock.Loc.Y}, shaft.Bounds.BR}))
		var ji int
		var jet byte
		for {
			dir := s2d.Point{}
			ji, jet = nextJet()

			if first {
				if hval, ok := hist[historyKey{ri, ji}]; ok {
					lh := hval[len(hval)-1]
					cycle := rockNr - lh.rockNr
					p1 := (2022 - rockNr) / cycle * sm.Abs(height-lh.height)
					fmt.Printf("SEEN: %d, %d, %d, %d, %d %d -- %d\n", rockNr, cycle, ri, ji, hval, height, p1)

					//rockNr = p1
				}
				hist[historyKey{ri, ji}] = append(hist[historyKey{ri, ji}], history{rockNr, height})
				first = false
			}

			switch jet {
			case '<':
				dir.X = -1
			case '>':
				dir.X = 1
			}

			shaft.StepShape(rock, dir)
			if ok := shaft.StepShape(rock, s2d.Point{0, 1}); !ok {
				break
			}
		}

		//fmt.Println("rock loc now:", rock.Loc)
		height = sm.Min(height, rock.Loc.Y)
	}

	fmt.Println("Part 1 Height:", -height)

	/*
		for ; i < 1000000000000; i++ {
			run()
		}
		fmt.Println("Part 2 Height:", -height)
	*/

	fmt.Println("Total Time", time.Now().Sub(t0))
}

func nextFunc[T any](list []T) func() (int, T) {
	next := 0
	return func() (i int, t T) {
		i = next
		t = list[next]
		next = (next + 1) % len(list)
		return
	}
}
