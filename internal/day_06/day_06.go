package day_06

import (
	"fmt"
	"regexp"

	"github.com/lorentzforces/advent-2024/internal/run"
)

func PartOne(input string) (int, error) {
	floorGrid := readGrid(input)
	startingCoords, found := floorGrid.findSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	startingLocation := guardLocation{
		coords: startingCoords,
		facing: directionFrom(up),
	}

	gameState := initGameState(floorGrid, startingLocation)
	for gameState.onMap {
		gameState.doMove()
	}

	return len(gameState.visitedLocations), nil
}

func PartTwo(input string) (int, error) {
	floorGrid := readGrid(input)
	startingCoords, found := floorGrid.findSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	startingLocation := guardLocation{
		coords: startingCoords,
		facing: directionFrom(up),
	}

	gameState := initGameState(floorGrid, startingLocation)
	for gameState.onMap {
		gameState.doMove()
	}

	return len(gameState.loopingObstaclePlacements), nil
}

const impassableTile rune = '#'

type gameState struct {
	floorMap grid

	location guardLocation

	// set of all coordinates visited
	visitedLocations map[vec]struct{}

	// set of all path elements visited - this includes coordinates as well as direction facing
	pathElements map[guardLocation]struct{}

	// set of all locations where it is determined that placing an obstacle would induce a loop
	loopingObstaclePlacements map[vec]struct{}

	// Whether the location in this state is within the bounds of the grid. This is probably
	// better if it were factored into something within the grid implementation itself, but
	// I'm not trying to do that level of effort right now.
	onMap bool
}

func initGameState(floorMap grid, startingLocation guardLocation) *gameState {
	state := gameState{
		floorMap: floorMap,
		location: startingLocation,
		visitedLocations: make(map[vec]struct{}, 0),
		pathElements: make(map[guardLocation]struct{}, 0),
		loopingObstaclePlacements: make(map[vec]struct{}, 0),
		onMap: true,
	}
	state.visitedLocations[state.location.coords] = run.Empty
	state.pathElements[state.location] = run.Empty
	return &state
}

func (self *gameState) doMove() {
	aheadCoords := self.location.coords.add(self.location.facing.unitVec)
	aheadTile := self.floorMap.charAt(aheadCoords)

	if aheadTile == impassableTile {
		self.location.facing = self.location.facing.rightAngleClockwise()
	} else {
		_, alreadyTrodden := self.visitedLocations[aheadCoords]
		// if the guard has already walked a path, there can't be an obstacle there
		if !alreadyTrodden && aheadTile != 0 && self.rightTurnWouldLoop() {
			self.loopingObstaclePlacements[aheadCoords] = run.Empty
		}
		self.location.coords = aheadCoords
	}

	if self.floorMap.charAt(self.location.coords) == 0 {
		self.onMap = false
	} else {
		self.visitedLocations[self.location.coords] = run.Empty
		self.pathElements[self.location] = run.Empty
	}
}

// TODO: this doesn't work, and I'm not 100% sure why (we do know that the correct answer is
// greater than ~800ish)

// Create a ghost and start walking right - if this intersects with a path element already taken,
// that means that turning right would put the guard into a loop. Since turning right at that
// location would loop, that means that placing an obstacle directly ahead would induce a loop.
func (self *gameState) rightTurnWouldLoop() bool {
	newObstacleCoords := self.location.coords.add(self.location.facing.unitVec)
	ghostLoc := guardLocation{
		coords: self.location.coords,
		facing: self.location.facing.rightAngleClockwise(),
	}
	ghostPath := make(map[guardLocation]struct{}, 0)
	ghostPath[ghostLoc] = run.Empty

	// check initial location
	if _, pathSeenBefore := self.pathElements[ghostLoc]; pathSeenBefore {
		return true
	}

	// stupid duplicated logic :'(
	// I'm not proud of it, but building an abstraction seemed more confusing than just doing this
	for {
		aheadCoords := ghostLoc.coords.add(ghostLoc.facing.unitVec)
		aheadTile := self.floorMap.charAt(aheadCoords)

		if aheadTile == impassableTile || newObstacleCoords.equals(aheadCoords) {
			ghostLoc.facing = ghostLoc.facing.rightAngleClockwise()
		} else {
			ghostLoc.coords = aheadCoords
		}

		// path has gone off-map
		if self.floorMap.charAt(ghostLoc.coords) == 0 {
			return false
		}

		if _, pathSeenBefore := self.pathElements[ghostLoc]; pathSeenBefore {
			return true
		}

		if _, ghostPathSeenBefore := ghostPath[ghostLoc]; ghostPathSeenBefore {
			return true
		}

		ghostPath[ghostLoc] = run.Empty
	}
}

type guardLocation struct {
	coords vec
	facing direction
}

func (self guardLocation) String() string {
	return fmt.Sprintf("{coords:%v, facing:%s}", self.coords, self.facing.label)
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

func (self vec) add(a vec) vec {
	return vec{
		x: a.x + self.x,
		y: a.y + self.y,
	}
}

func (self vec) equals(a vec) bool {
	return self.x == a.x && self.y == a.y
}
