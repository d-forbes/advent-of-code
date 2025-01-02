package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/d-forbes/aoc/utils"
)

func main() {
	input := loadInput("input.txt")

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(input), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(input), "in", time.Since(start))
}

func loadInput(file string) []int {
	fi, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var line string
	scanner := bufio.NewScanner(bytes.NewReader(fi))
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	raw := [][]int{}
	for i, id := 0, 0; i < len(line); i, id = i+2, id+1 {
		raw = append(raw, []int{id, utils.MustConv(string(line[i]))})
		var free int
		if i < len(line)-1 {
			free = utils.MustConv(string(line[i+1]))
		}
		raw = append(raw, []int{-1, free})
	}

	expandedRaw := []int{}
	for _, vals := range raw {
		for i := 0; i < vals[1]; i++ {
			expandedRaw = append(expandedRaw, vals[0])
		}
	}

	return expandedRaw
}

func solvePart1(input []int) int {
	disk := newDisk(input)
	disk.reorganise()

	var checksum int
	for i, n := range disk.data {
		if n == empty {
			break
		}
		checksum += i * n
	}

	return checksum
}

func solvePart2(input []int) int {
	return 0 // Implement part 2 solution
}

const empty = -1

type disk struct {
	data []int
}

func newDisk(raw []int) *disk {
	return &disk{
		data: raw,
	}
}

func (d *disk) reorganise() {
	for i := len(d.data) - 1; i >= 0; i-- {
		indx := slices.Index(d.data, empty)
		if indx >= i {
			break
		}

		d.data[i], d.data[indx] = d.data[indx], d.data[i]
	}
}
