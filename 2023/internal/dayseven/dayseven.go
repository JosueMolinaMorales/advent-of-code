package dayseven

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const (
	// where all cards' labels are distinct
	HIGH_CARD = iota
	// where two cards share one label, and the other three cards have a different label from the pair and each other
	ONE_PAIR
	// where two cards share one label, two other cards share a second label, and the remaining card has a third label
	TWO_PAIR
	// where three cards have the same label, and the remaining two cards are each different from any other card in the hand
	THREE_OF_KIND
	// where three cards have the same label, and the remaining two cards share a different label
	FULL_HOUSE
	// where four cards have the same label and one card has a different label
	FOUR_OF_KIND
	// where all five cards have the same label
	FIVE_OF_KIND
)

var HandMapping map[string]int = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"M": 0, // For part 2
}

type HandBid struct {
	Hand string
	Rank int
	Bid  int
}

func RunDaySeven() {
	input, err := os.ReadFile("./input/day7input.txt")
	if err != nil {
		panic("Cannot read day 7 file")
	}
	partTwo(string(input))
	// partTwo(testInput)
}

func partTwo(input string) {
	hands := parseInput(strings.ReplaceAll(input, "J", "M"))
	rankings := createRankings(&hands, 2)
	// Order rankings
	for _, v := range rankings {
		// fmt.Println("Before, ", v, k)
		sort.Slice(v, func(i, j int) bool {
			hand := v[i].Hand
			other := v[j].Hand

			idx := 0
			hv := HandMapping[string(hand[idx])]
			// fmt.Println(string(hand[idx]), hv)
			ov := HandMapping[string(other[idx])]
			// fmt.Println(string(other[idx]), hv)

			for hv == ov {
				idx += 1
				hv = HandMapping[string(hand[idx])]
				// fmt.Println(string(hand[idx]), hv)
				ov = HandMapping[string(other[idx])]
				// fmt.Println(string(other[idx]), ov)
			}

			return hv < ov
		})

		// fmt.Println("After, ", v)
	}

	ans := 0
	currRank := 1
	for _, rank := range []int{HIGH_CARD, ONE_PAIR, TWO_PAIR, THREE_OF_KIND, FULL_HOUSE, FOUR_OF_KIND, FIVE_OF_KIND} {
		hb := rankings[rank]
		for _, h := range hb {
			// fmt.Printf("%s: %d * %d\n", h.Hand, currRank, h.Bid)
			ans += (currRank * h.Bid)
			currRank += 1
		}
	}

	fmt.Println(ans)
}

func partOne(input string) {
	hands := parseInput(input)
	// Rank hands
	rankings := createRankings(&hands, 1)
	// Order rankings
	for _, v := range rankings {
		// fmt.Println("Before, ", v)
		sort.Slice(v, func(i, j int) bool {
			hand := v[i].Hand
			other := v[j].Hand

			idx := 0
			hv := HandMapping[string(hand[idx])]
			// fmt.Println(string(hand[idx]), hv)
			ov := HandMapping[string(other[idx])]
			// fmt.Println(string(other[idx]), hv)

			for hv == ov {
				idx += 1
				hv = HandMapping[string(hand[idx])]
				// fmt.Println(string(hand[idx]), hv)
				ov = HandMapping[string(other[idx])]
				// fmt.Println(string(other[idx]), ov)
			}

			return hv < ov
		})

		// fmt.Println("After, ", v)
	}

	ans := 0
	currRank := 1
	for _, rank := range []int{HIGH_CARD, ONE_PAIR, TWO_PAIR, THREE_OF_KIND, FULL_HOUSE, FOUR_OF_KIND, FIVE_OF_KIND} {
		hb := rankings[rank]
		for _, h := range hb {
			// fmt.Printf("%s: %d * %d\n", h.Hand, currRank, h.Bid)
			ans += (currRank * h.Bid)
			currRank += 1
		}
	}

	fmt.Println(ans)
}

func createRankings(hands *[]HandBid, part int) map[int][]HandBid {
	rankings := make(map[int][]HandBid, 0)
	for _, hb := range *hands {
		var rank int
		if part == 1 {
			rank = handRank(hb.Hand)
		} else {
			rank = handRankPartTwo(hb.Hand)
		}
		rankings[rank] = append(rankings[rank], hb)
	}
	return rankings
}

func parseInput(input string) []HandBid {
	hands := make([]HandBid, 0)
	for _, line := range strings.Split(input, "\n") {
		str := strings.Split(line, " ")
		hand := str[0]
		bid, err := strconv.Atoi(str[1])
		if err != nil {
			fmt.Println("Failed to convert", str[1], "to int")
		}

		hands = append(hands, HandBid{
			Hand: hand,
			Rank: 0,
			Bid:  bid,
		})
	}
	return hands
}

func handRankPartTwo(hand string) int {
	rank := handRank(hand)
	jCount := strings.Count(hand, "M")
	if jCount == 0 {
		return rank
	}
	switch rank {
	case FIVE_OF_KIND:
		return FIVE_OF_KIND
	case FOUR_OF_KIND:
		return FIVE_OF_KIND
	case FULL_HOUSE:
		return FIVE_OF_KIND
	case THREE_OF_KIND:
		return FOUR_OF_KIND
	case TWO_PAIR:
		if jCount == 1 {
			return FULL_HOUSE
		} else {
			return FOUR_OF_KIND
		}
	case ONE_PAIR:
		if jCount == 2 {
			return THREE_OF_KIND
		}
		if jCount == 1 {
			return THREE_OF_KIND
		}
	case HIGH_CARD:
		return ONE_PAIR
	}

	return rank
}

func handRank(hand string) int {
	characters := make(map[rune]int, 0)
	for _, ch := range hand {
		characters[ch] += 1
	}

	if len(characters) == 1 {
		// Five of kind
		return FIVE_OF_KIND
	}

	for _, v := range characters {
		if v == 4 {
			return FOUR_OF_KIND
		}
		if v == 3 && len(characters) == 2 {
			return FULL_HOUSE
		}
		if v == 3 {
			return THREE_OF_KIND
		}
	}

	if len(characters) == 3 {
		return TWO_PAIR
	}

	if len(characters) == 4 {
		return ONE_PAIR
	}

	return HIGH_CARD
}
