package main

import (
	"fmt"
	"log"
	"sort"
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
	var leftList []int
	var rightList []int

	for _, line := range lines {
		elements := strings.Split(line, "   ")
		var num int

		num, _ = strconv.Atoi(elements[0])
		leftList = append(leftList, num)

		num, _ = strconv.Atoi(elements[1])
		rightList = append(rightList, num)
	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	totalDiff := 0

	for i := range leftList {
		diff := utils.AbsDiff(leftList[i], rightList[i])
		totalDiff += diff
	}
	return totalDiff
}

func solvePart2(lines []string) int {
	var leftList []int
	var rightList []int

	for _, line := range lines {
		elements := strings.Split(line, "   ")
		var num int

		num, _ = strconv.Atoi(elements[0])
		leftList = append(leftList, num)

		num, _ = strconv.Atoi(elements[1])
		rightList = append(rightList, num)
	}

	totalScore := 0

	for i := range leftList {
		count := 0
		for j := range rightList {
			if leftList[i] == rightList[j] {
				count++
			}
		}
		score := leftList[i] * count
		totalScore += score
	}

	return totalScore
}
