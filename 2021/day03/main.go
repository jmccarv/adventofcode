package main

import (
	"bufio"
	"fmt"
	"os"
)

type data struct {
	str  string
	bits uint64
}

type common struct {
	most  uint64
	least uint64
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var input []data

	for s.Scan() {
		var nr uint64 = 0
		for _, c := range s.Text() {
			nr <<= 1
			if c == '1' {
				nr |= 1
			}
		}
		input = append(input, data{s.Text(), nr})
	}

	if len(input) == 0 {
		panic("No data")
	}

	part1(input)
	part2(input)
}

func calcCommon(input []data) common {
	var c common
	nrBits := len(input[0].str)
	mc := make([]int, nrBits)

	for _, d := range input {
		bp := len(d.str)
		for i := 0; i < len(d.str); i++ {
			bp--
			if d.bits&(1<<bp) == 0 {
				mc[i] -= 1
			} else {
				mc[i] += 1
			}
		}
	}

	for _, i := range mc {
		c.most <<= 1
		if i >= 0 {
			c.most |= 1
		}
	}

	c.least = ^c.most & (1<<nrBits - 1)
	return c
}

func part1(input []data) {
	c := calcCommon(input)
	fmt.Println(c.most * c.least)
}

func remove(x []data, ofs int) []data {
	if ofs >= len(x) {
		return x
	}
	x[ofs] = x[len(x)-1]
	return x[:len(x)-1]
}

func part2(input []data) {
	sieve := func(filter func(c common) uint64) uint64 {
		nums := append([]data{}, input...)

		for mask := uint64(1) << (len(input[0].str) - 1); mask > 0; {
			c := calcCommon(nums)

			for i := 0; i < len(nums); {
				if nums[i].bits&mask != filter(c)&mask {
					nums = remove(nums, i)
					continue
				}
				i++
			}
			if len(nums) == 1 {
				return nums[0].bits
			}

			mask >>= 1
		}
		return 0
	}

	ogr := sieve(func(c common) uint64 { return c.most })
	csr := sieve(func(c common) uint64 { return c.least })
	fmt.Println(ogr * csr)
}
