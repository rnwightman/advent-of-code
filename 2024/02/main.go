package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Level int
type Report []Level

func (r Report) IsSafe() bool {
	deltas := make([]int, 0, len(r)-1)
	for i := range len(r) - 1 {
		a := r[i]
		b := r[i+1]

		delta := int(b - a)

		if delta > 3 || delta < -3 || delta == 0 {
			return false
		}

		deltas = append(deltas, delta)
	}

	maxDelta := slices.Max(deltas)
	minDelta := slices.Min(deltas)
	if maxDelta > 0 && minDelta < 0 {
		return false
	}
	return true
}

func (r Report) GenerateDampendReports() []Report {
	variants := make([]Report, 0)
	for i := range r {
		variant := slices.Delete(slices.Clone(r), i, i+1)
		variants = append(variants, variant)
	}

	return variants
}

func (r Report) IsSafeWithDampener() bool {
	variants := r.GenerateDampendReports()
	// fmt.Println("variants of", r, variants)
	for _, variant := range variants {
		ok := variant.IsSafe()
		// fmt.Println("variant", variant, "ok?", ok)

		if ok {
			return true
		}
	}

	return false
}

func main() {
	reports := ReadReports()

	fmt.Println("reports", reports)
	safeReports := make([]Report, 0)
	for _, report := range reports {
		ok := report.IsSafe()

		if ok || report.IsSafeWithDampener() {
			safeReports = append(safeReports, report)
		}
	}

	fmt.Println("safe reports", safeReports)
	fmt.Println("number of safe reports", len(safeReports))
}

func ReadReports() []Report {
	reports := make([]Report, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		r := make(Report, 0)
		for _, field := range strings.Fields(line) {
			value, err := strconv.Atoi(field)
			if err != nil {
				log.Panic("unable to parse level", err)
			}

			level := Level(value)
			r = append(r, level)
		}

		reports = append(reports, r)
	}

	return reports
}
