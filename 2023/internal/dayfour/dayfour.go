package dayfour

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Winning []int
	Numbers []int
}

const testInput string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func RunDayFour() {
	input, err := os.ReadFile("./input/day4input.txt")
	if err != nil {
		panic("Could not read file for day 4")
	}
	res := partOne(string(input))
	fmt.Println("Part 1:", res)
	res = partTwo(string(input))
	fmt.Println("Part 2:", res)
}

func partOne(input string) int {
	cards := parseInput(input)
	sum := 0
	for _, card := range cards {
		count := 0
		points := 0
		for _, wn := range card.Winning {
			if slices.Contains(card.Numbers, wn) {
				points = int(math.Pow(float64(2), float64(count)))
				count++
			}
		}
		sum += points
	}

	return sum
}

func partTwo(input string) int {
	cards := parseInput(input)
	// Stores the number of wins for every game
	wins := make(map[int]int)
	// Stores the number of times each card is copied
	instances := make(map[int]int)
	for i, card := range cards {
		count := 0
		for _, win := range card.Winning {
			if slices.Contains(card.Numbers, win) {
				count++
			}
		}

		wins[i+1] = count
		instances[i+1] = 1
	}

	// Create queue
	queue := make([]int, 0)
	for k := range wins {
		queue = append(queue, k)
	}

	// Loop through queue, adding up instances
	for len(queue) != 0 {
		card := queue[0]
		// Get the number of wins for that card
		w := wins[card]
		// Add up the instances
		for _, c := range makeRange(card+1, card+w) {
			instances[c] += 1
			// Add to queue
			queue = append(queue, c)
		}
		queue = queue[1:]
	}

	// Get sum
	sum := 0
	for _, v := range instances {
		sum += v
	}

	return sum
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func parseInput(input string) []Card {
	cards := make([]Card, 0)
	for i, line := range strings.Split(input, "\n") {
		line = strings.ReplaceAll(line, fmt.Sprintf("Card %d: ", i+1), "")
		numbers := strings.Split(line, " | ")
		winning, nums := make([]int, 0), make([]int, 0)
		for _, win := range strings.Split(numbers[0], " ") {
			winNum, err := strconv.Atoi(strings.TrimSpace(win))
			if err != nil {
				// fmt.Printf("Could not convert %s to int\n", win)
				continue
			}
			winning = append(winning, winNum)
		}

		for _, n := range strings.Split(numbers[1], " ") {
			num, err := strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				// fmt.Printf("Could not convert %s to int\n", n)
				continue
			}

			nums = append(nums, num)
		}
		cards = append(cards, Card{Winning: winning, Numbers: nums})
	}

	return cards
}
