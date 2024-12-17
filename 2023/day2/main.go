package main

import (
	"fmt"
	"log"
	"strconv"
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
	sum, powers := 0, 0
	for _, line := range lines {
		game := strings.Split(line, ": ")
		mins := map[string]int{}
		for _, show := range strings.Split(game[1], ";") {
			for _, c := range strings.Split(show, ",") {
				cubes := strings.Split(strings.TrimSpace(c), " ")
				n, _ := strconv.ParseInt(cubes[0], 10, 32)
				mins[cubes[1]] = utils.Maxi(mins[cubes[1]], int(n))
			}
		}
		if mins["red"] <= 12 && mins["green"] <= 13 && mins["blue"] <= 14 {
			n, _ := strconv.ParseInt(strings.Split(game[0], " ")[1], 10, 32)
			sum += int(n)
		}
		powers += mins["red"] * mins["green"] * mins["blue"]
	}
	return sum
}

func solvePart2(lines []string) int {
	sum, powers := 0, 0
	for _, line := range lines {
		game := strings.Split(line, ": ")
		mins := map[string]int{}
		for _, show := range strings.Split(game[1], ";") {
			for _, c := range strings.Split(show, ",") {
				cubes := strings.Split(strings.TrimSpace(c), " ")
				n, _ := strconv.ParseInt(cubes[0], 10, 32)
				mins[cubes[1]] = utils.Maxi(mins[cubes[1]], int(n))
			}
		}
		if mins["red"] <= 12 && mins["green"] <= 13 && mins["blue"] <= 14 {
			n, _ := strconv.ParseInt(strings.Split(game[0], " ")[1], 10, 32)
			sum += int(n)
		}
		powers += mins["red"] * mins["green"] * mins["blue"]
	}
	return powers
}
