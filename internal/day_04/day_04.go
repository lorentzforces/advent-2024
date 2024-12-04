package day_04

import (
	"regexp"
)

func PartOne(input string) (int, error) {
	grid := readGrid(input)
	searchLetters := []rune("XMAS")

	startLocations := make([]vec, 0, 10)
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			pos := vec{x, y}
			if grid.charAt(pos) == searchLetters[0] {
				startLocations = append(startLocations, pos)
			}
		}
	}

	finds := 0
	for _, startPos := range startLocations {
		for _, dir := range basisVecs {
			for currChar := 1; currChar < len(searchLetters); currChar++ {
				checkPos := startPos.add(dir.mul(currChar))
				foundNext := grid.charAt(checkPos) == searchLetters[currChar]
				if !foundNext {
					break
				}
				if foundNext && currChar == len(searchLetters) - 1 {
					finds++
				}
			}
		}
	}

	return finds, nil
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

type grid struct {
	contents [][]rune
	height int
	width int
}

func (g grid) charAt(coords vec) rune {
	if coords.x < 0 || coords.x >= g.width || coords.y < 0 || coords.y >= g.height {
		return 0 // null character
	}
	return g.contents[coords.y][coords.x]
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

func (v vec) mul(a int) vec {
	return vec{
		x: v.x * a,
		y: v.y * a,
	}
}

var basisVecs = []vec{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}
