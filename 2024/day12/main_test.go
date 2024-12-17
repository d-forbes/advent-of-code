package main

import (
	"os"
	"testing"

	"github.com/d-forbes/aoc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	example  = "example.txt"
	example2 = "example2.txt"
	example3 = "example3.txt"
	example4 = "example4.txt"
	example5 = "example5.txt"
	input    = "input.txt"
)

func TestP1Example(t *testing.T) {
	lines, err := utils.ReadLines(example)
	require.NoError(t, err)
	assert.Equal(t, 140, solvePart1(lines))
}

func TestP1Example2(t *testing.T) {
	lines, err := utils.ReadLines(example2)
	require.NoError(t, err)
	assert.Equal(t, 772, solvePart1(lines))
}

func TestP1Example3(t *testing.T) {
	lines, err := utils.ReadLines(example3)
	require.NoError(t, err)
	assert.Equal(t, 1930, solvePart1(lines))
}

func TestP2Example(t *testing.T) {
	lines, err := utils.ReadLines(example)
	require.NoError(t, err)
	assert.Equal(t, 80, solvePart2(lines))
}

func TestP2Example2(t *testing.T) {
	lines, err := utils.ReadLines(example2)
	require.NoError(t, err)
	assert.Equal(t, 436, solvePart2(lines))
}

func TestP2Example3(t *testing.T) {
	lines, err := utils.ReadLines(example4)
	require.NoError(t, err)
	assert.Equal(t, 236, solvePart2(lines))
}

func TestP2Example4(t *testing.T) {
	lines, err := utils.ReadLines(example5)
	require.NoError(t, err)
	assert.Equal(t, 368, solvePart2(lines))
}

func TestExampleFileExists(t *testing.T) {
	_, err := os.Open(example)
	require.NoError(t, err)
}

func TestExampleFileNotEmpty(t *testing.T) {
	f, _ := os.Stat(example)
	assert.Greater(t, f.Size(), int64(0))
}

func TestInputFileExists(t *testing.T) {
	_, err := os.Open(input)
	require.NoError(t, err)
}

func TestInputFileNotEmpty(t *testing.T) {
	f, _ := os.Stat(input)
	assert.Greater(t, f.Size(), int64(0))
}
