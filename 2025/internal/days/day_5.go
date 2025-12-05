package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day5() {
	fmt.Println("2025 Day 5 Part 1:", day5Part1("inputs/day_5/input.txt"))
	fmt.Println("2025 Day 5 Part 2:", day5Part2("inputs/day_5/input.txt"))
}

func day5Part1(path string) int {
	ranges, ids, err := parseDatabase(path)
	if err != nil {
		log.Fatalf("ERROR 2025 Failed to Parse the Input Data: %s", err)
	}

	// Count how many ingredient IDs are fresh (fall within at least one range)
	validIds := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				validIds++
				break
			}
		}
	}
	return validIds
}

func day5Part2(path string) int {
	ranges, _, err := parseDatabase(path)
	if err != nil {
		log.Fatalf("ERROR 2025 Failed to Parse the Input Data: %s", err)
	}

	// Merge overlapping ranges, then count total IDs
	sortRanges(ranges)
	merged := mergeRanges(ranges)

	totalCount := 0
	for _, r := range merged {
		totalCount += r[1] - r[0] + 1
	}

	return totalCount
}

// sortRanges sorts ranges by their start position
func sortRanges(ranges [][]int) {
	n := len(ranges)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ranges[j][0] > ranges[j+1][0] {
				ranges[j], ranges[j+1] = ranges[j+1], ranges[j]
			}
		}
	}
}

// mergeRanges merges overlapping or adjacent ranges
func mergeRanges(ranges [][]int) [][]int {
	if len(ranges) == 0 {
		return [][]int{}
	}

	merged := [][]int{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		currentRange := ranges[i]
		lastMerged := merged[len(merged)-1]

		// Check if current range overlaps or is adjacent to the last merged range
		if currentRange[0] <= lastMerged[1]+1 {
			// Merge by extending the upper bound if needed
			if currentRange[1] > lastMerged[1] {
				lastMerged[1] = currentRange[1]
			}
		} else {
			// No overlap, add as new range
			merged = append(merged, currentRange)
		}
	}

	return merged
}

func parseDatabase(path string) ([][]int, []int, error) {
	input, err := io.ReadFileAsString(path)
	if err != nil {
		return nil, nil, err
	}

	parts := strings.Split(input, "\n\n")

	// Parse ranges (format: "lower-upper")
	ranges := make([][]int, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		r := strings.Split(line, "-")
		lower, err := strconv.Atoi(r[0])
		if err != nil {
			return nil, nil, err
		}
		upper, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, nil, err
		}
		ranges = append(ranges, []int{lower, upper})
	}

	// Parse ingredient IDs
	ids := make([]int, 0)
	for _, line := range strings.Split(parts[1], "\n") {
		id, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, id)
	}

	return ranges, ids, nil
}
