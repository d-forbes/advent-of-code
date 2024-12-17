package main

import (
	"fmt"
	"log"
	"strings"
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
	sum := 0

	for _, line := range lines {
		digits := []int{}
		for _, r := range line {
			//fmt.Println(i, r, string(r))
			if r >= '0' && r <= '9' {
				digits = append(digits, int(r-'0'))
				//fmt.Println(digits)
			}
		}
		sum += 10*digits[0] + digits[len(digits)-1]
	}
	return sum
}

func solvePart2(lines []string) int {
	sum := 0

	for _, line := range lines {
		digits := []int{}
		for i, r := range line {
			//fmt.Println(i, r, string(r))
			if r >= '0' && r <= '9' {
				digits = append(digits, int(r-'0'))
			}
			for j, n := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], n) {
					digits = append(digits, int(j))
				}
			}
		}
		sum += 10*digits[0] + digits[len(digits)-1]
	}
	return sum
}
