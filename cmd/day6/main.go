package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

func main() {

	// Map of the area.
	// # are obstacles
	// ^ Is the agent looking up
	// Guard turns right after colliding with a #, so it will stop at the position just before that
	s := pkg.MustReadFileLines("inputs/06.txt")

	obstacles := []pkg.Coord{}
	guard := pkg.Coord{}
	for j, line := range s {
		if trim(line) == "" {
			continue
		}
		for i, cha := range trim(line) {
			if cha == '#' {
				obstacles = append(obstacles, pkg.Coord{X: i, Y: j})
				continue
			}
			if cha == '^' {
				guard = pkg.Coord{X: i, Y: j}
				continue
			}
		}
	}

	stop := false
	i := 0
	positionsVisited := map[pkg.Coord]bool{guard: true}
	for { // Until the guard exits the area
		dir := pkg.Axii[i]
		initial := guard

		for { // Until The guard hits an obstacle or exits the map
			newGuardPos := guard.NextCoord(dir)
			if slices.Contains(obstacles, newGuardPos) {
				break
			}
			if newGuardPos.X < 0 || newGuardPos.Y < 0 || newGuardPos.X >= (len(s)) || newGuardPos.Y >= (len(s)) {
				stop = true
				break
			}
			guard = newGuardPos
			positionsVisited[guard] = true
		}
		final := guard
		fmt.Println(initial, " -> ", final)
		if stop {
			break
		}

		i = (i + 1) % 4
	}

	fmt.Println(len(positionsVisited)) // 5094 too low
}

func trim(s string) string {
	return strings.Trim(s, "\r\n")
}
