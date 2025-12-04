package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day2() {
	fmt.Println("2025 Day 2 Part 1:", day2Part1("inputs/day_2/input.txt"))
	fmt.Println("2025 Day 2 Part 2:", day2Part2("inputs/day_2/input.txt"))
}

func day2Part1(path string) int {
	input, err := io.ReadFileAsString(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
	}

	ranges := strings.Split(input, ",")
	// Find invalid IDs where the number splits into two identical halves
	// Examples: 22, 1111, 1212, 4545
	sum := 0

	for _, r := range ranges {
		start, end := parseRange(r)

		for i := start; i <= end; i++ {
			if hasRepeatingHalves(i) {
				sum += i
			}
		}
	}

	return sum
}

// parseRange parses a range string like "10-20" into start and end integers
func parseRange(rangeStr string) (int, int) {
	parts := strings.Split(rangeStr, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 2: invalid range start: %s", err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 2: invalid range end: %s", err)
	}
	return start, end
}

// hasRepeatingHalves checks if a number's string representation
// can be split into two identical halves
func hasRepeatingHalves(num int) bool {
	numStr := strconv.Itoa(num)
	length := len(numStr)

	// Can only split evenly if length is even
	if length%2 != 0 {
		return false
	}

	mid := length / 2
	return numStr[:mid] == numStr[mid:]
}

func day2Part2(path string) int {
	input, err := io.ReadFileAsString(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 2 Part 2: %s", err)
	}

	ranges := strings.Split(input, ",")
	// Find invalid IDs where a sequence repeats at least twice
	// Examples: 11, 1212, 135135
	invalidIDs := make(map[int]bool)

	for _, r := range ranges {
		start, end := parseRange(r)

		for i := start; i <= end; i++ {
			if hasRepeatingPattern(i) {
				invalidIDs[i] = true
			}
		}
	}

	sum := 0
	for num := range invalidIDs {
		sum += num
	}

	return sum
}

// hasRepeatingPattern checks if a number consists of a repeating pattern
// For example: 11 (1 repeats), 1212 (12 repeats), 135135 (135 repeats)
func hasRepeatingPattern(num int) bool {
	numStr := strconv.Itoa(num)
	length := len(numStr)

	tried := make(map[int]bool)

	// Try different pattern lengths (from 1 to half the total length)
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		gcf := util.GCF(length, patternLen)

		// Skip if we've already tried this GCF
		if tried[gcf] {
			continue
		}
		tried[gcf] = true

		// Check if the pattern repeats to form the entire number
		if length%gcf == 0 {
			pattern := numStr[:gcf]
			repeats := length / gcf

			// Build expected string by repeating the pattern
			expected := strings.Repeat(pattern, repeats)
			if expected == numStr {
				return true
			}
		}
	}

	return false
}
