package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	Entrance = 'A'
	Exit     = 'Z'
)

type Directions string

type Node struct {
	Label string

	Entrance bool
	Exit     bool

	Left  *Node
	Right *Node
}

func (n Node) String() string {
	return n.Label
}

func main() {
	dirs, startNodes := ParseInput(os.Stdin)
	steps := Solve(startNodes, dirs)

	fmt.Fprintln(os.Stdout, steps)
}

func AllAreExits(nodes []*Node) bool {
	for _, n := range nodes {
		if !n.Exit {
			return false
		}
	}

	return true
}

func Solve(curNodes []*Node, dirs Directions) int {
	fmt.Fprintln(os.Stderr, "Directions", dirs)
	fmt.Fprintln(os.Stderr, "Entrance", len(curNodes), curNodes)

	var step int
	for step = 0; !AllAreExits(curNodes); step++ {
		d := dirs[step%len(dirs)]
		fmt.Fprintf(os.Stderr, "S=%d; D=%c;\t@=%v\n", step, d, curNodes)

		var next *Node
		for i, n := range curNodes {
			switch d {
			case 'L':
				next = n.Left
			case 'R':
				next = n.Right

			default:
				panic("unexpected direction")
			}
			curNodes[i] = next
		}
	}
	fmt.Fprintln(os.Stderr, "Exit", len(curNodes), curNodes)

	return step
}

func ParseInput(f *os.File) (Directions, []*Node) {
	scanner := bufio.NewScanner(f)

	if !scanner.Scan() {
		panic("unable to read diections")
	}
	dirs := Directions(scanner.Text())

	if !scanner.Scan() {
		panic("unexpected end of input")
	}

	records := make([][3]string, 0)
	nodesByLabel := make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()

		label, left, right := ParseRecord(line)

		records = append(records, [3]string{
			label, left, right,
		})

		nodesByLabel[label] = &Node{
			Label:    label,
			Entrance: label[2] == Entrance,
			Exit:     label[2] == Exit,
		}
	}

	for _, r := range records {
		node, _ := nodesByLabel[r[0]]
		node.Left, _ = nodesByLabel[r[1]]
		node.Right, _ = nodesByLabel[r[2]]
	}

	entranceNodes := make([]*Node, 0)
	for _, n := range nodesByLabel {
		if n.Entrance {
			entranceNodes = append(entranceNodes, n)
		}
	}

	return dirs, entranceNodes
}

var nodeRexex = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func ParseRecord(s string) (label, left, right string) {
	matches := nodeRexex.FindStringSubmatch(s)
	return matches[1], matches[2], matches[3]
}
