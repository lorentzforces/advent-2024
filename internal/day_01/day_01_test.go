package day_01

import (
	"testing"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/stretchr/testify/assert"
)

var testInput string =
`3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 11, result)
}

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 31, result)
}
