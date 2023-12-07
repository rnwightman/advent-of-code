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

func (r Race) Solutions() []uint {
	// v := velocity
	// d := distance
	// t := time
	// t =

	slns := []uint{}
	for x := uint(1); x < r.TimeInMS; x++ {
		d := (r.TimeInMS - x) * x
		if d <= r.DistanceInMM {
			continue
		}

		slns = append(slns, x)
	}

	fmt.Fprintln(os.Stderr, "Race", r)
	fmt.Fprintln(os.Stderr, "Slns", len(slns), slns)

	return slns
}

func main() {
	races := ParseRaces(os.Stdin)

	var result uint = 1
	for _, race := range races {
		solutions := race.Solutions()
		numSolutions := uint(len(solutions))
		if numSolutions == 0 {
			continue
		}

		result *= numSolutions
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

	ns := make([]uint, len(records))
	for i := 0; i < len(records); i++ {
		n, err := strconv.ParseUint(records[i], 10, 64)
		if err != nil {
			panic(err)
		}
		ns[i] = uint(n)
	}

	return ns
}
