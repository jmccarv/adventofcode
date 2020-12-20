package main

import (
	//"fmt"
	"testing"
)

type gameTest struct {
	numbers    []int
	iterations int
	expect     int
}

var tests = []gameTest{
	{[]int{0, 3, 6}, 2020, 436},
	{[]int{1, 3, 2}, 2020, 1},
	{[]int{2, 1, 3}, 2020, 10},
	{[]int{1, 2, 3}, 2020, 27},
	{[]int{2, 3, 1}, 2020, 78},
	{[]int{3, 2, 1}, 2020, 438},
	{[]int{3, 1, 2}, 2020, 1836},
	{[]int{1, 20, 8, 12, 0, 14}, 2020, 492},
	{[]int{0, 3, 6}, 30000000, 175594},
	{[]int{1, 3, 2}, 30000000, 2578},
	{[]int{2, 1, 3}, 30000000, 3544142},
	{[]int{1, 2, 3}, 30000000, 261214},
	{[]int{2, 3, 1}, 30000000, 6895259},
	{[]int{3, 2, 1}, 30000000, 18},
	{[]int{3, 1, 2}, 30000000, 362},
	{[]int{1, 20, 8, 12, 0, 14}, 30000000, 63644},
}

func TestPart1(t *testing.T) {
	for _, test := range tests {
		if ret := part1(test.iterations, test.numbers); ret != test.expect {
			t.Errorf("%+v, got %d\n", test, ret)
		}
	}
}
