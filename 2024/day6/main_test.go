package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	example = "example.txt"
	input   = "input.txt"
)

func TestP1Example(t *testing.T) {
	grid, err := loadGrid(example)
	require.NoError(t, err)
	assert.Equal(t, 41, solvePart1(newGrid(grid)))
}

func TestP2Example(t *testing.T) {
	grid, err := loadGrid(example)
	require.NoError(t, err)
	assert.Equal(t, 6, solvePart2(newGrid(grid)))
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
