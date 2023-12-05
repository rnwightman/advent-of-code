package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type PartNumber struct {
	Value      int
	StartIndex int
	EndIndex   int
}

type Gear struct {
	Index int
}

type Line struct {
	PartNumbers []PartNumber
	Gears       []Gear
}

func (l *Line) FindPartsIntersecting(index int) []PartNumber {
	matches := []PartNumber{}

	for _, pn := range l.PartNumbers {
		ok := pn.Intersects(index-1) || pn.Intersects(index) || pn.Intersects(index+1)
		if ok {
			matches = append(matches, pn)
		}
	}

	return matches
}

func (pn *PartNumber) Intersects(index int) bool {
	return index >= pn.StartIndex && index < pn.EndIndex
}

func ParsePartNumber(s string, start, end int) PartNumber {
	ss := s[start:end]
	value, _ := strconv.Atoi(ss)

	return PartNumber{
		Value:      value,
		StartIndex: start,
		EndIndex:   end,
	}
}

func main() {
	sum := 0

	lines := readLines(os.Stdin)
	for lineIndex := range lines {
		var pLine, nLine Line
		if lineIndex-1 >= 0 {
			pLine = lines[lineIndex-1]
		}
		if lineIndex+1 < len(lines) {
			nLine = lines[lineIndex+1]
		}

		line := lines[lineIndex]

		ratios := identifyGearRatios(line, pLine, nLine)

		fmt.Fprintln(os.Stderr, " ", pLine)
		fmt.Fprintln(os.Stderr, ">", line)
		fmt.Fprintln(os.Stderr, " ", nLine)
		fmt.Fprintln(os.Stderr, "#", ratios)
		fmt.Fprintln(os.Stderr)

		for _, ratio := range ratios {
			sum += ratio
		}
	}

	fmt.Fprintln(os.Stdout, sum)
}

func identifyGearRatios(line, pLine, nLine Line) []int {
	ratios := []int{}

	for _, gear := range line.Gears {
		adjParts := []PartNumber{}

		adjParts = append(adjParts, pLine.FindPartsIntersecting(gear.Index)...)
		adjParts = append(adjParts, line.FindPartsIntersecting(gear.Index)...)
		adjParts = append(adjParts, nLine.FindPartsIntersecting(gear.Index)...)

		fmt.Fprintln(os.Stderr, "Adjacent part-numbers", adjParts)

		// gear is a gear iff adjacent to exactly two part numbers
		if len(adjParts) == 2 {
			ratio := adjParts[0].Value * adjParts[1].Value
			ratios = append(ratios, ratio)
		} else {
			fmt.Fprintf(os.Stderr, "not gear @ %d: #parts %d\n", gear.Index, len(adjParts))
		}
	}

	return ratios
}

func readLines(f *os.File) []Line {
	lines := []Line{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		line := Line{}

		// identify gears
		for _, indices := range gearRegex.FindAllStringIndex(text, -1) {
			line.Gears = append(line.Gears, Gear{
				Index: indices[0],
			})
		}

		// identify part-numbers
		for _, indicies := range numberRegex.FindAllStringIndex(text, -1) {
			pn := ParsePartNumber(text, indicies[0], indicies[1])
			line.PartNumbers = append(line.PartNumbers, pn)
		}

		lines = append(lines, line)
	}

	return lines
}

var numberRegex = regexp.MustCompile(`\d+`)
var gearRegex = regexp.MustCompile(`\*`)
