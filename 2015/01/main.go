package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	floor, pos := DecodeInstructions(text)

	fmt.Printf("Floor: %d\tPosition: %d\n", floor, pos)
}

func DecodeInstructions(input string) (int, int) {
	var floor, pos int

	for i, c := range input {
		switch c {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		}
		if pos == 0 && floor < 0 {
			pos = i + 1
		}
	}
	return floor, pos
}
