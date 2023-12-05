package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	partNumbers := []int{}

	lines := readLines(os.Stdin)
	for lineIndex := range lines {
		var pLine, nLine string
		if lineIndex-1 >= 0 {
			pLine = lines[lineIndex-1]
		}
		if lineIndex+1 < len(lines) {
			nLine = lines[lineIndex+1]
		}

		line := lines[lineIndex]

		partNumbersInLine := identifyPartNumbers(line, pLine, nLine)

		fmt.Fprintln(os.Stderr, " ", pLine)
		fmt.Fprintln(os.Stderr, ">", line)
		fmt.Fprintln(os.Stderr, " ", nLine)
		fmt.Fprintln(os.Stderr, "#", partNumbersInLine)
		fmt.Fprintln(os.Stderr)

		for _, partNumber := range partNumbersInLine {
			partNumbers = append(partNumbers, partNumber)
		}
	}

	sum := 0
	for _, partNumber := range partNumbers {
		sum = sum + partNumber
	}

	fmt.Fprintln(os.Stdout, sum)
}

var numberRegex = regexp.MustCompile(`\d+`)

func identifyPartNumbers(line, prevLine, nextLine string) []int {
	partNumbers := []int{}

	indexOfNumbers := numberRegex.FindAllIndex([]byte(line), -1)
	for _, indices := range indexOfNumbers {
		partNumber, _ := strconv.Atoi(line[indices[0]:indices[1]])
		partNumbers = append(partNumbers, partNumber)
	}

	return partNumbers
}

func readLines(f *os.File) []string {
	lines := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines

}
