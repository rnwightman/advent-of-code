package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Directions string

type Node struct {
	Label string
	Left  string
	Right string
}

type Map []Node

func main() {
	dirs, m := ParseInput(os.Stdin)

	fmt.Fprintln(os.Stderr, "Directions", dirs)
	fmt.Fprintf(os.Stderr, "Map contains %d nodes\n", len(m))
}

func ParseInput(f *os.File) (Directions, Map) {
	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		panic("unable to read diections")
	}
	dirs := Directions(scanner.Text())

	if !scanner.Scan() {
		panic("unexpected end of input")
	}

	var m Map
	for scanner.Scan() {
		line := scanner.Text()
		n := ParseNode(line)
		m = append(m, n)
	}

	return dirs, m
}

var nodeRexex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func ParseNode(s string) Node {
	matches := nodeRexex.FindStringSubmatch(s)
	return Node{
		Label: matches[1],
		Left:  matches[2],
		Right: matches[3],
	}
}
