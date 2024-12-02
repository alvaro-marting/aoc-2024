package main

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

func main() {
	s := pkg.MustReadFile(".inputs/01.txt")

	lines := strings.Split(s, "\n")
	loc1, loc2 := make([]int, len(lines)), make([]int, len(lines))
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
		loc1 = append(loc1, int1)
		loc2 = append(loc2, int2)
	}

	slices.Sort(loc1)
	slices.Sort(loc2)

	sum := float64(0)
	for i := range loc1 {
		sum += math.Abs(float64(loc1[i] - loc2[i]))
	}

	println(int(sum))

}

func splitLines(s string) []string {
	strings.Split(s, "\n")
	return []string{}
}
