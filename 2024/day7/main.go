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
	eqs := loadInput("input.txt")

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(eqs), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(eqs), "in", time.Since(start))
}

func loadInput(file string) []Equation {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	var eqs []Equation
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ": ")
		var is []int
		for _, n := range strings.Split(res[1], " ") {
			i, _ := strconv.Atoi(n)
			is = append(is, i)
		}
		tv, _ := strconv.Atoi(res[0])
		eqs = append(eqs, Equation{Numbers: is, TestValue: tv})
	}
	return eqs
}

func solvePart1(eqs []Equation) int {
	var total int
	for _, eq := range eqs {
		if IsTrueEquation(eq, []Op{Add, Mult}) {
			total += eq.TestValue
		}
	}
	return total
}

func solvePart2(eqs []Equation) int {
	var total int
	for _, eq := range eqs {
		if IsTrueEquation(eq, []Op{Add, Mult, Concat}) {
			total += eq.TestValue
		}
	}
	return total
}

type Equation struct {
	TestValue int
	Numbers   []int
}

type Op int

const (
	Add Op = iota
	Mult
	Concat
)

func cartesianProduct(elems []int, choose int) func() []int {
	// Putting these outside the closure and slicing means we're only using one backing array.
	// Plus, we only generate one permutation at a time - no massive materialized list. Good job, StackOverflow person.
	perms := make([]int, choose)
	ixes := make([]int, choose)
	return func() []int {
		perms = perms[:len(ixes)]
		// Assign our next permutation to the values specified by our index slice.
		for i, xi := range ixes {
			perms[i] = elems[xi]
		}
		// This loop essentially holds all but one index constant every time. It's easiest to see in action.
		// For elems == [7 8 9], choose == 4:
		// [0 0 0 0] -> selects [7 7 7 7]
		// [0 0 0 1] -> selects [7 7 7 8]
		// [0 0 0 2] -> selects [7 7 7 9]
		// [0 0 1 0] -> selects [7 7 8 7]
		// [0 0 1 1] -> selects [7 7 8 8]
		// [0 0 1 2] -> selects [7 7 8 9]
		// [0 0 2 0] -> selects [7 7 9 7]
		// ...and so on.
		for i := len(ixes) - 1; i >= 0; i-- {
			ixes[i]++
			if ixes[i] < len(elems) {
				break
			}
			ixes[i] = 0
			if i <= 0 {
				ixes = ixes[0:0]
				break
			}
		}
		return perms
	}
}

func IsTrueEquation(eq Equation, opTypes []Op) bool {
	// TODO: this is weird, I would do something else here, like use consistent types throughout.
	opInts := make([]int, len(opTypes))
	for ix := range opTypes {
		opInts[ix] = int(opTypes[ix])
	}
	gen := cartesianProduct(opInts, len(eq.Numbers))
	for {
		opsCartesian := gen()
		if len(opsCartesian) == 0 {
			break
		}
		var res int
		for ix, op := range opsCartesian {
			if res > eq.TestValue {
				break // short-circuit to prune some branches
			}
			if op == 0 {
				res += eq.Numbers[ix]
			} else if op == 1 {
				res *= eq.Numbers[ix]
			} else if op == 2 {
				res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, eq.Numbers[ix]))
			}
		}
		if res == eq.TestValue {
			return true
		}
	}
	return false
}
