package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/lorentzforces/advent-2024/internal/day_01"
	"github.com/lorentzforces/advent-2024/internal/day_02"
	"github.com/lorentzforces/advent-2024/internal/day_03"
	"github.com/lorentzforces/advent-2024/internal/day_04"
	"github.com/lorentzforces/advent-2024/internal/day_05"
	"github.com/lorentzforces/advent-2024/internal/day_06"
	"github.com/lorentzforces/advent-2024/internal/day_07"
	"github.com/lorentzforces/advent-2024/internal/run"
)

// TODO: add CLI parameters to run specific days
// TODO: parallelize running days
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
	{
		Day: 2,
		Part: 1,
		InputFile: "inputs/day_02_input.txt",
		Fn: func(s string) (any, error) { return day_02.PartOne(s) },
	},
	{
		Day: 2,
		Part: 2,
		InputFile: "inputs/day_02_input.txt",
		Fn: func(s string) (any, error) { return day_02.PartTwo(s) },
	},
	{
		Day: 3,
		Part: 1,
		InputFile: "inputs/day_03_input.txt",
		Fn: func(s string) (any, error) { return day_03.PartOne(s) },
	},
	{
		Day: 3,
		Part: 2,
		InputFile: "inputs/day_03_input.txt",
		Fn: func(s string) (any, error) { return day_03.PartTwo(s) },
	},
	{
		Day: 4,
		Part: 1,
		InputFile: "inputs/day_04_input.txt",
		Fn: func(s string) (any, error) { return day_04.PartOne(s) },
	},
	{
		Day: 4,
		Part: 2,
		InputFile: "inputs/day_04_input.txt",
		Fn: func(s string) (any, error) { return day_04.PartTwo(s) },
	},
	{
		Day: 5,
		Part: 1,
		InputFile: "inputs/day_05_input.txt",
		Fn: func(s string) (any, error) { return day_05.PartOne(s) },
	},
	{
		Day: 5,
		Part: 2,
		InputFile: "inputs/day_05_input.txt",
		Fn: func(s string) (any, error) { return day_05.PartTwo(s) },
	},
	{
		Day: 6,
		Part: 1,
		InputFile: "inputs/day_06_input.txt",
		Fn: func(s string) (any, error) { return day_06.PartOne(s) },
	},
	{
		Day: 6,
		Part: 2,
		InputFile: "inputs/day_06_input.txt",
		Fn: func(s string) (any, error) { return day_06.PartTwo(s) },
	},
	{
		Day: 7,
		Part: 1,
		InputFile: "inputs/day_07_input.txt",
		Fn: func(s string) (any, error) { return day_07.PartOne(s) },
	},
	{
		Day: 7,
		Part: 2,
		InputFile: "inputs/day_07_input.txt",
		Fn: func(s string) (any, error) { return day_07.PartTwo(s) },
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

		// we don't include file read time in our timing, just the actual solution execution
		start := time.Now()
		output, err := d.Fn(input)
		result.duration = time.Since(start)

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
	duration time.Duration
}

func (pr puzzleResult) String() string {
	return fmt.Sprintf(
		"Day %02d, Part %02d output: %s  [%s]\n",
		pr.day, pr.part, pr.output, pr.duration,
	)
}

func (pr puzzleResult) PrintErr() string {
	if pr.err == nil {
		return "No error!"
	}
	return fmt.Sprintf("ERROR: %s", pr.err.Error())
}
