package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID           int
	RevealedSets []CubeSet
}

func (g Game) MinimumSet() CubeSet {
	result := CubeSet{}

	for _, s := range g.RevealedSets {
		result.Red = max(result.Red, s.Red)
		result.Blue = max(result.Blue, s.Blue)
		result.Green = max(result.Green, s.Green)
	}

	return result
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func (cs *CubeSet) Power() int {
	return max(1, cs.Red) * max(1, cs.Green) * max(1, cs.Blue)
}

func main() {
	games := parseGames(os.Stdin)

	sum := 0
	for _, g := range games {
		minSet := g.MinimumSet()
		power := minSet.Power()

		sum += power
	}

	fmt.Fprintln(os.Stdout, sum)
}

func parseGames(f *os.File) []Game {
	gs := []Game{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		g := parseGame(line)

		fmt.Fprintln(os.Stderr, line, g)

		gs = append(gs, g)
	}

	return gs
}

func parseGame(line string) Game {
	before, setsLine, ok := strings.Cut(line, ": ")
	if !ok {
		panic("Cannot parse line")
	}

	rawId, ok := strings.CutPrefix(before, "Game ")
	if !ok {
		panic("Cannot parse game ID")
	}
	id, _ := strconv.Atoi(rawId)

	return Game{
		ID:           id,
		RevealedSets: parseSets(setsLine),
	}
}

func parseSets(input string) []CubeSet {
	inputs := strings.Split(input, "; ")
	sets := []CubeSet{}

	for _, setInput := range inputs {
		set := parseSet(setInput)
		sets = append(sets, set)
	}

	return sets
}

func parseSet(input string) CubeSet {
	is := strings.Split(input, ", ")

	cs := CubeSet{}
	for _, s := range is {
		rawCount, colour, _ := strings.Cut(s, " ")
		count, _ := strconv.Atoi(rawCount)

		switch colour {
		case "red":
			cs.Red = count

		case "blue":
			cs.Blue = count

		case "green":
			cs.Green = count
		}
	}

	return cs
}
