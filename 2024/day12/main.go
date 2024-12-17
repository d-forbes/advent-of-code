package main

import (
	"fmt"
	"log"
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

func solvePart1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				area, perimeter := calculateRegion(lines, visited, i, j, rows, cols)
				total += area * perimeter
			}
		}
	}
	return total
}

func calculateRegion(grid []string, visited [][]bool, row, col, rows, cols int) (int, int) {
	if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] {
		return 0, 0
	}

	char := grid[row][col]
	if visited[row][col] {
		return 0, 0
	}

	area := 1
	perimeter := 0
	visited[row][col] = true

	// Check all four sides for perimeter calculation
	if row == 0 || grid[row-1][col] != char {
		perimeter++
	}
	if row == rows-1 || grid[row+1][col] != char {
		perimeter++
	}
	if col == 0 || grid[row][col-1] != char {
		perimeter++
	}
	if col == cols-1 || grid[row][col+1] != char {
		perimeter++
	}

	// Recursively explore adjacent cells
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
			!visited[newRow][newCol] && grid[newRow][newCol] == char {
			a, p := calculateRegion(grid, visited, newRow, newCol, rows, cols)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func solvePart2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				area, sides := calculateRegionSides(lines, visited, i, j, rows, cols)
				total += area * sides
			}
		}
	}
	return total
}

func calculateRegionSides(grid []string, visited [][]bool, startRow, startCol, rows, cols int) (int, int) {
	if startRow < 0 || startRow >= rows || startCol < 0 || startCol >= cols || visited[startRow][startCol] {
		return 0, 0
	}

	char := grid[startRow][startCol]
	area := 0
	// Track cells in the region
	regionCells := make(map[[2]int]bool)

	// BFS to find all cells in the region
	queue := [][2]int{{startRow, startCol}}
	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] || grid[row][col] != char {
			continue
		}

		if visited[row][col] {
			continue
		}

		visited[row][col] = true
		area++
		regionCells[[2]int{row, col}] = true

		// Add unvisited neighbors of the same character to queue
		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow][newCol] && grid[newRow][newCol] == char {
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}

	// Create a padded grid for easier edge handling
	paddedGrid := make([][]bool, rows+2)
	counted := make([][][]bool, rows+2)
	for i := range paddedGrid {
		paddedGrid[i] = make([]bool, cols+2)
		counted[i] = make([][]bool, cols+2)
		for j := range paddedGrid[i] {
			counted[i][j] = make([]bool, 4)
		}
	}

	// Fill the padded grid with our region
	for cell := range regionCells {
		paddedGrid[cell[0]+1][cell[1]+1] = true
	}

	// Count sides using the same approach as the solution
	sides := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if paddedGrid[i][j] {
				// Check top edge
				if !paddedGrid[i-1][j] {
					counted[i][j][0] = true
					if !counted[i][j-1][0] {
						sides++
					}
				}
				// Check bottom edge
				if !paddedGrid[i+1][j] {
					counted[i][j][1] = true
					if !counted[i][j-1][1] {
						sides++
					}
				}
				// Check left edge
				if !paddedGrid[i][j-1] {
					counted[i][j][2] = true
					if !counted[i-1][j][2] {
						sides++
					}
				}
				// Check right edge
				if !paddedGrid[i][j+1] {
					counted[i][j][3] = true
					if !counted[i-1][j][3] {
						sides++
					}
				}
			}
		}
	}

	return area, sides
}
