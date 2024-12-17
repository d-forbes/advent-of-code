package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d-forbes/aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(lines), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(lines), "in", time.Since(start))
}

func solvePart1(lines []string) int {
	return 0 // Implement part 1 solution
}

func solvePart2(lines []string) int {
	return 0 // Implement part 2 solution
}
