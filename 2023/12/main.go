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

func Solutions(conds []Condition, damagedGroups []int64, damagedCount int64) int64 {
	// Evertying is done
	if len(conds) == 0 && len(damagedGroups) == 0 && damagedCount == 0 {
		return 1
	}

	// No more input and one group of damaged
	if len(conds) == 0 {
		if len(damagedGroups) == 1 && damagedCount == damagedGroups[0] {
			return 1
		} else {
			return 0
		}
	}

	var targetDamaged int64 = 0
	if len(damagedGroups) > 0 {
		targetDamaged = damagedGroups[0]
	}

	// Too many damaged no further solutions
	if damagedCount > targetDamaged {
		return 0
	}

	c := conds[0]

	if c == Unknown {
		altConds := make([]Condition, 0, len(conds))
		altConds = append(altConds, conds...)

		// Operational Reality
		altConds[0] = Operational
		opCount := Solutions(altConds, damagedGroups, damagedCount)

		// Damaged Reality
		altConds[0] = Damaged
		altCount := Solutions(altConds, damagedGroups, damagedCount)

		return opCount + altCount
	}

	if c == Operational {
		if damagedCount > 0 {
			if damagedCount != targetDamaged {
				return 0
			} else {
				damagedGroups = damagedGroups[1:]
				damagedCount = 0
			}
		}
	}
	if c == Damaged {
		damagedCount += 1
	}

	return Solutions(conds[1:], damagedGroups, damagedCount)
}

func (r Record) Solutions() int64 {
	return Solutions(r.Conditions, r.DamagedGroups, 0)
}

func main() {
	var sum int64 = 0

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
		solutions := r.Solutions()

		fmt.Fprintf(os.Stdout, "Row %d/%d ->\t%d\n", i+1, len(records), solutions)
		sum += solutions
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
