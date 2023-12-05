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
	Copies         uint
	WinningNumbers []uint
	Numbers        []uint
}

func (c Card) NumberOfMatches() int {
	var count int
	for _, n := range c.Numbers {
		ok := slices.Contains(c.WinningNumbers, n)
		if !ok {
			continue
		}

		count += 1
	}

	return count
}

func parseCard(s string) Card {
	label, after, _ := strings.Cut(s, ": ")
	sWinning, sNumbers, _ := strings.Cut(after, " | ")

	return Card{
		Label:          label,
		Copies:         1,
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

	for i := range cards {
		card := cards[i]
		n := card.NumberOfMatches()
		fmt.Fprintf(os.Stderr, "%v has %d matches\n", card, n)

		for j := i + 1; j <= i+n; j++ {
			cards[j].Copies += card.Copies
		}

		sum += card.Copies
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
