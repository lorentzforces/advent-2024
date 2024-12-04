package day_03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulExpr = `mul\((?<a>\d+),(?<b>\d+)\)`
var doExpr = `do\(\)`
var dontExpr = `don't\(\)`

func PartOne(input string) (int, error) {
	mulMatch := regexp.MustCompile(mulExpr)
	finds := mulMatch.FindAllStringSubmatch(input, -1)

	total := 0
	for _, find := range finds {
		value, err := getMulMatchValue(find)
		if err != nil { return 0, err }
		total += value
	}

	return total, nil
}

func PartTwo(input string) (int, error) {
	doDontExpr := strings.Join(
		[]string{wrapNonCapturing(mulExpr), wrapNonCapturing(doExpr), wrapNonCapturing(dontExpr)},
		"|",
	)
	doDontMatch := regexp.MustCompile(doDontExpr)
	mulMatch := regexp.MustCompile(mulExpr)

	start := 0
	enabled := true
	total := 0
	for {
		if start >= len(input) { break }
		matchRegion := doDontMatch.FindStringIndex(input[start:])
		if matchRegion == nil { break }
		matchStr := input[start + matchRegion[0]:start + matchRegion[1]]
		start = start + matchRegion[1]

		switch {
			case strings.HasPrefix(matchStr, "don't"):
				enabled = false
			case strings.HasPrefix(matchStr, "do"):
				enabled = true
			case strings.HasPrefix(matchStr, "mul"):
				if !enabled { continue }
				find := mulMatch.FindStringSubmatch(matchStr)
				value, err := getMulMatchValue(find)
				if err != nil { return 0, err }
				total += value
			default:
				return 0, fmt.Errorf("couldn't match expected prefixes in section \"%s\"", matchStr)
		}
	}

	return total, nil
}

func wrapNonCapturing(exp string) string {
	return `(?:` + exp + `)`
}

func getMulMatchValue(submatch []string) (int, error) {
	a, err := strconv.ParseInt(submatch[1], 10, 0)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse int a in section \"%s\": %w", submatch[0], err)
	}
	b, err := strconv.ParseInt(submatch[2], 10, 0)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse int b in section \"%s\": %w", submatch[0], err)
	}

	return int(a) * int(b), nil
}
