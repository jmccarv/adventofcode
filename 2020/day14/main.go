package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"regexp"
	"strconv"
)

type mask struct {
	or    uint64
	and   uint64
	float uint64
}

type instruction struct {
	msk mask   // mask currently in effect
	loc uint64 // location for 'mem' action
	val uint64 // value for to set in memory for 'mem' action
}

type instructions []instruction

const memMask = 0xFFFFFFFFF // 36 bit memory size

func main() {
	i := parse(os.Stdin)
	fmt.Println("part1:", i.part1())
	fmt.Println("part2:", i.uglyPart2())
}

var rInstr = regexp.MustCompile(`^\s*(mem\[([\d]+)\]|mask)\s*=\s*([X\d]+)`)

func (instr instructions) part1() uint64 {
	mem := make(map[uint64]uint64)

	for _, i := range instr {
		fmt.Println(i)
		mem[i.loc] = (i.val | i.msk.or) & i.msk.and
	}

	var sum uint64
	for _, x := range mem {
		sum += x
	}
	return sum
}

func (instr instructions) uglyPart2() uint64 {
	mem := make(map[uint64]bool)
	var sum uint64

	for i := len(instr) - 1; i >= 0; i-- {
		a := instr[i]
		if bits.OnesCount64(a.msk.float) > 24 {
			log.Fatal("Cowardly refusing to run part 2 on this input set -- too many floating addresses")
		}
		for _, loc := range a.addresses() {
			if _, ok := mem[loc]; !ok {
				sum += a.val
				mem[loc] = true
			}
		}
	}

	return sum
}

func (a instruction) addresses() []uint64 {
	ret := make(map[uint64]bool)

	addr := a.loc | a.msk.or
	bit := uint64(1)
	for i := 0; i < 36; i++ {
		if a.msk.float&bit == bit {
			for m, _ := range ret {
				ret[m|bit] = true
				ret[m&(^bit)] = true
			}
			ret[addr|bit] = true
			ret[addr&(^bit)] = true
		}
		bit <<= 1
	}
	/*
		fmt.Printf("%8d %36b\n", a.loc, a.loc)
		fmt.Printf("%-8s %36b\n", "or:", a.msk.or)
		fmt.Printf("%8d %36b\n", addr, addr)
		fmt.Printf("%-8s %36b\n", "float:", a.msk.float)
		for m, _ := range ret {
			fmt.Printf("%8s %36b\n", "", m)
		}
		fmt.Println()
	*/

	var z []uint64
	for x := range ret {
		z = append(z, x)
	}
	return z
}

func parse(r io.Reader) instructions {
	var ret []instruction
	var err error

	s := bufio.NewScanner(r)
	i := instruction{}
	for s.Scan() {
		line := s.Text()
		m := rInstr.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		fmt.Printf("%+v\n", m)

		switch m[1][0:3] {
		case "mas":
			//fmt.Println("mask: ", m[3])
			i.msk = newMask(m[3])
		case "mem":
			i.loc, err = strconv.ParseUint(m[2], 10, 36)
			if err != nil {
				log.Fatal("Invalid input (mem location): ", line)
			}
			i.val, err = strconv.ParseUint(m[3], 10, 36)
			if err != nil {
				log.Fatal("Invalid input (mem value): ", line)
			}
			ret = append(ret, i)
		}
	}

	return ret
}

func newMask(m string) mask {
	ret := mask{and: memMask}

	for _, c := range m {
		ret.or <<= 1
		ret.and <<= 1
		ret.float <<= 1

		switch c {
		case 'X':
			ret.and |= 0x01
			ret.float |= 0x01
		case '1':
			ret.or |= 0x01
			ret.and |= 0x01
		case '0':
		default:
			log.Fatal("Invalid mask input: ", m)
		}
		//fmt.Printf("nm(%c)\n    %036b\n    %036b\n", c, ret.or, ret.and)
	}
	return ret
}
