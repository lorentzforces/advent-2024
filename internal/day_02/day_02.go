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
		if rowIsSafe(row) { numSafe++ }
	}

	return numSafe, nil
}

// well at least we're limited to only ONE bad value, because if this went combinatoric I'd be in
// REAL trouble
func PartTwo(input string) (int, error) {
	lines := run.AsLines(input)

	numbers, err := parseInts(lines)
	if err != nil {
		return 0, fmt.Errorf("Parsing error: %w", err)
	}

	numSafe := 0
	for _, row := range numbers {
		if rowIsSafeWithSkip(row) { numSafe++ }
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

func rowIsSafe(row []int) bool {
	firstDiff := row[0] - row[1]

	for i := 0; i < len(row) - 1; i++ {
		if !pairIsSafe(row[i], row[i + 1], firstDiff) {
			return false
		}
	}

	return true
}

func rowIsSafeWithSkip(row []int) bool {
	overallDirection := overallDirection(row)
	badIntervals := make([]int, 0, 2)

	for i := 0; i < len(row) - 1; i++ {
		if !pairIsSafe(row[i], row[i + 1], overallDirection) {
			// if we already have two bad intervals, bail early
			if len(badIntervals) == 2 {
				return false
			}
			badIntervals = append(badIntervals, i)
		}
	}

	switch len(badIntervals) {
		case 0:
			return true
		case 1:
			// if there's only one bad interval, there's an obvious value to try skipping
			badInterval := badIntervals[0]
			if badInterval == 0 { // check start
				return pairIsSafe(row[1], row[2], overallDirection)
			} else if badInterval == len(row) - 2 { // check end
				return pairIsSafe(row[len(row) - 3], row[len(row) - 2], overallDirection)
			} else {
				return pairIsSafe(row[badInterval], row[badInterval + 2], overallDirection)
			}
		case 2:
			// if the bad intervals aren't next to each other, they're separate issues and we
			// can't skip just one number to fix them all
			if badIntervals[1] - badIntervals[0] > 1 {
				return false
			}
			// if they're next to each other, we skip the one value in common
			return pairIsSafe(row[badIntervals[0]], row[badIntervals[1] + 1], overallDirection)
		default:
			return false
	}
}

func pairIsSafe(a, b, prevDirection int) bool {
	diff := a - b
	sameDirection := (diff ^ prevDirection) >= 0 // same sign
	intervalGood := diff != 0 && diff <= 3 && diff >= -3
	return sameDirection && intervalGood
}

// number representing the overall direction of number progression in the row
func overallDirection(row []int) int {
	overallDirection := 0

	for i := 0; i < len(row) - 1; i++ {
		diff := row[i] - row[i + 1]
		if diff > 0 {
			overallDirection++
		}
		if diff < 0 {
			overallDirection--
		}
	}

	return overallDirection
}
