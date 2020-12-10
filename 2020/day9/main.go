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

	nr := part1(preamble, nums)
	fmt.Println("part1:", nr)
	fmt.Println("part2:", part2(nr, nums))
}

func part1(preamble int, input []int) int {
	// We only need to keep preamble-1 nodes in our ring
	// in order to keep track of all the necessary sums.
	// If preamble is 4, the ring will look like:
	// rb -> n1{ } -> n2{ } -> n3{ } -> n1
	rb := newRing(preamble - 1)

	// Prepouplate ring with all our sums
	for i, n1 := range input[0 : preamble-1] {
		rb.num = n1
		for _, n2 := range input[i+1 : preamble] {
			rb.sums = append(rb.sums, n1+n2)
		}
		rb = rb.next
	}

	// prev holds the next number we'll add to the ring.
	prev := input[preamble-1]

	for _, n := range input[preamble:] {
		fmt.Println(rb)

		if !rb.findSum(n) {
			return n
		}

		// Add the next number to the ring and hold on to our current n for next time
		rb = rb.add(prev)
		prev = n

		// Now compute sums of this number (n) + all others in the ring
		rb.walk(func(node *ringNode) {
			node.sums = append(node.sums, node.num+n)
		})
	}

	return 0
}

func part2(nr int, input []int) int {
	for i, n1 := range input[0:] {
		sum, min, max := n1, n1, n1

		for _, n2 := range input[i+1:] {
			if n2 < min {
				min = n2
			}
			if n2 > max {
				max = n2
			}

			sum += n2

			if sum == nr {
				return min + max
			}

			if sum > nr {
				break
			}
		}
	}
	return 0
}

// Create a new ring buffer with nrNodes nodes
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

// Walk each node in the ring calling callback() on each one
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

// Overwrites 'head' of the ring with num and advances
// the head of the ring to the next node.
// Returns a pointer to the new 'head' of the ring
func (r *ringNode) add(num int) *ringNode {
	r.num = num
	r.sums = []int{}
	return r.next
}

// Search the ring for the given sum. Returns true if found.
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
