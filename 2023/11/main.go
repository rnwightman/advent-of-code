package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	Row int
	Col int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.Row, c.Col)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (c Coordinate) DistanceTo(o Coordinate) int {
	cols := c.Col - o.Col
	rows := c.Row - o.Row

	return Abs(cols) + Abs(rows)
}

type Galaxy struct {
	ID       int
	Position Coordinate
}

func (g Galaxy) String() string {
	return fmt.Sprintf("%d@(%v)", g.ID, g.Position)
}

func main() {
	galaxies := readInput(os.Stdin)
	// TODO expand universe
	distances := calculateDistances(galaxies)

	fmt.Fprintln(os.Stderr, "Number of distances", len(distances))
	sumOfDistances := 0
	for _, dist := range distances {
		sumOfDistances += dist
	}

	fmt.Fprintln(os.Stdout, sumOfDistances)
}

func calculateDistances(galaxies []Galaxy) []int {
	ds := make([]int, 0)
	for x, g1 := range galaxies {
		for y, g2 := range galaxies {
			if y <= x {
				continue
			}

			d := g1.Position.DistanceTo(g2.Position)
			ds = append(ds, d)

			fmt.Fprintf(os.Stderr, "From %v\tTo %v\t %d\n", g1, g2, d)
		}
	}

	return ds
}

func readInput(f *os.File) []Galaxy {
	input := make([]Galaxy, 0)

	id := 0
	var row, col int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row += 1

		col = 0
		line := scanner.Text()
		for _, symbol := range line {
			col += 1

			if symbol == '#' {
				id += 1
				p := Coordinate{
					Row: row,
					Col: col,
				}
				g := Galaxy{
					ID:       id,
					Position: p,
				}

				input = append(input, g)
			}
		}
	}

	return input
}
