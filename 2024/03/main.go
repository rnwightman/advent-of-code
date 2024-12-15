package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	X int
	Y int
}

func (i Instruction) Result() int {
	return i.X * i.Y
}

func ParseInstructions(r *os.File) []Instruction {
	instructions := make([]Instruction, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		instructionsInLine := parseInstructions(line)

		instructions = append(instructions, instructionsInLine...)
	}

	return instructions
}

var instructionRegex *regexp.Regexp = regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)

func parseInstructions(input string) []Instruction {
	instructions := make([]Instruction, 0)

	matches := instructionRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		fmt.Println("values in", match)

		x, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("could not parse X from", match, err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal("could not parse Y from", match, err)
		}

		instruction := Instruction{
			X: x,
			Y: y,
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func main() {
	instructions := ParseInstructions(os.Stdin)
	fmt.Println("Instructions =", instructions)

	sum := 0
	for _, instruction := range instructions {
		sum = sum + instruction.Result()
	}

	fmt.Println("Sum of Products =", sum)
}
