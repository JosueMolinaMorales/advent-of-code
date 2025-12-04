package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day1() {
	fmt.Println("2025 Day 1 Part 1:", day1Part1("inputs/day_1/input.txt"))
	fmt.Println("2025 Day 1 Part 2:", day1Part2("inputs/day_1/input.txt"))
}

func day1Part1(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 1 Part 1: %s", err)
	}

	// Circular dial (0-99) starting at position 50
	// Spin left (L) or right (R) by given distance
	// Count the number of times the dial lands on 0
	const dialSize = 100
	const startPos = 50

	pos := startPos
	count := 0

	for _, action := range input {
		direction, distance := parseAction(action)
		pos = (pos + direction*distance)
		pos = util.EuclideanMod(pos, dialSize)

		if pos == 0 {
			count++
		}
	}

	return count
}

// parseAction parses an action string like "L25" or "R10" into direction and distance
// Returns -1 for left, +1 for right, and the numeric distance
func parseAction(action string) (int, int) {
	dir := action[0]
	distance, err := strconv.Atoi(action[1:])
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 1: invalid action: %s", err)
	}

	if dir == 'L' {
		return -1, distance
	}
	return 1, distance
}

func day1Part2(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 1 Part 2: %s", err)
	}

	// Count every time the dial passes through or lands on 0
	// Moving step-by-step to catch all passes
	const dialSize = 100
	const startPos = 50

	pos := startPos
	count := 0

	for _, action := range input {
		direction, distance := parseAction(action)

		// Move one step at a time to count all passes through 0
		for range distance {
			pos = (pos + direction)
			pos = util.EuclideanMod(pos, dialSize)

			if pos == 0 {
				count++
			}
		}
	}

	return count
}
