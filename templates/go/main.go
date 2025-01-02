package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d-forbes/aoc/utils"
)

func main() {
	input := parseInput("input.txt")

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(input), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(input), "in", time.Since(start))
}

func parseInput(file string) []string {
	fi, err := utils.ReadLines(file)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	return fi
}

func solvePart1(lines []string) int {
	return 0 // Implement part 1 solution
}

func solvePart2(lines []string) int {
	return 0 // Implement part 2 solution
}
