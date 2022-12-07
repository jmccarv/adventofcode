package main

import (
	"bufio"
	"fmt"
	"os"
)

type file struct {
	name   string
	size   int
	parent *dir
}

type dir struct {
	name   string
	files  []file
	size   int
	dirs   map[string]*dir
	parent *dir
}

var root = &dir{name: "/", dirs: make(map[string]*dir)}

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
			if cwd, ok = cwd.dirs[dname]; !ok {
				panic(fmt.Sprintf("Directory does not exist: %s", dname))
			}

		} else if nr, _ := fmt.Sscanf(s.Text(), "dir %s", &dname); nr == 1 {
			if _, ok := cwd.dirs[dname]; !ok {
				cwd.dirs[dname] = &dir{name: dname + "/", parent: cwd, dirs: make(map[string]*dir)}
			}

		} else if nr, _ := fmt.Sscanf(s.Text(), "%d %s", &size, &fn); nr == 2 {
			cwd.files = append(cwd.files, file{name: fn, size: size, parent: cwd})
		}
	}

	root.calc()
	//root.disp("", 0)
	part1()
	part2()
}

func part1() {
	sum := 0
	root.visitDirs(func(d dir) {
		if d.size <= 100000 {
			sum += d.size
		}
	})
	fmt.Println(sum)
}

func part2() {
	fsSize := 70000000
	needFree := 30000000
	needDel := needFree - (fsSize - root.size)

	min := *root
	root.visitDirs(func(d dir) {
		if d.size >= needDel && d.size < min.size {
			min = d
		}
	})
	fmt.Printf("%s %d\n", min.name, min.size)
}

func (d *dir) visitDirs(f func(dir)) {
	f(*d)
	for _, sub := range d.dirs {
		sub.visitDirs(f)
	}
}

func (d *dir) calc() int {
	d.size = 0

	for _, f := range d.files {
		d.size += f.size
	}

	for _, sub := range d.dirs {
		d.size += sub.calc()
	}
	return d.size
}

func (d *dir) disp(parent string, lvl int) {

	fmt.Printf("%*s%s%s %d\n", lvl, "", parent, d.name, d.size)

	for _, sub := range d.dirs {
		sub.disp(parent+d.name, lvl+2)
	}
}
