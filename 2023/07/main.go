package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Hand struct {
	Cards []Card
	Bid   int
}

type HandType int64

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

func (h HandType) String() string {
	switch h {
	case HighCard:
		return "high card"
	case OnePair:
		return "one pair"
	case TwoPair:
		return "two pair"
	case ThreeOfKind:
		return "three of a kind"
	case FullHouse:
		return "full house"
	case FourOfKind:
		return "four of a kind"
	case FiveOfKind:
		return "five of a kind"
	}

	return "unknown"
}

func (h Hand) UniqueCards() []Card {
	inResults := make(map[Card]bool)
	var results []Card

	for _, c := range h.Cards {
		if _, ok := inResults[c]; !ok {
			inResults[c] = true
			results = append(results, c)
		}
	}

	return results
}

func (h Hand) Type() HandType {
	uniqueCards := h.UniqueCards()
	counts := make([]int, len(uniqueCards))

	for i, c := range uniqueCards {
		for _, o := range h.Cards {
			if c != o {
				continue
			}
			counts[i] += 1
		}
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	// Score the hand
	switch {
	case counts[0] == 5:
		return FiveOfKind

	case counts[0] == 4:
		return FourOfKind

	case counts[0] == 3 && counts[1] == 2:
		return FullHouse

	case counts[0] == 3:
		return ThreeOfKind

	case counts[0] == 2 && counts[1] == 2:
		return TwoPair

	case counts[0] == 2 && counts[1] == 1:
		return OnePair

	default:
		return HighCard
	}
}

func (h Hand) Strength() int {
	ht := h.Type()
	strength := int(ht)

	return strength
}

type Card byte

func (c Card) String() string {
	return fmt.Sprintf("%c", c)
}

func (c Card) Strength() int {
	switch byte(c) {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14

	default:
		s, _ := strconv.Atoi(string(c))
		return s
	}
}

func main() {
	hands := ParseHands(os.Stdin)

	for _, h := range hands {
		fmt.Fprintln(os.Stderr, h, "t:", h.Type(), "s:", h.Strength())
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if n := cmp.Compare(a.Strength(), b.Strength()); n != 0 {
			return n
		}

		n := slices.CompareFunc(a.Cards, b.Cards, func(a, b Card) int {
			return cmp.Compare(a.Strength(), b.Strength())
		})
		fmt.Fprintln(os.Stderr, "Secondary comparison", a, b, n)

		return n
	})

	var result int
	for i, h := range hands {
		rank := i + 1
		winnings := rank * h.Bid

		fmt.Fprintf(os.Stderr, "%v; rank=%d; winnings: $%d\n", h, rank, winnings)

		result += winnings
	}

	fmt.Fprintln(os.Stdout, result)

}

func ParseHands(f *os.File) []Hand {
	hands := []Hand{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		b, a, ok := strings.Cut(line, " ")
		if !ok {
			panic("Unable to parse line")
		}
		hand := Hand{
			Cards: ParseCards(b),
			Bid:   ParseBid(a),
		}
		hands = append(hands, hand)
	}
	return hands
}

func ParseCards(s string) []Card {
	cards := make([]Card, 5)
	for i, char := range s {
		cards[i] = Card(char)
	}

	return cards
}

func ParseBid(s string) int {
	b, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(b)
}
