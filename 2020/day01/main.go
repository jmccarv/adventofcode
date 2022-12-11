package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var nums sort.IntSlice

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, i)
	}

	nums.Sort()

	a, b := find2(nums)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	a, b, c := find3(nums)
	fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
}

func find2(nums []int) (int, int) {
	for len(nums) > 1 {
		a := nums[0]
		nums = nums[1:]

		for _, b := range nums {
			if a+b == 2020 {
				return a, b
			} else if a+b > 2020 {
				break
			}
		}
	}
	return 0, 0
}

func find3(nums []int) (int, int, int) {
	for a := range nums {
		for b := range nums[a+1:] {
			for c := range nums[b+1:] {
				if nums[a]+nums[b]+nums[c] == 2020 {
					return nums[a], nums[b], nums[c]
				} else if nums[a]+nums[b]+nums[c] > 2020 {
					break
				}
			}
		}
	}

	return 0, 0, 0
}
