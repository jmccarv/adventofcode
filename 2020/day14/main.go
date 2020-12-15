package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	iMask = iota
	iMem
)

type mask struct {
	or  uint64
	and uint64
}

type instruction struct {
	act int    // action to perform
	msk mask   // mask to set
	loc int    // location for 'mem' action
	val uint64 // value for to set in memory for 'mem' action
}

const memMask = 0xFFFFFFFFF // 36 bit memory size

func main() {
	mem := make(map[int]uint64)
	msk := newMask("")

	for _, i := range parse(os.Stdin) {
		//fmt.Printf("%+v\n", p)
		switch i.act {
		case iMask:
			msk = i.msk
		case iMem:
			mem[i.loc] = (i.val | msk.or) & msk.and
		}
	}

	var sum uint64
	for _, x := range mem {
		sum += x
	}
	fmt.Println("part1:", sum)
}

var rInstr = regexp.MustCompile(`^\s*(mem\[([\d]+)\]|mask)\s*=\s*([X\d]+)`)

func parse(r io.Reader) []instruction {
	var ret []instruction
	var err error

	s := bufio.NewScanner(r)
	for s.Scan() {
		i := instruction{}

		line := s.Text()
		m := rInstr.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		fmt.Printf("%+v\n", m)

		switch m[1][0:3] {
		case "mas":
			//fmt.Println("mask: ", m[3])
			i.act = iMask
			i.msk = newMask(m[3])
		case "mem":
			i.act = iMem
			i.loc, err = strconv.Atoi(m[2])
			if err != nil {
				log.Fatal("Invalid input (mem location): ", line)
			}
			i.val, err = strconv.ParseUint(m[3], 10, 36)
			if err != nil {
				log.Fatal("Invalid input (mem value): ", line)
			}
		}

		ret = append(ret, i)
	}

	return ret
}

func newMask(m string) mask {
	ret := mask{and: memMask}

	for _, c := range m {
		ret.or <<= 1
		ret.and <<= 1

		switch c {
		case 'X':
			ret.and |= 0x01
		case '1':
			ret.or |= 0x01
			ret.and |= 0x01
		case '0':
		default:
			log.Fatal("Invalid mask input: ", m)
		}
		//fmt.Printf("nm(%c)\n    %036b\n    %036b\n", c, ret.orMask, ret.andMask)
	}
	return ret
}
