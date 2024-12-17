package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d-forbes/aoc/utils"
)

func main() {
	lines, err := utils.ReadLinestoInts("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	start := time.Now()
	fmt.Println("Part 1:", solvePart1(lines), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(lines), "in", time.Since(start))
}

func solvePart1(lines [][]int) int {
	var safeReports int
	for _, line := range lines {
		if checkReport(line) {
			safeReports++
		}
	}
	return safeReports
}

func solvePart2(lines [][]int) int {
	var safeReportsCompensated int
	for _, line := range lines {
		if checkReportCompensated(line) {
			safeReportsCompensated++
		}
	}
	return safeReportsCompensated
}

func checkReportCompensated(report []int) bool {
	if checkReport(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		reportCopy := make([]int, len(report))
		_ = copy(reportCopy, report)
		compensatedReport := append(reportCopy[:i], reportCopy[i+1:]...)

		if checkReport(compensatedReport) {
			return true
		}
	}
	return false
}

func checkReport(report []int) bool {
	var increased int
	var decreased int
	for i := 0; i < (len(report) - 1); i++ {
		ss := report[i : i+2]
		distance := ss[0] - ss[1]
		if distance == 0 {
			return false
		}
		if distance < 0 {
			decreased++
			distance = distance * -1
		} else {
			increased++
		}
		if distance > 3 {
			return false
		}
	}
	if (increased != 0) && (decreased != 0) {
		return false
	}
	return true
}
