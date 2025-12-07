package days

import (
	"fmt"
	"log"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/types"
)

func Day7() {
	fmt.Println("2025 Day 7 Part 1:", day7Part1("inputs/day_7/input.txt"))
	fmt.Println("2025 Day 7 Part 2:", day7Part2("inputs/day_7/input.txt"))
}

func day7Part1(path string) int {
	input, err := io.ReadFileAs2DArray(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 7 Part 1: %s", err)
	}

	startingPos, splitters := parseInput(input, false)
	seen := map[types.Point]bool{}
	runBeams(startingPos, splitters, len(input[0]), len(input), seen)

	// Count the splitters that were used
	res := 0
	for _, used := range splitters {
		if used {
			res++
		}
	}
	return res
}

func day7Part2(path string) int {
	input, err := io.ReadFileAs2DArray(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 7 Part 2: %s", err)
	}

	startingPos, splitters := parseInput(input, true)
	splitterList := getSortedSplitters(splitters)
	splitterIndex := createSplitterIndex(splitterList)

	memo := make(map[string]int)
	return countTimelinesOptimized(startingPos, 0, splitterList, splitterIndex, len(input[0]), len(input), memo)
}

func parseInput(input [][]string, initSplittersTrue bool) (types.Point, map[types.Point]bool) {
	startingPos := types.NewPoint(0, 0)
	splitters := map[types.Point]bool{}

	for i, row := range input {
		for j, col := range row {
			if col == "S" {
				startingPos = types.NewPoint(i, j)
			}
			if col == "^" {
				splitters[types.NewPoint(i, j)] = initSplittersTrue
			}
		}
	}

	return startingPos, splitters
}

func getSortedSplitters(splitters map[types.Point]bool) []types.Point {
	splitterList := make([]types.Point, 0, len(splitters))
	for splitter := range splitters {
		splitterList = append(splitterList, splitter)
	}

	// Sort by row (top to bottom)
	for i := 0; i < len(splitterList); i++ {
		for j := i + 1; j < len(splitterList); j++ {
			if splitterList[i].Row > splitterList[j].Row {
				splitterList[i], splitterList[j] = splitterList[j], splitterList[i]
			}
		}
	}

	return splitterList
}

func createSplitterIndex(splitterList []types.Point) map[types.Point]int {
	splitterIndex := make(map[types.Point]int, len(splitterList))
	for i, s := range splitterList {
		splitterIndex[s] = i
	}
	return splitterIndex
}

func countTimelinesOptimized(pos types.Point, splitterMask int, splitterList []types.Point,
	splitterIndex map[types.Point]int, maxCol, maxRow int, memo map[string]int,
) int {
	// Base case: out of bounds = 1 timeline completed
	if pos.Col < 0 || pos.Row < 0 || pos.Col >= maxCol || pos.Row >= maxRow {
		return 1
	}

	// Check memoization
	stateKey := fmt.Sprintf("%d,%d,%d", pos.Row, pos.Col, splitterMask)
	if val, exists := memo[stateKey]; exists {
		return val
	}

	// Move down
	nextPos := types.NewPoint(pos.Row+1, pos.Col)

	var result int

	// Check if we're about to hit a splitter
	if idx, isSplitter := splitterIndex[nextPos]; isSplitter {
		// Check if we've already used this splitter in this timeline
		if (splitterMask & (1 << idx)) != 0 {
			// Already used - this is a cycle, return 0
			result = 0
		} else {
			// Mark this splitter as used and split into two timelines
			newMask := splitterMask | (1 << idx)
			leftPos := types.NewPoint(nextPos.Row, nextPos.Col-1)
			rightPos := types.NewPoint(nextPos.Row, nextPos.Col+1)

			leftTimelines := countTimelinesOptimized(leftPos, newMask, splitterList, splitterIndex, maxCol, maxRow, memo)
			rightTimelines := countTimelinesOptimized(rightPos, newMask, splitterList, splitterIndex, maxCol, maxRow, memo)

			result = leftTimelines + rightTimelines
		}
	} else {
		// No splitter, continue down
		result = countTimelinesOptimized(nextPos, splitterMask, splitterList, splitterIndex, maxCol, maxRow, memo)
	}

	memo[stateKey] = result
	return result
}

func runBeams(currentBeam types.Point, splitters map[types.Point]bool, maxCol, maxRow int, cache map[types.Point]bool) {
	// Check if the beam is out of bounds
	if currentBeam.Col < 0 || currentBeam.Row < 0 || currentBeam.Col >= maxCol || currentBeam.Row >= maxRow {
		return
	}
	// Skip if we've seen this position
	if cache[currentBeam] {
		return
	}
	cache[currentBeam] = true

	// Move the beam down
	newLoc := types.NewPoint(currentBeam.Row+1, currentBeam.Col)

	// If the beam hits a splitter, split it
	if _, ok := splitters[newLoc]; ok {
		leftBeam := types.NewPoint(newLoc.Row, newLoc.Col-1)
		rightBeam := types.NewPoint(newLoc.Row, newLoc.Col+1)
		runBeams(leftBeam, splitters, maxCol, maxRow, cache)
		runBeams(rightBeam, splitters, maxCol, maxRow, cache)
		splitters[newLoc] = true
		return
	}

	// Continue moving down
	runBeams(newLoc, splitters, maxCol, maxRow, cache)
}
