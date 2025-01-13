package day_04

import (
	"regexp"

	"github.com/lorentzforces/advent-2024/internal/spatial"
)

func PartOne(input string) (int, error) {
	wordGrid := spatial.ReadGrid(input)
	results := findInGrid(wordGrid, "XMAS", cardinalAndDiagonalBasis)
	return len(results), nil
}

func PartTwo(input string) (int, error) {
	wordGrid := spatial.ReadGrid(input)
	results := findInGrid(wordGrid, "MAS", diagonalBasis)

	intersections := make(map[spatial.Vec2d]int)
	for _, result := range results {
		secondLetterLoc := result.loc.Add(result.dir)
		currCount, present := intersections[secondLetterLoc]
		if present {
			currCount += 1
		} else {
			currCount = 1
		}
		intersections[secondLetterLoc] = currCount
	}

	exCount := 0
	for _, intersectionCount := range intersections {
		if intersectionCount == 2 {
			exCount++
		}
	}

	return exCount, nil
}

// We assume our input is a string that represents a properly-formed grid, delimited by newlines.
// Don't make me a liar.
// Also I'm PRETTY sure that casting sub-slices to []rune will just map to slices of the original
// backing string/slice, and won't allocate more rune slices. PRETTY sure.
func readGrid(input string) spatial.Grid {
	newlinePattern := regexp.MustCompile("\n")
	width := newlinePattern.FindStringIndex(input)[0]
	stride := width + 1

	runeGrid := make([][]rune, 0, 1)
	runeGrid = append(runeGrid, []rune(input[0:width]))

	for i := stride; i < len(input); i += stride {
		runeGrid = append(runeGrid, []rune(input[i:i + width]))
	}

	return spatial.Grid{
			Contents: runeGrid,
			Height: len(runeGrid),
			Width: len(runeGrid[0]),
		}
}

// Assumes that the word to find is not a palindrome and is at least two letters.
func findInGrid(g spatial.Grid, word string, basisVecs []spatial.Vec2d) []findResult {
	searchLetters := []rune(word)

	startLocations := make([]spatial.Vec2d, 0, 10)
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			pos := spatial.Vec2d{X: x, Y: y}
			if g.CharAt(pos) == searchLetters[0] {
				startLocations = append(startLocations, pos)
			}
		}
	}

	finds := make([]findResult, 0, len(startLocations))
	for _, startPos := range startLocations {
		for _, dir := range basisVecs {
			for currChar := 1; currChar < len(searchLetters); currChar++ {
				checkPos := startPos.Add(dir.Mul(currChar))
				foundNext := g.CharAt(checkPos) == searchLetters[currChar]
				if !foundNext {
					break
				}
				if foundNext && currChar == len(searchLetters) - 1 {
					finds = append(finds, findResult{ loc: startPos, dir: dir })
				}
			}
		}
	}

	return finds
}

type findResult struct {
	loc spatial.Vec2d
	dir spatial.Vec2d
}

var cardinalAndDiagonalBasis = []spatial.Vec2d{
	{ X: -1, Y: -1},
	{ X: -1, Y: 0},
	{ X: -1, Y: 1},
	{ X: 0, Y: -1},
	{ X: 0, Y: 1},
	{ X: 1, Y: -1},
	{ X: 1, Y: 0},
	{ X: 1, Y: 1},
}

var diagonalBasis = []spatial.Vec2d{
	{ X: -1, Y: -1},
	{ X: -1, Y: 1},
	{ X: 1, Y: -1},
	{ X: 1, Y: 1},
}
