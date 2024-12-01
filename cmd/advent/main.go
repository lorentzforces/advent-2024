package main

import (
	"fmt"
	"slices"

	"github.com/lorentzforces/advent-2024/internal/day_01"
	"github.com/lorentzforces/advent-2024/internal/run"
)

// TODO: add CLI parameters to run specific days
func main() {
	results := runAll(runData)
	slices.SortFunc(
		results,
		func(a, b puzzleResult) int {
			if a.day == b.day {
				return a.part - b.part
			}
			return a.day - b.day
		},
	)

	for _, result := range results {
		fmt.Print(result)
		if result.err != nil {
			fmt.Printf("  %s\n", result.PrintErr())
		}
	}
}

var runData = []run.PuzzleData{
	{
		Day: 1,
		Part: 1,
		InputFile: "inputs/day_01_input.txt",
		Fn: func(s string) (any, error) { return day_01.PartOne(s) },
	},
	{
		Day: 1,
		Part: 2,
		InputFile: "inputs/day_01_input.txt",
		Fn: func(s string) (any, error) { return day_01.PartTwo(s) },
	},
}

func runAll(puzzles []run.PuzzleData) []puzzleResult {
	results := make([]puzzleResult, 0, len(puzzles))
	for _, d := range puzzles {
		result := puzzleResult{}
		result.day = d.Day
		result.part = d.Part

		input, err := run.GetFileContents(d.InputFile)
		if err != nil {
			result.err = err
			results = append(results, result)
			continue
		}

		output, err := d.Fn(input)
		result.err = err
		result.output = fmt.Sprint(output)
		results = append(results, result)
	}

	return results
}

type puzzleResult struct {
	day int
	part int
	output string
	err error
}

func (pr puzzleResult) String() string {
	return fmt.Sprintf("Day %02d, Part %02d output: %s\n", pr.day, pr.part, pr.output)
}

func (pr puzzleResult) PrintErr() string {
	if pr.err == nil {
		return "No error!"
	}
	return fmt.Sprintf("ERROR: %s", pr.err.Error())
}
