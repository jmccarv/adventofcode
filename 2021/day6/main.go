package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fish []int

func main() {
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		panic("Invalid input")
	}

	var school fish
	for _, n := range strings.Split(s.Text(), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic("Invalid input")
		}
		school = append(school, i)
	}
	//fmt.Println(f)

	part1(school)
}

func part1(school fish) {
	days := 80

	for ; days > 0; days-- {
		var spawn fish
		for i, f := range school {
			switch f {
			case 0:
				school[i] = 6
				spawn = append(spawn, 8)
			default:
				school[i]--
			}
		}
		school = append(school, spawn...)
	}
	fmt.Println(len(school))
}
