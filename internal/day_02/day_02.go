package day_02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lorentzforces/advent-2024/internal/run"
)

func PartOne(input string) (int, error) {
	lines := run.AsLines(input)

	numbers, err := parseInts(lines)
	if err != nil {
		return 0, fmt.Errorf("Parsing error: %w", err)
	}

	numSafe := 0
	for _, row := range numbers {
		if reportIsSafe(row) {
			numSafe++
		}
	}

	return numSafe, nil
}

func parseInts(lines []string) ([][]int, error) {
	results := make([][]int, len(lines))
	for i, line := range lines {
		lineNums := make([]int, 0, 5)
		for _, term := range strings.Fields(line) {
			number, err := strconv.ParseInt(term, 10, 0)
			if err != nil {
				return nil, fmt.Errorf("Error parsing number on line %d: %w", i, err)
			}

			lineNums = append(lineNums, int(number))
		}
		results[i] = lineNums
	}

	return results, nil
}

func reportIsSafe(row []int) bool {
	firstDiff := row[0] - row[1]

	for i := 0; i < len(row) - 1; i++ {
		diff := row[i] - row[i + 1]
		sameDir := (diff ^ firstDiff) >= 0 // same sign
		if !sameDir { return false }
		if diff == 0 || diff > 3 || diff < -3 { return false }
	}

	return true
}

type Direction int
const (
	Increasing Direction = iota
	Decreasing
)

func diffDirection(a, b int) Direction {
	if a < b { return Increasing }
	return Decreasing
}
