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

	IsEntrance bool
	IsExit     bool
}

func (n Node) String() string {
	return fmt.Sprintf("{ %s (%s, %s) }", n.Label, n.Left, n.Right)
}

func NewNode(label, left, right string) Node {
	return Node{
		Label: label,
		Left:  left,
		Right: right,

		IsEntrance: label[2] == 'A',
		IsExit:     label[2] == 'Z',
	}
}

type Map map[string]Node

func (m Map) FindNode(l string) Node {
	n, ok := m[l]
	if !ok {
		panic(fmt.Sprintf("Cannot locate node with label %s", l))
	}
	return n
}

func main() {
	dirs, m := ParseInput(os.Stdin)
	steps := Solve(m, dirs)

	fmt.Fprintln(os.Stdout, steps)
}

func AllAreExits(nodes []Node) bool {
	for _, n := range nodes {
		if !n.IsExit {
			return false
		}
	}

	return true
}

func Solve(m Map, dirs Directions) int {
	fmt.Fprintln(os.Stderr, "Directions", dirs)
	fmt.Fprintf(os.Stderr, "Map contains %d nodes\n", len(m))

	var curNodes []Node
	for _, node := range m {
		if node.IsEntrance {
			curNodes = append(curNodes, node)
		}
	}
	fmt.Fprintln(os.Stderr, "Entrance", len(curNodes), curNodes)

	var step int
	for step = 0; !AllAreExits(curNodes); step++ {
		d := dirs[step%len(dirs)]
		fmt.Fprintf(os.Stderr, "S=%d; D=%c;\t@=%v\n", step, d, curNodes)

		var nextLabel string
		for i, n := range curNodes {
			switch d {
			case 'L':
				nextLabel = n.Left
			case 'R':
				nextLabel = n.Right

			default:
				panic("unexpected direction")
			}
			curNodes[i] = m.FindNode(nextLabel)
		}
	}
	fmt.Fprintln(os.Stderr, "Exit", len(curNodes), curNodes)

	return step
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

	var m Map = make(Map, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n := ParseNode(line)
		m[n.Label] = n
	}

	return dirs, m
}

var nodeRexex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func ParseNode(s string) Node {
	matches := nodeRexex.FindStringSubmatch(s)
	return NewNode(matches[1], matches[2], matches[3])
}
