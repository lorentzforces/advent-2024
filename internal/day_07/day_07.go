package day_07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lorentzforces/advent-2024/internal/run"
)

func PartOne(input string) (uint64, error) {
	rawLines := run.AsLines(input)
	equations, err := parseEquationParams(rawLines, partOneOperators)
	if err != nil {
		return 0, err
	}

	total := uint64(0)
	for _, equation := range equations {
		if equation.hasValidOperators() {
			total += equation.target
		}
	}
	return total, nil
}

func PartTwo(input string) (uint64, error) {
	rawLines := run.AsLines(input)
	equations, err := parseEquationParams(rawLines, partTwoOperators)
	if err != nil {
		return 0, err
	}

	total := uint64(0)
	for _, equation := range equations {
		if equation.hasValidOperators() {
			total += equation.target
		}
	}
	return total, nil
}

type equationParams struct{
	target uint64
	terms []uint64
	numOperators int
	validOperators []operator
}

func parseEquationParams(lines []string, validOperators []operator) ([]equationParams, error) {
	results := make([]equationParams, 0, len(lines))
	for lineNr, line := range lines {
		allItems := strings.Fields(line)

		targetRunes := []rune(allItems[0])
		targetNumberRunes := targetRunes[:len(targetRunes) - 1] // chop off the final character, which is a ":"
		target, parseErr := strconv.ParseUint(string(targetNumberRunes), 10, 64)
		if parseErr != nil {
			return nil, fmt.Errorf(
				"Error parsing uint64 from \"%s\" on line %d: %w",
				string(targetNumberRunes), lineNr, parseErr,
			)
		}

		lineData := equationParams{
			target: target,
			terms: make([]uint64, len(allItems) - 1),
			numOperators: len(allItems) - 2,
			validOperators: validOperators,
		}

		for i, termStr := range allItems[1:] {
			val, parseErr := strconv.ParseUint(termStr, 10, 64)
			if parseErr != nil {
				return nil, fmt.Errorf(
					"Error parsing uint64 from \"%s\" on line %d: %w",
					termStr, lineNr, parseErr,
				)
			}
			lineData.terms[i] = val
		}

		results = append(results, lineData)
	}

	return results, nil
}

func (self equationParams) hasValidOperators() bool {
	return validOperatorRecursion(self, make([]operator, 0), uint64(self.terms[0]))
}

func validOperatorRecursion(equation equationParams, incOperators []operator, runningTotal uint64) bool {
	nextTerm := equation.terms[len(incOperators) + 1]
	for _, op := range equation.validOperators {
		opTotal := applyOperator(op, runningTotal, nextTerm)
		switch {
		// if we've already overshot the target, there's no hope
		case opTotal > equation.target: continue
		case len(incOperators) == equation.numOperators - 1: {
			if opTotal == equation.target {
				return true
			}
		}
		default:{
			newIncOperators := make([]operator, len(incOperators) + 1)
			copy(newIncOperators, incOperators)
			newIncOperators[len(incOperators)] = op
			if validOperatorRecursion(equation, newIncOperators, opTotal) {
				return true
			}
		}
		}
	}

	return false
}

func applyOperator(op operator, left uint64, right uint64) uint64 {
	switch op {
	case addition: return left + right
	case multiplication: return left * right
	case concatenation: {
		tenPower := uint64(10)
		for tenPower <= right {
			tenPower *= 10
		}
		return left * tenPower + right
		// The following was taken from a random person's solution to this same puzzle, and it runs
		// SIGNIFICANTLY slower. This is somewhat surprising since it (in theory) branches/loops less.
		// That being said
		// numDigits := int(math.Floor(math.Log10(float64(right))) + 1)
		// return left * uint64(math.Pow10(numDigits)) + right
	}
	default: panic(fmt.Sprintf(
		"Should be unreachable, determining operator from %v",
		op,
	))
	}
}

type operator int

const (
	addition operator = 0
	multiplication operator = 1
	concatenation operator = 2
)

var partOneOperators = []operator{
	addition,
	multiplication,
}
var partTwoOperators = []operator{
	addition,
	multiplication,
	concatenation,
}

func (self operator) String() string {
	switch self {
	case addition: return "ADD"
	case multiplication: return "MULT"
	case concatenation: return "CONCAT"
	default: panic(fmt.Sprintf("invalid math operator:%d", self))
	}
}
