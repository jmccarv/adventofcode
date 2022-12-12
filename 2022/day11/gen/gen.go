package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fh, err := os.Create("monkies.go")
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	fmt.Fprintf(fh, `// Code generated by elves, DO NOT EDIT.
package main

type monkey struct {
	nr int
	itemQueue []int
	op func(old int)
	inspected int
}

var p1monkies []*monkey
var p2monkies []*monkey

func init() {
`)
	var op, items, destTrue, destFalse string
	var divBy int
	lcm := 0
	x := 10
	monkey := -1

	for input.Scan() {
		if strings.HasPrefix(input.Text(), "Monkey") {
			monkey++
			x = 0
			continue
		}

		switch x {
		case 0: // starting items
			f := strings.Split(input.Text(), ":")
			items = f[1]
		case 1: // Operation
			f := strings.Split(input.Text(), "=")
			op = f[1]
		case 2: //test
			f := strings.Split(input.Text(), "by ")
			divBy, _ = strconv.Atoi(f[1])
			if lcm > 0 {
				lcm *= divBy
			} else {
				lcm = divBy
			}
		case 3: // If true
			f := strings.Fields(input.Text())
			destTrue = f[len(f)-1]
		case 4: // If false
			f := strings.Fields(input.Text())
			destFalse = f[len(f)-1]

			// we've read everything for the current monkey
			//fmt.Println("Monkey", monkey, items, op, divBy, destTrue, destFalse)
			fmt.Fprintf(fh, `
p1monkies = append(p1monkies, &monkey{
	nr: %d,
	itemQueue: []int{%s},
	op: func(old int) {
		old = (%s) / 3
		if old %% %d == 0 {
			p1monkies[%s].itemQueue = append(p1monkies[%s].itemQueue, old)
		} else {
			p1monkies[%s].itemQueue = append(p1monkies[%s].itemQueue, old)
		}
	},
})
p2monkies = append(p2monkies, &monkey{
	nr: %d,
	itemQueue: []int{%s},
	op: func(old int) {
		old = (%s) %% lcm
		if old %% %d == 0 {
			p2monkies[%s].itemQueue = append(p2monkies[%s].itemQueue, old)
		} else {
			p2monkies[%s].itemQueue = append(p2monkies[%s].itemQueue, old)
		}
	},
})
`, monkey, items, op, divBy, destTrue, destTrue, destFalse, destFalse,
				monkey, items, op, divBy, destTrue, destTrue, destFalse, destFalse)
		}
		x++
	}

	fmt.Fprintf(fh, `
}

var lcm = %d

`, lcm)
}
