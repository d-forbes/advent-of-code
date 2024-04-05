package day08

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/d-forbes/advent-of-code/util"
)

type choice struct {
	left  string
	right string
}

func importFile(input string) []string {
	file, err := os.ReadFile(input)
	util.Check(err)

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	return lines
}
func Day08() {
	lines := importFile("day08/day08Input.txt")
	start := time.Now()
	fmt.Println("Day 08 Part A Result:", partA(lines), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Day 08 Part B Result:", partB(lines), "in", time.Since(start))
}

func partA(lines []string) int {

	commands := lines[0]
	var split = func(r rune) bool {
		return r == '=' || r == '(' || r == ')' || r == ',' || r == ' '
	}

	var network = make(map[string]choice)
	for _, line := range lines[2:] {
		var fields = strings.FieldsFunc(line, split)
		network[fields[0]] = choice{fields[1], fields[2]}
	}

	current := "AAA"
	var res int
	for current != "ZZZ" {
		var index = res % len(commands)
		var command = commands[index]
		var choice = network[current]
		if command == 'L' {
			current = choice.left
		} else {
			current = choice.right
		}
		res++
	}
	return res

}

func partB(lines []string) int {
    commands := lines[0]
    var split = func(r rune) bool {
        return r == '=' || r == '(' || r == ')' || r == ',' || r == ' '
    }

    var network = make(map[string]choice)
    for _, line := range lines[2:] {
        var fields = strings.FieldsFunc(line, split)
        network[fields[0]] = choice{fields[1], fields[2]}
    }

    var startNodes []string
    for node := range network {
        if strings.HasSuffix(node, "A") {
            startNodes = append(startNodes, node)
        }
    }

    var res int
    for _, start := range startNodes {
        current := start
        for !strings.HasSuffix(current, "Z") {
            var index = res % len(commands)
            var command = commands[index]
            var choice = network[current]
            if command == 'L' {
                current = choice.left
            } else {
                current = choice.right
            }
            res++
        }
    }

    return res
}