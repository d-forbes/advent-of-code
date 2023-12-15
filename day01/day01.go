package day01

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Day01() {
	file, err := os.ReadFile("day01/day01Input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	//var p1 = day01a(lines)
	start := time.Now()
	fmt.Println("Day 01 Part A Result:", day01a(lines), "in", time.Since(start))

	//var p2 = day01b(lines)
	start1 := time.Now()
	fmt.Println("Day 01 Part B Result:", day01b(lines), "in", time.Since(start1))

}

func day01a(lines []string) int32 {
	sum := int32(0)

	for _, line := range lines {
		digits := []int32{}
		for _, r := range line {
			if r >= '0' && r <= '9' {
				digits = append(digits, r-'0')
			}
		}
		sum += 10*digits[0] + digits[len(digits)-1]
	}
	return sum

}

func day01b(lines []string) int32 {
	sum := int32(0)

	for _, line := range lines {
		digits := []int32{}
		for i, r := range line {
			if r >= '0' && r <= '9' {
				digits = append(digits, r-'0')
			}
			for j, n := range []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
				if strings.HasPrefix(line[i:], n) {
					digits = append(digits, int32(j))
				}
			}
		}
		sum += 10*digits[0] + digits[len(digits)-1]
	}
	return sum

}
