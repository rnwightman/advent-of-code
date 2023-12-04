package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var lookup map[string]int

func init() {
	lookup = make(map[string]int)

	lookup["one"] = 1
	lookup["two"] = 2
	lookup["three"] = 3
	lookup["four"] = 4
	lookup["five"] = 5
	lookup["six"] = 6
	lookup["seven"] = 7
	lookup["eight"] = 8
	lookup["nine"] = 9

	for i := range [10]int{} {
		word := fmt.Sprint(i)
		lookup[word] = i
	}
}

func main() {
	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		value := parseValue(line)

		result = result + value
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Fprintln(os.Stdout, result)
}

func parseValue(s string) int {
	var digits []int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			if v, ok := lookup[word]; ok {
				digits = append(digits, v)
				break
			}
		}
	}

	d1 := digits[0]
	d2 := digits[len(digits)-1]

	value, _ := strconv.Atoi(fmt.Sprintf("%d%d", d1, d2))

	fmt.Fprintln(os.Stderr, s, digits, value)

	return value
}
