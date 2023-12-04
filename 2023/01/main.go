package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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

	fmt.Fprintln(os.Stdout, "Sum of Calibration Values: ", result)
}

func parseValue(s string) int {
	var digits []int
	for _, value := range s {
		digit := 0
		if unicode.IsDigit(value) {
			digit = int(value - '0')
		}

		if digit > 0 {
			digits = append(digits, digit)
		}
	}

	d1 := digits[0]
	d2 := digits[len(digits)-1]

	value, _ := strconv.Atoi(fmt.Sprintf("%d%d", d1, d2))

	return value
}
