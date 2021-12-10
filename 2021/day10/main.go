package main

import (
	"bufio"
	"fmt"
	"os"
)

var closerFor = map[byte]byte{
	'[': ']',
	'(': ')',
	'{': '}',
	'<': '>',
}

var points = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

type tokenStack []byte

func main() {
	s := bufio.NewScanner(os.Stdin)
	var nextClose tokenStack
	ret := 0

	for s.Scan() {
		for _, tok := range s.Text() {
			tok := byte(tok)
			if c, ok := closerFor[tok]; ok {
				nextClose.push(c)
			} else if tok != nextClose.pop() {
				//fmt.Printf("Invalid closing token: %c %s\n", tok, nextClose)
				fmt.Printf("Invalid closing token: %c => %d\n", tok, points[tok])
				ret += points[tok]
				break
			}
		}
	}
	fmt.Println(ret)
}

func (s *tokenStack) push(tok byte) {
	*s = append(*s, tok)
	//fmt.Printf("push -> %s\n", s)
}

func (s *tokenStack) pop() byte {
	var ret byte
	l := len(*s)

	if len(*s) > 0 {
		ret = (*s)[l-1]
	}

	if len(*s) < 2 {
		*s = []byte{}
	} else {
		*s = (*s)[:l-1]
	}

	//fmt.Printf(" pop -> %s\n", s)
	return ret
}

func (s *tokenStack) String() string {
	return string(*s)
}
