package four

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay4() {
	res := solvePartOne()
	fmt.Println("Day 4 Part 1: ", res)
	res = solvePartTwo()
	fmt.Println("Day 4 Part 2: ", res)
}

func setup() [][]string {
	rawGrid, err := util.LoadFileAsString("./inputs/day_4.txt")
	if err != nil {
		panic(err)
	}
	crossword := make([][]string, 0)
	for _, row := range strings.Split(rawGrid, "\n") {
		crossword = append(crossword, strings.Split(row, ""))
	}

	return crossword
}

func solvePartOne() int {
	crossword := setup()
	xmasCount := 0
	directions := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	for i, row := range crossword {
		for j, col := range row {
			if col != "X" {
				continue
			}
			for _, dir := range directions {
				found := true
				for k := 1; k < 4; k++ {
					dx, dy := i+(dir[0]*k), j+(dir[1]*k)
					if isOutOfBounds(crossword, dx, dy) {
						found = false
						break
					}
					if crossword[dx][dy] != string("XMAS"[k]) {
						found = false
						break
					}
				}

				if found {
					xmasCount++
				}
			}

		}
	}
	return xmasCount
}

func solvePartTwo() int {
	cw := setup()
	xmasCount := 0

	for i, row := range cw {
		for j, col := range row {
			if col != "A" {
				continue
			}

			dx_1, dy_1 := i-1, j-1 // Top left
			dx_2, dy_2 := i+1, j+1 // Bottom Right
			dx_3, dy_3 := i-1, j+1 // Top Right
			dx_4, dy_4 := i+1, j-1 // Bottom Left
			if isOutOfBounds(cw, dx_1, dy_1) || isOutOfBounds(cw, dx_2, dy_2) || isOutOfBounds(cw, dx_3, dy_3) || isOutOfBounds(cw, dx_4, dy_4) {
				continue
			}
			if (cw[dx_1][dy_1] == "M" && cw[dx_2][dy_2] == "S" || cw[dx_1][dy_1] == "S" && cw[dx_2][dy_2] == "M") &&
				(cw[dx_3][dy_3] == "M" && cw[dx_4][dy_4] == "S" || cw[dx_3][dy_3] == "S" && cw[dx_4][dy_4] == "M") {
				xmasCount += 1
			}
		}
	}
	return xmasCount
}

func isOutOfBounds(cw [][]string, dx, dy int) bool {
	return dx < 0 || dx >= len(cw) || dy < 0 || dy >= len(cw[dx])
}
