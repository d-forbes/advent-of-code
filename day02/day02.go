package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day02() {

	// var p2 = part2(lines)
	// fmt.Println("Day 01 Part 2 Result: ", p2)
	data, _ := os.Open("day02/day02Input.txt")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	sum, powers := 0, 0
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ": ")
		mins := map[string]int{}
		for _, show := range strings.Split(game[1], ";") {
			for _, c := range strings.Split(show, ",") {
				cubes := strings.Split(strings.TrimSpace(c), " ")
				n, _ := strconv.ParseInt(cubes[0], 10, 32)
				mins[cubes[1]] = maxi(mins[cubes[1]], int(n))
			}
		}
		if mins["red"] <= 12 && mins["green"] <= 13 && mins["blue"] <= 14 {
			n, _ := strconv.ParseInt(strings.Split(game[0], " ")[1], 10, 32)
			sum += int(n)
		}
		powers += mins["red"] * mins["green"] * mins["blue"]
	}
	start := time.Now()
	fmt.Println("Day 02 Part A Result:", sum, "in", time.Since(start))
	start1 := time.Now()
	fmt.Println("Day 02 Part B Result:", powers, "in", time.Since(start1))

}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}
