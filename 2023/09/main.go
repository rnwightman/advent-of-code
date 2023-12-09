package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	histories := parseInput(os.Stdin)
	nextValues := make([]int, len(histories))

	fmt.Fprintln(os.Stderr, "Input", histories)

	for i := range histories {
		v := extrapolateValue(histories[i])
		nextValues[i] = v
	}

	fmt.Fprintln(os.Stderr, "Next", nextValues)

	sum := 0
	for _, v := range nextValues {
		sum += v
	}

	fmt.Fprintln(os.Stdout, sum)
}

func extrapolateValue(r []int) int {
	differences := make([]int, len(r)-1)

	ok := true
	for i := range differences {
		d := r[i+1] - r[i]
		differences[i] = d

		ok = ok && (d == 0)
	}

	first := r[0]
	if ok {
		return first
	} else {
		return first - extrapolateValue(differences)
	}
}

func parseInput(f *os.File) [][]int {
	histories := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rawReadings := strings.Fields(scanner.Text())
		readings := make([]int, len(rawReadings))
		for i, r := range rawReadings {
			v, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}

			readings[i] = v
		}
		histories = append(histories, readings)
	}

	return histories
}
