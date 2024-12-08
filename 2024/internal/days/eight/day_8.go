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

	grid := make([][]string, 0)
	frequencies := make(map[string][]types.Vector)
	for i, line := range strings.Split(rawMap, "\n") {
		row := make([]string, 0)
		for j, col := range strings.Split(line, "") {
			if col != "." {
				if frequencies[col] == nil {
					frequencies[col] = make([]types.Vector, 0)
				}
				frequencies[col] = append(frequencies[col], *types.NewVector(i, j))
			}
			row = append(row, col)
		}
		grid = append(grid, row)
	}

	return grid, frequencies
}

func isInBounds(dx, dy, xBound, yBound int) bool {
	return dx >= 0 && dx < xBound && dy >= 0 && dy < yBound
}

func solve(part2 bool) int {
	grid, freqs := setup()
	antinodes := hashset.New()
	for _, v := range freqs {
		for _, point := range v {
			for _, otherPoint := range v {
				if point == otherPoint {
					continue
				}

				dx, dy := point.X-otherPoint.X, point.Y-otherPoint.Y

				if part2 {
					x, y := point.X, point.Y
					for {
						if !isInBounds(x, y, len(grid), len(grid[0])) {
							break
						}
						antinodes.Add(*types.NewVector(x, y))
						x, y = x+dx, y+dy
					}
				} else {
					x, y := point.X+dx, point.Y+dy
					if isInBounds(x, y, len(grid), len(grid[0])) {
						antinodes.Add(*types.NewVector(x, y))
					}
				}
			}
		}
	}
	return antinodes.Size()
}
