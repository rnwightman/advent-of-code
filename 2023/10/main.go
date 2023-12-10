package main

import (
	"bufio"
	"fmt"
	"os"
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

type Coordinate struct {
	Row    int
	Column int
}

type Tile struct {
	Type     TileType
	Position Coordinate
}

func (t Tile) IsStartingTile() bool {
	return t.Type == StartingPos
}

type Board struct {
	Tiles [][]Tile
}

func (b Board) StartingTile() (*Tile, bool) {
	for _, row := range b.Tiles {
		for i := range row {
			t := &row[i]

			if t.IsStartingTile() {
				return t, true
			}
		}
	}

	return nil, false
}

func (b Board) TileAt(c Coordinate) (*Tile, bool) {
	if c.Row < 0 || c.Row >= len(b.Tiles) {
		return nil, false
	}

	row := b.Tiles[c.Row]
	if c.Column < 0 || c.Column >= len(row) {
		return nil, false
	}

	tile := &row[c.Column]
	return tile, true
}

func (t TileType) String() string {
	return fmt.Sprintf("%c", t)
}

func main() {
	input := readInput(os.Stdin)
	fmt.Fprintln(os.Stderr, "Input", input)

	board := parseBoard(input)
	fmt.Fprintln(os.Stderr, "Board", board)

	start, found := board.StartingTile()
	if !found {
		panic("could not locate starting position")
	}
	fmt.Fprintln(os.Stderr, "Starting Pos", start)

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
