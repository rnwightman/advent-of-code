package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pattern []string

func main() {
	var sum uint64

	patterns := parseInput(os.Stdin)
	for i, p := range patterns {
		answer := analyze(p)
		fmt.Fprintln(os.Stderr, "Pattern", i+1, "=", answer)
		for _, l := range p {
			fmt.Fprintln(os.Stderr, l)
		}
		fmt.Fprintln(os.Stderr)
	}

	fmt.Fprintln(os.Stdout, "Answer:", sum)
}

func analyze(p Pattern) uint64 {
	return 0
}

func parseInput(f *os.File) []Pattern {
	patterns := make([]Pattern, 0)

	var pattern Pattern
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			patterns = append(patterns, pattern)
			pattern = Pattern{}
			continue
		}

		pattern = append(pattern, line)
	}

	patterns = append(patterns, pattern)

	return patterns
}
