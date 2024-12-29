package run

import (
	"os"
	"strings"
	"testing"
)

type RunFunc func(string) (any, error)

type PuzzleData struct {
	Day int
	Part int
	InputFile string
	Fn RunFunc
}

func GetFileContents(path string) (string, error) {
	fileBuf, err := os.ReadFile(path)
	if err != nil { return "", err }
	return string(fileBuf), nil
}

func AsLines(s string) []string {
	lines := strings.Split(s, "\n")

	// trim trailing blank line (expected)
	if lines[len(lines) - 1] == "" {
		lines = lines[0:len(lines) - 1]
	}
	return lines
}

func AsLinesSplitOnBlanks(s string) [][]string {
	lines := AsLines(s)

	splits := make([][]string, 0, 1)
	start := 0
	for i, line := range lines {
		if line == "" {
			splits = append(splits, lines[start:i])
			start = i + 1
		}
	}

	splits = append(splits, lines[start:])
	return splits
}

func BailIfFailed(t *testing.T) {
	if t.Failed() { t.FailNow() }
}

var Empty struct{}
