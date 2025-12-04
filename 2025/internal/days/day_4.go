package days

import (
	"fmt"
	"log"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

const (
	roll        = "@"
	emptyCell   = "."
	minAdjacent = 4
)

func Day4() {
	fmt.Println("2025 Day 4 Part 1:", day4Part1("inputs/day_4/input.txt"))
	fmt.Println("2025 Day 4 Part 2:", day4Part2("inputs/day_4/input.txt"))
}

func day4Part1(path string) int {
	arrMap, err := io.ReadFileAs2DArray(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 4 Part 1: %s", err)
	}

	// Count rolls (@) that have fewer than 4 adjacent rolls
	count := 0
	for x, row := range arrMap {
		for y, cell := range row {
			if cell == roll && countAdjacentRolls(arrMap, x, y) < minAdjacent {
				count++
			}
		}
	}

	return count
}

// countAdjacentRolls counts how many rolls (@) are adjacent to the given position
// Checks all 8 directions: up, down, left, right, and diagonals
func countAdjacentRolls(grid [][]string, x, y int) int {
	directions := [][]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	count := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]

		// Check bounds
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			if grid[nx][ny] == roll {
				count++
			}
		}
	}

	return count
}

func day4Part2(path string) int {
	arrMap, err := io.ReadFileAs2DArray(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 4 Part 2: %s", err)
	}

	// Repeatedly remove rolls with fewer than 4 adjacent rolls
	// until no more rolls can be removed
	totalRemoved := 0

	for {
		rollsToRemove := findRollsToRemove(arrMap)

		if len(rollsToRemove) == 0 {
			break
		}

		totalRemoved += len(rollsToRemove)
		removeRolls(arrMap, rollsToRemove)
	}

	return totalRemoved
}

// findRollsToRemove identifies all rolls that have fewer than 4 adjacent rolls
func findRollsToRemove(grid [][]string) [][]int {
	toRemove := [][]int{}

	for x, row := range grid {
		for y, cell := range row {
			if cell == roll && countAdjacentRolls(grid, x, y) < minAdjacent {
				toRemove = append(toRemove, []int{x, y})
			}
		}
	}

	return toRemove
}

// removeRolls replaces the specified positions with empty cells
func removeRolls(grid [][]string, positions [][]int) {
	for _, pos := range positions {
		grid[pos[0]][pos[1]] = emptyCell
	}
}
