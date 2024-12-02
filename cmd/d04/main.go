package main

import (
	"log"
	"math"
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
		lvs := pkg.MapStringSliceToInt(strings.Split(line, " "))
		unsafeties := 0
		increasing := lvs[1] > lvs[0]
		for i := 1; i < len(lvs); i++ {
			if math.Abs(float64(lvs[i-1]-lvs[i])) < 1 || math.Abs(float64(lvs[i-1]-lvs[i])) > 3 || (lvs[i] > lvs[i-1] != increasing) {
				unsafeties++
			}
			if unsafeties > 1 {
				break
			}
			if i == len(lvs)-1 {
				safes++
			}
		}
	}

	log.Println(safes)

}
