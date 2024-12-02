package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

func main() {
	s := pkg.MustReadFile("./inputs/01.txt")

	lines := strings.Split(s, "\n")
	loc1, loc2 := make(map[int]int, 0), make(map[int]int, 0)
	for _, v := range lines {
		if v == "" {
			continue
		}
		groups := strings.Split(v, "   ")
		if len(groups) != 2 {
			log.Fatal("invalid split")
		}
		int1, err := strconv.Atoi(strings.Trim(groups[0], " "))
		if err != nil {
			log.Fatal(err)
		}
		int2, err := strconv.Atoi(strings.Trim(groups[1], " "))
		if err != nil {
			log.Fatal(err)
		}
		loc1[int1]++
		loc2[int2]++
	}

	acc := 0
	for k, v := range loc1 {
		acc += k * v * loc2[k]
	}

	log.Println(acc)
}
