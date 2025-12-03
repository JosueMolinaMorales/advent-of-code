package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day3() {
	fmt.Println("2025 Day 3 Part 1:", day3Part1("inputs/day_3/input.txt"))
	fmt.Println("2025 Day 3 Part 2:", day3Part2("inputs/day_3/input.txt"))
}

func day3Part1(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 3 Part 1: %s", err)
	}

	batteries := parseInputToBatteryBanks(input)

	result := 0
	for _, bank := range batteries {
		largestTwo := findLargestNDigits(bank, 2)
		result += largestTwo
	}
	return result
}

// parseInputToBatteryBanks converts input lines to a 2D slice of integers
func parseInputToBatteryBanks(lines []string) [][]int {
	batteries := make([][]int, 0, len(lines))
	for _, line := range lines {
		bank := make([]int, 0, len(line))
		for _, char := range strings.Split(line, "") {
			digit, err := strconv.Atoi(char)
			if err != nil {
				log.Fatalf("ERROR: Failed to parse digit: %s", err)
			}
			bank = append(bank, digit)
		}
		batteries = append(batteries, bank)
	}
	return batteries
}

func day3Part2(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 3 Part 2: %s", err)
	}

	batteries := parseInputToBatteryBanks(input)

	result := 0
	for _, bank := range batteries {
		largestNumber := findLargestNDigits(bank, 12)
		result += largestNumber
	}
	return result
}

// findLargestNDigits finds the largest N-digit number from left to right
// where each digit must appear in order in the original array
func findLargestNDigits(bank []int, n int) int {
	largest := make([]int, n)
	lastPos := -1

	// For each position in our result, find the largest available digit
	for position := 0; position < n; position++ {
		digitsRemaining := n - position

		// Search for the largest digit that still leaves room for remaining digits
		for j := lastPos + 1; j <= len(bank)-digitsRemaining; j++ {
			if bank[j] > largest[position] {
				largest[position] = bank[j]
				lastPos = j
			}
		}
	}

	return convertDigitsToNumber(largest)
}

// convertDigitsToNumber converts a slice of digits into a single integer
func convertDigitsToNumber(digits []int) int {
	var builder strings.Builder
	for _, digit := range digits {
		builder.WriteString(strconv.Itoa(digit))
	}

	number, err := strconv.Atoi(builder.String())
	if err != nil {
		log.Fatalf("ERROR: Failed to convert digits to number: %s", err)
	}
	return number
}
