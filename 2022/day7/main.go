package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var name string
		var size int

		if nr, _ := fmt.Sscanf(s.Text(), "$ cd %s", &name); nr == 1 {
			switch name {
			case "/":
				cwd = root
			case "..":
				cwd = cwd.parent
			default:
				var ok bool
				if cwd, ok = cwd.ents[name]; !ok {
					panic(fmt.Sprintf("Directory does not exist: %s", name))
				}
			}

		} else if nr, _ := fmt.Sscanf(s.Text(), "dir %s", &name); nr == 1 {
			if _, ok := cwd.ents[name]; !ok {
				cwd.ents[name] = &ent{typ: E_DIR, name: name, parent: cwd, ents: make(map[string]*ent)}
			}

		} else if nr, _ := fmt.Sscanf(s.Text(), "%d %s", &size, &name); nr == 2 {
			cwd.ents[name] = &ent{typ: E_FILE, name: name, size: size, parent: cwd}
		}
	}

	root.calc()
	root.disp(0)
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
	fmt.Printf("%d\n", min.size)
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

func (e *ent) disp(lvl int) {
	fmt.Printf("%*s- %s (dir, size=%d)\n", lvl, "", e.name, e.size)

	var names []string
	for n, _ := range e.ents {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, n := range names {
		if e.ents[n].typ == E_DIR {
			e.ents[n].disp(lvl + 2)
		} else {
			fmt.Printf("%*s- %s (file, size=%d)\n", lvl+2, "", n, e.ents[n].size)
		}
	}
}
