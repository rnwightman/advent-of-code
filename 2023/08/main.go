package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

type Directions string

type Node struct {
	Label string
	Left  string
	Right string
}

type Map []Node

func (m Map) IndexOfLabel(l string) int {
	return slices.IndexFunc(m, func(n Node) bool {
		return l == n.Label
	})
}

func main() {
	dirs, m := ParseInput(os.Stdin)
	steps := Solve(m, dirs)

	fmt.Fprintln(os.Stdout, steps)
}

func Solve(m Map, dirs Directions) int {
	fmt.Fprintln(os.Stderr, "Directions", dirs)
	fmt.Fprintf(os.Stderr, "Map contains %d nodes\n", len(m))

	startIndex := m.IndexOfLabel("AAA")
	exitIndex := m.IndexOfLabel("ZZZ")
	if startIndex == -1 || exitIndex == -1 {
		panic("Unable to locate start and/or exit nodes on map")
	}
	fmt.Fprintln(os.Stderr, "Start", startIndex, "Exit", exitIndex)

	steps := 0
	for i := startIndex; i != exitIndex; {
		d := dirs[steps%len(dirs)]

		cur := m[i]
		var nextLabel string
		switch d {
		case 'L':
			nextLabel = cur.Left
		case 'R':
			nextLabel = cur.Right
		default:
			panic("unexpected direction")
		}

		i = m.IndexOfLabel(nextLabel)
		steps += 1

		fmt.Fprintf(os.Stderr, "Step: %d goes from %s to %s\n", steps, cur.Label, nextLabel)
	}

	return steps
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
