package main

import (
	"fmt"

	"github.com/d-forbes/advent-of-code/day01"
	"github.com/d-forbes/advent-of-code/day02"
	"github.com/d-forbes/advent-of-code/day03"
	"github.com/d-forbes/advent-of-code/day04"
	"github.com/d-forbes/advent-of-code/day05"
	"github.com/d-forbes/advent-of-code/day06"
	"github.com/d-forbes/advent-of-code/day07"
	"github.com/d-forbes/advent-of-code/day08"
)

func main() {
	// now := time.Now()
	// defer func() {
	// 	fmt.Println(time.Since(now))
	// }()
	fmt.Println("Advent of Code 2023 Solutions")
	day01.Day01()
	day02.Day02()
	day03.Day03()
	day04.Day04()
	day05.Day05()
	day06.Day06()
	day07.Day07()
	day08.Day08()
}
