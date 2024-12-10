package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alvaro-marting/aoc-2024/pkg"
)

var grid int

func main() {

	// Map of the area.
	// # are obstacles
	// ^ Is the agent looking up
	// Guard turns right after colliding with a #, so it will stop at the position just before that
	s := pkg.MustReadFileLines("06.txt")

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
	grid = len(s)

	fmt.Println(task1(obstacles, guard))
	fmt.Println(task2(obstacles, guard))
}

func task1(obstacles []pkg.Coord, guard pkg.Coord) int {
	stop := false
	i := 0
	positionsVisited := map[pkg.Coord]bool{guard: true}
	for { // Until the guard exits the area
		dir := pkg.Axii[i]

		for { // Until The guard hits an obstacle or exits the map
			newGuardPos := guard.NextCoord(dir)
			if slices.Contains(obstacles, newGuardPos) {
				break
			}
			if newGuardPos.X < 0 || newGuardPos.Y < 0 || newGuardPos.X >= grid || newGuardPos.Y >= grid {
				stop = true
				break
			}
			guard = newGuardPos
			positionsVisited[guard] = true
		}
		if stop {
			break
		}

		i = (i + 1) % 4
	}

	return len(positionsVisited)
}

// Returns the number of possible obstacles added to cause a loop
func task2(obstacles []pkg.Coord, guard pkg.Coord) int {
	stop := false
	i := 0
	initialGuard := guard

	positionsVisited := map[pkg.Coord]bool{guard: true}
	for { // Until the guard exits the area
		dir := pkg.Axii[i]

		for { // Until The guard hits an obstacle or exits the map
			newGuardPos := guard.NextCoord(dir)
			if slices.Contains(obstacles, newGuardPos) {
				break
			}
			if newGuardPos.X < 0 || newGuardPos.Y < 0 || newGuardPos.X >= grid || newGuardPos.Y >= grid {
				stop = true
				break
			}
			guard = newGuardPos
			positionsVisited[guard] = true
		}
		if stop {
			break
		}

		i = (i + 1) % 4
	}

	posLoops := 0
	for k := range positionsVisited {
		if k != initialGuard && checkLoop(append(obstacles, k), initialGuard) {
			posLoops++
		}
	}

	return posLoops
}

func checkLoop(obstacles []pkg.Coord, guard pkg.Coord) bool {
	stop := false
	i := 0
	stoppedPositions := map[pkg.Coord]map[pkg.Coord]bool{}
	for { // Until the guard exits the area
		dir := pkg.Axii[i]

		for { // Until The guard hits an obstacle or exits the map
			newGuardPos := guard.NextCoord(dir)
			if slices.Contains(obstacles, newGuardPos) {
				dirs, ok := stoppedPositions[guard]
				if !ok {
					stoppedPositions[guard] = map[pkg.Coord]bool{dir: true}
				} else {
					_, ok := dirs[dir]
					if ok {
						return true
					} else {
						dirs[dir] = true
					}
				}

				break
			}
			if newGuardPos.X < 0 || newGuardPos.Y < 0 || newGuardPos.X >= grid || newGuardPos.Y >= grid {
				stop = true
				break
			}
			guard = newGuardPos
		}
		if stop {
			break
		}

		i = (i + 1) % 4
	}
	return false
}

func trim(s string) string {
	return strings.Trim(s, "\r\n")
}
