package day_06

import (
	"testing"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/stretchr/testify/assert"
)

var testInput =
`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 41, result)
}
