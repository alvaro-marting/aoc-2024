package main

import (
	"fmt"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

type coord struct {
	x, y int
}

func (c coord) nextCoord(c1 coord) coord {
	return coord{
		x: c.x + c1.x,
		y: c.y + c1.y,
	}
}

type xmasSearch struct {
	currentRune rune
	coord       coord
}

func (x *xmasSearch) nextCoord(addx, addy int) coord {
	return coord{
		x: x.coord.x + addx,
		y: x.coord.y + addy,
	}
}

func (x *xmasSearch) nextRune() (rune, bool) {
	r, ok := nextRune[x.currentRune]
	return r, ok
}
func (x *xmasSearch) restart() {
	x.currentRune = 'X'
}

var diagonals = []coord{
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

var directions = []coord{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

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
	xpos := []coord{}
	graph := make([][]rune, 0, len(s))
	for _, line := range s {
		if line == "" {
			continue
		}
		graph = append(graph, []rune(line))
		for i, char := range line {
			if char == 'X' {
				xpos = append(xpos, coord{i, len(graph) - 1})
			}
		}
	}

	xmasCount := 0

	for _, coord := range xpos {
		for _, dir := range directions {
			currentRune := 'X'
			newCoord := coord.nextCoord(dir)
			// While this direction is valid
			for {
				nextRune, ok := nextRune[currentRune]
				if !ok {
					xmasCount++
					currentRune = 'X'
					break
				}
				if newCoord.x < 0 || newCoord.y < 0 || newCoord.y >= len(graph) || newCoord.x >= len(graph[newCoord.y]) {
					break
				}
				if graph[newCoord.y][newCoord.x] != nextRune {
					currentRune = 'X'
					break
				}
				currentRune = nextRune
				newCoord = newCoord.nextCoord(dir)
			}

		}
	}

	return xmasCount
}

func task2(s []string) int {
	apos := []coord{}
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
				apos = append(apos, coord{i, len(graph) - 1})
			}
		}
	}

	xmasCount := 0

	for _, coord := range apos {

		newCoord1 := coord.nextCoord(diagonals[0])
		newCoord2 := coord.nextCoord(diagonals[1])
		newCoord3 := coord.nextCoord(diagonals[2])
		newCoord4 := coord.nextCoord(diagonals[3])

		if !((graph[newCoord1.y][newCoord1.x] == 'M' &&
			graph[newCoord2.y][newCoord2.x] == 'S') ||
			(graph[newCoord1.y][newCoord1.x] == 'S' &&
				graph[newCoord2.y][newCoord2.x] == 'M')) {
			continue
		}
		if !((graph[newCoord3.y][newCoord3.x] == 'M' &&
			graph[newCoord4.y][newCoord4.x] == 'S') ||
			(graph[newCoord3.y][newCoord3.x] == 'S' &&
				graph[newCoord4.y][newCoord4.x] == 'M')) {
			continue
		}
		xmasCount++
	}

	return xmasCount
}
