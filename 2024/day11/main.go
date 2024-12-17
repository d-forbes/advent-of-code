package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	ss := strings.Split(scanner.Text(), " ")
	stones := make([]int, len(ss))
	for ix := range ss {
		i, _ := strconv.Atoi(ss[ix])
		stones[ix] = i
	}
	fmt.Println(stones)

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(stones, 25), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(stones, 75), "in", time.Since(start))
}

func solvePart1(stones []int, nBlinks int) int {
	var total int
	for _, stone := range stones {
		total += CountStonesAfterBlinks(stone, nBlinks)
	}
	return total
}

func solvePart2(stones []int, nBlinks int) int {
	var total int
	for _, stone := range stones {
		total += CountStonesAfterBlinksCached(stone, nBlinks)
	}
	return total
}

func CountStonesAfterBlinks(stone int, nBlinks int) int {
	if nBlinks == 0 {
		return 1
	}
	if stone == 0 {
		return CountStonesAfterBlinks(1, nBlinks-1)
	} else if ss := strconv.Itoa(stone); len(ss)%2 == 0 { // even
		ns1, _ := strconv.Atoi(ss[:len(ss)/2])
		ns2, _ := strconv.Atoi(ss[len(ss)/2:])
		return CountStonesAfterBlinks(ns1, nBlinks-1) + CountStonesAfterBlinks(ns2, nBlinks-1)
	} else { // odd
		return CountStonesAfterBlinks(stone*2024, nBlinks-1)
	}
}

type StoneBlinkResult struct {
	Value     int
	NumBlinks int
}

var blinkCacher = make(map[StoneBlinkResult]int)

func CountStonesAfterBlinksCached(stone int, nBlinks int) int {
	if c, ok := blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: nBlinks}]; ok {
		return c
	}
	if nBlinks == 0 {
		return 1
	}
	if stone == 0 {
		return CountStonesAfterBlinksCached(1, nBlinks-1)
	} else if ss := strconv.Itoa(stone); len(ss)%2 == 0 { // even
		ns1, _ := strconv.Atoi(ss[:len(ss)/2])
		ns2, _ := strconv.Atoi(ss[len(ss)/2:])
		created := CountStonesAfterBlinksCached(ns1, nBlinks-1) + CountStonesAfterBlinksCached(ns2, nBlinks-1)
		blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: nBlinks}] = created
		return created
	} else { // odd
		created := CountStonesAfterBlinksCached(stone*2024, nBlinks-1)
		blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: nBlinks}] = created
		return created
	}
}
