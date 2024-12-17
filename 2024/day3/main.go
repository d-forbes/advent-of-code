package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/d-forbes/aoc/utils"
)

func main() {
	lines, err := utils.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(lines), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(lines), "in", time.Since(start))
}

func solvePart1(lines string) int {
	rex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := rex.FindAllStringSubmatch(lines, -1)

	sum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}
	return sum
}

func solvePart2(lines string) int {
	clean := regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`)
	cleanData := clean.ReplaceAllString(lines, "")

	rex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := rex.FindAllStringSubmatch(cleanData, -1)

	sum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}
	return sum
}
