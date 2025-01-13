package day_06

import (
	"fmt"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/lorentzforces/advent-2024/internal/spatial"
)

func PartOne(input string) (int, error) {
	floorGrid := spatial.ReadGrid(input)
	startingCoords, found := floorGrid.FindSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	startingLocation := guardLocation{
		coords: startingCoords,
		facing: spatial.DirectionFrom(spatial.Up),
	}

	gameState := initGameState(floorGrid, startingLocation, false)
	for gameState.onMap {
		gameState.doMove()
	}

	return len(gameState.visitedLocations), nil
}

func PartTwo(input string) (int, error) {
	floorGrid := spatial.ReadGrid(input)
	startingCoords, found := floorGrid.FindSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	startingLocation := guardLocation{
		coords: startingCoords,
		facing: spatial.DirectionFrom(spatial.Up),
	}

	gameState := initGameState(floorGrid, startingLocation, true)
	for gameState.onMap {
		gameState.doMove()
	}

	return len(gameState.loopingObstaclePlacements), nil
}

const impassableTile rune = '#'

type gameState struct {
	floorMap spatial.Grid

	location guardLocation

	// set of all coordinates visited
	visitedLocations map[spatial.Vec2d]struct{}

	// set of all path elements visited - this includes coordinates as well as direction facing
	pathElements map[guardLocation]struct{}

	trackingLoopingObstacles bool
	// set of all locations where it is determined that placing an obstacle would induce a loop
	loopingObstaclePlacements map[spatial.Vec2d]struct{}

	// Whether the location in this state is within the bounds of the grid. This is probably
	// better if it were factored into something within the grid implementation itself, but
	// I'm not trying to do that level of effort right now.
	onMap bool
}

func initGameState(
	floorMap spatial.Grid,
	startingLocation guardLocation,
	trackLoopingObstacles bool,
) *gameState {
	state := gameState{
		floorMap: floorMap,
		location: startingLocation,
		visitedLocations: make(map[spatial.Vec2d]struct{}, 0),
		pathElements: make(map[guardLocation]struct{}, 0),
		trackingLoopingObstacles: trackLoopingObstacles,
		loopingObstaclePlacements: make(map[spatial.Vec2d]struct{}, 0),
		onMap: true,
	}
	state.visitedLocations[state.location.coords] = run.Empty
	state.pathElements[state.location] = run.Empty
	return &state
}

func (self *gameState) doMove() {
	aheadCoords := self.location.coords.Add(self.location.facing.UnitVec)
	aheadTile := self.floorMap.CharAt(aheadCoords)

	if aheadTile == impassableTile {
		self.location.facing = rightAngleClockwise(self.location.facing)
	} else {
		_, alreadyTrodden := self.visitedLocations[aheadCoords]
		shouldCheckForLoops :=
			self.trackingLoopingObstacles &&
			// if the guard has already walked a path, there can't be an obstacle there
			!alreadyTrodden &&
			aheadTile != 0
		if shouldCheckForLoops && self.rightTurnWouldLoop() {
			self.loopingObstaclePlacements[aheadCoords] = run.Empty
		}
		self.location.coords = aheadCoords
	}

	if self.floorMap.CharAt(self.location.coords) == 0 {
		self.onMap = false
	} else {
		self.visitedLocations[self.location.coords] = run.Empty
		self.pathElements[self.location] = run.Empty
	}
}

// Create a ghost and start walking right - if this intersects with a path element already taken,
// that means that turning right would put the guard into a loop. Since turning right at that
// location would loop, that means that placing an obstacle directly ahead would induce a loop.
func (self *gameState) rightTurnWouldLoop() bool {
	newObstacleCoords := self.location.coords.Add(self.location.facing.UnitVec)
	ghostLoc := guardLocation{
		coords: self.location.coords,
		facing: rightAngleClockwise(self.location.facing),
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
		aheadCoords := ghostLoc.coords.Add(ghostLoc.facing.UnitVec)
		aheadTile := self.floorMap.CharAt(aheadCoords)

		if aheadTile == impassableTile || newObstacleCoords.Equals(aheadCoords) {
			ghostLoc.facing = rightAngleClockwise(ghostLoc.facing)
		} else {
			ghostLoc.coords = aheadCoords
		}

		// path has gone off-map
		if self.floorMap.CharAt(ghostLoc.coords) == 0 {
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
	coords spatial.Vec2d
	facing spatial.Direction
}

func (self guardLocation) String() string {
	return fmt.Sprintf("{coords:%v, facing:%s}", self.coords, self.facing.Label)
}

func rightAngleClockwise(d spatial.Direction) spatial.Direction {
	switch (d.Id) {
	case spatial.Up: return spatial.DirectionFrom(spatial.Right)
	case spatial.Right: return spatial.DirectionFrom(spatial.Down)
	case spatial.Down: return spatial.DirectionFrom(spatial.Left)
	case spatial.Left: return spatial.DirectionFrom(spatial.Up)
	default: panic(fmt.Sprintf(
		"Should be unreachable, determining clockwise right angle from %v",
		d,
	))
	}
}
