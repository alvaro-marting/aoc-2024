package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

func main() {
	s := pkg.MustReadFileLines("inputs/05.txt")

	rules, updates := readFile(s)

	fmt.Println(task1(rules, updates))
	fmt.Println(task2(rules, updates))
}

func readFile(s []string) (map[string][]string, [][]string) {
	rules := make(map[string][]string, 0)
	updates := make([][]string, 0)

	firstStop := false
	for _, v := range s {
		if strings.Trim(v, "\r\n") == "" {
			if !firstStop {
				firstStop = true
			}
			continue
		}
		if !firstStop {
			//rules
			rule := strings.Split(v, "|")
			rules[rule[0]] = append(rules[rule[0]], strings.Trim(rule[1], "\r\n"))
			continue
		}

		// updates
		update := strings.Split(v, ",")
		updates = append(updates, update)
	}
	return rules, updates
}

func task1(rules map[string][]string, updates [][]string) int {
	validUpdates := make([]int, 0)

	for x, line := range updates {
		valid := true
		// For v, there must not be a j position that needs to be printed after v,
		// That means that rule of v must not match with any j
		for i, v := range line {
			rule := rules[strings.Trim(v, "\r\n")]
			for j := i - 1; j >= 0; j-- {
				if slices.Contains(rule, line[j]) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			validUpdates = append(validUpdates, x)
		}
	}

	acc := 0

	for _, v := range validUpdates {
		line := updates[v]
		val, err := strconv.Atoi(line[int(math.Floor(float64(len(line))/float64(2)))])
		if err != nil {
			log.Fatal(err)
		}
		acc += val
	}
	return acc
}

func task2(rules map[string][]string, updates [][]string) int {
	invalidUpdates := make([]int, 0)

	for x, line := range updates {
		valid := true
		// For v, there must not be a j position that needs to be printed after v,
		// That means that rule of v must not match with any j
		for i, v := range line {
			rule := rules[strings.Trim(v, "\r\n")]
			for j := i - 1; j >= 0; j-- {
				if slices.Contains(rule, line[j]) {
					valid = false
					invalidUpdates = append(invalidUpdates, x)
					break
				}
			}
			if !valid {
				break
			}
		}
	}

	acc := 0
	for _, v := range invalidUpdates {
		line := updates[v]

		newLine := slices.Clone(line)
		// Rearrange line
		slices.SortFunc(newLine, func(s1, s2 string) int {
			s1Rule := rules[strings.Trim(s1, "\r\n")]
			s2Rule := rules[strings.Trim(s2, "\r\n")]
			if slices.Contains(s1Rule, strings.Trim(s2, "\r\n")) {
				return -1
			}
			if slices.Contains(s2Rule, strings.Trim(s1, "\r\n")) {
				return 1
			}
			return 0
		})

		val, err := strconv.Atoi(strings.Trim(newLine[int(math.Floor(float64(len(newLine))/float64(2)))], "\r\n"))
		if err != nil {
			log.Fatal(err)
		}
		acc += val
	}
	return acc
}
