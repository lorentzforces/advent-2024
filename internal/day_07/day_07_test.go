package day_07

import (
	"testing"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/stretchr/testify/assert"
)

var testInput =
`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, uint64(3749), result)
}

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, uint64(11387), result)
}
