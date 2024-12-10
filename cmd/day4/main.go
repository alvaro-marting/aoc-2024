package main

import (
	"fmt"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

var nextRune = map[rune]rune{
	'X': 'M',
	'M': 'A',
	'A': 'S',
}

func main() {
	// Find all XMAS words in graph
	// Brute forced solution:
	// Store all chars in a double array

	s := pkg.MustReadFileLines("./inputs/04.txt")

	fmt.Println(task1(s))
	fmt.Println(task2(s))
}

func task1(s []string) int {
	xpos := []pkg.Coord{}
	graph := make([][]rune, 0, len(s))
	for _, line := range s {
		if line == "" {
			continue
		}
		graph = append(graph, []rune(line))
		for i, char := range line {
			if char == 'X' {
				xpos = append(xpos, pkg.Coord{X: i, Y: len(graph) - 1})
			}
		}
	}

	xmasCount := 0

	for _, coord := range xpos {
		for _, dir := range pkg.Directions {
			currentRune := 'X'
			newCoord := coord.NextCoord(dir)
			// While this direction is valid
			for {
				nextRune, ok := nextRune[currentRune]
				if !ok {
					xmasCount++
					currentRune = 'X'
					break
				}
				if newCoord.X < 0 || newCoord.Y < 0 || newCoord.Y >= len(graph) || newCoord.X >= len(graph[newCoord.Y]) {
					break
				}
				if graph[newCoord.Y][newCoord.X] != nextRune {
					currentRune = 'X'
					break
				}
				currentRune = nextRune
				newCoord = newCoord.NextCoord(dir)
			}

		}
	}

	return xmasCount
}

func task2(s []string) int {
	apos := []pkg.Coord{}
	graph := make([][]rune, 0, len(s))
	for j, line := range s {
		if line == "" {
			continue
		}
		graph = append(graph, []rune(line))
		for i, char := range line {
			if i == 0 || j == 0 || i == len(line)-1 || j == len(line)-1 { // Assumption that the graph is square
				continue
			}
			if char == 'A' {
				apos = append(apos, pkg.Coord{X: i, Y: len(graph) - 1})
			}
		}
	}

	xmasCount := 0

	for _, coord := range apos {

		newCoord1 := coord.NextCoord(pkg.Diagonals[0])
		newCoord2 := coord.NextCoord(pkg.Diagonals[1])
		newCoord3 := coord.NextCoord(pkg.Diagonals[2])
		newCoord4 := coord.NextCoord(pkg.Diagonals[3])

		if !((graph[newCoord1.Y][newCoord1.X] == 'M' &&
			graph[newCoord2.Y][newCoord2.X] == 'S') ||
			(graph[newCoord1.Y][newCoord1.X] == 'S' &&
				graph[newCoord2.Y][newCoord2.X] == 'M')) {
			continue
		}
		if !((graph[newCoord3.Y][newCoord3.X] == 'M' &&
			graph[newCoord4.Y][newCoord4.X] == 'S') ||
			(graph[newCoord3.Y][newCoord3.X] == 'S' &&
				graph[newCoord4.Y][newCoord4.X] == 'M')) {
			continue
		}
		xmasCount++
	}

	return xmasCount
}
