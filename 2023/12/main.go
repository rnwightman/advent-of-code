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

func main() {
	fmt.Fprintln(os.Stderr, "WIP: Day 12")

	records := parseRecords(os.Stdin)
	for _, r := range records {
		fmt.Fprintln(os.Stderr, r)
	}
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
