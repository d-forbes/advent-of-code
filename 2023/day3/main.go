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
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	gears := map[string][]int{}

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(lines, gears), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(gears), "in", time.Since(start))
}

func solvePart1(lines []string, gears map[string][]int) int {
	re := regexp.MustCompile(`\d+`)
	total := 0
	lineLength := len(lines[0])
	linesLength := len(lines)

	for i, line := range lines {
		numbers := re.FindAllStringIndex(line, -1)
		//fmt.Println(i, line, numbers)

		if len(numbers) == 0 {
			continue
		}

		for _, number := range numbers {
			num := line[number[0]:number[1]]
			adjacent := false
			var symbol string
			symbolP := &symbol
			var symbolPosition int
			symbolPositionP := &symbolPosition
			var lineNum int
			lineNumP := &lineNum

			for _, index := range getRange(number[0], number[1]) {
				if i == 0 {
					if index == 0 {
						under, underRight := string(lines[i+1][0]), string(lines[i+1][1])
						if isAdjacent(under, symbolP, i+1, lineNumP, 0, symbolPositionP) || isAdjacent(underRight, symbolP, i+1, lineNumP, 1, symbolPositionP) {
							adjacent = true
							break
						}
						continue
					}

					if index == lineLength-1 {
						under, underLeft := string(lines[i+1][index]), string(lines[i+1][index-1])
						if isAdjacent(under, symbolP, i+1, lineNumP, index, symbolPositionP) || isAdjacent(underLeft, symbolP, i+1, lineNumP, index-1, symbolPositionP) {
							adjacent = true
							break
						}
						continue
					}

					left, underLeft, under, underRight, right := string(lines[i][index-1]), string(lines[i+1][index-1]), string(lines[i+1][index]), string(lines[i+1][index+1]), string(lines[i][index+1])
					if isAdjacent(left, symbolP, i, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(underLeft, symbolP, i+1, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(under, symbolP, i+1, lineNumP, index, symbolPositionP) ||
						isAdjacent(underRight, symbolP, i+1, lineNumP, index+1, symbolPositionP) ||
						isAdjacent(right, symbolP, i, lineNumP, index+1, symbolPositionP) {
						adjacent = true
						break
					} else {
						continue
					}
				}

				if i == linesLength-1 {
					if index == 0 {
						up, upRight := string(lines[i-1][0]), string(lines[i-1][1])
						if isAdjacent(up, symbolP, i-1, lineNumP, 0, symbolPositionP) ||
							isAdjacent(upRight, symbolP, i-1, lineNumP, 1, symbolPositionP) {
							adjacent = true
							break
						}
						continue
					}

					if index == lineLength-1 {
						up, upLeft := string(lines[i-1][index]), string(lines[i-1][index-1])
						if isAdjacent(up, symbolP, i-1, lineNumP, index, symbolPositionP) ||
							isAdjacent(upLeft, symbolP, i-1, lineNumP, index-1, symbolPositionP) {
							adjacent = true
							break
						}
						continue
					}

					left, upLeft, up, upRight, right := string(lines[i][index-1]), string(lines[i-1][index-1]), string(lines[i-1][index]), string(lines[i-1][index+1]), string(lines[i][index+1])
					if isAdjacent(left, symbolP, i, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(upLeft, symbolP, i-1, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(up, symbolP, i-1, lineNumP, index, symbolPositionP) ||
						isAdjacent(upRight, symbolP, i-1, lineNumP, index+1, symbolPositionP) ||
						isAdjacent(right, symbolP, i, lineNumP, index+1, symbolPositionP) {
						adjacent = true
						break
					} else {
						continue
					}
				}

				if index == 0 {
					up, upRight, right, downRight, down := string(lines[i-1][index]), string(lines[i-1][index+1]), string(lines[i][index+1]), string(lines[i+1][index+1]), string(lines[i+1][index])
					if isAdjacent(up, symbolP, i-1, lineNumP, index, symbolPositionP) ||
						isAdjacent(upRight, symbolP, i-1, lineNumP, index+1, symbolPositionP) ||
						isAdjacent(right, symbolP, i, lineNumP, index+1, symbolPositionP) ||
						isAdjacent(downRight, symbolP, i+1, lineNumP, index+1, symbolPositionP) ||
						isAdjacent(down, symbolP, i+1, lineNumP, index, symbolPositionP) {
						adjacent = true
						break
					}
					continue
				}

				if index == lineLength-1 {
					up, upLeft, left, downLeft, down := string(lines[i-1][index-1]), string(lines[i-1][index-1]), string(lines[i][index-1]), string(lines[i+1][index-1]), string(lines[i+1][index])
					if isAdjacent(up, symbolP, i-1, lineNumP, index, symbolPositionP) ||
						isAdjacent(upLeft, symbolP, i-1, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(left, symbolP, i, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(downLeft, symbolP, i+1, lineNumP, index-1, symbolPositionP) ||
						isAdjacent(down, symbolP, i+1, lineNumP, index, symbolPositionP) {
						adjacent = true
						break
					}
					continue
				}

				up, upLeft, left, downLeft, down, downRight, right, upRight := string(lines[i-1][index]), string(lines[i-1][index-1]), string(lines[i][index-1]), string(lines[i+1][index-1]), string(lines[i+1][index]), string(lines[i+1][index+1]), string(lines[i][index+1]), string(lines[i-1][index+1])
				if isAdjacent(up, symbolP, i-1, lineNumP, index, symbolPositionP) ||
					isAdjacent(upLeft, symbolP, i-1, lineNumP, index-1, symbolPositionP) ||
					isAdjacent(left, symbolP, i, lineNumP, index-1, symbolPositionP) ||
					isAdjacent(downLeft, symbolP, i+1, lineNumP, index-1, symbolPositionP) ||
					isAdjacent(down, symbolP, i+1, lineNumP, index, symbolPositionP) ||
					isAdjacent(upRight, symbolP, i-1, lineNumP, index+1, symbolPositionP) ||
					isAdjacent(right, symbolP, i, lineNumP, index+1, symbolPositionP) ||
					isAdjacent(downRight, symbolP, i+1, lineNumP, index+1, symbolPositionP) {
					adjacent = true
					break
				}
			}

			if adjacent {
				convertedNum := getNumber(num)
				total += convertedNum
				if symbol == "*" {
					key := fmt.Sprintf("%d-%d", lineNum, symbolPosition)
					gears[key] = append(gears[key], convertedNum)
				}
			}
		}
	}

	return total
}

func solvePart2(gears map[string][]int) int {
	total := 0
	for _, v := range gears {
		if len(v) == 2 {
			total += v[0] * v[1]
		}
	}

	return total
}

func isNan(number string) bool {
	_, err := strconv.Atoi(number)

	return err != nil
}

func getNumber(number string) int {
	if v, err := strconv.Atoi(number); err == nil {
		return v
	}

	return 0
}

func getRange(start int, end int) []int {
	numRange := make([]int, end-start)
	for i := range numRange {
		numRange[i] = i + start
	}

	return numRange
}

func isAdjacent(symbol string, symbolP *string, num int, numP *int, symbolPosition int, symbolPositionP *int) bool {
	if isNan(symbol) && symbol != "." {
		*symbolP = symbol
		*numP = num
		*symbolPositionP = symbolPosition
		return true
	}

	return false
}
