package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
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

var cache = sync.Map{}

func HasVisited(key string) (uint64, bool) {
	if answer, ok := cache.Load(key); ok {
		return answer.(uint64), true
	}

	return 0, false
}

func Visit(key string, answer uint64) uint64 {
	cache.Store(key, answer)
	return answer
}

func Solutions(conds []Condition, damagedGroups []int64, damagedCount int64) uint64 {
	// Evertying is done
	if len(conds) == 0 && len(damagedGroups) == 0 && damagedCount == 0 {
		return 1
	}

	key := fmt.Sprintf("%v::%v::%d", conds, damagedGroups, damagedCount)

	// No more input and one group of damaged
	if len(conds) == 0 {
		if len(damagedGroups) == 1 && damagedCount == damagedGroups[0] {
			return Visit(key, 1)
		} else {
			return Visit(key, 0)
		}
	}

	var targetDamaged int64 = 0
	if len(damagedGroups) > 0 {
		targetDamaged = damagedGroups[0]
	}

	// Too many damaged no further solutions
	if damagedCount > targetDamaged {
		return Visit(key, 0)
	}

	if answer, ok := HasVisited(key); ok {
		return answer
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

		answer := opCount + altCount
		return Visit(key, answer)
	}

	if c == Operational {
		if damagedCount > 0 {
			if damagedCount != targetDamaged {
				return Visit(key, 0)
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

func (r Record) Solutions() uint64 {
	return Solutions(r.Conditions, r.DamagedGroups, 0)
}

func main() {
	var numberOfCopies int8 = 1
	if len(os.Args) > 1 {
		n, err := strconv.ParseInt(os.Args[1], 10, 8)
		if err != nil {
			panic(err)
		}
		numberOfCopies = int8(n)
	}

	records := parseRecords(os.Stdin, numberOfCopies)
	solutions := make([]uint64, len(records))

	numWorkers := runtime.NumCPU()
	progress := sync.Map{}

	var wg sync.WaitGroup
	wg.Add(len(records))

	for wId := 0; wId < numWorkers; wId++ {
		go func(wId int) {
			for i, r := range records {
				if i%numWorkers != wId {
					continue
				}

				numSolutions := r.Solutions()
				solutions[i] = numSolutions
				wg.Done()

				percentComplete := (100 * i) / len(records)
				progress.Store(wId+1, percentComplete)

				// print progress
				progValues := make([]int, 0, numWorkers)
				progSum := 0
				progress.Range(func(_, v interface{}) bool {
					if percent, ok := v.(int); ok {
						progValues = append(progValues, percent)
						progSum += percent
					}
					return true
				})
				slices.Sort(progValues)
				fmt.Fprintln(os.Stderr, time.Now().Format(time.Stamp), "\t", progSum/numWorkers, "\t", progValues)
			}

			progress.Store(wId+1, 100)
		}(wId)
	}

	wg.Wait()

	var sum uint64 = 0
	for _, n := range solutions {
		sum += n
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
