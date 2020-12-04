package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fieldValidation = map[string]func(string) bool{
	"byr": func(s string) bool { return checkNumRange(s, 1920, 2002) },
	"iyr": func(s string) bool { return checkNumRange(s, 2010, 2020) },
	"eyr": func(s string) bool { return checkNumRange(s, 2020, 2030) },
	"hgt": checkHeight,
	"hcl": func(s string) bool { return checkRegexp(s, "^#[0-9a-f]{6}$") },
	"ecl": func(s string) bool { return checkRegexp(s, "^(amb|blu|brn|gry|grn|hzl|oth)$") },
	"pid": func(s string) bool { return checkRegexp(s, "^[0-9]{9}$") },
}

func main() {
	haveFields := 0
	valid := 0
	pass := map[string]string{}
	var pairs []string

	check := func(pass map[string]string) {
		if hasRequiredFields(pass) {
			haveFields++

			if validateFields(pass) {
				valid++
			}
		}
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		pairs = strings.Fields(s.Text())
		fmt.Println(pairs)

		if len(pairs) == 0 {
			// blank line, we have read a passport.
			check(pass)
			pass = map[string]string{}
			continue
		}

		// Parse the paris of key:val into our p map
		for _, p := range pairs {
			fields := strings.Split(p, ":")
			if len(fields) == 2 {
				pass[fields[0]] = fields[1]
			}
		}
	}

	// check final passport in the file
	if len(pairs) > 0 {
		check(pass)
	}

	fmt.Println(haveFields, "passports with required fields")
	fmt.Println(valid, "valid passports")
}

func hasRequiredFields(p map[string]string) bool {
	for k, _ := range fieldValidation {
		if _, ok := p[k]; !ok {
			return false
		}
	}
	return true
}

func validateFields(p map[string]string) bool {
	for k, f := range fieldValidation {
		if !f(p[k]) {
			fmt.Println("invalid: ", k, p[k])
			return false
		}
	}
	return true
}

func checkNumRange(s string, min, max int) bool {
	if n, err := strconv.Atoi(s); err == nil {
		return n >= min && n <= max
	}
	return false
}

func checkHeight(s string) bool {
	if len(s) < 3 {
		return false
	}

	var min, max int
	switch s[len(s)-2:] {
	case "in":
		min = 59
		max = 76
	case "cm":
		min = 150
		max = 193
	default:
		return false
	}

	return checkNumRange(s[0:len(s)-2], min, max)
}

func checkRegexp(s, r string) bool {
	if ok, err := regexp.MatchString(r, s); err == nil {
		return ok
	}
	return false
}
