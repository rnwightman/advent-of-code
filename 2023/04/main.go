package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Card struct {
	Label          string
	WinningNumbers []uint
	Numbers        []uint
}

func (c Card) Score() uint {
	var pts uint
	for _, n := range c.Numbers {
		ok := slices.Contains(c.WinningNumbers, n)
		if !ok {
			continue
		}

		pts = max(1, pts*2)
	}

	return pts
}

func parseCard(s string) Card {
	label, after, _ := strings.Cut(s, ": ")
	sWinning, sNumbers, _ := strings.Cut(after, " | ")

	return Card{
		Label:          label,
		WinningNumbers: parseNumbers(sWinning),
		Numbers:        parseNumbers(sNumbers),
	}
}

func parseNumbers(s string) []uint {
	ns := []uint{}
	for _, r := range strings.Split(s, " ") {
		n, err := strconv.ParseUint(r, 10, 64)
		if err != nil {
			continue
		}

		ns = append(ns, uint(n))
	}
	return ns
}

func main() {
	var sum uint

	cards := readCards(os.Stdin)

	for _, card := range cards {
		points := card.Score()
		fmt.Fprintln(os.Stdout, card, points)

		sum += points
	}

	fmt.Fprintln(os.Stdout, sum)
}

func readCards(f *os.File) []Card {
	cards := []Card{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		card := parseCard(text)

		cards = append(cards, card)
	}

	return cards
}
