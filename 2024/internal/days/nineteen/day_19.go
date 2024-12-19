package nineteen

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/iterators"
)

func SolveDay19() {
	fmt.Println("Day 19 Part 1: ", solvePartOne())
	fmt.Println("Day 19 Part 2: ", solvePartTwo())
}

func setup() ([]string, []string) {
	input, err := util.LoadFileAsString("./inputs/day_19.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")
	towels := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	return towels, designs
}

func solvePartOne() int {
	towels, designs := setup()

	return len(iterators.Filter(designs, func(design string) bool {
		return makeDesign(design, towels, &map[string]int{}) > 0
	}))
}

func solvePartTwo() int {
	towels, designs := setup()
	count := 0
	for _, design := range designs {
		count += makeDesign(design, towels, &map[string]int{})
	}
	return count
}

func makeDesign(design string, patterns []string, cache *map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	if val, ok := (*cache)[design]; ok {
		return val
	}

	filtered := iterators.Filter(patterns, func(p string) bool {
		return strings.HasPrefix(design, p)
	})

	sum := 0
	for _, f := range filtered {
		sum += makeDesign(strings.TrimPrefix(design, f), patterns, cache)
	}
	(*cache)[design] = sum
	return sum
}
