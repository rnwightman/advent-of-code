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
	digits := ""
	for _, value := range s {
		if unicode.IsDigit(value) {
			digits = fmt.Sprintf("%s%c", digits, value)
		}
	}

	d1 := digits[0]
	d2 := digits[len(digits)-1]

	digits = fmt.Sprintf("%c%c", d1, d2)

	value, _ := strconv.Atoi(digits)
	return value
}
