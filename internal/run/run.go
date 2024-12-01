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

func BailIfFailed(t *testing.T) {
	if t.Failed() { t.FailNow() }
}
