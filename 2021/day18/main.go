package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	literal = iota
	value
)

type token struct {
	typ  int
	char rune
	val  int
}

type number []token

func main() {
	s := bufio.NewScanner(os.Stdin)
	var nums []number
	for s.Scan() {
		nums = append(nums, parseNum(s.Text()))
	}
	//fmt.Println(nums)
	part1(nums)
	part2(nums)

	//testExplode(nums)
}

/*
func testExplode(nums []number) {
	for _, n := range nums {
		fmt.Println(n)
		fmt.Println(n.reduce())
		fmt.Println()
	}
}
*/

func part1(nums []number) {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		n = n.add(nums[i])
	}
	fmt.Println(n.magnitude())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part2(nums []number) {
	m := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			m = max(m, nums[i].add(nums[j]).magnitude())
			m = max(m, nums[j].add(nums[i]).magnitude())
		}
	}
	fmt.Println(m)
}

func (n number) add(n1 number) number {
	//fmt.Println("Add", n, "+", n1)
	ret := number{token{typ: literal, char: '['}}
	ret = append(ret, n...)
	ret = append(ret, token{typ: literal, char: ','})
	ret = append(ret, n1...)
	ret = append(ret, token{typ: literal, char: ']'})
	//fmt.Println("=(unreduced)", ret)
	ret = ret.reduce()
	//fmt.Println("=", ret)
	return ret
}

// this is ugly. I don't care.
func (n number) reduce() number {
	done := false

	for !done {
		done = true
		n1 := number{}
		depth := 0
		//fmt.Println(n)
		for i, t := range n {
			if t.typ == literal {
				switch t.char {
				case '[':
					depth++
				case ']':
					depth--
				}
			} else {
				if depth > 4 {
					// explode
					done = false
					//fmt.Printf("E %d ", i)
					for j := i - 1; j > 0; j-- {
						if n[j].typ == value {
							n[j].val += t.val
							break
						}
					}
					for j := i + 3; j < len(n); j++ {
						if n[j].typ == value {
							n[j].val += n[i+2].val
							break
						}
					}
					n[i-1] = n[i]
					n[i-1].val = 0
					n1 = append(n1, n[0:i]...)
					n1 = append(n1, n[i+4:]...)
					n = n1
					break
				}
			}
		}

		if done {
			for i, t := range n {
				if t.val > 9 {
					// split
					done = false
					//fmt.Printf("S %d ", i)
					l := t.val / 2
					r := t.val - l

					n1 = append(n1, n[0:i]...)
					n1 = append(n1, token{typ: literal, char: '['})
					n1 = append(n1, token{typ: value, val: l})
					n1 = append(n1, token{typ: literal, char: ','})
					n1 = append(n1, token{typ: value, val: r})
					n1 = append(n1, token{typ: literal, char: ']'})
					n1 = append(n1, n[i+1:]...)
					n = n1
					break
				}
			}
		}
	}
	return n
}

type magStack []int

func (n number) magnitude() int {
	m := &magStack{}

	for _, t := range n {
		if t.typ == value {
			m.push(t.val)
			continue
		}
		switch t.char {
		case ']':
			m.push(m.pop()*2 + m.pop()*3)
		}
	}
	//fmt.Println("stack len", len(*m))
	return m.pop()
}

func (m *magStack) push(x int) {
	*m = append(*m, x)
}

func (m *magStack) pop() int {
	ret := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return ret
}

func parseNum(s string) number {
	var num []token
	for i, r := range s {
		switch r {
		case '[', ']', ',':
			num = append(num, token{typ: literal, char: r})
		default:
			num = append(num, token{typ: value, val: atoi(s[i : i+1])})
		}
	}
	return num
}

func atoi(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func (n number) String() (s string) {
	for _, t := range n {
		switch t.typ {
		case literal:
			s = s + string(t.char)
		case value:
			s = s + fmt.Sprintf("%d", t.val)
		}
	}
	return
}
