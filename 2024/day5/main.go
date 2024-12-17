package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/d-forbes/aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(lines), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(lines), "in", time.Since(start))
}

type rule struct {
	before, after int
}

func parseRules(lines []string) ([]rule, int) {
	var rules []rule
	emptyLineIdx := 0

	for i, line := range lines {
		if line == "" {
			emptyLineIdx = i
			break
		}
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		after, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		rules = append(rules, rule{before, after})
	}

	return rules, emptyLineIdx
}

func parseUpdate(line string) []int {
	var nums []int
	for _, s := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(strings.TrimSpace(s))
		nums = append(nums, num)
	}
	return nums
}

func isValidUpdate(update []int, rules []rule) bool {
	// Create a set of pages in the update for quick lookup
	pages := make(map[int]bool)
	for _, page := range update {
		pages[page] = true
	}

	// Check each rule
	for _, r := range rules {
		// Only check rules where both pages are in the update
		if pages[r.before] && pages[r.after] {
			// Find positions of both pages
			beforeIdx := -1
			afterIdx := -1
			for i, page := range update {
				if page == r.before {
					beforeIdx = i
				}
				if page == r.after {
					afterIdx = i
				}
			}
			// Rule is violated if 'after' comes before 'before'
			if beforeIdx > afterIdx {
				return false
			}
		}
	}
	return true
}

func buildGraph(rules []rule, pages []int) map[int][]int {
	// Create adjacency list representation of the graph
	graph := make(map[int][]int)

	// Initialize graph with all pages
	for _, page := range pages {
		if _, exists := graph[page]; !exists {
			graph[page] = []int{}
		}
	}

	// Add edges from rules where both pages are in the update
	pagesSet := make(map[int]bool)
	for _, page := range pages {
		pagesSet[page] = true
	}

	for _, r := range rules {
		if pagesSet[r.before] && pagesSet[r.after] {
			graph[r.before] = append(graph[r.before], r.after)
		}
	}

	return graph
}

func topologicalSort(graph map[int][]int) []int {
	// Count incoming edges for each node
	inDegree := make(map[int]int)
	for node := range graph {
		inDegree[node] = 0
	}
	for _, edges := range graph {
		for _, dest := range edges {
			inDegree[dest]++
		}
	}

	// Initialize queue with nodes that have no incoming edges
	var queue []int
	for node := range graph {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	var result []int
	for len(queue) > 0 {
		// Remove a node from the queue
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// Decrease in-degree for all neighbors
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If we couldn't visit all nodes, there's a cycle
	if len(result) != len(graph) {
		return nil
	}

	return result
}

func orderUpdate(update []int, rules []rule) []int {
	// Build directed graph from rules
	graph := buildGraph(rules, update)

	// Perform topological sort
	ordered := topologicalSort(graph)
	if ordered == nil {
		// If there's a cycle, return the original order
		return update
	}

	return ordered
}

func solvePart1(lines []string) int {
	rules, emptyLineIdx := parseRules(lines)
	sum := 0

	// Process each update
	for _, line := range lines[emptyLineIdx+1:] {
		if line == "" {
			continue
		}
		update := parseUpdate(line)
		if isValidUpdate(update, rules) {
			// For valid updates, add the middle page number
			middleIdx := len(update) / 2
			sum += update[middleIdx]
		}
	}

	return sum
}

func solvePart2(lines []string) int {
	rules, emptyLineIdx := parseRules(lines)
	sum := 0

	// Process each update
	for _, line := range lines[emptyLineIdx+1:] {
		if line == "" {
			continue
		}
		update := parseUpdate(line)
		if !isValidUpdate(update, rules) {
			// For invalid updates, reorder and get middle number
			ordered := orderUpdate(update, rules)
			middleIdx := len(ordered) / 2
			sum += ordered[middleIdx]
		}
	}

	return sum
}
