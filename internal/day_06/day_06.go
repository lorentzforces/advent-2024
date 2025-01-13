package day_06

import (
	"fmt"

	"github.com/lorentzforces/advent-2024/internal/spatial"
	"github.com/lorentzforces/advent-2024/internal/stores"
)

func PartOne(input string) (int, error) {
	floorGrid := spatial.ReadGrid(input)
	startingCoords, found := floorGrid.FindSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	startingLocation := guardLocation{
		coords: startingCoords,
		facing: spatial.Up.Into(),
	}

	gameState := initGameState(floorGrid, startingLocation, false)
	for gameState.onMap {
		gameState.doMove()
	}

	return gameState.visitedLocations.Len(), nil
}

func PartTwo(input string) (int, error) {
	floorGrid := spatial.ReadGrid(input)
	startingCoords, found := floorGrid.FindSingleChar('^')
	if !found {
		return 0, fmt.Errorf("Starting character not found in grid")
	}

	startingLocation := guardLocation{
		coords: startingCoords,
		facing: spatial.Up.Into(),
	}

	gameState := initGameState(floorGrid, startingLocation, true)
	for gameState.onMap {
		gameState.doMove()
	}

	return gameState.loopingObstaclePlacements.Len(), nil
}

const impassableTile rune = '#'

type gameState struct {
	floorMap spatial.Grid

	location guardLocation

	// set of all coordinates visited
	visitedLocations stores.Set[spatial.Vec2d]

	// set of all path elements visited - this includes coordinates as well as direction facing
	pathElements stores.Set[guardLocation]

	trackingLoopingObstacles bool
	// set of all locations where it is determined that placing an obstacle would induce a loop
	loopingObstaclePlacements stores.Set[spatial.Vec2d]

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
		visitedLocations: stores.EmptySet[spatial.Vec2d](),
		pathElements: stores.EmptySet[guardLocation](),
		trackingLoopingObstacles: trackLoopingObstacles,
		loopingObstaclePlacements: stores.EmptySet[spatial.Vec2d](),
		onMap: true,
	}
	state.visitedLocations.Put(state.location.coords)
	state.pathElements.Put(state.location)
	return &state
}

func (self *gameState) doMove() {
	aheadCoords := self.location.coords.Add(self.location.facing.UnitVec)
	aheadTile := self.floorMap.CharAt(aheadCoords)

	if aheadTile == impassableTile {
		self.location.facing = self.location.facing.StepClockwise()
	} else {
		alreadyTrodden := self.visitedLocations.Contains(aheadCoords)
		shouldCheckForLoops :=
			self.trackingLoopingObstacles &&
			// if the guard has already walked a path, there can't be an obstacle there
			!alreadyTrodden &&
			aheadTile != 0
		if shouldCheckForLoops && self.rightTurnWouldLoop() {
			self.loopingObstaclePlacements.Put(aheadCoords)
		}
		self.location.coords = aheadCoords
	}

	if self.floorMap.IsOutOfBounds(self.location.coords) {
		self.onMap = false
	} else {
		self.visitedLocations.Put(self.location.coords)
		self.pathElements.Put(self.location)
	}
}

// Create a ghost and start walking right - if this intersects with a path element already taken,
// that means that turning right would put the guard into a loop. Since turning right at that
// location would loop, that means that placing an obstacle directly ahead would induce a loop.
func (self *gameState) rightTurnWouldLoop() bool {
	newObstacleCoords := self.location.coords.Add(self.location.facing.UnitVec)
	ghostLoc := guardLocation{
		coords: self.location.coords,
		facing: self.location.facing.StepClockwise(),
	}
	ghostPath := stores.EmptySet[guardLocation]()
	ghostPath.Put(ghostLoc)

	// check initial location
	if self.pathElements.Contains(ghostLoc) {
		return true
	}

	// stupid duplicated logic :'(
	// I'm not proud of it, but building an abstraction seemed more confusing than just doing this
	for {
		aheadCoords := ghostLoc.coords.Add(ghostLoc.facing.UnitVec)
		aheadTile := self.floorMap.CharAt(aheadCoords)

		if aheadTile == impassableTile || newObstacleCoords.Equals(aheadCoords) {
			ghostLoc.facing = ghostLoc.facing.StepClockwise()
		} else {
			ghostLoc.coords = aheadCoords
		}

		// path has gone off-map
		if self.floorMap.IsOutOfBounds(ghostLoc.coords) {
			return false
		}

		if self.pathElements.Contains(ghostLoc) {
			return true
		}

		if ghostPath.Contains(ghostLoc) {
			return true
		}

		ghostPath.Put(ghostLoc)
	}
}

type guardLocation struct {
	coords spatial.Vec2d
	facing spatial.Direction
}

func (self guardLocation) String() string {
	return fmt.Sprintf("{coords:%v, facing:%s}", self.coords, self.facing.Label)
}
