package twentyfive

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay25() {
	fmt.Println("Day 25 Part 1: ", solvePartOne())
}

func solvePartOne() int {
	input, err := util.LoadFileAsString("./inputs/day_25.txt")
	if err != nil {
		panic(err)
	}
	keys := [][]int{}
	locks := [][]int{}

	for _, section := range strings.Split(input, "\n\n") {
		grid := [][]string{}
		for _, line := range strings.Split(section, "\n") {
			grid = append(grid, strings.Split(line, ""))
		}
		columns := countColumns(grid)
		if grid[0][0] == "#" {
			// Lock
			locks = append(locks, columns)
		} else {
			// Key
			keys = append(keys, columns)
		}
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			works := true
			for i := 0; i < len(key); i++ {
				if (key[i] + lock[i]) > 5 {
					works = false
				}
			}
			if works {
				count++
			}
		}
	}
	return count
}

func countColumns(grid [][]string) []int {
	cols := len(grid[0])
	columns := make([]int, cols)
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == "#" {
				columns[i]++
			}
		}
		columns[i]--
	}

	return columns
}
