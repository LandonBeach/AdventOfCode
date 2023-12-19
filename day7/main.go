package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
	Type  string
}

// Determine if the hand is one of the following:
// - five of a kind
//   - where all five cards have the same label: `AAAAA`
//
// - four of a kind
//   - where four cards have the same label and one card has a different label: `AA8AA`
//
// - full house
//   - where three cards have the same label, and the remaining two cards share a different label: `23332`
//
// - three of a kind
//   - where three cards have the same label, and the remaining two cards are each different from any other card in the hand: `TTT98`
//
// - two pair
//   - where two cards share one label, two other cards share a second label, and the remaining card has a third label: `23432`
//
// - one pair
//   - where two cards share one label, and the other three cards have a different label from the pair and each other: `A23A4`
//
// - high card
//   - where all cards' labels are distinct: `23456`
func (ch *Hand) findHandType() {
	cards := make(map[string]int)

	// Count all cards in the hand
	for _, label := range ch.Cards {
		cards[string(label)]++
	}

	// The number of unique cards in the hand
	switch len(cards) {
	case 1:
		ch.Type = "five of a kind"
	case 2:
		for _, occurances := range cards {
			if occurances == 4 {
				ch.Type = "four of a kind"
				break
			} else {
				ch.Type = "full house"
			}
		}
	case 3:
		for _, occurances := range cards {
			if occurances == 3 {
				ch.Type = "three of a kind"
				break
			} else {
				ch.Type = "two pair"
			}
		}
	case 4:
		ch.Type = "one pair"
	case 5:
		ch.Type = "high card"
	}
}

// Initialize a new instance of a Hand
func NewHand(hand string, bid string) *Hand {
	ch := Hand{
		Cards: hand,
		Bid:   utils.StringToInt(bid),
	}

	ch.findHandType()
	return &ch
}

func main() {
	file := utils.OpenFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make(map[string][]Hand)

	// Get all hands from the input file and categorize them by type (i.e. "full house", "high card", etc.)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		h := NewHand(line[0], line[1])
		hands[h.Type] = append(hands[h.Type], *h)
	}

	// Sort the hands in ascending order within each category
	for _, hand := range hands {
		sort.Slice(hand, func(i, j int) bool {
			first, second := hand[i].Cards, hand[j].Cards

			// Define the order of face cards ('A' > 'K' > 'Q' > 'J' > 'T')
			order := map[byte]int{'A': 5, 'K': 4, 'Q': 3, 'J': 2, 'T': 1}

			for k := range first {
				firstCard, secondCard := first[k], second[k]

				// Compare cards based on the defined face card order
				if order[firstCard] != order[secondCard] {
					return order[firstCard] < order[secondCard]
				}

				// If the face values are the same, compare based on the individual card values
				if firstCard != secondCard {
					return firstCard < secondCard
				}
			}

			// If all cards are the same up to this point, return false (no need to swap)
			return false
		})
	}

	rank := 1
	part1Total := 0
	order := []string{"high card", "one pair", "two pair", "three of a kind", "full house", "four of a kind", "five of a kind"}

	for _, handType := range order {
		for _, hand := range hands[handType] {
			part1Total += (hand.Bid * rank)
			rank++
		}
	}

	fmt.Printf("Part1 total: %d\n", part1Total)

}
