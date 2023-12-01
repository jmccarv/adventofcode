package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	//sm "github.com/jmccarv/adventofcode/util/math"
	//s2d "github.com/jmccarv/adventofcode/util/simple2d"
)

var p1 int
var p2 int

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		part1(s.Text(), &p1)
		part2(s.Text(), &p2)
	}
	fmt.Println("part1", p1)
	fmt.Println("part2", p2)
}

func part1(text string, res *int) {
	var nums string
	for _, c := range text {
		if unicode.IsDigit(c) {
			nums += string(c)
		}
	}
	if len(nums) > 0 {
		*res += int((nums[0]-'0')*10) + int(nums[len(nums)-1]-'0')
	}
}

func part2(text string, res *int) {
	nums := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for ofs := 0; ofs < len(text)-1; ofs++ {
		left := text[0:ofs]
		right := text[ofs:]
		for k, v := range nums {
			if strings.Index(right, k) == 0 {
				// leave the last character of the number's name as it might overlap and we need to transform both
				// i.e. the string "eightwo"  must turn into "82" (well "82o") is thankfully also acceptible
				text = left + v + right[len(k)-1:]
				break
			}
		}
	}

	fmt.Println(text)
	part1(text, res)
}
