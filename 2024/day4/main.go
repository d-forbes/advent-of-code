package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(grid), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(grid), "in", time.Since(start))
}

func solvePart1(grid [][]rune) int {
	return FindXmasUni(grid)
}

func solvePart2(grid [][]rune) int {
	return CountXDashMas(grid)
}

var (
	startRune = 'X'
	endRune   = '-'
	targetMap = map[rune]rune{
		startRune: 'M',
		'M':       'A',
		'A':       'S',
		'S':       endRune,
	}
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 1},
		{1, 0},
		{1, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
	}
)

func FindXmasUni(grid [][]rune) int {
	var total int
	for y, rl := range grid {
		for x := range rl {
			// we could early-exit here if not on 'X', but whatever, we'll try everything every time
			for _, dir := range directions {
				if directionHasXmas(grid, x, y, dir) {
					total++
				}
			}
		}
	}
	return total
}

func directionHasXmas(grid [][]rune, x, y int, direction []int) bool {
	nextRune := startRune
	for grid[y][x] == nextRune {
		nextRune = targetMap[nextRune]
		if nextRune == endRune {
			return true
		}
		x += direction[0]
		y += direction[1]
		if !inBounds(x, y, grid) {
			return false // bounds check before for-loop check
		}
	}
	return false
}

var (
	corners = [][]int{
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
)

func CountXDashMas(grid [][]rune) int {
	var total int
	for y, rl := range grid {
		for x := range rl {
			if grid[y][x] != 'A' {
				continue
			}
			xmas := make(map[rune]int)
			for _, c := range corners {
				if ok, r := isValidCorner(grid, x, y, c); ok {
					// this makes sure we have the right number of the right letters
					// isValidCorner checks that they're arranged correctly; this checks they're counted correctly.
					// is this smart? who's to say!!
					xmas[r]++
				}
			}
			if reflect.DeepEqual(xmas, map[rune]int{'M': 2, 'S': 2}) {
				total++
			}
		}
	}
	return total
}

func isValidCorner(grid [][]rune, x, y int, c []int) (bool, rune) {
	newX := x + c[0]
	newY := y + c[1]
	oppositeX := x - c[0]
	oppositeY := y - c[1]
	// not on the grid
	if !inBounds(newX, newY, grid) || !inBounds(oppositeX, oppositeY, grid) {
		return false, '\x10'
	}
	// not a valid letter
	if grid[newY][newX] != 'M' && grid[newY][newX] != 'S' {
		return false, '\x10'
	}
	// MAM + SAS
	if grid[newY][newX] == grid[oppositeY][oppositeX] {
		return false, '\x10'
	}
	return true, grid[newY][newX]
}

func inBounds(x, y int, grid [][]rune) bool {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return false
	}
	return true
}
