package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type rangeMap struct {
	srcStart  []int
	destStart []int
	length    []int
}

func (rm *rangeMap) Add(src, dest, length int) {
	rm.srcStart = append(rm.srcStart, src)
	rm.destStart = append(rm.destStart, dest)
	rm.length = append(rm.length, length)
}

func (rm *rangeMap) Len() int {
	return len(rm.srcStart)
}

func (rm *rangeMap) String() string {
	var ret string
	for i := 0; i < rm.Len(); i++ {
		ret += fmt.Sprintf("%d %d %d\n", rm.srcStart[i], rm.destStart[i], rm.length[i])
	}
	return ret
}

func (rm *rangeMap) Lookup(src int) int {
	for i, start := range rm.srcStart {
		//fmt.Println("lookup", src, i, start, rm.destStart[i], rm.length[i])
		if start <= src && src < start+rm.length[i] {
			//fmt.Println("found")
			return rm.destStart[i] + src - start
		}
	}
	return src
}

func main() {
	var maps []*rangeMap
	var seeds []int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	for _, nr := range strings.Fields(s.Text())[1:] {
		x, _ := strconv.Atoi(nr)
		seeds = append(seeds, x)
	}

	var src, dest, length int
	rm := &rangeMap{}
	for s.Scan() {
		n, err := fmt.Sscan(s.Text(), &dest, &src, &length)
		if err == nil && n == 3 {
			rm.Add(src, dest, length)
		} else {
			if rm.Len() > 0 {
				maps = append(maps, rm)
				rm = &rangeMap{}
			}
		}

	}
	if rm.Len() > 0 {
		maps = append(maps, rm)
	}

	// part 1
	p1 := math.MaxInt
	for _, seed := range seeds {
		fmt.Println("seed", seed, mapSeed(seed, maps))
		p1 = min(p1, mapSeed(seed, maps))
	}

	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println("seed range", seeds[i], seeds[i+1])
		wg.Add(1)
		go func(i int) {
			nr := math.MaxInt
			for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
				nr = min(nr, mapSeed(seed, maps))
			}
			ch <- nr
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	p2 := math.MaxInt
	for nr := range ch {
		p2 = min(p2, nr)
	}
	fmt.Println("Part 1", p1)
	fmt.Println("Part 2", p2)
}

func mapSeed(src int, maps []*rangeMap) int {
	for _, m := range maps {
		src = m.Lookup(src)
		//fmt.Println("got:", i, src)
	}
	return src
}
