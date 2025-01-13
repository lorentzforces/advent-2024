package spatial

import (
	"fmt"
	"regexp"
)

type Vec2d struct {
	X int
	Y int
}

func (self Vec2d) Add(a Vec2d) Vec2d {
	return Vec2d{
		X: a.X + self.X,
		Y: a.Y + self.Y,
	}
}

func (self Vec2d) Mul(a int) Vec2d {
	return Vec2d{
		X: self.X * a,
		Y: self.Y * a,
	}
}

func (self Vec2d) Equals(a Vec2d) bool {
	return self.X == a.X && self.Y == a.Y
}

type Grid struct {
	Contents [][]rune
	Height int
	Width int
	LocationsOfConcern []Vec2d
}

// Create a grid out of the given input string
// We assume our input is a string that represents a properly-formed grid, delimited by newlines.
// Don't make me a liar.
// Also I'm PRETTY sure that casting sub-slices to []rune will just map to slices of the original
// backing string/slice, and won't allocate more rune slices. PRETTY sure.
func ReadGrid(input string) Grid {
	newlinePattern := regexp.MustCompile("\n")
	width := newlinePattern.FindStringIndex(input)[0]
	stride := width + 1

	runeGrid := make([][]rune, 0, 1)
	runeGrid = append(runeGrid, []rune(input[0:width]))

	for i := stride; i < len(input); i += stride {
		runeGrid = append(runeGrid, []rune(input[i:i + width]))
	}

	return Grid{
			Contents: runeGrid,
			Height: len(runeGrid),
			Width: len(runeGrid[0]),
		}
}

func (self Grid) CharAt(coords Vec2d) rune {
	if coords.X < 0 || coords.X >= self.Width || coords.Y < 0 || coords.Y >= self.Height {
		return 0 // null character
	}
	return self.Contents[coords.Y][coords.X]
}

// Finds a single rune in the grid. If there is more than one instance of that rune in the grid,
// returns the first one found (closest to 0, 0).
func (self Grid) FindSingleChar(c rune) (location Vec2d, found bool) {
	for x := 0; x < self.Width; x++ {
		for y := 0; y < self.Height; y++ {
			pos := Vec2d{X: x, Y: y}
			if self.CharAt(pos) == c {
				return pos, true
			}
		}
	}

	return Vec2d{}, false
}

// We'll see how this pattern goes. This is basically trying to implement some kind of enumeration
// over directions, but I'm starting to think it's not better than just having a bunch of named
// "constant" struct values.
// Good lord Golang just give us sum types. They don't even need to be fancy.
type DirectionId uint8
const (
	Up DirectionId = iota
	Down
	Left
	Right
)

type Direction struct {
	Id DirectionId
	Label string
	UnitVec Vec2d
}

var directions = map[DirectionId]Direction{
	Up: {
		Up,
		"UP",
		Vec2d{X: 0, Y: -1},
	},
	Down: {
		Down,
		"DOWN",
		Vec2d{X: 0, Y: 1},
	},
	Left: {
		Left,
		"LEFT",
		Vec2d{X: -1, Y: 0},
	},
	Right: {
		Right,
		"RIGHT",
		Vec2d{X: 1, Y: 0},
	},
}

func DirectionFrom(id DirectionId) Direction {
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
