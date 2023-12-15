package day06

import (
	"fmt"
	"time"

	"github.com/d-forbes/advent-of-code/util"
)

func Day06() {
	start := time.Now()
	fmt.Println("Day 06 Part A Result:", partA(), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Day 06 Part B Result:", partB(), "in", time.Since(start))
}
func partA() int {
	//fmt.Println("Day 05 Part B Result:", partB())
	lines := util.GetInput("day06/day06input.txt")
	test := make([][]int, len(lines))
	distances := make([][]int, 4)

	for i, line := range lines {
		numbers := util.GetNumbersOfLine(line)
		test[i] = append(test[i], numbers...)
	}
	//fmt.Printf("%#v\n", test)

	for i, time := range test[0] {
		wins := 0
		for charge := 0; charge < time+1; charge++ {

			drive := time - charge
			distance := charge * drive
			if distance > test[1][i] {
				wins++
			}
			// fmt.Println(wins)

		}
		distances[i] = append(distances[i], wins)
	}
	return distances[0][0] * distances[1][0] * distances[2][0] * distances[3][0]
}

// charge + drive = time
func partB() int {
	lines := util.GetInput("day06/day06input copy.txt")
	test := make([][]int, len(lines))
	distances := make([][]int, 1)

	for i, line := range lines {
		numbers := util.GetNumbersOfLine(line)
		test[i] = append(test[i], numbers...)
	}
	//fmt.Printf("%#v\n", test)

	for i, time := range test[0] {
		wins := 0
		for charge := 0; charge < time+1; charge++ {

			drive := time - charge
			distance := charge * drive
			if distance > test[1][i] {
				wins++
			}
			//fmt.Println(wins)

		}
		distances[i] = append(distances[i], wins)
	}
	return distances[0][0]
}
