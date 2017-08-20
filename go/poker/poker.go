package poker

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

const testVersion = 5

type Card struct {
	rank string
	suit string
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.rank, c.suit)
}

type Hand [5]Card

func (h *Hand) String() string {
	return strings.Join([]string{h[0].String(),
		h[1].String(),
		h[2].String(),
		h[3].String(),
		h[4].String()}, " ")
}

var validSuit = map[string]bool{"♤": true, "♡": true, "♢": true, "♧": true}

var rankValue = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

func NewCard(input string) (Card, error) {
	s := input[len(input)-1:]
	r := input[:len(input)-1]

	if _, ok := validSuit[s]; !ok {
		return Card{}, errors.New("invalid suit")
	}
	if _, ok := rankValue[r]; !ok {
		return Card{}, errors.New("invalid rank")
	}

	return Card{rank: r, suit: s}, nil
}

func NewHand(input string) (Hand, error) {
	hand := Hand{}
	inputCards := strings.Split(input, " ")
	if len(inputCards) != 5 {
		return Hand{}, errors.New("invalid input format, did't get 5 cards")
	}
	for i, card := range inputCards {
		c, err := NewCard(card)
		if err != nil {
			return Hand{}, err
		}
		hand[i] = c
	}
	return hand, nil
}

func BestHand(inputHands []string) ([]string, error) {
	hands := make([]Hand, len(inputHands))
	for i := range inputHands {
		hand, err := NewHand(inputHands[i])
		if err != nil {
			return nil, err
		}
	}

	sort.Slice(hands, lessHand(hands))
	return []string{hands[0].String()}, nil
}

func lessHand(hands []Hand) func(i, j int) bool {

	// Sorting with the less func make the lowest elements at the
	// beginning, so the winning hand is less than the losing hand
	return func(i, j int) bool {
		tests := []HandRankingTest{}
		for _, t := range tests {
			if t(hands[i]) && t(hands[j]) {
				// Depending on the function, need to break either by the highest pair or
				// by the kicker.
			}
			if t(hands[i]) {
				return true
			}
			if t(hands[j]) {
				return false
			}
		}
		// If no test are successful, break tie with high card.

	}

}

type HandRankingTest func(h Hand) bool

func hasRoyalFlush(h Hand) bool {
	var rankSet = map[string]bool{
		"10": true,
		"J":  true,
		"Q":  true,
		"K":  true,
		"A":  true,
	}

	handSet := make(map[string]bool)
	suit := h[0].suit
	for _, c := range h {
		if c.suit != suit {
			return false
		}
		handSet[c.rank] = true
	}
	return reflect.DeepEqual(handSet, rankSet)
}

func hasStraightFlush(h Hand) bool {
	return hasStraight(h) && hasFlush(h)
}

func hasFourOfAKind(h Hand) bool {

}

func hasFullHouse(h Hand) bool {

}

func hasFlush(h Hand) bool {

}

func hasStraight(h Hand) bool {
}
func hasThreeOfAKind(h Hand) bool {

}

func hasTwoPairs(h Hand) bool {

}

func hasOnePair(h Hand) bool {

}

func highCard(h Hand) Card {

}
