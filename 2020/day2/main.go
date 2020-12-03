package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var v1, v2 int

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := strings.FieldsFunc(s.Text(), func(r rune) bool {
			return r == '-' || r == ' ' || r == ':'
		})
		if len(fields) != 4 {
			log.Fatal("Invalid input", s.Text())
		}

		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal("Invalid number", fields[0])
		}

		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal("Invalid number", fields[1])
		}

		l := fields[2][0]
		pwd := fields[3]

		if valid1(n1, n2, l, pwd) {
			v1++
		}

		if valid2(n1, n2, l, pwd) {
			v2++
		}
	}

	fmt.Println("part 1", v1, "valid passwords")
	fmt.Println("part 2", v2, "valid passwords")
}

func valid1(n1, n2 int, l byte, pwd string) bool {
	nr := strings.Count(pwd, string(l))
	return nr >= n1 && nr <= n2
}

func valid2(n1, n2 int, l byte, pwd string) bool {
	return (pwd[n1-1] == l) != (pwd[n2-1] == l)
}
