package main

import (
	"iter"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

func main() {
	lines := pkg.MustReadFileLines("./inputs/02.txt")
	// Each line is a report
	// Safe if only decreasing or increasing
	// Safe if the difference is max 3
	safes := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		lvs := mapStringSliceToInt(strings.Split(line, " "))
		increasing := lvs[1] > lvs[0]
		for i := 1; i < len(lvs); i++ {
			if math.Abs(float64(lvs[i-1]-lvs[i])) < 1 || math.Abs(float64(lvs[i-1]-lvs[i])) > 3 || (lvs[i] > lvs[i-1] != increasing) {
				break
			}
			if i == len(lvs)-1 {
				safes++
			}
		}
	}

	log.Println(safes)

}

func mapStringSliceToInt(s []string) []int {
	return slices.Collect(mapSlice(slices.Values(s), func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return i
	}))
}

func mapSlice[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for a := range seq {
			if !yield(f(a)) {
				return
			}
		}
	}
}
