package day_06

import (
	"fmt"
	"regexp"

	"github.com/lorentzforces/advent-2024/internal/run"
)

func PartOne(input string) (int, error) {
	floorGrid := readGrid(input)
	startingLocation, found := floorGrid.findSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	location := guardLocation{
		currLocation: startingLocation,
		facing: directionFrom(up),
		onMap: true,
	}

	visitedLocations := make(map[vec]struct{}, 0)
	for location.onMap {
		visitedLocations[location.currLocation] = run.Empty
		doMove(floorGrid, &location)
	}

	return len(visitedLocations), nil
}

const impassableTile = '#'

func doMove(floorMap grid, location *guardLocation) {
	aheadLocation := location.currLocation.add(location.facing.unitVec)
	aheadTile := floorMap.charAt(aheadLocation)
	if aheadTile == impassableTile {
		location.facing = location.facing.rightAngleClockwise()
	} else {
		location.currLocation = aheadLocation
		if floorMap.charAt(aheadLocation) == 0 {
			location.onMap = false
		}
	}
}

type guardLocation struct {
	currLocation vec
	facing direction
	onMap bool
}

type directionId uint8
const (
	up directionId = iota
	down
	left
	right
)

type direction struct {
	id directionId
	label string
	unitVec vec
}

var directions = map[directionId]direction{
	up: {
		up,
		"UP",
		vec{0, -1},
	},
	down: {
		down,
		"DOWN",
		vec{0, 1},
	},
	left: {
		left,
		"LEFT",
		vec{-1, 0},
	},
	right: {
		right,
		"RIGHT",
		vec{1, 0},
	},
}

func directionFrom(id directionId) direction {
	dir, found := directions[id]
	if !found {
		panic(fmt.Sprintf(
			"Bad direction: was given a direction ID (enum), but the given value did not match " +
				"any known value: %v\n",
			id,
		))
	}
	return dir
}

func (self direction) rightAngleClockwise() direction {
	switch (self.id) {
	case up: return directionFrom(right)
	case right: return directionFrom(down)
	case down: return directionFrom(left)
	case left: return directionFrom(up)
	default: panic(fmt.Sprintf(
		"Should be unreachable, determining clockwise right angle from %v",
		self,
	))
	}
}

// We assume our input is a string that represents a properly-formed grid, delimited by newlines.
// Don't make me a liar.
// Also I'm PRETTY sure that casting sub-slices to []rune will just map to slices of the original
// backing string/slice, and won't allocate more rune slices. PRETTY sure.
func readGrid(input string) grid {
	newlinePattern := regexp.MustCompile("\n")
	width := newlinePattern.FindStringIndex(input)[0]
	stride := width + 1

	runeGrid := make([][]rune, 0, 1)
	runeGrid = append(runeGrid, []rune(input[0:width]))

	for i := stride; i < len(input); i += stride {
		runeGrid = append(runeGrid, []rune(input[i:i + width]))
	}

	return grid{
			contents: runeGrid,
			height: len(runeGrid),
			width: len(runeGrid[0]),
		}
}

// TODO: factor out grid/vec/other storage to a dedicated, shared file
// see also: day 04
type grid struct {
	contents [][]rune
	height int
	width int
}

func (self grid) charAt(coords vec) rune {
	if coords.x < 0 || coords.x >= self.width || coords.y < 0 || coords.y >= self.height {
		return 0 // null character
	}
	return self.contents[coords.y][coords.x]
}

// Finds a single rune in the grid. If there is more than one instance of that rune in the grid,
// returns the first one found (closest to 0, 0).
func (self grid) findSingleChar(c rune) (location vec, found bool) {
	for x := 0; x < self.width; x++ {
		for y := 0; y < self.height; y++ {
			pos := vec{x, y}
			if self.charAt(pos) == c {
				return pos, true
			}
		}
	}

	return vec{}, false
}

type vec struct {
	x int
	y int
}

func (v vec) add(a vec) vec {
	return vec{
		x: a.x + v.x,
		y: a.y + v.y,
	}
}
