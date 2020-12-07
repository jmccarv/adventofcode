package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color   string
	carries map[string]int
}

var bags = make(map[string]bag)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		b := parseBag(s.Text())
		bags[b.color] = b
	}

	//fmt.Printf("%+v\n", bags)
	fmt.Println("part1", part1("shiny gold"))
	fmt.Println("part2", part2("shiny gold"))
}

func parseBag(line string) bag {
	var b bag

	l := strings.Split(line, " bags contain ")
	if len(l) != 2 {
		log.Fatal("Invalid input:", line)
	}

	b.color = l[0]
	b.carries = make(map[string]int)

	if l[1] == "no other bags." {
		return b
	}

	re := regexp.MustCompile(`(\d+)\s*(.*?)\s+bags?`)

	for _, carried := range strings.Split(l[1], ",") {
		m := re.FindStringSubmatch(carried)
		if len(m) != 3 {
			log.Fatal("(parse) Invalid input:", carried)
		}
		nr, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal("(numconv) Invalid input: ", m[1], " from ", carried)
		}
		b.carries[m[2]] = nr
	}

	//fmt.Println(b, "\n")
	return b
}

func part2(color string) int {
	ret := 0
	for c, nr := range bags[color].carries {
		ret += nr + nr*part2(c) // our bag carries nr bags of color c, so count nr c bags + all the bags c contains
	}
	return ret
}

func part1(color string) int {
	// We want to carry the bag of color 'color' in at least one other bag
	// Return how many different bag colors would be valid for the outermost bag

	// We loop through all the outer bags (those in the bags map)
	// and check all the bags contained in them to see if our passed in 'color'
	// is found nested in there...
	nr := 0
	for _, b := range bags {
		if find1(color, b) {
			nr++
		}
	}
	return nr
}

func find1(color string, b bag) bool {
	// Look through all the bags carried by 'b' and if any
	// of them contain the bag 'color' return true
	for k, _ := range b.carries {
		if color == k {
			return true
		} else {
			if find1(color, bags[k]) {
				return true
			}
		}
	}

	return false
}
