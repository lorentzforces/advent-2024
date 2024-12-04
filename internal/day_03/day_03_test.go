package day_03

import (
	"testing"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/stretchr/testify/assert"
)

var testPartOneInput string = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testPartOneInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 161, result)
}

var testPartTwoInput string =
	"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testPartTwoInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 48, result)
}
