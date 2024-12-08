package eight

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay8() {
	fmt.Println("Day 8 Part 1: ", solve(false))
	fmt.Println("Day 8 Part 2: ", solve(true))
}

func setup() ([][]string, map[string][]types.Vector) {
	rawMap, err := util.LoadFileAsString("./inputs/day_8.txt")
	if err != nil {
		panic(err)
	}

	grid := [][]string{}
	frequencies := make(map[string][]types.Vector)

	for i, line := range strings.Split(rawMap, "\n") {
		row := strings.Split(line, "")
		grid = append(grid, row)

		for j, col := range row {
			if col == "." {
				continue
			}
			frequencies[col] = append(frequencies[col], *types.NewVector(i, j))
		}
	}
	return grid, frequencies
}

func isInBounds(x, y, xBound, yBound int) bool {
	return x >= 0 && x < xBound && y >= 0 && y < yBound
}

func solve(part2 bool) int {
	grid, freqs := setup()
	antinodes := hashset.New()

	for _, points := range freqs {
		for _, point := range points {
			for _, otherPoint := range points {
				if point == otherPoint {
					continue
				}

				dx, dy := point.X-otherPoint.X, point.Y-otherPoint.Y
				x, y := point.X, point.Y

				if part2 {
					// Follow the vector until out of bounds
					for isInBounds(x, y, len(grid), len(grid[0])) {
						antinodes.Add(*types.NewVector(x, y))
						x, y = x+dx, y+dy
					}
				} else {
					// Single step in the direction of the vector
					x, y = x+dx, y+dy
					if isInBounds(x, y, len(grid), len(grid[0])) {
						antinodes.Add(*types.NewVector(x, y))
					}
				}
			}
		}
	}
	return antinodes.Size()
}
