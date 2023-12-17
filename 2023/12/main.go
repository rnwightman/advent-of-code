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

func ToDamagedGroups(conds []Condition) []int64 {
	groups := make([]int64, 0)

	var count int64 = 0
	for _, c := range conds {
		if c == Damaged {
			count += 1
		} else if count > 0 {
			groups = append(groups, count)
			count = 0
		}
	}

	if count > 0 {
		groups = append(groups, count)
	}

	return groups
}

func (r Record) Matches(conds []Condition) bool {
	groups := ToDamagedGroups(conds)
	if len(groups) != len(r.DamagedGroups) {
		return false
	}

	for i := 0; i < len(groups); i++ {
		if groups[i] != r.DamagedGroups[i] {
			return false
		}
	}

	return true
}

func Expand(conds []Condition) [][]Condition {
	if len(conds) == 0 {
		return [][]Condition{}
	}

	expanded := make([][]Condition, 0)
	prefix := make([]Condition, 0)
	for i, c := range conds {
		if c != Unknown {
			prefix = append(prefix, c)
			continue
		}

		remainder := conds[i+1:]
		expRemainder := Expand(remainder)

		if len(expRemainder) == 0 {
			// good
			expandedGood := make([]Condition, 0, len(conds))
			expandedGood = append(expandedGood, prefix...)
			expandedGood = append(expandedGood, Operational)

			// bad
			expandedBad := make([]Condition, 0, len(conds))
			expandedBad = append(expandedBad, prefix...)
			expandedBad = append(expandedBad, Damaged)

			expanded = append(expanded, expandedGood)
			expanded = append(expanded, expandedBad)
		}

		for _, rem := range expRemainder {
			// good
			expandedGood := make([]Condition, 0, len(conds))
			expandedGood = append(expandedGood, prefix...)
			expandedGood = append(expandedGood, Operational)
			expandedGood = append(expandedGood, rem...)

			// bad
			expandedBad := make([]Condition, 0, len(conds))
			expandedBad = append(expandedBad, prefix...)
			expandedBad = append(expandedBad, Damaged)
			expandedBad = append(expandedBad, rem...)

			expanded = append(expanded, expandedGood)
			expanded = append(expanded, expandedBad)
		}
		break
	}

	if len(expanded) == 0 {
		expanded = append(expanded, prefix)
	}

	return expanded
}

func (r Record) PossibleArrangements() [][]Condition {
	return Expand(r.Conditions)
}

func main() {
	sum := 0

	records := parseRecords(os.Stdin)
	for _, r := range records {
		fmt.Fprintln(os.Stderr, r)

		arrangements := r.PossibleArrangements()
		for _, a := range arrangements {
			if len(a) != len(r.Conditions) {
				panic("Arrangement has different length than record")
			}
			ok := r.Matches(a)

			if ok {
				fmt.Fprintln(os.Stderr, "\t", a)
				sum += 1
			}
		}

		fmt.Fprintln(os.Stderr)
	}

	fmt.Fprintln(os.Stdout, sum)
}

func parseRecords(f *os.File) []Record {
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

		rs = append(rs, r)
	}

	return rs
}
