package day_05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/lorentzforces/advent-2024/internal/run"
	"github.com/lorentzforces/advent-2024/internal/stores"
)

func PartOne(input string) (int, error) {
	lineSplits := run.AsLinesSplitOnBlanks(input)
	if len(lineSplits) != 2 {
		return 0, fmt.Errorf(
			"expected 2 chunks of lines separated by blanks, but found %d instead",
			len(lineSplits),
		)
	}

	constraints, err := parseOrderings(lineSplits[0])
	if err != nil {
		return 0, fmt.Errorf("error parsing orderings: %w", err)
	}

	sequences, err := parseLists(lineSplits[1])
	if err != nil {
		return 0, fmt.Errorf("error parsing ")
	}

	total := 0
	for _, seq := range sequences {
		if len(seq) % 2 != 1 {
			return 0, fmt.Errorf("sequence %v was an even number of elements", seq)
		}
		if sequenceValid := constraints.meetsConstraints(seq); sequenceValid {
			total += getMiddleValue(seq)
		}
	}

	return total, nil
}

func PartTwo(input string) (int, error) {
	lineSplits := run.AsLinesSplitOnBlanks(input)
	if len(lineSplits) != 2 {
		return 0, fmt.Errorf(
			"expected 2 chunks of lines separated by blanks, but found %d instead",
			len(lineSplits),
		)
	}

	constraints, err := parseOrderings(lineSplits[0])
	if err != nil {
		return 0, fmt.Errorf("error parsing orderings: %w", err)
	}

	sequences, err := parseLists(lineSplits[1])
	if err != nil {
		return 0, fmt.Errorf("error parsing ")
	}

	total := 0
	for _, seq := range sequences {
		if len(seq) % 2 != 1 {
			return 0, fmt.Errorf("sequence %v was an even number of elements", seq)
		}

		// Editorial note: this only works if the input rules cover all potential comparisons
		// between adjacent values. For example, if we have the following scenario:
		// - a sequence of [1, 11, 200]
		// - a rule that 1 must come after 200
		// - no rule relating 1 and 11
		// ... then any sorting method relying on comparing adjacent values will consider this
		// sequence to be already sorted.
		// Luckily for us, it appears that the input satisfies this requirement.
		if sequenceValid := constraints.meetsConstraints(seq); !sequenceValid {
			constraints.sortSequence(seq) // sorts in-place
			total += getMiddleValue(seq)
		}
	}

	return total, nil
}

// Parses orderings of the form "11|99" from the given lines, where the number before the pipe must
// appear BEFORE the number after the pipe.
// Result value is a mapping of those orderings BY their "after" value.
func parseOrderings(orderings []string) (constraintData, error) {
	constraints := makeConstraintData()
	for _, line := range orderings {
		vals := strings.Split(line, "|")
		if len(vals) != 2 {
			return constraintData{}, fmt.Errorf(
				"expected a line with 2 pipe-delimited integer values, but found %d instead with " +
					"line \"%s\"",
				len(vals),
				line,
			)
		}

		rawBefore, err:= strconv.ParseInt(vals[0], 10, 0)
		if err != nil {
			return constraintData{}, fmt.Errorf("error parsing integer: %s %w", vals[0], err)
		}
		rawAfter, err:= strconv.ParseInt(vals[1], 10, 0)
		if err != nil {
			return constraintData{}, fmt.Errorf("error parsing integer: %s %w", vals[1], err)
		}
		before := int(rawBefore)
		after := int(rawAfter)

		constraints.registerConstraint(ordering{before, after})
	}

	return constraints, nil
}

// parse lists of values, referred to in puzzle as "updates"
func parseLists(lines []string) ([][]int, error) {
	lists := make([][]int, len(lines))
	for i, line := range lines {
		rawVals := strings.Split(line, ",")

		vals := make([]int, len(rawVals))
		for j, rawVal := range(rawVals) {
			val, err := strconv.ParseInt(rawVal, 10, 0)
			if err != nil {
				return nil, fmt.Errorf(
					"error parsing integer from %s on line \"%s\" %w",
					rawVal, line, err,
				)
			}
			vals[j] = int(val)
		}
		lists[i] = vals
	}

	return lists, nil
}

type constraintData struct {
	// mapping of orderings according to their "after" value
	pool map[int]stores.Set[ordering]
	currBadValues stores.Set[int]
}

func makeConstraintData() constraintData {
	return constraintData {
		pool: make(map[int]stores.Set[ordering]),
		currBadValues: stores.EmptySet[int](),
	}
}

func (self *constraintData) meetsConstraints(seq []int) bool {
	self.resetCurrentConstraints()
	for _, val := range seq {
		if valueValid := self.seeValue(val); !valueValid {
			return false
		}
	}
	return true
}

func (self *constraintData) resetCurrentConstraints() {
	self.currBadValues = stores.EmptySet[int]()
}

func (self *constraintData) registerConstraint(o ordering) {
	orderingsForKey := getOrInit(self.pool, o.after)
	orderingsForKey.Put(o)
	self.pool[o.after] = orderingsForKey
}

// Registers seeing a value.
// Returns false if the value violates ordering rules based on prior values seen since
// resetCurrentConstraints() was last called.
func (self *constraintData) seeValue(val int) bool {
	if self.currBadValues.Contains(val) {
		return false
	}

	orderingsForKey, hasOrderings := self.pool[val]
	if hasOrderings {
		for ordering := range orderingsForKey.Vals() {
			self.currBadValues.Put(ordering.before)
		}
	}
	return true
}

// Sort the given sequence according to the constraint rules.
// A & B
// if we have an ordering (a, b), then:
// - b comes after a
// - a comes before b
func (self *constraintData) sortSequence(seq []int) {
	slices.SortFunc(seq, func(a, b int) int {
		aRules := self.pool[a]
		bRules := self.pool[b]
		aComesAfter := hasOrderingWithValBefore(aRules, b)
		bComesAfter := hasOrderingWithValBefore(bRules, a)

		switch {
		case bComesAfter: return -1
		case aComesAfter: return 1
		default: return 0
		}
	})
}

func getOrInit(orderings map[int]stores.Set[ordering], key int) stores.Set[ordering] {
	val, present := orderings[key]
	if !present {
		val = stores.EmptySet[ordering]()
	}
	return val
}

type ordering struct {
	before int
	after int
}

type intSet map[int]struct{}

type orderingSet map[ordering]struct{}

func hasOrderingWithValBefore(set stores.Set[ordering], n int) bool {
	for ord := range set.Vals() {
		if ord.before == n { return true }
	}
	return false
}

func hasOrderingWithValAfter(set stores.Set[ordering], n int) bool {
	for ord := range set.Vals() {
		if ord.after == n { return true }
	}
	return false
}

func getMiddleValue(seq []int) int {
	middleIndex := len(seq) / 2
	return seq[middleIndex]
}
