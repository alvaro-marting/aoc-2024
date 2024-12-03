package main

import (
	"regexp"
	"strconv"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

var exp = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

func main() {
	s := pkg.MustReadFile("./inputs/03.txt")

	matches := exp.FindAllStringSubmatch(s, -1)
	acc := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		acc += a * b
	}

	println(acc)

}
