package day_01

import "testing"

var testInput string =
`3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPartOneSampleInput(t *testing.T) {
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("ERROR: %s", err)
	}

	if result != 11 {
		t.Fatalf("Expected output to be 11, but was %d instead", result)
	}
}
