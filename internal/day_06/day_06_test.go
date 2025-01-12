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

func TestPartTwoSampleInput(t *testing.T) {
	result, err := PartTwo(testInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 6, result)
}

var testCornerInput =
`....#..
.......
....^..
...#...
....#..
.......
.......
`

func TestPartTwoCornerLoop(t *testing.T) {
	result, err := PartTwo(testCornerInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 1, result)
}

// The loop which is possible in this setup doesn't intersect with the unblocked path, which I did
// not initially account for.
var testDisconnectedLoopInput =
`.....
..#..
....#
.#...
^..#.
`

func TestPartTwoDisconnectedLoop(t *testing.T) {
	result, err := PartTwo(testDisconnectedLoopInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 1, result)
}

// A naive implementation might consider an obstacle outside the grid boundaries (on the top) as
// something that might induce a loop, but this is incorrect.
var testBoundaryConditionInput =
`....#
.....
#....
...#.
.^...
`

func TestPartTwoOutOfBoundsLoop(t *testing.T) {
	result, err := PartTwo(testBoundaryConditionInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 0, result)
}

// Initial implementation didn't consider the new obstacle as part of further collisions beyond the
// very first right-hand turn. In this case, the loop is a square on the left, which will not be
// recognized if the obstacle "disappears" while the ghost is walking.
var testNewObstacleIsPartOfLoopInput =
`.......
.#.....
....O..
......#
.......
#......
...#^#.
`

func TestPartTwoNewObstacleIntegralToLoop(t *testing.T) {
	result, err := PartTwo(testNewObstacleIsPartOfLoopInput)
	assert.NoError(t, err)
	run.BailIfFailed(t)
	assert.Equal(t, 1, result)
}
