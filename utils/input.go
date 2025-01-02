package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the content of a file and trims whitespace.
func ReadFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return strings.TrimSpace(string(file)), nil
}

// ReadLines reads a file and splits it into lines.
func ReadLines(path string) ([]string, error) {
	content, err := ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines from %s: %w", path, err)
	}
	return strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n"), nil
}

// ReadLinestoInts reads a file and splits it into slices of slices of ints.
func ReadLinestoInts(path string) ([][]int, error) {
	content, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	var lines [][]int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var ints []int
		for i, r := range line {
			n, err := strconv.Atoi(r)
			if err != nil {
				return nil, fmt.Errorf("failed to parse integer at line %d, position %d: %w", len(lines)+1, i+1, err)
			}
			ints = append(ints, n)
		}
		lines = append(lines, ints)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file %s: %w", path, err)
	}

	return lines, nil
}

func MustConv(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
