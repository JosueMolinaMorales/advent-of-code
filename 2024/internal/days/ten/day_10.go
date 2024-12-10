package ten

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay10() {
	fmt.Println("Day 10 Part 1: ", solve(false))
	fmt.Println("Day 10 Part 2: ", solve(true))
}

func setup() ([][]int, hashset.Set) {
	rawMap, err := util.LoadFileAsString("./inputs/day_10.txt")
	if err != nil {
		panic(rawMap)
	}

	topMap := [][]int{}
	startingPoints := hashset.New()
	for i, line := range strings.Split(rawMap, "\n") {
		row := []int{}
		for j, s := range strings.Split(line, "") {
			n := util.ToInt(s)
			if n == 0 {
				startingPoints.Add(*types.NewVector(i, j))
			}

			row = append(row, n)
		}
		topMap = append(topMap, row)
	}

	return topMap, *startingPoints
}

func solve(findAllPaths bool) int {
	topMap, startingPoints := setup()

	// Use DFS to find path from 0 - 9
	count := 0
	for _, point := range startingPoints.Values() {
		p := point.(types.Vector)
		count += search(p, findAllPaths, *hashset.New(), topMap)
	}
	return count
}

func search(curr_pos types.Vector, findAllPaths bool, visited hashset.Set, topMap [][]int) int {
	// Mark current position on visited
	visited.Add(curr_pos)

	// Check to see if the curr_point is 9
	if topMap[curr_pos.X][curr_pos.Y] == 9 {
		return 1
	}

	// Find all adjacent points
	found := 0
	for _, adjacent := range adj(curr_pos, topMap) {
		// Dont continue if we have visited unless we are attempting to find all paths
		if !visited.Contains(adjacent) || findAllPaths {
			found += search(adjacent, findAllPaths, visited, topMap)
		}
	}

	return found
}

func adj(curr_pos types.Vector, topMap [][]int) []types.Vector {
	directions := [][]int{
		{0, 1},  // Right
		{0, -1}, // Left
		{-1, 0}, // Down
		{1, 0},  // Up
	}

	adjacents := []types.Vector{}
	for _, dir := range directions {
		dx, dy := curr_pos.X+dir[0], curr_pos.Y+dir[1]
		if dx < 0 || dx >= len(topMap) || dy < 0 || dy >= len(topMap[0]) {
			continue
		}

		np := topMap[dx][dy]
		if (topMap[curr_pos.X][curr_pos.Y] + 1) == np {
			adjacents = append(adjacents, *types.NewVector(dx, dy))
		}
	}
	return adjacents
}
