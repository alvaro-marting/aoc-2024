package main

import (
	"regexp"
	"strconv"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

var exp = regexp.MustCompile("(do\\(\\))|(don't\\(\\))|mul\\((\\d{1,3}),(\\d{1,3})\\)")

func main() {

	s := pkg.MustReadFile("./inputs/03.txt")

	matches := exp.FindAllStringSubmatch(s, -1)
	enabled := true
	acc := 0
	for _, match := range matches {
		if match[1] != "" {
			enabled = true
			continue
		}
		if match[2] != "" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		a, _ := strconv.Atoi(match[3])
		b, _ := strconv.Atoi(match[4])
		acc += a * b
	}

	println(acc)

}
