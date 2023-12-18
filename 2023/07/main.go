package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("filename is empty")
		return
	}
	filename := os.Args[1]

	lines, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Total Winnings:", totalWinnings(lines, false))
	fmt.Println("Total Winnings (Jokers):", totalWinnings(lines, true))
}

func totalWinnings(lines []string, specialJoker bool) int {
	sum := 0

	handWithBids := []handWithBid{}
	for _, line := range lines {
		handWithBids = append(handWithBids, parse(line))
	}
	var sorted []handWithBid
	if specialJoker {
		sorted = sortByRank(handWithBids, specialJoker)
	} else {
		sorted = sortByRank(handWithBids, specialJoker)
	}

	for i, handWithBid := range sorted {
		sum += int(handWithBid.bid) * (i + 1)
	}

	return sum
}

func sortByRank(handWithBids []handWithBid, specialJoker bool) []handWithBid {
	sort.Slice(handWithBids, func(i, j int) bool {
		handTypeI := handWithBids[i].hand.handType()
		handTypeJ := handWithBids[j].hand.handType()

		if specialJoker {
			handTypeI = handWithBids[i].hand.handTypeSpecialJoker()
			handTypeJ = handWithBids[j].hand.handTypeSpecialJoker()
		}

		if handTypeI == handTypeJ {
			// If hand types are equal, compare the cards in the hands
			for k := 0; k < len(handWithBids[i].hand.cards); k++ {
				cardI := cardRank(handWithBids[i].hand.cards[k], specialJoker)
				cardJ := cardRank(handWithBids[j].hand.cards[k], specialJoker)

				if cardI != cardJ {
					return cardI < cardJ
				}
			}
		}

		return handTypeI < handTypeJ
	})

	return handWithBids
}

func cardRank(card string, specialJoker bool) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if specialJoker {
			return 1
		}
		return 11
	case "T":
		return 10
	default:
		v, _ := strconv.Atoi(card)
		return v
	}
}

type hand struct {
	cards []string
}

type bid int

type handWithBid struct {
	hand hand
	bid  bid
}

type handType int

const (
	HighCard handType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h hand) handType() handType {
	cardCounts := map[string]int{}

	for _, card := range h.cards {
		cardCounts[card]++
	}

	pairs := 0
	threes := 0
	fours := 0

	for _, count := range cardCounts {
		switch count {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		case 5:
			return FiveOfAKind
		}
	}

	if fours > 0 {
		return FourOfAKind
	}

	if threes > 0 && pairs > 0 {
		return FullHouse
	}

	if threes > 0 {
		return ThreeOfAKind
	}

	if pairs == 2 {
		return TwoPairs
	}

	if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func (h hand) handTypeSpecialJoker() handType {
	cardCounts := map[string]int{}

	for _, card := range h.cards {
		cardCounts[card]++
	}

	pairs := 0
	threes := 0
	fours := 0
	jokers := 0

	for card, count := range cardCounts {
		if card == "J" {
			jokers = count
			if jokers == 5 {
				return FiveOfAKind
			}
			continue
		}
		switch count {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		case 5:
			return FiveOfAKind
		}
	}

	if (fours == 1 && jokers == 1) || (threes == 1 && jokers == 2) || (pairs == 1 && jokers == 3) || (jokers >= 4) {
		return FiveOfAKind
	}
	if (fours == 1) || (threes == 1 && jokers == 1) || (pairs == 1 && jokers == 2) || (jokers == 3) {
		return FourOfAKind
	}
	if (threes == 1 && pairs == 1) || (pairs == 2 && jokers == 1) || (pairs == 1 && jokers == 2) {
		return FullHouse
	}
	if (threes == 1) || (pairs == 1 && jokers == 1) || (jokers == 2) {
		return ThreeOfAKind
	}
	if (pairs == 2) || (pairs == 1 && jokers == 1) {
		return TwoPairs
	}
	if (pairs == 1) || (jokers == 1) {
		return OnePair
	}
	return HighCard

}

func parse(line string) handWithBid {
	parts := strings.Split(line, " ")

	hand := hand{}
	for _, card := range parts[0] {
		hand.cards = append(hand.cards, string(card))
	}
	bid_value, _ := strconv.Atoi(parts[1])

	return handWithBid{hand, bid(bid_value)}
}

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return strings.Split(string(f), "\n"), nil
}
