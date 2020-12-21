package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type note struct {
	name  string
	valid map[int]struct{}
}

type ticket []int

type input struct {
	notes   []note
	me      ticket
	tickets []ticket
}

func main() {
	data := parse(os.Stdin)
	fmt.Printf("%+v\n", data)

	fmt.Println("part1: ", data.part1())
}

func parse(r io.Reader) input {
	data := input{}

	s := bufio.NewScanner(r)
	data.parseNotes(s)

	data.me = data.parseNextTicket(s)
	data.parseTickets(s)

	return data
}

func (d *input) part1() int {
	var ret int

	for _, t := range d.tickets {
		for _, num := range t {
			if d.findNote(num).valid == nil {
				ret += num
			}
		}
	}

	return ret
}

func (d *input) findNote(num int) note {
	for _, note := range d.notes {
		if _, ok := note.valid[num]; ok {
			fmt.Println("Found note for", num)
			return note
		}
	}
	return note{}
}

func (d *input) parseNotes(s *bufio.Scanner) {
	for s.Scan() {
		//fields := strings.Split(s.Text(), ":")
		text := s.Text()
		c := strings.Index(text, ":")
		if c < 2 {
			return
		}

		n := note{valid: make(map[int]struct{})}

		n.name = text[0:c]
		fields := text[c+1:]

		for _, r := range strings.Fields(fields) {
			if nums := strings.Split(r, "-"); len(nums) == 2 {
				n.addRange(nums[0], nums[1])
			}
		}

		d.notes = append(d.notes, n)
	}
}

func (n *note) addRange(low, high string) {
	l, err := strconv.Atoi(low)
	if err != nil {
		log.Fatal("Invalid low range: ", low)
	}
	h, err := strconv.Atoi(high)
	if err != nil {
		log.Fatal("Invalid hight range: ", high)
	}

	for ; l <= h; l++ {
		n.valid[l] = struct{}{}
	}
}

func (d *input) parseNextTicket(s *bufio.Scanner) ticket {
	var t ticket

	for s.Scan() {
		nums := strings.Split(s.Text(), ",")
		if len(nums) < 2 {
			continue
		}

		for _, nstr := range nums {
			num, err := strconv.Atoi(nstr)
			if err != nil {
				log.Fatal("Invalid ticket number: ", nstr)
			}
			t = append(t, num)
		}
		return t
	}

	return nil
}

func (d *input) parseTickets(s *bufio.Scanner) {
	for {
		t := d.parseNextTicket(s)
		if t == nil {
			return
		}

		d.tickets = append(d.tickets, t)
	}
}
