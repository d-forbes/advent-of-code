package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// Reverses the order of characters in a string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Counts how many times a substring appears in a string.
func CountOccurrences(s, substr string) int {
	count := 0
	for {
		idx := strings.Index(s, substr)
		if idx == -1 {
			break
		}
		count++
		s = s[idx+len(substr):]
	}
	return count
}

// Extracts all numbers from a string and returns them as a slice of integers.
func GetNumbersOfLine(line string) []int {
	re := regexp.MustCompile(`\d+`)

	f := re.FindAllStringIndex(line, -1)

	numbers := make([]int, 0)

	for _, match := range f {
		valueStr := line[match[0]:match[1]]
		val, _ := strconv.Atoi(valueStr)

		numbers = append(numbers, val)
	}

	return numbers
}
