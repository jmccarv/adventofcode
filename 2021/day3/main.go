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
	mc := make([]int, len(input[0].str))

	for _, d := range input {
		bp := len(d.str)
		for i := 0; i < len(d.str); i++ {
			bp--
			Δ := 1
			if d.bits&(1<<bp) == 0 {
				Δ = -1
			}
			mc[i] += Δ
		}
	}

	var mask uint64
	for _, i := range mc {
		mask <<= 1
		mask |= 1

		c.most <<= 1
		if i >= 0 {
			c.most |= 1
		}
	}
	c.least = ^c.most & mask
	return c
}

func part1(input []data) {
	c := calcCommon(input)
	fmt.Println(c.most * c.least)
}

func part2(input []data) {
	sieve := func(filter func(c common) uint64) uint64 {
		var mask uint64 = 1 << (len(input[0].str) - 1)
		var nums []data
		for _, x := range input {
			nums = append(nums, x)
		}

		for mask > 0 {
			var newNums []data
			c := calcCommon(nums)

			for _, d := range nums {
				want := filter(c) & mask
				if d.bits&mask == want {
					newNums = append(newNums, d)
				}
			}
			nums = newNums
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
