package day_02

import (
	"testing"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/stretchr/testify/assert"
)

var testInput string =
`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 2, result)
}

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 4, result)
}

func TestPartTwoWithFirstNumberBad(t *testing.T) {
	result, err := PartTwo("5 6 4 2 1")
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 1, result)
}

func TestPartTwoWithLastNumberBad(t *testing.T) {
	result, err := PartTwo("7 6 4 2 3")
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 1, result)
}
