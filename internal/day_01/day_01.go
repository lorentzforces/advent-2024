package day_01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/lorentzforces/advent-2024/internal/run"
)

func PartOne(input string) (uint, error) {
	lines := run.AsLines(input)

	left, right, err := parseTwoLists(lines)
	if err != nil {
		return 0, fmt.Errorf("Parsing error: %w", err)
	}

	slices.Sort(left)
	slices.Sort(right)

	var totalDistance uint = 0
	for i := 0; i < len(left); i++ {
		leftVal := left[i]
		rightVal := right[i]
		if leftVal < rightVal {
			totalDistance += rightVal - leftVal
		} else {
			totalDistance += leftVal - rightVal
		}
	}

	return totalDistance, nil
}

func parseTwoLists(lines []string) (left []uint, right []uint, err error) {
	left = make([]uint, len(lines))
	right = make([]uint, len(lines))
	for i, line := range lines {
		strVals := strings.Fields(line)
		if len(strVals) != 2 {
			err = fmt.Errorf(
				"Line %d parsed with %d values (not 2)",
				i + 1, len(strVals),
			)
			return nil, nil, err
		}

		leftVal, err := strconv.ParseUint(strVals[0], 10, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("Error parsing left value on line %d: %w", i + 1, err)
		}
		rightVal, err := strconv.ParseUint(strVals[1], 10, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("Error parsing right value on line %d: %w", i + 1, err)
		}

		left[i] = uint(leftVal)
		right[i] = uint(rightVal)
	}

	return left, right, nil
}
