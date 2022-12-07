package main

import (
	"bufio"
	"fmt"
	"os"
)

type entType int

const (
	_ entType = iota
	E_FILE
	E_DIR
)

type ent struct {
	name   string
	typ    entType
	size   int
	ents   map[string]*ent
	parent *ent
}

var root = &ent{name: "/", typ: E_DIR, ents: make(map[string]*ent)}

func main() {
	cwd := root

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var fn, dname string
		var size int

		if nr, _ := fmt.Sscanf(s.Text(), "$ cd %s", &dname); nr == 1 {
			switch dname {
			case "/":
				cwd = root
				continue
			case "..":
				cwd = cwd.parent
				continue
			}

			var ok bool
			if cwd, ok = cwd.ents[dname]; !ok {
				panic(fmt.Sprintf("Directory does not exist: %s", dname))
			}

		} else if nr, _ := fmt.Sscanf(s.Text(), "dir %s", &dname); nr == 1 {
			if _, ok := cwd.ents[dname]; !ok {
				cwd.ents[dname] = &ent{typ: E_DIR, name: dname + "/", parent: cwd, ents: make(map[string]*ent)}
			}

		} else if nr, _ := fmt.Sscanf(s.Text(), "%d %s", &size, &fn); nr == 2 {
			cwd.ents[fn] = &ent{typ: E_FILE, name: fn, size: size, parent: cwd}
		}
	}

	root.calc()
	//root.disp("", 0)
	part1()
	part2()
}

func part1() {
	sum := 0
	root.visit(func(e *ent) {
		if e.typ == E_DIR && e.size <= 100000 {
			sum += e.size
		}
	})
	fmt.Println(sum)
}

func part2() {
	fsSize, needFree := 70000000, 30000000
	needDel := needFree - (fsSize - root.size)

	min := root
	root.visit(func(e *ent) {
		if e.typ == E_DIR && e.size >= needDel && e.size < min.size {
			min = e
		}
	})
	fmt.Printf("%s %d\n", min.name, min.size)
}

func (e *ent) visit(f func(*ent)) {
	for _, sub := range e.ents {
		sub.visit(f)
	}
	f(e)
}

func (e *ent) calc() {
	e.visit(func(x *ent) {
		if x.parent != nil {
			x.parent.size += x.size
		}
	})
}

func (e *ent) disp(parent string, lvl int) {
	fmt.Printf("%*s%s%s %d\n", lvl, "", parent, e.name, e.size)

	for _, sub := range e.ents {
		if sub.typ == E_DIR {
			sub.disp(parent+e.name, lvl+2)
		}
	}
}
