package main

import (
	"os"
	"testing"

	"github.com/d-forbes/aoc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	example = "example.txt"
	input   = "input.txt"
)

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

func TestP1Example(t *testing.T) {
	_, err := utils.ReadFile(example)
	require.NoError(t, err)
	lines, err := utils.ReadLines(example)
	require.NoError(t, err)
	cards := new(Storage)
	assert.Equal(t, 13, solvePart1(lines, cards))
}

// func TestP2Example(t *testing.T) {
// 	_, err := utils.ReadFile(example)
// 	require.NoError(t, err)
// 	lines, err := utils.ReadLines(example)
// 	require.NoError(t, err)
// 	assert.Equal(t, 30, solvePart2(cards, len(lines)))
// }