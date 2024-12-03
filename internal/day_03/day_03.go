package day_03

import (
	"fmt"
	"regexp"
	"strconv"
)

func PartOne(input string) (int, error) {
	mulMatch := regexp.MustCompile(`mul\((?<a>\d+),(?<b>\d+)\)`)
	finds := mulMatch.FindAllStringSubmatch(input, -1)

	total := 0
	for _, find := range finds {
		a, err := strconv.ParseInt(find[1], 10, 0)
		if err != nil {
			return 0, fmt.Errorf("couldn't parse int a in section \"%s\" %w", find[0], err)
		}
		b, err := strconv.ParseInt(find[2], 10, 0)
		if err != nil {
			return 0, fmt.Errorf("couldn't parse int b in section \"%s\" %w", find[0], err)
		}

		total += int(a) * int(b)
	}

	return total, nil
}
