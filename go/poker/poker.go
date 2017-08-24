package poker

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode/utf8"
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

	rankLen := 1
	r := input[:rankLen]
	if r == "1" {
		rankLen = 2
		r = input[:rankLen]
	}
	s, _ := utf8.DecodeLastRuneInString(input)
	if utf8.RuneCountInString(input) != len(r)+1 {
		return Card{}, errors.New("unexpected card format")
	}

	if _, ok := validSuit[string(s)]; !ok {
		return Card{}, errors.New("invalid suit")
	}
	if _, ok := rankValue[r]; !ok {
		return Card{}, errors.New("invalid rank")
	}

	return Card{rank: r, suit: string(s)}, nil
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
		hands[i] = hand
	}

	sort.Slice(hands, lessHand(hands))
	if len(hands) > 1 && isTie(hands[0], hands[1]) {
		return []string{hands[0].String(), hands[1].String()}, nil
	}
	return []string{hands[0].String()}, nil
}

func lessHand(hands []Hand) func(i, j int) bool {

	// Sorting with the less func make the lowest elements at the
	// beginning, so the winning hand is less than the losing hand
	// In short, if hand i win, must return true
	return func(i, j int) bool {
		tests := []HandRankingTest{
			hasRoyalFlush,
			hasStraightFlush,
			hasFourOfAKind,
			hasFullHouse,
			hasFlush,
			hasStraight,
			hasThreeOfAKind,
			hasTwoPairs,
			hasOnePair,
		}

		handValueTests := make([]HandValueTest, len(tests))
		handValueTests[3] = getFullHouseHandValue
		handValueTests[4] = getFlushHandValue
		handValueTests[5] = getStraightHandValue
		handValueTests[6] = getThreeOfAKindHandValue
		handValueTests[7] = getTwoPairsHandValue
		handValueTests[8] = getOnePairHandValue

		for k, t := range tests {
			if t(hands[i]) && t(hands[j]) {
				// Depending on the function, get the associated function
				// to calculate the winning hand
				if a := handValueTests[k]; a != nil {
					return a(hands[i]) > a(hands[j])
				}
			}
			if t(hands[i]) {
				return true
			}
			if t(hands[j]) {
				return false
			}
		}
		// If no test are successful, break tie with high card.
		RankI := rankValue[highCard(hands[i]).rank]
		rankJ := rankValue[highCard(hands[j]).rank]
		return RankI > rankJ

	}

}

type HandRankingTest func(h Hand) bool
type HandValueTest func(h Hand) int

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
	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	for _, count := range cardKind {
		if count == 4 {
			return true
		}
	}
	return false
}

func hasFullHouse(h Hand) bool {
	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	var hasThree, hasTwo bool
	for _, count := range cardKind {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func getFullHouseHandValue(h Hand) (value int) {
	return 10*getThreeOfAKindHandValue(h) + getOnePairHandValue(h)
}

func hasFlush(h Hand) bool {
	suit := h[0].suit
	for _, c := range h {
		if c.suit != suit {
			return false
		}
	}
	return true
}

func getFlushHandValue(h Hand) (value int) {
	for _, c := range h {
		value += (1 << uint(rankValue[c.rank]))
	}
	return
}

func hasStraight(h Hand) bool {
	// I don't want to sort them so here is a cheap hack
	cardKind := make(map[int]int)
	smallestRank := 15
	for _, card := range h {
		cardKind[rankValue[card.rank]]++
		if cardKind[rankValue[card.rank]] > 1 {
			return false
		}
		if rankValue[card.rank] < smallestRank {
			smallestRank = rankValue[card.rank]
		}
		if card.rank == "A" {
			smallestRank = 1
			cardKind[1]++
		}
	}
	for i := smallestRank; i < smallestRank+5; i++ {
		if _, ok := cardKind[i]; !ok {
			// Not sequential
			// If we have an ace, check for the end too!
			for j := rankValue["A"]; j > rankValue["A"]-5; j-- {
				if _, ok := cardKind[j]; !ok {
					return false
				}
			}
		}
	}

	return true
}

func getStraightHandValue(h Hand) (value int) {
	for _, c := range h {
		if c.rank == "A" {
			continue
		}
		value += (1 << uint(rankValue[c.rank]))
	}
	return
}

func hasThreeOfAKind(h Hand) bool {
	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	for _, count := range cardKind {
		if count == 3 {
			return true
		}
	}
	return false

}

func getThreeOfAKindHandValue(h Hand) (value int) {
	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	for rank, count := range cardKind {
		if count == 3 {
			value += 100 * rankValue[rank]
		} else {
			value += rankValue[rank]
		}

	}
	return
}

func hasTwoPairs(h Hand) bool {
	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	var firstPair bool
	for _, count := range cardKind {
		if count == 2 && firstPair {
			return true
		}
		if count == 2 {
			firstPair = true
		}
	}
	return false

}

func getTwoPairsHandValue(h Hand) (value int) {

	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	var firstPair bool
	for rank, count := range cardKind {
		if count == 2 && firstPair {
			value += (1 << uint(rankValue[rank]))
		}
		if count == 2 {
			value += (1 << uint(rankValue[rank]))
		}
	}
	return
}

func hasOnePair(h Hand) bool {
	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	for _, count := range cardKind {
		if count == 2 {
			return true
		}
	}
	return false

}

func getOnePairHandValue(h Hand) (value int) {

	cardKind := make(map[string]int)
	for _, card := range h {
		cardKind[card.rank]++
	}
	for rank, count := range cardKind {
		if count == 2 {
			value += rankValue[rank]
		}
	}
	return
}

func highCard(h Hand) Card {

	highPos := 0
	for pos, card := range h {
		if rankValue[card.rank] > rankValue[h[highPos].rank] {
			highPos = pos
		}

	}
	return h[highPos]
}

func isTie(i, j Hand) bool {
	// Is a tie if both hands have the same ranks cards
	RankMapI := make(map[string]int)
	RankMapJ := make(map[string]int)

	for _, c := range i {
		RankMapI[c.rank]++
	}
	for _, c := range j {
		RankMapJ[c.rank]++
	}

	return reflect.DeepEqual(RankMapI, RankMapJ)
}
