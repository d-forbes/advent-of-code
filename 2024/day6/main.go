package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	grid, err := loadGrid("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	start := time.Now()
	fmt.Println("Part 1:", solvePart1(newGrid(grid)), "in", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", solvePart2(newGrid(grid)), "in", time.Since(start))
}

func solvePart1(grid *Grid) int {
	startPos, ok := grid.Scan(Start)
	if !ok {
		return -1
	}
	g := Guard{Direction: Up, Position: startPos}
	grid.MarkTravel(g)
	for {
		for grid.Get(g.NextStep()) != Block {
			if grid.Get(g.NextStep()) == OOB {
				return grid.traveledSpots
			}
			g.TakeStep()
			grid.MarkTravel(g)
		}
		g.Direction = g.Direction.Next()
	}
}

func solvePart2(grid *Grid) int {
	startPos, ok := grid.Scan(Start)
	if !ok {
		return -1
	}

	height := len(grid.grid)
	width := len(grid.grid[0])
	var loopCount int
	blocks := make(map[XY]bool)

	// Process each position
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := XY{x, y}
			if blocks[pos] {
				continue
			}

			// Create a copy of the grid for this position
			localGrid := newGrid(grid.raw)
			localGrid.Set(pos, Block)

			// Initialize guard at start position
			g := Guard{Direction: Up, Position: startPos}

			// Check for loop
			if hasLoop(localGrid, g) {
				loopCount++
				blocks[pos] = true
			}
		}
	}

	return loopCount
}

func hasLoop(grid *Grid, g Guard) bool {
	visited := make(map[Guard]bool)
	for {
		step := g.NextStep()
		for grid.Get(step) != Block {
			if grid.Get(step) == OOB {
				return false
			}
			if visited[g] {
				return true
			}
			visited[g] = true
			g.TakeStep()
			step = g.NextStep()
		}
		g.Direction = g.Direction.Next()
	}
}

func loadGrid(filePath string) ([][]rune, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid, nil
}

const (
	OOB       = '\x10'
	Traversed = 'X'
	Block     = '#'
	Start     = '^'
)

type Grid struct {
	raw           [][]rune // this copy for resets is a little
	grid          [][]rune
	traveledSpots int
}

func newGrid(raw [][]rune) *Grid {
	return &Grid{
		raw:  raw, // for resets
		grid: newGridFromRaw(raw),
	}
}

func newGridFromRaw(raw [][]rune) [][]rune {
	grid := make([][]rune, len(raw))
	for ix, g := range raw {
		newG := make([]rune, len(g))
		copy(newG, g)
		grid[ix] = newG
	}
	return grid
}

func (g *Grid) Get(pos XY) rune {
	if pos.X >= len(g.grid[0]) || pos.Y >= len(g.grid) || pos.X < 0 || pos.Y < 0 {
		return OOB
	}
	return g.grid[pos.Y][pos.X]
}

func (g *Grid) Scan(r rune) (XY, bool) {
	for y, line := range g.grid {
		for x := range line {
			p := XY{x, y}
			if g.Get(p) == r {
				return p, true
			}
		}
	}
	return XY{}, false
}

func (g *Grid) MarkTravel(guard Guard) {
	if r := g.Get(guard.Position); r != Traversed {
		g.Set(guard.Position, Traversed)
		g.traveledSpots++
	}
}

func (g *Grid) Set(pos XY, r rune) {
	g.grid[pos.Y][pos.X] = r
}

type XY struct {
	X int
	Y int
}

func (p XY) Add(that XY) XY {
	return XY{X: p.X + that.X, Y: p.Y + that.Y}
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) ToStep() XY {
	var pos XY
	switch d {
	case Up:
		pos = XY{0, -1}
	case Right:
		pos = XY{1, 0}
	case Down:
		pos = XY{0, 1}
	case Left:
		pos = XY{-1, 0}
	}
	return pos
}

func (d Direction) Next() Direction {
	return (d + 1) % 4
}

type Guard struct {
	Direction Direction
	Position  XY
}

func (g *Guard) NextStep() XY {
	return g.Position.Add(g.Direction.ToStep())
}

func (g *Guard) TakeStep() {
	g.Position = g.NextStep()
}
