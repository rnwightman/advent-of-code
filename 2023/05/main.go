package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	seeds, seedCount := readSeeds(reader)
	seedToSoil := readMap(reader)
	soilToFertilizer := readMap(reader)
	fertilizerToWater := readMap(reader)
	waterToLight := readMap(reader)
	lightToTemperature := readMap(reader)
	temperatureToHumidity := readMap(reader)
	humidityToLocation := readMap(reader)

	fmt.Fprintln(os.Stderr, "parsed input")

	seedsProcessed := 0
	nearestLocation := math.MaxInt
	fmt.Fprintln(os.Stdout, "Seed Ranges", len(seeds))

	for _, r := range seeds {
		for i := 0; i < r.Length; i++ {
			seed := r.Start + i

			soil := seedToSoil.Lookup(seed)
			fertilizer := soilToFertilizer.Lookup(soil)
			water := fertilizerToWater.Lookup(fertilizer)
			light := waterToLight.Lookup(water)
			temp := lightToTemperature.Lookup(light)
			humidity := temperatureToHumidity.Lookup(temp)
			location := humidityToLocation.Lookup(humidity)

			nearestLocation = min(nearestLocation, location)
		}

		seedsProcessed += r.Length
		percentComplete := int(100 * float64(seedsProcessed) / float64(seedCount))
		fmt.Fprintf(os.Stderr, "progress %d of %d: %d\n", seedsProcessed, seedCount, percentComplete)
	}

	fmt.Fprintln(os.Stdout, nearestLocation)
}

func readSeeds(reader *bufio.Reader) ([]Range, int) {
	line, err := Readln(reader)
	if err != nil {
		panic(err)
	}

	rSeeds, ok := strings.CutPrefix(line, "seeds: ")
	if !ok {
		panic("Missing seeds section")
	}

	numbers := []int{}
	for _, r := range strings.Split(rSeeds, " ") {
		n, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}
	// consume blank link
	Readln(reader)

	seedCount := 0

	seedCategories := make([]Range, len(numbers)/2)
	for i := 0; i < len(numbers); i += 2 {
		baseCategory := numbers[i]
		length := numbers[i+1]

		seedCategories[i/2] = Range{
			Start:  baseCategory,
			Length: length,
		}
		seedCount += length
	}

	return seedCategories, seedCount
}

func readMap(reader *bufio.Reader) Mapping {
	line, err := Readln(reader)
	if err != nil {
		panic(err)
	}

	label, ok := strings.CutSuffix(line, " map:")
	if !ok {
		panic("Missing mapping header")
	}
	mapping := NewMapping(label)
	for {
		line, err := Readln(reader)
		if err != nil {
			panic(err)
		}

		if line == "" {
			break
		}

		values := strings.Split(line, " ")
		if len(values) != 3 {
			panic(fmt.Sprintf("expected length 3, got %d (%v)", len(values), values))
		}
		destRangeStart, _ := strconv.Atoi(values[0])
		sourceRangeStart, _ := strconv.Atoi(values[1])
		rangeLength, _ := strconv.Atoi(values[2])

		mapping.AddMapping(sourceRangeStart, destRangeStart, rangeLength)
	}

	return mapping
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

type Range struct {
	Start  int
	Length int
}

type Mapping struct {
	Label  string
	Ranges []RangeMapping
}

type RangeMapping struct {
	SourceRangeStart int
	DestRangeStart   int
	RangeLength      int
}

func (m RangeMapping) Lookup(source int) (int, bool) {
	delta := source - m.SourceRangeStart
	if delta < 0 || delta > m.RangeLength {
		return 0, false
	}

	return m.DestRangeStart + delta, true
}

func NewMapping(label string) Mapping {
	return Mapping{
		Label: label,
	}
}

func (m *Mapping) AddMapping(source, dest, length int) {
	m.Ranges = append(m.Ranges, RangeMapping{
		SourceRangeStart: source,
		DestRangeStart:   dest,
		RangeLength:      length,
	})
}

func (m Mapping) Lookup(source int) int {
	for _, r := range m.Ranges {
		if dest, ok := r.Lookup(source); ok {
			return dest
		}
	}

	return source
}
