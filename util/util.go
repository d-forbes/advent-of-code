package util

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetInput(input string) []string {
	file, err := os.ReadFile(input)
	Check(err)

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	return lines
}

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

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
