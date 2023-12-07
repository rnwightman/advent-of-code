package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	TimeInMS     uint
	DistanceInMM uint
}

func (r Race) TestChargeTime(x uint) bool {
	d := (r.TimeInMS - x) * x
	return d > r.DistanceInMM
}

func (r Race) NumberOfSolutions() uint {
	var minSln, maxSln uint
	minVelocity := r.DistanceInMM / r.TimeInMS

	for x := minVelocity; x < r.TimeInMS; x++ {
		ok := r.TestChargeTime(x)
		if ok {
			minSln = x
			break
		}
	}

	for x := r.TimeInMS; x > minSln; x-- {
		ok := r.TestChargeTime(x)
		if ok {
			maxSln = x
			break
		}
	}

	count := maxSln - minSln + 1
	return count
}

func main() {
	races := ParseRaces(os.Stdin)

	var result uint = 1
	for _, race := range races {
		n := race.NumberOfSolutions()
		if n == 0 {
			continue
		}

		result *= n
	}

	fmt.Fprintln(os.Stderr, "Races", races)
	fmt.Fprintln(os.Stdout, result)
}

func ParseRaces(f *os.File) []Race {
	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		panic("Unable to read times")
	}
	rawTimes, ok := strings.CutPrefix(scanner.Text(), "Time:")
	if !ok {
		panic("Unable to parse times")
	}
	times := ParseNumbers(rawTimes)

	if !scanner.Scan() {
		panic("Unable to read distances")
	}
	rawDistances, ok := strings.CutPrefix(scanner.Text(), "Distance:")
	if !ok {
		panic("Unable to parse distances")
	}
	distances := ParseNumbers(rawDistances)

	if len(times) != len(distances) {
		panic("Times and Distances have different count")
	}

	races := make([]Race, len(times))
	for i := 0; i < len(races); i++ {
		time := times[i]
		distance := distances[i]
		race := Race{
			TimeInMS:     time,
			DistanceInMM: distance,
		}

		races[i] = race
	}
	return races
}

func ParseNumbers(s string) []uint {
	records := strings.Fields(s)

	record := strings.Join(records, "")
	n, err := strconv.ParseUint(record, 10, 64)
	if err != nil {
		panic(err)
	}

	return []uint{
		uint(n),
	}
}
