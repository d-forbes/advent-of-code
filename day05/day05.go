package day05

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/d-forbes/advent-of-code/util"
)

func Day05() {
	start := time.Now()
	fmt.Println("Day 05 Part A Result:", partA(), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Day 05 Part B Result:", partB(), "in", time.Since(start))
}

type rangeSeed struct {
	start int
	end   int
}

type seq struct {
	start int
	end   int
	diff  int
}

type set = []seq

func scanSeeds(line string) []int {
	return util.GetNumbersOfLine(line)
}

func scanSeedsv2(line string) []rangeSeed {
	numbers := util.GetNumbersOfLine(line)
	seeds := make([]rangeSeed, 0)

	i := 0
	for i < len(numbers) {

		seeds = append(seeds, rangeSeed{
			start: numbers[i],
			end:   numbers[i] + numbers[i+1] - 1,
		})

		i = i + 2
	}

	return seeds
}

func scanSeq(line string) seq {
	numbers := util.GetNumbersOfLine(line)

	source := numbers[1]
	dest := numbers[0]
	size := numbers[2]

	return seq{
		start: source,
		end:   source + size - 1,
		diff:  dest - source,
	}
}

func getSeed(result int, sets []set) int {
	curr := result

	i := len(sets) - 1

	for i >= 0 {
		var seqSelected seq
		found := false

		j := len(sets[i]) - 1

		for j >= 0 {
			seq := sets[i][j]

			if curr >= seq.start+seq.diff && curr <= seq.end+seq.diff {
				seqSelected = seq
				found = true
			}

			j--
		}

		if found {
			curr = curr - seqSelected.diff
		}

		i--
	}

	return curr
}

func getLocation(seed int, sets []set) int {
	curr := seed

	for _, set := range sets {
		var seqSelected seq
		found := false

		for _, seq := range set {
			if curr >= seq.start && curr <= seq.end {
				seqSelected = seq
				found = true
			}
		}

		if found {
			curr = curr + seqSelected.diff
		}
	}

	return curr
}

func getSets(lines []string) []set {
	sets := make([]set, 0)

	for i, line := range lines {
		if i == 0 || len(line) == 0 {
			continue
		}

		if strings.Contains(line, "map:") {
			sets = append(sets, make([]seq, 0))
			continue
		}

		lastSetIdx := len(sets) - 1

		sets[lastSetIdx] = append(sets[lastSetIdx], scanSeq(line))
	}

	return sets
}

func getResult(seeds []int, sets []set) int {
	solution := math.Inf(1)

	seedsResult := make([]int, len(seeds))

	for i, seed := range seeds {
		seedsResult[i] = getLocation(seed, sets)
	}

	for _, res := range seedsResult {
		if res < int(solution) {
			solution = float64(res)
		}
	}

	return int(solution)
}

func partA() int {
	lines := util.GetInput("day05/day05Input.txt")

	seeds := scanSeeds(lines[0])
	sets := getSets(lines)

	return getResult(seeds, sets)
}

func partB() int {
    lines := util.GetInput("day05/day05Input.txt")

    seeds := scanSeedsv2(lines[0])
    sets := getSets(lines)

    // Sort the seeds
    sort.Slice(seeds, func(i, j int) bool {
        return seeds[i].start < seeds[j].start
    })

    low, high := 0, math.MaxInt32

    for low <= high {
        mid := low + (high-low)/2
        seedSolution := getSeed(mid, sets)

        found := false
        for _, seed := range seeds {
            if seedSolution >= seed.start && seedSolution <= seed.end {
                found = true
                break
            }
        }

        if found {
            return mid
        } else if seedSolution < seeds[0].start {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1 // return an error value if no solution is found
}
