package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition rune

func (c Condition) String() string {
	return fmt.Sprintf("%c", c)
}

const (
	Operational Condition = '.'
	Damaged               = '#'
	Unknown               = '?'
)

type Record struct {
	Conditions    []Condition
	DamagedGroups []int64
}

func Solutions(conds []Condition, damagedGroups []int64, damagedCount int64) ([][]Condition, bool) {
	solutions := make([][]Condition, 0)

	// Evertying is done
	if len(conds) == 0 && len(damagedGroups) == 0 && damagedCount == 0 {
		return solutions, true
	}

	// No more input and one group of damaged
	if len(conds) == 0 {
		return solutions, len(damagedGroups) == 1 && damagedCount == damagedGroups[0]
	}

	var targetDamaged int64 = 0
	if len(damagedGroups) > 0 {
		targetDamaged = damagedGroups[0]
	}

	// Too many damaged no further solutions
	if damagedCount > targetDamaged {
		return solutions, false
	}

	c := conds[0]

	if c == Unknown {
		var altSols [][]Condition
		altConds := make([]Condition, 0, len(conds))
		altConds = append(altConds, conds...)

		// Operational Reality
		altConds[0] = Operational
		altSols, ok := Solutions(altConds, damagedGroups, damagedCount)
		if ok {
			solutions = append(solutions, altSols...)
		}

		// Damaged Reality
		altConds[0] = Damaged
		altSols, ok = Solutions(altConds, damagedGroups, damagedCount)
		if ok {
			solutions = append(solutions, altSols...)
		}

		return solutions, len(solutions) > 0
	}

	if c == Operational {
		if damagedCount > 0 {
			if damagedCount != targetDamaged {
				return [][]Condition{}, false
			} else {
				damagedGroups = damagedGroups[1:]
				damagedCount = 0
			}
		}
	}
	if c == Damaged {
		damagedCount += 1
	}

	remSolutions, ok := Solutions(conds[1:], damagedGroups, damagedCount)
	if !ok {
		return solutions, false
	}

	for _, remSol := range remSolutions {
		solution := make([]Condition, 0, len(conds))
		solution = append(solution, c)
		solution = append(solution, remSol...)

		solutions = append(solutions, solution)
	}

	if len(solutions) == 0 {
		solution := []Condition{c}
		solutions = append(solutions, solution)
	}

	return solutions, true
}

func (r Record) Solutions() [][]Condition {
	sols, ok := Solutions(r.Conditions, r.DamagedGroups, 0)
	if !ok {
		panic("unable to find solution")
	}

	return sols
}

func main() {
	sum := 0

	var numberOfCopies int8 = 1
	if len(os.Args) > 1 {
		n, err := strconv.ParseInt(os.Args[1], 10, 8)
		if err != nil {
			panic(err)
		}
		numberOfCopies = int8(n)
	}

	records := parseRecords(os.Stdin, numberOfCopies)
	for i, r := range records {
		fmt.Fprintln(os.Stderr, r)
		solutions := r.Solutions()

		fmt.Fprintf(os.Stdout, "Row %d/%d ->\t%d\n", i+1, len(records), len(solutions))

		for _, sol := range solutions {
			fmt.Fprintln(os.Stderr, "\t", sol)
			sum += 1
		}

		fmt.Fprintln(os.Stderr)
	}

	fmt.Fprintln(os.Stdout, sum)
}

func parseRecords(f *os.File, numCopies int8) []Record {
	rs := make([]Record, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rConds, rGroups, ok := strings.Cut(line, " ")
		if !ok {
			continue
		}

		// Parse conditions
		conds := make([]Condition, 0, len(rConds))
		for _, rCond := range rConds {
			cond := Condition(rCond)
			conds = append(conds, cond)
		}

		// Parse groups
		groups := make([]int64, 0)
		for _, rGroup := range strings.Split(rGroups, ",") {
			n, _ := strconv.ParseInt(rGroup, 10, 64)

			groups = append(groups, n)
		}

		r := Record{
			Conditions:    conds,
			DamagedGroups: groups,
		}

		// Expand input
		for i := int8(1); i < numCopies; i++ {
			r.Conditions = append(r.Conditions, Unknown)
			r.Conditions = append(r.Conditions, conds...)

			r.DamagedGroups = append(r.DamagedGroups, groups...)
		}

		rs = append(rs, r)
	}

	return rs
}
