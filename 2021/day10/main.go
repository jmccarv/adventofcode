package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var closerFor = map[byte]byte{
	'[': ']',
	'(': ')',
	'{': '}',
	'<': '>',
}

var p1Points = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var p2Points = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

type tokenStack []byte

func main() {
	s := bufio.NewScanner(os.Stdin)
	var incomplete []tokenStack
	p1 := 0

	for s.Scan() {
		good := true
		var nextClose tokenStack
		for _, tok := range s.Text() {
			tok := byte(tok)
			if c, ok := closerFor[tok]; ok {
				nextClose.push(c)
			} else if tok != nextClose.pop() {
				//fmt.Printf("Invalid closing token: %c => %d\n", tok, points[tok])
				p1 += p1Points[tok]
				good = false
				break
			}
		}
		if good {
			incomplete = append(incomplete, nextClose.clone())
		}
	}
	fmt.Println("p1", p1)

	// part 2
	var scores []int
	for _, i := range incomplete {
		score := 0
		for j := len(i) - 1; j >= 0; j-- {
			score = score*5 + p2Points[i[j]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println("p2", scores[len(scores)/2])
}

func (s tokenStack) clone() tokenStack {
	x := make([]byte, len(s))
	copy(x, s)
	return x
}

func (s *tokenStack) push(tok byte) {
	*s = append(*s, tok)
	//fmt.Printf("push -> %s\n", s)
}

func (s *tokenStack) pop() byte {
	var ret byte
	l := len(*s)

	if l > 0 {
		ret = (*s)[l-1]
	}
	*s = (*s)[:l-1]

	//fmt.Printf(" pop -> %s\n", s)
	return ret
}

func (s *tokenStack) String() string {
	return string(*s)
}
