package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	seeds := readSeeds(reader)
	seedToSoil := readMap(reader)
	soilToFertilizer := readMap(reader)
	fertilizerToWater := readMap(reader)
	waterToLight := readMap(reader)
	lightToTemperature := readMap(reader)
	temperatureToHumidity := readMap(reader)
	humidityToLocation := readMap(reader)

	fmt.Fprintln(os.Stderr, "parsed input")

	locations := []int{}
	for _, seed := range seeds {
		soil := seedToSoil.Lookup(seed)
		fertilizer := soilToFertilizer.Lookup(soil)
		water := fertilizerToWater.Lookup(fertilizer)
		light := waterToLight.Lookup(water)
		temp := lightToTemperature.Lookup(light)
		humidity := temperatureToHumidity.Lookup(temp)
		location := humidityToLocation.Lookup(humidity)

		locations = append(locations, location)
	}

	fmt.Fprintln(os.Stderr, "resolved locations")
	fmt.Fprintln(os.Stderr, locations)

	closest := locations[0]
	for _, location := range locations {
		closest = min(closest, location)
	}

	fmt.Fprintln(os.Stdout, closest)
}

func readSeeds(reader *bufio.Reader) []int {
	line, err := Readln(reader)
	if err != nil {
		panic(err)
	}

	rSeeds, ok := strings.CutPrefix(line, "seeds: ")
	if !ok {
		panic("Missing seeds section")
	}

	seedCategories := []int{}
	for _, r := range strings.Split(rSeeds, " ") {
		sc, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}

		seedCategories = append(seedCategories, sc)
	}

	// consume blank link
	Readln(reader)

	return seedCategories
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

type Mapping struct {
	Label string
	m     map[int]int
}

func NewMapping(label string) Mapping {
	return Mapping{
		Label: label,
		m:     make(map[int]int),
	}
}

func (m *Mapping) AddMapping(source, dest, count int) {
	for i := 0; i < count; i++ {
		m.m[source+i] = dest + i
	}
}

func (m Mapping) Lookup(source int) int {
	if dest, ok := m.m[source]; ok {
		return dest
	} else {
		return source
	}
}
