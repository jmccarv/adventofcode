package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ringNode struct {
	num  int
	sums []int
	next *ringNode
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s preamble_length", os.Args[0])
	}
	preamble, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid preamble length: ", os.Args[1])
	}

	var nums []int
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		x, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal("Invalid input", s.Text())
		}
		nums = append(nums, x)
	}

	fmt.Println("part1:", part1(preamble, nums))
}

func part1(preamble int, input []int) int {
	rb := newRing(preamble - 1)

	for i, n1 := range input[0 : preamble-1] {
		rb.num = n1
		for _, n2 := range input[i+1 : preamble] {
			rb.sums = append(rb.sums, n1+n2)
		}
		rb = rb.next
	}

	prev := input[preamble-1]
	for _, n := range input[preamble:] {
		fmt.Println(rb)

		if !rb.findSum(n) {
			return n
		}

		// Add our new number to our ring, which we keep in prev
		rb = rb.add(prev)
		prev = n

		// Now compute sums of this number + all others in the ring
		rb.walk(func(node *ringNode) {
			node.sums = append(node.sums, node.num+n)
			fmt.Println("walk: n=", n, "  num=", node.num, "  sums:", node.sums)
		})
	}

	return 0
}

func newRing(nrNodes int) *ringNode {
	rb := &ringNode{}
	node := rb

	for i := 1; i < nrNodes; i++ {
		rn := &ringNode{}
		node.next = rn
		node = rn
	}
	node.next = rb

	return rb
}

func (r *ringNode) String() string {
	ret := ""
	r.walk(func(n *ringNode) {
		ret += fmt.Sprintf("%d : %+v\n", n.num, n.sums)
	})
	return ret
}

// If callback returns false, traversal halts and walk() returns false
// Returns true if we completed our walk
func (r *ringNode) walk(callback func(*ringNode)) {
	p := r
	for {
		callback(p)
		if p.next == r {
			return
		}
		p = p.next
	}
}

// Returns a pointer to the newly added node (now at the 'end' of the ring)
// and the new 'head' of the ring
func (r *ringNode) add(num int) *ringNode {
	r.num = num
	r.sums = []int{}
	return r.next
}

func (r *ringNode) findSum(num int) bool {
	found := false

	r.walk(func(node *ringNode) {
		for _, x := range node.sums {
			if x == num {
				found = true
			}
		}
	})

	return found
}
