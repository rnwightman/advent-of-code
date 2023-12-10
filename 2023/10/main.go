package main

import (
	"bufio"
	"fmt"
	"os"

	"slices"
)

type TileType rune

const (
	Ground         TileType = '.'
	VerticalPipe   TileType = '|'
	HorizontalPipe TileType = '-'
	BendNorthEast  TileType = 'L'
	BendNorthWest  TileType = 'J'
	BendSouthWest  TileType = '7'
	BendSouthEast  TileType = 'F'
	StartingPos    TileType = 'S'
)

func (t TileType) String() string {
	return fmt.Sprintf("%c", t)
}

type Coordinate struct {
	Row    int
	Column int
}

func (c Coordinate) North() Coordinate {
	return Coordinate{
		Row:    c.Row - 1,
		Column: c.Column,
	}
}

func (c Coordinate) South() Coordinate {
	return Coordinate{
		Row:    c.Row + 1,
		Column: c.Column,
	}
}

func (c Coordinate) East() Coordinate {
	return Coordinate{
		Row:    c.Row,
		Column: c.Column + 1,
	}
}

func (c Coordinate) West() Coordinate {
	return Coordinate{
		Row:    c.Row,
		Column: c.Column - 1,
	}
}

type Tile struct {
	Type     TileType
	Position Coordinate
}

func (t Tile) IsStartingTile() bool {
	return t.Type == StartingPos
}

func (t Tile) AdjacentPositions() []Coordinate {
	switch t.Type {
	case VerticalPipe:
		return []Coordinate{t.Position.North(), t.Position.South()}

	case HorizontalPipe:
		return []Coordinate{t.Position.East(), t.Position.West()}

	case BendNorthEast:
		return []Coordinate{t.Position.North(), t.Position.East()}
	case BendNorthWest:
		return []Coordinate{t.Position.North(), t.Position.West()}
	case BendSouthEast:
		return []Coordinate{t.Position.South(), t.Position.East()}
	case BendSouthWest:
		return []Coordinate{t.Position.South(), t.Position.West()}

	case StartingPos:
		return []Coordinate{
			t.Position.North(),
			t.Position.East(),
			t.Position.South(),
			t.Position.West(),
		}

	default:
		return []Coordinate{}
	}
}

type Board struct {
	Tiles [][]Tile
}

func (b Board) StartingTile() (Tile, bool) {
	for _, row := range b.Tiles {
		for _, t := range row {
			if t.IsStartingTile() {
				return t, true
			}
		}
	}

	return Tile{}, false
}

func (b Board) TileAt(c Coordinate) (Tile, bool) {
	if c.Row < 0 || c.Row >= len(b.Tiles) {
		return Tile{}, false
	}

	row := b.Tiles[c.Row]
	if c.Column < 0 || c.Column >= len(row) {
		return Tile{}, false
	}

	tile := row[c.Column]
	return tile, true
}

func (b Board) TileConnectsTo(c Tile) []Tile {
	tiles := make([]Tile, 0)

	for _, p := range c.AdjacentPositions() {
		t, ok := b.TileAt(p)
		if !ok {
			continue
		}

		mutuallyAdjacent := slices.Contains(t.AdjacentPositions(), c.Position)
		if mutuallyAdjacent {
			tiles = append(tiles, t)
		}
	}

	return tiles
}

func (b Board) NextTile(p, c Tile) Tile {
	for _, aPos := range c.AdjacentPositions() {
		if aPos == p.Position {
			continue
		}

		t, ok := b.TileAt(aPos)
		if !ok {
			panic("walked off the board")
		}
		return t
	}

	panic("unable to find nex tile")
}

func (b Board) Explore(p, c Tile) (int, bool) {
	if c.Type == StartingPos {
		return 1, true
	}

	n := b.NextTile(p, c)
	s, ok := b.Explore(c, n)
	if !ok {
		return 0, false
	}

	return s + 1, true
}

func main() {
	input := readInput(os.Stdin)
	board := parseBoard(input)

	start, found := board.StartingTile()
	if !found {
		panic("could not locate starting position")
	}
	fmt.Fprintln(os.Stderr, "Starting Pos", start)

	adjToStart := board.TileConnectsTo(start)
	fmt.Fprintln(os.Stderr, "Adj to Start", adjToStart)
	if len(adjToStart) != 2 {
		panic("unexpected number of candidate loops")
	}

	steps, ok := board.Explore(start, adjToStart[0])
	if !ok {
		panic("unable to explore the loop")
	}

	fmt.Fprintln(os.Stderr, "Length", steps)

	dist := steps / 2
	fmt.Fprintln(os.Stdout, dist)
}

func parseBoard(inputTiles [][]TileType) Board {
	nRows := len(inputTiles)
	tiles := make([][]Tile, nRows)
	for row := 0; row < nRows; row++ {
		rInput := inputTiles[row]

		nCols := len(rInput)
		rTiles := make([]Tile, nCols)
		for col := 0; col < nCols; col++ {
			tt := rInput[col]
			rTiles[col] = Tile{
				Type:     tt,
				Position: Coordinate{Row: row, Column: col},
			}
		}

		tiles[row] = rTiles
	}
	return Board{
		Tiles: tiles,
	}
}

func readInput(f *os.File) [][]TileType {
	tiles := make([][]TileType, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rTiles := make([]TileType, len(line))

		for i, c := range line {
			rTiles[i] = TileType(c)
		}
		tiles = append(tiles, rTiles)
	}

	return tiles
}
