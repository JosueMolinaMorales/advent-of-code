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
	res := partOne(string(input))
	fmt.Println("Part 1:", res)
	res = partTwo(string(input))
	fmt.Println("Part 2:", res)
}

func partTwo(input string) int {
	hands := parseInput(input)
	// Now change the value of J in HandMapping
	HandMapping["J"] = 0
	// Rank Hands
	rankings := createRankings(&hands, 2)
	// Order rankings
	orderRankings(&rankings)
	// Calculate score
	return calculateBidScore(rankings)
}

func partOne(input string) int {
	hands := parseInput(input)
	// Rank hands
	rankings := createRankings(&hands, 1)
	// Order rankings
	orderRankings(&rankings)
	// Calculate score
	return calculateBidScore(rankings)
}

// Function to order rankings based on hand values
func orderRankings(rankings *map[int][]HandBid) {
	for _, v := range *rankings {
		sort.Slice(v, func(i, j int) bool {
			hand := v[i].Hand
			other := v[j].Hand

			idx := 0
			var hv, ov int
			for hv == ov {
				hv = HandMapping[string(hand[idx])]
				ov = HandMapping[string(other[idx])]
				idx += 1
			}

			return hv < ov
		})
	}
}

// Function to calculate and print the final score
func calculateBidScore(rankings map[int][]HandBid) int {
	ans := 0
	currRank := 1
	for _, rank := range []int{HIGH_CARD, ONE_PAIR, TWO_PAIR, THREE_OF_KIND, FULL_HOUSE, FOUR_OF_KIND, FIVE_OF_KIND} {
		hb := rankings[rank]
		for _, h := range hb {
			ans += (currRank * h.Bid)
			currRank++
		}
	}
	return ans
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
	// Calculate the original rank
	rank := handRank(hand)
	// Count the number of J's
	jCount := strings.Count(hand, "J")
	if jCount == 0 {
		return rank
	}
	switch rank {
	case FIVE_OF_KIND:
		// If the rank of the hand is five of a kind, then return
		return FIVE_OF_KIND
	case FOUR_OF_KIND:
		// If the rank of the hand is four of a kind, then the extra J will make it five of a kind
		// AAAAJ -> AAAAA
		return FIVE_OF_KIND
	case FULL_HOUSE:
		// If the rank of the hand is full house, then the extra J will make it five of a kind
		// AAAJJ -> AAAAA
		return FIVE_OF_KIND
	case THREE_OF_KIND:
		// If the rank of the hand is three of a kind, then the extra J will make it four of a kind
		// AAAKJ -> AAAAK
		return FOUR_OF_KIND
	case TWO_PAIR:
		// If the rank of the hand is two pair (AAKKJ or AAJJK), then
		if jCount == 1 {
			// If there is one J, then it will become a full house (AAKKJ --> AAAKK)
			return FULL_HOUSE
		} else {
			// If there are two J's, then it will become four of a kind (AAJJK --> AAAAK)
			return FOUR_OF_KIND
		}
	case ONE_PAIR:
		// If the rank of the hand is one pair (AAKJQ) then the extra J will make it three of a kind
		return THREE_OF_KIND
	case HIGH_CARD:
		// If all cards are different, changing the J to another card will make a one pair
		return ONE_PAIR
	}

	return rank
}

func handRank(hand string) int {
	// Find the number of distinct characters
	characters := make(map[rune]int, 0)
	for _, ch := range hand {
		characters[ch] += 1
	}

	// If there is only 1 distinct character, then it is five of a kind
	if len(characters) == 1 {
		// Five of kind
		return FIVE_OF_KIND
	}

	// Loop through the distinct characters
	for _, v := range characters {
		// If the count of the character is 4, then it is four of a kind
		if v == 4 {
			return FOUR_OF_KIND
		}
		// If the count of the character is 3 and there are only 2 distinct characters, then it is a full house
		if v == 3 && len(characters) == 2 {
			return FULL_HOUSE
		}
		// If the count of the character is 3 and there are more than 2 distinct characters, then it is three of a kind
		if v == 3 {
			return THREE_OF_KIND
		}
	}

	// If there are 3 distinct characters, then it is two pair
	if len(characters) == 3 {
		return TWO_PAIR
	}

	// If there are 4 distinct characters, then it is one pair
	if len(characters) == 4 {
		return ONE_PAIR
	}
	// If there are 5 distinct characters, then it is high card
	return HIGH_CARD
}
